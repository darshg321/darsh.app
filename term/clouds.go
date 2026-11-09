package main

// The landing's clouds drift one by one. The photo is pulled apart once into a clear sky (a smooth
// gradient fitted to the blue around the clouds) and the clouds themselves (how much each pixel is
// cloud rather than sky, split into separate puffs); every frame lays each puff back over the sky,
// blown right at its own speed and wrapping round when it leaves the screen.

import (
	"math"
	"sync"
)

// a puff is one cloud, cut out of the photo: alpha and colour over its bounding box
type puff struct {
	x0, y0, w, h int
	a            []float64
	c            [][3]float64
}

type cloudPic struct {
	p     *pic
	bg    [3][6]float64 // per channel, a quadratic in x and y over the photo
	puffs []puff
}

var (
	cloudOnce sync.Once
	cloudData *cloudPic
)

func quad(x, y float64) [6]float64 { return [6]float64{1, x, y, x * x, x * y, y * y} }

func (cp *cloudPic) sky(x, y float64) [3]float64 {
	q := quad(x/float64(cp.p.W), y/float64(cp.p.H))
	var c [3]float64
	for ch := range 3 {
		for i, v := range q {
			c[ch] += cp.bg[ch][i] * v
		}
	}
	return c
}

func loadClouds() *cloudPic {
	cloudOnce.Do(func() {
		p := loadPic("/img/sky.jpg")
		if p == nil {
			return
		}
		cp := &cloudPic{p: p}
		W, H := p.W, p.H
		px := func(x, y int) [3]float64 {
			o := (y*W + x) * 3
			return [3]float64{float64(p.px[o]), float64(p.px[o+1]), float64(p.px[o+2])}
		}
		sat := func(c [3]float64) float64 {
			mx, mn := math.Max(c[0], math.Max(c[1], c[2])), math.Min(c[0], math.Min(c[1], c[2]))
			if mx == 0 {
				return 0
			}
			return (mx - mn) / mx
		}

		// the sky: least squares over the clearly blue pixels
		var ata [6][6]float64
		var atb [3][6]float64
		for y := 0; y < H; y += 2 {
			for x := 0; x < W; x += 2 {
				c := px(x, y)
				if sat(c) < 0.5 {
					continue
				}
				q := quad(float64(x)/float64(W), float64(y)/float64(H))
				for i := range 6 {
					for j := range 6 {
						ata[i][j] += q[i] * q[j]
					}
					for ch := range 3 {
						atb[ch][i] += q[i] * c[ch]
					}
				}
			}
		}
		for ch := range 3 {
			cp.bg[ch] = solve6(ata, atb[ch])
		}

		// how cloudy each pixel is: how far it's gone from the sky towards white
		white := [3]float64{245, 247, 250}
		alpha := make([]float64, W*H)
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				c, s := px(x, y), cp.sky(float64(x), float64(y))
				var num, den float64
				for ch := range 3 {
					num += (c[ch] - s[ch]) * (white[ch] - s[ch])
					den += (white[ch] - s[ch]) * (white[ch] - s[ch])
				}
				a := math.Max(0, math.Min(1, num/den))
				alpha[y*W+x] = smooth(0.05, 0.55, a)
			}
		}

		// the puffs: everything touching, flood filled
		label := make([]int, W*H)
		for i := range label {
			label[i] = -1
		}
		var stack []int
		for start := range alpha {
			if alpha[start] < 0.03 || label[start] >= 0 {
				continue
			}
			id := len(cp.puffs)
			x0, y0, x1, y1 := W, H, 0, 0
			var area float64
			label[start] = id
			stack = append(stack[:0], start)
			var members []int
			for len(stack) > 0 {
				i := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				members = append(members, i)
				area += alpha[i]
				x, y := i%W, i/W
				x0, y0, x1, y1 = min(x0, x), min(y0, y), max(x1, x), max(y1, y)
				for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
					nx, ny := x+d[0], y+d[1]
					if nx < 0 || ny < 0 || nx >= W || ny >= H {
						continue
					}
					if j := ny*W + nx; label[j] < 0 && alpha[j] >= 0.03 {
						label[j] = id
						stack = append(stack, j)
					}
				}
			}
			if area < 6 {
				continue // a speck; it stays in the sky as nothing
			}
			pf := puff{x0: x0, y0: y0, w: x1 - x0 + 1, h: y1 - y0 + 1}
			pf.a = make([]float64, pf.w*pf.h)
			pf.c = make([][3]float64, pf.w*pf.h)
			// a puff the photo cut off at the side would show a hard edge once it moves, so it fades out there
			fade := float64(W) / 5
			for _, i := range members {
				x, y := i%W, i/W
				a := alpha[i]
				if x0 == 0 {
					a *= math.Pow(smooth(0, fade, float64(x)), 1.5)
				}
				if x1 == W-1 {
					a *= math.Pow(smooth(0, fade, float64(W-1-x)), 1.5)
				}
				k := (y-y0)*pf.w + x - x0
				pf.a[k], pf.c[k] = a, px(x, y)
			}
			cp.puffs = append(cp.puffs, pf)
		}
		cloudData = cp
	})
	return cloudData
}

func smooth(e0, e1, x float64) float64 {
	t := math.Max(0, math.Min(1, (x-e0)/(e1-e0)))
	return t * t * (3 - 2*t)
}

