import { InjectionKey, UserState } from 'vue';
import {createStore, Store} from 'vuex'

export const key : InjectionKey<Store<UserState>> = Symbol()
const userState = {
    state : () => ({
        isLoggedIn  : false,
        name        : '',
        avatarUrl   : '',
        type        : -1,
        jwt         : '',
    }),
    mutations: {
        setLogin (state: any) {
            state.isLoggedIn = true;
        },
        setJwt (state: any, jwt: string) {
            state.jwt = jwt;
        }
    },
    action : {
        login (state: any, jwt: string) {
            setLogin(state, jwt);

        }
    }
}

export const store = createStore({
    modules : {
        UserState : userState,
    }
})