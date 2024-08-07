package usecase

import (
	"github.com/Beelzebub0/go-crud-boilerplate/src/business/entity"
	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateArea(c *gin.Context, params entity.AreaInput) (entity.Area, error) {
	return uc.dom.CreateArea(c, params)
}

func (uc *usecase) GetAreaByID(c *gin.Context, aid int64) (entity.Area, error) {
	return uc.dom.GetAreaByID(c, aid)
}

func (uc *usecase) GetArea(c *gin.Context, params entity.AreaParams) ([]entity.Area, entity.Pagination, error) {
	return uc.dom.GetArea(c, params)
}

func (uc *usecase) UpdateArea(c *gin.Context, params entity.UpdateAreaInput) (entity.Area, error) {
	return uc.dom.UpdateArea(c, params)
}

func (uc *usecase) DeleteArea(c *gin.Context, aid int64) error {
	return uc.dom.DeleteArea(c, aid)
}
