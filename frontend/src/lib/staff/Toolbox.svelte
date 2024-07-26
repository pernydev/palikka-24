<script lang="ts">
	import { initialLoad } from "$lib/canvas/grid";
	import { tool } from "$lib/canvas/tool";
	import { Brush, Scan } from "lucide-svelte";
</script>

<div class="toolbar" data-loaded={$initialLoad}>
    <button data-selected={$tool === "paint"} onclick={() => $tool = "paint"}>
        <Brush />
    </button>
    <button data-selected={$tool === "area"} onclick={() => $tool = "area"}>
        <Scan />
    </button>
</div>

<style>
	.toolbar {
		position: fixed;
		left: 1rem;
		top: 50%;
		transform: translateY(-50%);

		display: flex;
        flex-direction: column;
		justify-content: center;
		align-items: center;
		gap: 0.5rem;
		padding: 0.5rem;
		background: #302358bd;
		backdrop-filter: blur(0.5rem);
		border-radius: 0.5rem;

        opacity: 0;
        transition: opacity 3s;
	}

    .toolbar[data-loaded="true"] {
        opacity: 1;
    }
    
    button {
        all: unset;
        cursor: pointer;
        color: white;
        background: #403d5877;
        border-radius: 0.5rem;
        padding: 0.5rem;
        display: flex;
        justify-content: center;
        align-items: center;
        transition: background 0.5s;
    }

    button[data-selected="true"] {
        background: #403d58;
    }
</style>
