package main

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	// the landing clock only needs to keep up with the minute
	tickEvery = 15 * time.Second
	idleAfter = 10 * time.Minute
)

type tickMsg time.Time

type model struct {
	c    *Content
	prof Profile

	w, h   int
	page   int // pageLanding, pageAbout, ...
	detail int // index into c.Projects being read, or -1
	navSel int // the landing nav's pick

	filter, sel int
	off         int
	listOff     int // where the project list was scrolled to before opening one
	rows        [][2]int

	lastIn  time.Time
	ticking bool
	wantSel bool

	layer *Layer
	dirty bool
	help  bool // the ? overlay
}

func newModel(c *Content, prof Profile, w, h int) *model {
	return &model{c: c, prof: prof, w: max(w, 20), h: max(h, 8), detail: -1, dirty: true, lastIn: time.Now()}
}

func tick() tea.Cmd { return tea.Tick(tickEvery, func(t time.Time) tea.Msg { return tickMsg(t) }) }

func (m *model) Init() tea.Cmd {
	m.ticking = true
	return tick()
}

func (m *model) touch() tea.Cmd {
	m.lastIn = time.Now()
	if !m.ticking {
		m.ticking = true
		return tick()
	}
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.w, m.h = max(msg.Width, 20), max(msg.Height, 8)
		m.dirty = true
		return m, nil

	case tickMsg:
		if time.Since(m.lastIn) > idleAfter {
			m.ticking = false
			return m, nil
		}
		if m.page == pageLanding {
			m.dirty = true
		}
		return m, tick()

	case tea.MouseMsg:
		cmd := m.touch()
		switch msg.Button {
		case tea.MouseButtonWheelUp:
			m.off -= 3
		case tea.MouseButtonWheelDown:
			m.off += 3
		case tea.MouseButtonLeft:
			if msg.Action == tea.MouseActionPress {
				m.click(msg.X, msg.Y)
			}
		}
		return m, cmd

	case tea.KeyMsg:
		cmd := m.touch()
		if m.help && msg.String() != "?" && msg.String() != "q" && msg.String() != "ctrl+c" {
			// any key closes the overlay
			m.help = false
			return m, cmd
		}
		if m.key(msg.String()) {
			return m, tea.Quit
		}
		return m, cmd
	}
	return m, nil
}

// key handles a key press; true means quit
func (m *model) key(k string) bool {
	switch k {
	case "q", "ctrl+c":
		return true
	case "?":
		m.help = !m.help
		return false
	case "1", "2", "3", "4":
		m.go_(int(k[0] - '0'))
		return false
	}
	if m.page == pageLanding {
		switch k {
		case "right", "l", "tab":
			m.navSel = (m.navSel + 1) % len(navItems)
		case "left", "h", "shift+tab":
			m.navSel = (m.navSel + len(navItems) - 1) % len(navItems)
		case "down", "j", "enter", " ":
			m.go_(m.navSel + 1)
		case "esc":
			return true
		}
		m.dirty = true
		return false
	}
	switch k {
	case "esc", "backspace":
		m.back()
	case "tab", "right", "l":
		m.go_(m.page%len(navItems) + 1)
	case "shift+tab", "left", "h":
		m.go_((m.page+len(navItems)-2)%len(navItems) + 1)
	case "pgdown", " ":
		m.off += m.h - 4
	case "pgup":
		m.off -= m.h - 4
	case "g", "home":
		m.off = 0
	case "G", "end":
		m.off = 1 << 20
	case "up", "k":
		switch {
		case m.page == pageProjects && m.detail < 0 && m.sel > 0:
			m.listKey(k)
		case m.page == pageProjects && m.detail < 0, m.off <= 0:
			// up at the top closes the tab (or the project, back to the list)
			m.back()
		default:
			m.off -= 2
		}
	case "down", "j", "enter", "f":
		if m.page == pageProjects && m.detail < 0 {
			m.listKey(k)
		} else if k == "down" || k == "j" {
			m.off += 2
		}
	}
	return false
}

func (m *model) go_(p int) {
	if p >= 1 && p <= len(navItems) {
		m.navSel = p - 1
	}
	m.page, m.detail, m.off = p, -1, 0
	m.dirty = true
}

func (m *model) back() {
	if m.detail >= 0 {
		m.detail = -1
		m.off = m.listOff
		m.wantSel = true
	} else {
		m.page = pageLanding
	}
	m.dirty = true
}

func (m *model) open(slug string) {
	for k, p := range m.shown() {
		if p.Slug == slug {
			m.sel = k
		}
	}
	for i := range m.c.Projects {
		if m.c.Projects[i].Slug == slug {
			m.listOff = m.off
			if m.page != pageProjects {
				m.listOff = 0
			}
			m.detail, m.page, m.off = i, pageProjects, 0
			m.dirty = true
		}
	}
}

func (m *model) listKey(k string) {
	n := len(m.shown())
	if n == 0 {
		return
	}
	switch k {
	case "up", "k":
		m.sel = max(0, m.sel-1)
	case "down", "j":
		m.sel = min(n-1, m.sel+1)
	case "f":
		m.filter = (m.filter + 1) % len(filters)
		m.sel = 0
	case "enter":
		m.open(m.shown()[m.sel].Slug)
		return
	}
	m.dirty = true
	m.wantSel = true
}

