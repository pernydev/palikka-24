<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import Backdrop from "$lib/Backdrop.svelte";
	import { PUBLIC_API_URL } from '$env/static/public';
	import { goto } from '$app/navigation';

    let out = false;
    let reloadTimeout: any;
    onMount(async () => {
        const resp = await fetch(`${PUBLIC_API_URL}/auth`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/string',
            },
            body: new URLSearchParams(location.search).get('code')
        });

        if (resp.ok) {
            const data = await resp.json();
            localStorage.setItem('token', data.token);
            localStorage.setItem('staff', data.staff ? 'true' : 'false');
            out = true;
            setTimeout(() => {
                goto('/');
            }, 1300);
        } else {
            console.error(await resp.text());
        }

        reloadTimeout = setTimeout(() => {
            location.reload();
        }, 2000);
    });

    onDestroy(() => {
        clearTimeout(reloadTimeout);
    });
</script>


<div id="tranisitoner" data-out={out}></div>

<Backdrop />


<style>
    #tranisitoner {
        position: fixed;
        inset: 0;
        background: #0a012100;
        transition: background 1s;
        z-index: 100;
    }

    #tranisitoner[data-out="true"] {
        background: #0a0121;
    }
</style>

