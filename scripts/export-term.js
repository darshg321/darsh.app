// Dumps the site's data to term/content.json, and copies the sky photo into term/assets, for the SSH
// version (term/). Run after editing src/lib/data.
import { cpSync, mkdirSync, rmSync, writeFileSync } from 'node:fs';
import { dirname } from 'node:path';
import { createServer } from 'vite';

// the landing page's backdrop, drawn in ASCII
const images = ['/img/sky.jpg'];

const vite = await createServer({ server: { middlewareMode: true }, appType: 'custom', logLevel: 'error' });
try {
	const { projects, statusLabel, bench, early, repo } = await vite.ssrLoadModule('/src/lib/data/projects.ts');
	const { events, otherEvents } = await vite.ssrLoadModule('/src/lib/data/events.ts');
	const about = await vite.ssrLoadModule('/src/lib/data/about.ts');

	const out = {
		projects: projects.map((p) => ({
			slug: p.slug,
			title: p.title,
			tagline: p.tagline,
			when: p.when,
			year: p.year,
			kind: p.kind,
			status: p.status,
			statusLabel: statusLabel[p.status],
			mask: p.mask,
			stack: p.stack,
			with: (p.with ?? []).map((w) => w.name),
			// events it lists, and events that list it, like ProjectCard does
			events: events.filter((e) => p.events?.includes(e.slug) || e.projects?.includes(p.slug)).map((e) => e.slug),
			award: p.award ?? '',
			repo: repo(p) ?? '',
			links: p.links,
			body: p.body
		})),
		events: events.map((e) => ({
			slug: e.slug,
			name: e.name,
			kind: e.kind,
			role: e.role ?? '',
			when: e.when,
			place: e.place,
			theme: e.theme,
			url: e.url ?? '',
			summary: e.summary,
			stats: e.stats ?? [],
			award: e.award ?? '',
			projects: e.projects ?? []
		})),
		otherEvents,
		bench,
		early,
		contact: about.contact,
		roles: about.roles.map(({ title, org, when, points, mask }) => ({ title, org, when, points, mask })),
		frc: about.frc,
		awards: about.awards.map(({ what, who, when }) => ({ what, who, when })),
		offline: about.offline
	};
	writeFileSync('term/content.json', JSON.stringify(out, null, '\t') + '\n');

	rmSync('term/assets', { recursive: true, force: true });
	let copied = 0;
	for (const src of images) {
		mkdirSync(dirname('term/assets' + src), { recursive: true });
		cpSync('static' + src, 'term/assets' + src);
		copied++;
	}
	console.log(`term/content.json: ${out.projects.length} projects, ${out.events.length} events; ${copied} image in term/assets`);
} finally {
	await vite.close();
}
