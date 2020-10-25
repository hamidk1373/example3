package users

import (
	"encoding/json"
	"fmt"
	"hamid/example3/helpers"
	"hamid/example3/settings"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func readAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := getAllFromDB()
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "Error getting data from database", err)
		return
	}

	helpers.SendSuccessResponse(w, http.StatusOK, helpers.JSONResponse{"users": users})
}

func readOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ID, _ := strconv.Atoi(params.ByName("id"))
	user, err := getOneFromDB(uint(ID))
	if err != nil {
		helpers.SendError(w, http.StatusNotFound, "User not Found", err)
		return
	}

	helpers.SendSuccessResponse(w, http.StatusOK, helpers.JSONResponse{"user": user})
}

func create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := new(User)

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "Error Saveing user into databas1", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Your password is not strong enough!", err)
		return
	}

	user.Password = string(hashedPassword)

	err = user.createOneToDB()

	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == pgerrcode.UniqueViolation {
			helpers.SendError(w, http.StatusConflict, fmt.Sprintf("%s alreary exists", pgErr.ConstraintName), nil)
		}
		// helpers.SendError(w, http.StatusInternalServerError, "Error Saveing user into databas2", err)
		return
	}
	helpers.SendSuccessResponse(w, http.StatusCreated, helpers.JSONResponse{"message": "User created successfuly"})
}

func update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ID, _ := strconv.Atoi(params.ByName("id"))
	user := new(User)
	user.ID = uint(ID)

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Invalid user data. Please check your inputs", err)
		return
	}

	err = user.updateOneToDB()
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "Error updating user into databas", err)
		return
	}
	helpers.SendSuccessResponse(w, http.StatusOK, helpers.JSONResponse{"message": "User updated successfully"})
}

func delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ID, _ := strconv.Atoi(params.ByName("id"))
	err := deleteOneFromDB(uint(ID))
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "Error deleting user from databas", err)
		return
	}
	helpers.SendSuccessResponse(w, http.StatusOK, helpers.JSONResponse{"message": "User deleted successfully"})
}

func register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := new(User)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "Error Saveing user into databas1", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Your password is not strong enough!", err)
		return
	}

	user.Password = string(hashedPassword)

	err = user.createOneToDB()
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == pgerrcode.UniqueViolation {
			helpers.SendError(w, http.StatusConflict, fmt.Sprintf("%s alreary exists", strings.TrimPrefix(pgErr.ConstraintName, "idx_users_")), nil)
		}
		return
	}
	helpers.SendSuccessResponse(w, http.StatusCreated, helpers.JSONResponse{"message": "User created successfuly"})
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loginFormData := new(LoginUser)

	err := json.NewDecoder(r.Body).Decode(&loginFormData)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "invalid data", nil)
		return
	}
	user, err := getOneByEmailFromDB(loginFormData.Email)
	if err != nil {
		helpers.SendError(w, http.StatusNotFound, "Email isn't registered.", err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginFormData.Password))
	if err != nil {
		helpers.SendError(w, http.StatusUnauthorized, "Wrong password", err)
		return
	}

	claims := &settings.JWTCustomClaims{
		user.Email,
		// u.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, "Server Error!!", err)
	}

	helpers.SendSuccessResponse(w, http.StatusOK, helpers.JSONResponse{"message": "logged in successfully", "token": t})
}
