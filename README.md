# Community-E-Home

## API 接口文档

说明：
1. 下方文档中，参数和返回值使用了**非**标准的json语法格式。（标准的json中没有注释）
2. 所有请求使用json格式，请注意Json格式和PostForm的区别。可以参考：https://www.jianshu.com/p/d0a4de05786b
3. 

### 1. [请求] Login 登陆请求

url: `/api/login`
method: `POST`
params:
```json
{
    'email' = string, // 电子邮箱
    'phone' = string, // 电话号
    // 注意上述两个字段只需要提供一个
    'method' = int,
    // 0 -- 使用电子邮箱登陆
    // 1 -- 使用电话号登陆
    'pwd' = 'string' // 密码，使用MD5进行哈希加密。
}
```

return values:
```json
{

    'status' = int,
    // status   解释           状态码
    // 0        成功           200
    // -1       查无此人        406
    // -2       密码错误        406
    // -100     其他错误        400
    'jwt' = string
    // 作为登陆凭证，48小时有效，请前端使用cookie等方式保存，大多数请求将使用jwt进行验证。
    // 登陆失败则返回空值
}
```

### 2. [请求] 注册 (未实现)
url: `/api/register`
method : `POST`
params:
```json
{
    "email" = string, // 电子邮箱地址
    "phone" = string, // 电话号码
    "type"  = int,  // 用户类型
    // 0 for 志愿者
    // 1 for 子女
    "pwd"   = string // 密码，使用MD5进行哈希加密
}
```

return values:
```json
    "status" : int  // 返回状态
    // 0 for 成功
    // -1 for 存在此用户
    // -100 其他原因
```

注：注册不包括登陆，请注册完毕后再调用登陆API。

### 3. 获取用户信息 (未实现)
url: `/api/user`
method: `GET`
params:
```json
{
    "email" : string,   // 电子邮箱 
    "phone" : string    // 电话号码
    // 上述两项只传一项即可
}
```