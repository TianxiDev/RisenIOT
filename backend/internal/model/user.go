package model

// LoginUserRes 登录返回
type LoginUserRes struct {
	Id           uint64 `orm:"id,primary"       json:"id"`           // 数据库自增主键
	UserName     string `orm:"user_name,unique" json:"userName"`     // 用户名
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	UserPassword string `orm:"user_password"    json:"userPassword"` // 登录密码; rsa 加密
	UserSalt     string `orm:"user_salt"        json:"userSalt"`     // 加密盐
	UserStatus   uint   `orm:"user_status"      json:"userStatus"`   // 用户状态; 0:禁用, 1:正常, 2:未验证
	IsAdmin      int    `orm:"is_admin"         json:"isAdmin"`      // 是否后台管理员; 1:是 ,0:否
	Avatar       string `orm:"avatar"           json:"avatar"`       // 头像
}
