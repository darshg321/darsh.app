<script lang="ts">
	import ProjectCard from '#lib/components/ProjectCard.svelte';
	import { projects, bench, early, type Project } from '#lib/data/projects.js';

	const filters: { id: string; label: string; test: (p: Project) => boolean }[] = [
		{ id: 'all', label: 'everything', test: () => true },
		{ id: 'hardware', label: 'hardware', test: (p) => p.kind !== 'software' },
		{ id: 'software', label: 'software', test: (p) => p.kind !== 'hardware' },
		{ id: 'built', label: 'built and working', test: (p) => ['built', 'shipped', 'ongoing'].includes(p.status) },
		{ id: 'designed', label: 'designed, not built', test: (p) => p.status === 'designed' || p.status === 'assembled' },
		{ id: 'awards', label: 'won something', test: (p) => !!p.award },
		{ id: 'friends', label: 'with friends', test: (p) => !!p.with?.length }
	];
	let active = $state('all');
	const shown = $derived(projects.filter(filters.find((f) => f.id === active)!.test));
	const count = (f: (typeof filters)[number]) => projects.filter(f.test).length;
</script>

<svelte:head>
	<title>projects · darsh.app</title>
	<meta name="description" content="everything i've built: circuit boards, drones, robots, games, apps and the sites for the events i run." />
</svelte:head>

<h1 class="page-title">things i've made</h1>
<p class="on-sky intro">
	everything, roughly in order of how much of me went into it. the pad on each board says where it's at: a
	<span class="status built">filled pad</span> is built and working, a
	<span class="status designed">dashed one</span> is designed but not built yet, and a
	<span class="status ongoing">green one</span> is what i'm working on now.
</p>

<div class="filters" role="group" aria-label="filter projects">
	{#each filters as f}
		<button type="button" class:on={active === f.id} aria-pressed={active === f.id} onclick={() => (active = f.id)}>
			{f.label} <span>{count(f)}</span>
		</button>
	{/each}
</div>

<section class="grid cast" aria-live="polite">
	{#each shown as p (p.slug)}<ProjectCard {p} />{/each}
</section>

<section class="lists">
	<div class="cast"><div class="board black list">
		<h2 class="silk-h">what's on my bench</h2>
		<ul>
			{#each bench as t}
				<li>
					<strong>{t.title}</strong>
					<span class="when">{t.when}</span>
					<p>{t.note}</p>
				</li>
			{/each}
		</ul>
	</div></div>
	<div class="cast" id="early"><div class="board white list">
		<h2 class="silk-h">where it started, 2022–24</h2>
		<ul>
			{#each early as t}
				<li>
					{#if t.url}<a href={t.url}>{t.title}</a>{:else}<strong>{t.title}</strong>{/if}
					<span class="when">{t.when}</span>
					<p>{t.note}</p>
				</li>
			{/each}
		</ul>
	</div></div>
</section>

<style>
	.intro { max-width: 76ch; margin-bottom: 26px; }
	.intro .status { font-size: 15px; color: #fff; text-shadow: var(--sky-text); vertical-align: 1px; }

	.filters { display: flex; flex-wrap: wrap; gap: 8px; margin-bottom: 26px; }
	.filters button {
		font: 600 14px var(--mono); font-variation-settings: 'wdth' 87.5;
		background: rgba(15, 39, 66, 0.55); color: #fff; border: 2px solid rgba(255, 255, 255, 0.7);
		padding: 8px 12px; border-radius: 4px; cursor: pointer; backdrop-filter: blur(2px);
	}
	.filters button span { opacity: 0.7; margin-left: 4px; }
	.filters button:hover { border-color: #fff; }
	.filters button.on { background: var(--yellow); color: var(--black); border-color: var(--yellow); }

	.grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 22px; margin-bottom: 64px; }
	/* the bright masks wash out against the sky, so this page lifts the boards harder:
	   a deeper hard shadow, a soft one under it, and a dark board edge */
	.grid.cast { filter: drop-shadow(-8px 10px 0 rgba(8, 28, 58, 0.45)) drop-shadow(0 14px 18px rgba(8, 28, 58, 0.35)); }
	.grid :global(.card) { box-shadow: inset 0 0 0 2px rgba(0, 0, 0, 0.28), inset 0 -4px 0 rgba(0, 0, 0, 0.12); }
	.intro { background: rgba(10, 32, 64, 0.35); padding: 12px 16px; border-radius: 4px; }

	.lists { display: grid; grid-template-columns: 1fr 1fr; gap: 22px; align-items: start; }
	.list { padding: 26px 30px 24px; }
	.silk-h { font: 700 18px var(--mono); color: inherit; text-shadow: none; margin-bottom: 14px; }
	.list ul { list-style: none; display: grid; gap: 12px; }
	.list a, .list strong { font: 700 15px var(--mono); text-decoration-color: var(--gold); text-underline-offset: 3px; }
	.when { font: 12px var(--mono); opacity: 0.65; margin-left: 8px; }
	.list p { font-size: 14px; line-height: 1.4; opacity: 0.9; }

	@media (max-width: 1100px) { .grid { grid-template-columns: repeat(3, 1fr); } }
	@media (max-width: 820px) { .grid { grid-template-columns: repeat(2, 1fr); } .lists { grid-template-columns: 1fr; } }
	@media (max-width: 520px) { .grid { grid-template-columns: 1fr; } }
</style>
