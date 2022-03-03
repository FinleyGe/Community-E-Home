import request from '../utilities/request'
import {    ILoginInfo, ILoginRequest, 
            IRegisterInfo, IRequest, 
            IVerifyInfo, IUserInfo } from '../types/user'
export const LoginApi = async (loginInfo: ILoginInfo) => {
    const res = await request<ILoginRequest>({
        url: '/api/login',
        method: 'POST',
        data: loginInfo
    })
    return res
}

export const RegisterApi = async (registerInfo: IRegisterInfo) => {
    const res = await request<IRequest>({
        url: '/api/register',
        method: 'POST',
        data: registerInfo
    })
    return res
}
export const VerifySendApi = async (verifyInfo: IVerifyInfo) => {
    return await request<IRequest>({
        url: '/api/user/vertify',
        method: 'POST',
        data: verifyInfo
    })
}
export const GetUserInfoApi = async (jwt: string) => {
    return await request<IUserInfo>({
        url: '/api/user/info',
        method: 'GET',
        headers:{
            Authorization: jwt,
        },
    })
}

export const UpdateUserInfoApi = async (userinfo:IUserInfo, jwt: string) => {
    return await request<IRequest>({
        url: '/api/user/info',
        method: 'POST',
        headers:{
            Authorization: jwt
        },
        data: userinfo,
    })
}