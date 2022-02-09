# Community-E-Home

## API 接口文档

说明：
1. 下方文档中，参数和返回值使用了**非**标准的json语法格式。（标准的json中没有注释）
### [请求] Login 登陆请求

url: `/login`
method: `POST`
params:
```json
{
    'e-mail' = string // 电子邮箱
    'phone' = string // 电话号
    // 注意上述两个字段只需要提供一个
    'method' = int
    // 0 -- 使用电子邮箱登陆
    // 1 -- 使用电话号登陆
    'pwd' = 'string' // 密码，使用MD5进行哈希加密。
}
```

return values:
```json
{

    'status' = int
    // 0 成功
    // -1 查无此人
    // -2 密码错误
    'jwt' = string
    // 作为登陆凭证，48小时有效，请前端使用cookie等方式保存，大多数请求将使用jwt进行验证。
}
```

### [请求] 注册
url: `/register`
method : `POST`
params:
```json
{
}
```
<!-- TODO: Unfinished Here -->