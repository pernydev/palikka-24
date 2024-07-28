// Tools are basically just for staff, default tool is "paint" which is just pixel placing. Staff check is done server side so no need to worry about that.

import { writable, type Writable } from "svelte/store";

type Tool = 'paint' | 'area'

export const tool: Writable<Tool> = writable('paint');