package main

// Screens are drawn onto a Layer of cells (a character, two colours, attributes, maybe a link), then
// written out as ANSI that only changes what changes from one cell to the next.

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RGB struct{ R, G, B uint8 }

func hex(s string) RGB {
	v, _ := strconv.ParseUint(strings.TrimPrefix(s, "#"), 16, 32)
	return RGB{uint8(v >> 16), uint8(v >> 8), uint8(v)}
}

// mix goes from a (t=0) to b (t=1)
func mix(a, b RGB, t float64) RGB {
	t = math.Max(0, math.Min(1, t))
	f := func(x, y uint8) uint8 { return uint8(float64(x)*(1-t) + float64(y)*t + 0.5) }
	return RGB{f(a.R, b.R), f(a.G, b.G), f(a.B, b.B)}
}

func (c RGB) lum() float64 {
	return (0.2126*float64(c.R) + 0.7152*float64(c.G) + 0.0722*float64(c.B)) / 255
}

type Kind uint8

const (
	Empty Kind = iota // shows what's under it
	Solid
)

type Attr uint8

const (
	Bold Attr = 1 << iota
	Faint
	Italic
	Underline
)

type Cell struct {
	R      rune
	Fg, Bg RGB
	A      Attr
	K      Kind
	Link   uint16 // 1-based index into Layer.links
}

// Hit is a clickable area and what clicking it does
type Hit struct {
	X, Y, W, H int
	Action     string
}

type Layer struct {
	W, H  int
	C     []Cell
	links []string
	Hits  []Hit
}

func NewLayer(w, h int) *Layer { return &Layer{W: w, H: h, C: make([]Cell, w*h)} }

// grow makes room down to row y; a page doesn't know how long it is until it's drawn
func (l *Layer) grow(y int) {
	if y < l.H {
		return
	}
	h := max(y+1, l.H*2)
	c := make([]Cell, l.W*h)
	copy(c, l.C)
	l.C, l.H = c, h
}

func (l *Layer) at(x, y int) *Cell {
	if x < 0 || x >= l.W || y < 0 {
		return nil
	}
	l.grow(y)
	return &l.C[y*l.W+x]
}

// Trim cuts the layer to h rows
func (l *Layer) Trim(h int) {
	l.grow(h)
	l.H = h
	l.C = l.C[:l.W*h]
}

func (l *Layer) Hit(x, y, w, h int, action string) { l.Hits = append(l.Hits, Hit{x, y, w, h, action}) }

func (l *Layer) Fill(x, y, w, h int, bg RGB) {
	for j := y; j < y+h; j++ {
		for i := x; i < x+w; i++ {
			if c := l.at(i, j); c != nil {
				*c = Cell{K: Solid, R: ' ', Fg: bg, Bg: bg}
			}
		}
	}
}

// Put writes one character over whatever background is there
func (l *Layer) Put(x, y int, r rune, fg RGB, a Attr) {
	if c := l.at(x, y); c != nil {
		c.R, c.Fg, c.A, c.K = r, fg, a, Solid
	}
}

// Text writes s on one line and returns the column after it
func (l *Layer) Text(x, y int, s string, fg RGB, a Attr) int {
	for _, r := range s {
		l.Put(x, y, r, fg, a)
		x++
	}
	return x
}

// Linkify makes a run of cells a link, for terminals that do OSC 8
func (l *Layer) Linkify(x, y, w int, url string) {
	if url == "" {
		return
	}
	id := uint16(0)
	for i, u := range l.links {
		if u == url {
			id = uint16(i + 1)
		}
	}
	if id == 0 {
		l.links = append(l.links, url)
		id = uint16(len(l.links))
	}
	for i := x; i < x+w; i++ {
		if c := l.at(i, y); c != nil {
			c.Link = id
		}
	}
}

