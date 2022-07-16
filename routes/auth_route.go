package routes

import (
	"net/http"
	"server/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthSuccessResponse struct {
	Message string `json:"message"`
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func (rt *Route) loginRoute(c echo.Context) error {
	var payload AuthRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse("", "One field might be missing, fill in the missing fields."))
	}

	if len(payload.Username) == 0 || len(payload.Password) == 0 {
		return c.JSON(http.StatusBadRequest, utils.NewResponse("", "One field might be missing, fill in the missing fields."))
	}

	user, err := rt.DB.GetUserByUsername(c.Request().Context(), payload.Username)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewResponse(nil, "User doesn't exist."))
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		return c.JSON(http.StatusForbidden, utils.NewResponse(nil, "Incorrect username or password."))
	}

	access_token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(user.Username, user.Email, string(user.UserRole))), rt.Cfg.JWT_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	refresh_token, err := utils.NewJwtToken(utils.NewJwtClaims(utils.NewJwtPayload(user.Username, user.Email, string(user.UserRole))), rt.Cfg.JWT_REFRESH_SECRET_KEY)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}
	return c.JSON(http.StatusOK, utils.NewResponse(AuthSuccessResponse{Message: "Logged in.", Access: access_token, Refresh: refresh_token}, ""))
}
