<script lang="ts">
	import { onMount } from 'svelte';
	import { CANVAS_HEIGHT, CANVAS_WIDTH, PIXEL_SIZE } from '../constants';
	import { grid } from './canvas/grid';
	import { getTexture } from './canvas/texture';

	let canvas: HTMLCanvasElement;
	let ctx: CanvasRenderingContext2D;

	onMount(async () => {
		ctx = canvas.getContext('2d') as CanvasRenderingContext2D;
		setTimeout(() => {
			reRender($grid);
            setTimeout(() => {
                reRender($grid);
            }, 100);
		}, 100);
	});

	function reRender(canvasGrid: Record<string, number>) {
		console.log('reRender');
		if (!ctx) {
			console.log('no ctx');
			return;
		}
		// clear canvas
		ctx.clearRect(0, 0, CANVAS_WIDTH * PIXEL_SIZE, CANVAS_HEIGHT * PIXEL_SIZE);

		ctx.fillStyle = 'white';
		ctx.fillRect(0, 0, CANVAS_WIDTH * PIXEL_SIZE, CANVAS_HEIGHT * PIXEL_SIZE);

		Object.entries(canvasGrid).forEach(([coords, texture]) => {
			const [x, y] = coords.split(',').map(Number);
			console.log(x, y, texture);
			ctx.drawImage(getTexture(texture), x * PIXEL_SIZE, y * PIXEL_SIZE, PIXEL_SIZE, PIXEL_SIZE);
		});
	}

	grid.subscribe((v) => {
		if (!ctx) return;
		reRender(v);
	});
</script>

<canvas bind:this={canvas} width={CANVAS_WIDTH * PIXEL_SIZE} height={CANVAS_HEIGHT * PIXEL_SIZE}
></canvas>

<style>
	canvas {
		image-rendering: pixelated;
	}
</style>
