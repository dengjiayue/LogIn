package funcmod

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"fmt"
	"goweb/day7/gee7"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/dchest/captcha"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

// 允许跨域访问中间件
func Dfot() gee7.HandlerFunc {
	return func(ctx *gee7.Context) {
		// 处理预检请求
		if ctx.Req.Method == "OPTIONS" {
			// 验证预检请求的来源、头部字段和请求方法是否符合预期
			// ...

			// 设置响应头部字段
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// 返回状态码 200 和空响应体
			ctx.Writer.WriteHeader(http.StatusOK)
			return
		}
	}
}

// 临时.........身份验证中间件(是否处于登录状态)
// func Authentication(db *sql.DB) gee7.HandlerFunc {
// 	return func(ctx *gee7.Context) {

// 		var signInData SignInData
// 		err := ctx.Getjson(&signInData)
// 		if err != nil {
// 			ctx.Returnfunc(401, err.Error(),nil)
// 		}
// 		if err = SignInVerification(db, &signInData); err != nil {
// 			ctx.Returnfunc(401, "失败:"+err.E,nilrror())
// 		} else {
// 			ctx.String(200, "成功")
// 		}
// 	}
// }

// 统一输出日志
// 程序正常运行
func Lognormal(msg string) {
	log.Println("normal:", msg)
}

// 警告
func Logwarning(msg string) {
	log.Println("warning:", msg)
}

// 错误
func Logerror(msg string) {
	log.Panicln("error:", msg)
}

// 添加用户操作中间件
func Addoperationmid(db *sql.DB) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {
		ctx.Next()
		//没有id信息,直接退出无需插入操作信息
		if ctx.Userid == 0 {
			return
		}
		defer func() {
			if ctx.StatusCode >= 400 {
				Addoperation(db, ctx.Userid, "操作"+ctx.Path+"失败")
			} else {
				Addoperation(db, ctx.Userid, "操作"+ctx.Path+"成功")
			}
		}()
	}
}

// 发送图片验证码图片中间件(用户提交请求时)
func Picverfi(db *sql.DB) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {

		// 生成验证码
		id := uuid.New().String() //生成唯一标识符
		code := captcha.RandomDigits(4)
		img := captcha.NewImage(id, code, captcha.StdWidth, captcha.StdHeight)
		// 生成验证码
		var buf bytes.Buffer
		_, err := img.WriteTo(&buf)
		if err != nil {
			ctx.Returnfunc(401, "出错了,点击重试", nil)
		}

		// 将验证码和图片存储到数据库中
		encodedImage := base64.StdEncoding.EncodeToString(buf.Bytes())
		verificationCode := VerificationCode{PicCode: code, CreatedAt: time.Now()}
		_, err = db.Exec("INSERT INTO verification_codes (id, code, created_at) VALUES (?, ?, ?)", id, verificationCode.PicCode, verificationCode.CreatedAt)
		if err != nil {
			ctx.Returnfunc(http.StatusInternalServerError, "Internal Server Error:"+err.Error(), nil)
			return
		}

		// 返回图片与唯一标识符
		ctx.Returnfunc(200, "ok", gee7.H{
			"picid": id,
			"png":   encodedImage,
		})

		fmt.Printf("id=%#v\ncode=%d", id, code)
	}
}

// 用于检验用户传入的数据是否合理
// 正则表达式，匹配长度为 3-40 的 ASCII 字符串
// 合理为false不合理为true
func Validatedata(password string) bool {
	regex := regexp.MustCompile(`^[[:ascii:]]{3,40}$`)
	return !regex.MatchString(password)
}

// 密码改密
func ResetpaswBypassword(db *sql.DB, signinlest SignInList) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {
		data := ResetpaswData{}
		//读取用户传入的Returnfunc数据
		err := ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(410, "解析数据错误"+err.Error(), nil)
			return
		}
		if signinlest[data.IDcode] == nil {
			ctx.Returnfunc(420, "非法访问", nil)
			return
		}
		ctx.Userid = signinlest[data.IDcode].ID
		if Validatedata(data.Password) {
			ctx.Returnfunc(401, "password不合理,请规范填写数据", nil)
			return
		}
		if Validatedata(data.Newpassword) {
			ctx.Returnfunc(401, "newpassword不合理,请规范填写数据", nil)
			return
		}
		//身份验证
		err = SignInVerification(signinlest, data.IDcode)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		//验证密码
		err = Verifypassword(db, signinlest[data.IDcode].ID, data.Password)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		//验证成功:修改密码
		err = Changepassword(db, signinlest[data.IDcode].ID, data.Newpassword)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		ctx.Returnfunc(200, "ok", nil)
	}
}

// 验证改密第一步:1. 图片验证,;2. 发送邮件验证码
func ResetpaswByEmail1(db *sql.DB, dialer *gomail.Dialer, signinlest SignInList) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {
		data := ResetpaswbyemailData1{}
		//读取用户传入的Returnfunc数据
		err := ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(410, err.Error(), nil)
			return
		}
		if signinlest[data.IDcode] == nil {
			ctx.Returnfunc(420, "非法访问", nil)
			return
		}
		ctx.Userid = signinlest[data.IDcode].ID
		//验证图片验证码是否正确
		err = Verifypiccode(db, data.PicID, data.PicCode)
		if err != nil {
			ctx.Returnfunc(401, "图片验证码错误"+err.Error(), nil)
			return
		}
		// 验证成功:获取用户邮箱
		emial, err := GetuserEmail(db, signinlest[data.IDcode].ID)
		if err != nil {
			ctx.Returnfunc(410, "未找到邮箱信息,请联系管理员"+err.Error(), nil)
			return
		}
		verification := GetuVerification(db, emial)
		//储存验证码到数据库
		err = InsertVerification(db, &EmailVerificationCode{
			Email:        emial,
			Verification: verification})
		if err != nil {
			ctx.Returnfunc(401, "发送失败,稍后再试"+err.Error(), nil)
			return
		}
		//发送验证码
		err = SendEmail(dialer, emial, verification)
		if err != nil {
			ctx.Returnfunc(401, "邮箱错误"+err.Error(), nil)
			return
		}
		ctx.Returnfunc(200, "ok", nil)
	}
}

// 完成验证码改密功能:验证验证码是否正确,插入数据
func ResetpaswByEmail2(db *sql.DB, signinlest SignInList) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {
		var data ResetpaswbyemailData2
		err := ctx.Getjson(&data)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		if signinlest[data.IDcode] == nil {
			ctx.Returnfunc(420, "非法访问", nil)
			return
		}
		if Validatedata(data.Newpassword) {
			ctx.Returnfunc(402, "password不合理,请规范填写数据", nil)
			return
		}
		//获取邮箱
		email, err := GetuserEmail(db, signinlest[data.IDcode].ID)
		if err != nil {
			ctx.Returnfunc(401, "验证码错误"+err.Error(), nil)
			return
		}
		// 验证用户邮箱验证码是否正确
		err = Verify(db, email, data.Verification)
		if err != nil {
			ctx.Returnfunc(501, "验证码错误"+err.Error(), nil)
			return
		}
		//验证成功,获取id信息
		id, err := GetuserId(db, email)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		//修改密码
		err = Changepassword(db, id, data.Newpassword)
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		// 添加操作信息
		err = Addoperation(db, id, "改密成功")
		if err != nil {
			ctx.Returnfunc(401, err.Error(), nil)
			return
		}
		ctx.Returnfunc(200, "ok", nil)
	}
}
