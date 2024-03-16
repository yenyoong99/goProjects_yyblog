import instance from '../http'
import { APP_CONFIG } from '../config'

export var LIST_BLOG = (params) => {
    return instance({
        url: APP_CONFIG.API_URL + 'blogs',
        method: 'get',
        params: params
    })
}

export var GET_BLOG = (id, params) => {
    return instance({
        url: APP_CONFIG.API_URL + `blogs/${id}`,
        method: 'get',
        params: params
    })
}

export var CRATE_BLOG = (data) => {
    return instance({
        url: APP_CONFIG.API_URL + 'blogs/',
        method: 'post',
        params: data
    })
}

export var UPDATE_BLOG = (id, data) => {
    return instance({
        url: APP_CONFIG.API_URL + `blogs/${id}`,
        method: 'patch',
        params: data
    })
}

export var DELETE_BLOG = (id) => {
    return instance({
        url: APP_CONFIG.API_URL + `blogs/${id}`,
        method: 'delete',
    })
}
