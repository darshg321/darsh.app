<script lang="ts">
	import ProjectCard from '#lib/components/ProjectCard.svelte';
	import Gallery from '#lib/components/Gallery.svelte';
	import Media from '#lib/components/Media.svelte';
	import { themeFonts } from '#lib/fonts.js';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const e = $derived(data.e);
</script>

<svelte:head>
	<title>{e.name} · darsh.app</title>
	<meta name="description" content="{e.name}: {e.summary}" />
	<link rel="stylesheet" href={themeFonts} />
</svelte:head>

<div class="themed t-{e.theme}">
	<div class="themed-bg"></div>

	<a class="back" href="/events">← all events</a>

	<header class="ehead">
		<div>
			{#if e.logo}<img class="logo" src={e.logo} alt="{e.name} logo" />{/if}
			<h1>{e.name}</h1>
			<p class="when">{e.when} · {e.place}</p>
			{#if e.role}<p class="role">{e.role}</p>{/if}
			{#if e.award}<p class="award"><b>★</b> {e.award}</p>{/if}
		</div>
		{#if e.cover}<img class="cover" class:site={e.cover.startsWith('/img/s/')} src={e.cover} alt="" />{/if}
	</header>

	{#if e.stats}
		<div class="stats">
			{#each e.stats as s}<div class="panel stat"><strong>{s.n}</strong><span>{s.label}</span></div>{/each}
		</div>
	{/if}

	<div class="cols">
		<section class="panel prose">
			{@html data.body}
			{#if e.links?.length}
				<p class="links">{#each e.links as l}<a href={l.url}>{l.label} ↗</a>{/each}</p>
			{/if}
		</section>
		{#if e.video}
			<div class="panel vid"><Media video={e.video} /></div>
		{:else if e.decor?.length}
			<div class="decor" aria-hidden="true">{#each e.decor as src}<img {src} alt="" />{/each}</div>
		{/if}
	</div>

	{#if e.gallery?.length}
		<section class="block">
			<h2>pictures</h2>
			<Gallery images={e.gallery} />
		</section>
	{/if}

	{#if data.made.length}
		<section class="block">
			<h2>{e.kind === 'ran' ? 'what i made for it' : 'what i made there'}</h2>
			<div class="made cast">{#each data.made as p}<ProjectCard {p} />{/each}</div>
		</section>
	{/if}
</div>

<style>
	.themed { margin: -20px -32px 0; padding: 20px 32px 10px; }
	.back { display: inline-block; margin-bottom: 22px; font: 600 15px var(--mono); color: inherit; text-decoration-color: var(--accent); text-underline-offset: 4px; }
	.ehead { display: grid; grid-template-columns: minmax(0, 1.2fr) minmax(0, 1fr); gap: 36px; align-items: center; margin-bottom: 30px; }
	.logo { max-height: 120px; width: auto; margin-bottom: 14px; }
	h1 { font: 700 64px/1.02 var(--head); color: inherit; letter-spacing: -1px; }
	.when { font-size: 18px; margin-top: 10px; opacity: 0.9; }
	.role {
		display: inline-block; margin-top: 14px; font: 700 14px var(--mono); padding: 7px 11px; border-radius: 3px;
		background: var(--accent); color: var(--ink);
	}
	/* same sticker as the cards, just bigger */
	.award {
		display: inline-block; margin: 18px 0 0 14px; font: 800 20px/1.2 var(--mono); padding: 11px 16px;
		background: #ffd23f; color: #16181b; border: 2px solid #16181b; box-shadow: -5px 5px 0 #16181b; rotate: -2deg; border-radius: 3px;
	}
	.award b { color: #c8372d; }
	.cover { border-radius: var(--radius); max-height: 340px; width: 100%; object-fit: cover; }
	/* site screenshots are shown whole */
	.cover.site { max-height: none; height: auto; object-fit: contain; }
	.t-outpost .cover, .t-dayoneof .cover { border: 3px solid var(--line); box-shadow: 6px 6px 0 var(--line); }
	.t-fallout .cover { border: 6px solid #8f6a55; }
	.t-fallout .logo { max-height: 150px; }
	.t-hackthenorth .cover { border: 3px solid #fff; }
	/* Hack the North's stats are its yellow game cards */
	.t-hackthenorth .stats { display: flex; flex-wrap: wrap; }
	.t-hackthenorth .stat {
		width: clamp(140px, 42vw, 170px); aspect-ratio: 413 / 532; justify-content: center; align-items: center; text-align: center; padding: 30px 12px 0 !important;
		background: url('/img/e/htn/gamecard.webp') center / contain no-repeat !important; border: 0 !important; box-shadow: none !important; color: #3a2a0a;
	}
	.t-hackthenorth .stat strong { font-size: 34px; }
	.t-hackthenorth .stat span { font: 700 17px var(--text); opacity: 1; margin: 2px 0 0; }
	.t-hackthenorth .decor img { width: 36%; }
	.t-hackthenorth .decor img:first-child { width: 50%; }
	.t-polaris .logo { filter: invert(1) hue-rotate(180deg); }
	.t-undercity .cover { border: 2px solid var(--line); }
	.t-offscript .cover { border: 1px solid var(--line); object-position: center 30%; }
	.t-offscript .logo { max-height: 34px; }
	.t-offscript .role { border-radius: 999px; padding: 8px 16px; font-family: 'Inter Tight', sans-serif; font-weight: 800; text-transform: uppercase; }
	.t-offscript .stat strong { background: var(--wtw); -webkit-background-clip: text; background-clip: text; color: transparent; }

	.stats { display: grid; grid-template-columns: repeat(auto-fit, minmax(150px, 1fr)); gap: 16px; margin-bottom: 26px; }
	.stat { padding: 18px 22px !important; display: flex; flex-direction: column; }
	.stat strong { font: 700 38px/1 var(--head); }
	.t-fraserhacks .stat strong { text-shadow: var(--glow); }
	.stat span { font-size: 14px; margin-top: 6px; opacity: 0.85; }

	.cols { display: grid; grid-template-columns: minmax(0, 1.6fr) minmax(0, 1fr); gap: 24px; align-items: start; }
	.cols:not(:has(.vid, .decor)) { grid-template-columns: minmax(0, 820px); }
	/* borrowed stickers, scattered rather than lined up */
	.decor { position: relative; display: flex; flex-wrap: wrap; justify-content: center; align-items: center; gap: 4px 10px; }
	.decor img { width: 46%; height: auto; filter: drop-shadow(0 6px 0 rgba(0, 0, 0, 0.12)); }
	.decor img:nth-child(odd) { rotate: -5deg; }
	.decor img:nth-child(even) { rotate: 4deg; margin-top: 30px; }
	.decor img:first-child:nth-last-child(odd) { width: 62%; }
	.t-fallout .decor img:first-child { width: 88%; rotate: -2deg; border: 4px solid #61453a; border-radius: 14px; filter: none; }
	.t-fallout .decor img:nth-child(2) { width: 56%; margin-top: 4px; filter: none; }
	.prose { font-size: 18px; }
	.links { display: flex; flex-wrap: wrap; gap: 8px 18px; }
	.links a { font: 700 15px var(--mono); }
	.vid { padding: 16px !important; }

	.block { margin-top: 48px; }
	.block h2 { font: 700 28px var(--head); color: inherit; text-shadow: none; margin-bottom: 16px; }
	.t-fraserhacks .block h2, .t-undercity .block h2 { text-shadow: var(--glow); }
	.t-fallout .block h2, .t-fallout .ehead { color: #fff; text-shadow: 0 2px 0 rgba(0, 60, 110, 0.45); }
	/* their hand-lettered headings: white, outlined in the frame's brown, like the wordmark */
	.t-fallout .block h2 { font-size: 46px; font-weight: 400; -webkit-text-stroke: 7px #61453a; paint-order: stroke fill; text-shadow: none; letter-spacing: 1px; }
	.made { display: grid; grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); gap: 20px; }
	.made :global(.card) { color: var(--ink); font-family: var(--body); }

	@media (max-width: 900px) {
		.ehead, .cols { grid-template-columns: 1fr; }
		h1 { font-size: 44px; }
	}
	@media (max-width: 640px) { .themed { margin: -10px -18px 0; padding: 10px 18px; } }
</style>
