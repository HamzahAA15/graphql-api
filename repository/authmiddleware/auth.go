package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"sirclo/gql/entities"
	"sirclo/gql/pkg/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// type UserR struct {
// 	userRepo user.RepositoryUser
// }

// type Middleware interface {
// 	Middleware() func(http.Handler) http.Handler
// }

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// if !strings.Contains(header, "Bearer") {

			// }

			// // Allow unauthenticated users in
			// if header == "" {
			// 	fmt.Println("haho")
			// 	fmt.Errorf("unauthorized")
			// 	next.ServeHTTP(w, r)
			// 	return
			// }
			if !strings.Contains(header, "Bearer") {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := ""
			arrayToken := strings.Split(header, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}

			//validate jwt token
			// tokenStr := header
			name, err := jwt.ParseToken(tokenString)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			userS := entities.User{Name: name}
			// id, err := user.GetUserIdByUsername(name)
			// if err != nil {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }
			// userS.Id = id
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &userS)
			fmt.Println("ini ctx", ctx)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *entities.User {
	raw, _ := ctx.Value(userCtxKey).(*entities.User)
	return raw
}
