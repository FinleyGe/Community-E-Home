import {createStore} from 'vuex'

const userConfig = {
    state: () => {
        isLoggedIn : Boolean
    }
}

const store = createStore({
    modules : {
        UserConfig : userConfig,
    }
})