package unit_tests

import (
	"github.com/appointment-octopus/auth/services"
	"github.com/appointment-octopus/auth/services/models"
	"github.com/appointment-octopus/auth/settings"
	"net/http"
	"os"
	"testing"
	"github.com/stretchr/testify/suite")

type AuthenticationServicesTestSuite struct {
	suite.Suite
}

func (suite *AuthenticationServicesTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
}

func (suite *AuthenticationServicesTestSuite) TestLogin() {
	user := models.User{
		Username: "haku",
		Password: "testing",
	}
	response, token := services.Login(&user)
	suite.Equal(http.StatusOK, response)
	suite.NotEmpty(token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginIncorrectPassword() {
	user := models.User{
		Username: "haku",
		Password: "Password",
	}
	response, token := services.Login(&user)
	suite.Equal(http.StatusUnauthorized, response)
	suite.Empty(token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginIncorrectUsername() {
	user := models.User{
		Username: "Username",
		Password: "testing",
	}
	response, token := services.Login(&user)
	suite.Equal(http.StatusUnauthorized, response)
	suite.Empty(token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginEmptyCredentials() {
	user := models.User{
		Username: "",
		Password: "",
	}
	
	response, token := services.Login(&user)


	suite.Equal(http.StatusUnauthorized, response)
	suite.Empty(token)
}

func (suite *AuthenticationServicesTestSuite) TestRefreshToken() {
	user := models.User{
		Username: "haku",
		Password: "testing",
	}

	newToken := services.RefreshToken(&user)
	suite.NotEmpty(newToken)
}


func TestAuthenticationServicesTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationServicesTestSuite))
}