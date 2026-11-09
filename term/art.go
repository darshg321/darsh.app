package main

// The landing page's backdrop is the site's sky photo as coloured ASCII, the way term.samliu.dev shows its
// webcam: each cell takes the colour of its patch of the photo, and a character dense enough for how
// bright it is, on a ground that's the same colour pushed most of the way to night.

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	_ "image/jpeg"
	"math"
	"sync"
)

//go:embed assets
var assets embed.FS

type cache[K comparable, V any] struct {
	mu  sync.Mutex
	m   map[K]V
	cap int
}

func newCache[K comparable, V any](n int) *cache[K, V] { return &cache[K, V]{m: map[K]V{}, cap: n} }

func (c *cache[K, V]) get(k K, make func() V) V {
	c.mu.Lock()
	if v, ok := c.m[k]; ok {
		c.mu.Unlock()
		return v
	}
	c.mu.Unlock()
	v := make()
	c.mu.Lock()
	if len(c.m) >= c.cap {
		c.m = map[K]V{}
	}
	c.m[k] = v
	c.mu.Unlock()
	return v
}

// pic is the decoded photo, shrunk to at most 480px on its long side
type pic struct {
	W, H int
	px   []uint8 // rgb
}

var pics = newCache[string, *pic](4)

func loadPic(path string) *pic {
	return pics.get(path, func() *pic {
		b, err := assets.ReadFile("assets" + path)
		if err != nil {
			return nil
		}
		img, _, err := image.Decode(bytes.NewReader(b))
		if err != nil {
			return nil
		}
		r := img.Bounds()
		scale := max(1, (max(r.Dx(), r.Dy())+479)/480)
		p := &pic{W: r.Dx() / scale, H: r.Dy() / scale}
		p.px = make([]uint8, p.W*p.H*3)
		for y := 0; y < p.H; y++ {
			for x := 0; x < p.W; x++ {
				var s [3]int
				for j := 0; j < scale; j++ {
					for i := 0; i < scale; i++ {
						c := color.RGBAModel.Convert(img.At(r.Min.X+x*scale+i, r.Min.Y+y*scale+j)).(color.RGBA)
						s[0] += int(c.R)
						s[1] += int(c.G)
						s[2] += int(c.B)
					}
				}
				n := scale * scale
				o := (y*p.W + x) * 3
				p.px[o], p.px[o+1], p.px[o+2] = uint8(s[0]/n), uint8(s[1]/n), uint8(s[2]/n)
			}
		}
		return p
	})
}

// average colour of a patch, in photo pixels
func (p *pic) patch(x0, y0, x1, y1 float64) RGB {
	ix0, iy0 := max(0, int(x0)), max(0, int(y0))
	ix1, iy1 := min(p.W, max(ix0+1, int(x1+0.5))), min(p.H, max(iy0+1, int(y1+0.5)))
	var r, g, b, n int
	for y := iy0; y < iy1; y++ {
		for x := ix0; x < ix1; x++ {
			o := (y*p.W + x) * 3
			r += int(p.px[o])
			g += int(p.px[o+1])
			b += int(p.px[o+2])
			n++
		}
	}
	if n == 0 {
		return RGB{}
	}
	return RGB{uint8(r / n), uint8(g / n), uint8(b / n)}
}

// darkest to brightest, the same kind of ramp as the reference: thin strokes in the blue, letters in the clouds
var skyRamp = []rune(".:;!|ilzrvuUQ$")

var night = hex("#050b16")

type skyKey struct{ w, h int }

var skies = newCache[skyKey, *Layer](32)

// Sky is the backdrop at w×h, cover-fitted and cropped a little above centre
func Sky(w, h int) *Layer {
	return skies.get(skyKey{w, h}, func() *Layer {
		l := NewLayer(w, h)
		p := loadPic("/img/sky.jpg")
		if p == nil {
			l.Fill(0, 0, w, h, night)
			return l
		}
		// a cell is about twice as tall as it is wide
		s := min(float64(p.W)/float64(w), float64(p.H)/float64(h)/2)
		sx, sy := s, s*2
		ox := (float64(p.W) - sx*float64(w)) / 2
		oy := (float64(p.H) - sy*float64(h)) * 0.35
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				x0, y0 := ox+float64(x)*sx, oy+float64(y)*sy
				c := posterise(p.patch(x0, y0, x0+sx, y0+sy), 12)
				lum := c.lum()
				// stretch the contrast so the blue gets thin strokes and the clouds get letters
				t := math.Max(0, math.Min(0.999, (lum-0.12)/0.78))
				r := skyRamp[int(t*float64(len(skyRamp)))]
				l.C[y*w+x] = Cell{K: Solid, R: r, Fg: mix(c, white, 0.15+0.35*t), Bg: mix(c, night, 0.52)}
			}
		}
		return l
	})
}

// posterise steps brightness and saturation in HSV but keeps the hue, so neighbouring cells share
// colours (less ANSI per frame) without the sky turning teal in places; q is the brightness steps
func posterise(c RGB, q int) RGB {
	r, g, b := float64(c.R)/255, float64(c.G)/255, float64(c.B)/255
	mx, mn := math.Max(r, math.Max(g, b)), math.Min(r, math.Min(g, b))
	v := mx
	s := 0.0
	if mx > 0 {
		s = (mx - mn) / mx
	}
	h := 0.0
	if d := mx - mn; d > 0 {
		switch mx {
		case r:
			h = math.Mod((g-b)/d, 6)
		case g:
			h = (b-r)/d + 2
		default:
			h = (r-g)/d + 4
		}
	}
	h = math.Round(h*4) / 4
	s = math.Round(s*8) / 8
	v = math.Round(v*float64(q)) / float64(q)
	cc := v * s
	xx := cc * (1 - math.Abs(math.Mod(h, 2)-1))
	var rr, gg, bb float64
	switch {
	case h < 1:
		rr, gg = cc, xx
	case h < 2:
		rr, gg = xx, cc
	case h < 3:
		gg, bb = cc, xx
	case h < 4:
		gg, bb = xx, cc
	case h < 5:
		rr, bb = xx, cc
	default:
		rr, bb = cc, xx
	}
	m := v - cc
	u := func(f float64) uint8 { return uint8(math.Max(0, math.Min(255, (f+m)*255+0.5))) }
	return RGB{u(rr), u(gg), u(bb)}
}
