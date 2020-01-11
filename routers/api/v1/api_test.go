package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/haodiaodemingzi/cloudfeet/models"
	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/auth"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/pac"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/proxy"

	"net/http"
	"net/http/httptest"
	"testing"
)

//var token string
var router *gin.Engine

func init(){
	settings.Setup()
	models.Setup()

	gin.SetMode("release")
	router = gin.Default()
	// map url
	router.GET(settings.Config.URL.ProxyInfo, proxy.GetProxy)
	router.POST(settings.Config.URL.AuthToken, auth.GenToken)
	router.GET(settings.Config.URL.PullDomains, pac.PullDomains)
	router.POST(settings.Config.URL.UploadDomains, pac.UploadDomains)
	router.PUT(settings.Config.URL.UpdateDomains, pac.UpdateDomains)
	router.POST(settings.Config.URL.UploadDNSFile, pac.UploadDomainFile)
	router.GET(settings.Config.URL.PacConfig, pac.DownloadBoxConfig)
	router.GET(settings.Config.URL.InitScript, pac.DownloadBoxScript)
}


func performRequest(r http.Handler, method, path string, reader io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, reader)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func prepareArgs(params map[string]interface{}) io.Reader{
	authInfo, _ := json.Marshal(params)
	reader := bytes.NewReader(authInfo)
	return reader
}


func TestGetAuthToken(t *testing.T) {
	hashmap := map[string]interface{}{
		"username": "cloudfeet",
		"password": "b44cf893756edd7d6a2245eb116341f8",
	}
	args := prepareArgs(hashmap)

	w := performRequest(router, "POST", settings.Config.URL.AuthToken, args)
	assert.Equal(t, http.StatusOK, w.Code)
	//var response map[string]string
	//_ = json.Unmarshal([]byte(w.Body.String()), &response)
	//fmt.Printf("resp : %+v", response)
}

func TestPullDomains(t *testing.T) {
	fmt.Println("hello domains")
	hashmap := map[string]interface{}{}
	args := prepareArgs(hashmap)

	w := performRequest(router, "GET", settings.Config.URL.PullDomains, args)
	fmt.Printf("body resp %+v", w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	//var response map[string]string
	//_ = json.Unmarshal([]byte(w.Body.String()), &response)
	//fmt.Printf("resp : %+v", response)
}

func TestUploadDomains(t *testing.T) {
	fmt.Println("hello domains")
	hashmap := map[string]interface{}{
		"source": "test",
		"domains": "cloudfeet-test007.com,cloudfeet-test008.com",
	}
	args := prepareArgs(hashmap)

	w := performRequest(router, "POST", settings.Config.URL.UploadDomains, args)
	fmt.Printf("body resp %+v", w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	//var response map[string]string
	//_ = json.Unmarshal([]byte(w.Body.String()), &response)
	//fmt.Printf("resp : %+v", response)
}

func TestUpdateDomains(t *testing.T) {
	domainResults := map[string]string{
		"cloudfeet-test001.com": "1",
		"cloudfeet-test002.com": "2",
	}
	hashmap := map[string]interface{}{
		"source": "test",
		"domains": domainResults,
	}
	args := prepareArgs(hashmap)

	w := performRequest(router, "PUT", settings.Config.URL.UpdateDomains, args)
	assert.Equal(t, http.StatusOK, w.Code)
	//var response map[string]string
	//_ = json.Unmarshal([]byte(w.Body.String()), &response)
	//fmt.Printf("resp : %+v", response)
}

func TestGetPacConfig(t *testing.T) {
	hashmap := map[string]interface{}{}
	args := prepareArgs(hashmap)

	w := performRequest(router, "GET", settings.Config.URL.PacConfig, args)
	fmt.Printf("body resp %+v", w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	//var response map[string]string
	//_ = json.Unmarshal([]byte(w.Body.String()), &response)
	//fmt.Printf("resp : %+v", response)
}

func TestGetInitScript(t *testing.T) {
	hashmap := map[string]interface{}{}
	args := prepareArgs(hashmap)

	w := performRequest(router, "GET", settings.Config.URL.InitScript, args)
	fmt.Printf("body resp %+v", w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	//var response map[string]string
	//_ = json.Unmarshal([]byte(w.Body.String()), &response)
	//fmt.Printf("resp : %+v", response)
}

func TestGetProxy(t *testing.T) {
	hashmap := map[string]interface{}{}
	args := prepareArgs(hashmap)

	w := performRequest(router, "GET", settings.Config.URL.ProxyInfo, args)
	fmt.Printf("body resp %+v", w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	//var response map[string]string
	//_ = json.Unmarshal([]byte(w.Body.String()), &response)
	//fmt.Printf("resp : %+v", response)
}
