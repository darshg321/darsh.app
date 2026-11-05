<script lang="ts">
	import '../app.css';
	import '#lib/styles/themes.css';
	import { page } from '$app/state';
	import type { LayoutProps } from './$types';

	let { children }: LayoutProps = $props();

	const nav = [
		{ href: '/', label: 'home' },
		{ href: '/projects', label: 'projects' },
		{ href: '/events', label: 'events' },
		{ href: '/about', label: 'about' }
	];
	const current = (href: string) =>
		href === '/' ? page.url.pathname === '/' : page.url.pathname.startsWith(href);

</script>

<div class="page">
	<header class="top">
		<a class="logo" href="/"><img src="/img/cut/badge.webp" alt="" width="46" height="46" />darsh.app</a>
		<nav aria-label="Main">
			{#each nav as n}
				<a href={n.href} aria-current={current(n.href) ? 'page' : undefined}>{n.label}</a>
			{/each}
		</nav>
	</header>

	<main>
		{@render children()}
	</main>

	<footer>
		<p class="foot">
			by darsh, with (productive) procrastination &lt;3
		</p>
	</footer>
</div>

<style>
	.top { display: flex; align-items: center; justify-content: space-between; gap: 16px; margin-bottom: 56px; }
	.logo {
		display: flex; align-items: center; gap: 12px; text-decoration: none;
		font: 700 22px var(--mono); color: #fff; text-shadow: var(--sky-text);
	}
	.logo img { width: 46px; height: auto; filter: drop-shadow(0 2px 3px rgba(12, 45, 90, 0.4)); }
	nav { display: flex; gap: 28px; flex-wrap: wrap; }
	nav a {
		font: 600 16px var(--mono); text-decoration: none; color: #fff; text-shadow: var(--sky-text);
		padding-bottom: 4px; border-bottom: 3px solid transparent;
	}
	nav a:hover { border-bottom-color: rgba(255, 255, 255, 0.6); }
	nav a[aria-current] { border-bottom-color: var(--yellow); }

	footer { margin-top: 72px; }
	.foot { font-size: 14px; color: #fff; text-shadow: 0 1px 3px rgba(12, 45, 90, 0.6); }

	@media (max-width: 640px) {
		.top { flex-direction: column; align-items: flex-start; margin-bottom: 36px; }
		nav { gap: 20px; }
	}
</style>
