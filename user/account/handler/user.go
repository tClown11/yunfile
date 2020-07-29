package handler

import (
	"context"
	"log"
	"strconv"
	"user-server/comm"
	"user-server/proto"
)

type AccountService struct {
}

func (a *AccountService) Signin(c context.Context, req *proto.ReqSignin, res *proto.ResSignin) error {
	username := req.Username
	password := req.Password

	//信息校验
	//生成token 获取用户id
	token, uid := usersignin(username, password)
	if token == "" && uid == 0 {
		//res需要赋值
		res.Token = ""
		res.Uid = 0
		return nil
	}
	rdb := comm.NewRedisClient()
	err := rdb.Set(c, strconv.FormatInt(uid, 10), token, 24*60*60*1000*1000*1000).Err() //超时时间单位为纳秒
	if err != nil {
		log.Fatal(err)
		return nil
	}
	//res的message 赋值
	res.Token = token
	res.Uid = int64(uid)
	return nil
}

func (a *AccountService) Signup(c context.Context, req *proto.ReqSignup, res *proto.ResSignup) error {
	username := req.Username
	password := req.Password

	if userSignUp(username, password) {
		//message赋值
		res.Message = "success"
		res.Code = 200
		return nil
	} else {
		//message赋值
		res.Message = "注册失败"
		res.Code = 0
		return nil
	}
}

func (a *AccountService) SelectUserFileAll(c context.Context, req *proto.ReqGetUserFile, res *proto.ResGetUserFile) error {
	uid := req.Uid

	fileinfo := selectAll(uid)

	res.Userfile = fileinfo

	return nil
}

func (a *AccountService) RenameFile(c context.Context, req *proto.ReqReName, res *proto.ResReName) error {
	filehash := req.Filehash
	filename := req.Filename
	uid := req.Uid

	//获取文件表的数据与用户数据表的数据
	//修改用户表
	file, err := getUserFile(filehash, uid)
	if err != nil {
		log.Println(err)
		//message赋值
		res.Ok = false
		res.Message = "文件不存在"
		return nil
	}
	file.FileName = filename
	ok, err := updateFileName(file)
	if err != nil || !ok {
		log.Println(err)
		//message赋值
		res.Ok = false
		res.Message = "重命名失败"
		return nil
	}

	//修改文件表
	filemeta, err := getFileMeta(filehash)
	if err != nil {
		log.Println(err)
		//message赋值
		res.Ok = false
		res.Message = "文件表不存在文件"
		return nil
	}
	filemeta.FileHash = filehash
	ok, err = updateFileMetaName(filemeta)
	if err != nil || !ok {
		log.Println(err)
		//message赋值
		res.Ok = false
		res.Message = "重命名失败"
		return nil
	}
	if ok {
		//message赋值
		res.Ok = true
		res.Message = "重命名成功"
		return nil
	}
	return nil
}

func (a *AccountService) UserInfo(c context.Context, req *proto.ReqUserInfo, res *proto.ResUserInfo) error {
	uid := req.Uid
	user := userinfo(uid)
	//message
	userTomessage(user, res)
	return nil
}
