import { ComponentCustomProperties } from "vue";
import {Store} from 'vuex'

declare module "@vue/runtime-core" {
    interface UserState {
        name : string;
        pwd : string;
        avatarUrl : string;
        type : number;
        jwt : string;
    }
    interface ComponentCustomProperties{
        $store : Store<UserState>;
    }
}