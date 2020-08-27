package authpkg

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"encoding/json"
	"chat/base"
)

type UserDao struct {
	pool *redis.Pool
}

var ConstDao *UserDao

func NewDao(pool *redis.Pool) (userdao *UserDao) {
	return &UserDao{
		pool : pool,
	}
}

func (this *UserDao) GetUserById(con redis.Conn, id int) (user *UserInfo, err error) {
	r1, err := redis.String(con.Do("HGET", "user", id))

	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOEXIST
		}
		return
	}

	user = &UserInfo{}
	err = json.Unmarshal([]byte(r1), user)
	if err != nil {
		fmt.Println("反序列化用户信息出错：", err)
		return
	}
	return
}

func (this *UserDao) CheckUserInfo(id int, pwd string) (user *UserInfo, err error) {
	c := this.pool.Get()
	user, err = this.GetUserById(c, id)
	if err != nil {
		return
	}

	if user.Upwd != pwd {
		err = ERROR_USER_PWD
		return
	}

	defer c.Close()
	return
}

func (this *UserDao) Register(user *base.UserInfo) (err error) {
	c := this.pool.Get()
	_, err = this.GetUserById(c, user.UserId)
	if err == nil {
		err = ERROR_USER_EXIST
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("序列化注册用户信息出现错误：", err)
		return
	}

	_, err = c.Do("HSET", "user", user.UserId, string(data))
	if err != nil {
		fmt.Println("注册用户，写入redis出错：", err)
		return
	}
	defer c.Close()
	return
}
