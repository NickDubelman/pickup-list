import { writable } from 'svelte/store'
import type { Profile } from './profile'

interface List {
  id: number
  name: string
  people: { realName: string; nbaName: string }[]
}

function createLists() {
  const { subscribe, update, set } = writable<List[]>([])

  return {
    subscribe,
    set,
    addList: (id: number, name: string) => {
      update(prev => [...prev, { id, name: name.trim(), people: [] }])
    },
    joinList: (listID: number, profile: Profile) => {
      update(lists =>
        lists.map(list => {
          if (list.id === listID) {
            return { ...list, people: [...list.people, profile] }
          }
          return list
        })
      )
    }
  }
}

export const lists = createLists()
