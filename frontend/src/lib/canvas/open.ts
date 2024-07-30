import { PUBLIC_API_URL } from "$env/static/public";
import { writable } from "svelte/store";
import { end } from "./secret/end";

export const open = writable(false);

export async function fetchOpenState() {
    console.log('fetching open state');
    const res = await fetch(`${PUBLIC_API_URL}/open`);
    const data = await res.json();
    open.set(data.open);
}

export function isOpenChangeMessage(message: Uint8Array): boolean {
    if (message.length !== 1) {
        return false;
    }

    if (message[0] === 2) {
        window.location.reload();
        return true;
    }

    console.log('open state change', (message[0] === (new Uint8Array([3]))[0]));

    if (message[0] === (new Uint8Array([3]))[0]) {
        end.set(true);
        return true;
    }

    return message[0] === 1;
}