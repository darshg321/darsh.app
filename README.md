# darsh.app

My personal site: SvelteKit 3 + Svelte 5, prerendered to static files with `adapter-static`.
The look is the "sunny" concept: a real sky photo, and every panel is a circuit board with plated mounting holes.

```sh
npm install
npm run dev       # local dev
npm run build     # static site in build/
npm run check     # types
```

## Where things live

| What | File |
|---|---|
| every project, in display order (most effort / current first) | `src/lib/data/projects.ts` |
| events i run and events i went to, each with a theme | `src/lib/data/events.ts` |
| roles, FRC, awards | `src/lib/data/about.ts` |
| event themes (colours, fonts, textures borrowed from each event's site) | `src/lib/styles/themes.css` |
| the board look | `src/app.css` |

Pages: `/` home, `/projects` (filterable grid), `/events` (themed cards), `/about`.
There are no project or event pages: a project card links to the project's GitHub repo (its first GitHub link) and
lists the events it went to, and an event card links to the event's own site (its `url`).

## Images

`scripts/images.sh` builds `static/img/**.webp` from the archived catalogue in `archive/content` (repo images,
contributions, and `archive/content/remote/` for event photos, logos and site screenshots). Add a line to its
manifest and re-run it. Videos in `static/video/` were re-encoded from the repos and the Off-Script recap.

## Status labels

Project status is deliberately honest: **built** means it physically exists and works, **designed** means
CAD/PCB/BOM only, **assembled, debugging** is Minicopter, **working on it now** is current work.

## archive/

Everything that was in this folder before the site: the design concepts (`concept/`, including `sunny/`), the
inspiration screenshots, and the full content catalogue (`content/CATALOGUE.md`, every repo's README, journal
and images).

## term/ (the site over SSH)

`ssh darsh.app` is the site as a terminal app, styled after term.samliu.dev. The landing is the sky photo as
coloured ASCII with a `+---+` panel (figlet name, nav, email), a live box (Mississauga time and what i'm up to, flipped
with the arrow keys) and a hint. The inner pages (about, projects, events, contact) read like a terminal document: a
strip of the sky with the nav hanging off it, `====` headings, wrapped text and a `~ / page` breadcrumb, with the
site's colours as accents (projects in their solder-mask colours, events in their own, yellow for what's picked).
Projects open into their own page. It's a Go server on [Wish](https://github.com/charmbracelet/wish) and
[Bubble Tea](https://github.com/charmbracelet/bubbletea); the figlet font is `small` from the FIGlet distribution.

It reads `term/content.json` and `term/assets/img/sky.jpg`, both written by `npm run term:export` from
`src/lib/data` and `static/img`, so re-export after editing the data. The words that live in the Svelte pages rather
than the data (the about intro, page intros, the landing's tagline and "what i'm up to") are copied into `term/pages.go` and
`term/landing.go`, so update both.

```sh
npm run term:export                       # src/lib/data + the sky photo -> term/
cd term && go run .                       # listens on :2222, host key in term/.ssh/
ssh -p 2222 localhost
go run . -preview 120x40 -keys "2 down enter"   # print one frame without ssh
go test ./...                             # every page and project at 60x20 to 200x60, and the ANSI size budget
```

It runs on a free Google Cloud e2-micro (project `darsh-term-5d0b61`, VM `term`, us-central1-a, 35.209.174.254).
`darsh.app`'s A record points at the VM so `ssh darsh.app` works; Caddy on the VM (`term/Caddyfile`) redirects web
visitors to `www.darsh.app`, which Vercel serves. The VM's own admin SSH is on port 2200. After editing the data,
`term/deploy.sh` re-exports, rebuilds and restarts it (needs `go` and `gcloud` on the PATH). A billing budget emails
if the project ever costs anything.
