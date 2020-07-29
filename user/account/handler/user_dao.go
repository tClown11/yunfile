package handler

import (
	"fmt"
	"log"
	"time"
	"user-server/comm"
	"user-server/proto"
)

func updateFileMetaName(f *FileMeta) (bool, error) {
	db := comm.GetDB()

	err := db.Table("filemeta").Update("file_name", f.FileName).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func getFileMeta(filehash string) (*FileMeta, error) {
	db := comm.GetDB()

	f := &FileMeta{}
	err := db.Table("filemeta").Where("file_hash=?", filehash).Find(f).Error
	return f, err
}

func updateFileName(u *UserFile) (bool, error) {
	db := comm.GetDB()
	//警告:当使用struct更新时，FORM将仅更新具有非空值的字段
	err := db.Table("userfile").Where("ID = ?", u.ID).Update("file_name", u.FileName).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func getUserFile(filehash string, uid int64) (*UserFile, error) {
	db := comm.GetDB()
	f := &UserFile{}
	err := db.Table("userfile").Where("file_hash=? AND user_id=?", filehash, uid).Find(&f).Error
	return f, err
}

func selectAll(uid int64) (f []*proto.UserFile) {
	db := comm.GetDB()

	err := db.Table("userfile").Where("user_id=?", uid).Limit(15).Find(&f).Error
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func userSignUp(username, password string) bool {
	db := comm.GetDB()

	salt := comm.GetSalt()
	pw := comm.Sha1([]byte(password + salt))

	//INSERT INTO users(username, password, salt) values(username, pw, salt)
	//db.Exec("INSERT INTO users (user_name, password, salt) value (?, ?, ?)", username, pws, salt)
	err := db.Create(&User{UserName: username, Password: pw, Salt: salt}).Error
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func usersignin(username, password string) (string, int64) {
	db := comm.GetDB()
	user := &User{}

	//SELECT * FROM users WHERE name = username limit 1;
	db.Where("user_name=?", username).First(&user)
	pw := comm.Sha1([]byte(password + user.Salt))
	if pw == user.Password && username == user.UserName {
		token := GetToken(username)
		return token, user.ID
	} else {
		return "", 0
	}
}

func userinfo(uid int64) *User {
	db := comm.GetDB()
	user := &User{}

	//SELECT * FROM users WHERE id = uid limit 1;
	db.Where("id=?", uid).First(&user)
	return user
}

func GetToken(username string) string {
	//32字符: md5(username + timestamp + token_salt)
	ts := fmt.Sprintf("%x", time.Now().Unix())
	return comm.MD5([]byte(username + ts + "_token_salt"))
}
