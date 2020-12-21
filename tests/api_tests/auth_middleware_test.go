package api_tests

import (
	"github.com/appointment-octopus/auth/services"
	"github.com/appointment-octopus/auth/routers"
	"github.com/appointment-octopus/auth/settings"
	"github.com/appointment-octopus/auth/core/authentication"
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
	"net/http/httptest"
	"os"
	"github.com/stretchr/testify/suite"
	"testing"
)

var token string
var server *negroni.Negroni

type MiddlewaresTestSuite struct {
	suite.Suite
}

func (suite *MiddlewaresTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
}

func (suite *MiddlewaresTestSuite) SetupTest() {
	authBackend := authentication.InitJWTAuthenticationBackend()
	suite.NotNil(authBackend)
	token, _ = authBackend.GenerateToken("1234")

	router := routers.InitRoutes()
	server = negroni.Classic()
	server.UseHandler(router)
}

func (suite *MiddlewaresTestSuite) TestRequireTokenAuthentication() {
	resource := "/test/hello"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	server.ServeHTTP(response, request)

	suite.Equal(response.Code, http.StatusOK)
}

func (suite *MiddlewaresTestSuite) TestRequireTokenAuthenticationInvalidToken() {
	resource := "/test/hello"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "token"))
	server.ServeHTTP(response, request)

	suite.Equal(response.Code, http.StatusUnauthorized)
}

func (suite *MiddlewaresTestSuite) TestRequireTokenAuthenticationEmptyToken() {
	resource := "/test/hello"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", ""))
	server.ServeHTTP(response, request)

	suite.Equal(response.Code, http.StatusUnauthorized)
}

func (suite *MiddlewaresTestSuite) TestRequireTokenAuthenticationWithoutToken() {
	resource := "/test/hello"

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	server.ServeHTTP(response, request)

	suite.Equal(response.Code, http.StatusUnauthorized)
}

func (suite *MiddlewaresTestSuite) TestRequireTokenAuthenticationAfterLogout() {
	resource := "/test/hello"

	requestLogout, _ := http.NewRequest("GET", resource, nil)
	requestLogout.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	services.Logout(requestLogout)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", resource, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	server.ServeHTTP(response, request)

	suite.Equal(response.Code, http.StatusUnauthorized)
}

func TestAuthenticationBackendTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewaresTestSuite))
}