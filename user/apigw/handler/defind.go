package handler

import (
	"user-client/comm"
	"user-client/proto"
)

type UserFile struct {
	ID       int64  `gorm:"column:id"`
	UserId   int64  `gorm:"column:user_id"`
	FileHash string `gorm:"column:file_hash"`
	FileName string `gorm:"column:file_name"`
	FileSize string `gorm:"column:file_size"`
}

func messageTouserfile(list *proto.UserFile) *UserFile {
	userfile := &UserFile{}
	userfile.ID = list.Id
	userfile.UserId = list.Userid
	userfile.FileHash = list.Filehash
	userfile.FileName = list.Filename
	userfile.FileSize = comm.FormatFileSize(list.Filesize)
	return userfile
}
