package main

import (
	"fmt"
	"sync"
	"time"
)

var instance *uxmpp
var once sync.Once

// Init Singleton UXMPP
func GetInstance() *uxmpp {
	fmt.Println("=================== uxmpp.GetInstance ===================")
	once.Do(func() {
		var ux uxmpp
		var err error
		ux, err = InitXMPP()
		if err != nil {
			fmt.Println("[xxxxxx] ERROR:", err)
			//log.Fatal(err)
		}
		// fmt.Println(ux)
		// Consumer thread receive messages PoW Result.
		ux.Listen()
		// ping Server
		ux.PingServer(time.Duration(5))
		instance = &ux
	})
	return instance
}
