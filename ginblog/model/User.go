package model

import (
	"encoding/base64"
	"ginblog/utils/errormsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"lable:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" lable:"密码""`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role"validate:"required,gte=2" lable:"角色码"`
}

//查询用户是否存在

func CheckUser(name string) int {
	var users User
	SqlDb.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errormsg.ERROR_USERNAME_USED
	}
	return errormsg.SUCCESS
}

//添加用户

func CreateUser(data *User) int {
	err := SqlDb.Create(&data).Error
	if err != nil {
		return errormsg.ERROR //500
	}
	return errormsg.SUCCESS
}

//查询用户列表

func GetUsers(username string, pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int

	if username == "" {
		SqlDb.Select("id,username,role").Where("username LIKE ?", username+"%").Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize)
		return users, total
	}
	SqlDb.Select("id,username,role").Where("username LIKE ?", username+"%").Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize)
	//一般涉及到用户列表的都是要进行分页的
	//limit和offset传-1为不使用限制

	//我们要考虑到返回err的原因，查询错误，或者为空
	if err == gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}
func GetUser(id int) (User, int) {
	var user User
	err := SqlDb.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errormsg.ERROR
	}
	return user, errormsg.SUCCESS
}

//删除用户

func DeleteUser(id int) int {
	var user User
	//执行的软删除
	err := SqlDb.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

//编辑用户信息

func EditUser(id int, data *User) int {
	var user User
	//用map传参数，如果用struct为零的字段是不会进行更新的
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	SqlDb.Model(&user).Where("id = ?", id).Update(maps)
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

//密码加密,使用的是scrypt加盐算法，一般还是使用比较熟悉的加密算法，因为性能开销啥的还没考虑

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{'E', 'T', 'E', 'R', 'N', 'A', 'L'}
	pwd, err := scrypt.Key([]byte(password), salt, 1<<15, 2, 3, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpwd := base64.StdEncoding.EncodeToString(pwd)
	return fpwd
}

//钩子函数,这个方法名是约定执行的，gorm包的一个特性

func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

//登入验证
func CheckLogin(username string, password string) int {
	var user User
	SqlDb.Where("username = ? ", username).First(&user)
	if user.ID == 0 {
		return errormsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errormsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errormsg.ERROR_USER_NO_RIGHT
	}
	return errormsg.SUCCESS
}
