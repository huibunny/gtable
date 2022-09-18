package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gtable/config"
	"gtable/internal/usecase"
	"gtable/pkg/logger"
)

type tableRoutes struct {
	t      usecase.Table
	l      logger.Interface
	tables []config.Table
}

func newtableRoutes(handler *gin.RouterGroup, t usecase.Table, l logger.Interface, tables []config.Table) {
	r := &tableRoutes{t, l, tables}

	h := handler.Group("/table")
	{
		h.GET("/load", r.load)
	}
}

type appResponse struct {
	ErrCode int         `json:"errcode" example:"0"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data" example:"{}"`
}

// @Summary     Load
// @Description Load data from table
// @ID          load
// @Tags  	    load
// @Accept
// @Produce     json
// @Param
// @Success     200 {object} loginResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /table/load [get]
func (r *tableRoutes) load(c *gin.Context) {

	data, errcode, err := r.t.Load(c.Request.Context(), r.tables)

	message := "success"
	if err != nil {
		message = "fail to load data from table"
	}

	c.JSON(http.StatusOK, appResponse{ErrCode: errcode, Message: message, Data: data})
}
