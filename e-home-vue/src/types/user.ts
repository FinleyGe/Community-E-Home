export interface ILoginInfo {
    emailPhone: string;
    pwd: string;
    method: number;
}

export interface ILoginRequest {
    jwt : string;
}

export interface IRegisterInfo {
    email: string;
    phone: string;
    pwd: string;
    name: string;
    code: string;
    type: number;
}

export interface IRequest {
    message: string;
}

export interface IVerifyInfo {
    email: string;
}

export interface IUserInfo {
    email: string;
    name: string;
    type: number;
    gender: number;
    province: string;
    city: string;
    phone: string;
    profile: string;
}