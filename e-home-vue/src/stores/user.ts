import { defineStore } from "pinia";
import {IUserInfo} from '../types/user'
var data:IUserInfo = {
    email: "",
    name: "",
    type: 0,
    gender: 0,
    province: "",
    city: "",
    phone: "",
    profile: "",
}
export const useUserStore = defineStore(
    'user',
    {
        state: () => {
            return { 
                isLogin: false,
                avatarUrl: '',
                jwt: '',
                data: data,
                // todo: and something else
            }
        }
    }
)