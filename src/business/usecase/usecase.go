package usecase

import (
	"github.com/Beelzebub0/go-crud-boilerplate/src/business/domain"
	"github.com/Beelzebub0/go-crud-boilerplate/src/business/entity"
	"github.com/gin-gonic/gin"
)

type Usecase interface {

	// Area
	GetArea(c *gin.Context, params entity.AreaParams) ([]entity.Area, entity.Pagination, error)
	GetAreaByID(c *gin.Context, aid int64) (entity.Area, error)
	CreateArea(c *gin.Context, params entity.AreaInput) (entity.Area, error)
	UpdateArea(c *gin.Context, params entity.UpdateAreaInput) (entity.Area, error)
	DeleteArea(c *gin.Context, aid int64) error
}

type usecase struct {
	dom domain.Domain
}

func Init(dom domain.Domain) Usecase {
	return &usecase{
		dom: dom,
	}
}
