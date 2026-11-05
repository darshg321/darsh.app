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
| roles, FRC, awards, school, timeline | `src/lib/data/about.ts` |
| event themes (colours, fonts, textures borrowed from each event's site) | `src/lib/styles/themes.css` |
| the board look | `src/app.css` |

Pages: `/` home, `/projects` (filterable grid), `/events` and `/events/[slug]` (themed), `/about`.
There are no project pages: a project card links to the project's GitHub repo (its first GitHub link) and lists
the events it went to.

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
