package handlers

import (
	"api_gateway/api/http"
	"api_gateway/genproto/auth_service"
	"api_gateway/pkg/util"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param profile body auth_service.CreateUser true "CreateUserRequestBody"
// @Success 201 {object} auth_service.User "UserResponse"
// @Failure 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var User auth_service.CreateUser
	err := c.ShouldBindJSON(&User)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	fmt.Println("just created")

	resp, err := h.services.UserService().Create(
		c.Request.Context(),
		&User,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, resp)
}

// GetUserByID godoc
// @ID get_user_by_id
// @Router /user/{id} [GET]
// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} auth_service.User "UserResponse"
// @Failure 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUserById(c *gin.Context) {
	UserID := c.Param("id")
	if !util.IsValidUUID(UserID) {
		h.handleResponse(c, http.InvalidArgument, "User ID is an invalid UUID")
		return
	}
	resp, err := h.services.UserService().GetByID(
		context.Background(),
		&auth_service.UserPK{
			Id: UserID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, resp)
}

// GetMyself godoc
// @ID get_user_myself
// @Router /myself [GET]
// @Summary Get Myself
// @Description Get Myself
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} auth_service.User "User"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetMyself(c *gin.Context) {
	
    resp, err := h.services.UserService().GetUserByName(
        context.Background(),
        &auth_service.GetByName{
            Name: "abdurahmon", 
        },
    )
    if err != nil {
        h.handleResponse(c, http.GRPCError, err.Error())
        return
    }

    h.handleResponse(c, http.OK, resp)
}


//  @Security ApiKeyAuth
// GetUserList godoc
// @ID get_user_list
// @Router /user [GET]
// @Summary Get User List
// @Description Get User List
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "Offset"
// @Param limit query integer false "Limit"
// @Param search query string false "Search"
// @Success 200 {object} auth_service.UserListResponse "UserListResponse"
// @Failure 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) GetUserList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.UserService().GetUserList(
		context.Background(),
		&auth_service.UserListRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
