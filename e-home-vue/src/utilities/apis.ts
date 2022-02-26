import axios from 'axios';

const url = `http://127.0.0.1:8989`
export interface IUserInfo {
    name    : string;
    pwd     : string;
    email   : string;
    phone   : string;
    type    : number;
}

export function register(userData: IUserInfo):void {
    axios.post(url + '/api/register', userData)
    .then((response) => {
        alert('注册成功');
        console.log(response);
    }).catch((error) => {
        console.log(error);
    })
}