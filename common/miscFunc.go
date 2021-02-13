package common

import (
	"GoCode/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendJsonBack(feedback string,check string,c *gin.Context)  {
	messageMap := map[string]interface{}{
		"msg": feedback,
		"check" : check,
	}
	c.JSON(http.StatusOK, messageMap)
}

func GetcookieStruct(c *gin.Context) model.User {
	cookie, err := c.Cookie("user_cookie")
	cookie, err = DePwdCode(cookie)
	var uinfo model.User
	err = json.Unmarshal([]byte(cookie), &uinfo)
	if err!=nil{
		fmt.Println("cookie 解密出错")
	}
	return uinfo
}