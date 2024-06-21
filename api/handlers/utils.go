package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/vert-capital/vertc-go-admin/api/entity"
	"github.com/vert-capital/vertc-go-admin/entity"
)

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return true
	}
	return false
}

func jsonResponse(c *gin.Context, httpStatus int, data any) {
	c.JSON(httpStatus, data)
}

func runAction(table entity.Table, action string) (response api.ResponseCreateUpdateDelete, err error) {
	action_func := table.Actions[action]
	if action_func == nil {
		return api.ResponseCreateUpdateDelete{
			Message: "Action not found",
		}, nil
	}
	err = action_func()
	if err != nil {
		return api.ResponseCreateUpdateDelete{
			Message: err.Error(),
		}, err
	}
	return api.ResponseCreateUpdateDelete{
		Message: "Action executed",
	}, nil

}
