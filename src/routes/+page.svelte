<script lang="ts">
	import ProjectCard from '#lib/components/ProjectCard.svelte';
	import EventCard from '#lib/components/EventCard.svelte';
	import { projects, bySlug, repo, repoOf } from '#lib/data/projects.js';
	import { events, eventUrl } from '#lib/data/events.js';
	import { contact } from '#lib/data/about.js';
	import { themeFonts } from '#lib/fonts.js';

	const featured = ['icey', 'ultisense', 'arcadetris'].map((s) => bySlug(s)!);
	const ran = events.filter((e) => e.kind === 'ran');
	const recent = [
		{ slug: 'icey', note: 'an FPGA board, a course, and now a chip', lang: 'Verilog' },
		{ slug: 'dayoneof', note: 'the site and Slack bot for dayoneof', lang: 'Svelte' },
		{ slug: 'cs-club-site', note: 'our club, its prizes, and what members build', lang: 'Svelte' }
	];
	const built = projects.filter((p) => p.kind !== 'software' && p.status === 'built').length;

	// the site over ssh (term/ in this repo); clicking copies the command
	const command = 'ssh darsh.app';
	let note = $state('');
	async function copy() {
		try {
			await navigator.clipboard.writeText(command);
			note = 'copied! paste it into a terminal';
		} catch {
			note = "couldn't copy, type it into a terminal";
		}
		setTimeout(() => (note = ''), 3500);
	}
</script>

<svelte:head>
	<title>darsh.app</title>
	<meta name="description" content="darsh: a high school student in Mississauga who builds circuit boards, robots and drones, and runs FraserHacks and Dayoneof." />
	<link rel="stylesheet" href={themeFonts} />
</svelte:head>

<section class="about">
	<div>
		<h1 class="big">hiya!</h1>
		<p class="on-sky">
			i'm darsh, a grade 12 student in Mississauga, Ontario, and i make hardware: circuit boards, robots, drones and
			the occasional thing that lights up.
		</p>
		<p class="on-sky">
			i started out writing software in 2022. these days it's mostly KiCad, a soldering iron and a 3D printer, and right
			now an FPGA board and a chip to go with it.
		</p>
		<p class="on-sky">
			i co-lead <a href={eventUrl('fraserhacks')}>FraserHacks</a>, run <a href={eventUrl('dayoneof')}>Dayoneof</a> with Hack Club,
			lead the <a href={eventUrl('csclub')}>CS Club</a>, and led controls on <a href="/about#frc">FRC team 1325</a>.
		</p>
	</div>
	<div class="cast term-wrap">
		<button type="button" class="term" onclick={copy} aria-label="copy the command: {command}">
			<span class="term-bar" aria-hidden="true"><i></i><i></i><i></i><span>darsh@home: ~</span></span>
			<span class="term-body">
				<span class="term-note"># try it !!</span>
				<span class="term-line"><span class="prompt">~ $</span> {command}<span class="cursor" aria-hidden="true"></span></span>
				<span class="term-out">the whole site, in your terminal</span>
			</span>
			<span class="term-hint">
				{note || 'click to copy the command'}
			</span>
		</button>
	</div>
</section>

<div class="live-head">
	<h2>what i'm up to</h2>
	<h2 class="contact-heading">get in touch</h2>
