package handler

import (
	"context"
	"first_init/model"
	"first_init/proto/login"
	"first_init/service"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthenticationHandler(service *service.UserService) *LoginHandler {
	return &LoginHandler{
		UserService: service,
	}
}

type LoginHandler struct {
	login.UnimplementedLoginServiceServer
	UserService *service.UserService
	writer      http.ResponseWriter
}

type UserHandler struct {
	UserService *service.UserService
}

func parseToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func userClaimFromToken(claims jwt.MapClaims) string {

	sub, ok := claims["role"].(string)
	if !ok {
		return ""
	}

	return sub
}
func userClaimmFromToken(claims jwt.MapClaims) string {

	sub, ok := claims["username"].(string)
	if !ok {
		return ""
	}

	return sub
}

func checkRole(ctx context.Context) string {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ""
	}

	tokenInfo, _ := parseToken(token)

	role := userClaimFromToken(tokenInfo)

	fmt.Println("role is: " + role)
	return role
}
func checkUsername(ctx context.Context) string {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ""
	}

	tokenInfo, _ := parseToken(token)

	role := userClaimmFromToken(tokenInfo)

	fmt.Println("username is: " + role)
	return role
}

func (h LoginHandler) GreetFromLogin(ctx context.Context, request *login.Request) (*login.Response, error) {
	return &login.Response{
		Greeting: fmt.Sprintf("Hihi %s!", request.Name),
	}, nil
}
func (h LoginHandler) CreateUser(ctx context.Context, request *login.CreateUserRequest) (*login.CreateUserResponse, error) {
	//var User *model.User
	var User = model.User{}
	User.Name = request.GetReg().Name
	User.Surname = request.GetReg().Surname
	User.Email = request.GetReg().Email
	User.Username = request.GetReg().Username
	User.Password = request.GetReg().Password
	User.Role = "User"
	// print the JSON string
	h.UserService.Create(&User)
	fmt.Println("Iznad je request")
	return &login.CreateUserResponse{
		Reg: &login.User{},
	}, nil
}
func (h LoginHandler) UpdateUser(ctx context.Context, request *login.UpdateRequest) (*login.UpdateResponse, error) {
	//var User *model.User
	role := checkRole(ctx)
	if role == "User" || role == "Host" {
		var User = model.User{}
		User.ID, _ = primitive.ObjectIDFromHex(request.GetReg().Id)
		User.Name = request.GetReg().Name
		User.Surname = request.GetReg().Surname
		User.Email = request.GetReg().Email
		User.Username = request.GetReg().Username
		User.Password = request.GetReg().Password
		User.Role = "User"
		// print the JSON string
		h.UserService.UpdateUser(&User)
		fmt.Println("Iznad je request")
		return &login.UpdateResponse{
			Reg: &login.UpdateInfo{},
		}, nil
	}
	return &login.UpdateResponse{
		Reg: &login.UpdateInfo{},
	}, status.Errorf(codes.Unauthenticated, "You don't have permissions for this action")
}
func (h LoginHandler) DeleteUser(ctx context.Context, request *login.DeleteRequest) (*login.DeleteResponse, error) {
	//var User *model.User
	// print the JSON string
	role := checkRole(ctx)
	usr := checkUsername(ctx)
	if role == "User" || role == "Host" {
		fmt.Println(request.Dlt)
		h.UserService.DeleteUser(usr)
		fmt.Println("Iznad je request")
		return &login.DeleteResponse{
			Dlt: "Deleted",
		}, nil
	}
	return &login.DeleteResponse{
		Dlt: "Deleted",
	}, status.Errorf(codes.Unauthenticated, "You don't have permissions for this action")
}
func (h LoginHandler) GetUser(ctx context.Context, request *login.RequestGetUser) (*login.ResponseGetUser, error) {
	usr := checkUsername(ctx)
	Userr, _ := h.UserService.FindByUsername(usr)
	fmt.Println("Iznad je request")
	accoInfo := &login.UserInfo{
		Id:       Userr.ID.Hex(),
		Name:     Userr.Name,
		Surname:  Userr.Surname,
		Username: Userr.Username,
		Password: Userr.Password,
		Email:    Userr.Email,
	}
	return &login.ResponseGetUser{
		Usr: accoInfo,
	}, nil
}

func (h LoginHandler) Login(ctx context.Context, request *login.LoginRequest) (*login.LoginResponse, error) {
	//var User *model.User
	var User = model.User{}
	fmt.Println(User)
	tokenString := ""
	Userr, er := h.UserService.FindUserForLogin(request.Logg.Username, request.Logg.Password)
	if er != nil {
		//fmt.Sprintf("Error")
	}
	fmt.Println(Userr)
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: request.Logg.Username,
		Role:     Userr.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		//fmt.Sprintf("Error")
	}

	return &login.LoginResponse{
		Token: tokenString,
	}, nil
}

// func (handler *UserHandler) Get(writer http.ResponseWriter, req *http.Request) {
// 	id := mux.Vars(req)["id"]
// 	User, err := handler.UserService.FindUser(id)
// 	writer.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		writer.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusOK)
// 	json.NewEncoder(writer).Encode(User)
// }

// func (handler *UserHandler) Register(writer http.ResponseWriter, req *http.Request) {
// 	var User *model.User
// 	err := json.NewDecoder(req.Body).Decode(&User)
// 	User.Role = "User"
// 	if err != nil {
// 		println("Error while parsing json")
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = handler.UserService.Create(User)
// 	if err != nil {
// 		println("Error while creating a new User")
// 		writer.WriteHeader(http.StatusExpectationFailed)
// 		return
// 	}
// 	writer.WriteHeader(http.StatusCreated)
// 	writer.Header().Set("Content-Type", "application/json")
// }
// func (handler *UserHandler) Login(writer http.ResponseWriter, req *http.Request) {
// 	var User *model.User
// 	err := json.NewDecoder(req.Body).Decode(&User)
// 	if err != nil {
// 		println("Error while parsing json")
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = handler.UserService.FindUserByUsernameAndPassword(User.Username, User.Password)
// 	if err == nil {
// 		writer.WriteHeader(http.StatusCreated)
// 		writer.Header().Set("Content-Type", "application/json")
// 	}

// }

var jwtKey = []byte("secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// func (handler *UserHandler) Loginn(w http.ResponseWriter, r *http.Request) {
// 	var credentials *Credentials
// 	err := json.NewDecoder(r.Body).Decode(&credentials)
// 	fmt.Println(credentials.Username + " " + credentials.Password)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	var Userr *model.User
// 	Userr, er := handler.UserService.FindUserForLogin(credentials.Username, credentials.Password)
// 	if er != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	expirationTime := time.Now().Add(time.Minute * 5)

// 	claims := &Claims{
// 		Username: credentials.Username,
// 		Role:     Userr.Role,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	http.SetCookie(w, &http.Cookie{
// 		Name:    "token",
// 		Value:   tokenString,
// 		Expires: expirationTime,
// 	})
// 	w.WriteHeader(http.StatusOK)
// 	a, _ := json.Marshal(tokenString)
// 	w.Write(a)
// }

// func (handler *UserHandler) Home(w http.ResponseWriter, r *http.Request) {

// 	cookie := r.Header.Get("Authorization")
// 	if cookie == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	tokenStr := cookie
// 	claims := &Claims{}
// 	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	a, _ := json.Marshal("Hello, " + claims.Username)
// 	w.Write(a)

// }
// func (handler *UserHandler) Cao(w http.ResponseWriter, r *http.Request) {

// 	w.WriteHeader(http.StatusOK)
// 	a, _ := json.Marshal("cao")
// 	w.Write(a)

// }
