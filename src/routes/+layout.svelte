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

	let copied = $state(false);
	const button = `<a href="https://darsh.app"><img src="https://darsh.app/img/button.png" width="88" height="31" alt="darsh.app"></a>`;
	async function copy() {
		await navigator.clipboard.writeText(button);
		copied = true;
		setTimeout(() => (copied = false), 1600);
	}
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
		<section class="cast"><div class="board black buttons" aria-label="88 by 31 buttons">
			<div class="mine">
				<button class="b88 b-darsh" type="button" onclick={copy}>darsh<b>.app</b></button>
				<p aria-live="polite">{copied ? 'copied!' : 'click my button to copy its code'}</p>
			</div>
			<div class="row">
				<a class="b88 b-hc" href="https://hackclub.com">hack club</a>
				<span class="b88 b-kicad">made in KiCad</span>
				<span class="b88 b-jlc">fabbed in shenzhen</span>
				<span class="b88 b-solder">60/40 solder inside</span>
				<span class="b88 b-frc">FRC 1325</span>
				<span class="b88 b-human">made by a human</span>
			</div>
		</div></section>
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
	.buttons { padding: 22px 34px; }
	.mine {
		display: flex; align-items: center; gap: 16px; padding-bottom: 14px; margin-bottom: 14px;
		border-bottom: 1px dashed rgba(247, 247, 242, 0.3);
	}
	.mine p { font-size: 14px; opacity: 0.8; }
	.row { display: flex; flex-wrap: wrap; gap: 8px; }
	.b88 {
		width: 88px; height: 31px; display: grid; place-items: center; text-align: center;
		font: 700 8px/1.1 var(--mono); text-transform: uppercase; border: 1px solid; color: inherit; text-decoration: none;
	}
	button.b88 { cursor: pointer; }
	.b-darsh { background: #3b8ad6; color: #fff; border-color: #fff; font-size: 11px; text-transform: none; }
	.b-darsh b { color: var(--yellow); }
	.b-hc { background: #ec3750; color: #fff; border-color: #fff; }
	.b-kicad { background: #314cb0; color: #fff; border-color: #9cc4ff; }
	.b-jlc { background: var(--green); color: #fff; border-color: var(--gold); }
	.b-solder { background: #c0c6cc; color: #222; border-color: #666; }
	.b-frc { background: #111; color: #f2c230; border-color: #f2c230; }
	.b-human { background: #fff; color: #111; border-color: #111; }
	.foot { margin-top: 22px; font-size: 14px; color: #fff; text-shadow: 0 1px 3px rgba(12, 45, 90, 0.6); }

	@media (max-width: 640px) {
		.top { flex-direction: column; align-items: flex-start; margin-bottom: 36px; }
		nav { gap: 20px; }
	}
</style>
