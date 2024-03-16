//  install:  npm i @vueuse/core， Use responsive localStorage objects
//  use state.isLogin
import { useStorage } from '@vueuse/core'

// Help encapsulate LocalStorage into a responsive Ref object
export var state = useStorage(
    'yyblog-store',
    {isLogin: false, username: ''},
    localStorage,
    { mergeDefaults: true }
)
