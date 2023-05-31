package funcmod

//对数据库操作的所有函数
import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

// 添加操作信息,向表OperationRecord(操作列表)添加
func Addoperation(db *sql.DB, id int, operation string) error {

	_, err := db.Exec("insert into OperationRecord (id,behavior,time) VALUES (?, ?, ?)", id, operation, time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

// 获取:表OperationRecord(操作列表)的数据
func Getoperations(db *sql.DB, id int) ([]OperationRecord, error) {

	rows, err := db.Query("SELECT behavior, time FROM OperationRecord WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := []OperationRecord{}
	for rows.Next() {
		p := &OperationRecord{ID: id}
		err = rows.Scan(&p.Behavior, &p.Time)
		if err != nil {
			return result, err
		}
		result = append(result, *p)
	}
	if err := rows.Err(); err != nil {
		return result, err
	}
	return result, nil
}

// user表------------------------------------------------------------------

// add:向user(注册列表)添加元素
func AddLogindata(db *sql.DB, logindata *Logindata) (int, error) {

	r, err := db.Exec("insert into user (name,password,Email) VALUES (?, ?, ?)", logindata.Name, encryptPassword(logindata.Password), logindata.Email)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	return int(id), err
}

// ((id是操作用的,不是给用户登录用的))
// get:获取:向user(注册列表)获取id,登陆成功,获取id
func GetuserId(db *sql.DB, email string) (int, error) {

	id := -1
	err := db.QueryRow("SELECT id FROM user WHERE Email = ?", email).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// get:获取:向user(注册列表)获取email(改密时查Email,不能由用户传入Email,)
func GetuserEmail(db *sql.DB, id int) (string, error) {

	Email := ""
	err := db.QueryRow("SELECT Email FROM user WHERE id = ?", id).Scan(&Email)
	if err != nil {
		return "", err
	}
	return Email, nil
}

// 查:验证user(注册列表)密码
func Verifypassword(db *sql.DB, id int, password string) error {

	count := 0
	//SELECT COUNT(*)这个函数一般不会报错,如果找不到值,那么count=0
	err := db.QueryRow("SELECT COUNT(*) FROM user WHERE id = ? AND password = ? ", id, encryptPassword(password)).Scan(&count)
	fmt.Printf("count=%#v\n", count)
	if err != nil || count == 0 {
		return fmt.Errorf("出错了:密码错误:%#v", err)
	}
	return nil
}

// 改:修改:user(注册列表)密码
func Changepassword(db *sql.DB, id int, newpassword string) error {

	_, err := db.Exec("UPDATE user SET password = ? WHERE id = ?", encryptPassword(newpassword), id)
	if err != nil {
		return err
	}
	return nil
}

// -SignInlist-----------删除其他操作,使用内存管理登录列表---------------------------------------------------------

// 查:SignIn查询身份验证功能:传入SignInlist
func SignInVerification(signinlest SignInList, idcode string) error {

	// 查询身份认证码是否正确

	if signinlest[idcode] == nil {
		return fmt.Errorf("非法访问")
	}

	// 验证身份认证码是否过期
	if time.Now().Unix() > signinlest[idcode].ExpirationTime {
		return fmt.Errorf("身份验证过期,请重新登陆")
	}

	// 验证成功.更新身份认证码的过期时间
	signinlest[idcode].ExpirationTime = time.Now().Unix() + 30*60
	return nil
}

// // 增:向登录列表插入信息
// func InsertSignInData(db *sql.DB, data *SignInData) error {

// 	// 准备SQL语句
// 	stmt, err := db.Prepare("REPLACE INTO SignInList (id, IDcode, ExpirationTime) VALUES (?, ?, ?)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// 执行插入操作
// 	result, err := stmt.Exec(data.ID, data.IDCode, time.Now().Unix()+30*60)
// 	if err != nil {
// 		return err
// 	}

// 	// 输出插入结果
// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // delete:删除登录信息; 登出删除登录列表元素
// func Deletesignindata(db *sql.DB, id int) error {
// 	_, err := db.Exec("DELETE FROM SignInList WHERE id = ?", id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// VerificationCode(验证码储存库)----------------------------------------------------------

// 获取验证码
// 检查是否进行验证过,用于重新发送验证码时调用,检查用户是否验证过
// 如果没有就发送与原来相同的验证码,否则就发送新验证码
func GetuVerification(db *sql.DB, email string) int {

	// 查询验证码是否正确
	var expirationTime int64 //储存过期时间

	verification, t := 0, -1

	err := db.QueryRow("SELECT time, Verification, ExpirationTime FROM VerificationCode WHERE Email = ?", email).Scan(&t, &verification, &expirationTime)
	//当没有验证码发送记录||已经验证过||超出时间限制  则需要重新生成验证码
	if err != nil || t == 1 || expirationTime+2500*60 < time.Now().Unix() {
		fmt.Printf("重新生成验证码%#v\n")
		rand.Seed(time.Now().UnixNano())
		// 生成100000到999999之间的随机数,作为验证码
		return rand.Intn(899999) + 100000
	} else {
		fmt.Printf("旧验证码=%#v\n", verification)
		return verification
	}
}

// 查: 验证验证码
func Verify(db *sql.DB, email string, verification int) error {

	// 查询验证码是否正确
	var expirationTime int64 //储存过期时间
	//代表已经验证过
	_, err := db.Exec("UPDATE VerificationCode SET time = ? WHERE Email = ?", 1, email)
	if err != nil {
		return fmt.Errorf("出错了" + err.Error())
	}
	err = db.QueryRow("SELECT ExpirationTime FROM VerificationCode WHERE Email = ? AND Verification = ?", email, verification).Scan(&expirationTime)
	if err != nil {
		return fmt.Errorf("验证码错误" + err.Error())
	}

	// 验证身份认证码是否过期
	if time.Now().Unix() > expirationTime {
		return fmt.Errorf("过期")
	}
	return nil
}

// 增:验证码列表插入信息
func InsertVerification(db *sql.DB, data *EmailVerificationCode) error {

	// 准备SQL语句

	stmt, err := db.Prepare("REPLACE INTO VerificationCode (Email, Verification, ExpirationTime, time) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行插入操作
	_, err = stmt.Exec(data.Email, data.Verification, time.Now().Unix()+5*60, 0)
	if err != nil {
		return err
	}
	return nil
}

// get:获取:验证次数,重发验证码0无,1有
func GetuVerificationtime(db *sql.DB, email string) (int, error) {
	t := -1
	err := db.QueryRow("SELECT time FROM user WHERE Email = ?", email).Scan(&t)
	if err != nil {
		return -1, err
	}
	return t, nil
}

// // ----------------------------------------------------------
// func Insertuserdata(db *sql.DB, data Logindata) error {

// 	// 准备SQL语句

// 	stmt, err := db.Prepare("INSERT INTO user (Email, name, password) VALUES (?, ?, ?)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

//		// 执行插入操作
//		_, err = stmt.Exec(data.Email, data.Name, encryptPassword(data.Password))
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//
// 图片验证码验证操作:用户传入图片id(唯一标识符发送图片信息时给的),+验证码
func Verifypiccode(db *sql.DB, id string, code []byte) error {
	var createdAt []uint8
	err := db.QueryRow("SELECT created_at FROM verification_codes WHERE id = ? AND code = ?", id, code).Scan(&createdAt)
	if err != nil {
		return err
	}
	created_ats := string(createdAt)
	created_at, err := time.Parse("2006-01-02 15:04:05", created_ats)
	if err != nil {
		return err
	}
	// 删除验证码
	db.Exec("DELETE FROM verification_codes WHERE id = ?", id)
	if time.Now().Unix() > created_at.Unix()+5*60 { //时效5分钟
		return fmt.Errorf("验证码过期")
	}
	return nil
}
