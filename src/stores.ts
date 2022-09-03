import { writable } from 'svelte-local-storage-store'

export const history = writable('history', {})
