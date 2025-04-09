package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/suleymanoff/kursach/internal/pkg/jwt"
	"github.com/suleymanoff/kursach/internal/pkg/middlewares"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middlewares.SetCORSHeaders(w)

		var newUser struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Password hashing failed", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec(
			"INSERT INTO users (email, password_hash, role, full_name) VALUES ($1, $2, $3, $4)",
			newUser.Email,
			string(hashedPassword),
			newUser.Role,
			newUser.Name,
		)

		if err != nil {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middlewares.SetCORSHeaders(w)

		var creds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		var storedUser struct {
			ID           int
			PasswordHash string
			Role         string
		}

		err := db.QueryRow(
			"SELECT id, password_hash, role FROM users WHERE email = $1",
			creds.Email,
		).Scan(&storedUser.ID, &storedUser.PasswordHash, &storedUser.Role)

		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword(
			[]byte(storedUser.PasswordHash),
			[]byte(creds.Password),
		); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token, err := jwt.GenerateJWT(storedUser.ID, storedUser.Role)
		if err != nil {
			http.Error(w, "Token generation failed", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"status": "success",
			"token":  token,
		})
	}
}
