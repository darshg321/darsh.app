// Every project, in the order i'd want someone to see them: most time and care first,
// what i'm working on now near the top. Status is honest: "designed" means CAD/PCB only.

export type Status = 'ongoing' | 'built' | 'shipped' | 'assembled' | 'designed' | 'demoed';
export type Kind = 'hardware' | 'software' | 'both';
export type Mask = 'green' | 'blue' | 'purple' | 'red' | 'yellow' | 'white' | 'black';

export interface Img {
	src: string;
	alt: string;
}

export interface Project {
	slug: string;
	title: string;
	tagline: string;
	when: string;
	year: number;
	kind: Kind;
	status: Status;
	mask: Mask;
	cover: string;
	/** cutouts float on the board; photos and renders get cropped to fill */
	photo?: boolean;
	stack: string[];
	with?: { name: string; url?: string }[];
	/** slugs from events.ts */
	events?: string[];
	award?: string;
	hours?: number;
	links: { label: string; url: string }[];
	body: string;
	gallery: Img[];
	video?: { src: string; caption: string };
	youtube?: { id: string; caption: string };
	/** file in src/lib/content/journals, and the repo its relative images live in */
	journal?: { file: string; repo: string };
	writeup?: string;
}

export const statusLabel: Record<Status, string> = {
	ongoing: 'working on it now',
	built: 'built',
	shipped: 'shipped',
	assembled: 'assembled, debugging',
	designed: 'designed, not built',
	demoed: 'demoed at a hackathon'
};

const gh = (repo: string) => ({ label: 'github', url: `https://github.com/${repo}` });
const g = (dir: string, ...items: [string, string][]): Img[] =>
	items.map(([file, alt]) => ({ src: `/img/p/${dir}/${file}.webp`, alt }));

