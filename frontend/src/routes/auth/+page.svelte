<script>
	import { goto } from '$app/navigation';
	import Backdrop from '$lib/Backdrop.svelte';
	import { onMount } from 'svelte';

    const URL = "https://discord.com/oauth2/authorize?client_id=1266025772976705567&response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A5173%2Fcallback.html&scope=identify+role_connections.write&prompt=none";
    let out = false;

    function openAuthDialog() {
        const popup = window.open(URL, 'auth', 'width=400,height=600');
        popup?.focus();
    }
    
    onMount(() => {
        window.addEventListener('message', (event) => {
            out = true;
            setTimeout(() => {
                goto("/auth/loading?code=" + event.data);
            }, 1000);
        });
    });
</script>

<Backdrop />

<div id="content" data-out={out}>
    <h1>Tervetuloa</h1>
    <p>Kirjaudu sisään piirtääksesi.</p>

    <button on:click={openAuthDialog}>Kirjaudu sisään</button>
</div>


<style>
    #content {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        text-align: center;

        font-family: "system-ui", sans-serif;
        color: white;

        background: #221e42;
        padding: 2em;
        width: 70vw;
        max-width: 50rem!important;
        border-radius: 1rem;

        animation: in 0.9s forwards;
    }

    #content[data-out="true"] {
        animation: out 0.9s forwards;
    }

    @keyframes in {
        from {
            transform: translate(-50%, -20%) scale(0.9);
            opacity: 0;
        }
        to {
            transform: translate(-50%, -50%) scale(1);
            opacity: 1;
        }
    }

    @keyframes out {
        from {
            transform: translate(-50%, -50%) scale(1);
            opacity: 1;
        }
        to {
            transform: translate(-50%, -20%) scale(0.9);
            opacity: 0;
        }
    }

    button {
        all: unset;
        cursor: pointer;
        padding: 0.5em 1em;
        background: #403d58;
        color: white;
        border-radius: 0.5rem;
        transition: background 0.5s;
    }

</style>
