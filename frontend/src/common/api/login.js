import instance from '../http'
import { APP_CONFIG } from '../config'

export const LOGIN = (data) => {
    return instance({
        url: APP_CONFIG.API_URL + 'tokens/',
        method: 'post',
        data: data
    })
}
