package api

import (
	"blogSystem/service"
	"blogSystem/utils"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateBlog(c *gin.Context) {
	var createBlog service.BlogService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createBlog); err == nil {
		res := createBlog.CreateBlogService(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func GetBlog(c *gin.Context) {
	var getBlogService service.BlogService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getBlogService); err == nil {
		res := getBlogService.GetBlogService(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func GetBlogs(c *gin.Context) {
	var getBlogService service.BlogService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getBlogService); err == nil {
		res := getBlogService.GetBlogsService(claim.Id, c.Param("uid"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func UpdateBlog(c *gin.Context) {
	var updateBlog service.BlogService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateBlog); err == nil {
		res := updateBlog.UpdateBlogService(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func SearchBlogs(c *gin.Context) {
	var searchBlogService service.SearchBlogService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchBlogService); err == nil {
		res := searchBlogService.SearchBlogsService(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteBlog(c *gin.Context) {
	var deleteBlog service.BlogService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteBlog); err == nil {
		res := deleteBlog.DeleteBlogService(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

func DeleteBlogs(c *gin.Context) {
	var deleteBlogsService service.BlogService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteBlogsService); err == nil {
		res := deleteBlogsService.DeleteBlogsService(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}
