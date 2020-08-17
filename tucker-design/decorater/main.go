package main

import (
	"fmt"

	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var receiveData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	//Send data
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipdata, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipdata))
}

type EncryptComponent struct {
	Key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptdata, err := cipher.Encrypt([]byte(data), self.Key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptdata))
}

type DecryptComponent struct {
	Key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.Key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipdata, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipdata))
}

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	receiveData = data
}

func main() {
	sender := &EncryptComponent{
		Key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("Hello world")
	fmt.Println(sentData)

	reciver := &UnzipComponent{
		com: &DecryptComponent{
			Key: "abcde",
			com: &ReadComponent{},
		},
	}

	reciver.Operator(sentData)
	fmt.Println(receiveData)
}
