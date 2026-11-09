package main

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

// the site's palette (src/app.css), as accents on a terminal's dark
var (
	white  = hex("#ffffff")
	silk   = hex("#f7f7f2")
	yellow = hex("#efc419")
	gold   = hex("#d6a542")
	sky    = hex("#3b8ad6")
	// inner pages: a night-sky ground, body text, and the quieter text
	ground = hex("#08111f")
	text   = hex("#d5dce6")
	dim    = hex("#7d8aa0")
	// the landing panels
	panel  = hex("#050a12")
	border = hex("#5d6b82")
	hcRed  = hex("#ec3750") // Hack Club's red
	green  = hex("#46e07c") // the "working on it now" pad
)

// the solder masks, lifted so they read as text on the dark; black and white become greys
var maskText = map[string]RGB{
	"green": hex("#4fc27f"), "blue": hex("#6aa6f2"), "purple": hex("#b08ae8"), "red": hex("#f0675c"),
	"yellow": hex("#efc419"), "white": hex("#e8e8e0"), "black": hex("#9aa4b2"),
}

// each event's own accent, from src/lib/styles/themes.css
var eventAccent = map[string]RGB{
	"fraserhacks": hex("#e84bff"), "dayoneof": hex("#e8574c"), "offscript": hex("#00bfff"), "csclub": hex("#ef4256"),
	"polaris": hex("#8ef2d6"), "outpost": hex("#fe9802"), "fallout": hex("#c9977f"), "alleyway": hex("#f3bf10"),
	"undercity": hex("#ff4fdc"), "hackthenorth": hex("#e8562a"), "uottahack": hex("#ff8fd2"),
}

// Span is a run of styled text; a paragraph is a list of them
type Span struct {
	S    string
	A    Attr
	Fg   *RGB
	Link string
}

func plain(s string) []Span { return []Span{{S: s}} }

var (
	mdBold = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	mdLink = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	mdCode = regexp.MustCompile("`([^`]+)`")
	mdAny  = regexp.MustCompile("\\*\\*[^*]+\\*\\*|\\[[^\\]]+\\]\\([^)]+\\)|`[^`]+`")
)

// markdown handles the little the bodies use: **bold**, [links](url) and `code`
func markdown(s string, linkFg, codeFg RGB) []Span {
	s = strings.ReplaceAll(s, "\n", " ")
	var out []Span
	last := 0
	for _, m := range mdAny.FindAllStringIndex(s, -1) {
		if m[0] > last {
			out = append(out, Span{S: s[last:m[0]]})
		}
		tok := s[m[0]:m[1]]
		switch {
		case mdBold.MatchString(tok):
			out = append(out, Span{S: mdBold.FindStringSubmatch(tok)[1], A: Bold})
		case mdLink.MatchString(tok):
			g := mdLink.FindStringSubmatch(tok)
			out = append(out, Span{S: g[1], A: Underline, Fg: &linkFg, Link: g[2]})
		default:
			out = append(out, Span{S: mdCode.FindStringSubmatch(tok)[1], Fg: &codeFg})
		}
		last = m[1]
	}
	if last < len(s) {
		out = append(out, Span{S: s[last:]})
	}
	return out
}

// word is a piece of a line that can't be broken, with the space after it
type word struct {
	spans []Span
	space bool
}

func width(s string) int { return runewidth.StringWidth(s) }

// Wrap breaks spans into lines no wider than w
func Wrap(spans []Span, w int) [][]Span {
	w = max(w, 4)
	var words []word
	cur := word{}
	for _, sp := range spans {
		for i, part := range strings.Split(sp.S, " ") {
			if i > 0 {
				cur.space = true
				words = append(words, cur)
				cur = word{}
			}
			if part != "" {
				p := sp
				p.S = part
				cur.spans = append(cur.spans, p)
			}
		}
	}
	words = append(words, cur)

	var lines [][]Span
	var line []Span
	lw := 0
	wordW := func(wd word) int {
		n := 0
		for _, s := range wd.spans {
			n += width(s.S)
		}
		return n
	}
	for i, wd := range words {
		ww := wordW(wd)
		if ww == 0 {
			continue
		}
		// a word longer than the line gets cut
		for ww > w {
			if lw > 0 {
				lines, line, lw = append(lines, line), nil, 0
			}
			head, tail := splitSpans(wd.spans, w)
			lines = append(lines, head)
			wd.spans, ww = tail, ww-w
		}
		if lw > 0 && lw+1+ww > w {
			lines, line, lw = append(lines, line), nil, 0
		}
		if lw > 0 {
			// the space takes the style of the word before it, so links underline straight through
			prev := words[i-1].spans
			sp := Span{S: " "}
			if len(prev) > 0 && len(wd.spans) > 0 && prev[len(prev)-1].A&Underline != 0 && prev[len(prev)-1].A == wd.spans[0].A && prev[len(prev)-1].Link == wd.spans[0].Link {
				sp = prev[len(prev)-1]
				sp.S = " "
			}
			line = append(line, sp)
			lw++
		}
		line = append(line, wd.spans...)
		lw += ww
	}
	if lw > 0 || len(lines) == 0 {
		lines = append(lines, line)
	}
	return lines
}

func splitSpans(spans []Span, w int) (head, tail []Span) {
	n := 0
	for i, s := range spans {
		sw := width(s.S)
		if n+sw <= w {
			head = append(head, s)
			n += sw
			continue
		}
		cut := 0
		for k := range s.S {
			if n+width(s.S[:k]) > w {
				break
			}
			cut = k
		}
		a, b := s, s
		a.S, b.S = s.S[:cut], s.S[cut:]
		if a.S != "" {
			head = append(head, a)
		}
		tail = append([]Span{b}, spans[i+1:]...)
		return
	}
	return
}

// Line draws one wrapped line; spans without their own colour use fg
func (l *Layer) Line(x, y int, line []Span, fg RGB, base Attr) int {
	for _, sp := range line {
		c := fg
		if sp.Fg != nil {
			c = *sp.Fg
		}
		start := x
		x = l.Text(x, y, sp.S, c, base|sp.A)
		l.Linkify(start, y, x-start, sp.Link)
	}
	return x
}

// Para wraps and draws a paragraph, returning the row after it
func (l *Layer) Para(x, y, w int, spans []Span, fg RGB, base Attr) int {
	for _, line := range Wrap(spans, w) {
		l.Line(x, y, line, fg, base)
		y++
	}
	return y
}

func paraHeight(spans []Span, w int) int { return len(Wrap(spans, w)) }

func lineWidth(line []Span) int {
	n := 0
	for _, s := range line {
		n += width(s.S)
	}
	return n
}

// fit shortens s to w columns with an ellipsis
func fit(s string, w int) string {
	if width(s) <= w {
		return s
	}
	if w <= 1 {
		return "…"[:min(len("…"), w*3)]
	}
	return runewidth.Truncate(s, w, "…")
}

func runeCount(s string) int { return utf8.RuneCountInString(s) }
