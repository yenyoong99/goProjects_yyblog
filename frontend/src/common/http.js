import axios from "axios";
import { APP_CONFIG } from './config.js'

const instance = axios.create({
    baseURL: APP_CONFIG.DOMAIN_NAME,
    timeout: 5000,
    headers: {'Content-Type': 'application/json'}

});

instance.interceptors.response.use(
    (resp) => {
        return resp.data
    },
    (error) => {
        console.log(error)
        var message = error.message
        if (error.response && error.response.data) {
            message = error.response.data.message
        }
        return Promise.reject(new Error(message));
    }
)

export default instance
