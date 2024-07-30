<script lang="ts">
	import { loggedIn } from './canvas/auth/auth';
	import { open } from './canvas/open';
	import { onDestroy, onMount } from 'svelte';
	import panzoom, { type PanZoom } from 'panzoom';
	import { CANVAS_HEIGHT, CANVAS_WIDTH, PIXEL_SIZE } from '../constants';
	import { place } from './canvas/place';
	import ImageGrid from './ImageGrid.svelte';
	import { getCanvas, grid, init, initialLoad } from './canvas/grid';
	import { hotbar, selected } from './canvas/hotbar';
	import { closeInventory, inventoryOpen, setOpenInventory } from './canvas/inventory';
	import { tool } from './canvas/tool';
	import { deleteArea } from './canvas/staff/area';
	import { cooldown } from './canvas/cooldown';
	import { get } from 'svelte/store';
	import { end } from './canvas/secret/end';

	let canvas: HTMLDivElement;

	let selection = $state({
		x: 0,
		y: 0
	});

	let mouse = {
		x: 0,
		y: 0
	};

	let initial = {
		x: 0,
		y: 0
	};

	let area = $state({
		step: 0,
		x: 0,
		y: 0,
		width: 1,
		height: 1
	});

	let isMouseDown = false;

	let panzoomInstance: PanZoom;

	onMount(() => {
		panzoomInstance = panzoom(canvas, {
			minZoom: 0.5,
			maxZoom: 8,
			bounds: true,
			boundsPadding: 0.1,
			autocenter: true,
			zoomDoubleClickSpeed: 1, // Disable double click zoom
			onTouch
		});

		init();
	});

	function onTouch(event: TouchEvent) {
		console.log('touch');
		// get the item touched on. If it has data-ontouch-action switch statement for action
		const target = event.target as HTMLElement;
		const action = target.getAttribute('data-ontouch-action');
		if (!action) {
			touchPlace(event);
			return;
		}

		switch (action) {
			case 'inventory':
				const slot = target.getAttribute('data-inventory-item');
				if (!slot) return;
				$selected = 0;
				$hotbar[0] = parseInt(slot);
				console.log('inventory', slot);
				closeInventory();
				break;
			case 'place':
				place(selection.x, selection.y);
				break;
		}
	}

	$effect(() => {
		if (!panzoom) return;
		if ($inventoryOpen) panzoomInstance.pause();
		else panzoomInstance.resume();
	});

	function setSelection() {
		const rect = canvas.getBoundingClientRect();
		const x = (mouse.x - rect.left) / panzoomInstance.getTransform().scale;
		const y = (mouse.y - rect.top) / panzoomInstance.getTransform().scale;

		selection.x = Math.min(Math.max(0, Math.floor(x / PIXEL_SIZE)), CANVAS_WIDTH);
		selection.y = Math.min(Math.max(0, Math.floor(y / PIXEL_SIZE)), CANVAS_HEIGHT);
	}

	function mouseDown(event: MouseEvent) {
		initial = {
			x: event.clientX,
			y: event.clientY
		};

		if (event.which === 2) {
			if (get(end)) return;
			if ($selected === -1) $selected = 0;
			const block = $grid[`${selection.x},${selection.y}`];
			if (block) {
				$hotbar[$selected] = block;
			}
			return;
		}

		isMouseDown = true;
	}

	function mouseMove(event: MouseEvent) {
		mouse = {
			x: event.clientX,
			y: event.clientY
		};

		setSelection();

		if (!mouseDown) return;
		let diffX = Math.abs(event.clientX - initial.x);
		let diffY = Math.abs(event.clientY - initial.y);
		if (diffX > 5 || diffY > 5) {
			isMouseDown = false;
		}

		switch ($tool) {
			case 'area':
				if (area.step === 0) {
					area = {
						x: selection.x,
						y: selection.y,
						width: 1,
						height: 1,
						step: 0
					};
				} else {
					area.width = selection.x - area.x + 1;
					area.height = selection.y - area.y + 1;
					if (area.width > 20) {
						area.width = 20;
					}
					if (area.height > 20) {
						area.height = 20;
					}
				}
		}
	}

	function canvasTouchStart(event: TouchEvent) {
		if (event.touches.length > 1) return;
		initial = {
			x: event.touches[0].clientX,
			y: event.touches[0].clientY
		};
		isMouseDown = true;
	}

	function touchMove(event: TouchEvent) {
		if (!mouseDown) return;
		let diffX = Math.abs(event.touches[0].clientX - initial.x);
		let diffY = Math.abs(event.touches[0].clientY - initial.y);
		if (diffX > 5 || diffY > 5) {
			isMouseDown = false;
		}
	}

	function touchPlace(event: TouchEvent) {
		if (!mouseDown) return;

		if ($tool === 'paint') {
			if (!$open) return;
			if (!$loggedIn) return;
			if ($cooldown > 0) return;

			const rect = canvas.getBoundingClientRect();
			const x = (event.touches[0].clientX - rect.left) / panzoomInstance.getTransform().scale;
			const y = (event.touches[0].clientY - rect.top) / panzoomInstance.getTransform().scale;

			selection.x = Math.min(Math.max(0, Math.floor(x / PIXEL_SIZE)), CANVAS_WIDTH);
			selection.y = Math.min(Math.max(0, Math.floor(y / PIXEL_SIZE)), CANVAS_HEIGHT);
		}
		isMouseDown = false;
	}

	function mouseUp(event: Event) {
		if (window.innerWidth < 768) return;

		if (isMouseDown) {
			switch ($tool) {
				case 'paint':
					if (!$open) return;
					if (!$loggedIn) return;
					if ($cooldown > 0) return;
					place(selection.x, selection.y);
					break;
				case 'area':
					console.log(area);
					if (area.step === 0) {
						area.step = 1;
					} else {
						area.step = 0;
						deleteArea(area.x, area.y, area.width, area.height);
					}
			}
		}
		isMouseDown = false;
	}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div
	bind:this={canvas}
	id="canvas"
	onmousemove={mouseMove}
	tabindex="-1"
	onmousedown={mouseDown}
	onmouseup={mouseUp}
	ontouchstart={canvasTouchStart}
	ontouchmove={touchMove}
	data-loaded={$initialLoad}
