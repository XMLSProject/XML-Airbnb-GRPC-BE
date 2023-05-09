package handler

import (
	"context"
	"first_init/proto/login"
	"first_init/service"
	"fmt"
)

func NewAuthenticationHandler(service *service.UserService) *LoginHandler {
	return &LoginHandler{
		UserService: service,
	}
}

type LoginHandler struct {
	login.UnimplementedLoginServiceServer
	UserService *service.UserService
}

type UserHandler struct {
	UserService *service.UserService
}

func (h LoginHandler) GreetFromLogin(ctx context.Context, request *login.Request) (*login.Response, error) {

	fmt.Println("Uso u hendler")
	User, _ := h.UserService.FindUser("Joca")
	fmt.Println(User)
	fmt.Println("Evo greske")
	return &login.Response{
		Greeting: fmt.Sprintf("Hihi %s!", request.Name),
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

// var jwtKey = []byte("secret_key")

// type Credentials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type Claims struct {
// 	Username string `json:"username"`
// 	Role     string `json:"role"`
// 	jwt.StandardClaims
// }

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
