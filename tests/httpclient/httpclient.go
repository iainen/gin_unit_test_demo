/**
* @author : yi.zhang
* @description : api 描述
* @date   : 2020-09-18 14:01
 */

package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	mysql.Init()
// 	router.Init()
// }

// var r = router.Get()
func NewClient(r *gin.Engine) *Client {
	return &Client{r: r}
}

type Client struct {
	r *gin.Engine
}

func (c *Client) Get(uri string) (int, []byte) {
	// 构造get请求

	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	c.r.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	code := result.StatusCode
	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return code, body
}

func (c *Client) Post(uri string, data interface{}) (int, []byte) {
	fmt.Println(data)
	bytParams, _ := json.Marshal(data)
	body := bytes.NewBuffer(bytParams)
	fmt.Println(string(bytParams))
	req, _ := http.NewRequest("POST", uri, body)
	w := httptest.NewRecorder()
	c.r.ServeHTTP(w, req)
	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	code := result.StatusCode
	bodyResult, _ := ioutil.ReadAll(result.Body)
	return code, bodyResult
}
