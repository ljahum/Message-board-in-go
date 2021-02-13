package Controller

import (
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"GoCode/common"
)

type TestController struct {

}

func (hc *TestController)Router(engine *gin.Engine)  {
	engine.GET("/test",hc.test)
	engine.GET("/to_ajax_test",hc.AjaxTest)
	engine.POST("/post_ajax",hc.PostAjax)

	engine.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("user_cookie")
		fmt.Println(cookie)

		if err != nil {
			fmt.Println("No set")

			value := []byte("this a cookie")
			cookieStr ,err:= common.EnPwdCode(value)
			if(err != nil){
				fmt.Println("加密失误")
			}
			c.SetCookie("user_cookie", cookieStr, 3600, "/", "localhost", false, true)
		} else {
			fmt.Println("already have a cookie")
			cookie, err = common.DePwdCode(cookie)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})



}

func (hc *TestController)test(ct *gin.Context)  {
	ct.JSON(200,gin.H{
		"message":	"Hello",
	})
}

func (hc *TestController) AjaxTest(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ajax_test.html", nil)
}

func (hc *TestController) PostAjax(ctx *gin.Context) {
	name := ctx.PostForm("name")
	age := ctx.PostForm("age")
	fmt.Println(name)
	fmt.Println(age)
	messgae_map := map[string]interface{}{
		"code":200,
		"msg":"提交成功",
	}
	ctx.JSON(http.StatusOK,messgae_map)
}