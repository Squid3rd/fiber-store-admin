const BASE_URL = "http://localhost:8080/api/v1"

function readStringEnv(key: string, fallback: string){
    const value = import.meta.env[key]

    return typeof value === 'string' && value.trim() !== '' ? value : fallback
}

export const env = {
    apiBaseUrl : readStringEnv('VITE_API_URL', BASE_URL)
}
