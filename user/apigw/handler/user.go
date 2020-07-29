package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"user-client/comm"
	"user-client/proto"

	"github.com/gin-gonic/gin"
)

var (
	AccountServiceClient proto.AccountService
)

func Signin(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	if username == "" || password == "" {
		return
	}

	// 通过AccountService服务的client，调用 AccountRegister 方法
	res, err := AccountServiceClient.Signin(context.TODO(), &proto.ReqSignin{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println(err)
		return
	}

	c.SetCookie("username", username, 86400*1000000000, "/", "119.23.229.247", false, true)
	c.SetCookie("uid", strconv.FormatInt(res.Uid, 10), 86400*1000000000, "/", "119.23.229.247", false, true)
	c.SetCookie("token", res.Token, 86400*1000000000, "/", "119.23.229.247", false, true)

	//将token放进浏览器的cookie里，另外也放到redis里来进行对token的验证
	//登录成功重定向到首页
	resp := comm.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://119.23.229.247:3000/user/home",
			Username: username,
			Token:    res.Token,
		},
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
}

func SignUp(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	//TODO 用户名长度和密码长度验证

	res, err := AccountServiceClient.Signup(context.TODO(), &proto.ReqSignup{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println(err)
		return
	}
	if res.Code == int64(200) {
		resp := comm.RespMsg{
			Code: 200,
			Msg:  "OK",
		}
		c.Data(http.StatusOK, "application/json", resp.JSONBytes())
		return
	} else {
		resp := comm.RespMsg{
			Code: 0,
		}
		c.Data(http.StatusOK, "application/json", resp.JSONBytes())
		return
	}
}

func UserInfo(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		log.Fatal("Failed to get uid, err: %s", err)
	}

	uId, err := strconv.Atoi(uid)
	if err != nil {
		log.Fatal("Failed to get uid, err: %s", err)
	}
	res, err := AccountServiceClient.UserInfo(context.TODO(), &proto.ReqUserInfo{
		Uid: int64(uId),
	})
	if err != nil {
		log.Println(err)
	}

	resp := comm.RespMsg{
		Code: 200,
		Msg:  "OK",
		Data: res.Username,
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
	return
}

func SelectUserFileAll(c *gin.Context) {
	id, err := c.Cookie("uid")
	if err != nil {
		log.Println(err)
	}
	uid, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}

	res, err := AccountServiceClient.SelectUserFileAll(context.TODO(), &proto.ReqGetUserFile{
		Uid: int64(uid),
	})
	if err != nil {
		log.Println(err)
	}
	var list []*UserFile
	for i := 0; i < len(res.Userfile); i++ {
		info := messageTouserfile(res.Userfile[i])
		list = append(list, info)
	}

	resp := comm.RespMsg{
		Code: 200,
		Msg:  "success",
		Data: list,
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
	return
}

func Logout(c *gin.Context) {
	c.SetCookie("username", "null", -1, "/", "119.23.229.247", false, true)
	c.SetCookie("uid", "null", -1, "/", "119.23.229.247", false, true)
	c.SetCookie("token", "null", -1, "/", "119.23.229.247", false, true)

	resp := comm.RespMsg{
		Code: 200,
		Msg:  "OK",
		Data: struct {
			Code int
		}{
			Code: 200,
		},
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
}

//修改文件名字
func RenameFile(c *gin.Context) {
	//获取post参数
	filehash := c.Request.FormValue("filehash")
	filename := c.Request.FormValue("filename")
	userid, err := c.Cookie("uid")
	if err != nil {
		log.Println(err)
		return
	}
	uid, err := strconv.Atoi(userid)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := AccountServiceClient.RenameFile(context.TODO(), &proto.ReqReName{
		Filehash: filehash,
		Filename: filename,
		Uid:      int64(uid),
	})

	if res.Ok {
		resp := comm.RespMsg{
			Code: 200,
			Msg:  "success",
		}
		c.Data(http.StatusOK, "application/json", resp.JSONBytes())
	}
}
