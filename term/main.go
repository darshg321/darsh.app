// darsh.app over SSH: `ssh darsh.app`. A Wish server that hands every session a Bubble Tea program
// with the site in a terminal, from content.json and the sky photo in assets (scripts/export-term.js).
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	gossh "golang.org/x/crypto/ssh"
)

//go:embed content.json
var contentJSON []byte

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// profile guesses whether the visitor's terminal does 24-bit colour. ssh doesn't pass COLORTERM by
// default, so anything that says 256color gets truecolor too; nearly every terminal that does 256 does both.
func profile(term string, environ []string) Profile {
	for _, e := range environ {
		if e == "COLORTERM=truecolor" || e == "COLORTERM=24bit" {
			return TrueColor
		}
	}
	t := strings.ToLower(term)
	for _, s := range []string{"256color", "truecolor", "direct", "kitty", "alacritty", "ghostty", "wezterm", "foot", "rio", "contour"} {
		if strings.Contains(t, s) {
			return TrueColor
		}
	}
	return ANSI256
}

func main() {
	preview := flag.String("preview", "", "draw one frame at WxH to stdout instead of serving, e.g. 120x40")
	keys := flag.String("keys", "", "with -preview: keys to press first, space separated (e.g. \"2 down down enter\")")
	flag.Parse()

	var c Content
	if err := json.Unmarshal(contentJSON, &c); err != nil {
		log.Fatalf("content.json: %v", err)
	}

	if *preview != "" {
		var w, h int
		fmt.Sscanf(*preview, "%dx%d", &w, &h)
		m := newModel(&c, TrueColor, w, h)
		for _, k := range strings.Fields(*keys) {
			var msg tea.KeyMsg
			switch k {
			case "up":
				msg = tea.KeyMsg{Type: tea.KeyUp}
			case "down":
				msg = tea.KeyMsg{Type: tea.KeyDown}
			case "left":
				msg = tea.KeyMsg{Type: tea.KeyLeft}
			case "right":
				msg = tea.KeyMsg{Type: tea.KeyRight}
			case "enter":
				msg = tea.KeyMsg{Type: tea.KeyEnter}
			case "esc":
				msg = tea.KeyMsg{Type: tea.KeyEsc}
			case "pgdn":
				msg = tea.KeyMsg{Type: tea.KeyPgDown}
			case "tab":
				msg = tea.KeyMsg{Type: tea.KeyTab}
			default:
				if strings.HasPrefix(k, "scroll") {
					var n int
					fmt.Sscanf(k, "scroll%d", &n)
					m.View()
					m.off += n
					continue
				}
				msg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(k)}
			}
			m.View()
			m.Update(msg)
		}
		fmt.Print(m.View())
		return
	}

	addr := net.JoinHostPort(env("HOST", "0.0.0.0"), env("PORT", "2222"))
	srv, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithHostKeyPath(env("HOST_KEY", ".ssh/id_ed25519")),
		// anyone can come in, no keys or passwords
		wish.WithPublicKeyAuth(func(ssh.Context, ssh.PublicKey) bool { return true }),
		wish.WithKeyboardInteractiveAuth(func(ssh.Context, gossh.KeyboardInteractiveChallenge) bool { return true }),
		wish.WithIdleTimeout(30*time.Minute),
		wish.WithMiddleware(
			bm.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				pty, _, _ := s.Pty()
				return newModel(&c, profile(pty.Term, s.Environ()), pty.Window.Width, pty.Window.Height),
					[]tea.ProgramOption{tea.WithAltScreen(), tea.WithMouseCellMotion()}
			}),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	log.Printf("listening on %s", addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
