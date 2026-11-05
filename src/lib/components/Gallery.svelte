<script lang="ts">
	import type { Img } from '#lib/data/projects.js';

	let { images }: { images: Img[] } = $props();
	let open = $state<number | null>(null);
	let dialog = $state<HTMLDialogElement>();

	function show(i: number) {
		open = i;
		dialog?.showModal();
	}
	function step(d: number) {
		if (open === null) return;
		open = (open + d + images.length) % images.length;
	}
	function key(e: KeyboardEvent) {
		if (e.key === 'ArrowRight') step(1);
		if (e.key === 'ArrowLeft') step(-1);
	}
</script>

<div class="grid">
	{#each images as img, i}
		<button type="button" class="thumb" onclick={() => show(i)} aria-label="view: {img.alt}">
			<img src={img.src} alt={img.alt} loading="lazy" decoding="async" />
		</button>
	{/each}
</div>

<dialog bind:this={dialog} onclose={() => (open = null)} onkeydown={key}>
	{#if open !== null}
		<figure>
			<img src={images[open].src} alt={images[open].alt} />
			<figcaption>{images[open].alt} <span>{open + 1} / {images.length}</span></figcaption>
		</figure>
		<div class="ctl">
			{#if images.length > 1}
				<button type="button" onclick={() => step(-1)} aria-label="previous">←</button>
				<button type="button" onclick={() => step(1)} aria-label="next">→</button>
			{/if}
			<button type="button" onclick={() => dialog?.close()}>close</button>
		</div>
	{/if}
</dialog>

<style>
	.grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(170px, 1fr)); gap: 12px; }
	.thumb {
		aspect-ratio: 4 / 3; border: 0; padding: 0; cursor: zoom-in; border-radius: 4px; overflow: hidden;
		background: rgba(255, 255, 255, 0.92);
	}
	.thumb img { width: 100%; height: 100%; object-fit: contain; transition: scale 0.2s; }
	.thumb:hover img { scale: 1.04; }
	dialog {
		margin: auto; border: 0; background: transparent; max-width: min(1100px, 94vw); max-height: 94vh; color: #fff;
	}
	dialog::backdrop { background: rgba(8, 20, 40, 0.88); }
	figure img { max-height: 80vh; width: auto; margin: 0 auto; border-radius: 4px; background: #fff; }
	figcaption { font: 14px var(--body); margin-top: 10px; text-align: center; }
	figcaption span { opacity: 0.6; margin-left: 8px; font-family: var(--mono); font-size: 12px; }
	.ctl { display: flex; justify-content: center; gap: 10px; margin-top: 12px; }
	.ctl button {
		font: 600 14px var(--mono); background: var(--yellow); color: var(--black); border: 0;
		padding: 8px 14px; border-radius: 3px; cursor: pointer;
	}
</style>