// solve6 is Gaussian elimination for the sky fit
func solve6(a [6][6]float64, b [6]float64) [6]float64 {
	for c := range 6 {
		p := c
		for r := c + 1; r < 6; r++ {
			if math.Abs(a[r][c]) > math.Abs(a[p][c]) {
				p = r
			}
		}
		a[c], a[p] = a[p], a[c]
		b[c], b[p] = b[p], b[c]
		for r := c + 1; r < 6; r++ {
			f := a[r][c] / a[c][c]
			for k := c; k < 6; k++ {
				a[r][k] -= f * a[c][k]
			}
			b[r] -= f * b[c]
		}
	}
	var x [6]float64
	for r := 5; r >= 0; r-- {
		s := b[r]
		for k := r + 1; k < 6; k++ {
			s -= a[r][k] * x[k]
		}
		x[r] = s / a[r][r]
	}
	return x
}

// a sprite is a puff at a screen's scale: premultiplied colour and alpha per cell
type sprite struct {
	x, y, w, h int // where it starts, in cells
	a          []float64
	c          [][3]float64
	speed      float64 // cells a second
}

type scene struct {
	sky     [][3]float64
	sprites []sprite
}

var scenes = newCache[skyKey, *scene](32)

func cloudScene(w, h int) *scene {
	return scenes.get(skyKey{w, h}, func() *scene {
		cp := loadClouds()
		if cp == nil {
			return nil
		}
		p := cp.p
		// the same crop as Sky
		s := min(float64(p.W)/float64(w), float64(p.H)/float64(h)/2)
		sx, sy := s, s*2
		ox := (float64(p.W) - sx*float64(w)) / 2
		oy := (float64(p.H) - sy*float64(h)) * 0.35
		sc := &scene{sky: make([][3]float64, w*h)}
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				sc.sky[y*w+x] = cp.sky(ox+(float64(x)+0.5)*sx, oy+(float64(y)+0.5)*sy)
			}
		}
		for idx, pf := range cp.puffs {
			cx0 := int(math.Floor((float64(pf.x0) - ox) / sx))
			cx1 := int(math.Floor((float64(pf.x0+pf.w-1) - ox) / sx))
			cy0 := max(0, int(math.Floor((float64(pf.y0)-oy)/sy)))
			cy1 := min(h-1, int(math.Floor((float64(pf.y0+pf.h-1)-oy)/sy)))
			if cy1 < cy0 {
				continue
			}
			sp := sprite{x: cx0, y: cy0, w: cx1 - cx0 + 1, h: cy1 - cy0 + 1}
			sp.a = make([]float64, sp.w*sp.h)
			sp.c = make([][3]float64, sp.w*sp.h)
			for j := 0; j < sp.h; j++ {
				for i := 0; i < sp.w; i++ {
					X0, Y0 := ox+float64(cx0+i)*sx, oy+float64(cy0+j)*sy
					var a, n float64
					var c [3]float64
					for y := int(Y0); y < int(Y0+sy+0.5); y++ {
						for x := int(X0); x < int(X0+sx+0.5); x++ {
							n++
							u, v := x-pf.x0, y-pf.y0
							if u < 0 || v < 0 || u >= pf.w || v >= pf.h {
								continue
							}
							k := v*pf.w + u
							a += pf.a[k]
							for ch := range 3 {
								c[ch] += pf.a[k] * pf.c[k][ch]
							}
						}
					}
					if n > 0 {
						sp.a[j*sp.w+i] = a / n
						for ch := range 3 {
							c[ch] /= n
						}
						sp.c[j*sp.w+i] = c
					}
				}
			}
			// lower clouds are nearer, so they go faster; a little jitter so neighbours separate
			depth := (float64(pf.y0) + float64(pf.h)/2) / float64(p.H)
			jitter := float64((idx*37)%11) / 10
			sp.speed = (0.35 + 0.9*depth + 0.35*jitter) * float64(w) / 120
			sc.sprites = append(sc.sprites, sp)
		}
		return sc
	})
}

// SkyAt is the backdrop t seconds in: each cloud blown right at its own pace over a still sky
func SkyAt(w, h int, t float64) *Layer {
	sc := cloudScene(w, h)
	if sc == nil {
		return Sky(w, h)
	}
	buf := make([][3]float64, w*h)
	copy(buf, sc.sky)
	for _, sp := range sc.sprites {
		// it leaves on the right, waits a moment off screen and comes back in on the left
		period := float64(w + sp.w + 12)
		pos := math.Mod(float64(sp.x+sp.w)+sp.speed*t, period)
		pos -= float64(sp.w)
		for x := max(0, int(math.Floor(pos))); x < min(w, int(math.Ceil(pos))+sp.w); x++ {
			u := float64(x) - pos // where in the sprite this column falls
			i := int(math.Floor(u))
			f := u - float64(i)
			for j := 0; j < sp.h; j++ {
				y := sp.y + j
				if y >= h {
					break
				}
				// between two of the sprite's columns, for motion smoother than a cell at a time
				var a float64
				var c [3]float64
				for _, k := range [2]struct {
					i  int
					wt float64
				}{{i, 1 - f}, {i + 1, f}} {
					if k.i < 0 || k.i >= sp.w || k.wt == 0 {
						continue
					}
					a += k.wt * sp.a[j*sp.w+k.i]
					for ch := range 3 {
						c[ch] += k.wt * sp.c[j*sp.w+k.i][ch]
					}
				}
				o := &buf[y*w+x]
				for ch := range 3 {
					o[ch] = o[ch]*(1-a) + c[ch]
				}
			}
		}
	}
	l := NewLayer(w, h)
	u := func(v float64) uint8 { return uint8(math.Max(0, math.Min(255, v+0.5))) }
	for i, c := range buf {
		l.C[i] = skyCell(RGB{u(c[0]), u(c[1]), u(c[2])})
	}
	return l
}
