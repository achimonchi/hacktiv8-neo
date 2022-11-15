package controllers

import (
	"fmt"
	"latihan1/params/request"
	"latihan1/params/views"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// GetTodos godoc
// @Summary Get All TODOS
// @Schemes
// @Description get all todos
// @Tags TODOS
// @Accept json
// @Produce json
// @Success 200 {object} views.GetTodosSuccessSwag
// @Router /todos [get]
func GetAll(c *gin.Context) {

}

// @BasePath /api/v1
// CreateTodo godoc
// @Summary Create TODO
// @Schemes
// @Description create todo
// @Tags TODOS
// @Accept json
// @Produce json
// @Param request body request.CreateTodo  true  "Request Body"
// @Success 200 {object} views.CreateTodoSuccessSwag
// @Failure      400  {object}  views.CreateTodoFailureSwag
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var req request.CreateTodo

	err := c.ShouldBindJSON(&req)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	var payload = views.CreateTodoPayload{
		ID:          1,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	resp := views.GeneralSuccessPayload{
		Status:  http.StatusCreated,
		Message: "CREATE TODO SUCCESS",
		Payload: payload,
	}

	c.JSON(resp.Status, resp)
}

func GetByID(c *gin.Context) {

}

func UpdateByID(c *gin.Context) {

}

func DeleteByID(c *gin.Context) {

}
