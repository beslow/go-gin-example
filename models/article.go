package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
	"github.com/unknwon/com"
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
func GetArticleTotal(params interface{}) (count int) {
	db.Model(&Article{}).Where(params).Count(&count)

	return
}

// AddArticle 添加文章
func AddArticle(article map[string]interface{}) bool {
	_tagID := 0
	if tagID, ok := article["tag_id"]; ok {
		switch tagID.(type) {
		case int:
			_tagID = tagID.(int)
		case string:
			_tagID = com.StrTo(tagID.(string)).MustInt()
		case float64:
			_tagID = int(tagID.(float64))
		}
	}
	_state := -1
	if state, ok := article["tag_id"]; ok {
		switch state.(type) {
		case int:
			_state = state.(int)
		case string:
			_state = com.StrTo(state.(string)).MustInt()
		case float64:
			_state = int(state.(float64))
		}
	}
	db.Create(&Article{
		TagID:     _tagID,
		Title:     article["title"].(string),
		Desc:      article["desc"].(string),
		Content:   article["content"].(string),
		CreatedBy: article["created_by"].(string),
		State:     _state,
	})

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

// ValidationArticle 校验文章
func ValidationArticle(params map[string]interface{}, action string) validation.Validation {
	valid := validation.Validation{}

	if tagID, ok := params["tag_id"]; ok {
		_tagID := -1
		switch tagID.(type) {
		case string:
			_tagID = com.StrTo(tagID.(string)).MustInt()
		case int:
			_tagID = tagID.(int)
		case float64:
			_tagID = int(tagID.(float64))
		}
		valid.Min(_tagID, 1, "tag_id").Message("标签ID必须大于0")

		if !ExistTagByID(_tagID) {
			valid.SetError("tag_id", "关联标签必须存在")
		}
	}

	if title, ok := params["title"]; ok || action == "create" {
		valid.Required(title, "title").Message("文章标题不能为空")
		valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	}

	if desc, ok := params["desc"]; ok || action == "create" {
		valid.Required(desc, "desc").Message("文章简述不能为空")
		valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	}

	if content, ok := params["content"]; ok || action == "create" {
		valid.Required(content, "content").Message("文章内容不能为空")
		valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	}

	if state, ok := params["state"]; ok || action == "create" {
		fmt.Println(reflect.TypeOf(state))
		_state := -1
		switch state.(type) {
		case string:
			_state = com.StrTo(state.(string)).MustInt()
		case int:
			_state = state.(int)
		case float64:
			_state = int(state.(float64))
		}
		fmt.Println(_state)
		valid.Range(_state, 0, 1, "state").Message("状态只能是0或1")
	}

	if modifiedBy, ok := params["modified_by"]; ok {
		valid.Required(modifiedBy, "modified_by").Message("文章修改人不能为空")
	}

	return valid
}
