package unit_tests

import (
	"os"
	"testing"

	"github.com/appointment-octopus/auth/core/authentication"
	"github.com/appointment-octopus/auth/core/db"
	"github.com/appointment-octopus/auth/services/models"
	"github.com/appointment-octopus/auth/settings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/suite"
)

type AuthenticationBackendTestSuite struct {
	suite.Suite
}

func (suite *AuthenticationBackendTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
}

func (suite *AuthenticationBackendTestSuite) TestInitJWTAuthenticationBackend() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	suite.NotNil(authBackend)
	suite.NotNil(authBackend.PublicKey)
}

func (suite *AuthenticationBackendTestSuite) TestGenerateToken() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenString, err := authBackend.GenerateToken(1234)

	suite.Nil(err)
	suite.NotEmpty(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})

	suite.Nil(err)
	suite.True(token.Valid)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticate() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	user := &models.User{
		Email:    "haku@email.com",
		Password: "testing",
	}
	suite.Equal(authBackend.Authenticate(user), true)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticateIncorrectPass() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	user := models.User{
		Email:    "haku@email.com",
		Password: "Password",
	}
	suite.Equal(authBackend.Authenticate(&user), false)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticateIncorrectEmail() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	user := &models.User{
		Email:    "Haku@email.com",
		Password: "testing",
	}
	suite.Equal(authBackend.Authenticate(user), false)
}

func (suite *AuthenticationBackendTestSuite) TestLogout() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenString, err := authBackend.GenerateToken(1234)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	err = authBackend.Logout(tokenString, token)

	suite.Nil(err)

	redisValue, err := db.RedisGetValue(tokenString)
	suite.Nil(err)
	suite.NotEmpty(redisValue)
}

func (suite *AuthenticationBackendTestSuite) TestIsInBlacklist() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenString, err := authBackend.GenerateToken(1234)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	err = authBackend.Logout(tokenString, token)
	suite.Nil(err)

	suite.True(authBackend.IsInBlacklist(tokenString))
}

func (suite *AuthenticationBackendTestSuite) TestIsNotInBlacklist() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	suite.False(authBackend.IsInBlacklist("1234"))
}

func TestAuthenticationBackendTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationBackendTestSuite))
}
