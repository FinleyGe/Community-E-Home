import { defineStore } from "pinia";

export const useUserStore = defineStore(
    'user',
    {
        state: () => {
            return { 
                isLogin: false,
                name: '',
                avatarUrl: '',
                jwt: '',
                // todo: and something else
            }
        }
    }
)