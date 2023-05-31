package main

import (
	"fmt"
	"goweb/day7/gee7"
	"goweb/logindata/funcmod"
	"goweb/logindata/signup"
	"log"
	"time"
	// "github.com/rs/cors"
)

func main() {
	r := gee7.New()
	// r.Use(cors.Default())
	//设置统一日志输出(日期+时间+输出文件+行+信息)
	r.Use(funcmod.Dfot())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db, err := funcmod.Opensql()
	if err != nil {
		fmt.Printf("数据库出错=%#v\n", err)
		return
	}
	defer db.Close()

	//反向映射表
	signinlest2 := make(map[int]string)

	signinlest := funcmod.SignInList{}
	signinlest["ebbeb952-27d2-42cc-979c-bd99e517f0cd"] = &funcmod.SignInData{
		ID:             10007,
		ExpirationTime: time.Now().Unix() + 30*60,
	}
	r.Use(funcmod.DBMid(db))
	// funcmod.Changepassword(db, 10007, "abcd1234")
	//连接邮箱
	dailer := funcmod.CreatDialer()
	//获取图片验证码
	r.GET("/verifypic", funcmod.Picverfi(db))
	//注册1,1. 图片验证,;2. 发送邮件验证码
	r.POST("/signup1", func(ctx *gee7.Context) {
		data := funcmod.Login1data{}
		//读取用户传入的json数据
		err = ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(410, "解析数据错误"+err.Error(), nil)
			funcmod.Logwarning("解析数据错误" + err.Error())
			return
		}
		if funcmod.Validatedata(data.Email) {
			ctx.Returnfunc(402, "email不合理,请规范填写数据", nil)
			funcmod.Logwarning("email不合理")
			return
		}
		//验证图片验证码是否正确
		err = funcmod.Verifypiccode(db, data.PicID, data.PicCode)
		if err != nil {
			ctx.Returnfunc(401, "验证码错误"+err.Error(), nil)
			funcmod.Logwarning("验证码错误")
			return
		}
		_, err = funcmod.GetuserId(db, data.Email)
		if err == nil {
			ctx.Returnfunc(410, "账号已经注册了", nil)

			return
		}

		verification := funcmod.GetuVerification(db, data.Email)
		//储存验证码到数据库
		err = funcmod.InsertVerification(db, &funcmod.EmailVerificationCode{
			Email:        data.Email,
			Verification: verification})
		if err != nil {
			ctx.Returnfunc(401, "发送失败,稍后再试"+err.Error(), nil)
			return
		}
		//发送验证码
		err = funcmod.SendEmail(dailer, data.Email, verification)
		if err != nil {
			ctx.Returnfunc(401, "发送失败"+err.Error(), nil)
			return
		}

		ctx.Returnfunc(200, "ok", nil)
	})
	//login2:1.验证验证码,;2. 向用户列表插入数据
	r.POST("/signup2", func(ctx *gee7.Context) {
		data := funcmod.Logindata{}
		//读取用户传入的json数据
		err = ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(410, "解析数据错误"+err.Error(), nil)
			return
		}
		if funcmod.Validatedata(data.Name) {
			ctx.Returnfunc(402, "name不合理,请规范填写数据", nil)
			return
		}
		if funcmod.Validatedata(data.Password) {
			ctx.Returnfunc(402, "password不合理,请规范填写数据", nil)
			return
		}
		//验证用户验证码是否正确
		err = funcmod.Verify(db, data.Email, data.Verification)
		if err != nil {
			ctx.Returnfunc(501, "验证码错误"+err.Error(), nil)
			return
		}
		//插入用户数据
		id, err := funcmod.AddLogindata(db, &data)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
		} else {
			// 添加操作信息
			err = funcmod.Addoperation(db, id, "注册成功")
			if err != nil {
				ctx.Returnfunc(401, err.Error(), nil)
			} else {
				ctx.Returnfunc(200, "ok", nil)
			}
		}
	})
	r.POST("/loginpsw", signup.SigninBypassword(db, signinlest, signinlest2))
	r.POST("/loginemail1", signup.SigninByEmail1(db, dailer))
	r.POST("/loginemail2", signup.SigninByEmail2(db, signinlest, signinlest2))

	home := r.Group("/home")
	home.Use(funcmod.Addoperationmid(db))
	//修改密码
	//密码改密
	home.POST("/resetpaswbypasw", funcmod.ResetpaswBypassword(db, signinlest))
	//验证码改密
	home.POST("/resetpaswbyemail1", funcmod.ResetpaswByEmail1(db, dailer, signinlest))
	home.POST("/resetpaswbyemail2", funcmod.ResetpaswByEmail2(db, signinlest))
	// 登出
	r.POST("/signout", func(ctx *gee7.Context) {
		data := funcmod.IDCode{}
		//读取用户传入的json数据
		err = ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(410, "解析数据错误"+err.Error(), nil)
			return
		}
		if signinlest[data.IDcode] == nil {
			ctx.Returnfunc(420, "非法访问", nil)
			return
		}
		ctx.Userid = signinlest[data.IDcode].ID

		//身份验证

		if signinlest[data.IDcode] == nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		//删除登录列表信息
		delete(signinlest, data.IDcode)

		ctx.Returnfunc(200, "ok", nil)

	})
	// 获取操作信息
	r.POST("/getoperations", func(ctx *gee7.Context) {
		data := funcmod.IDCode{}
		//读取用户传入的json数据
		err = ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(410, "解析数据错误"+err.Error(), nil)
			return
		}
		if signinlest[data.IDcode] == nil {
			ctx.Returnfunc(420, "非法访问", nil)
			return
		}
		r, err := funcmod.Getoperations(db, signinlest[data.IDcode].ID)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		ctx.Returnfunc(200, "ok", gee7.H{
			"operations": r,
		})
	})

	r.Run(":8888")
}
