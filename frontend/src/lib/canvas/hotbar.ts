import { writable, type Writable } from "svelte/store";

export const hotbar: Writable<Record<number, number>> = writable({  
    0: 81,
});

export const selected = writable(0);