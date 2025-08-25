package handler

import (
	"homework_4_blog/internal/model"
	"homework_4_blog/internal/service"
	"homework_4_blog/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePost 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
func CreatePost(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		util.Logger.Error("CreatePost", zap.String("errorMessage", "获取用户信息失败"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return
	}
	var userID uint
	if claims, ok := claims.(*util.Claims); ok {
		userID = claims.UserID
	}

	util.Logger.Debug("CreatePost", zap.Uint("userID", userID))

	// 创建文章
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		util.Logger.Error("CreatePost", zap.String("errorMessage", "绑定参数失败"))
		c.JSON(http.StatusBadRequest, "绑定参数失败")
	}
	post.UserID = userID

	// 写入数据库
	err := service.CreatePost(&post)
	if err != nil {
		util.Logger.Error("CreatePost", zap.String("errorMessage", "创建文章失败"))
		c.JSON(http.StatusInternalServerError, "创建文章失败")
	}

	util.Logger.Info("CreatePost", zap.String("info", "创建文章成功"))
	c.JSON(http.StatusOK, "创建文章成功")
}

// GetPostList 获取用户所有文章列表
func GetPostList(c *gin.Context) {

}

// GetPostInfo 获取单个文章的信息
func GetPostInfo(c *gin.Context) {

}

// UpdatePost 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func UpdatePost(c *gin.Context) {

}

// DeletePost 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func DeletePost(c *gin.Context) {

}
