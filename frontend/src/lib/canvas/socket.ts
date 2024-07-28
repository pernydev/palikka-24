import { goto } from "$app/navigation";
import { PUBLIC_WS_URL } from "$env/static/public";
import { writable } from "svelte/store";
import { blobToArrayBuffer } from "../..";
import { parseGridUpdate } from "./grid";
import { isOpenChangeMessage, open } from "./open";

console.log("Connecting to WebSocket");
export let socket = new WebSocket(PUBLIC_WS_URL);
export const connected = writable(false);

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
    if (isOpenChangeMessage(data)) {
        open.set(data[0] === 1);
        return;
    }
    parseGridUpdate(data);
};

let reconnectAttempts = 0;

socket.onclose = () => {
    connected.set(false);

    if (reconnectAttempts > 5) goto("/offline");
    setTimeout(() => {
        reconnectAttempts++;
        socket = new WebSocket(PUBLIC_WS_URL);
    }, 1000);
};
