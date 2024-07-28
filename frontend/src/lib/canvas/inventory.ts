import { writable } from "svelte/store";

export const inventoryOpen = writable(false);
export let openInventory = () => {};
export let closeInventory = () => {};

export function setOpenInventory(fn: () => void) {
    openInventory = fn;
}

export function setCloseInventory(fn: () => void) {
    closeInventory = fn;
}