func (m *model) click(x, y int) {
	if m.layer == nil {
		return
	}
	if m.page != pageLanding {
		y += m.off
	}
	for i := len(m.layer.Hits) - 1; i >= 0; i-- {
		h := m.layer.Hits[i]
		if x < h.X || x >= h.X+h.W || y < h.Y || y >= h.Y+h.H {
			continue
		}
		switch a := h.Action; {
		case strings.HasPrefix(a, "page:"):
			m.go_(atoi(a[5:]))
		case strings.HasPrefix(a, "project:"):
			m.open(a[8:])
		case strings.HasPrefix(a, "filter:"):
			m.filter, m.sel = atoi(a[7:]), 0
			m.dirty = true
		case a == "back":
			m.back()
		}
		return
	}
}

func (m *model) View() string {
	if m.dirty || m.layer == nil || m.layer.W != m.w {
		if m.page == pageLanding {
			m.layer = m.landing()
		} else {
			m.layer = m.inner()
		}
		m.dirty = false
	}
	if m.page == pageLanding {
		out := Over(m.layer, NewLayer(m.w, 1), 0)
		m.overlay(out)
		return Render(out, m.prof)
	}
	if m.wantSel {
		m.keepSelVisible()
		m.wantSel = false
	}
	m.off = max(0, min(m.off, m.layer.H-m.h))
	base := NewLayer(m.w, m.h)
	base.Fill(0, 0, m.w, m.h, ground)
	out := Over(base, m.layer, m.off)
	m.keyBar(out)
	m.overlay(out)
	return Render(out, m.prof)
}

type keyHelp struct{ k, what string }

// keys is what works where you are, most useful first
func (m *model) keys() []keyHelp {
	switch {
	case m.page == pageLanding:
		return []keyHelp{{"←→", "pick a tab"}, {"↓ enter", "open it"}, {"1-4", "jump to a tab"}, {"q esc", "quit"}}
	case m.detail >= 0:
		return []keyHelp{{"esc", "back to projects"}, {"↑↓", "scroll"}, {"↑ at top", "back"}, {"←→", "other tabs"}, {"1-4", "jump"}, {"q", "quit"}}
	case m.page == pageProjects:
		return []keyHelp{{"esc", "back"}, {"↑↓", "pick"}, {"enter", "open"}, {"f", "filter"}, {"←→", "tabs"}, {"↑ at top", "close"}, {"1-4", "jump"}, {"q", "quit"}}
	}
	return []keyHelp{{"esc", "back"}, {"↑↓", "scroll"}, {"↑ at top", "close"}, {"←→", "tabs"}, {"space", "page down"}, {"1-4", "jump"}, {"q", "quit"}}
}

// keyBar pins the keys along the bottom of an inner page, as many as fit, then "? all keys"
func (m *model) keyBar(out *Layer) {
	y := m.h - 1
	out.Fill(0, y, m.w, 1, panel)
	more := "? all keys"
	x := 2
	for _, k := range m.keys() {
		w := width(k.k) + 1 + width(k.what) + 3
		if x+w+width(more) > m.w-2 {
			break
		}
		x = out.Text(x, y, k.k, yellow, Bold)
		x = out.Text(x+1, y, k.what, text, 0) + 3
	}
	out.Text(m.w-2-width(more), y, more, dim, 0)
}

// overlay is the ? panel: every key, wherever you are
func (m *model) overlay(out *Layer) {
	if !m.help {
		return
	}
	all := []keyHelp{
		{"←→ / tab", "move between tabs"},
		{"↓ / enter", "open the tab (from the start screen)"},
		{"1 2 3 4", "about, projects, events, contact"},
		{"↑↓ / j k", "scroll, or pick a project"},
		{"enter", "open the picked project"},
		{"f", "change the projects filter"},
		{"esc", "go back a level"},
		{"↑ at the top", "close the tab, or the project"},
		{"space pgdn", "page down (pgup up)"},
		{"g G", "top, bottom"},
		{"mouse", "scroll, click tabs and projects"},
		{"q", "quit"},
	}
	kw := 0
	for _, k := range all {
		kw = max(kw, width(k.k))
	}
	w := min(m.w-2, kw+2+36+6)
	h := min(m.h-2, len(all)+4)
	x, y := (m.w-w)/2, (m.h-h)/2
	out.Fill(x, y, w, h, panel)
	box(out, x, y, w, h, border)
	out.Text(x+3, y+1, "keys", white, Bold)
	for i, k := range all {
		if 2+i >= h-1 {
			break
		}
		out.Text(x+3, y+2+i, k.k, yellow, Bold)
		out.Text(x+3+kw+2, y+2+i, fit(k.what, w-kw-8), text, 0)
	}
	note := "any key closes"
	out.Text(x+w-3-width(note), y+1, note, dim, 0)
}

func (m *model) keepSelVisible() {
	if m.detail >= 0 || m.sel >= len(m.rows) {
		return
	}
	r := m.rows[m.sel]
	if r[0]-1 < m.off {
		m.off = r[0] - 1
	} else if r[0]+r[1]+1 > m.off+m.h {
		m.off = r[0] + r[1] + 1 - m.h
	}
}
