/**
* @author : yi.zhang
* @description : api 描述
* @date   : 2020-09-18 14:01
 */

package httpclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func NewClient(r *gin.Engine) *Client {
	return &Client{r: r}
}

type Client struct {
	r *gin.Engine
}

func (c *Client) Get(uri string) (int, []byte) {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	c.r.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	code := result.StatusCode
	body, _ := ioutil.ReadAll(result.Body)

	return code, body
}

func (c *Client) Post(uri string, data interface{}) (int, []byte) {
	bytParams, _ := json.Marshal(data)
	body := bytes.NewBuffer(bytParams)

	req, _ := http.NewRequest("POST", uri, body)
	w := httptest.NewRecorder()

	c.r.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	code := result.StatusCode
	bodyResult, _ := ioutil.ReadAll(result.Body)

	return code, bodyResult
}
