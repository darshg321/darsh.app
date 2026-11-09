package main

import "strings"

// Inner pages read like a terminal document, after the reference: a nav bar of blocks, a heading
// underlined with =, plain wrapped text, a dashed rule and a ~ / breadcrumb. Colour is the accent:
// projects in their solder-mask colours, events in their own, yellow for whatever's picked.

const (
	pageLanding = iota
	pageAbout
	pageProjects
	pageEvents
	pageContact
)

// doc is the page being drawn: a cursor down a column of text
type doc struct {
	m    *model
	l    *Layer
	x, w int
	y    int
}

func (m *model) frameW() (x, w int) {
	w = min(m.w-8, 80)
	if m.w < 50 {
		w = m.w - 4
	}
	return max(2, min(4, (m.w-w)/2)), w
}

func (m *model) inner() *Layer {
	x, w := m.frameW()
	d := &doc{m: m, l: NewLayer(m.w, m.h), x: x, w: w, y: 1}
	// a strip of the landing's sky across the top, so the inner pages keep its blue
	strip := Sky(m.w, 5)
	copy(d.l.C, strip.C)
	d.y = 4
	d.navBar()
	var crumbs []string
	switch {
	case m.detail >= 0:
		p := &m.c.Projects[m.detail]
		d.project(p)
		crumbs = []string{"projects", strings.ToLower(p.Title)}
	case m.page == pageAbout:
		d.about()
		crumbs = []string{"about"}
	case m.page == pageProjects:
		d.projects()
		crumbs = []string{"projects"}
	case m.page == pageEvents:
		d.events()
		crumbs = []string{"events"}
	case m.page == pageContact:
		d.contact()
		crumbs = []string{"contact"}
	}
	d.footer(crumbs)
	d.l.Trim(max(d.y+2, m.h)) // room for the key bar pinned along the bottom
	return d.l
}

// navBar is the row of blocks; the page you're on drops out of the bar in yellow
func (d *doc) navBar() {
	items := append([]string{"~"}, navItems...)
	x := d.x
	for i, it := range items {
		s := "  " + it + "  "
		cur := i == d.m.page && d.m.detail < 0 || i == pageProjects && d.m.detail >= 0
		if cur {
			// the page you're on is a gap in the bar, on the night ground
			d.l.Fill(x, d.y, width(s), 1, ground)
			x = d.l.Text(x, d.y, s, yellow, Bold)
		} else {
			d.l.Fill(x, d.y, width(s), 1, sky)
			d.l.Text(x, d.y, s, white, Bold)
			x += width(s)
		}
		d.l.Hit(x-width(s), d.y, width(s), 1, "page:"+itoa(i))
	}
	d.y += 3
}

// h1 is the page heading, underlined with =
func (d *doc) h1(s string, under RGB) {
	d.l.Text(d.x, d.y, s, white, Bold)
	d.l.Text(d.x, d.y+1, strings.Repeat("=", width(s)), under, 0)
	d.y += 3
}

// h2 is a section heading, underlined with -
func (d *doc) h2(s string) {
	d.y++
	d.l.Text(d.x, d.y, s, white, Bold)
	d.l.Text(d.x, d.y+1, strings.Repeat("-", width(s)), dim, 0)
	d.y += 3
}

func (d *doc) para(spans []Span) {
	d.y = d.l.Para(d.x, d.y, d.w, spans, text, 0) + 1
}

// row puts left on the line and right flush to the edge
func (d *doc) row(left []Span, right []Span, fg RGB) {
	lw, rw := lineWidth(left), lineWidth(right)
	if lw+2+rw > d.w {
		d.l.Line(d.x, d.y, left, fg, 0)
		d.y++
		d.l.Line(d.x, d.y, right, fg, 0)
	} else {
		d.l.Line(d.x, d.y, left, fg, 0)
		d.l.Line(d.x+d.w-rw, d.y, right, fg, 0)
	}
	d.y++
}

// bullet is a hanging-indented line starting with mark
func (d *doc) bullet(mark string, markFg RGB, spans []Span) {
	d.l.Text(d.x, d.y, mark, markFg, 0)
	d.y = d.l.Para(d.x+width(mark)+1, d.y, d.w-width(mark)-1, spans, text, 0)
}

