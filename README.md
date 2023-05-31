# login功能设计 

## 需求
实现一个登录功能

实现的功能

1. 注册
2. 登录,githubs授权登录
3. 重置密码
4. 查看操作记录(登录, 注册, 重置密码, 登出. 都算操作)
5. 登出





### 后端接口设计

#### 1. 人机验证
**只要下面出现 人机验证 的功能都需要使用到该接口**

路径:http://localhost:8888/verifypic

方法get

功能:获取图片验证码

请求参数: 无

响应数据(示例)
```
{
    "code": 200,
    "msg": "ok",
    "data": {
        "picid": "52fa7b86-cf25-4822-a12c-d48e5b4d7f31",
        "png": "iVBORw0KGgoAAAANSUhEUgAAAPAAAABQCAMAAAAQlwhOAAAAP1BMVEUAAABNKGM1EEu9mNOUb6pvSoVFIFtZNG96VZA2EUy5lM93Uo2ifbhEH1pFIFt/WpVkP3qrhsFkP3pKJWBnQn3MHHtHAAAAAXRSTlMAQObYZgAABDFJREFUeJzsWuuSsyAMTdZpd6b2Mjv7/u/6zacFchGINkXdevaPi1xyOCEELBw4cGCH+F7bgMb4/v44xmsb8Ll4PB6q7GsVS9rg8dCMv77+MOMnYYx/8AkKI4GsECbhbwABR75MYPIWxkmwce667k12uiHQmRI4FAxvLZS7bvOMkWooFA5l49sJX1fYMOFgfOKqawSWgXMdmyX8n8L4kI3LyJdubFHGNvlSveKj2nl1/NojYgAKAifdOF+rE28ZmKJuJJP31N3zDVtLoi0UlvX3TTg4sllhW6ZhTkjagrB9ShzMpNELZMahSmW3xs25NaKooFQNvp223NhoqIyYTyvRnH41hs4M5aaEOvsIkzQxFaTjLStM7GL/osilIZXTqD7Zra/Cfd+79DN9IiDPIpVmjTB7WBrlN6ZfFvS9L2NWQNawOizRRtwNZKe+W5cnYa4OglZ4fO66LlQWbiH6QNyywjL88kSaEBlOOnQ7yiwD8q9nyOr7/selI6VwLrKOhInCopHu0lNhAPj5cWIs/mVkTumFUli0Ar2knTclH74KjO+JM4aMwpzkAoWvXtYvALPxJN/lUkpKkmZqtjGv1/UYl1XJKQzIFX6W7UBhtS2rt4MaE0ELdd61xcRSoZz+DnyvVxDbMKUb82zvKP0mlK0MhNNGjBpxOvagMBatRAiEAymhLLuw34PC9SWMIaIOCnNVISaUiFjuajMoWTl1YoqeDfGEDJDk3T7jgi5CM/ovcWExI43MXo4CX3FiILsQkutsovsqCs8cUG6kgMlhhb4kXkmFeXEbKBczDsxFQfGllF7zCbLAvTvNQQvoTdE61bwiD7h8xbKTAhdfXpO8gyIzmm8WQD4i1MeWjLnnhtIo5/SdFv9q/F7G7DoRVSZQHVspPLGspaNmPJcM/yqrrLXyJjlukskRK6NPzYq8EZn8Oq6L9FrwBqbkZ/p1nXHdOvNdM1/IpiazYJrQuovVrLMbLxaVrZH19wFm37nVPoGUO5klFln81nb2X0QYGd9ut2hJyfGLoxhNCi3Cg6n+bL61ijfaIhdqyqMYTXq2AP5hztqu0qsp/k62Mo0VLnUWHOhR5m3mZudsf0uvFlJMqVQMlzoL8kTeaE7AO5/Pyjhk1w5zbYkmVB11ILxEXsXRvCT+KxwNo1kbvJqgp17znYyEl06pPIzMyOkRdc7ocE9WOV4gwsD3hf5Zb2avpmJwLDZFWRLW9NSVs8soQC+CDHVBHddcUzbiQuyOzvlUi3ODbFq8nmYA926Ml+liWn8dhpm9Et95wEQZJvhgv78OjGfmqA0Q3VrDhe/MxMsN9/s986aBLSswvt/zjFtgdvR6FUXCpwYG+G80ZZQIsx87vA0rMM6+a8HXP6/ZBfbwyc0ZW/yV8YEDBzaKy+WytglNcbl8GOOD8J/Hp/HdJP4FAAD//55dD54XIp4vAAAAAElFTkSuQmCC"
    }
}
```

png为base64编码数据,解码可以得到验证码图片,需要展示给用户

picid为图片的唯一标识符,验证时需要带上才能正常验证,否则将报错


#### 2. 注册1:发送注册验证码:

