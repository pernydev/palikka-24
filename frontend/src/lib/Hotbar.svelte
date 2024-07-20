<script lang="ts">
	import { inventoryOpen } from './canvas/inventory';
	import HotbarItem from './HotbarItem.svelte';
</script>

<div class="hotbar">
	<HotbarItem slot="menu" />
    <span class="seprator"></span>
	{#each Array.from({ length: 9 }) as _, i}
    <HotbarItem slot={i} />
	{/each}
	<span class="seprator"></span>
	<HotbarItem slot="remove" />
</div>

<style>
	.hotbar {
		position: fixed;
		bottom: 2rem;
		left: 50%;
		transform: translateX(-50%);

		display: flex;
		justify-content: center;
		align-items: center;
		gap: 0.5rem;
		padding: 0.5rem;
		background: #302358bd;
		backdrop-filter: blur(0.5rem);
		border-radius: 0.5rem;
	}

    .seprator {
        width: 2px;
        height: 2rem;
        background: #fff;
        opacity: 0.5;
        margin-inline: 0.5rem;
    }
</style>
