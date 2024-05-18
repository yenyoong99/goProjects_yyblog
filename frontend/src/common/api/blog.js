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
        data: data
    })
}

export var UPDATE_BLOG = (id, data) => {
    return instance({
        url: APP_CONFIG.API_URL + `blogs/${id}`,
        method: 'patch',
        data: data
    })
}

export var UPLOAD_BLOG_IMAGES = (files) => {
    const formData = new FormData();
    files.forEach((file) => {
        formData.append('images', file);
    });

    return instance({
        url: APP_CONFIG.API_URL + 'blogs/upload',
        method: 'post',
        data: formData,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

export var DELETE_BLOG = (id) => {
    return instance({
        url: APP_CONFIG.API_URL + `blogs/${id}`,
        method: 'delete',
    })
}
