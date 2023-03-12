package midleware

import (
	"Tugas-Pert4/service"
	"Tugas-Pert4/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.ErrorResponse("No token found", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.ErrorResponse("No token found", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.ErrorResponse("Invalid token", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := utils.ErrorResponse("Invalid token", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		teamRole, err := jwtService.GetRoleByToken(string(authHeader))
		fmt.Println("ROLE", teamRole)
		if err != nil || (teamRole != "admin" && teamRole != role) {
			response := utils.ErrorResponse("Failed to process request", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		// get userID from token
		teamID, err := jwtService.GetTeamIDByToken(authHeader)
		if err != nil {
			response := utils.ErrorResponse("Failed to process request", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		fmt.Println("ROLE", teamRole)
		c.Set("teamID", teamID)
		c.Next()
	}
}
