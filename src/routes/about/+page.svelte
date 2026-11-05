<script lang="ts">
	import { roles, frc, awards, offline, contact } from '#lib/data/about.js';
	import { md } from '#lib/md.js';
	import { bySlug } from '#lib/data/projects.js';

	const mushu = bySlug('mushu')!;
</script>

<svelte:head>
	<title>about · darsh.app</title>
	<meta name="description" content="about darsh: what i do, FRC 1325, awards and the rest." />
</svelte:head>

<section class="intro">
	<div class="cast photo-wrap">
		<figure class="board white photo">
			<img src="/img/me/camera.webp" alt="me holding a Nikon camera up to my face, in front of a graffiti wall" />
			<figcaption class="silk">me, behind the camera</figcaption>
		</figure>
	</div>
	<div>
		<h1 class="page-title">about me</h1>
		<p class="on-sky">
			i'm Darsh Gupta, in grade 12 at John Fraser Secondary School in Mississauga, graduating in 2027. i build
			hardware: circuit boards in KiCad, mechanisms in Onshape, firmware in C and C++, and lately an FPGA board and a
			chip.
		</p>
		<p class="on-sky">
			i started with Python games in grade 8, moved to websites and apps, and got hooked on hardware the summer before
			grade 11, when a Hack Club Hackpad turned into a drone, a keyboard, and a trip to Undercity. most of what i care
			about now is getting other people building too: FraserHacks, the CS Club, Dayoneof, and Off-Script.
		</p>
		<p class="on-sky">
			say hi at <a href="mailto:{contact.email}">{contact.email}</a>
			{#each contact.links as l}, <a href={l.url}>{l.label}</a>{/each}.
		</p>
	</div>
</section>

<section class="section">
	<div class="section-head"><h2>what i do</h2></div>
	<div class="roles cast">
		{#each roles as r}
			<article class="board {r.mask} role">
				<p class="silk">{r.when}</p>
				<h3>{r.org}</h3>
				<p class="title">{r.title}</p>
				<ul>{#each r.points as pt}<li>{pt}</li>{/each}</ul>
				{#if r.href}<a class="more" href={r.href}>more →</a>{/if}
			</article>
		{/each}
	</div>
</section>

<section class="section" id="frc">
	<div class="section-head"><h2>FRC team 1325</h2></div>
	<div class="frc cast">
		<div class="board black frc-board">
			<div class="prose">{@html md(frc.body)}</div>
			<ul class="results">{#each frc.results as r}<li>★ {r}</li>{/each}</ul>
		</div>
		<div class="board black mushu">
			<img src={mushu.cover} alt="Mushu, 1325's 2025 robot" />
			<span class="silk">Mushu, 2025</span>
		</div>
	</div>
</section>

<section class="section">
	<div class="split">
		<div class="cast"><div class="board yellow list">
			<h2 class="silk-h">awards</h2>
			<ul>
				{#each awards as a}
					<li>
						{#if a.href}<a href={a.href}><strong>{a.what}</strong></a>{:else}<strong>{a.what}</strong>{/if}
						<span>{a.who}</span><em>{a.when}</em>
					</li>
				{/each}
			</ul>
		</div></div>
		<div class="stack">
			<div class="cast"><div class="board green list">
				<h2 class="silk-h">away from the desk</h2>
				<ul>{#each offline as o}<li><strong>{o.what}</strong> <span>{o.note}</span></li>{/each}</ul>
			</div></div>
		</div>
	</div>
</section>

<style>
	.intro { display: grid; grid-template-columns: auto minmax(0, 1fr); gap: 48px; align-items: center; margin-bottom: 64px; }
	.intro p { max-width: 62ch; margin-bottom: 12px; }
	.photo { width: 280px; padding: 30px 26px 14px; rotate: -3deg; }
	.photo img { width: 100%; border-radius: 3px; }
	.photo figcaption { display: block; margin-top: 10px; text-align: center; }

	.roles { display: grid; grid-template-columns: repeat(auto-fill, minmax(330px, 1fr)); gap: 20px; }
	.role { padding: 28px 28px 22px; display: flex; flex-direction: column; }
	.role h3 { font: 700 21px/1.15 var(--mono); margin-top: 4px; }
	.role .title { font-weight: 600; margin-bottom: 10px; opacity: 0.9; }
	.role ul { padding-left: 1.1em; font-size: 15px; display: grid; gap: 4px; margin-bottom: 12px; }
	.more { margin-top: auto; font: 700 14px var(--mono); text-decoration-color: var(--gold); }

	.frc { display: grid; grid-template-columns: minmax(0, 1.6fr) minmax(0, 1fr); gap: 22px; }
	.frc-board { padding: 32px 36px; }
	.results { list-style: none; display: grid; gap: 6px; font: 600 14px var(--mono); color: var(--gold); }
	.mushu { display: grid; place-items: center; padding: 30px; text-decoration: none; }
	.mushu img { filter: drop-shadow(-5px 7px 0 rgba(0, 0, 0, 0.4)); }

	.split { display: grid; grid-template-columns: 1.2fr 1fr; gap: 22px; align-items: start; }
	.stack { display: grid; gap: 22px; }
	.list { padding: 26px 30px; }
	.silk-h { font: 700 18px var(--mono); color: inherit; text-shadow: none; margin-bottom: 12px; }
	.list ul { list-style: none; display: grid; gap: 9px; font-size: 15px; }
	.list li span { opacity: 0.85; }
	.list li em { font: normal 12px var(--mono); opacity: 0.6; margin-left: 8px; }
	.list li strong { font-weight: 700; }
	.list a { text-decoration-color: var(--black); }

	@media (max-width: 900px) {
		.intro, .frc, .split { grid-template-columns: 1fr; }
		.photo { width: 220px; }
	}
</style>
