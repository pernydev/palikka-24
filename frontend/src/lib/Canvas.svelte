<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import panzoom, { type PanZoom } from 'panzoom';
	import { CANVAS_HEIGHT, CANVAS_WIDTH, PIXEL_SIZE } from '../constants';
	import { place } from './canvas/place';
	import { socket } from './canvas/socket';
	import ImageGrid from './ImageGrid.svelte';
	import { getCanvas } from './canvas/grid';

	let canvas: HTMLDivElement;

	let selection = $state({
		x: 0,
		y: 0
	});

	let mouse = {
		x: 0,
		y: 0
	};

	let panzoomInstance: PanZoom;
	let panning = false;

	onMount(() => {
		panzoomInstance = panzoom(canvas, {
			minZoom: 0.5,
			maxZoom: 8,
			bounds: true,
			boundsPadding: 0.1,
			autocenter: true,
			zoomDoubleClickSpeed: -1, // Disable double click zoom
		});

		panzoomInstance.on('zoomend', () => {
			setSelection();
		});

		panzoomInstance.on('pan', () => {
			panning = true;
		});

		panzoomInstance.on('panend', () => {
			panning = false;
			setSelection();
		});

		getCanvas();
	});

	onDestroy(() => {
		socket.close();
	});

	function onClick(event: Event) {
		if (panning) return;
		place(selection.x, selection.y, 0);
	}

	function mouseMove(event: MouseEvent) {
		mouse = {
			x: event.clientX,
			y: event.clientY
		};

		setSelection();
	}

	function setSelection() {
		const rect = canvas.getBoundingClientRect();
		const x = (mouse.x - rect.left) / panzoomInstance.getTransform().scale;
		const y = (mouse.y - rect.top) / panzoomInstance.getTransform().scale;

		selection.x = Math.min(Math.max(0, Math.floor(x / PIXEL_SIZE)), CANVAS_WIDTH);
		selection.y = Math.min(Math.max(0, Math.floor(y / PIXEL_SIZE)), CANVAS_HEIGHT);
	}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<div bind:this={canvas} id="canvas" onmousemove={mouseMove} tabindex="-1" onmouseup={onClick}>
	<img
		id="selection"
		style="--x: {selection.x}; --y: {selection.y};"
		tabindex="-1"
		src="/assets/blocks/0.png"
		alt=""
	/>
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
	}
</style>
