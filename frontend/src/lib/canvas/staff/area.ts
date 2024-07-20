import { PUBLIC_API_URL } from "$env/static/public";

export async function deleteArea(from_x: number, from_y: number, width: number, height: number) {
    const payload = [];
    for (let y = from_y; y < from_y + height; y++) {
        for (let x = from_x; x < from_x + width; x++) {
            console.log(`Deleting block at ${x},${y}`);
            payload.push(x);
            payload.push(y);
            payload.push(0);
        }
    }
    
    await fetch(`${PUBLIC_API_URL}/execute`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/octet-stream',
        },
        body: new Uint8Array(payload),
    });
}