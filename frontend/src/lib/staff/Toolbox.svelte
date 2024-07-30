<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	import { initialLoad } from '$lib/canvas/grid';
	import { open } from '$lib/canvas/open';
	import { tool } from '$lib/canvas/tool';
	import { Brush, FileQuestion, Lock, Scan, Unlock } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let showTools = $state(false);
	let isGod = $state(false);

	onMount(() => {
		showTools = (localStorage.getItem('staff') ?? 'false') === 'true';
		isGod = (localStorage.getItem('god') ?? 'false') === 'true';
	});

	function sendLockStatus() {
		fetch(`${PUBLIC_API_URL}/open`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: localStorage.getItem('token') ?? ''
			}
		});
	}

	function sendEndStatus() {
		fetch(`${PUBLIC_API_URL}/end`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: localStorage.getItem('token') ?? ''
			}
		});
	}
</script>

{#if showTools}
	<div class="toolbar" data-loaded={$initialLoad}>
		<button data-selected={$tool === 'paint'} onclick={() => ($tool = 'paint')}>
			<Brush />
		</button>
		<button data-selected={$tool === 'area'} onclick={() => ($tool = 'area')}>
			<Scan />
		</button>
		{#if isGod}
			<button onclick={() => sendLockStatus()}>
				{#if $open}
					<Lock />
				{:else}
					<Unlock />
				{/if}
			</button>
			<button onclick={() => (sendEndStatus())}>
				<FileQuestion />
			</button>
		{/if}
	</div>
{/if}

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

	.toolbar[data-loaded='true'] {
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

	button[data-selected='true'] {
		background: #403d58;
	}
</style>
