import { writable } from 'svelte/store'

function createLists() {
	const { subscribe, set, update } = writable([
		{ id: 1, name: 'First list', people: [] },
		{ id: 2, name: 'Second list', people: [] }
	])

	return {
		subscribe,
		addList: (id: number, name: string) => {
			update(prev => [...prev, { id, name: name.trim(), people: [] }])
		},
		joinList: (listID: number, realName: string, nbaName: string) => {
			update(lists =>
				lists.map(list => {
					if (list.id === listID) {
						return { ...list, people: [...list.people, { realName, nbaName }] }
					}
					return list
				})
			)
		}
	}
}

export const lists = createLists()
