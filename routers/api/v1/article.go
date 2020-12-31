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
	valid := validation.Validation{}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")
	}

	var tagID = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagID = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagID
		valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	var msg string

	if !valid.HasErrors() {
		articles := models.GetArticles(util.GetPage(c), setting.PageSize, maps)

		data["list"] = articles
		data["total"] = models.GetArticleTotal(maps)

		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors {
			msg += err.Message + ", "
		}
	}

	if msg == "" {
		msg = e.GetMsg(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// GetArticle 获取指定文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	var data interface{}

	code := e.INVALID_PARAMS
	msg := ""
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			msg += err.Message + ", "
		}
	}
	if msg == "" {
		msg = e.GetMsg(code)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	articleParams := make(map[string]interface{})
	c.BindJSON(&articleParams)

	valid := models.ValidationArticle(articleParams, "create")

	code := e.SUCCESS
	msg := e.GetMsg(code)

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		msg = ""
		for _, err := range valid.Errors {
			msg += err.Message + ", "
		}
	} else {
		models.AddArticle(articleParams)
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
		var articleParams map[string]interface{}

		c.ShouldBind(&articleParams)

		valid := models.ValidationArticle(articleParams, "update")

		if valid.HasErrors() {
			code = e.INVALID_PARAMS
			msg = ""
			for _, err := range valid.Errors {
				msg += err.Message + ", "
			}
		} else {
			models.EditArticle(id, articleParams)
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
