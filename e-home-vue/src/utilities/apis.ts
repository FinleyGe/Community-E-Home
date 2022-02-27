import axios from 'axios';
import 
const url = `http://127.0.0.1:8989`
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

export function login(userData: IUserLoginInfo):void {
    axios.post(url + '/api/login', userData).
    then((response) => {
        console.log(response);
        response.data.jwt
    }).
    catch((error) => {
        console.log(error);
    })
}

export function EditUserInfo(userData: IEditUserInfo):void {
    // TODO
}

export function ValidEmail(userData: IValidEmail): void {
    // TODO
}
