<script lang="ts">
	import { onMount } from "svelte";

    const last = 226;

    let images: number[] = $state([]);

    onMount(() => {
        for (let i = 1; i < last; i++) {
            images.push(i);
        }
    });
    
    function moveUp(i: number) {
        const index = images.indexOf(i);
        if (index === 0) return;
        const temp = images[index - 1];
        images[index - 1] = images[index];
        images[index] = temp;
    }

    function moveDown(i: number) {
        const index = images.indexOf(i);
        if (index === images.length - 1) return;
        const temp = images[index + 1];
        images[index + 1] = images[index];
        images[index] = temp;
    }

    function exportOrder() {
        const order = images.map((i, index) => `${i}:${index + 1}`).join(',');
        const orderJson = JSON.stringify({ order });
        const orderEncoded = btoa(orderJson);
        
        navigator.clipboard.writeText(orderEncoded);
        console.log(orderEncoded);
    }

</script>

<button onclick={exportOrder}>Export order</button>

<ul>
    {#each images as i}
        <li>
            <img src={`/assets/blocks/${i}.png`} alt="" />
            <button onclick={() => moveUp(i)}>Up</button>
            <button onclick={() => moveDown(i)}>Down</button>
        </li>
    {/each}
</ul>

<style>
    img {
        width: 3rem;
        height: 3rem;
        image-rendering: pixelated;
    }

    li {
        display: flex;
        align-items: center;
        gap: 1rem;
    }

    button {
        height: 3rem;
    }
</style>