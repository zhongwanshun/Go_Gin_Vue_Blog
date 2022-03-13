package model

import (
	"context"
	"ginweibo/utils/errmsg"
	"github.com/sirupsen/logrus"
	"time"
)

type Article struct {
	//使用Gin的模型绑定和验证
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Title        string    `json:"title"`
	Desc         string    `json:"desc"`
	Content      string    `json:"content"`
	CommentCount int       `json:"comment_count"`
	ReadCount    int       `json:"read_count"`
}

func fetchArticle(ctx context.Context, query string, args ...interface{}) ([]Article, error) {
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

	result := make([]Article, 0)
	for rows.Next() {
		t := Article{}
		err = rows.Scan(&t.ID, &t.Title, &t.CreatedAt, &t.Desc, &t.Content, &t.CommentCount, &t.ReadCount)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

// CreateArt 新增微博
func CreateArt(ctx context.Context, data *Article) int {
	query := "INSERT article SET created_at = ?, title = ?, `desc` = ?, content = ?, comment_count = ?, read_count = ?"
	stmt, err := DB().PrepareContext(ctx, query)
	if err != nil {
		return errmsg.ERROR
	}
	res, err := stmt.ExecContext(ctx, time.Now(), data.Title, data.Desc, data.Content, data.CommentCount, data.ReadCount)

	if err != nil {
		return errmsg.ERROR // 500
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errmsg.ERROR
	}
	if affect != 1 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArtInfo 查询单个微博
func GetArtInfo(ctx context.Context, id int) (Article, int) {
	var art Article
	query1 := "SELECT created_at, title, `desc`, content, comment_count, read_count FROM article WHERE id = ?"
	query2 := "UPDATE article SET read_count = ? WHERE id = ?"
	tx, err := DB().Begin()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error("Unable to rollback!")
		}
		return art, errmsg.ERROR
	}
	err = tx.QueryRowContext(ctx, query1, id).Scan(&art.CreatedAt, &art.Title, &art.Desc, &art.Content, &art.CommentCount, &art.ReadCount)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error("Unable to rollback!")
		}
		return art, errmsg.ERROR
	}
	_, err = tx.ExecContext(ctx, query2, art.ReadCount+1, id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			logrus.Error("Unable to rollback!")
		}
		return art, errmsg.ERROR
	}
	if err = tx.Commit(); err != nil {
		logrus.Error("Error when committing")
		err = tx.Rollback()
		if err != nil {
			logrus.Error("Unable to rollback!")
		}
		return art, errmsg.ERROR
	}
	return art, errmsg.SUCCESS
}

// GetArt 查询微博列表
func GetArt(ctx context.Context, pageSize int, pageNum int) ([]Article, int, int64) {
	var total int64
	query := "SELECT id, title, created_at, `desc`, content, comment_count, read_count FROM article ORDER BY created_at DESC LIMIT ? OFFSET ?"
	list, err := fetchArticle(ctx, query, pageSize, (pageNum-1)*pageSize)

	if err != nil {
		return make([]Article, 0), errmsg.ERROR, 0
	}
	total = int64(len(list))
	return list, errmsg.SUCCESS, total

}

// SearchArticle 搜索微博标题
func SearchArticle(ctx context.Context, title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var total int64
	query := "SELECT id, title, created_at, `desc`, content, comment_count, read_count FROM article WHERE title LIKE ? ORDER BY created_at DESC LIMIT ? OFFSET ?"
	list, err := fetchArticle(ctx, query, title+"%", pageSize, (pageNum-1)*pageSize)
	if err != nil {
		return make([]Article, 0), errmsg.ERROR, 0
	}
	total = int64(len(list))
	return list, errmsg.SUCCESS, total
}

// EditArt 编辑微博
func EditArt(ctx context.Context, id int, data *Article) int {
	query := "UPDATE article SET title = ?, `desc` = ?, content = ? WHERE id = ?"
	_, err := DB().ExecContext(ctx, query, data.Title, data.Desc, data.Content, id)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteArt 删除微博
func DeleteArt(ctx context.Context, id int) int {
	query := "DELETE FROM article WHERE id = ?"
	_, err := DB().ExecContext(ctx, query, id)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
