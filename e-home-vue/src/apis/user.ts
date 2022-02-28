import request from '../utilities/request'
import { ILoginInfo, ILoginRequest } from '../types/user'
export const LoginApi = async (loginInfo: ILoginInfo) => {
    const res = await request<ILoginRequest>({
        url: '/api/login',
        method: 'POST',
        data: loginInfo
    })
    return res
}