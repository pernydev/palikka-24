<script lang="ts">
	import { goto } from '$app/navigation';
	import { loggedIn } from './canvas/auth/auth';
	import { initialLoad } from './canvas/grid';
	import { inventoryOpen } from './canvas/inventory';
	import { open } from './canvas/open';
	import HotbarItem from './HotbarItem.svelte';
</script>

{#if $open || !$loggedIn}
	<div class="hotbar" data-loaded={$initialLoad} data-loggedin={$loggedIn}>
		{#if $loggedIn}
			<HotbarItem slot="menu" />
			<span class="seprator"></span>
			{#each Array.from({ length: 9 }) as _, i}
				<HotbarItem slot={i} />
			{/each}
			<span class="seprator second"></span>
			<HotbarItem slot="remove" />
		{:else}
			<p>Kirjaudu sisään piirtääksesi</p>
			<a href="/auth">Kirjaudu sisään</a>
		{/if}
	</div>
{/if}

<style>
	.hotbar {
		position: fixed;
		bottom: 2rem;
		left: 50%;
		transform: translateX(-50%) translateY(200%);

		display: flex;
		justify-content: center;
		align-items: center;
		gap: 0.5rem;
		padding: 0.5rem;
		background: #302358bd;
		backdrop-filter: blur(0.5rem);
		border-radius: 0.5rem;

		z-index: 1000;
	}

	.hotbar[data-loaded='true'] {
		animation: in 0.5s 1s forwards;
	}

	.hotbar[data-loggedin='false'] {
		color: white;
		padding-inline: 1rem;
		display: flex;
		flex-direction: column;
	}

	p {
		margin: 0;
	}

	a {
		all: unset;
		cursor: pointer;
		color: white;
		background: #403d5877;
		border-radius: 0.5rem;
		padding: 0.5rem;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	@keyframes in {
		from {
			transform: translateX(-50%) translateY(100%);
			opacity: 0;
		}
		to {
			transform: translateX(-50%) translateY(0);
			opacity: 1;
		}
	}

	.seprator {
		width: 2px;
		height: 2rem;
		background: #fff;
		opacity: 0.5;
		margin-inline: 0.5rem;
	}

	@media (max-width: 768px) {
		.seprator {
			display: none;
		}
	}
</style>
