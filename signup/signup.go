package signup

import (
	"database/sql"
	"fmt"
	"goweb/day7/gee7"
	"goweb/logindata/funcmod"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

func SigninBypassword(db *sql.DB, signinlest funcmod.SignInList) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {
		var data funcmod.Signupbypasw
		err := ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
		}
		// 验证图片验证码是否正确
		err = funcmod.Verifypiccode(db, data.PicID, data.PicCode)
		if err != nil {
			ctx.Returnfunc(401, "验证码错误"+err.Error(), nil)
			return
		}
		//验证成功,获取id信息
		id, err := funcmod.GetuserId(db, data.Email)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		// 验证登录密码是否正确
		err = funcmod.Verifypassword(db, id, data.Password)
		if err != nil {
			ctx.Returnfunc(401, "密码错误:"+err.Error(), nil)
			return
		}
		//密码验证成功:
		idcode := uuid.New().String() //随机生成唯一标识符,作为身份码
		signindata := funcmod.SignInData{
			ID:             id,
			ExpirationTime: time.Now().Unix() + 30*60,
		}
		//向登录列表中加入元素
		signinlest[idcode] = &signindata
		// 添加操作信息
		funcmod.Addoperation(db, signindata.ID, "signin成功")
		ctx.Returnfunc(200, "ok", gee7.H{
			"idcode": idcode,
		})
	}
}

// 登录第一步:1. 图片验证,;2. 发送邮件验证码
func SigninByEmail1(db *sql.DB, dialer *gomail.Dialer) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {
		data := funcmod.Signupbyemail1{}
		//读取用户传入的json数据
		err := ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(410, "解析数据错误"+err.Error(), nil)
			return
		}
		// fmt.Printf("=%#v\n", data)
		if funcmod.Validatedata(data.Email) {
			ctx.Returnfunc(402, "email不合理,请规范填写数据", nil)
		}
		//验证图片验证码是否正确
		err = funcmod.Verifypiccode(db, data.PicID, data.PicCode)
		if err != nil {
			ctx.Returnfunc(401, "验证码错误"+err.Error(), nil)
			return
		}
		_, err = funcmod.GetuserId(db, data.Email)
		if err != nil {
			ctx.Returnfunc(410, "账号未注册注册,请先注册", nil)
			return
		}
		//获取验证码函数
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
		err = funcmod.SendEmail(dialer, data.Email, verification)
		if err != nil {
			ctx.Returnfunc(401, "邮箱错误"+err.Error(), nil)
			return
		}
		ctx.Returnfunc(200, "ok", gee7.H{
			"msg": "ok",
		})
	}
}

// 完成验证码登录功能:验证验证码是否正确,插入数据
func SigninByEmail2(db *sql.DB, signinlest funcmod.SignInList) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {
		var data funcmod.Signupbyemail2
		err := ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(404, err.Error(), nil)
			return
		}
		fmt.Printf("=%#v\n", data)
		// 验证用户邮箱验证码是否正确
		err = funcmod.Verify(db, data.Email, data.Verification)
		if err != nil {
			ctx.Returnfunc(501, "验证码错误"+err.Error(), nil)
			return
		}

		//验证成功,获取id信息
		id, err := funcmod.GetuserId(db, data.Email)
		if err != nil {
			ctx.Returnfunc(403, "未发现账号"+err.Error(), nil)
			return
		}
		funcmod.Lognormal(fmt.Sprintf("id=%d", id))
		idcode := uuid.New().String() //随机生成唯一标识符,作为身份码
		signindata := funcmod.SignInData{
			ID:             id,
			ExpirationTime: time.Now().Unix() + 30*60,
		}
		// 插入用户登录信息,登陆成功
		signinlest[idcode] = &signindata
		//插入操作信息
		funcmod.Addoperation(db, signindata.ID, "signin成功")
		//登录成功返回id与身份码
		ctx.Returnfunc(200, "ok", gee7.H{
			"idcode": idcode,
		})

	}
}
