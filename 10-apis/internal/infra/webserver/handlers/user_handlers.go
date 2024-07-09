package handlers

import (
	"10-apis/internal/dto"
	"10-apis/internal/entity"
	"10-apis/internal/infra/db"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string
}

type UserHandler struct {
	UserDB db.UserInterface
}

func NewUserHandler(db db.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

// @Summary      Get JWT
// @Description  Get JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body  dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate-token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var user dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errorMsg := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500 {object}  Error
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	println("jiba")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMsg := Error{Message: err.Error()}
		println("jeba")
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMsg := Error{Message: err.Error()}
		println("jeba")
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMsg := Error{Message: err.Error()}
		println("jeba")
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get users"))
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user"))
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
