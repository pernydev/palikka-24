import { PUBLIC_API_URL } from "$env/static/public";
import { grid } from "./grid";

export async function place(x: number, y: number, texture: number) {
	grid.update((grid) => {
		return { ...grid, [`${x},${y}`]: texture };
	});
	
    const body = new Uint8Array([x, y, texture]);
    
	const response = await fetch(`${PUBLIC_API_URL}/place`, {
		method: 'POST',
		headers: {
            'Content-Type': 'application/octet-stream',
		},
        body,
	});

	if (!response.ok) {
		console.error('Failed to place block');
	}
}