</div>
<section class="live cast" aria-label="what i'm up to">
	<!-- dressed like Dayoneof's own site, same as its event card -->
	<a class="t-dayoneof running" href={eventUrl('dayoneof')}>
		<h3 class="silk">running</h3>
		<p class="big-line">Dayoneof</p>
		<p>a month of daily videos, with Hack Club. i built the site and post mine on @darsh.builds.</p>
		<img class="dayo-mark" src="/img/e/dayoneof-wordmark.webp" alt="" />
	</a>

	<a class="board green bench" href={repoOf('icey')}>
		<h3 class="silk">on the bench</h3>
		<p class="big-line">a chip</p>
		<p>my design for Hack Club's shared GF180 tapeout. next: testing it on Icey before it goes to the fab.</p>
	</a>

	<a class="board purple playing" href={repoOf('bing-bang-dream-adventure')}>
		<h3 class="silk">just finished</h3>
		<div class="playing-body">
			<img src="/img/p/bingbang/gameplay.webp" alt="Bing Bang Dream Adventure mid-song" />
			<div>
				<p class="song">Bing Bang Dream Adventure</p>
				<p class="note">a rhythm game played by flicking two Ultisense boards. 2nd at Horizons Polaris</p>
			</div>
		</div>
	</a>

	<div class="board blue github">
		<h3 class="silk">recently on github</h3>
		<ul class="repos">
			{#each recent as r}
				{@const p = bySlug(r.slug)!}
				<li><a href={repo(p)}>{p.title}</a><p>{r.note}</p><span class="lang">{r.lang}</span></li>
			{/each}
		</ul>
	</div>

	<div class="contacts">
		<a class="board red contact" href="mailto:{contact.email}"><strong>email</strong><span>{contact.email}</span></a>
		{#each contact.links as l, i}
			<a class="board {['black', 'purple', 'blue'][i]} contact" href={l.url}><strong>{l.label}</strong><span>{l.handle}</span></a>
		{/each}
	</div>
</section>

<section class="section">
	<div class="section-head">
		<h2>(some) things i've made</h2>
		<a class="sky-link" href="/projects">all {projects.length} projects</a>
	</div>
	<div class="row3 cast">
		{#each featured as p}<ProjectCard {p} big />{/each}
	</div>
	<p class="on-sky tally">{built} hardware builds that actually exist, plus a pile of designs waiting on parts.</p>
</section>

<section class="section">
	<div class="section-head">
		<h2>things i run</h2>
		<a class="sky-link" href="/events">and where i've been</a>
	</div>
	<div class="row4">
		{#each ran as e}<EventCard {e} />{/each}
	</div>
</section>

<style>
	.about { display: grid; grid-template-columns: minmax(0, 1fr) auto; gap: 48px; align-items: center; margin-bottom: 64px; }
	.about p { max-width: 60ch; margin-bottom: 12px; }
	/* a terminal window, darker than the sky, tilted like a board would be */
	.term {
		display: block; width: 300px; rotate: 3deg; text-align: left; cursor: pointer; color: #d5dce6;
		background: rgba(5, 10, 18, 0.92); border: 1.5px solid rgba(255, 255, 255, 0.18); border-radius: 8px; overflow: hidden;
		font: 500 15px/1.5 var(--mono);
	}
	.term-bar { display: flex; align-items: center; gap: 6px; padding: 8px 10px; background: rgba(255, 255, 255, 0.07); }
	.term-bar i { width: 10px; height: 10px; border-radius: 50%; background: #ff5f57; }
	.term-bar i:nth-child(2) { background: #febc2e; }
	.term-bar i:nth-child(3) { background: #28c840; }
	.term-bar span { margin-left: 8px; font-size: 12px; color: #7d8aa0; }
	.term-body { display: grid; gap: 4px; padding: 16px 18px 12px; }
	.term-note { color: #7d8aa0; }
	.term-line { color: #fff; font-weight: 700; font-size: 17px; white-space: nowrap; }
	.prompt { color: #46e07c; margin-right: 4px; }
	.cursor {
		display: inline-block; width: 0.6em; height: 1.1em; margin-left: 3px; vertical-align: -0.2em;
		background: var(--yellow); animation: cursor 1.1s steps(1) infinite;
	}
	@keyframes cursor { 50% { opacity: 0; } }
	.term-out { color: #7d8aa0; font-size: 13px; }
	.term-hint {
		display: block; padding: 8px 18px 12px; font-size: 12px; color: var(--yellow);
		border-top: 1px dashed rgba(255, 255, 255, 0.14);
	}
	.term:hover { border-color: var(--yellow); }
	.term:hover .term-line { text-decoration: underline; text-decoration-color: var(--yellow); text-underline-offset: 4px; }
	@media (prefers-reduced-motion: reduce) { .cursor { animation: none; } }

	.live-head { display: grid; grid-template-columns: 2.25fr 0.8fr; gap: 22px; align-items: end; margin-bottom: 16px; }
	.contact-heading { font-size: 20px; }
	.live {
		display: grid; gap: 22px; margin-bottom: 64px;
		grid-template-columns: 1fr 1.25fr 0.8fr;
		grid-template-areas: 'a b e' 'c d e';
	}
	.running { grid-area: a; }
	.t-dayoneof.running {
		position: relative; display: block; padding: 22px 26px; overflow: hidden;
		background: var(--bg); color: var(--fg); font-family: var(--text);
		border: 2px solid var(--line); border-radius: var(--radius); box-shadow: 6px 6px 0 var(--line);
	}
	.t-dayoneof.running h3 { color: var(--accent); }
	.t-dayoneof.running .big-line { font: 700 40px/1 var(--head); }
	.dayo-mark { position: absolute; right: 18px; top: 16px; width: 96px; rotate: 5deg; }
	.live a.running:hover .big-line { text-decoration-color: var(--accent); }
	.bench { grid-area: b; }
	.playing { grid-area: c; }
	.github { grid-area: d; }
	.contacts { grid-area: e; display: grid; grid-template-rows: repeat(4, 1fr); gap: 14px; }
	.live a.board, .live a.running { text-decoration: none; }
	.live a.board:hover .big-line, .live a.board:hover .song, .live a.running:hover .big-line { text-decoration: underline; text-decoration-color: var(--gold); text-underline-offset: 5px; }
	h3.silk { margin-bottom: 12px; }
	.big-line { font: 800 36px/1.05 var(--mono); margin-bottom: 10px; }
	.running p:not(.big-line), .bench p:not(.big-line) { max-width: 34ch; font-size: 16px; }

	.playing-body { display: flex; gap: 18px; align-items: center; }
	.playing-body img { width: 120px; aspect-ratio: 4 / 3; object-fit: cover; border: 2px solid rgba(255, 255, 255, 0.8); border-radius: 4px; }
	.song { font: 700 22px/1.1 var(--mono); margin-bottom: 4px; }
	.note { font-size: 14px; opacity: 0.85; }

	.repos { list-style: none; display: grid; gap: 10px; }
	.repos li { display: grid; grid-template-columns: auto 1fr; column-gap: 12px; align-items: baseline; }
	.repos li p { grid-column: 1 / -1; }
	.repos a { font: 700 16px var(--mono); text-decoration-color: var(--gold); text-underline-offset: 3px; }
	.repos p { font-size: 14px; margin: 2px 0 0; line-height: 1.35; }
	.lang { font-size: 13px; color: var(--gold); }

	.contact { display: flex; flex-direction: column; justify-content: center; text-decoration: none; padding: 14px 26px 14px 30px; }
	.contact strong { font: 700 18px var(--mono); }
	.contact span { font-size: 13px; overflow-wrap: anywhere; opacity: 0.85; }
	.contact:hover strong { text-decoration: underline; text-decoration-color: var(--gold); text-underline-offset: 4px; }

	.row3 { display: grid; grid-template-columns: repeat(3, 1fr); gap: 22px; }
	.tally { margin-top: 18px; font-size: 16px; width: fit-content; }
	.row4 { display: grid; grid-template-columns: repeat(4, 1fr); gap: 18px; }

	@media (max-width: 1100px) {
		.live { grid-template-columns: 1fr 1fr; grid-template-areas: 'a b' 'c d' 'e e'; }
		.contact-heading { display: none; }
		.contacts { grid-template-rows: none; grid-template-columns: repeat(4, 1fr); }
		.row4 { grid-template-columns: 1fr 1fr; }
	}
	@media (max-width: 800px) {
		.row3 { grid-template-columns: 1fr; }
		.contacts { grid-template-columns: 1fr 1fr; }
	}
	@media (max-width: 640px) {
		.about { grid-template-columns: 1fr; }
		.term-wrap { order: -1; }
		.term { width: 100%; max-width: 300px; rotate: 0deg; }
		.live { grid-template-columns: 1fr; grid-template-areas: 'a' 'b' 'c' 'd' 'e'; }
		.repos { grid-template-columns: 1fr; }
		.row4 { grid-template-columns: 1fr; }
	}
</style>
