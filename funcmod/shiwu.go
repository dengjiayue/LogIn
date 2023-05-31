package funcmod

import (
	"database/sql"
	"goweb/day7/gee7"
)

// 事务中间件,如果用户响应出错,则数据库回滚到响应之前的状态
func DBMid(db *sql.DB) gee7.HandlerFunc {
	return func(ctx *gee7.Context) {

		// 开始事务
		tx, err := db.Begin()
		if err != nil {
			ctx.FailJson(401, err.Error())
			return
		}

		// 调用下一个处理器
		ctx.Next()

		defer func() {
			// 如果响应出现错误，回滚事务,500以内的错误回滚
			if ctx.StatusCode >= 400 && ctx.StatusCode <= 500 {
				tx.Rollback()
				return
			}
			// 提交事务
			tx.Commit()
		}()
	}
}
