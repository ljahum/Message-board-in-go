package Controller

import (
	"GoCode/common"
	"GoCode/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) Router(ur *gin.Engine) {

	ur.GET("/login", uc.loginView)
	ur.POST("/login", uc.checkLogin)

	ur.GET("/regedit", uc.register)
	ur.POST("/regedit", uc.checkRegister)

	ur.GET("/logout", uc.logout) // 访问这个来清空cookie

}

//POST路由

func (uc *UserController) checkLogin(c *gin.Context) { //post

	var feedback string
	check := "false"

	cookie, _ := c.Cookie("user_cookie")
	// 从表单获取信息
	Email := c.PostForm("Email")
	passwd := c.PostForm("passwd")
	if cookie != "" {
		feedback = "宁已经登录🌶!"
		check = "true"
	} else if Email == "" || passwd == "" {
		feedback = "用户名或密码错误"
	} else {

		// 按name读用户信息

		fmt.Println("一位朋友登录了:", Email, passwd)
		db := common.GetDB()
		sqlStr := "select id,name, passwd ,mail from user_tab where mail=?"

		//check
		var u model.User
		err := db.QueryRow(sqlStr, Email).Scan(&u.Id, &u.Name, &u.Passwd, &u.Mail)
		if err != nil { // 没有查询到该用户
			feedback = "用户名或密码错误"
		} else if u.Mail == Email && u.Passwd == passwd {
			fmt.Printf("name:%s passwd:%s mail:%s\n", u.Name, u.Passwd, u.Mail)
			feedback = "登录成功"
			check = "true"
			setUserCookie(u, c)

		} else {
			feedback = "用户名或密码错误"
		}
	}
	common.SendJsonBack(feedback, check, c)

}

func (uc UserController) checkRegister(c *gin.Context) {
	var feedback string
	var check string

	cookie, _ := c.Cookie("user_cookie")
	name := c.PostForm("username")
	passwd := c.PostForm("passwd")
	mail := c.PostForm("mail")

	if cookie != "" {
		feedback = "宁已经登录🌶!"
		check = "true"
	} else if name == "" || passwd == "" || mail == "" {
		feedback = "用户名或密码错误"
	} else {

		db := common.GetDB()
		sqlStr := "select * from user_tab where mail=?"

		var u model.User
		err := db.QueryRow(sqlStr, mail).Scan(&u.Id, &u.Name, &u.Passwd, &u.Mail)
		if err == nil {
			feedback = "这个邮箱已被注册🌶！！！"
		} else {
			sqlStr := "insert into user_tab(name,passwd,mail) values (?,?,?)"
			ret, err := db.Exec(sqlStr, name, passwd, mail)
			if err != nil {
				fmt.Printf("insert failed, err:%v\n", err)
				return
			}
			theID, err := ret.LastInsertId() // 新插入数据的id
			if err != nil {
				fmt.Printf("get lastinsert ID failed, err:%v\n", err)
				return
			}
			fmt.Printf("朋友 %s %s 注册了本站第 %d 个账号\n", name, mail, theID)
			user := model.User{
				Id:     theID,
				Name:   name,
				Passwd: passwd,
				Mail:   mail,
			}
			feedback = "注册成功"
			check = "true"
			setUserCookie(user, c)
		}
	}
	common.SendJsonBack(feedback, check, c)

}

// func

func setUserCookie(u model.User, c *gin.Context) {
	value, err := json.Marshal(u)
	if err != nil {
		fmt.Println("json序列化出错")
	}
	cookieStr, err := common.EnPwdCode(value)

	c.SetCookie("user_cookie", cookieStr, 3600, "/", common.Main_domain, false, true)
}

// GET 路由

func (uc UserController) register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)

}

func (uc *UserController) loginView(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)

}

//logout
func (uc *UserController) logout(c *gin.Context) {
	c.SetCookie("user_cookie", "nil", -1, "/", common.Main_domain, false, true)
	c.HTML(http.StatusOK, "login.html", nil)

}
