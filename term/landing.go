package main

import (
	"strings"
	"time"
	_ "time/tzdata"
)

// the landing nav, in order; each opens an inner page
var navItems = []string{"about", "projects", "events", "contact"}

// what i'm up to, from the site's home page; it sits beside the name
type status struct {
	label, what string
	col         RGB
}

var statuses = []status{
	{"on the bench", "a chip for hack club's tapeout", maskText["green"]},
	{"running", "dayoneof, with hack club", eventAccent["dayoneof"]},
	{"just finished", "bing bang dream adventure", maskText["purple"]},
}

var toronto, _ = time.LoadLocation("America/Toronto")

func (m *model) landing() *Layer {
	l := NewLayer(m.w, m.h)
	copy(l.C, Sky(m.w, m.h).C)

	name := small.Render("darsh")
	nameW := 0
	for _, r := range name {
		nameW = max(nameW, width(r))
	}
	now := time.Now().In(toronto)
	where := "mississauga, on · " + strings.ToLower(now.Format("3:04 pm MST"))
	labelW, infoW := 0, width(where)
	for _, st := range statuses {
		labelW = max(labelW, width(st.label))
	}
	for _, st := range statuses {
		infoW = max(infoW, labelW+2+width(st.what))
	}
	tag := []Span{{S: "making hardware and running hackathons @ "}, {S: "hack club", Fg: &hcRed}}
	email := m.c.Contact.Email

	// the tabs: the picked one is a yellow key with an arrow saying it opens downwards
	tabW := 0
	for i, it := range navItems {
		tabW += width(it) + 2
		if i == m.navSel {
			tabW += 2
		}
	}
	tabW += len(navItems) - 1

	inner := max(nameW+5+infoW, lineWidth(tag), tabW+6+width(email))
	beside := inner <= m.w-8 // the info beside the name, or under it on a narrow screen
	if !beside {
		inner = min(m.w-8, max(nameW, infoW, lineWidth(tag), tabW))
	}
	tagLines := Wrap(tag, inner)
	split := tabW+6+width(email) > inner
	infoRows := 1 + 1 + len(statuses)
	top := len(name)
	if beside {
		top = max(top, infoRows)
	} else {
		top += 1 + infoRows
	}
	h := 2 + top + 1 + len(tagLines) + 1 + 1 + 1 + 1
	if split {
		h++
	}
	w := inner + 6
	x := max(1, min(m.w/24, m.w-w-1))
	y := max(1, m.h-h-max(1, m.h/7))
	l.Fill(x, y, w, h, panel)
	box(l, x, y, w, h, border)

	cy := y + 2
	for i, r := range name {
		l.Text(x+3, cy+i, r, white, 0)
	}
	ix, iy := x+3+nameW+5, cy
	if !beside {
		ix, iy = x+3, cy+len(name)+1
	}
	l.Text(ix, iy, where, dim, 0)
	for i, st := range statuses {
		l.Text(ix, iy+2+i, st.label, dim, 0)
		l.Text(ix+labelW+2, iy+2+i, fit(st.what, inner-(ix-x-3)-labelW-2), st.col, 0)
	}
	cy += top + 1
	for _, ln := range tagLines {
		l.Line(x+3, cy, ln, text, 0)
		cy++
	}
	cy++
	nx := x + 3
	for i, it := range navItems {
		if i > 0 {
			nx = l.Text(nx, cy, "|", border, 0)
		}
		s := " " + it + " "
		if i == m.navSel {
			s = " " + it + " ↓ "
			l.Fill(nx, cy, width(s), 1, yellow)
			l.Text(nx, cy, s, panel, Bold)
		} else {
			l.Text(nx, cy, s, text, 0)
		}
		l.Hit(nx, cy, width(s), 1, "page:"+itoa(i+1))
		nx += width(s)
	}
	ex, ey := x+w-3-width(email), cy
	if split {
		ex, ey = x+3, cy+1
	}
	end := l.Text(ex, ey, email, text, 0)
	l.Linkify(ex, ey, end-ex, "mailto:"+email)

	// bottom right: how to get around
	hint := "←→ pick a tab · ↓ or enter opens it · ? all keys"
	if width(hint)+8 > m.w {
		hint = "←→ pick · ↓ open · ? keys"
	}
	if m.w >= 34 {
		hw := width(hint) + 4
		hx, hy := m.w-hw-2, min(m.h-2, y+h-2)
		if hx < x+w+1 {
			hy = min(m.h-2, y+h+1)
		}
		if hy < m.h-1 {
			l.Fill(hx, hy-1, hw, 3, panel)
			l.Text(hx+2, hy, hint, text, 0)
		}
	}
	return l
}

// box draws a +---+ border
func box(l *Layer, x, y, w, h int, c RGB) {
	for i := x + 1; i < x+w-1; i++ {
		l.Put(i, y, '-', c, 0)
		l.Put(i, y+h-1, '-', c, 0)
	}
	for j := y + 1; j < y+h-1; j++ {
		l.Put(x, j, '|', c, 0)
		l.Put(x+w-1, j, '|', c, 0)
	}
	for _, p := range [][2]int{{x, y}, {x + w - 1, y}, {x, y + h - 1}, {x + w - 1, y + h - 1}} {
		l.Put(p[0], p[1], '+', c, 0)
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	neg := n < 0
	if neg {
		n = -n
	}
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	if neg {
		s = "-" + s
	}
	return s
}
