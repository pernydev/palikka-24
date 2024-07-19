import { writable, type Writable } from 'svelte/store';
import './socket.ts';
import { PUBLIC_API_URL } from '$env/static/public';

export const grid: Writable<Record<string, number>> = writable({});

export async function getCanvas() {
    const resp = await fetch(`${PUBLIC_API_URL}/grid`);
    const data = await resp.arrayBuffer();
    const gridData = new Uint8Array(data);
    parseGridUpdate(gridData);
}

export function parseGridUpdate(data: Uint8Array) {
    const gridChange: Record<string, number> = {};

    for (let i = 0; i < data.length; i += 3) {
        const x = data[i];
        const y = data[i + 1];
        const texture = data[i + 2];

        gridChange[`${x},${y}`] = texture;
        console.log(`Setting ${x},${y} to ${texture}`);
    }

    grid.update((grid) => {
        return { ...grid, ...gridChange };
    });
}