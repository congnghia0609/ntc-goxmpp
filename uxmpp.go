package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mattn/go-xmpp"
)

type umessage struct {
	to, body string
}

func (m umessage) To() string {
	return m.to
}

func (m umessage) Body() string {
	return m.body
}

type uxmpp struct {
	opt     xmpp.Options
	client  *xmpp.Client
	msgChan chan string
}

func genUUID() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

var server = "message.example.com:80"
var username = "user@example.com"
var password = "password"
var status = "xa"
var statusMessage = "I online"
var notls = true
var debug = false
var session = true
var insecure = true
var startTLS = false
var resource = fmt.Sprintf("GoClient_%s", genUUID())

func (ux uxmpp) Send(m umessage) (int, error) {
	return ux.client.Send(xmpp.Chat{Remote: m.to, Type: "chat", Text: m.body})
}

func (ux uxmpp) SendTo(toUser string, msg string) (int, error) {
	return ux.client.Send(xmpp.Chat{Remote: toUser, Type: "chat", Text: msg})
}

func serverName(host string) string {
	return strings.Split(host, ":")[0]
}

func (ux uxmpp) IsConnect() bool {
	return ux.client != nil
}

func (ux uxmpp) PingServer(seconds time.Duration) {
	// ping Server
	go func() {
		i := uint64(1)
		for _ = range time.Tick(seconds * time.Second) {
			fmt.Println("====== Ping server time:", i)
			ux.client.PingC2S("", "")
			i++
		}
	}()
}

func (ux uxmpp) Listen() {
	// Consumer thread receive messages PoW Result.
	go func() {
		for {
			chat, err := ux.client.Recv()
			if err != nil {
				log.Fatal(err)
				//fmt.Println(err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				fmt.Println(v.Remote, v.Text)
				// put message to channel
				ux.msgChan <- v.Text
				// // string json to map
				// in := []byte(v.Text)
				// var mapData map[string]interface{}
				// json.Unmarshal(in, &mapData)
				// // mapData["count"] = 1
				// // map to string json and print
				// out, _ := json.Marshal(mapData)
				// fmt.Println(string(out))
				// // get data
				// type_xmpp := mapData["type_xmpp"]
				// type_msg := mapData["type_msg"]
				// blockId := mapData["blockId"]
				// blockHeader := mapData["blockHeader"]
				// seedHash := mapData["seedHash"]
				// difficulty := mapData["difficulty"]
				// mixHash := mapData["mixHash"]
				// nonce := mapData["nonce"]
				// fmt.Println("type_xmpp:", type_xmpp, ", type_msg:", type_msg, ", blockId:", blockId, ", blockHeader:", blockHeader,
				// 	", seedHash:", seedHash, ", difficulty:", difficulty, ", mixHash:", mixHash, ", nonce:", nonce)
				// fmt.Println("")
			case xmpp.Presence:
				fmt.Println("xmpp.Presence:", v.From, v.Show)
			}
		}
	}()
}

func (ux uxmpp) CreateDataJsonPoW(blockId uint64, blockHeader string, seedHash string, difficulty string) string {
	rs := ""
	rs = fmt.Sprintf("{\"type_xmpp\":\"XMPP_Topic_Disc_TaskMn\",\"type_msg\":\"DTM_Task_PoW\",\"blockId\":%d,\"blockHeader\":\"%s\",\"seedHash\":\"%s\",\"difficulty\":\"%s\"}", blockId, blockHeader, seedHash, difficulty)
	return rs
}

// InitXMPP function.
func InitXMPP() (uxmpp, error) {
	var ux uxmpp
	var err error
	if !notls {
		xmpp.DefaultConfig = tls.Config{
			ServerName:         serverName(server),
			InsecureSkipVerify: false,
		}
	} else {
		xmpp.DefaultConfig = tls.Config{
			ServerName:         serverName(server),
			InsecureSkipVerify: true,
		}
	}
	fmt.Println("=================== NewClient { username:", username, ", password:", password, "}")

	ux.opt = xmpp.Options{
		Host:                         server,
		User:                         username,
		Password:                     password,
		Resource:                     resource,
		NoTLS:                        notls,
		Debug:                        debug,
		Session:                      session,
		Status:                       status,
		StatusMessage:                statusMessage,
		InsecureAllowUnencryptedAuth: insecure,
		StartTLS:                     startTLS,
	}
	ux.client, err = ux.opt.NewClient()
	if err != nil {
		fmt.Println("[xxxxxx] ERROR:", err)
		// log.Fatal(err)
	}
	ux.msgChan = make(chan string)
	return ux, err
}

func makeDataPoW(blockId uint64, blockHeader string, seedHash string, difficulty string) string {
	rs := ""
	rs = fmt.Sprintf("{\"type_xmpp\":\"XMPP_Topic_Disc_TaskMn\",\"type_msg\":\"DTM_Task_PoW\",\"blockId\":%d,\"blockHeader\":\"%s\",\"seedHash\":\"%s\",\"difficulty\":\"%s\"}", blockId, blockHeader, seedHash, difficulty)
	return rs
}

func sendMsg(cli *xmpp.Client, toUser string, msg string) (n int, err error) {
	return cli.Send(xmpp.Chat{Remote: toUser, Type: "chat", Text: msg})
}
