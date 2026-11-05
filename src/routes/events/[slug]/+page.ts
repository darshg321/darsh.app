import { error } from '@sveltejs/kit';
import { events, eventBySlug } from '#lib/data/events.js';
import { projects } from '#lib/data/projects.js';
import { md } from '#lib/md.js';
import type { EntryGenerator, PageLoad } from './$types';

export const entries: EntryGenerator = () => events.map((e) => ({ slug: e.slug }));

export const load: PageLoad = ({ params }) => {
	const e = eventBySlug(params.slug);
	if (!e) error(404, 'no event here');
	const made = projects.filter((p) => e.projects?.includes(p.slug) || p.events?.includes(e.slug));
	return { e, body: md(e.body), made };
};