export const projects: Project[] = [
	{
		slug: 'icey',
		title: 'Icey',
		tagline: 'an Uno-sized FPGA board, a course to learn it with, and now a chip',
		when: 'Jul 2026 – now',
		year: 2026,
		kind: 'both',
		status: 'ongoing',
		mask: 'green',
		cover: '/img/cut/icey.webp',
		photo: true,
		stack: ['iCE40UP5K', 'Verilog', 'KiCad', 'Yosys', 'nextpnr', 'JavaScript'],
		events: ['offscript'],
		award: 'OSHWA certified open source hardware (CA000078)',
		links: [
			gh('darshg321/Icey'),
			{ label: 'OSHWA certificate', url: 'https://certification.oshwa.org/ca000078.html' }
		],
		body: `Icey is an FPGA board the size of an Arduino Uno, built around the iCE40UP5K. it has a VGA port for display out, USB host for keyboards, USB-C programming over JTAG or UART, QSPI flash, a microSD slot, broken-out headers, LEDs and buttons. PCBWay fabbed and assembled it.

the board is half of it. the other half is **Learn FPGAs with Icey**, a course that runs the real open-source toolchain in your browser: Yosys compiled to WebAssembly synthesises your Verilog, then the exact gates get simulated on a virtual board with LEDs, switches, a 7-segment display, a serial terminal and a VGA monitor. hit "build for Icey" and it places and routes for the real chip and flashes it over WebUSB. 25 lessons go from lighting an LED to building your own CPU and an Apple I running the Woz Monitor.

there's also a 14-session **make a chip** track that ends with a VGA graphics chip laid out in GF180 for Hack Club's shared #asic tapeout. that's what i'm working on right now: my own design for the shared chip. it hasn't been on Icey yet; that's the next step before it goes to the fab.

i brought the board to Off-Script so people could see it up close. it's certified open source hardware by OSHWA.`,
		gallery: g(
			'icey',
			['photo', 'the assembled Icey board with a VGA connector'],
			['render', 'render of Icey from above'],
			['render2', 'render of Icey at an angle'],
			['back', 'the back of the board, with the open source hardware mark'],
			['pcb', 'the PCB layout in KiCad'],
			['schematic', 'the schematic']
		),
		video: {
			src: '/video/icey-life.mp4',
			caption: "conway's game of life on Icey: 128 cells computed every clock"
		}
	},
	{
		slug: 'dayoneof',
		title: 'Dayoneof',
		tagline: 'a Hack Club program: post a short video every day for a month',
		when: 'Sep 2026 – now',
		year: 2026,
		kind: 'software',
		status: 'ongoing',
		mask: 'yellow',
		cover: '/img/e/dayoneof-reel.webp',
		photo: true,
		stack: ['SvelteKit', 'Svelte 5', 'Airtable', 'Slack API', 'Hack Club Auth'],
		with: [{ name: 'Celestial' }, { name: 'Zach Latta' }],
		events: ['dayoneof'],
		links: [
			{ label: 'dayoneof.hackclub.com', url: 'https://dayoneof.hackclub.com' },
			gh('hackclub/dayoneof')
		],
		body: `Dayoneof is a Hack Club "you ship, we ship" program for teens: post a short-form video every day for a month, about anything at all, keep the streak alive and earn prizes. i run it with Celestial and Zach Latta, and i built the site and everything behind it (44 commits so far).

under the hood: SvelteKit and Svelte 5, Airtable as the database, sign-in with Hack Club Auth, and a Slack bot on the events API. cron jobs settle each person's day (a streak freeze or a break), post the leaderboard at 9pm, and DM anyone with a streak who hasn't posted by 8pm their time. views come in from Hack Club's unified socials database, and there's a gallery, an FAQ and a page for every participant.

i also added it to Hack Club's DNS and put a promo card for it on Stardance. my own daily videos go up on [@darsh.builds](https://www.instagram.com/darsh.builds/).`,
		gallery: [
			{ src: '/img/s/dayoneof.webp', alt: 'the Dayoneof landing page' },
			{ src: '/img/e/dayoneof-reel.webp', alt: 'cover of my day one reel' }
		]
	},
	{
		slug: 'fraserhacks-sites',
		title: 'FraserHacks Platform',
		tagline: 'the check-in system and the sites for three years of our hackathon',
		when: '2024 – now',
		year: 2026,
		kind: 'software',
		status: 'shipped',
		mask: 'purple',
		cover: '/img/s/fh26.webp',
		photo: true,
		stack: ['React', 'MongoDB', 'Google Cloud', 'Vite', 'Tailwind'],
		with: [{ name: 'Jason Chou', url: 'https://github.com/JasonChou0105' }],
		events: ['fraserhacks'],
		links: [
			{ label: 'fraserhacks.dev', url: 'https://fraserhacks.dev' },
			{ label: '2026.fraserhacks.dev', url: 'https://2026.fraserhacks.dev' },
			gh('FraserHacks/FraserHacks2027'),
			gh('FraserHacks/FH26'),
			gh('FraserHacks/fraserhacks24website')
		],
		body: `for FraserHacks 2024 i was on the tech team: i built the check-in system in React and MongoDB on Google Cloud, and it handled 100+ attendees checking in live without going down. i also worked on the site and spent the day helping hackers get unstuck.

since then i've co-led the event, and Jason Chou and i made the sites: the space-themed one for FraserHacks 2026 (83 hackers, 28 projects, 7 judges), and the sunset one for 2027 that's taking interest forms now.`,
		gallery: [
			{ src: '/img/s/fh26.webp', alt: 'the FraserHacks 2026 site' },
			{ src: '/img/s/fh27.webp', alt: 'the FraserHacks 2027 site' },
			{ src: '/img/e/fh24/3.webp', alt: 'hackers at FraserHacks 2024' }
		]
	},
	{
		slug: 'ultisense',
		title: 'Ultisense',
		tagline: 'an ESP32-S3 dev board made to sense as many things as possible',
		when: 'Nov 2025 – Mar 2026',
		year: 2026,
		kind: 'hardware',
		status: 'built',
		mask: 'purple',
		cover: '/img/cut/ultisense.webp',
		stack: ['ESP32-S3', 'KiCad', 'C++', 'I²C', 'SPI'],
		events: ['polaris'],
		links: [gh('darshg321/Ultisense')],
		body: `my favourite board so far. Ultisense packs an NFC antenna, a 6-DOF IMU, a VOC air sensor, humidity and temperature, colour and ambient light, a magnetometer, a microphone, two speakers and a microSD slot onto one ESP32-S3 board.

when the boards arrived, none of the I²C sensors would talk. the folks in the 1BitSquared Discord pointed me at a logic analyzer (which i happened to have from an old AliExpress order), and by the time i'd soldered headers on to probe it, it had started working. after a fight with Arduino board definitions i made a custom board package, then got a real VOC index out of the gas sensor with Sensirion's algorithm: about 90–120 in my room, and up near 200 when i breathe on it, which is exactly what it should do.

two Ultisense boards later became the flick controllers for Bing Bang Dream Adventure.`,
		gallery: g(
			'ultisense',
			['build', 'the assembled purple Ultisense board'],
			['pcb3d', '3D render of Ultisense'],
			['pcb', 'the PCB layout'],
			['schematic', 'the schematic']
		),
		journal: { file: 'ultisense', repo: 'darshg321/Ultisense' }
	},
	{
		slug: 'bing-bang-dream-adventure',
		title: 'Bing Bang Dream Adventure',
		tagline: 'a Bad Apple!! rhythm game you play by flicking two IMU controllers',
		when: 'Aug – Oct 2026',
		year: 2026,
		kind: 'both',
		status: 'built',
		mask: 'blue',
		cover: '/img/p/bingbang/intro.webp',
		photo: true,
		stack: ['Godot 4', 'Python', 'Arduino', 'ESP32-S3', 'Ultisense'],
		with: [{ name: '@flyingf15h', url: 'https://github.com/flyingf15h' }],
		events: ['polaris'],
		award: '2nd place at Horizons Polaris',
		links: [gh('flyingf15h/bing-bang-dream-adventure')],
		body: `a rhythm game inspired by Project Sekai and Beat Saber, synced to Bad Apple!!. notes come in from six directions and you flick the controllers to hit them. the controllers are Ultisense boards, read through their IMUs.

Godot can't open a serial port, so a small Python bridge reads the boards over USB or WiFi, decides what counts as a flick, and posts each one to the game over localhost. there's a beatmap select, a results screen and a leaderboard page that reads the scores the game writes.

we made it at Horizons Polaris in Toronto and took 2nd place, mostly for how tightly the hardware and software fit together.`,
		gallery: g(
			'bingbang',
			['intro', 'the title screen'],
			['gameplay', 'gameplay, with notes coming in on a ring'],
			['select', 'picking a beatmap'],
			['results', 'the results screen'],
			['leaderboard', 'the leaderboard page']
		)
	},
	{
		slug: 'minicopter',
		title: 'Minicopter',
		tagline: 'an ESP32 drone where the PCB is the frame',
		when: 'Jul – Aug 2025',
		year: 2025,
		kind: 'hardware',
		status: 'assembled',
		mask: 'green',
		cover: '/img/cut/minicopter.webp',
		stack: ['ESP32-C3', 'KiCad', 'MicroPython', 'PID'],
		events: ['undercity'],
		hours: 35,
		links: [gh('darshg321/Minicopter')],
		body: `a small quadcopter where the circuit board is the frame: ESP32-C3, brushed motors on MOSFETs, LiPo charging over USB-C, an LDO and an IMU, with PID control and the goal of flying a set path from my room to the couch and landing.

i designed it, ordered it and soldered it, then spent a lot of hours with a multimeter. the LDO's enable pin was held low by the wrong resistor, the USB CC pull-downs weren't soldered properly, and the USB ESD chip was too small to rework by hand. i got the 3.3 V rail up off a bench supply, but it hasn't flown yet. everything i learned went into Mini-Air.

Minicopter was one of my submissions to Hack Club Highway, which got me to Undercity.`,
		gallery: g(
			'minicopter',
			['render', 'render of the Minicopter board with motors'],
			['board1', 'the assembled board'],
			['board2', 'the board with its ESP32 module soldered'],
			['board3', 'debugging the board'],
			['bench', 'the board powered from a bench supply'],
			['cad', 'the CAD'],
			['pcb3d', 'the 3D PCB']
		),
		journal: { file: 'minicopter', repo: 'darshg321/Minicopter' }
	},
	{
		slug: 'arcadetris',
		title: 'Arcadetris',
		tagline: 'a two-player LED matrix Tetris arcade, built at Undercity',
		when: 'Jul 2025',
		year: 2025,
		kind: 'hardware',
		status: 'built',
		mask: 'red',
		cover: '/img/cut/arcadetris.webp',
		stack: ['C++', 'Arduino', 'LED matrix'],
		events: ['undercity'],
		links: [gh('darshg321/arcadetris')],
		body: `built from scratch in four days at Hack Club Undercity at GitHub HQ: an LED matrix in a cardboard cabinet, two sets of controls, and two-player Tetris where each side keeps its own score. it worked by the end, and it's still on the project wall of the CS Club site.`,
		gallery: g(
			'arcadetris',
			['start', 'Arcadetris showing PRESS ANY BUTTON TO START'],
			['working', 'Arcadetris running two games at once'],
			['posture', 'building it at Undercity'],
			['posture2', 'more building at Undercity'],
			['build', 'the inside of the cabinet']
		)
	},
	{
		slug: 'steady',
		title: 'Steady',
		tagline: 'a double inverted pendulum that balances itself on a rack and pinion',
		when: 'Jun 2026',
		year: 2026,
		kind: 'hardware',
		status: 'designed',
		mask: 'blue',
		cover: '/img/cut/steady.webp',
		stack: ['Teensy 4.1', 'C++', 'Onshape', 'Blender', 'control'],
		links: [
			gh('darshg321/Steady'),
			{
				label: 'onshape',
				url: 'https://cad.onshape.com/documents/bdb7c14a6d723ab75184049c/w/bbc527d9ae0d21a2dad8c051/e/2942f8800dd81ea3c4a498f0'
			}
		],
		body: `a cart on a linear rail, driven by a 12 V motor through a rack and pinion, balancing two hinged arms on top of each other. a Teensy 4.1 will run the control loop at 2 kHz, two time-of-flight sensors find the cart, two AS5600 magnetic encoders read both joints, and a DRV8874 drives the motor.

i made it to get properly good at CAD, and the assembly works: everything fits with zero overlap, and the relations move the way the real thing will. i learned where bearings go, how to mount an encoder in an awkward spot (the second joint), and that real people don't use hex shafts. i printed my own anyway. the design, BOM and wiring are done; it isn't built yet.`,
		gallery: g(
			'steady',
			['render', 'render of Steady'],
			['cad', 'the CAD assembly'],
			['planning', 'planning sketches'],
			['wiring', 'the wiring diagram']
		),
		journal: { file: 'steady', repo: 'darshg321/Steady' }
	},
	{
		slug: 'spotiplay',
		title: 'Spotiplay',
		tagline: 'a little wifi "now playing" screen for Spotify',
		when: 'Apr – Jul 2026',
		year: 2026,
		kind: 'hardware',
		status: 'built',
		mask: 'purple',
		cover: '/img/cut/spotiplay.webp',
		stack: ['ESP32-C3', 'C++', 'Onshape', 'Spotify API'],
		links: [gh('darshg321/Spotiplay')],
		body: `an ESP32-C3 and a 1.8" TFT in a 3D-printed case. it joins wifi, signs in to Spotify, polls what's playing and draws it, and three Cherry MX switches on top do play, pause and skip. it sits on my desk.

i made it after seeing a Spotify car display, and as a reason to relearn CAD before bigger projects. when Spotify changed their API, i sent a fix for the guide to Hack Club Stasis.`,
		gallery: g(
			'spotiplay',
			['build', 'Spotiplay on my desk showing a song'],
			['cad', 'the case in CAD'],
			['wiring', 'the wiring diagram']
		),
		journal: { file: 'spotiplay', repo: 'darshg321/Spotiplay' }
	},
	{
		slug: 'easyrp2040',
		title: 'EasyRP2040',
		tagline: 'my own Pico-format RP2040 board, with debug LEDs and an IMU',
		when: 'Oct 2025',
		year: 2025,
		kind: 'hardware',
		status: 'built',
		mask: 'blue',
		cover: '/img/cut/easyrp2040.webp',
		stack: ['RP2040', 'KiCad', 'JLCPCB'],
		links: [gh('darshg321/EasyRP2040')],
		body: `a Raspberry Pi Pico-sized RP2040 board with USB-C, debug LEDs and a 9-DOF IMU, and silkscreen art i drew in Excalidraw. my first full routing job taught me to place ICs first, then capacitors, and to keep vias and back-side signals down. JLC assembled most of it; the IMU was standard-only and would've added about $50, so i placed that one myself with solder paste and a hot air gun.`,
		gallery: g(
			'easyrp2040',
			['photo', 'the assembled EasyRP2040 on a table'],
			['build', 'render of EasyRP2040'],
			['pcb', 'the PCB layout']
		),
		journal: { file: 'easyrp2040', repo: 'darshg321/EasyRP2040' }
	},
	{
		slug: 'fume-extractor',
		title: 'Fume-Extractor',
		tagline: 'pulls solder fumes away from my face, quietly',
		when: 'May – Sep 2026',
		year: 2026,
		kind: 'hardware',
		status: 'built',
		mask: 'yellow',
		cover: '/img/cut/fume.webp',
		stack: ['ESP32-S3', 'Onshape', 'PWM'],
		links: [gh('darshg321/Fume-Extractor')],
		body: `i kept breathing in solder fumes, which probably wasn't great for my brain cells. so: a Noctua fan and a carbon filter in a 3D-printed case, with an ESP32-S3 setting the fan speed over PWM, all off a 12 V barrel jack. it was also a simple project to practise CAD on, with heat-set inserts and a back plate that closes it up.`,
		gallery: g(
			'fume',
			['build', 'the printed fume extractor with its carbon filter'],
			['poster', 'the project zine'],
			['render', 'a render of it floating in front of the night sky'],
			['cad', 'the CAD'],
			['wiring', 'the wiring diagram']
		),
		youtube: { id: 'nK1IK6mNC_M', caption: 'demo' }
	},
	{
		slug: 'boreas',
		title: 'Boreas',
		tagline: 'an STM32F7 flight computer for a rocket',
		when: 'Jun – Aug 2026',
		year: 2026,
		kind: 'hardware',
		status: 'designed',
		mask: 'white',
		cover: '/img/cut/boreas.webp',
		photo: true,
		stack: ['STM32F7', 'bare-metal C', 'KiCad', 'Make'],
		links: [gh('darshg321/Boreas')],
		body: `a flight computer for my first rocket, which i'd like to get past 1 km. an STM32F7 reads an IMU, a barometer and a magnetometer and logs it all to an SD card, so i can graph each flight and work out what to fix. it has three servo outputs, a buzzer, LEDs, buttons and LiPo charging over USB-C.

the firmware is register-level C with my own Makefile, flashed over ST-Link with \`make flash\`. it was mostly practice at reading datasheets and routing a dense board, and a step towards my own drone flight controller.`,
		gallery: g(
			'boreas',
			['poster', 'the Boreas zine'],
			['render', 'render of the board'],
			['pcb-jlc', 'the board in JLCPCB'],
			['pcb', 'the PCB layout'],
			['schematic', 'the schematic']
		)
	},
	{
		slug: 'mini-air',
		title: 'Mini-Air',
		tagline: 'an analog FPV drone on my own Betaflight flight controller',
		when: 'Jun – Aug 2026',
		year: 2026,
		kind: 'hardware',
		status: 'designed',
		mask: 'green',
		cover: '/img/cut/miniair.webp',
		stack: ['Betaflight', 'Bluejay', 'ELRS', 'KiCad', 'Onshape'],
		links: [gh('darshg321/Mini-Air')],
		body: `the drone i actually want to fly. the flight controller is mine: it runs Betaflight and has an on-board OSD, IMU, barometer and SPI flash. around it go a Bluejay ESC, ELRS radio, GPS and an analog camera on a 2S battery with 2-inch props, about 124 g, built for speed. everything is analog to keep latency low (and because 480p FPV footage with an OSD looks great). the VTX, GPS, receiver and ESC are all modular, so any of them can be swapped.

it's Minicopter's lessons, done properly.`,
		gallery: g(
			'miniair',
			['fc', 'the Mini-Air flight controller'],
			['drone', 'render of the whole drone'],
			['fc2', 'the flight controller at an angle'],
			['poster', 'the Mini-Air zine'],
			['pcb', 'the PCB layout'],
			['schematic', 'the schematic']
		)
	},
	{
		slug: 'toparm',
		title: 'TopArm',
		tagline: 'a robot arm on a top hat',
		when: 'Oct 2025',
		year: 2025,
		kind: 'hardware',
		status: 'built',
		mask: 'red',
		cover: '/img/cut/toparm.webp',
		stack: ['Arduino', 'PCA9685', 'SG90 servos', '3D printing'],
		events: ['alleyway'],
		award: '3rd in the Hackaday Supercon 2025 costume contest',
		links: [gh('darshg321/TopArm')],
		body: `a 3D-printed arm on micro servos, strapped to a top hat and driven with a joystick, so you can grab things off the top shelf, wave at friends, or just gesture dramatically while you talk. an Arduino and a PCA9685 servo driver run it, and LEDs add some bling.

i wore it at Hackaday Supercon 2025. it took 3rd in the costume contest and won me a $5,000 lab-grade waveform generator, and it got a lot of serious hardware people (including Framework's CEO) to stop and talk.`,
		gallery: g('toparm', ['build', 'TopArm on the hat'], ['concept', 'the first sketch']),
		youtube: { id: 'ynK_KVnrBHw', caption: 'TopArm moving' },
		writeup: 'toparm'
	},
	{
		slug: 'aura-generator',
		title: 'Aura Generator',
		tagline: 'a constant cloud of mist that follows you around',
		when: 'Sep – Oct 2025',
		year: 2025,
		kind: 'hardware',
		status: 'built',
		mask: 'purple',
		cover: '/img/cut/aura.webp',
		stack: ['ultrasonic atomizers', 'USB-C power'],
		events: ['alleyway', 'fallout'],
		links: [gh('darshg321/aura-generator')],
		body: `stage-smoke energy for everyday life. USB-C ultrasonic atomizers vibrate hundreds of thousands of times a second under water to make mist, fed by cotton wicks from a bottle and powered from a battery bank. it was meant for my shoe, but tape and walking don't mix, so it lived on my hat. dye the water and the mist changes colour.

i brought it to Supercon, and later gave a talk about it to 150 people at the Fallout kickoff.`,
		gallery: g('aura', ['build', 'the Aura Generator on a shoe'], ['concept', 'the first sketch']),
		youtube: { id: 'ul6SC5NBSu8', caption: 'Aura Generator demo' },
		writeup: 'aura'
	},
	{
		slug: 'cs-club-site',
		title: 'CS Club Site',
		tagline: 'cs.fraserhacks.dev: our club, its prizes, and everything members build',
		when: 'Oct 2026 – now',
		year: 2026,
		kind: 'software',
		status: 'shipped',
		mask: 'red',
		cover: '/img/s/cs.webp',
		photo: true,
		stack: ['SvelteKit', 'Supabase', 'Vercel'],
		events: ['csclub'],
		links: [
			{ label: 'cs.fraserhacks.dev', url: 'https://cs.fraserhacks.dev' },
			gh('FraserHacks/csclub')
		],
		body: `the site for the Fraser Comp Sci Club, which i run. it's where people sign up, see the trinket of the week and the prizes, and find the contests we go after: the CCC, Peel Science Fair and Peel Skills, plus the hackathons we go to as a group.

every meeting gets its own page with the whole lesson on it (the LED heart speed build, the Arduino challenges, the Smash tourney with live sign-ups and a bracket), and anything a member builds goes up in the project directory with their name on it, so they can point to it on applications. there's an admin side for editing projects and running brackets, backed by Supabase.`,
		gallery: [
			{ src: '/img/s/cs.webp', alt: 'the CS Club site' },
			{ src: '/img/e/cs/meet1-1.webp', alt: 'an LED heart from meeting 1' },
			{ src: '/img/e/cs/meet1-4.webp', alt: 'an LED cluster and button from meeting 1' }
		]
	},
	{
		slug: 'productite',
		title: 'Productite',
		tagline: 'a student productivity app, on the App Store',
		when: 'Jul – Dec 2024',
		year: 2024,
		kind: 'software',
		status: 'shipped',
		mask: 'white',
		cover: '/img/p/productite/2.webp',
		photo: true,
		stack: ['React Native', 'Expo', 'TypeScript'],
		links: [gh('darshg321/Productite'), { label: 'cider PR', url: 'https://github.com/hackclub/cider/pull/64' }],
		body: `a to-do list, a daily diary, goals and a time tracker with charts of where your day went. i built it in React Native with Expo and took it all the way through App Store submission. it got me an Apple developer licence from Hack Club Cider, $150+ in grants, and a spot in the Hack Club Arcade showcase.`,
		gallery: g(
			'productite',
			['2', 'the time tracker'],
			['1', 'the home screen'],
			['3', 'a pie chart of tracked time']
		),
		video: { src: '/video/productite.mp4', caption: 'a walk through the app' }
	},
	{
		slug: 'jollipop',
		title: 'Jollipop',
		tagline: 'a motorised propeller hat and a comically oversized lollipop',
		when: 'Jul 2026',
		year: 2026,
		kind: 'hardware',
		status: 'built',
		mask: 'yellow',
		cover: '/img/p/jollipop/build.webp',
		stack: ['XIAO RP2040', 'DC motor', 'LED strips'],
		with: [{ name: '@penguinissus', url: 'https://github.com/penguinissus' }],
		events: ['fallout'],
		award: 'funniest project at Fallout',
		links: [gh('penguinissus/jollynator')],
		body: `become jolly today. a XIAO RP2040 spins the propeller, runs rainbow LED strips and plays a tune on a speaker, off AA batteries through a buck converter. life is all sunshine and rainbows. made at Fallout in Shenzhen.`,
		gallery: g('jollipop', ['build', 'the propeller hat'], ['poster', 'the Jollipop poster'], ['wiring', 'the wiring']),
		youtube: { id: 'aW7lNljLVsI', caption: 'the demo video' }
	},
	{
		slug: 'keychain-tris',
		title: 'Keychain-tris',
		tagline: 'Tetris on a keychain, played by tilting it',
		when: 'Jun 2026',
		year: 2026,
		kind: 'hardware',
		status: 'designed',
		mask: 'yellow',
		cover: '/img/cut/keychaintris.webp',
		photo: true,
		stack: ['ESP32-C3', 'MPU6050', 'KiCad'],
		events: ['fallout'],
		links: [gh('darshg321/Keychain-tris')],
		body: `a speedrun before the end of Fallout: an ESP32-C3 board with an OLED, LiPo charging, a 3.3 V LDO and an MPU6050. you play Tetris by moving the whole keychain: slide it right to move right, drop it to drop. six axes means plenty of controls.`,
		gallery: g(
			'keychaintris',
			['poster', 'the Keychain-tris zine'],
			['render', 'render of the board'],
			['pcb', 'the PCB layout'],
			['schematic', 'the schematic']
		)
	},
	{
		slug: 'whale-badge',
		title: 'Whale Badge',
		tagline: 'a whale-shaped NFC badge with an LED eye',
		when: 'Jul 2026',
		year: 2026,
		kind: 'hardware',
		status: 'built',
		mask: 'white',
		cover: '/img/cut/badge.webp',
		stack: ['NT3H2111 NFC', 'KiCad'],
		events: ['fallout'],
		links: [gh('darshg321/fallout-badge')],
		body: `a small whale PCB for Fallout. tap a phone on it and it opens whatever link you've written to it, and it has an LED for the eye. it's also the logo of this site.`,
		gallery: g('badge', ['build', 'the built whale badge'], ['pcb', 'the PCB'])
	},
	{
		slug: 'fancyboard',
		title: 'FancyBoard',
		tagline: 'a hot-swappable, low-profile wireless keyboard',
		when: 'Jul – Oct 2025',
		year: 2025,
		kind: 'hardware',
		status: 'designed',
		mask: 'black',
		cover: '/img/cut/fancyboard.webp',
		stack: ['nice!nano', 'BlueMicro', 'KiCad', 'Fusion 360'],
		events: ['undercity'],
		links: [gh('darshg321/FancyBoard')],
		body: `a full low-profile mechanical keyboard with hot-swap sockets, Bluetooth on a nice!nano, a screen and a plate with art on it. i started on ZMK, found it lowkey gibberish, and switched to BlueMicro. KiCad doesn't export stretched keys properly, so i placed those by hand (it's sooo slow). it was my other Highway submission.`,
		gallery: g('fancyboard', ['render', 'render of FancyBoard'], ['cad', 'the case CAD'], ['pcb', 'the PCB layout']),
		journal: { file: 'fancyboard', repo: 'darshg321/FancyBoard' }
	},
	{
		slug: 'hackpad',
		title: "Darsh's Hackpad",
		tagline: 'a 4-key macropad for playing Tetris',
		when: 'Jun 2025',
		year: 2025,
		kind: 'hardware',
		status: 'built',
		mask: 'green',
		cover: '/img/cut/hackpad.webp',
		photo: true,
		stack: ['XIAO RP2040', 'KMK', 'KiCad'],
		links: [gh('darshg321/darshs-hackpad')],
		body: `one of my first PCBs, made for Hack Club's Hackpad program: four keys on a XIAO RP2040 in a 3D-printed case, with KMK firmware mapping them to the arrow keys for Tetris. it's where the whole hardware thing started.`,
		gallery: g('hackpad', ['built', 'the built hackpad'], ['render', 'render of the hackpad'], ['pcb', 'the PCB'])
	},
	{
		slug: 'crane-advantage',
		title: 'Crane Advantage',
		tagline: 'a crane that shows mechanical advantage with real numbers',
		when: 'May – Jun 2026',
		year: 2026,
		kind: 'hardware',
		status: 'designed',
		mask: 'red',
		cover: '/img/cut/crane.webp',
		stack: ['ESP32-S3', 'Onshape', 'planetary gearbox'],
		links: [gh('darshg321/Crane-Advantage')],
		body: `a 12 V motor lifts a load through my own 10:1 planetary gearbox and a pulley system. force sensors under the motor lever and in the basket, plus an accelerometer, put the input force, the load and its acceleration on an OLED, so you can watch the motor push with less force than the weight it's lifting. made to learn bearings, pulleys and gears.`,
		gallery: g(
			'crane',
			['poster', 'the Crane Advantage zine'],
			['concept', 'the concept sketch'],
			['cad', 'the CAD'],
			['wiring', 'the wiring']
		)
	},
	{
		slug: 'deadline',
		title: 'Deadline',
		tagline: 'a guillotine that creeps towards a stickman as a deadline gets closer',
		when: 'Aug 2026',
		year: 2026,
		kind: 'hardware',
		status: 'designed',
		mask: 'black',
		cover: '/img/cut/deadline.webp',
		stack: ['ESP32-S3', 'NEMA17', 'AS5600', 'Google Calendar API'],
		links: [gh('darshg321/Deadline')],
		body: `putting the dead in deadline. an ESP32-S3 reads my next Google Calendar event, and a NEMA17 on a ball screw moves a little guillotine towards a stickman: 0% of the way at the start, touching him when the time's up. there's no limit switch: it homes by running into the hard stop and catching the stall on the AS5600, which also watches for lost steps.`,
		gallery: g('deadline', ['cad', 'the CAD'], ['plan', 'the plan'], ['wiring', 'the wiring']),
		journal: { file: 'deadline', repo: 'darshg321/Deadline' }
	},
	{
		slug: 'mushu',
		title: 'Mushu',
		tagline: "FRC team 1325's 2025 robot",
		when: '2025 season',
		year: 2025,
		kind: 'hardware',
		status: 'built',
		mask: 'black',
		cover: '/img/cut/mushu.webp',
		stack: ['Java', 'WPILib', 'Limelight', 'PID'],
		with: [{ name: 'FRC Team 1325 Inverse Paradox' }],
		award: 'top 3% of 3,600+ teams worldwide',
		links: [],
		body: `a team robot, so the credit goes to all of 1325. my part was controls: i programmed the elevator and claw with closed-loop PID, and the Limelight pipeline that finds game pieces. Mushu ranked #13 of 130 in Ontario, and the team finished the season in the top 3% worldwide. more about the team on the [about page](/about#frc).`,
		gallery: [{ src: '/img/cut/mushu.webp', alt: 'Mushu, the 2025 robot' }]
	},
	{
		slug: 'guessrgo',
		title: 'GuessrGo',
		tagline: 'a geo-guessing app, made at Hack the North 2025',
		when: 'Sep 2025',
		year: 2025,
		kind: 'software',
		status: 'demoed',
		mask: 'blue',
		cover: '/img/p/guessrgo/map.webp',
		photo: true,
		stack: ['React Native', 'Expo'],
		with: [{ name: 'James Lian', url: 'https://github.com/James-Lian' }],
		events: ['hackthenorth'],
		links: [gh('James-Lian/GuessrGo')],
		body: `a geo-guessing game in Expo and React Native, made with James Lian over a weekend at Hack the North. i wrote the geo side (23 commits' worth).`,
		gallery: [{ src: '/img/p/guessrgo/map.webp', alt: 'an illustration of guessing where a photo was taken' }]
	},
	{
		slug: 'diarymuse',
		title: 'DiaryMuse',
		tagline: 'turns diary entries into songs you can share',
		when: 'Jan 2025',
		year: 2025,
		kind: 'software',
		status: 'demoed',
		mask: 'purple',
		cover: '/img/p/diarymuse/1.webp',
		photo: true,
		stack: ['TypeScript', 'Python'],
		with: [{ name: 'Jason Chou', url: 'https://github.com/JasonChou0105' }],
		events: ['uottahack'],
		links: [gh('JasonChou0105/DiaryMuse')],
		body: `our uOttaHack 7 project. write about your day and DiaryMuse turns the entry into a personal song, then you can share it with friends on a little social feed, where people like, comment and remix each other's songs.`,
		gallery: g('diarymuse', ['1', 'writing a diary entry'], ['2', 'the song page'], ['3', 'the social feed'])
	},
	{
		slug: 'bench',
		title: 'Bench',
		tagline: 'tracks every part and tool i have at home',
		when: 'Jul 2026',
		year: 2026,
		kind: 'software',
		status: 'shipped',
		mask: 'green',
		cover: '/img/p/bench/demo.webp',
		photo: true,
		stack: ['TypeScript', 'Firestore', 'Google sign-in'],
		links: [{ label: 'hardwarebench.vercel.app', url: 'https://hardwarebench.vercel.app' }, gh('darshg321/bench')],
		body: `a personal hardware inventory, so i always know what i have and what i need next. paste a screenshot of a store page and it fills in the part and its specs; it can also suggest projects from what's in the drawers. it syncs across devices through Firestore.`,
		gallery: g('bench', ['demo', 'the bench inventory'])
	},
	{
		slug: 'submerge',
		title: 'Submerge',
		tagline: 'a late-night koi pond theme for KDE Plasma',
		when: 'Oct 2026',
		year: 2026,
		kind: 'software',
		status: 'shipped',
		mask: 'blue',
		cover: '/img/s/submerge.webp',
		photo: true,
		stack: ['QML', 'Qt 5', 'KDE Plasma'],
		with: [{ name: '@flyingf15h', url: 'https://github.com/flyingf15h' }],
		links: [gh('flyingf15h/submerge')],
		body: `@flyingf15h made a live koi pond wallpaper and HUD for Plasma 6, and i ported it to Plasma 5.27 for my laptop. then i kept going: power and GPU readouts, a battery readout, a lock and sign-in screen, pause modes and a big cut in GPU use.`,
		gallery: [{ src: '/img/s/submerge.webp', alt: 'the Submerge desktop' }]
	},
	{
		slug: 'win-convert',
		title: 'Win-Convert',
		tagline: 'right-click any image in Windows to convert or compress it',
		when: 'Aug 2026',
		year: 2026,
		kind: 'software',
		status: 'shipped',
		mask: 'white',
		cover: '/img/p/winconvert/demo.webp',
		photo: true,
		stack: ['C#', 'PowerShell'],
		links: [gh('darshg321/Win-Convert')],
		body: `a right-click menu for Explorer that converts, compresses or optimises images between PNG, JPG, WEBP, HEIC, PDF and SVG. all six work both ways.`,
		gallery: g('winconvert', ['demo', 'the Win-Convert right-click menu'])
	},
	{
		slug: 'claudle',
		title: 'Claudle',
		tagline: 'DM a Discord bot to drive my PC with Claude Code',
		when: 'Aug – Oct 2026',
		year: 2026,
		kind: 'software',
		status: 'shipped',
		mask: 'black',
		cover: '/img/p/claudle/chat.webp',
		photo: true,
		stack: ['Python', 'discord.py', 'Claude Code'],
		links: [gh('darshg321/claudle')],
		body: `send it a message and it runs it as a Claude Code prompt in a persistent session on my Windows PC, streaming tool calls back and posting a screenshot when it's done. there are direct commands for the shell, screen, mouse and keyboard, files, processes and power too. only my account can talk to it; everyone else is ignored before a byte is parsed.`,
		gallery: [{ src: '/img/p/claudle/chat.webp', alt: 'an illustration of a Discord DM with Claudle' }]
	},
	{
		slug: 'face-recognizer',
		title: 'Face-Recognizer',
		tagline: 'recognises faces in a live video stream from a folder of photos',
		when: '2023, updated 2026',
		year: 2023,
		kind: 'software',
		status: 'shipped',
		mask: 'blue',
		cover: '/img/p/face/frame.webp',
		photo: true,
		stack: ['Python', 'MTCNN', 'PyQt5'],
		links: [gh('darshg321/Face-Recognizer')],
		body: `MTCNN finds faces in each frame, then face embeddings are compared against a folder of known people, with adjustable confidence and match thresholds. i first wrote it in 2023 and came back in 2026 to swap tkinter for a threaded PyQt5 video stream with live settings, plus a slideshow on how face recognition works and where it's biased.`,
		gallery: [{ src: '/img/p/face/frame.webp', alt: 'an illustration of a face being detected and matched' }]
	}
];

