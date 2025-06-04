 // PATH: go-auth/utils/GenerateHashPassword.go

 package utils

 import  (
     "go-auth/models"
	 "golang.org/x/crypto/bcrypt"
     "github.com/dgrijalva/jwt-go"
 )

 func GenerateHashPassword(password string) (string, error) {
     bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
     return string(bytes), err
 }


  // PATH: go-auth/utils/CompareHashPassword.go

 func CompareHashPassword(password, hash string) bool {
     err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
     return err == nil
 }


  // PATH: go-auth/utils/ParseToken.go

 

 func ParseToken(tokenString string) (claims *models.Claims, err error) {
     token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
         return []byte("DavidGoggins@123456789"), nil
     })

     if err != nil {
         return nil, err
     }

     claims, ok := token.Claims.(*models.Claims)

     if !ok {
         return nil, err
     }

     return claims, nil
 }

