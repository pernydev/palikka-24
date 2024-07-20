import { writable } from "svelte/store";

export const inventoryOpen = writable(false);
export let openInventory = () => {};

export function setOpenInventory(fn: () => void) {
    openInventory = fn;
}
