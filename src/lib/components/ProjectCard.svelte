<script lang="ts">
	import { statusLabel, repo, type Project } from '#lib/data/projects.js';
	import { events } from '#lib/data/events.js';

	let { p, big = false }: { p: Project; big?: boolean } = $props();
	const url = $derived(repo(p));
	// where it's been: events it lists, and events that list it
	const been = $derived(events.filter((e) => p.events?.includes(e.slug) || e.projects?.includes(p.slug)));
</script>

<!-- the title link stretches over the whole card; the event chips sit above it as their own links -->
<div class="board {p.mask} card" class:big class:linked={url}>
	<div class="art" class:photo={p.photo}>
		{#if p.cover}
			<img src={p.cover} alt="" loading="lazy" decoding="async" />
		{:else}
			<span class="silk-art" aria-hidden="true">{p.title}</span>
		{/if}
	</div>
	<div class="text">
		<h3>{#if url}<a class="go" href={url}>{p.title}</a>{:else}{p.title}{/if}</h3>
		<p>{p.tagline}</p>
	</div>
	<div class="meta">
		<span class="status {p.status}">{statusLabel[p.status]}</span>
		<span class="silk">{p.year}</span>
	</div>
	{#if p.award}<span class="award silk">★ {p.award}</span>{/if}
	{#if been.length}
		<div class="been">{#each been as e}<a href="/events/{e.slug}">@ {e.name}</a>{/each}</div>
	{/if}
</div>

<style>
	.card { position: relative; display: flex; flex-direction: column; padding: 28px 22px 20px; height: 100%; }
	.go { color: inherit; text-decoration: none; }
	.go::after { content: ''; position: absolute; inset: 0; z-index: 1; }
	.go:focus-visible { outline: none; }
	.go:focus-visible::after { outline: 3px solid var(--gold); outline-offset: 2px; }
	/* nothing gets cropped: cutouts and photos both scale to fit, centred in the same box */
	.art { height: 160px; margin-bottom: 16px; flex-shrink: 0; display: flex; align-items: center; justify-content: center; }
	.art img { max-height: 100%; max-width: 100%; width: auto; height: auto; filter: drop-shadow(-4px 6px 0 rgba(0, 0, 0, 0.3)); }
	/* photos sit on the board like a print: thin white border, same hard shadow */
	.art.photo img { border: 3px solid #f7f7f2; border-radius: 2px; filter: none; box-shadow: -4px 5px 0 rgba(0, 0, 0, 0.3); }
	.big .art { height: 220px; }
	.silk-art {
		width: 100%; height: 100%; display: grid; place-items: center; text-align: center;
		border: 1.5px dashed color-mix(in srgb, currentColor 45%, transparent); border-radius: 4px;
		font: 700 22px var(--mono); font-variation-settings: 'wdth' 75; opacity: 0.75; padding: 10px;
	}
	h3 { font: 700 20px/1.15 var(--mono); font-variation-settings: 'wdth' 87.5; }
	.big h3 { font-size: 24px; }
	.text p { font-size: 15px; line-height: 1.4; margin: 5px 0 12px; }
	.meta { margin-top: auto; display: flex; justify-content: space-between; align-items: center; gap: 8px; }
	.award { display: block; margin-top: 8px; color: var(--gold); line-height: 1.3; }
	.yellow .award, .white .award { color: #7a5300; }
	.linked:hover h3 { text-decoration: underline; text-decoration-color: var(--gold); text-underline-offset: 5px; }
	/* the events it went to, as little silkscreen tags */
	.been { position: relative; z-index: 2; display: flex; flex-wrap: wrap; gap: 6px; margin-top: 10px; }
	.been a {
		font: 600 12px/1 var(--mono); padding: 5px 7px; border-radius: 3px; text-decoration: none; color: inherit;
		border: 1.5px solid color-mix(in srgb, currentColor 55%, transparent);
	}
	.been a:hover { border-color: var(--gold); background: color-mix(in srgb, currentColor 12%, transparent); }
</style>
