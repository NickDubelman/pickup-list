import { writable } from 'svelte/store'

export const lists = writable([
	{ id: 1, name: 'First list' },
	{ id: 2, name: 'Second list' }
])
