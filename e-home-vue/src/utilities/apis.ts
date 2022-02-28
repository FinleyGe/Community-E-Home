import axios, { AxiosRequestConfig } from 'axios';

const url = `http://127.0.0.1:8989`
// const store = useStore(key)

export interface IUserInfo {
    name    : string;
    pwd     : string;
    email   : string;
    phone   : string;
    type    : number;
}

export interface IUserLoginInfo {
    emailPhone : string;
    method : number;
    pwd   : string;
}

export interface IEditUserInfo {
    profile : string;
}

export interface IValidEmail {
    code : string;
}

const request = async (config: AxiosRequestConfig): Promise<> => {}

export function register(userData: IUserInfo):void {
    axios.post(url + '/api/register', userData)
    .then((response) => {
        alert('注册成功');
        console.log(response);
        // TODO change to edit the user info page
    })
    .catch((error) => {
        if(error.response.status === 406) {
            alert('用户已经存在！')
        }else{
            alert('其他错误！')
        }
    })
}

export function login(userData: IUserLoginInfo):string {
    axios.post(url + '/api/login', userData).
    then((response) => {
        // console.log(response.data.jwt);
        var jwt: string = response.data.jwt;
        return jwt
    }).
    catch((error) => {
        console.log(error);
        return ''
    })
}

export function EditUserInfo(userData: IEditUserInfo):void {
    // TODO
}

export function ValidEmail(userData: IValidEmail): void {
    // TODO
}
