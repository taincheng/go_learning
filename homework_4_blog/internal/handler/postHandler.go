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

	claims := getClaims(c)
	var userID uint
	if claims != nil {
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
	c.JSON(http.StatusOK, gin.H{
		"message": "创建文章成功",
	})
}

// GetPostList 获取用户所有文章列表
func GetPostList(c *gin.Context) {
	claims := getClaims(c)
	var userID uint
	if claims != nil {
		userID = claims.UserID
	}
	postList, err := service.SelectPostList(userID)
	if err != nil {
		util.Logger.Error("GetPostList", zap.String("errorMessage", err.Error()))
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取文章列表成功",
		"data":    postList,
	})
}

// GetPostInfo 获取单个文章的信息
func GetPostInfo(c *gin.Context) {

	postTitle := c.Query("postTitle")
	if postTitle == "" {
		util.Logger.Error("GetPostInfo", zap.String("errorMessage", "文章名称不能为空"))
		c.JSON(http.StatusBadRequest, "文章名称不能为空")
		return
	}
	postInfo, err := service.SelectPostInfoByTitle(postTitle)
	if err != nil {
		util.Logger.Error("GetPostInfo", zap.String("errorMessage", err.Error()))
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取文章信息成功",
		"data":    postInfo,
	})
}

// UpdatePost 实现文章的更新功能，只有文章的作者才能更新自己的文章。
func UpdatePost(c *gin.Context) {
	claims := getClaims(c)
	var userID uint
	if claims != nil {
		userID = claims.UserID
	}
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		util.Logger.Error("UpdatePost", zap.String("errorMessage", "绑定参数失败"))
		c.JSON(http.StatusBadRequest, "绑定参数失败")
	}

	if post.UserID != userID {
		util.Logger.Error("UpdatePost", zap.String("errorMessage", "用户ID不匹配"))
		c.JSON(http.StatusNotAcceptable, "用户ID不匹配")
		return
	}

	if err := service.UpdatePost(&post); err != nil {
		util.Logger.Error("UpdatePost", zap.String("errorMessage", "更新文章失败"))
		c.JSON(http.StatusInternalServerError, "更新文章失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "更新文章成功",
	})
}

// DeletePost 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func DeletePost(c *gin.Context) {
	claims := getClaims(c)
	var userID uint
	if claims != nil {
		userID = claims.UserID
	}
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		util.Logger.Error("DeletePost", zap.String("errorMessage", "绑定参数失败"))
		c.JSON(http.StatusBadRequest, "绑定参数失败")
	}
	if post.UserID != userID {
		util.Logger.Error("DeletePost", zap.String("errorMessage", "用户ID不匹配"))
		c.JSON(http.StatusNotAcceptable, "用户ID不匹配")
		return
	}
	if post.ID == 0 {
		util.Logger.Error("DeletePost", zap.String("errorMessage", "文章ID不能为空"))
		c.JSON(http.StatusBadRequest, "文章ID不能为空")
		return
	}
	if err := service.DeletePost(&post); err != nil {
		util.Logger.Error("DeletePost", zap.String("errorMessage", "删除文章失败"))
		c.JSON(http.StatusInternalServerError, "删除文章失败")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除文章成功",
	})

}

func getClaims(c *gin.Context) *util.Claims {
	claims, exists := c.Get("claims")
	if !exists {
		util.Logger.Error("CreatePost", zap.String("errorMessage", "获取用户信息失败"))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取用户信息失败"})
		return nil
	}
	if claims, ok := claims.(*util.Claims); ok {
		return claims
	}
	return nil
}
