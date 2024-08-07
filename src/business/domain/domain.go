package domain

import (
	"github.com/Beelzebub0/go-crud-boilerplate/src/business/entity"
	"github.com/Beelzebub0/go-crud-boilerplate/src/lib/database"
	"github.com/gin-gonic/gin"
)

type Domain interface {
	// Area
	GetArea(c *gin.Context, params entity.AreaParams) ([]entity.Area, entity.Pagination, error)
	GetAreaByID(c *gin.Context, aid int64) (entity.Area, error)
	CreateArea(c *gin.Context, params entity.AreaInput) (entity.Area, error)
	UpdateArea(c *gin.Context, params entity.UpdateAreaInput) (entity.Area, error)
	DeleteArea(c *gin.Context, aid int64) error
}

type domain struct {
	sql database.SQL
}

func Init(sql database.SQL) Domain {
	return &domain{
		sql: sql,
	}
}
