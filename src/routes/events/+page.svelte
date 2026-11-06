<script lang="ts">
	import EventCard from '#lib/components/EventCard.svelte';
	import { events, otherEvents } from '#lib/data/events.js';
	import { themeFonts } from '#lib/fonts.js';

	const ran = events.filter((e) => e.kind === 'ran');
	const went = events.filter((e) => e.kind === 'went');
</script>

<svelte:head>
	<title>events · darsh.app</title>
	<meta name="description" content="the events i run (FraserHacks, Dayoneof, Off-Script, CS Club) and the ones i've been to (Undercity, Fallout, Outpost, Polaris, Supercon)." />
	<link rel="stylesheet" href={themeFonts} />
</svelte:head>

<h1 class="page-title">events</h1>
<p class="on-sky intro">
	the ones i run, and the ones i've flown out to. each one is dressed in its own colours, and links to its own site.
</p>

<section class="section">
	<div class="section-head"><h2>things i run</h2></div>
	<div class="ran">{#each ran as e}<EventCard {e} wide />{/each}</div>
</section>

<section class="section">
	<div class="section-head"><h2>places i've been</h2><p>newest first. hackathons, mostly hardware.</p></div>
	<div class="went">{#each went as e}<EventCard {e} />{/each}</div>
</section>

<section class="cast">
	<div class="board white other">
		<h2 class="silk-h">and also</h2>
		<ul>
			{#each otherEvents as o}<li><strong>{o.name}</strong> <span>{o.note}</span> <em>{o.when}</em></li>{/each}
		</ul>
	</div>
</section>

<style>
	.intro { max-width: 64ch; margin-bottom: 40px; }
	.ran { display: grid; gap: 24px; }
	.went { display: grid; grid-template-columns: repeat(2, 1fr); gap: 22px; }
	.ran :global(.ev), .went :global(.ev) { box-shadow: -8px 10px 0 rgba(8, 28, 58, 0.35); }
	.other { padding: 26px 30px; }
	.silk-h { font: 700 18px var(--mono); color: inherit; text-shadow: none; margin-bottom: 12px; }
	.other ul { list-style: none; display: grid; grid-template-columns: 1fr 1fr; gap: 8px 30px; }
	.other strong { font: 700 15px var(--mono); }
	.other span { font-size: 15px; }
	.other em { font: normal 12px var(--mono); opacity: 0.6; margin-left: 6px; }
	
	@media (max-width: 760px) { .went, .other ul { grid-template-columns: 1fr; } }
</style>
