<script lang="ts">
	import type { Event } from '#lib/data/events.js';

	let { e, wide = false }: { e: Event; wide?: boolean } = $props();
	// screenshots of a site are shown whole in a little browser window, never cropped
	const site = $derived(e.cover?.startsWith('/img/s/'));
</script>

<a class="ev t-{e.theme}" class:wide href="/events/{e.slug}">
	<div class="ev-body">
		{#if e.logo}<img class="ev-logo" src={e.logo} alt="" loading="lazy" />{/if}
		<span class="ev-name">{e.name}</span>
		<span class="ev-when">{e.when} · {e.place}</span>
		<span class="ev-sum">{e.summary}</span>
		{#if wide && e.stats}
			<span class="ev-stats">
				{#each e.stats.slice(0, 3) as s}<span><b>{s.n}</b>{s.label}</span>{/each}
			</span>
		{/if}
		{#if e.role}<span class="ev-role">{e.role}</span>{/if}
	</div>
	{#if wide && e.cover}
		<div class="ev-cover" class:site>
			{#if site}<span class="ev-browser" aria-hidden="true"><i></i><i></i><i></i></span>{/if}
			<img src={e.cover} alt="" loading="lazy" />
		</div>
	{/if}
	{#if e.award}<span class="ev-award"><b>★</b> {e.award}</span>{/if}
</a>
