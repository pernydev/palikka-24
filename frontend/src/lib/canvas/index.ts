import { writable, type Writable } from 'svelte/store';

export const grid: Writable<Record<string, string>> = writable(
    {
        "0,0": "orange",
    }
);