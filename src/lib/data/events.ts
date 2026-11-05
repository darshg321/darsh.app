// Events i've run and events i've been to. `theme` picks a look borrowed from the event's own site
// (see src/lib/styles/themes.css); projects link back by slug.

import type { Img } from './projects';

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
	| 'uottahack'
	| 'plain';

export interface Event {
	slug: string;
	name: string;
	kind: 'ran' | 'went';
	role?: string;
	when: string;
	place: string;
	theme: Theme;
	logo?: string;
	cover?: string;
	// art borrowed from the event's site, stuck beside the write-up
	decor?: string[];
	summary: string;
	body: string;
	stats?: { n: string; label: string }[];
	award?: string;
	gallery?: Img[];
	video?: { src: string; caption: string };
	links?: { label: string; url: string }[];
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
		theme: 'fraserhacks',
		logo: '/img/e/fh26-logo.webp',
		cover: '/img/e/fh26/leads.webp',
		summary: "Mississauga's largest high school hackathon. i co-lead it.",
		body: `FraserHacks is our school's hackathon, and the biggest high school one in Mississauga. i joined as a tech exec for 2024: i built the check-in system (React, MongoDB, Google Cloud) that took 100+ live check-ins without going down, worked on the site, and spent the day helping teams get unstuck.

since 2025 i've co-led it. i run fundraising, logistics and operations with a 10-person exec team: we got sponsorship from Rotary and AP starting from zero contacts and almost no budget, sorted the custom merch with suppliers, brought in industry judges and workshops, and pushed CS to students across Mississauga and Oakville. on the day i mentor teams through ideas, prototypes and live demos.

FraserHacks 2026 ran on March 26: 83 hackers shipped 28 projects (a 98% submission rate) for 7 judges. now we're planning 2027.`,
		stats: [
			{ n: '83', label: 'hackers in 2026' },
			{ n: '28', label: 'projects shipped' },
			{ n: '98%', label: 'submission rate' },
			{ n: '10', label: 'person exec team' }
		],
		gallery: [
			{ src: '/img/e/fh26/leads.webp', alt: 'the two co-leads in FraserHacks hoodies' },
			{ src: '/img/e/fh26/stage2.webp', alt: 'opening the event on stage' },
			{ src: '/img/e/fh26/hall.webp', alt: 'the full hall of hackers' },
			{ src: '/img/e/fh26/mentoring.webp', alt: 'helping a team at their laptop' },
			{ src: '/img/e/fh26/merch.webp', alt: 'hackers with the custom merch' },
			{ src: '/img/e/fh26/rotary.webp', alt: 'a speaker from Rotary, one of our sponsors' },
			{ src: '/img/e/fh26/tables.webp', alt: 'teams working' },
			{ src: '/img/e/fh24/1.webp', alt: 'the library at FraserHacks 2024' },
			{ src: '/img/e/fh24/4.webp', alt: 'teams at FraserHacks 2024' }
		],
		links: [
			{ label: 'fraserhacks.dev (2027)', url: 'https://fraserhacks.dev' },
			{ label: '2026.fraserhacks.dev', url: 'https://2026.fraserhacks.dev' }
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
		theme: 'dayoneof',
		logo: '/img/e/dayoneof-wordmark.webp',
		cover: '/img/s/dayoneof.webp',
		summary: 'a Hack Club program: post a short video every day for a month.',
		body: `most people spend too much time scrolling. Dayoneof asks teens to post a short-form video every day for a month instead, about anything at all, and give each other feedback along the way. keep the streak and you earn prizes.

i run it with Celestial and Zach Latta from Hack Club, and i built the site, the streak tracking and the Slack bot. i'm doing it too: my day one reel is the one on the front page, and the rest are on [@darsh.builds](https://www.instagram.com/darsh.builds/).`,
		gallery: [
			{ src: '/img/s/dayoneof.webp', alt: 'the Dayoneof site' },
			{ src: '/img/e/dayoneof-reel.webp', alt: 'my day one reel' }
		],
		links: [
			{ label: 'dayoneof.hackclub.com', url: 'https://dayoneof.hackclub.com' },
			{ label: '@darsh.builds', url: 'https://www.instagram.com/darsh.builds/' }
		],
		projects: ['dayoneof']
	},
	{
		slug: 'offscript',
		name: 'Off-Script',
		kind: 'ran',
		role: 'host',
		when: 'Aug – Sep 2026',
		place: 'University of Waterloo, Waterloo Tech Week',
		theme: 'offscript',
		logo: '/img/e/wtw-logo.svg',
		cover: '/img/e/offscript/table.webp',
		summary: 'a technical showcase i hosted at UWaterloo as part of Waterloo Tech Week, as a high schooler.',
		body: `Off-Script was a showcase of things people had actually built, held at the University of Waterloo during Waterloo Tech Week. i organised and hosted it while still in high school.

120+ people RSVP'd and 50+ came, mostly engineering, CS and math students, along with PhD researchers and people from industry. i curated 10+ projects with live demos across FPGA hardware, embedded systems and hardware/software integration. i brought Icey and Bing Bang Dream Adventure.`,
		stats: [
			{ n: '120+', label: 'RSVPs' },
			{ n: '50+', label: 'came' },
			{ n: '10+', label: 'projects demoed' }
		],
		gallery: [
			{ src: '/img/e/offscript/table.webp', alt: 'people crowded round a table, talking over laptops' },
			{ src: '/img/e/offscript/pizza.webp', alt: 'the room filling up around the pizza' },
			{ src: '/img/e/offscript/demo2.webp', alt: 'someone showing a project on a laptop' },
			{ src: '/img/e/offscript/whiteboard.webp', alt: 'two people sketching on the whiteboard' },
			{ src: '/img/e/offscript/chairs.webp', alt: 'a circle of attendees in the chairs by the window' },
			{ src: '/img/e/offscript/plates.webp', alt: 'attendees with pizza, waving at the camera' },
			{ src: '/img/e/offscript/wide.webp', alt: 'the room early on, before most people arrived' }
		],
		video: { src: '/video/offscript.mp4', caption: 'the recap' },
		projects: ['icey', 'bing-bang-dream-adventure']
	},
	{
		slug: 'csclub',
		name: 'Fraser Comp Sci Club',
		kind: 'ran',
		role: 'president',
		when: 'Oct 2023 – now',
		place: 'John Fraser SS, Thursdays at lunch, room 183',
		theme: 'csclub',
		logo: '/img/e/cs-logo.webp',
		cover: '/img/s/cs.webp',
		summary: 'build cool stuff, win prizes, compete. Thursdays at lunch.',
		body: `build cool stuff, win prizes, compete, and take home a 3D-printed trinket every meeting. no applications and no experience needed. i joined in grade 9, was the teaching lead on the exec in grade 10, and now i run it with Yichen Li and Jason Chou.

every Thursday is something you build or play with: a speed build of an LED heart on a breadboard (fastest one wins a Tims card), Arduino challenges, a two-week Smash Ultimate tourney, and coming up, soldering a coin-cell keychain, a macropad keychain, and designing a first PCB. we also go after the CCC, Peel Science Fair and Peel Skills, head to local hackathons as a group, and run FraserHacks.

anything a member makes goes up in the project directory on the club site with their name on it.`,
		stats: [
			{ n: '20+', label: 'members every week' },
			{ n: '1', label: 'trinket every meeting' },
			{ n: '3', label: 'competitions we train for' }
		],
		gallery: [
			{ src: '/img/s/cs.webp', alt: 'the club site' },
			{ src: '/img/e/cs/meet1-1.webp', alt: 'an LED heart on a breadboard' },
			{ src: '/img/e/cs/meet1-2.webp', alt: 'a button-triggered LED light show' },
			{ src: '/img/e/cs/meet1-3.webp', alt: 'button traffic lights' },
			{ src: '/img/e/cs/meet1-5.webp', alt: 'LEDs driven through an IC' },
			{ src: '/img/e/cs/meet1-6.webp', alt: 'an Arduino LED heart' }
		],
		links: [
			{ label: 'cs.fraserhacks.dev', url: 'https://cs.fraserhacks.dev' },
			{ label: '@jfss_cs', url: 'https://instagram.com/jfss_cs' }
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
		theme: 'polaris',
		logo: '/img/e/horizons-logo.svg',
		cover: '/img/p/bingbang/intro.webp',
		summary: "one of Hack Club's six Horizons hackathons, \"in the heart of Canada\".",
		body: `Hack Club ran six Horizons hackathons around the world this summer, and Polaris was the Toronto one. we built Bing Bang Dream Adventure, a rhythm game played with two Ultisense boards as flick controllers, and took **2nd place** for the hardware/software integration.`,
		award: '2nd place',
		projects: ['bing-bang-dream-adventure', 'ultisense'],
		links: [{ label: 'horizons.hackclub.com', url: 'https://horizons.hackclub.com' }]
	},
	{
		slug: 'outpost',
		name: 'Outpost',
		kind: 'went',
		when: 'Jul 14–20, 2026',
		place: 'San Francisco + Open Sauce',
		theme: 'outpost',
		logo: '/img/e/outpost-logo.webp',
		cover: '/img/s/outpost.webp',
		summary: 'a 4-day Hack Club hardware hackathon, then two days showing projects at Open Sauce.',
		body: `Hack Club × Open Sauce: you qualify by making hardware projects, then fly to San Francisco for a 4-day hackathon followed by a 2-day showcase at Open Sauce, the maker faire run by YouTubers.`,
		links: [{ label: 'outpost.hackclub.com', url: 'https://outpost.hackclub.com' }]
	},
	{
		slug: 'fallout',
		name: 'Fallout',
		kind: 'went',
		when: 'Jul 1–7, 2026',
		place: 'Shenzhen, China',
		theme: 'fallout',
		logo: '/img/e/fallout/logo.webp',
		cover: '/img/s/fallout.webp',
		decor: ['/img/e/fallout/comic.webp', '/img/e/fallout/person.webp'],
		summary: 'build hardware for free, then go to Shenzhen. a week in the hardware capital of the world.',
		body: `Hack Club's hardware program that ends in Shenzhen, next to the markets and fabs that make most of the world's electronics. i gave a talk on the Aura Generator to about 150 people at the kickoff, made a whale NFC badge, speedran Keychain-tris before the deadline, and built Jollipop, which won **funniest project** (one of only five awards).`,
		award: 'funniest project',
		projects: ['jollipop', 'keychain-tris', 'whale-badge', 'aura-generator'],
		links: [{ label: 'fallout.hackclub.com', url: 'https://fallout.hackclub.com' }]
	},
	{
		slug: 'alleyway',
		name: 'Alleyway × Hackaday Supercon',
		kind: 'went',
		when: 'Nov 2025',
		place: 'Pasadena, California',
		theme: 'alleyway',
		logo: '/img/e/hackaday-logo.svg',
		cover: '/img/p/toparm/build.webp',
		summary: 'Hack Club picked me to bring projects to Hackaday Supercon.',
		body: `Supercon is Hackaday's hardware conference: badge hacking, talks and a lot of very serious hardware people. through Hack Club's Alleyway i brought TopArm and the Aura Generator. TopArm took **3rd in the costume contest** and won me a $5,000 lab-grade waveform generator, and the best part was walking around talking to people about what they build for work, then dragging them back to see mine.

a talk there, Cyril Engmann's "What Makes a Robot Feel Alive?", got me hooked on actuators and legged robots.`,
		award: '3rd, costume contest ($5k lab instrument)',
		projects: ['toparm', 'aura-generator'],
		links: [{ label: 'Hackaday Supercon', url: 'https://hackaday.io/superconference/' }]
	},
	{
		slug: 'hackthenorth',
		name: 'Hack the North 2025',
		kind: 'went',
		when: 'Sep 2025',
		place: 'University of Waterloo',
		theme: 'hackthenorth',
		cover: '/img/s/htn.webp',
		decor: ['/img/e/htn/polaroid.webp', '/img/e/htn/cassette.webp', '/img/e/htn/vinyl.webp'],
		summary: "Canada's biggest hackathon.",
		stats: [
			{ n: '1000+', label: 'hackers' },
			{ n: '36', label: 'hours' }
		],
		body: `three days at Waterloo. James Lian and i built GuessrGo, a geo-guessing app in React Native.`,
		projects: ['guessrgo']
	},
	{
		slug: 'undercity',
		name: 'Undercity',
		kind: 'went',
		when: 'Jul 11–14, 2025',
		place: 'GitHub HQ, San Francisco',
		theme: 'undercity',
		logo: '/img/e/undercity-logo.svg',
		cover: '/img/p/arcadetris/start.webp',
		summary: "Hack Club's 4-day hardware hackathon at GitHub HQ.",
		body: `to qualify you ship hardware through Hack Club Highway. i submitted Minicopter and FancyBoard, after making a Hackpad earlier that summer. then four days at GitHub HQ, where i built Arcadetris from scratch: two-player Tetris on an LED matrix in a cardboard cabinet. `,
		projects: ['arcadetris', 'minicopter', 'fancyboard', 'hackpad'],
		links: [{ label: 'undercity.hackclub.com', url: 'https://undercity.hackclub.com' }]
	},
	{
		slug: 'uottahack',
		name: 'uOttaHack 7',
		kind: 'went',
		when: 'Jan 2025',
		place: 'University of Ottawa',
		theme: 'uottahack',
		cover: '/img/s/uottahack.webp',
		summary: 'a weekend hackathon in Ottawa.',
		body: `Jason Chou and i built DiaryMuse, which turns diary entries into songs you can share.`,
		projects: ['diarymuse']
	},
	{
		slug: 'byssi',
		name: 'BYSSI hackathon',
		kind: 'went',
		when: 'summer 2023',
		place: 'Ontario',
		theme: 'plain',
		summary: 'my first hackathon, the summer before high school. we won.',
		body: `my first hackathon, the summer after grade 8. we took **1st place**.`,
		award: '1st place'
	}
];

/** smaller things that don't need their own page */
export const otherEvents = [
	{ name: 'UTMUN', note: 'Model UN at U of T', when: '2025 – 26' },
	{ name: 'KingsMUN', note: 'Model UN', when: '2025 – 26' },
	{ name: 'UofT Science Unlimited', note: 'summer camp; finalist with an AI garbage sorter', when: '2025' }
];

export const eventBySlug = (slug: string) => events.find((e) => e.slug === slug);
