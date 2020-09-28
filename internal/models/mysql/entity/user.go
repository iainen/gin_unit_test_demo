/**
* @author : yi.zhang
* @description : entity 描述
* @date   : 2020-08-17 18:25
 */

package entity

type User struct {
	ID       int    `mysql:"column:id;primary_key;size:36;auto_increment" json:"id"`
	UserName string `mysql:"column:user_name;size:64;index;default:'';not null;" json:"name"` // 用户名
	Password string `mysql:"column:password;size:40;default:'';not null;" json:"password"`    // 密码(sha1(md5(明文))加密)
	Email    string `mysql:"column:email;size:255;index;" json:"email"`                       // 邮箱
}

func (User) TableName() string {
	return "user"
}
