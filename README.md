# Community-E-Home
## 使用后端代码进行调试
(下述Shell指令为 Linux Shell)
0. 配置好 `Golang` 环境
1. 配置`mysql`服务器，*建议*安装mysql的可视化工具(或者使用vscode 并安装本项目推荐拓展)。
2. 建议使用 `vscode` 作为 IDE 本项目提供数个 `拓展建议`
3. 在你的某个目录下，通过shell执行： 
   `git clone git@github.com:FinleyGe/Community-E-Home.git`
4. `cd Community-E-Home/server/config`
5. 复制`config.example.yaml`为副本`config.yaml`
   `cp config.example.yaml config.yaml`
6. 使用文本编辑器编辑
7. `cd ..`
8. `go mod tidy`
9.  `go get`
10. `go run main.go`
11. 则服务器将运行于你所设定的端口上

## API 接口文档

说明：
1. 下方文档中，参数和返回值使用了**非**标准的json语法格式。（标准的json中没有注释）
2. 所有请求使用json格式，请注意Json格式和PostForm的区别。可以参考：https://www.jianshu.com/p/d0a4de05786b
3. 部分指令需要认证：在Header 中 Authorization中传入 jwt 以认证
   如果认证不成功则返回401错误以及
   ```
   {
       "status" : "unauthorized"
   }
   ```
   如果有其他错误则返回500。

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
```
{

    'status' : int,
    // status   解释           状态码
    // 0        成功           200
    // -1       查无此人        406
    // -2       密码错误        406
    // -100     其他错误        400
    'jwt' : string
    // 作为登陆凭证，48小时有效，请前端使用cookie等方式保存，大多数请求将使用jwt进行验证。
    // 登陆失败则返回空值
}
```

### 2. [请求] 注册
url: `/api/register`
method : `POST`
params:
```
{
    "email" : string, // 电子邮箱地址
    "phone" : string, // 电话号码
    "type"  : int,  // 用户类型
    // 0 for 志愿者
    // 1 for 子女
    "pwd"   : string // 密码，使用MD5进行哈希加密
}
```

return values:
```
{
    "status" : int  // 返回状态
    // 0 for 成功
    // -1 for 存在此用户
    // -100 其他原因
}
```

注：注册不包括登陆，请注册完毕后再调用登陆API。

### 3. 获取用户信息 为实现
url: `/api/user`
method: `GET`
params:
不需要传入参数

return values:
```
{
    "status"        : int, 
    // 返回值        意义
    // 0            success
    // -100         其他错误
    "name"          : string,
    "email"         : string,
    "phone"         : string,
    "avatar_url"    : string,
    "type"          : int,
    "profile"       : string,
}
```