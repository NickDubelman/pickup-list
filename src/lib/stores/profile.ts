import { writable } from 'svelte/store'

export interface Profile {
	realName: string
	nbaName: string
}

function createProfile() {
	const { subscribe, set, update } = writable<Profile>(null)

	return { subscribe, set }
}

export const profile = createProfile()
