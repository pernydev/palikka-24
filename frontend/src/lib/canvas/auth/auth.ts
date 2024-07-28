import { writable } from 'svelte/store';

export const loggedIn = writable(true);

export function checkIsLoggedIn() {
	//if (localStorage.getItem('token')) loggedIn.set(true);
}
