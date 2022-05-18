package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goapitemplate/pkg/helper"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ResponseReturn() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		// Before request
		c.Next()
		// After request

		responseBody := bodyLogWriter.body.String()

		var responseCode int
		var responseMsg string
		var responseData interface{}
		fmt.Println("Response 1: ", responseBody)

		if responseBody != "" {
			fmt.Println("Response 2: ", responseBody)

			response := helper.Response{}
			err := json.Unmarshal([]byte(responseBody), &response)
			if err == nil {
				responseCode = response.StatusCode
				responseMsg = response.Message
				responseData = response.Data
			}
		}

		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}

		//Log format
		accessLogMap := make(map[string]interface{})

		accessLogMap["TIME"] = time.Now()
		accessLogMap["METHOD"] = c.Request.Method
		accessLogMap["URL"] = c.Request.RequestURI
		accessLogMap["PROTO"] = c.Request.Proto
		accessLogMap["USER_AGENT"] = c.Request.UserAgent()
		accessLogMap["REFERER"] = c.Request.Referer()
		accessLogMap["POST_DATA"] = c.Request.PostForm.Encode()
		accessLogMap["CLIENT_IP"] = c.ClientIP()

		accessLogMap["STATUS"] = responseCode
		accessLogMap["MESSAGE"] = responseMsg
		accessLogMap["DATA"] = responseData

		//accessLogJson, _ := util.JsonEncode(accessLogMap)

		// if f, err := os.OpenFile(config.AppAccessLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
		// 	log.Println(err)
		// } else {
		// 	f.WriteString(accessLogJson + "\n")
		// }

		fmt.Println(json.Marshal(accessLogMap))
	}
}
