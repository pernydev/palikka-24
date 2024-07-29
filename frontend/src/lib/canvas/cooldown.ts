import { get, writable } from "svelte/store";

export const cooldown = writable(0);

export function startCooldown() {
    cooldown.set(5);
    const interval = setInterval(() => {
        cooldown.update((c) => c - 1);
        if (get(cooldown) <= 0) {
            clearInterval(interval);
        }
    }, 1000);
}