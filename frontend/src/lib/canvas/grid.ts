import { get, writable, type Writable } from 'svelte/store';
import './socket.ts';
import { PUBLIC_API_URL } from '$env/static/public';
import { fetchOpenState } from './open.js';
import { connected } from './socket';

export const grid: Writable<Record<string, number>> = writable({});
export const initialLoad = writable(false);

export async function init() {
    await Promise.all([fetchOpenState(), getCanvas()]);
    initialLoad.set(true);
}

export async function getCanvas() {
    const resp = await fetch(`${PUBLIC_API_URL}/grid`);
    if (!resp.ok) {
        connected.set(false);
        return;
    }
    const data = await resp.arrayBuffer();
    const gridData = new Uint8Array(data);
    parseGridUpdate(gridData);
}

export function parseGridUpdate(data: Uint8Array) {
    for (let i = 0; i < data.length; i += 3) {
        const x = data[i];
        const y = data[i + 1];
        const texture = data[i + 2];

        switch (texture) {
            case 0:
                console.log('Deleting block');
                grid.update((grid) => {
                    delete grid[`${x},${y}`];
                    return grid;
                });
                break;
            default:
                grid.update((grid) => {
                    grid[`${x},${y}`] = texture;
                    return grid;
                });
                break;
        }
    }

    console.log(get(grid));
}