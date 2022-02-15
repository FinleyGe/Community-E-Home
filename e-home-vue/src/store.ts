import {createStore} from 'vuex'

const userState = {
    state : () => ({
        isLoggedIn  : Boolean,
        name        : String,
        avatarUrl   : String,
        type        : String,

    }),
    mutations: {
        login (state: any) {
            state.isLoggedIn = true
        },
        logout (state: any) {
            state.isLoggedIn = false
        },
        
    }
}

const store = createStore({
    modules : {
        UserState : userState,
    }
})