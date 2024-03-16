import instance from '../http'
import { APP_CONFIG } from '../config'

export var LIST_BLOG = (params) => {
    return instance.get(APP_CONFIG.API_URL + 'blogs', params)
}

export var GET_BLOG = (id, params) => {
    return instance.get(APP_CONFIG.API_URL + 'blogs/${id}', params)
}

export var CRATE_BLOG = (data) => {
    return instance.post(APP_CONFIG.API_URL + 'blogs/', data)
}

export var UPDATE_BLOG = (data) => {
    return instance.patch(APP_CONFIG.API_URL + 'blogs/${id}', data)
}
