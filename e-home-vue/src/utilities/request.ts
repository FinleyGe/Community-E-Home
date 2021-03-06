import axios, { AxiosRequestConfig } from 'axios';
import config from '../config';
import {ResponseType} from '../types/apis'
// const request = axios.create({
//     baseURL: 'http://127.0.0.1:8989',
// })
const instance = axios.create({
    baseURL: 'http://127.0.0.1:8989'
})

const request = async <T = any>(config:AxiosRequestConfig): Promise<ResponseType<T>> => {
    try{
        const {data} = await instance.request<ResponseType<T>>(config)
        console.log(data)
        return data
    } catch (err:any) {
        const message = err.message || '请求失败'
        console.error(message)
        alert(message)
        return {
            message: message,
            data: null as any
        }
    }
}

export default request