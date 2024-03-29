package controller

import (
	"latihan_api/exception"
	"latihan_api/helper"
	"latihan_api/model/web"
	"latihan_api/service"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

var JWT_KEY = []byte("my_secret_key")

type Claims struct {
	Username string
	RoleID   int
	jwt.RegisteredClaims
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	controller.UserService.Register(userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Register successful, please login",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody web.UserLoginRequest
	helper.ReadFromRequestBody(request, &requestBody)
	userResponse, err := controller.UserService.Login(requestBody)
	if err != nil {
		exception.ErrorHandler(writer, request, exception.NewNotFoundError(err.Error()))
		return
	}

	expTime := time.Now().Add(time.Hour * 1)
	claims := &Claims{
		Username: requestBody.Username,
		RoleID:   userResponse.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWT_KEY)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}
	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Login successful, hello " + userResponse.Username,
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cookie, err := request.Cookie("token")
	if err != nil {
		panic(exception.NewNotAuthorize("login first"))

	}
	tokenString := cookie.Value
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})
	if err != nil || !token.Valid {
		panic(exception.NewNotAuthorize(err.Error()))

	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		panic(exception.NewNotAuthorize("login first"))

	}

	expirationTime := claims.ExpiresAt.Time.Unix()
	if expirationTime < time.Now().Unix() {
		panic(exception.NewNotAuthorize("login first"))

	}

	if claims.RoleID != 1 {
		panic(exception.NewNotAuthorize("Only for admin"))

	}

	userResponses := controller.UserService.FindAll()

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, err := request.Cookie("token")
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "No user to log out",
			Data:   "You haven't logged in yet",
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "logout successful",
	}
	helper.WriteToResponseBody(writer, webResponse)
}
