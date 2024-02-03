package midware

// import (
// 	"hr/configs/models"
// 	"math/rand"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// // 生成随机密钥
// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// func generateRandomString(length int) string {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	rng := rand.New(source)

// 	result := make([]byte, length)
// 	for i := range result {
// 		result[i] = charset[rng.Intn(len(charset))]
// 	}

// 	return string(result)
// }

// var jwtKey = []byte(generateRandomString(10))

// type Claims struct {
// 	Userid     string `json:"userid"`
// 	Role       string `json:"role"`
// 	Grade      int    `json:"grade"`
// 	Profession string `json:"profession"`
// 	jwt.StandardClaims
// }

// func AuthenticateMiddleware(c *gin.Context, allowedRoles ...string) {
// 	authorizationHeader := c.GetHeader("Authorization")
// 	if authorizationHeader == "" {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	if !token.Valid {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	claims, ok := token.Claims.(*Claims)
// 	if !ok {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	user, ok := c.Get("currentUser")

// 	currentUser, ok := user.(models.CurrentUser)
// 	if !ok {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}
// 	if claims.Userid != currentUser.UserId { //从上下文中的用户信息中获取用户id与claims核对
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 5*time.Minute {
// 		// 生成新的Token
// 		newToken, err := GenerateNewToken(currentUser)
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusInternalServerError)
// 			return
// 		}

// 		// 将新的Token发送给客户端
// 		c.Header("Authorization", "Bearer "+newToken)
// 	}

// 	// 检查用户的角色是否在允许的角色列表中
// 	roleAllowed := false
// 	for _, allowedRole := range allowedRoles {
// 		if claims.Role == allowedRole {
// 			roleAllowed = true
// 			break
// 		}
// 	}

// 	if !roleAllowed {
// 		c.AbortWithStatus(http.StatusForbidden)
// 		return
// 	}
// 	c.Next()
// }

// // 生成token
// func GenerateNewToken(currentUser models.CurrentUser) (string, error) {
// 	expirationTime := time.Now().Add(30 * time.Minute) // 新Token有效期为30分钟

// 	claims := &Claims{
// 		Userid:     currentUser.UserId,
// 		Role:       currentUser.Role,
// 		Grade:      currentUser.Grade,
// 		Profession: currentUser.Profession,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }
