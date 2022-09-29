import { writable } from "svelte/store";

export const started = writable(false)
export const username = writable("")
export const canConnect = writable(false)
export const userID = writable("")
export const messages = writable([])
export const serverAddress = "ws://localhost:8080/ws"