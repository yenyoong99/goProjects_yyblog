import axios from "axios";
import { Message } from '@arco-design/web-vue'


const instance = axios.create({
    timeout: 5000,
    headers: {'Content-Type': 'application/json'},
    withCredentials: true
});

instance.interceptors.response.use(
    (response) => {
        return response.data
    },
    (err) => {
        // console.log(err)
        var msg = err.message
        if (err.response.data && err.response.data.message) {
            msg = err.response.data.message
            if (err.response.data.code === 401) {
                window.location.assign('/login')
            }
        }

        Message.error({
            content: msg,
            duration: 2000
        })
        return Promise.reject(err)
    }
)

export default instance
