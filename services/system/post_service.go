package service

import (
	request "resetgoapi.com/rest_go_api/common/request/system"
	"resetgoapi.com/rest_go_api/global"
	models "resetgoapi.com/rest_go_api/models/system"
	"resetgoapi.com/rest_go_api/pkg/db/scopes"
)

var PostService = postService{}

type postService struct{}

type IPostService interface {
	Page(request *request.PostListRequest) (*models.SysPost, int64, error)
	Create(request *request.CreatePostRequest) error
	Update(request *request.UpdatePostRequest) error
	Delete(id int64) error
	FindOne(id int64) (*models.SysPost, error)
	Find() ([]*models.SysPost, error)
}

func (p postService) Page(request *request.PostListRequest) (posts *models.SysPost, total int64, err error) {
	err = global.GORM.Scopes(scopes.Paginate(&request.PageRequest)).Find(&posts).Count(&total).Error
	return
}

func (p postService) Create(request *request.CreatePostRequest) (err error) {
	err = global.GORM.Create(&models.SysPost{
		PostCode: request.PostCode,
		PostName: request.PostName,
		PostSort: request.PostSort,
		Status:   request.Status,
		Remark:   request.Remark,
	}).Error
	return
}

func (p postService) Update(request *request.UpdatePostRequest) (err error) {
	var post models.SysPost
	err = global.GORM.Model(&post).Updates(&models.SysPost{
		PostCode: request.PostCode,
		PostName: request.PostName,
		PostSort: request.PostSort,
		Status:   request.Status,
		Remark:   request.Remark,
	}).Error
	return
}

func (p postService) Delete(id int64) (err error) {
	err = global.GORM.Where("id = ?", id).Delete(&models.SysPost{}).Error
	return
}

func (p postService) FindOne(id int64) (post *models.SysPost, err error) {
	err = global.GORM.Where("id = ?", id).First(&post).Error
	return
}

func (p postService) Find() (posts []*models.SysPost, err error) {
	err = global.GORM.Find(&posts).Error
	return
}
