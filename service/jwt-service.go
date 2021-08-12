package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// new jwtservice method to create a new instance JWTService
func NewJwtService() JWTService {
	return &jwtService{
		issuer:    "fastbeet",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "fastbeet"
	}
	return secretKey
}

func (s *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}


func (j *jwtService) ValidateToken(token string) (*jwt.Token, error){
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error)){
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Error("Unexpected signing method%v", t.Header["alg"])
		}
		return []byte(jwt.secretKey), nil
	}
}