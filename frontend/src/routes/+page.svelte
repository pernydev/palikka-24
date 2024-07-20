<script>
	import Canvas from '$lib/Canvas.svelte';
	import { connected } from '$lib/canvas/socket';
	import Hotbar from '$lib/Hotbar.svelte';
	import Inventory from '$lib/Inventory.svelte';
	import Toolbox from '$lib/staff/Toolbox.svelte';

	let isStaff = $state(true);
</script>

{#if !$connected}
	<div id="disconnected-overlay">
		<h2>Yhteys katkesi!</h2>
		<p>Yritetään uudelleen...</p>
	</div>
{/if}

<Canvas />
<Inventory />
<Hotbar />
{#if isStaff}
	<Toolbox />
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
</style>