// Over lays top (scrolled down by off rows) over base; Empty cells let base show through
func Over(base, top *Layer, off int) *Layer {
	out := NewLayer(base.W, base.H)
	// both link tables, the top's numbered after the base's
	out.links = append(append([]string{}, base.links...), top.links...)
	shift := uint16(len(base.links))
	for y := 0; y < base.H; y++ {
		for x := 0; x < base.W; x++ {
			c := base.C[y*base.W+x]
			if ty := y + off; ty >= 0 && ty < top.H && x < top.W {
				if t := top.C[ty*top.W+x]; t.K != Empty {
					c = t
					if c.Link != 0 {
						c.Link += shift
					}
				}
			}
			out.C[y*base.W+x] = c
		}
	}
	return out
}

// ---------------- ANSI ----------------

type Profile int

const (
	TrueColor Profile = iota
	ANSI256
)

func sgrColor(b *strings.Builder, base int, c RGB, p Profile) {
	if p == TrueColor {
		fmt.Fprintf(b, ";%d;2;%d;%d;%d", base, c.R, c.G, c.B)
	} else {
		fmt.Fprintf(b, ";%d;5;%d", base, to256(c))
	}
}

// Render writes the screen as ANSI, sending only the attributes that change from cell to cell
func Render(l *Layer, p Profile) string {
	var b strings.Builder
	b.Grow(l.W * l.H * 6)
	for y := 0; y < l.H; y++ {
		var cur Cell
		var link uint16
		first := true
		for x := 0; x < l.W; x++ {
			c := l.C[y*l.W+x]
			if c.R == 0 {
				c.R = ' '
			}
			if c.Link != link {
				if c.Link == 0 {
					b.WriteString("\x1b]8;;\x1b\\")
				} else {
					b.WriteString("\x1b]8;;" + l.links[c.Link-1] + "\x1b\\")
				}
				link = c.Link
			}
			// a blank shows only its background, so it can keep whatever foreground came before
			if !first && c.R == ' ' && c.A&Underline == 0 && cur.A&Underline == 0 {
				c.Fg, c.A = cur.Fg, cur.A
			}
			switch {
			case first || c.A != cur.A:
				b.WriteString("\x1b[0")
				for _, a := range []struct {
					f Attr
					s string
				}{{Bold, ";1"}, {Faint, ";2"}, {Italic, ";3"}, {Underline, ";4"}} {
					if c.A&a.f != 0 {
						b.WriteString(a.s)
					}
				}
				sgrColor(&b, 38, c.Fg, p)
				sgrColor(&b, 48, c.Bg, p)
				b.WriteByte('m')
			case c.Fg != cur.Fg || c.Bg != cur.Bg:
				b.WriteString("\x1b[")
				var s strings.Builder
				if c.Fg != cur.Fg {
					sgrColor(&s, 38, c.Fg, p)
				}
				if c.Bg != cur.Bg {
					sgrColor(&s, 48, c.Bg, p)
				}
				b.WriteString(strings.TrimPrefix(s.String(), ";"))
				b.WriteByte('m')
			}
			cur, first = c, false
			b.WriteRune(c.R)
		}
		if link != 0 {
			b.WriteString("\x1b]8;;\x1b\\")
		}
		b.WriteString("\x1b[0m")
		if y < l.H-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}

// to256 picks the nearest xterm-256 colour, from the 6x6x6 cube or the grey ramp
func to256(c RGB) int {
	steps := []int{0, 95, 135, 175, 215, 255}
	near := func(v uint8) int {
		best, bi := 1<<30, 0
		for i, s := range steps {
			if d := (int(v) - s) * (int(v) - s); d < best {
				best, bi = d, i
			}
		}
		return bi
	}
	r, g, bl := near(c.R), near(c.G), near(c.B)
	cube := RGB{uint8(steps[r]), uint8(steps[g]), uint8(steps[bl])}
	grey := (int(c.R) + int(c.G) + int(c.B)) / 3
	gi := max(0, min(23, (grey-8+5)/10))
	gv := uint8(8 + gi*10)
	dist := func(a RGB) int {
		dr, dg, db := int(a.R)-int(c.R), int(a.G)-int(c.G), int(a.B)-int(c.B)
		return dr*dr + dg*dg + db*db
	}
	if dist(RGB{gv, gv, gv}) < dist(cube) {
		return 232 + gi
	}
	return 16 + 36*r + 6*g + bl
}
