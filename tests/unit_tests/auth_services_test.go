package unit_tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/appointment-octopus/auth/services"
	"github.com/appointment-octopus/auth/services/models"
	"github.com/appointment-octopus/auth/settings"
	"github.com/stretchr/testify/suite"
)

type AuthenticationServicesTestSuite struct {
	suite.Suite
}

func (suite *AuthenticationServicesTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
}

func (suite *AuthenticationServicesTestSuite) TestLogin() {
	user := models.User{
		Email:    "haku@email.com",
		Password: "testing",
	}
	response, token := services.Login(&user)
	suite.Equal(http.StatusOK, response)
	suite.NotEmpty(token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginIncorrectPassword() {
	user := models.User{
		Email:    "haku@email.com",
		Password: "Password",
	}
	response, token := services.Login(&user)
	suite.Equal(http.StatusUnauthorized, response)
	suite.Empty(token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginIncorrectEmail() {
	user := models.User{
		Email:    "Email",
		Password: "testing",
	}
	response, token := services.Login(&user)
	suite.Equal(http.StatusUnauthorized, response)
	suite.Empty(token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginEmptyCredentials() {
	user := models.User{
		Email:    "",
		Password: "",
	}

	response, token := services.Login(&user)

	suite.Equal(http.StatusUnauthorized, response)
	suite.Empty(token)
}

func (suite *AuthenticationServicesTestSuite) TestRefreshToken() {
	user := models.User{
		Email:    "haku@email.com",
		Password: "testing",
	}

	newToken := services.RefreshToken(&user)
	suite.NotEmpty(newToken)
}

func TestAuthenticationServicesTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationServicesTestSuite))
}
