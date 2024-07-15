<script lang="ts">
	import { onMount } from 'svelte';
	import panzoom, { type PanZoom } from 'panzoom';
	import { CANVAS_HEIGHT, CANVAS_WIDTH, PIXEL_SIZE } from '../constants';
	import { grid } from './canvas';

	let canvas: HTMLDivElement;

	let selection = $state({
		x: 0,
		y: 0
	});

    let panzoomInstance: PanZoom;

	onMount(() => {
		panzoomInstance = panzoom(canvas, {
			minZoom: 0.5,
			maxZoom: 4,
			bounds: true,
			boundsPadding: 0.1,
			autocenter: true
		});
	});

	function mouseMove(event: MouseEvent) {
		const rect = canvas.getBoundingClientRect();
        const x = (event.clientX - rect.left) / panzoomInstance.getTransform().scale;
        const y = (event.clientY - rect.top) / panzoomInstance.getTransform().scale;

        selection.x = Math.min(Math.max(0, Math.floor(x / PIXEL_SIZE)), CANVAS_WIDTH);
        selection.y = Math.min(Math.max(0, Math.floor(y / PIXEL_SIZE)), CANVAS_HEIGHT);
	}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div bind:this={canvas} id="canvas" onmousemove={mouseMove} tabindex="-1">
	<div id="selection" style="--x: {selection.x}; --y: {selection.y};"></div>
    {#each Object.entries($grid) as [coords, value]}
        <div style="--x: {coords.split(",")[0]}; --y: {coords.split(",")[1]}; background-color: {value};"></div>
    {/each}
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

	#canvas > * {
		position: absolute;
		top: calc(var(--y) * var(--pixel-size));
		left: calc(var(--x) * var(--pixel-size));
		width: var(--pixel-size);
		height: var(--pixel-size);
        z-index: -1;
	}

	#selection {
		background-color: #ff0000;
        opacity: 0.5;
        z-index: 0;
         
	}
</style>
