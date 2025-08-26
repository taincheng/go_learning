package handler

import (
	"homework_4_blog/internal/model"
	"homework_4_blog/internal/service"
	"homework_4_blog/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 实现评论的创建功能，已认证的用户可以对文章发表评论。
// 实现评论的读取功能，支持获取某篇文章的所有评论列表。

func CreateComment(c *gin.Context) {
	var userID uint
	claims := getClaims(c)
	if claims != nil {
		userID = claims.UserID
	}
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		util.Logger.Error("CreateComment", zap.String("errorMessage", err.Error()))
		c.JSON(http.StatusBadRequest, "绑定参数失败")
		return
	}
	comment.UserID = userID
	err := service.CreateComment(&comment)
	if err != nil {
		util.Logger.Error("CreateComment", zap.String("errorMessage", "创建评论失败"))
		c.JSON(http.StatusInternalServerError, "创建评论失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "创建评论成功",
	})
}

func GetCommentList(c *gin.Context) {
	postId, err := strconv.ParseUint(c.Query("postId"), 10, 64)
	if err != nil {
		util.Logger.Error("GetCommentList", zap.String("errorMessage", "参数类型错误"))
		c.JSON(http.StatusBadRequest, "参数类型错误")
		return
	}

	commentList, err := service.GetCommentList(uint(postId))
	if err != nil {
		util.Logger.Error("GetCommentList", zap.String("errorMessage", err.Error()))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取评论列表成功",
		"data":    commentList,
	})
}