>
	{#if $tool === 'paint' && $open && $loggedIn}
		<div
			id="selection"
			style="--x: {selection.x}; --y: {selection.y}; --src: url('/assets/blocks/{$hotbar[
				$selected
			]}.png');"
			tabindex="-1"
			data-cooldown={$cooldown > 0}
			data-remover={$selected === -1}
		></div>
	{:else if $tool === 'area'}
		<div
			id="area"
			style="--x: {area.x}; --y: {area.y}; --width: {area.width}; --height: {area.height};"
			tabindex="-1"
		></div>
	{/if}
	<ImageGrid />
</div>

<style>
	* {
		user-select: none;
	}
	#canvas {
		width: calc(var(--canvas-width) * var(--pixel-size));
		height: calc(var(--canvas-height) * var(--pixel-size));
		background-color: #fff;

		transform: scale(var(--scale));
		transition: transform 0.08s;

		position: relative;

		opacity: 0;
		filter: blur(5px);
	}

	#canvas[data-loaded='true'] {
		animation: fade-in 1s forwards;
	}

	@keyframes fade-in {
		from {
			opacity: 0;
			filter: blur(5px);
		}
		to {
			opacity: 1;
			filter: blur(0);
		}
	}

	#selection {
		position: absolute;
		top: calc(var(--y) * var(--pixel-size));
		left: calc(var(--x) * var(--pixel-size));
		width: var(--pixel-size);
		height: var(--pixel-size);
		z-index: -1;
		opacity: 0.5;
		z-index: 0;
		image-rendering: pixelated;
		background: var(--src);
	}

	#selection[data-remover='true'] {
		background: white;
	}

	#selection[data-cooldown='true'] {
		filter: grayscale(75%);
	}

	#area {
		position: absolute;
		top: calc(var(--y) * var(--pixel-size));
		left: calc(var(--x) * var(--pixel-size));
		width: calc(var(--width) * var(--pixel-size));
		height: calc(var(--height) * var(--pixel-size));
		background: #802929b6;
		outline: 1px solid #ff0000;
		z-index: 1;
	}
</style>
