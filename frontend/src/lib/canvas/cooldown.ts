import { get, writable } from "svelte/store";
import { end } from "./secret/end";

export const cooldown = writable(0);

export function startCooldown() {
    if (get(end)) return;
    cooldown.set(5);
    const interval = setInterval(() => {
        cooldown.update((c) => c - 1);
        if (get(cooldown) <= 0) {
            clearInterval(interval);
        }
    }, 1000);
}