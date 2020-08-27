package main

import (
	"fmt"
	"net"
	"io"
	"chat/server/model"
)

func initUserDao() {
	authpkg.ConstDao = authpkg.NewDao(pool)

}

// 获取消息信息
func readMess(con net.Conn) {
	Processor := &Processor{
		Con : con,
	}
	err := Processor.MainProcess()
	if err != nil {
		if err == io.EOF {
			fmt.Println("没有新的消息")
		}
	}

	defer con.Close()
}

func main() {

	initPool("localhost:6379") // 初始化redis连接池
	initUserDao()

	// 连接池
	/*pool = &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 10 * time.Second,
		Dial: func () (redis.Conn, error) {
			return redis.Dial("tcp", "0.0.0.0:6379")
		},
	}*/

	// pool.Close() // 关闭连接池

	// c := pool.Get()

	/*c, err := redis.Dial("tcp", "0.0.0.0:6379")
	EchoErr(err)*/

	/*r, err := c.Do("MSET", "k1", "a", "k2", "b")
	EchoErr(err)
	fmt.Println(r)*/

	/*r2, err := redis.Strings(c.Do("MGET", "k1", "k2"))
	EchoErr(err)
	fmt.Println(r2)*/

	// defer c.Close()


	li, err := net.Listen("tcp", "0.0.0.0:8856")
	if err != nil {
		fmt.Println(err)
	}
	for {
		con, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go readMess(con)
	}
	defer li.Close()
}
