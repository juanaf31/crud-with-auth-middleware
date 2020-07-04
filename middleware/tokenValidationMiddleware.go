package middleware

import (
	"fmt"
	"liveCodeAPI/utils"
	"net/http"
	"strings"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
		} else {
			tokenVal := strings.Split(token, "Bearer ")
			// resDecode, err := tools.JwtDecoder(tokenVal[1])
			_, err := utils.JwtDecoder(tokenVal[1])
			// fmt.Println(resDecode["customKey"])
			if err != nil {
				http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
				fmt.Println(err)
			} else {
				next.ServeHTTP(w, r)
				// if resDecode["customKey"] == "rahasiadong" {
				// 	next.ServeHTTP(w, r)
				// } else {
				// 	fmt.Println(err)
				// 	http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
				// }
			}
		}
	})
}
