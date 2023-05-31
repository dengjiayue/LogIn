package funcmod

import "time"

// 登录列表方便查找,避免查找sql 提升性能.key为身份码
type SignInList map[string]*SignInData

// 用于 储存登录信息
type SignInData struct {
	ID             int   `json:"id"`
	ExpirationTime int64 `json:"ExpirationTime"`
}

// 身份验证信息(用于接收用户传入的身份码)
type IDCode struct {
	IDcode string `json:"idcode"`
}

// 邮箱验证码信息(验证码需要真人验证之后再发送)(登录时可以直接使用)
type EmailVerificationCode struct {
	Email        string `json:"email"`
	Verification int    `json:"verification"`
}

// 注册传入信息2,邮箱验证码信息+用户基本信息
type Logindata struct {
	EmailVerificationCode `json:"emailverificationcode"`
	Name                  string `json:"name"`
	Password              string `json:"password"`
}

// 注册传入信息1,id为图片的唯一标识符,code为用户传入 的图片验证码, Email为需要发送验证邮件的邮箱
type Login1data struct {
	PicID   string `json:"picid"`
	PicCode []byte `json:"piccode"`
	Email   string `json:"email"`
}

// 操作信息列表
type OperationRecord struct {
	ID       int    `json:"id"`
	Behavior string `json:"behavior"`
	Time     int    `json:"time"`
}

// 密码:改密接受信息
type ResetpaswData struct {
	IDCode
	Password    string `json:"password"`
	Newpassword string `json:"newpassword"`
}

// 改密接受信息1:图片验证+身份验证
type ResetpaswbyemailData1 struct {
	IDCode
	PicID   string `json:"picid"`
	PicCode []byte `json:"piccode"`
}

// 验证码验证+新密码
type ResetpaswbyemailData2 struct {
	IDCode
	Verification int    `json:"verification"`
	Newpassword  string `json:"newpassword"`
}

// 验证码数据模型
type VerificationCode struct {
	PicID     string    `json:"picid"`
	PicCode   []byte    `json:"piccode"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

// 密码登录信息(图片验证信息+邮箱+密码)
type Signupbypasw struct {
	PicID    string `json:"picid"`
	PicCode  []byte `json:"piccode"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 发送验证码时上传的信息(图片验证信息+身份验证信息)
type Signupbyemail1 struct {
	PicID   string `json:"picid"`
	PicCode []byte `json:"piccode"`
	Email   string `json:"email"`
}

// 使用邮箱验证码登录时上传的信息(图片验证信息+身份验证信息)
type Signupbyemail2 struct {
	EmailVerificationCode
}
