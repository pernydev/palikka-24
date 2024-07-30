import { PUBLIC_WS_URL } from '$env/static/public';
import { writable } from 'svelte/store';
import { blobToArrayBuffer } from '../..';
import { getCanvas, parseGridUpdate } from './grid';
import { open } from './open';
import { end } from './secret/end';
import { hotbar } from './hotbar';


export const connected = writable(false);
let socket: WebSocket;
export function connect() {
	socket = new WebSocket(PUBLIC_WS_URL);
	socket.onopen = () => {
		connected.set(true);
		console.log('WebSocket connection established');
	};

	/*
The socket sends custom binary messages to the client. Here is what the server sends:
- X Coordinate (1 byte)
- Y Coordinate (1 byte)
- Texture Index (1 byte)

And that's it. A packet can contain multiple of those sets of data. Nothing is seprating the sets, but the sets are always 3 bytes long.

BUT! A packet can also be an open state change, in this case the length of the packet is 1 byte, and the byte is either 0 or 1.
*/
	socket.onmessage = async (event) => {
		console.log('Received message from WebSocket');
		const arrayBuffer = await blobToArrayBuffer(event.data);
		const data = new Uint8Array(arrayBuffer);
		if (executeEvent(data)) {
			return;
		}
		parseGridUpdate(data);
	};

	socket.onclose = () => {
		connected.set(false);
        setTimeout(() => {
            connect();
        }, 1000);
	};
}

export function disconnect() {
	socket.close();
}

export function executeEvent(message: Uint8Array): boolean {
    if (message.length !== 1) {
        return false;
    }

    if (message[0] === (new Uint8Array([2]))[0]) {
		getCanvas();
        return true;
    }
	
    if (message[0] === (new Uint8Array([3]))[0]) {
        end.set(true);
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
        return true;
    }

	if (message[0] === 1 || message[0] === 0) {
		open.set(message[0] === 1);
		return true;
	}

	return false;
}