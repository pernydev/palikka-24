<script lang="ts">
	import { CANVAS_HEIGHT, CANVAS_WIDTH, PIXEL_SIZE } from '../constants';
	import './app.css';

	let { children } = $props();
</script>

<div
	id="rootl"
	style="--canvas-width: {CANVAS_WIDTH}; --canvas-height: {CANVAS_HEIGHT}; --pixel-size: {PIXEL_SIZE}px;"
>
	{@render children()}
</div>