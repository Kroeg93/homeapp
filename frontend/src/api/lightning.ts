import {apiFetchJson} from "./helper/fetchHelper";

export const fetchLightbulbs = async () =>
    await apiFetchJson("http://127.0.0.1:3000/api/lightning/lightbulb")