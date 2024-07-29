<script>
	import Canvas from '$lib/Canvas.svelte';
	import { connect, connected, disconnect } from '$lib/canvas/socket';
	import Hotbar from '$lib/Hotbar.svelte';
	import Inventory from '$lib/Inventory.svelte';
	import Toolbox from '$lib/staff/Toolbox.svelte';
	import { onDestroy, onMount } from 'svelte';
	import './canvas.css';
	import { cooldown } from '$lib/canvas/cooldown';
	import { checkIsLoggedIn } from '$lib/canvas/auth/auth';

	let isStaff = $state(true);
	let allowDisconnect = $state(true);

	onMount(() => {
		connect();
		checkIsLoggedIn();
		setTimeout(() => {
			allowDisconnect = false;
		}, 2000);
	});

	onDestroy(() => {
		allowDisconnect = true;
		disconnect();
	});
</script>

{#if !$connected && !allowDisconnect}
	<div id="disconnected-overlay">
		<h2>Yhteys katkesi!</h2>
		<p>Jos tässä kestää hetki, uudelleenlataa sivu. </p>
	</div>
{/if}

<Canvas />
<Inventory />
<Hotbar />
{#if isStaff}
	<Toolbox />
{/if}

{#if $cooldown}
	<div id="cooldown-overlay">
		<span>{$cooldown}s</span>
	</div>
{/if}

<style>
	#disconnected-overlay {
		z-index: 1000;
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgb(0, 0, 0);
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		color: white;
		font-size: 2em;
		pointer-events: none;

		animation: fade-in 0.5s;
	}

	@keyframes fade-in {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}

	#cooldown-overlay {
		position: fixed;
		top: 1rem;
		left: 50%;
		transform: translateX(-50%);
		padding: 0.5rem 1rem;
		background: #403d589c;
		backdrop-filter: blur(0.5rem);
		color: white;
		border-radius: 0.5rem;
		font-size: 1.5em;
		z-index: 1000;
	}
</style>
