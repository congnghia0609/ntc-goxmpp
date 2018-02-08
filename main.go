package main

import (
	"fmt"
)

func main() {
	TCF{
		Try: func() {
			fmt.Println("I tried")
			Throw("Oh,...sh...")
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally...")
		},
	}.Do()

	// 1. Init directive
	// var ux uxmpp
	// var err error
	// ux, err = InitXMPP()
	// if err != nil {
	// 	log.Fatal(err)
	// 	//fmt.Println(err)
	// }
	// fmt.Println("=================== ux ===================")
	// fmt.Println(ux)
	// // Consumer thread receive messages PoW Result.
	// ux.Listen()
	// // ping Server
	// ux.PingServer(time.Duration(5))

	// 2. Init Singleton
	// var ux *uxmpp
	// ux = GetInstance()

	// // listening message channel
	// go func() {
	// 	var msg string
	// 	for msg = range ux.msgChan {
	// 		// string json to map
	// 		in := []byte(msg)
	// 		var mapData map[string]interface{}
	// 		json.Unmarshal(in, &mapData)
	// 		// mapData["count"] = 1
	// 		// map to string json and print
	// 		out, _ := json.Marshal(mapData)
	// 		fmt.Println(string(out))
	// 		// get data
	// 		type_xmpp := mapData["type_xmpp"]
	// 		type_msg := mapData["type_msg"]
	// 		blockId := mapData["blockId"]
	// 		blockHeader := mapData["blockHeader"]
	// 		seedHash := mapData["seedHash"]
	// 		difficulty := mapData["difficulty"]
	// 		mixHash := mapData["mixHash"]
	// 		nonce := mapData["nonce"]
	// 		fmt.Println("type_xmpp:", type_xmpp, ", type_msg:", type_msg, ", blockId:", blockId, ", blockHeader:", blockHeader,
	// 			", seedHash:", seedHash, ", difficulty:", difficulty, ", mixHash:", mixHash, ", nonce:", nonce)
	// 		fmt.Println("")
	// 	}
	// }()

	// // send to User.
	// toUser := "user@example.com"
	// // send message PoW.
	// blockId := uint64(10)
	// blockHeader := "0x04a4fcf765d61e99fc2a9c785f4505f32de74c38ec2d0d120b5c278d5659e087"
	// seedHash := "0x0000000000000000000000000000000000000000000000000000000000000000"
	// difficulty := "0x00007fe007fe007fe007fe007fe007fe007fe007fe007fe007fe007fe007fe00"
	// for i := 0; i < 10; i++ {
	// 	// msg := makeDataPoW(blockId, blockHeader, seedHash, difficulty)
	// 	// fmt.Println("Send:", toUser, "<----->", msg)
	// 	// ux.SendTo(toUser, msg)
	// 	body := ux.CreateDataJsonPoW(blockId, blockHeader, seedHash, difficulty)
	// 	fmt.Println("Send:", toUser, "<----->", body)
	// 	m := umessage{to: toUser, body: body}
	// 	ux.Send(m)
	// 	blockId++
	// 	time.Sleep(5 * time.Second)
	// }

	// fmt.Println("=========== XMPP-GO-Client is running...")
	// //swing current thread.
	// max32i := 1<<32 - 1
	// time.Sleep(time.Duration(max32i) * time.Second)

}
