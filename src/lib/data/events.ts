// Events i've run and events i've been to. `theme` picks a look borrowed from the event's own site
// (see src/lib/styles/themes.css); `url` is the event's own site. projects link back by slug.

export type Theme =
	| 'fraserhacks'
	| 'dayoneof'
	| 'offscript'
	| 'csclub'
	| 'polaris'
	| 'outpost'
	| 'fallout'
	| 'alleyway'
	| 'undercity'
	| 'hackthenorth'
	| 'uottahack';

export interface Event {
	slug: string;
	name: string;
	kind: 'ran' | 'went';
	role?: string;
	when: string;
	place: string;
	theme: Theme;
	/** the event's own site; the card links straight to it */
	url?: string;
	logo?: string;
	cover?: string;
	summary: string;
	stats?: { n: string; label: string }[];
	award?: string;
	projects?: string[];
}

export const events: Event[] = [
	// ---------------- things i run ----------------
	{
		slug: 'fraserhacks',
		name: 'FraserHacks',
		kind: 'ran',
		role: 'co-lead',
		when: '2024 – now',
		place: 'John Fraser SS, Mississauga',
		url: 'https://fraserhacks.dev',
		theme: 'fraserhacks',
		logo: '/img/e/fh26-logo.webp',
		cover: '/img/e/fh26/leads.webp',
		summary: "Mississauga's largest high school hackathon. i co-lead it.",
		stats: [
			{ n: '83', label: 'hackers in 2026' },
			{ n: '28', label: 'projects shipped' },
			{ n: '98%', label: 'submission rate' },
			{ n: '10', label: 'person exec team' }
		],
		projects: ['fraserhacks-sites']
	},
	{
		slug: 'dayoneof',
		name: 'Dayoneof',
		kind: 'ran',
		role: 'host',
		when: 'Sep 2026 – now',
		place: 'Hack Club, online',
		url: 'https://dayoneof.hackclub.com',
		theme: 'dayoneof',
		logo: '/img/e/dayoneof-wordmark.webp',
		cover: '/img/s/dayoneof.webp',
		summary: 'a Hack Club program: post a short video every day for a month.',
		projects: ['dayoneof']
	},
	{
		slug: 'offscript',
		name: 'Off-Script',
		kind: 'ran',
		role: 'host',
		when: 'Aug – Sep 2026',
		place: 'University of Waterloo, Waterloo Tech Week',
		url: 'https://wygo.world/off-script-ai-hackathon',
		theme: 'offscript',
		logo: '/img/e/wtw-logo.svg',
		cover: '/img/e/offscript/table.webp',
		summary: 'a technical showcase i hosted at UWaterloo as part of Waterloo Tech Week, as a high schooler.',
		stats: [
			{ n: '120+', label: 'RSVPs' },
			{ n: '50+', label: 'came' },
			{ n: '10+', label: 'projects demoed' }
		],
		projects: ['icey', 'bing-bang-dream-adventure']
	},
	{
		slug: 'csclub',
		name: 'Fraser Comp Sci Club',
		kind: 'ran',
		role: 'president',
		when: 'Oct 2023 – now',
		place: 'John Fraser SS, Thursdays at lunch, room 183',
		url: 'https://cs.fraserhacks.dev',
		theme: 'csclub',
		logo: '/img/e/cs-logo.webp',
		cover: '/img/s/cs.webp',
		summary: 'build cool stuff, win prizes, compete. Thursdays at lunch.',
		stats: [
			{ n: '20+', label: 'members every week' },
			{ n: '1', label: 'trinket every meeting' },
			{ n: '3', label: 'competitions we train for' }
		],
		projects: ['cs-club-site', 'arcadetris']
	},

	// ---------------- places i've been ----------------
	{
		slug: 'polaris',
		name: 'Horizons Polaris',
		kind: 'went',
		when: 'Aug 7–9, 2026',
		place: 'Toronto',
		url: 'https://horizons.hackclub.com',
		theme: 'polaris',
		logo: '/img/e/horizons-logo.svg',
		cover: '/img/p/bingbang/intro.webp',
		summary: "one of Hack Club's six Horizons hackathons, \"in the heart of Canada\".",
		award: '2nd place',
		projects: ['bing-bang-dream-adventure', 'ultisense'],
	},
	{
		slug: 'outpost',
		name: 'Outpost',
		kind: 'went',
		when: 'Jul 14–20, 2026',
		place: 'San Francisco + Open Sauce',
		url: 'https://outpost.hackclub.com',
		theme: 'outpost',
		logo: '/img/e/outpost-logo.webp',
		cover: '/img/s/outpost.webp',
		summary: 'a 4-day Hack Club hardware hackathon, then two days showing projects at Open Sauce.',
	},
	{
		slug: 'fallout',
		name: 'Fallout',
		kind: 'went',
		when: 'Jul 1–7, 2026',
		place: 'Shenzhen, China',
		url: 'https://fallout.hackclub.com',
		theme: 'fallout',
		logo: '/img/e/fallout/logo.webp',
		cover: '/img/s/fallout.webp',
		summary: 'build hardware for free, then go to Shenzhen. a week in the hardware capital of the world.',
		award: 'funniest project',
		projects: ['jollipop', 'keychain-tris', 'whale-badge', 'aura-generator'],
	},
	{
		slug: 'alleyway',
		name: 'Alleyway × Hackaday Supercon',
		kind: 'went',
		when: 'Nov 2025',
		place: 'Pasadena, California',
		url: 'https://hackaday.io/superconference/',
		theme: 'alleyway',
		logo: '/img/e/hackaday-logo.svg',
		cover: '/img/p/toparm/build.webp',
		summary: 'Hack Club picked me to bring projects to Hackaday Supercon.',
		award: '3rd, costume contest ($5k lab instrument)',
		projects: ['toparm', 'aura-generator'],
	},
	{
		slug: 'hackthenorth',
		name: 'Hack the North 2025',
		kind: 'went',
		when: 'Sep 2025',
		place: 'University of Waterloo',
		url: 'https://hackthenorth.com',
		theme: 'hackthenorth',
		cover: '/img/s/htn.webp',
		summary: "Canada's biggest hackathon.",
		stats: [
			{ n: '1000+', label: 'hackers' },
			{ n: '36', label: 'hours' }
		],
		projects: ['guessrgo']
	},
	{
		slug: 'undercity',
		name: 'Undercity',
		kind: 'went',
		when: 'Jul 11–14, 2025',
		place: 'GitHub HQ, San Francisco',
		url: 'https://undercity.hackclub.com',
		theme: 'undercity',
		logo: '/img/e/undercity-logo.svg',
		cover: '/img/p/arcadetris/start.webp',
		summary: "Hack Club's 4-day hardware hackathon at GitHub HQ.",
		projects: ['arcadetris', 'minicopter', 'fancyboard', 'hackpad'],
	},
	{
		slug: 'uottahack',
		name: 'uOttaHack 7',
		kind: 'went',
		when: 'Jan 2025',
		place: 'University of Ottawa',
		url: 'https://uottahack.ca',
		theme: 'uottahack',
		cover: '/img/s/uottahack.webp',
		summary: 'a weekend hackathon in Ottawa.',
		projects: ['diarymuse']
	}
];

export const eventUrl = (slug: string) => events.find((e) => e.slug === slug)?.url;

/** smaller things that don't get a card */
export const otherEvents = [
	{ name: 'UTMUN', note: 'Model UN at U of T', when: '2025 – 26' },
	{ name: 'KingsMUN', note: 'Model UN', when: '2025 – 26' },
	{ name: 'UofT Science Unlimited', note: 'summer camp; finalist with an AI garbage sorter', when: '2025' }
];

