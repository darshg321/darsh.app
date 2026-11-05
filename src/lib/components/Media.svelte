<script lang="ts">
	let {
		video,
		youtube
	}: { video?: { src: string; caption: string }; youtube?: { id: string; caption: string } } = $props();
</script>

{#if video}
	<figure class="media">
		<!-- svelte-ignore a11y_media_has_caption -->
		<video src={video.src} controls muted playsinline preload="metadata"></video>
		<figcaption class="silk">{video.caption}</figcaption>
	</figure>
{/if}
{#if youtube}
	<figure class="media">
		<iframe
			src="https://www.youtube-nocookie.com/embed/{youtube.id}"
			title={youtube.caption}
			loading="lazy"
			allow="accelerometer; encrypted-media; gyroscope; picture-in-picture"
			allowfullscreen
		></iframe>
		<figcaption class="silk">{youtube.caption}</figcaption>
	</figure>
{/if}

<style>
	.media { margin-bottom: 14px; }
	video { width: 100%; max-height: 70vh; border-radius: 4px; background: #000; }
	iframe { width: 100%; aspect-ratio: 16 / 9; border: 0; border-radius: 4px; background: #000; }
	figcaption { margin-top: 8px; opacity: 0.85; }
</style>
