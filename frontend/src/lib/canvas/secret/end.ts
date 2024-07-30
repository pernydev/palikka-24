import { PUBLIC_API_URL } from "$env/static/public";
import { writable } from "svelte/store";
import { hotbar } from "../hotbar";

export const end = writable(false);

export async function isEnd() {
    const response = await fetch(`${PUBLIC_API_URL}/end`);
    const data = await response.json();
    end.set(data.end);

    if (!data.end) return;

    hotbar.set({
        0: 164,
        1: 164,
        2: 164,
        3: 164,
        4: 164,
        5: 164,
        6: 164,
        7: 164,
        8: 164,
        9: 164,
    });
}