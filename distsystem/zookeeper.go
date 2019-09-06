package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)

	}

	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}

	println("lock success,doyour business logic")
	time.Sleep(time.Second * 10)

	l.Unlock()
	println("unlock sucess,finish business logic")
}
