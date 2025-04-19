package service

import (
	"blogSystem/models"
	"blogSystem/serializer"
)

// BlogService 博客增删改查接口
type BlogService struct {
	BlogName    string `json:"blog_name"`
	BlogContent string `json:"blog_content"`
	User        models.User
}

// CreateBlogService 创建博客接口
func (service *BlogService) CreateBlogService(uid uint) serializer.Response {
	var user models.User
	models.DB.First(&user, uid)
	blog := models.Blog{
		BlogName:    service.BlogName,
		BlogContent: service.BlogContent,
		User:        user,
		UserID:      user.ID,
	}
	err := models.DB.Create(&blog).Error
	if err != nil {
		return serializer.Response{
			Status:  300,
			Message: "创建博客失败",
		}
	} else {
		user.BlogCount = user.BlogCount + 1
		models.DB.Save(&user)
		return serializer.Response{
			Status:  200,
			Message: "博客创建成功",
			Data:    serializer.BuildBlog(blog),
		}
	}
}

// GetBlogService 查询单条博客
func (service *BlogService) GetBlogService(uid uint, id string) serializer.Response {
	var blog models.Blog
	err := models.DB.First(&blog, id).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "查询失败",
		}
	} else {
		return serializer.Response{
			Status:  200,
			Message: "信息如下",
			Data:    serializer.BuildBlog(blog),
		}
	}
}

// GetBlogsService 查询用户发布的所有博客
func (service *BlogService) GetBlogsService(loginId uint, uid string) serializer.Response {
	var blogs []models.Blog
	var count int64
	models.DB.Model(&models.Blog{}).Preload("User").Where("user_id = ?", uid).Count(&count).Find(&blogs)
	return serializer.Response{
		Status:  200,
		Message: "用户发布如下博客内容",
		Data:    serializer.BuildBlogs(blogs),
		Count:   count,
	}
}

// UpdateBlogService 更新博客内容
func (service *BlogService) UpdateBlogService(uid uint, id string) serializer.Response {
	var blog models.Blog
	err := models.DB.First(&blog, id).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "博客不存在",
		}
	} else {
		blog.BlogName = service.BlogName
		blog.BlogContent = service.BlogContent
		models.DB.Save(&blog)
		return serializer.Response{
			Status:  200,
			Message: "更新完成",
			Data:    serializer.BuildBlog(blog),
		}
	}
}

// SearchBlogService 模糊查询接口
type SearchBlogService struct {
	Info     string `form:"info" json:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

// SearchBlogsService 模糊查询博客内容
func (service *SearchBlogService) SearchBlogsService(uid uint) serializer.Response {
	var blogs []models.Blog
	var count int64
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	models.DB.Model(&models.Blog{}).Preload("User").Where("user_id = ?", uid).
		Where("blog_name LIKE ? OR blog_content LIKE ?", service.Info+"%", "%"+service.Info+"%").
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&blogs)
	return serializer.Response{
		Status:  200,
		Message: "查询结果如下",
		Data:    serializer.BuildBlogs(blogs),
		Count:   count,
	}
}

// DeleteBlogService 删除单条博客
func (service *BlogService) DeleteBlogService(uid uint, id string) serializer.Response {
	var blog models.Blog
	var user models.User
	models.DB.Model(&models.User{}).Where("user_id = ?", uid)
	models.DB.First(&blog, id)
	err := models.DB.Delete(&blog).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "删除失败",
			Error:   err.Error(),
		}
	} else {
		if user.BlogCount >= 1 {
			user.BlogCount = user.BlogCount - 1
		} else {
			user.BlogCount = 0
		}
		models.DB.Save(&user)
		return serializer.Response{
			Status:  200,
			Message: "删除成功",
		}
	}
}

// DeleteBlogsService 删除用户所有博客
func (service *BlogService) DeleteBlogsService(uid uint, id string) serializer.Response {
	var blogs []models.Blog
	var count int64
	var user models.User
	models.DB.Model(&models.User{}).Where("user_id = ?", uid)
	err := models.DB.Model(&models.Blog{}).Where("user_id = ?", uid).Count(&count).Find(&blogs).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "删除失败",
		}
	} else {
		user.BlogCount = 0
		models.DB.Save(&user)
		models.DB.Delete(&blogs)
		return serializer.Response{
			Status:  200,
			Message: "删除成功",
			Count:   count,
		}
	}
}
