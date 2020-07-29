package handler

type User struct {
	ID       int64  `gorm:"column:id"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Salt     string `gorm:"cloumn:salt"`
}

//与数据库交互的数据结构
type UserFile struct {
	ID       int64  `gorm:"column:id"`
	UserId   int64  `gorm:"column:user_id"`
	FileHash string `gorm:"column:file_hash"`
	FileName string `gorm:"column:file_name"`
	FileSize int64  `gorm:"column:file_size"`
}

type FileMeta struct {
	ID       int64
	FileHash string `gorm:"column:file_hash"`
	FileName string `gorm:"column:file_name"`
	FileSize int64  `gorm:"column:file_size"`
	Location string `gorm:"column:file_location"`
}
