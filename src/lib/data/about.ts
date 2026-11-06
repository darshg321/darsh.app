// Roles, awards and the rest of the about page.

import { repoOf } from './projects';
import { eventUrl } from './events';

export const contact = {
	email: 'darshg321@gmail.com',
	links: [
		{ label: 'github', handle: '@darshg321', url: 'https://github.com/darshg321' },
		{ label: 'instagram', handle: '@darsh.builds', url: 'https://www.instagram.com/darsh.builds/' },
		{ label: 'linkedin', handle: 'in/darshg321', url: 'https://www.linkedin.com/in/darshg321' }
	]
};

export interface Role {
	title: string;
	org: string;
	when: string;
	points: string[];
	/** where to read more on this site */
	href?: string;
	mask: 'green' | 'blue' | 'purple' | 'red' | 'yellow' | 'white' | 'black';
}

export const roles: Role[] = [
	{
		title: 'co-lead',
		org: 'FraserHacks',
		when: '2024 – now',
		mask: 'purple',
		href: eventUrl('fraserhacks'),
		points: [
			'co-lead fundraising, logistics and operations with a 10-person exec team',
			'got sponsorship from Rotary and AP from zero contacts, and handle merch suppliers, judges and workshops',
			'built the 2024 check-in system (React, MongoDB, Google Cloud): 100+ live check-ins, no downtime',
			'mentor teams from idea to live demo; 2026 shipped 28 projects at a 98% submission rate'
		]
	},
	{
		title: 'host',
		org: 'Dayoneof, Hack Club',
		when: 'Sep 2026 – now',
		mask: 'yellow',
		href: eventUrl('dayoneof'),
		points: [
			'run a Hack Club program for teens to post a short video every day for a month',
			'built the site, streak tracking and Slack bot in SvelteKit, Airtable and Hack Club Auth'
		]
	},
	{
		title: 'controls engineering lead',
		org: 'FRC Team 1325 Inverse Paradox',
		when: 'Sep 2023 – May 2026',
		mask: 'black',
		href: '#frc',
		points: [
			'3rd in Canada of 120+ teams (2026); top 3% worldwide of 3,600+ (2024 and 2025)',
			'built a game-piece detection pipeline in Java on a Limelight, with filtering for changing arena light',
			'closed-loop PID on the elevator and claw, checked against match logs'
		]
	},
	{
		title: 'president',
		org: 'Fraser Comp Sci Club',
		when: 'Oct 2023 – now',
		mask: 'red',
		href: eventUrl('csclub'),
		points: [
			'run a weekly build club for 20+ students: LED speed builds, Arduino, soldering, PCBs',
			'train members for the CCC, Peel Science Fair and Peel Skills, and built the club site'
		]
	},
	{
		title: 'host',
		org: 'Off-Script',
		when: 'Aug – Sep 2026',
		mask: 'white',
		href: eventUrl('offscript'),
		points: [
			'hosted a technical showcase at the University of Waterloo during Waterloo Tech Week',
			'120+ RSVPs, 50+ attendees, 10+ live demos of FPGA, embedded and hardware/software projects'
		]
	},
	{
		title: 'app developer',
		org: 'Productite',
		when: 'Jul – Sep 2024',
		mask: 'blue',
		href: repoOf('productite'),
		points: ['shipped a student productivity app to the iOS App Store; $150+ in grants and a Hack Club Arcade feature']
	},
	{
		title: 'volunteer',
		org: 'Achēv / Youth2Action',
		when: 'Jul 2023 – Apr 2024',
		mask: 'green',
		points: [
			'outreach for the Y2A academic fair: contacted schools and organisations',
			'MC for the University of Windsor Law panel: 100+ participants, 1,800 registered',
			'helped organise a political town hall that drew 80+ students',
			'graduated the program with 180+ volunteer hours'
		]
	}
];

export const frc = {
	team: 'FRC Team 1325 Inverse Paradox',
	body: `i joined 1325 in grade 9 on the programming subteam, writing the intake subsystem and parts of the engineering notebook. in grade 10 i worked as human player and pit crew and programmed the elevator and the rest of the controls on Mushu, and i've since coded the shooter. i led controls engineering and taught an 11-lesson controls curriculum to the junior members: command-based code, PID tuning, and getting from a sensor to an actuator.

i also helped a bunch of local teams with their code (10167, 4015 and 5031 among them), and helped a lot with our alphabot.`,
	results: [
		'3rd in Canada of 120+ teams (2026)',
		'top 3% worldwide of 3,600+ teams (2024 and 2025)',
		'World Championship playoffs',
		'2× Excellence in Engineering award',
		'Mushu: #13 of 130 in Ontario (2025)'
	]
};

export const awards = [
	{ what: 'Outstanding High School Student Science Award', who: 'University of Toronto, Andre Simpson Lab', when: '2025' },
	{ what: '3rd place, costume contest ($5,000 lab instrument)', who: 'Hackaday Supercon, for TopArm', when: '2025', href: repoOf('toparm') },
	{ what: '2nd place', who: 'Horizons Polaris, for bing bang dream adventure', when: '2026', href: eventUrl('polaris') },
	{ what: 'funniest project', who: 'Fallout, for Jollipop', when: '2026', href: eventUrl('fallout') },
	{ what: 'OSHWA open source hardware certification', who: 'for Icey (CA000078)', when: '2026', href: repoOf('icey') },
	{ what: 'silver medal', who: 'Waterloo Junior FLC', when: '2025' },
	{ what: 'finalist', who: 'UofT Science Unlimited, with an AI garbage sorter', when: '2025' },
	{ what: 'Hack Club Arcade showcase + Cider grant', who: 'for Productite', when: '2024', href: repoOf('productite') }
];

export const offline = [
	{ what: 'junior rugby 7s', note: '2nd at ROPSSAA' },
	{ what: 'badminton', note: 'JFSS team, and a lot of it outside school' },
	{ what: 'cross country', note: 'grades 8 and 9' },
	{ what: 'cooking', note: 'my crispy pork belly bao got featured in the Cookwell newsletter' },
	{ what: 'Model UN', note: 'UTMUN and KingsMUN' },
	{ what: 'volunteering', note: 'Mississauga Little Aces, Youth2Action' }
];
