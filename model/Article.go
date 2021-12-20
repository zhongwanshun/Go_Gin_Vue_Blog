package model

import "github.com/jinzhu/gorm"

//文章模型
type Article struct {
	//文章分类
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	//文章id
	Cid int `gorm:"type:int;not null" json:"cid"`
	//描述
	Desc string `gorm:"type:varchar(200)" json:"desc"`
	//目录
	Content string `gorm:"type:longtext" json:"content"`
	//文章图片
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}