/** what i actually build with */
export const bench = [
	{ title: 'KiCad', note: 'every PCB on this site, from the hackpad to Icey', when: 'since 2025' },
	{ title: 'Onshape', note: 'CAD for Steady, Crane Advantage, Deadline, Fume-Extractor and Spotiplay', when: 'since 2026' },
	{ title: 'Fusion 360 + SolidWorks', note: 'before Onshape: FancyBoard, Minicopter and my first cases', when: '2025' },
	{ title: 'JLCPCB + PCBWay', note: 'fab and assembly. PCBWay made Icey', when: 'since 2025' },
	{ title: '3D printer', note: 'every case and mount, TopArm, and my own (very fake) hex shafts', when: 'always on' },
	{ title: 'solder paste + hot air', note: "for the parts the fab won't place, like EasyRP2040's IMU", when: 'since 2025' },
	{ title: 'logic analyzer', note: 'what finally got Ultisense talking over I²C', when: 'since 2026' },
	{ title: 'bench supply + multimeter', note: 'how Minicopter got its 3.3 V rail back', when: 'since 2025' },
	{ title: 'waveform generator', note: 'the $5,000 one TopArm won at Supercon', when: 'since 2025' }
];

/** where it started */
export const early = [
	{ title: 'HelpConnect', note: 'a mutual-aid platform that matches people who want to help with people who need it', when: '2023', url: 'https://github.com/darshg321/HelpConnect' },
	{ title: 'Sprig: Timing', note: 'a game merged into Hack Club Sprig', when: '2024', url: 'https://github.com/hackclub/sprig/pull/1824' },
	{ title: 'ChatGPT Voice Assistant', note: 'Whisper speech-to-text into ChatGPT, deployed on Google App Engine', when: '2023', url: 'https://github.com/darshg321/ChatGPT-Voice-Assistant' },
	{ title: 'ChatGPT-Explain-Linux', note: 'a CLI that explains Linux commands', when: '2023', url: 'https://github.com/darshg321/ChatGPT-Explain-Linux' },
	{ title: 'Dino Game', note: "Chrome's dinosaur game in JavaScript, my first GitHub Page", when: '2023', url: 'https://darshg321.github.io/' },
	{ title: 'Undyne Simulator', note: "Undyne's boss fight from Undertale, with random attacks", when: '2022', url: 'https://github.com/darshg321/Undyne-Simulator' },
	{ title: 'Flappy Bird, Rock Paper Scissors and Apple Catching', note: 'my first Python games, one with a high-score database', when: '2022' },
	{ title: 'Weather and Date Apps', note: 'a C# console weather app and a web page with the day', when: '2023' }
];

export const bySlug = (slug: string) => projects.find((p) => p.slug === slug);
// there are no project pages: a project links straight to its repo (the first GitHub link)
export const repo = (p: Project) => p.links.find((l) => l.url.startsWith('https://github.com/'))?.url;
export const repoOf = (slug: string) => repo(bySlug(slug)!);
