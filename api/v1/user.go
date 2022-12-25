package v1

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurmuhammaddeveloper/blog_db/api/models"
	"github.com/nurmuhammaddeveloper/blog_db/storage/repo"
)

// @Security ApiKeyAuth
// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User"
// @Success 201 {object} models.User
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		req models.CreateUserRequest
	)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	resp, err := h.Storage.User().Create(&repo.User{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		PhoneNumber:     req.PhoneNumber,
		Email:           req.Email,
		Gender:          req.Gender,
		UserName:        req.UserName,
		ProfileImageUrl: req.ProfileImageUrl,
		Type:            req.Type,
		Password:        req.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.User{
		ID:              resp.ID,
		FirstName:       resp.FirstName,
		LastName:        resp.LastName,
		PhoneNumber:     resp.PhoneNumber,
		Email:           resp.Email,
		Gender:          resp.Gender,
		UserName:        resp.UserName,
		ProfileImageUrl: resp.ProfileImageUrl,
		Type:            resp.Type,
		CreatedAt:       resp.CreatedAt,
	})
}

// @Router /users/{id} [get]
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.User
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
func (h *handlerV1) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	resp, err := h.Storage.User().Get(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.User{
		ID:              resp.ID,
		FirstName:       resp.FirstName,
		LastName:        resp.LastName,
		PhoneNumber:     resp.PhoneNumber,
		Email:           resp.Email,
		Gender:          resp.Gender,
		UserName:        resp.UserName,
		ProfileImageUrl: resp.ProfileImageUrl,
		Type:            resp.Type,
		CreatedAt:       resp.CreatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /users/me [get]
// @Summary Get user by token
// @Description Get user by token
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} models.User
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
func (h *handlerV1) GetUserProfile(c *gin.Context) {
	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	resp, err := h.Storage.User().Get(payload.UserID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.User{
		ID:              resp.ID,
		FirstName:       resp.FirstName,
		LastName:        resp.LastName,
		PhoneNumber:     resp.PhoneNumber,
		Email:           resp.Email,
		Gender:          resp.Gender,
		UserName:        resp.UserName,
		ProfileImageUrl: resp.ProfileImageUrl,
		Type:            resp.Type,
		CreatedAt:       resp.CreatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /users/{id} [put]
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param User body models.CreateUserRequest true "User"
// @Success 201 {object} models.User
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		req models.CreateUserRequest
	)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err = c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	result, err := h.Storage.User().Update(&repo.User{
		ID:              id,
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		PhoneNumber:     req.PhoneNumber,
		Email:           req.Email,
		Gender:          req.Gender,
		UserName:        req.UserName,
		ProfileImageUrl: req.ProfileImageUrl,
		Type:            req.Type,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.User{
		ID:              result.ID,
		FirstName:       result.FirstName,
		LastName:        result.LastName,
		PhoneNumber:     result.PhoneNumber,
		Email:           result.Email,
		Gender:          result.Gender,
		UserName:        result.UserName,
		ProfileImageUrl: result.ProfileImageUrl,
		Type:            result.Type,
	})
}

// @Security ApiKeyAuth
// @Router /users/{id} [delete]
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err = h.Storage.User().Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Success: "Successfully deleted!",
	})
}

// @Router /users [get]
// @Summary Get user by giving limit, page and search for something.
// @Description Get user by giving limit, page and search for something.
// @Tags user
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 201 {object} models.GetAllUsersResponse
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) GetAllUsers(c *gin.Context) {
	params, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	result, err := h.Storage.User().GetAll(&repo.GetAllUserParams{
		Limit:  int32(params.Limit),
		Page:   int32(params.Page),
		Search: params.Search,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, getUsersResponse(result))
}

func getUsersResponse(data *repo.GetAllUsersResult) *models.GetAllUsersResponse {
	response := models.GetAllUsersResponse{
		Users: make([]*models.User, 0),
		Count: data.Count,
	}

	for _, user := range data.Users {
		u := parseUserModel(user)
		response.Users = append(response.Users, &u)
	}

	return &response
}

func parseUserModel(user *repo.User) models.User {
	return models.User{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		PhoneNumber:     user.PhoneNumber,
		Email:           user.Email,
		Gender:          user.Gender,
		UserName:        user.UserName,
		ProfileImageUrl: user.ProfileImageUrl,
		Type:            user.Type,
		CreatedAt:       user.CreatedAt,
	}
}
