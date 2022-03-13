package model

import (
	"context"
	"database/sql"
	"fmt"
	"ginweibo/utils/errmsg"
	"github.com/sirupsen/logrus"
	"time"
)

type Comment struct {
	//使用Gin的模型绑定和验证
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserId    uint      `json:"user_id"`
	ArticleId uint      `json:"article_id"`
	Title     string    `json:"article_title"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Status    int8      `json:"status"`
}

func fetchComment(ctx context.Context, query string, args ...interface{}) ([]Comment, error) {
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

	result := make([]Comment, 0)
	for rows.Next() {
		t := Comment{}
		err = rows.Scan(&t.ID, &t.UserId, &t.ArticleId, &t.Title, &t.Username, &t.Content, &t.Status, &t.CreatedAt)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

// AddComment 新增评论
func AddComment(ctx context.Context, data *Comment) int {
	query := `INSERT comment SET created_at = ?, user_id = ?, article_id = ?, title = ?, username = ?, content = ?, status = ?`
	query1 := `SELECT comment_count FROM article WHERE id = ?`
	query2 := `UPDATE article SET comment_count = ? WHERE id = ?`
	tx, err := DB().Begin()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	res, err := tx.ExecContext(ctx, query, time.Now(), data.UserId, data.ArticleId, data.Title, data.Username, data.Content, data.Status)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	affect, err := res.RowsAffected()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	if affect != 1 {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	var commentCount int
	err = tx.QueryRowContext(ctx, query1, data.ArticleId).Scan(&commentCount)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	res, err = tx.ExecContext(ctx, query2, commentCount+1, data.ArticleId)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	affect, err = res.RowsAffected()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error(err.Error())
		}
		return errmsg.ERROR
	}
	if affect != 1 {
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

// GetComment 查询单个评论
func GetComment(ctx context.Context, id int) (Comment, int) {
	var res Comment
	query := `SELECT id, user_id, article_id, title, username, content, status, created_at FROM comment WHERE id = ?`
	list, err := fetchComment(ctx, query, id)
	if err != nil {
		return Comment{}, errmsg.ERROR
	}
	if len(list) > 0 {
		res = list[0]
	} else {
		return res, errmsg.ERROR
	}
	return res, errmsg.SUCCESS
}

// GetCommentList 后台所有获取评论列表
func GetCommentList(ctx context.Context, pageSize int, pageNum int) ([]Comment, int64, int) {

	var commentList []Comment
	var total int64
	query := `SELECT comment.id, user_id, article_id, article.title, username, comment.content, status, comment.created_at
				FROM comment
				LEFT JOIN article ON article.id = comment.article_id 
				ORDER BY created_at DESC LIMIT ? OFFSET ?`
	commentList, err := fetchComment(ctx, query, pageSize, (pageNum-1)*pageSize)
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	total = int64(len(commentList))
	return commentList, total, errmsg.SUCCESS
}

// GetCommentCount 获取评论数量
func GetCommentCount(ctx context.Context, id int) (int64, int) {
	var total int64
	query := `SELECT COUNT(*) FROM comment WHERE article_id = ? AND status = 1`
	rows, err := DB().QueryContext(ctx, query, id)
	if err != nil {
		return -1, errmsg.ERROR
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(rows)
	for rows.Next() {
		err = rows.Scan(&total)
		if err != nil {
			return -1, errmsg.ERROR
		}
	}
	return total, errmsg.SUCCESS
}

// GetCommentListFront 展示页面获取评论列表
func GetCommentListFront(ctx context.Context, id int, pageSize int, pageNum int) ([]Comment, int64, int) {
	var total int64
	query := `SELECT comment.id, user_id, article_id, article.title, username, comment.content, status, comment.created_at
				FROM comment
				LEFT JOIN article ON article.id = comment.article_id
				WHERE article_id = ? AND status = 1 
				ORDER BY created_at DESC LIMIT ? OFFSET ?`
	list, err := fetchComment(ctx, query, id, pageSize, (pageNum-1)*pageSize)
	if err != nil {
		return make([]Comment, 0), 0, errmsg.ERROR
	}
	total = int64(len(list))
	return list, total, errmsg.SUCCESS
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, id uint) int {
	query := `DELETE FROM comment WHERE id = ?`
	stmt, err := DB().PrepareContext(ctx, query)
	if err != nil {
		return errmsg.ERROR
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return errmsg.ERROR
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errmsg.ERROR
	}
	if rowsAffected != 1 {
		logrus.Error(fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAffected))
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}
