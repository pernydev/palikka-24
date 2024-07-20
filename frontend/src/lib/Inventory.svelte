<script lang="ts">
	import { hotbar, selected } from './canvas/hotbar';
	import { onMount } from 'svelte';
	import { inventoryOpen, setOpenInventory } from './canvas/inventory';

	let dialog: HTMLDialogElement;
	const itemCount = 250;

	function onKeyUp(event: KeyboardEvent) {
		switch (event.key.toLowerCase()) {
			case 'escape':
				dialog.close();
				break;
			case 'e':
				$inventoryOpen = true;
				if (dialog.open) dialog.close();
				else dialog.showModal();
				break;
		}
	}

	onMount(() => {
		setOpenInventory(() => {
			$inventoryOpen = true;
			dialog.showModal();
		});

		dialog.addEventListener('close', () => {
			console.log('close');	
			$inventoryOpen = false;
		});
	});
</script>

<svelte:window on:keyup={onKeyUp} />

<dialog bind:this={dialog}>
	<div class="hotbar">
		{#each Array.from({ length: 9 }) as _, i}
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<button
				ondragover={(event) => {
					event.preventDefault();
				}}
				ondrop={(event) => {
					console.log('drop');
					event.preventDefault();
					const index = parseInt(event.dataTransfer?.getData('text/plain') || '');
					console.log(index);
					$hotbar[i] = index;
				}}
				onclick={() => {
					delete $hotbar[i];
				}}
			>
				{#if $hotbar[i]}
					<img src={`/assets/blocks/${$hotbar[i]}.png`} alt="" />
				{/if}
			</button>
		{/each}
	</div>
	<div class="items">
		{#each Array.from({ length: itemCount }) as _, i}
			<button
				ondragstart={(event) => {
					console.log('dragstart');
					event.dataTransfer?.setData('text/plain', (i + 1).toString());
				}}
				draggable
				onclick={() => {
					for (let h = 0; h < 9; h++) {
						if ($hotbar[h]) continue;
						$hotbar[h] = i + 1;
						$selected = h;
						break;
					}
				}}
			>
				<img src={`/assets/blocks/${i + 1}.png`} alt="" />
			</button>
		{/each}
	</div>
</dialog>

<style>
	dialog {
		color: white;
		border: none;
		border-radius: 1rem;
		background: #34314d;
		backdrop-filter: blur(0.5rem);

		padding: 1rem;
	}

	.hotbar {
		display: flex;
		justify-content: space-between;
		gap: 0.5rem;
		margin-bottom: 1rem;
	}

	.hotbar > button {
		all: unset;
		cursor: pointer;
		width: 4rem;
		height: 4rem;
		border-radius: 0.5rem;
		background: #403d58;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.items {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(3rem, 1fr));
		gap: 0.25rem;
		max-width: 50rem;
	}

	button {
		all: unset;
		cursor: pointer;
	}

	img {
		width: 3rem;
		height: 3rem;
		image-rendering: pixelated;
		border-radius: 0.5rem;
	}
</style>
