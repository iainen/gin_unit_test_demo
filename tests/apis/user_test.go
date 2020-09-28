package apis

import (
	"demo/internal/router"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestInit(t *testing.T) {
	suite.Run(t, new(ApiSuite))
}

func (s *ApiSuite) Test_CreateUser() {
	uri := router.UriApiCreateUser

	tests := []struct {
		name     string
		password string
		email    string
		resCode  int
	}{
		{"", "", "email", http.StatusBadRequest},
		{"zhangyi", "123456", "email", http.StatusOK},
		{"test", "test", "test", http.StatusInternalServerError},
	}

	for _, test := range tests {
		data := struct {
			Name     string `json:"name"`
			Password string `json:"password"`
			Email    string `json:"email"`
		}{
			Name:     test.name,
			Password: test.password,
			Email:    test.email,
		}

		code, _ := s.httpClient.Post(uri, data)

		assert.Equal(s.T(), test.resCode, code, "http reponse status_code not equal")
	}
}

func (s *ApiSuite) Test_GetUser() {
	uri := router.UriApiGetUser

	tests := []struct {
		name    string
		resCode int
	}{
		{"test", http.StatusOK},
		{"not found", http.StatusNotFound},
		{"", http.StatusBadRequest}, // test must
	}

	for _, test := range tests {
		URL, _ := url.Parse(uri)
		params := url.Values{}
		params.Set("name", test.name)
		URL.RawQuery = params.Encode()

		code, _ := s.httpClient.Get(URL.String())

		assert.Equal(s.T(), test.resCode, code, "http reponse status_code not equal")
	}
}
