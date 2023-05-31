package funcmod

import (
	"crypto/sha256"
	"encoding/hex"
)

//密码哈希加密算法

func encryptPassword(password string) string {
	// 将字符串类型的密码转换为字节类型
	passwordBytes := []byte(password)
	// 创建SHA-256哈希对象
	sha256Hash := sha256.New()
	// 将密码字节数据传入哈希对象
	sha256Hash.Write(passwordBytes)
	// 获取哈希值的字节数据
	hashBytes := sha256Hash.Sum(nil)
	// 将字节数据转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)
	// 返回十六进制字符串类型的哈希值
	return hashString
}
