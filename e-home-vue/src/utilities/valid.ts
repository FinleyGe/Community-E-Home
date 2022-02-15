export function isEmail(email: string):Boolean {
    var reg = /^[0-9a-zA-Z_.-]+[@][0-9a-zA-Z_.-]+([.][a-zA-Z]+){1,2}$/;
    return reg.test(email);
}

export function isPhone(phone: string):Boolean {
    var reg = /[0-9]{11}/;
    return reg.test(phone);
}
