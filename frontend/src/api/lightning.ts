import {apiFetchJson} from "./helper/fetchHelper";

export const fetchLightbulbs = async () =>
    await apiFetchJson("http://127.0.0.1:3000/api/lightning/lightbulb")

export const updateLightbulbStatus = async (id: number) => {
    await apiFetchJson(`http://127.0.0.1:3000/api/lightning/lightbulb/${id}/state`, {
        method: 'POST',
        credentials: 'include',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ on: false })
    })
}