路径: http://localhost:8888/loginemail1

方法post

功能:**人机验证**+邮箱验证码发送

请求数据(示例) :数据格式json
```
{
{
  "picid": "c19c20d2-ebea-487a-9c7d-f9c98705728f",
  "email": "...@gmail.com",
  "piccode": [
    2,
    3,
    9,
    9
  ]
}
```
picid为人机验证执行http://localhost:8888/verifypic时返回储存的,

piccode为图片中展示的验证码,需要用户输入

email是用户输入的注册邮箱

响应数据:数据格式json
```
{
code:状态码判断响应状态
msg:响应提示信息,一般为出错的时候展示给用户
data:nil(data为interface,根据需要进行返回,这里不需要返回值,所以为空nil)
}
```

![](https://djy1-1306563712.cos.ap-shanghai.myqcloud.com/20230414200825.png)

#### 3. 注册2:验证+信息填写
路径: http://localhost:8888/signup2

方法post

请求参数示例:
```
{
    "name": "小明",
    "password": "abcd123",
    "emailverificationcode": {
        "email": "...@gmail.com",
        "verification": 987639
    }
}
```

name,password,verification为用户输入,Email为注册1时填写的数据

返回:数据格式json
```
{
  "code": 200,
  "msg": "ok",
  "data": nil
}

```
提示用户成败信息

#### 4. 邮箱密码登录
路径: http://localhost:8888/loginpsw

方法:post

功能:人机验证+密码登录

请求参数:数据格式json
```
{
    "picid": "3a00e077-0e2f-47c0-9123-04a827235a6e",
    "piccode": [
       7,1,9,3
    ],
    "email": "...@gmail.com",
    "password": "abcd123"
}
```
piccode,Email,password为用户输入,picid为人机验证功能执行方法



返回:数据格式json
```

{
  "code": 200,
  "msg": "ok",
  "data": {
    idcode:身份码(string)
  }
}
```
这里如果成功data中会包含一条idcode信息,需要储存cookie,后面home下功能需要用到

![](https://djy1-1306563712.cos.ap-shanghai.myqcloud.com/20230414201123.png)

#### 5. 邮箱验证码登录1:发送邮箱验证码
路径: http://localhost:8888/loginemail1

方法post

功能:人机验证+邮箱验证码发送

请求参数:数据格式json
```
{
    "picid": "a3e6bcf5-e23b-4406-bd1a-14edf55d82b4",
    "piccode": [
        1,7,9,9
    ],
    "email": "...@gmail.com"
}
```

响应数据:数据格式json
```
{
  "code": 200,
  "msg": "ok",
  "data":nil
}

```
需要处理响应结果

#### 6. 验证码登录2
路径:http://localhost:8888/loginemail2

方法post

功能:验证验证码

请求参数:数据格式json
```

{
    "email": "...@gmail.com",
    "verification": 715562
}

```
Email为邮箱验证登录1时填写的数据

verification为用户输入

返回数据:数据格式json
```
{
    "code": 200,
    "msg": "ok",
    "data": {
        "idcode": "384a0457-7857-4da4-88b3-72c080634cf5"
    }
}
```
idcode存cookie,便于后期使用




![](https://djy1-1306563712.cos.ap-shanghai.myqcloud.com/20230414200940.png)

#### 7. 密码改密
http://localhost:8888/home/resetpaswbypasw

方法post

功能:身份验证+密码验证+改密

参数:数据格式json
```
{
    "password": "abcd123",
    "newpassword": "abcd1234",
    "idcode": "aed48dfc-868d-46c7-a7ed-23ec3fb09257"
}
```
idcode从cookie中获取,其他为用户填写,需要确保信息的合理性

返回:数据格式json
```
{
    "code": 200,
    "msg": "ok",
    "data": nil
}
```
处理成败信息

![](https://djy1-1306563712.cos.ap-shanghai.myqcloud.com/20230414201938.png)

#### 8. 查询操作信息
http://localhost:8888/getoperations

方法post

功能:身份验证+查询操作信息

请求参数:数据格式json
```
{
idcode:身份码
}
```
从cookie中获取

返回数据:数据格式json
```
{
code:200
msg:"ok"
{
  "operations": [
    {
      "id": 10004,
      "behavior": "注册成功",
      "time": 1682166659
    },
    {
      "id": 10004,
      "behavior": "signin成功",
      "time": 1682167322
    },
    {
      "id": 10004,
      "behavior": "登出成功",
      "time": 1682167868
    },
    {
      "id": 10004,
      "behavior": "signin成功",
      "time": 1682169145
    },
    {
      "id": 10004,
      "behavior": "signin成功",
      "time": 1682170421
    },
    {
      "id": 10004,
      "behavior": "signin成功",
      "time": 1682170825
    },
    {
      "id": 10004,
      "behavior": "登出成功",
      "time": 1682171024
    },
    {
      "id": 10004,
      "behavior": "signin成功",
      "time": 1682171528
    },
    {
      "id": 10004,
      "behavior": "改密成功",
      "time": 1682171658
    }
  ],
  "msg": "ok"
}
```
成功:展示"operations"中的 信息,失败返回失败原因

![](https://djy1-1306563712.cos.ap-shanghai.myqcloud.com/20230414202210.png)

#### 9. 验证码改密1
路径http://localhost:8888/home/resetpaswbyemail1

方法post



功能:人机验证+身份验证+邮件验证码发送

参数:数据格式json
```
{
    "picid": "52fa7b86-cf25-4822-a12c-d48e5b4d7f31",
    "piccode": [
        7,5,6,4
    ],
    "idcode": "ebbeb952-27d2-42cc-979c-bd99e517f0cd"
}
```
piccode为人机验证功能时返回图片中的验证码,由用户输入

返回:数据格式json
```
{
    "code": 200,
    "msg": "ok",
    "data": null
}
```
处理成败信息

#### 10. 登出
http://localhost:8888/logout

方法post

功能:身份验证+查询操作信息

请求参数:数据格式json
```
{
idcode:身份码
}
```
从cookie中获取

返回数据:数据格式json
```
{
code:200
msg:"ok"
data:nil
```
处理成败信息

#### 9. 验证码改密2
路径http://localhost:8888/home/resetpaswbyemail2

方法post

功能:身份验证+邮箱验证码验证+修改密码




参数:数据格式json
```
{
    "newpassword": "abcdef1234ad",
    "idcode": "ebbeb952-27d2-42cc-979c-bd99e517f0cd",
    "verification": 123455
}
```


返回(错误示例):数据格式json
```
{
    "code": 501,
    "msg": "验证码错误验证码错误sql: no rows in result set",
    "data": null
}
```

上面示例为验证码错误产生的返回

![](https://djy1-1306563712.cos.ap-shanghai.myqcloud.com/20230414201800.png)

![](https://djy1-1306563712.cos.ap-shanghai.myqcloud.com/20230414201321.png)

## 前端实现方案


#### 注册1 signup1(注册)

3个按钮:

1 获取图片验证:访问图片验证码接口,展示base64验证图片并储存picid

2 提交:访问注册1接口,提交数据

3 登录:跳转页面到LoginPsw(登录页面)

#### 注册2 signup2
3个按钮:

1 重新发送验证码: 退回到signup1页面

2 提交:访问注册2接口

3 登录:跳转页面到LoginPsw(登录页面)

#### 密码登录loginpsw

4个按钮:

1 获取验证码图片:访问获取验证码图片接口,展示返回的图片并储存picid

2 忘记密码: 跳转到到loginemail1页面

3 提交:访问密码登录接口接口,并处理响应结果

如果成功储存返回的idcode(cookie)并跳转页面到home页面;失败展示失败信息

#### 验证码登录 

1. loginemail1


4个按钮:

1 获取图片验证:访问图片验证码接口,展示base64验证图片并储存picid

2 提交:访问验证码登录1接口

3 注册:跳转页面到sianup1(注册页面)

4 使用密码登录:跳转到loginpsw(密码登录页面)

2. loginemail2

4个按钮:

1 重新发送验证码: 退回到loginemail1页面

2 提交:访问验证码登录2接口

3 注册:跳转页面到signup1(注册页面)

4 使用密码登录:跳转到loginpsw(密码登录页面)

### home页面
/home
* 定义三个按钮 sign out(登出),reset password(重置密码),view(查看操作记录)

 sign out(登出):访问登出接口,提交json数据,退回到loginpsw(登录页面)
 
 查询:访问查询功能接口,并处理响应数据,如果成功展示操作信息,如果失败,展示失败信息
 
 ![](https://djy1-1306563712.cos.ap-shanghai.myqcloud.com/20230521232627.png)

 #### repasswordpsw 旧密码密码改密

三个按钮:

1  忘记密码:跳转到repasswordemail页面(验证码改密)

2 返回:退回到home页面

3 改密:访问密码改密接口

#### repasswordemail 验证码改密1

3个按钮:

1 获取图片验证:访问图片验证码接口,展示base64验证图片并储存picid

2 提交:访问验证码改密1接口

3 使用密码改密:跳转页面到repasswordPsw(密码改密页面)

4 返回:退回到home页面

2. 验证码改密2 repasswordemail2

4个按钮:

1 重新发送验证码: 退回到repasswordemail1页面

2 提交:访问验证码验证码改密2接口

3 返回home:跳转页面到home(注册页面)

4 使用密码改密:跳转到repasswordpsw(密码改密页面)


注意进行错误处理,详情见代码# LogIn
