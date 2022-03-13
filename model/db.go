package model

import (
	"database/sql"
	"fmt"
	"ginweibo/utils"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var sqlDB *sql.DB

// InitDb 连接配置数据库
func InitDb() {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	sqlDB, err = sql.Open(`mysql`, dns)
	//错误处理
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数:", err)
		os.Exit(1)
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}

func DB() *sql.DB {
	return sqlDB
}
