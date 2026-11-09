package main

// FIGlet's "small" font by Glenn Chappell (BSD, from the figlet distribution), with kerning: each letter
// slides left until it would touch the one before, the way `figlet -k` sets it.

import (
	_ "embed"
	"strings"
)

//go:embed fonts/small.flf
var smallFLF string

var small = parseFLF(smallFLF)

type flf struct {
	h      int
	glyphs map[rune][]string
}

func parseFLF(src string) *flf {
	lines := strings.Split(strings.ReplaceAll(src, "\r", ""), "\n")
	hdr := strings.Fields(lines[0])
	hard := hdr[0][len(hdr[0])-1:]
	f := &flf{h: atoi(hdr[1]), glyphs: map[rune][]string{}}
	body := lines[1+atoi(hdr[5]):]
	i := 0
	for r := rune(32); r < 127 && i+f.h <= len(body); r++ {
		rows := make([]string, f.h)
		for k := range rows {
			row := body[i+k]
			end := row[len(row)-1:]
			rows[k] = strings.ReplaceAll(strings.TrimRight(row, end), hard, " ")
		}
		f.glyphs[r] = rows
		i += f.h
	}
	return f
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			break
		}
		n = n*10 + int(c-'0')
	}
	return n
}

// Render sets s as rows of text, kerned
func (f *flf) Render(s string) []string {
	out := make([][]rune, f.h)
	for _, ch := range s {
		g, ok := f.glyphs[ch]
		if !ok {
			g = f.glyphs[' ']
		}
		gw := 0
		gl := make([][]rune, f.h)
		for k, row := range g {
			gl[k] = []rune(row)
			gw = max(gw, len(gl[k]))
		}
		for k := range gl {
			for len(gl[k]) < gw {
				gl[k] = append(gl[k], ' ')
			}
		}
		ow := 0
		for k := range out {
			ow = max(ow, len(out[k]))
		}
		for k := range out {
			for len(out[k]) < ow {
				out[k] = append(out[k], ' ')
			}
		}
		// the furthest the glyph can slide left without two inked cells landing on each other
		shift := 0
		for s := 1; s <= min(gw, ow); s++ {
			clash := false
			for k := 0; k < f.h && !clash; k++ {
				for i := 0; i < s; i++ {
					if out[k][ow-s+i] != ' ' && gl[k][i] != ' ' {
						clash = true
						break
					}
				}
			}
			if clash {
				break
			}
			shift = s
		}
		for k := range out {
			for i := 0; i < gw; i++ {
				x := ow - shift + i
				if x < len(out[k]) {
					if gl[k][i] != ' ' {
						out[k][x] = gl[k][i]
					}
				} else {
					out[k] = append(out[k], gl[k][i])
				}
			}
		}
	}
	rows := make([]string, f.h)
	for k := range out {
		rows[k] = strings.TrimRight(string(out[k]), " ")
	}
	// drop rows that are blank at the bottom (small has a spare descender row)
	for len(rows) > 0 && strings.TrimSpace(rows[len(rows)-1]) == "" {
		rows = rows[:len(rows)-1]
	}
	return rows
}
