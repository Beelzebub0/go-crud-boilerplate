package domain

import (
	"github.com/Beelzebub0/go-crud-boilerplate/src/business/entity"
	"github.com/gin-gonic/gin"
)

func (dom *domain) CreateArea(c *gin.Context, params entity.AreaInput) (entity.Area, error) {
	return dom.sqlCreateArea(c.Request.Context(), params)
}

func (dom *domain) GetAreaByID(c *gin.Context, aid int64) (entity.Area, error) {
	return dom.sqlGetAreaByID(c.Request.Context(), aid)
}

func (dom *domain) GetArea(c *gin.Context, params entity.AreaParams) ([]entity.Area, entity.Pagination, error) {
	return dom.sqlGetArea(c.Request.Context(), params)
}

func (dom *domain) UpdateArea(c *gin.Context, params entity.UpdateAreaInput) (entity.Area, error) {
	return dom.sqlUpdateArea(c.Request.Context(), params)
}

func (dom *domain) DeleteArea(c *gin.Context, aid int64) error {
	return dom.sqlDeleteArea(c.Request.Context(), aid)
}
