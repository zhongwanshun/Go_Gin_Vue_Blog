package model

import (
	"context"
	"database/sql"
	"fmt"
	"ginweibo/utils/errmsg"
	"github.com/sirupsen/logrus"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//数据库的具体实现

//声明结构体
type ZwsUser struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	//使用Gin的模型绑定和验证
	Username string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `json:"role" validate:"required,gte=2" label:"角色码"`
}

func fetchUser(ctx context.Context, query string, args ...interface{}) ([]ZwsUser, error) {
	rows, err := DB().QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result := make([]ZwsUser, 0)
	for rows.Next() {
		t := ZwsUser{}
		err = rows.Scan(&t.ID, &t.CreatedAt, &t.Username, &t.Role)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

// CheckUser 查询用户是否存在
func CheckUser(ctx context.Context, name string) (code int) {
	//引入User
	var user ZwsUser
	query := `SELECT id FROM user WHERE username = ?`
	err := DB().QueryRowContext(ctx, query, name).Scan(&user.ID)
	if err != sql.ErrNoRows {
		return errmsg.ErrorUsernameUsed //1001
	}
	return errmsg.SUCCESS
}

// CreateUser 新增用户
func CreateUser(ctx context.Context, data *ZwsUser) int {
	data.Password = ScryptPw(data.Password)
	query := `INSERT user SET created_at = ?, username = ?, password = ?, role = ?`
	stmt, err := DB().PrepareContext(ctx, query)
	if err != nil {
		return errmsg.ERROR // 500
	}
	_, err = stmt.ExecContext(ctx, time.Now(), data.Username, data.Password, data.Role)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUser 查询用户
func GetUser(ctx context.Context, id int) (ZwsUser, int) {
	var user ZwsUser
	query := `SELECT id, created_at, username, role FROM user WHERE id = ?`
	err := DB().QueryRowContext(ctx, query, id).Scan(&user.ID, &user.CreatedAt, &user.Username, &user.Role)
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 查询用户列表(用户名,pagesize通过query传入,当前页数，返回一个用户切片)
func GetUsers(ctx context.Context, username string, pageSize int, pageNum int) ([]ZwsUser, int64) {
	var total int64
	query := `SELECT id, created_at, username, role FROM user WHERE username LIKE ? LIMIT ? OFFSET ?`
	// 数据库查找
	if username != "" {
		list, err := fetchUser(ctx, query, username+"%", pageSize, (pageNum-1)*pageSize)
		if err != nil {
			return nil, total
		}
		total = int64(len(list))
		return list, total
	} else {
		return nil, total
	}
}

// EditUser 编辑用户信息(一个id ,一个字典)
func EditUser(ctx context.Context, id int, data *ZwsUser) int {
	query1 := `SELECT id, username FROM user WHERE username = ?`
	query2 := `UPDATE user SET username = ?, role = ? WHERE id = ?`

	tx, err := DB().Begin()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	var ids uint
	err = tx.QueryRowContext(ctx, query1, data.Username).Scan(&ids)
	if err == nil {
		if ids != data.ID {
			err = tx.Rollback()
			if err != nil {
				logrus.Error(err.Error())
			}
			return errmsg.ErrorUsernameUsed
		}
	}
	_, err = tx.ExecContext(ctx, query2, data.Username, data.Role, id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
	}
	return errmsg.SUCCESS
}

// ChangePassword 修改密码
func ChangePassword(ctx context.Context, id int, data *ZwsUser) int {
	data.Password = ScryptPw(data.Password)
	query := `UPDATE user SET password = ? WHERE id = ?`
	stmt, err := DB().PrepareContext(ctx, query)
	if err != nil {
		return errmsg.ERROR
	}
	res, err := stmt.ExecContext(ctx, data.Password, id)
	if err != nil {
		return errmsg.ERROR
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errmsg.ERROR
	}
	if affect != 1 {
		logrus.Error(fmt.Errorf("Weird  Behavior. Total Affected: %d", affect))
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(ctx context.Context, id int) int {
	query := `DELETE FROM user WHERE id = ?`
	stmt, err := DB().PrepareContext(ctx, query)
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return errmsg.ERROR
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errmsg.ERROR
	}
	if affect != 1 {
		logrus.Error(fmt.Errorf("Weird  Behavior. Total Affected: %d", affect))
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ScryptPw 生成密码(将密码传入)
func ScryptPw(password string) string {
	const cost = 10 //数值越大,破解难度越大
	//利用hash加密(passwd的byte,将lenkey传入)
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	//返回结果hash,只能生成,不能解析
	return string(HashPw)
}

// CheckLogin 登录验证
func CheckLogin(ctx context.Context, username string, password string) (ZwsUser, int) {
	var user ZwsUser
	var PasswordErr error
	query := `SELECT id, created_at, username, role, password FROM user WHERE username = ?`
	err := DB().QueryRowContext(ctx, query, username).Scan(&user.ID, &user.CreatedAt, &user.Username, &user.Role, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errmsg.ErrorUserNotExist
		}
		return user, errmsg.ERROR
	}
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if PasswordErr != nil {
		return user, errmsg.ErrorPasswordWrong
	}
	return user, errmsg.SUCCESS
}
