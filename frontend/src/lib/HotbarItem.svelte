<script lang="ts">
	import { hotbar, selected } from './canvas/hotbar';
	import { LayoutGrid, Trash2 } from 'lucide-svelte';
	import { openInventory } from './canvas/inventory';

	let { slot }: { slot: number | 'remove' | 'menu' } = $props();

	function onClick() {
		switch (slot) {
			case 'remove':
				$selected = -1;
				break;
			case 'menu':
				openInventory();
				break;
			default:
				$selected = slot;
				break;
		}
	}

	function onKeydown(event: KeyboardEvent) {
		if (slot === 'remove' && event.key === '§') {
			onClick();
			return;
		}
		if (slot === 'menu' && event.key === 'e') {
			return;
		}
		if (slot === 'menu' || slot === 'remove') return;

		if (event.key === String(slot + 1)) {
			onClick();
		}
	}
</script>

<svelte:window on:keydown={onKeydown} />

<button data-selected={slot === $selected} onclick={onClick} data-index={slot}>

	{#if slot === 'remove'}
		<Trash2 />
	{:else if slot === 'menu'}
		<LayoutGrid />
	{:else if $hotbar[slot]}
		<img src={`/assets/blocks/${$hotbar[slot]}.png`} alt="" data-ontouch-action="place" />	
	{/if}
</button>

<style>
	button {
		width: 3rem;
		height: 3rem;
		border: none;
		background: none;
		border-radius: 0.5rem;
		cursor: pointer;
		padding: 0;
		color: white;
		transition: transform 0.1s;
		position: relative;
	}

	button[data-selected='true'] {
		transform: translateY(-0.25rem);
		border: 2px solid white;
	}

	img {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		border-radius: 0.5rem;
		image-rendering: pixelated;
		z-index: 2;
	}

	button[data-index='menu'] {
		display: block!important;
	}

	button[data-index='0'] {
		display: block!important;
	}

	@media (max-width: 768px) {
		button {
			display: none;
			transform: translateY(0)!important;
			border: none!important;
		}
	}
</style>
