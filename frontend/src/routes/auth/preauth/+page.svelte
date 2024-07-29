<script lang="ts">
	import { goto } from "$app/navigation";
	import { page } from "$app/stores";
	import { onMount } from "svelte";

    onMount(() => {
        const token = new URL($page.url).searchParams.get('token');
        const staff = new URL($page.url).searchParams.get('staff');
        localStorage.setItem('token', token || "");
        localStorage.setItem('staff', staff || "");

        goto('/');
    });
</script>