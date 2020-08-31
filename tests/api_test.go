package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-starter/config"
	"gin-starter/controllers"
	"gin-starter/mappers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	v1 := r.Group("/v1")
	{
		api := new(controllers.Api)

		v1.POST("/user/login", api.Login)
		v1.POST("/user/register", api.Register)
	}

	return r
}


var loginCookie string

var testEmail = "test@gmail.com"
var testPassword = "123456"

var accessToken string
var refreshToken string

var articleID int


func TestIntDB(t *testing.T) {
config.InitDB()
}


func TestRegister(t *testing.T) {
	testRouter := SetupRouter()
	var registerForm mappers.RegisterForm
	registerForm.FullName = "testing"
	registerForm.Email = testEmail
	registerForm.Password = testPassword
	data, _ := json.Marshal(registerForm)
	req, err := http.NewRequest("POST", "/v1/user/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}


func main() {
	r := SetupRouter()
	config.InitDB()
	r.Run()
}