// field is "key   value", the value wrapped beside the key
func (d *doc) field(key string, spans []Span) {
	d.l.Text(d.x, d.y, key, dim, 0)
	d.y = d.l.Para(d.x+10, d.y, d.w-10, spans, text, 0)
}

func (d *doc) link(label, shown, url string, pad int) {
	d.l.Text(d.x, d.y, label, dim, 0)
	ux := d.x + pad
	if pad+width(shown) > d.w {
		d.y++
		ux = d.x + 2
	}
	end := d.l.Text(ux, d.y, fit(shown, d.w-(ux-d.x)), text, Underline)
	d.l.Linkify(ux, d.y, end-ux, url)
	d.y++
}

func (d *doc) footer(crumbs []string) {
	d.y += 2
	d.l.Text(d.x, d.y, strings.Repeat("-", d.w), dim, 0)
	d.y += 2
	x := d.l.Text(d.x, d.y, "~", yellow, Underline)
	d.l.Hit(d.x, d.y, 1, 1, "page:0")
	for i, c := range crumbs {
		x = d.l.Text(x, d.y, " / ", dim, 0)
		start := x
		x = d.l.Text(x, d.y, c, dim, 0)
		if i == 0 && len(crumbs) > 1 {
			d.l.Hit(start, d.y, x-start, 1, "back")
		}
	}
	d.y += 2
}

// ---------------- about ----------------

func (d *doc) about() {
	c := d.m.c
	d.h1("about me", yellow)
	d.para(plain("i'm Darsh Gupta, in grade 12 at John Fraser Secondary School in Mississauga, graduating in 2027. i build hardware: circuit boards in KiCad, mechanisms in Onshape, firmware in C and C++, and lately an FPGA board and a chip."))
	d.para(plain("i started with Python games in grade 8, moved to websites and apps, and got hooked on hardware the summer before grade 11, when a Hack Club Hackpad turned into a drone, a keyboard, and a trip to Undercity. most of what i care about now is getting other people building too: FraserHacks, the CS Club, Dayoneof, and Off-Script."))

	d.h2("what i do")
	for _, r := range c.Roles {
		mc := maskText[r.Mask]
		d.row([]Span{{S: r.Org, A: Bold, Fg: &mc}, {S: "  " + r.Title, Fg: &text}}, []Span{{S: r.When, Fg: &dim}}, text)
		for _, p := range r.Points {
			d.bullet("-", dim, plain(p))
		}
		d.y++
	}

	d.h2("frc team 1325")
	for _, p := range strings.Split(c.FRC.Body, "\n\n") {
		d.para(markdown(p, yellow, yellow))
	}
	for _, r := range c.FRC.Results {
		d.bullet("★", yellow, plain(r))
	}

	d.h2("awards")
	for _, a := range c.Awards {
		d.bullet("★", yellow, []Span{{S: a.What, A: Bold, Fg: &white}, {S: " · " + a.Who + " · " + a.When, Fg: &dim}})
	}

	d.h2("away from the desk")
	for _, o := range c.Offline {
		d.bullet("-", dim, []Span{{S: o.What, A: Bold, Fg: &white}, {S: " · " + o.Note}})
	}
}

// ---------------- projects ----------------

type filter struct {
	label string
	test  func(p *Project) bool
}

var filters = []filter{
	{"everything", func(*Project) bool { return true }},
	{"hardware", func(p *Project) bool { return p.Kind != "software" }},
	{"software", func(p *Project) bool { return p.Kind != "hardware" }},
	{"built", func(p *Project) bool { return p.Status == "built" || p.Status == "shipped" || p.Status == "ongoing" }},
	{"on paper", func(p *Project) bool { return p.Status == "designed" || p.Status == "assembled" }},
	{"won something", func(p *Project) bool { return p.Award != "" }},
	{"with friends", func(p *Project) bool { return len(p.With) > 0 }},
}

func (m *model) shown() []*Project {
	var ps []*Project
	for i := range m.c.Projects {
		if filters[m.filter].test(&m.c.Projects[i]) {
			ps = append(ps, &m.c.Projects[i])
		}
	}
	return ps
}

// pad is the status mark the site puts on each board: filled when it's real, dashed when it's on
// paper, half when it's assembled, green when it's what i'm on now
func pad(status string) (string, RGB) {
	switch status {
	case "designed":
		return "◌", gold
	case "assembled":
		return "◐", gold
	case "ongoing":
		return "●", green
	case "demoed":
		return "■", gold
	}
	return "●", gold
}

