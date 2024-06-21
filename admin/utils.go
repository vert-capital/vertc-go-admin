package vertc_go_admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func runAction(table Table, action string) (response ResponseCreateUpdateDelete, err error) {
	actions := *table.Actions
	action_func := actions[action]
	if action_func == nil {
		return ResponseCreateUpdateDelete{
			Message: "Action not found",
		}, nil
	}
	err = action_func()
	if err != nil {
		return ResponseCreateUpdateDelete{
			Message: err.Error(),
		}, err
	}
	return ResponseCreateUpdateDelete{
		Message: "Action executed",
	}, nil

}
