import { writable } from 'svelte/store'

interface List {
	id: number
	name: string
	people: { realName: string; nbaName: string }[]
}

function createLists() {
	const { subscribe, update } = writable<List[]>([])

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
