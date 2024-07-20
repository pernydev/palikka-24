<script lang="ts">
	import { CANVAS_HEIGHT, CANVAS_WIDTH, PIXEL_SIZE } from '../constants';

	let { children } = $props();
</script>

<div
	id="rootl"
	tabindex="-1"
	style="--canvas-width: {CANVAS_WIDTH}; --canvas-height: {CANVAS_HEIGHT}; --pixel-size: {PIXEL_SIZE}px;"
>
	{@render children()}
</div>
