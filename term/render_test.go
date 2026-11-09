package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func content(t *testing.T) *Content {
	var c Content
	if err := json.Unmarshal(contentJSON, &c); err != nil {
		t.Fatal(err)
	}
	return &c
}

// Every page draws at common sizes without panicking, fills the screen exactly, and stays inside the
// ANSI budget: a full repaint (scrolling, the sky drifting) shouldn't cost a visitor much.
func TestPagesFitAndStayLight(t *testing.T) {
	c := content(t)
	for _, sz := range [][2]int{{60, 20}, {80, 24}, {120, 40}, {200, 60}} {
		for page := 0; page <= pageContact; page++ {
			m := newModel(c, TrueColor, sz[0], sz[1])
			m.page = page
			out := m.View()
			if lines := strings.Count(out, "\n") + 1; lines != sz[1] {
				t.Errorf("%dx%d page %d: %d lines, want %d", sz[0], sz[1], page, lines, sz[1])
			}
			if kb := len(out) / 1024; kb > sz[0]*sz[1]/60 {
				t.Errorf("%dx%d page %d: %d KB a frame", sz[0], sz[1], page, kb)
			}
		}
		// every project opens
		for i := range c.Projects {
			m := newModel(c, ANSI256, sz[0], sz[1])
			m.open(c.Projects[i].Slug)
			m.View()
		}
	}
}