func (d *doc) projects() {
	m := d.m
	d.h1("things i've made", yellow)
	d.para(plain("everything, roughly in order of how much of me went into it. each one opens: pick it with ↑↓ (or click it) and press enter. the mark says where it's at:"))
	for _, k := range []struct{ s, label string }{{"built", "built and working"}, {"designed", "designed, not built yet"}, {"assembled", "assembled, debugging"}, {"ongoing", "what i'm working on now"}} {
		p, c := pad(k.s)
		d.bullet(p, c, plain(k.label))
	}
	d.y++

	// filters: the one that's on is in brackets, in yellow
	x := d.x
	x = d.l.Text(x, d.y, "show ", dim, 0)
	for i, f := range filters {
		n := 0
		for k := range m.c.Projects {
			if f.test(&m.c.Projects[k]) {
				n++
			}
		}
		s := f.label + " " + itoa(n)
		if i == m.filter {
			s = "[" + s + "]"
		} else {
			s = " " + s + " "
		}
		if x > d.x && x+width(s) > d.x+d.w {
			x, d.y = d.x+5, d.y+1
		}
		c, a := dim, Attr(0)
		if i == m.filter {
			c, a = yellow, Bold
		}
		end := d.l.Text(x, d.y, s, c, a)
		d.l.Hit(x, d.y, end-x, 1, "filter:"+itoa(i))
		x = end + 1
	}
	d.y += 3

	m.rows = m.rows[:0]
	for i, p := range m.shown() {
		sel := i == m.sel
		mc := maskText[p.Mask]
		mark, pc := pad(p.Status)
		top := d.y
		if sel {
			d.l.Text(d.x, d.y, ">", yellow, Bold)
		}
		d.l.Text(d.x+2, d.y, mark, pc, 0)
		var name []Span
		if sel {
			// the picked one sits on its own solder mask, like a board, with a nudge to open it
			d.l.Fill(d.x+4, d.y, width(p.Title)+2, 1, masks[p.Mask])
			d.l.Text(d.x+5, d.y, p.Title, ink(p.Mask), Bold)
			open := "  enter ↵ to open"
			d.l.Text(d.x+6+width(p.Title), d.y, open, yellow, 0)
			name = []Span{{S: " " + p.Title + " " + open}}
		} else {
			d.l.Text(d.x+4, d.y, p.Title, mc, Bold)
			name = []Span{{S: p.Title}}
		}
		right := p.StatusLabel + "  " + itoa(p.Year)
		if 4+lineWidth(name)+2+width(right) <= d.w {
			d.l.Text(d.x+d.w-width(right), d.y, right, dim, 0)
		}
		d.y++
		tc := dim
		if sel {
			tc = text
		}
		d.l.Text(d.x+4, d.y, fit(p.Tagline, d.w-4), tc, 0)
		d.y += 2
		d.l.Hit(d.x, top, d.w, 2, "project:"+p.Slug)
		m.rows = append(m.rows, [2]int{top, 2})
	}

	d.h2("what's on my bench")
	for _, b := range m.c.Bench {
		d.bullet("-", dim, []Span{{S: b.Title, A: Bold, Fg: &white}, {S: " · " + b.Note}, {S: " · " + b.When, Fg: &dim}})
	}
	d.h2("where it started, 2022–24")
	for _, e := range m.c.Early {
		t := Span{S: e.Title, A: Bold, Fg: &white}
		if e.URL != "" {
			t.A |= Underline
			t.Link = e.URL
		}
		d.bullet("-", dim, []Span{t, {S: " · " + e.Note}, {S: " · " + e.When, Fg: &dim}})
	}
}

