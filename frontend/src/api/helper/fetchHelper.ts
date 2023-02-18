export const apiFetch = async (url: any, opts: any = {}) => {
    opts = Object.assign({
        credentials: 'include',
        headers: Object.assign({ 'Content-Type': 'application/json' }, (opts.headers || {}))
    }, opts)


    console.log(url)

    return await fetch(url, opts).then(async response => {
        if (response.status >= 400)
            return response.json().then(
                json => Promise.reject(new Error(`Error: ${json}`)),
                _ => Promise.reject(new Error("ERROR: fetch"))
            )
        else
            return Promise.resolve(response)
    })
}

export const apiFetchJson = async(url: any, opts: any = {}) => {
    const response = await apiFetch(url, opts)

    try {
        return await response.json()
    } catch(e) {
        throw new Error("ERROR: apiFetchJson")
    }
}
