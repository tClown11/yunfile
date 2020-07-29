package handler

import (
	"user-server/proto"
)

func userTomessage(u *User, res *proto.ResUserInfo) {
	res.Id = u.ID
	res.Username = u.UserName
}

func userfileTomessage(u *UserFile, res *proto.UserFile) {
	res.Id = u.ID
	res.Userid = u.UserId
	res.Filehash = u.FileHash
	res.Filename = u.FileName
	res.Filesize = u.FileSize
}

func filemetaTomessage(u *FileMeta) (res *proto.MetaFile) {
	res.Id = u.ID
	res.Filehash = u.FileHash
	res.Filename = u.FileName
	res.Filesize = u.FileSize
	res.Location = u.Location
	return
}
