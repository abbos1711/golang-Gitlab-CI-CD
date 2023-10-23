package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	"gitlab.com/tizim-back/api/models"
	"gitlab.com/tizim-back/pkg/logger"
)

// @Summary 	Login
// @Description This api can login
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param body	body models.LoginReq true "Login"
// @Success 201 {object} models.LoginRes
// @Failure 400 string Error response
// @Router /v1/auth/login [post]
func (h *handlerV1) Login(ctx *gin.Context) {
	var req models.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	user, err := h.Storage.User().GetUserByUserName(ctx, req.UserName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "user not found",
			})
			return
		}
		h.log.Error("Error while getting user by username", logger.Any("user", err))
		return
	}

	if req.Password != user.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username or password is wrong",
		})
		h.log.Error("Error while matching password", logger.Any("user", err))
		return
	}

	h.jwthandler.Role = "admin"
	h.jwthandler.Aud = []string{"user-api"}
	h.jwthandler.SigninKey = h.cfg.SignKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]
	if err != nil {
		h.log.Error("error occured while generating tokens")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.LoginRes{
		Token: accessToken,
	})
}

// InitializeRedisClient initializes and returns a Redis client.
func InitializeRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default database
	})

	return client
}
