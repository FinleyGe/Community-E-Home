export function isEmail(email: string):boolean {
    var reg = /^[0-9a-zA-Z_.-]+[@][0-9a-zA-Z_.-]+([.][a-zA-Z]+){1,2}$/;
    return reg.test(email);
}

export function isPhone(phone: string):boolean {
    var reg = /[0-9]{11}/;
    return reg.test(phone);
}

export function isPwdValid(pwd: string):boolean {
    var reg = /^[a-zA-Z]\w{5,17}$/
    return reg.test(pwd)
}

export function isEmailOrPhone(emailOrPhone: string):number {
    if(isEmail(emailOrPhone)){
        return 0;
    }else if (isPhone(emailOrPhone)){
        return 1;
    }else{
        return -1;
    }
}