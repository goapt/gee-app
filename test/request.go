package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"

	"app/api/router"
)

func NewUrlEncodeRequest(path string, values *url.Values) *http.Request {
	req := NewRequest(path, values.Encode())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	return req
}

func NewJsonRequest(path string, data map[string]interface{}) *http.Request {
	body, _ := json.Marshal(data)
	req := NewRequest(path, string(body))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	return req
}

func NewXMLRequest(path, body string) *http.Request {
	req := NewRequest(path, body)
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	return req
}

func NewRequest(path, body string) *http.Request {
	buf := bytes.NewBufferString(body)
	req, _ := http.NewRequest("POST", path, buf)
	return req
}

func Run(t *testing.T, req *http.Request) *httptest.ResponseRecorder {
	r := router.Engine()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	return w
}

func GetJsonBody(resp *httptest.ResponseRecorder, path string) gjson.Result {
	body, _ := ioutil.ReadAll(resp.Body)
	return gjson.GetBytes(body, path)
}

func IsSuccess(t *testing.T, resp *httptest.ResponseRecorder) {
	assert.Equal(t, GetJsonBody(resp, "code").Int(), int64(gee.SuccessCode))
}

func IsFail(t *testing.T, resp *httptest.ResponseRecorder) {
	assert.NotEqual(t, GetJsonBody(resp, "code").Int(), int64(gee.SuccessCode))
}
