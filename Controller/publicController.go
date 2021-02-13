package Controller

import (
	"GoCode/common"
	"GoCode/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PubController struct {

}

func (pc PubController) Router(pr *gin.Engine)  {
	pr.Use(Validate2()) //假把意识用个中间件试一哈

	pr.GET("/",pc.indexView)
	pr.GET("/index", pc.toHome)
	pr.GET("/room", pc.room)
	pr.POST("/room",pc.getMsg)
}
//poster
func (pc PubController) getMsg(c *gin.Context)  {
	var feedback string
	var check string
	var uinfo model.User
	var msg model.Comment
	//get cookie
	uinfo = common.GetcookieStruct(c)
	fmt.Println(uinfo.Mail)
	//get msg
	Msg := c.PostForm("message")
	if len(Msg) > 128{
		feedback  = "消息太长了"
		check= "false"
	}else {
		msg.Mail = uinfo.Mail
		msg.Name = uinfo.Name
		msg.Content = Msg
		msg.Time = time.Now().Format("2006-01-02 15:04:05")
		db := common.GetDB()
		sqlStr := "INSERT INTO liuyan(name,content,time,mail )VALUES(?,?,?,?)"
		_, err := db.Exec(sqlStr, msg.Name, msg.Content,msg.Time,msg.Mail)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)

			feedback = "数据库插入出错"
			check = "false"
			return
		}else {
			feedback = "发表成功"
			check = "true"
		}

	}


	common.SendJsonBack(feedback,check,c)
}

//geter
func Validate2() gin.HandlerFunc{
	return func(c *gin.Context){
		var tmp model.User                                 //定义一个结构体变量
		tmp = common.GetcookieStruct(c)
		if checkCookie(tmp.Mail,tmp.Passwd) {
			c.Next()
		}else {
			c.SetCookie("user_cookie", "nil", -1, "/", "localhost", false, true)
			c.HTML(http.StatusOK,"index.html",nil)
			c.Abort()
		}
	}
}


func checkCookie(Email string,passwd string) bool {
	db := common.GetDB()
	sqlStr := "select id,name, passwd ,mail from user_tab where mail=?"
	var u model.User
	err := db.QueryRow(sqlStr,Email).Scan(&u.Id, &u.Name, &u.Passwd, &u.Mail)
	if err !=  nil{ // 没有查询到该用户
		return false
	}else if u.Mail == Email && u.Passwd == passwd{
		return true
	}else {
		return false
	}
}




func (pc PubController) indexView(c *gin.Context)  {
	c.HTML(200,"index_loged.html",nil)
}

func (pc PubController) toHome(c *gin.Context)  {
	c.Redirect(http.StatusMovedPermanently, "/")
}

//goland:noinspection SqlNoDataSourceInspection
func (pc PubController) room(c *gin.Context)  {
//读取留言
	db := common.GetDB()
	sqlStr := "select * from liuyan ORDER BY id DESC LIMIT 20"
	rows, err := db.Query(sqlStr)
	if err != nil {
		panic("fail to connect databse,err:")
	}
	defer rows.Close()
	var liuyanData []*model.Comment
	for rows.Next() {
		var com model.Comment
		err := rows.Scan(&com.Id, &com.Name, &com.Content,&com.Mail,&com.Time)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
		}
		liuyanData = append(liuyanData,&com)
	}
//读取个人信息
	var uinfo model.User
	uinfo = common.GetcookieStruct(c)

	c.HTML(http.StatusOK, "arr.html", gin.H{
		"title":  "Gin",
		//"stuArr": [1]*model.Comment{com1},
		"stuArr": liuyanData,
		"uinfo":uinfo,
	})
}
