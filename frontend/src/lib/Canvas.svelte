<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import panzoom, { type PanZoom } from 'panzoom';
	import { CANVAS_HEIGHT, CANVAS_WIDTH, PIXEL_SIZE } from '../constants';
	import { place } from './canvas/place';
	import { socket } from './canvas/socket';
	import ImageGrid from './ImageGrid.svelte';
	import { getCanvas } from './canvas/grid';
	import { hotbar, selected } from './canvas/hotbar';
	import { inventoryOpen } from './canvas/inventory';
	import { tool } from './canvas/tool';

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
			zoomDoubleClickSpeed: 1 // Disable double click zoom
		});

		getCanvas();
	});

	$effect(() => {
		if (!panzoom) return;
		if ($inventoryOpen) panzoomInstance.pause();
		else panzoomInstance.resume();
	});

	onDestroy(() => {
		socket.close();
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
	}

	function mouseUp(event: Event) {
		if (isMouseDown) {
			switch ($tool) {
				case 'paint':
					place(selection.x, selection.y);
					break;
				case 'area':
					console.log(area);
					if (area.step === 0) {
						area = {
							x: selection.x,
							y: selection.y,
							width: 1,
							height: 1,
							step: 1
						};
					} else {
						area.width = selection.x - area.x + 1;
						area.height = selection.y - area.y + 1;
						area.step = 0;
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
>
	{#if $tool === 'paint'}
		<div
			id="selection"
			style="--x: {selection.x}; --y: {selection.y}; --src: url('/assets/blocks/{$hotbar[
				$selected
			]}.png');"
			tabindex="-1"
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
