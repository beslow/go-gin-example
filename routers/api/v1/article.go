package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/beslow/go-gin-example/models"
	"github.com/beslow/go-gin-example/pkg/e"
	"github.com/beslow/go-gin-example/pkg/setting"
	"github.com/beslow/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetArticles 获取文章列表
func GetArticles(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	var state = 0
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	maps["state"] = state

	articles := models.GetArticles(util.GetPage(c), setting.PageSize, maps)

	data["list"] = articles
	data["total"] = models.GetArticleTotal(articles)

	code := e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle 添加文章
func AddArticle(c *gin.Context) {

	var articleParams models.Article
	c.BindJSON(&articleParams)

	tagID := articleParams.TagID
	title := articleParams.Title
	content := articleParams.Content

	valid := validation.Validation{}
	valid.Required(title, "title").Message("文章标题不能为空")
	valid.Required(content, "content").Message("文章内容不能为空")

	if tagID != 0 && !models.ExistTagByID(tagID) {
		valid.SetError("TagID", "关联的标签不存在1")
	}

	code := e.SUCCESS
	msg := e.GetMsg(code)

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		msg = ""
		for _, err := range valid.Errors {
			msg += err.Message + ", "
		}
	} else {
		models.AddArticle(tagID, title, articleParams.Desc, content, articleParams.State)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]string),
	})
}

// EditArticle 更新文章
func EditArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	msg := e.GetMsg(code)
	if models.ExistArticleByID(id) {
		maps := make(map[string]interface{})
		var articleParams models.Article
		valid := validation.Validation{}
		c.ShouldBind(&articleParams)
		tagID := articleParams.TagID
		if tagID != 0 {
			if tagID != 0 && !models.ExistTagByID(tagID) {
				valid.SetError("TagID", "关联的标签不存在")
			} else {
				maps["TagID"] = tagID
			}
		}

		title := articleParams.Title
		if title != "" {
			maps["Title"] = title
		}

		desc := articleParams.Desc
		if desc != "" {
			maps["Desc"] = desc
		}

		content := articleParams.Content
		if content != "" {
			maps["Content"] = content
		}

		maps["State"] = articleParams.State

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
			msg = ""
			for _, err := range valid.Errors {
				msg += err.Message + ", "
			}
		} else {
			models.EditArticle(id, maps)
		}

	} else {
		code = e.INVALID_PARAMS
		msg = "文章不存在"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]string),
	})

}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS
	msg := e.GetMsg(code)
	if models.ExistArticleByID(id) {
		models.DeleteArticle(id)
	} else {
		code = e.INVALID_PARAMS
		msg = "文章不存在"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]string),
	})
}
