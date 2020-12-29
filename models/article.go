package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article 文章
type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate 创建前回调
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// BeforeUpdate 更新前回调
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// GetArticles 获取文章列表
func GetArticles(pageNum int, pageSize int, params interface{}) (articles []Article) {
	db.Preload("Tag").Where(params).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

// GetArticle 获得文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

// GetArticleTotal 获得文章总数量
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

// AddArticle 添加文章
func AddArticle(article *Article) bool {
	db.Create(article)

	return true
}

// ExistArticleByID 根据ID判断文章是否存在
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}

// EditArticle 更新文章
func EditArticle(id int, params map[string]interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(params)

	return true
}

// DeleteArticle 删除文章
func DeleteArticle(id int) {
	db.Where("id = ?", id).Delete(&Article{})
}
