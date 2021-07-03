package usersvc

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(gGroup gin.IRouter)
}
type router struct {
	ctrl Controller
}

func NewRouter(ctrl Controller) Router {
	return &router{ctrl: ctrl}
}

func (r *router) Register(gGroup gin.IRouter) {
	g := gGroup.Group("/users")
	{
		g.POST("", r.createUser)
		g.GET("/:id", r.getUserByID)
	}
}

func (r *router) createUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBind(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := r.ctrl.CreateUser(c, req)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	rsp := CreateUserResponse{
		Code:    successCode,
		Message: "success",
		Data:    user,
	}
	c.JSON(http.StatusOK, rsp)
}

func (r *router) getUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := r.ctrl.GetUserByID(c, id)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	rsp := GetUserByIDResponse{
		Code:    successCode,
		Message: "success",
		Data:    user,
	}
	c.JSON(http.StatusOK, rsp)
}
