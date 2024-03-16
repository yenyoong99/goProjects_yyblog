import axios from "axios";

const instance = axios.create({
    timeout: 5000,
    headers: {'Content-Type': 'application/json'},
    withCredentials: true
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