// project is one project, up close (the site's card links out to GitHub; here it opens)
func (d *doc) project(p *Project) {
	c := d.m.c
	mc := maskText[p.Mask]
	d.h1(strings.ToLower(p.Title), mc)
	d.para([]Span{{S: p.Tagline, Fg: &white}})
	mark, pc := pad(p.Status)
	d.field("status", []Span{{S: mark + " ", Fg: &pc}, {S: p.StatusLabel}})
	d.field("when", plain(p.When))
	kind := map[string]string{"both": "hardware and software"}[p.Kind]
	if kind == "" {
		kind = p.Kind
	}
	d.field("kind", plain(kind))
	if len(p.Stack) > 0 {
		d.field("stack", plain(strings.Join(p.Stack, " · ")))
	}
	if len(p.With) > 0 {
		d.field("with", plain(strings.Join(p.With, ", ")))
	}
	if len(p.Events) > 0 {
		var at []Span
		for i, s := range p.Events {
			if e := c.event(s); e != nil {
				if i > 0 {
					at = append(at, Span{S: ", "})
				}
				ac := eventAccent[e.Theme]
				at = append(at, Span{S: e.Name, Fg: &ac, Link: e.URL})
			}
		}
		d.field("at", at)
	}
	if p.Award != "" {
		d.field("award", []Span{{S: "★ " + p.Award, Fg: &yellow}})
	}
	d.y++
	for _, para := range strings.Split(p.Body, "\n\n") {
		d.para(markdown(para, mc, mc))
	}
	if len(p.Links) > 0 {
		d.h2("links")
		pad := 0
		for _, lk := range p.Links {
			pad = max(pad, width(lk.Label))
		}
		for _, lk := range p.Links {
			d.link(lk.Label, lk.URL, lk.URL, pad+3)
		}
	}
}

// ---------------- events ----------------

func (d *doc) events() {
	c := d.m.c
	d.h1("events", yellow)
	d.para(plain("the ones i run, and the ones i've flown out to."))
	for _, sec := range []struct{ kind, title string }{{"ran", "things i run"}, {"went", "places i've been"}} {
		d.h2(sec.title)
		for i := range c.Events {
			e := &c.Events[i]
			if e.Kind != sec.kind {
				continue
			}
			ac := eventAccent[e.Theme]
			right := e.When
			if e.Role != "" {
				right = e.Role + " · " + e.When
			}
			d.row([]Span{{S: e.Name, A: Bold, Fg: &ac, Link: e.URL}}, []Span{{S: right, Fg: &dim}}, text)
			d.l.Para(d.x, d.y, d.w, plain(e.Place), dim, 0)
			d.y += paraHeight(plain(e.Place), d.w)
			d.y = d.l.Para(d.x, d.y, d.w, plain(e.Summary), text, 0)
			if len(e.Stats) > 0 {
				var st []Span
				for k, s := range e.Stats {
					if k > 0 {
						st = append(st, Span{S: " · ", Fg: &dim})
					}
					st = append(st, Span{S: s.N, A: Bold, Fg: &ac}, Span{S: " " + s.Label})
				}
				d.y = d.l.Para(d.x, d.y, d.w, st, text, 0)
			}
			if e.Award != "" {
				d.y = d.l.Para(d.x, d.y, d.w, []Span{{S: "★ " + e.Award, Fg: &yellow}}, text, 0)
			}
			if e.URL != "" {
				end := d.l.Text(d.x, d.y, fit(e.URL, d.w), dim, Underline)
				d.l.Linkify(d.x, d.y, end-d.x, e.URL)
				d.y++
			}
			d.y++
		}
	}
	d.h2("and also")
	for _, o := range c.OtherEvents {
		d.bullet("-", dim, []Span{{S: o.Name, A: Bold, Fg: &white}, {S: " · " + o.Note}, {S: " · " + o.When, Fg: &dim}})
	}
}

// ---------------- contact ----------------

func (d *doc) contact() {
	c := d.m.c
	d.h1("contact", yellow)
	d.para(plain("say hi."))
	d.link("email", c.Contact.Email, "mailto:"+c.Contact.Email, 12)
	for _, l := range c.Contact.Links {
		d.link(l.Label, l.Handle, l.URL, 12)
	}
	d.link("web", "www.darsh.app", "https://www.darsh.app", 12)
	d.y++
	d.para([]Span{{S: "my daily videos for dayoneof go up on "}, {S: "@darsh.builds", Fg: &yellow, Link: "https://www.instagram.com/darsh.builds/"}, {S: "."}})
}

// the masks themselves, for the picked project's highlight
var masks = map[string]RGB{
	"green": hex("#1e6b3d"), "blue": hex("#1d4f9e"), "purple": hex("#4d2b7c"), "red": hex("#a92a22"),
	"yellow": hex("#efc419"), "white": hex("#f4f4ef"), "black": hex("#16181b"),
}

func ink(mask string) RGB {
	if mask == "yellow" || mask == "white" {
		return hex("#16181b")
	}
	return silk
}
