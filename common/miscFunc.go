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
	var uinfo model.User
	if err!=nil{
		fmt.Println("届不到的cookie")
		return uinfo
	}
	cookie, err = DePwdCode(cookie)
	if err!=nil{
		fmt.Println("cookie 解密出错")
		return uinfo
	}
	err = json.Unmarshal([]byte(cookie), &uinfo)
	if err!=nil{
		fmt.Println("cookie 序列化出错")
		return uinfo
	}
	fmt.Println("一位朋友成功访问了权限网页:",uinfo.Name,uinfo.Mail)
	
	return uinfo
}