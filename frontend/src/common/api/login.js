import instance from '../http'
import { APP_CONFIG } from '../config'

export var LOGIN = (data) => {
    return instance.post(APP_CONFIG.APP_NAME + APP_CONFIG.API_PREFIX + 'tokens', data)
}
