import { PUBLIC_API_URL } from '$env/static/public';
import { get } from 'svelte/store';
import { grid } from './grid';
import { hotbar, selected } from './hotbar';

export async function place(x: number, y: number) {
	const texture = get(hotbar)[get(selected)];
	grid.update((grid) => {
		if (texture === undefined) {
			delete grid[`${x},${y}`];
			return grid;
		} else {
			return { ...grid, [`${x},${y}`]: texture };
		}
	});

	const body = new Uint8Array([x, y, texture]);

	const response = await fetch(`${PUBLIC_API_URL}/place`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/octet-stream'
		},
		body
	});

	if (!response.ok) {
		console.error('Failed to place block');
		if (get(grid)[`${x},${y}`] === texture) {
			grid.update((grid) => {
				delete grid[`${x},${y}`];
				return grid;
			});
		}
	}
}
