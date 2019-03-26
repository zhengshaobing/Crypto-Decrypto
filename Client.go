package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net"
	"os"
)

// 发送消息，并对数据进行签名

func signature(src string, filename string) ([]byte) {
	// 打开私钥文件
	privatekey,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	// 读取私钥文件信息
	info, _ := privatekey.Stat()
	recbuf := make([]byte,info.Size())
	privatekey.Read(recbuf)
	// 解码私钥文件
	block, _ :=pem.Decode(recbuf)
	Privkey,_ :=x509.ParsePKCS1PrivateKey(block.Bytes)
	// 获取签名信息摘要
	h :=md5.New()
	h.Write([]byte(src))
	hashed :=h.Sum(nil)
	// 进行签名
	opts := rsa.PSSOptions{rsa.PSSSaltLengthAuto,crypto.MD5}
	sig,_ := rsa.SignPSS(rand.Reader,Privkey,crypto.MD5,hashed,&opts)
	return sig
}
func SendTcp(src []byte)  {
	// 建立tcp链接
	conn,_ :=net.ResolveTCPAddr("tcp4","222.28.46.98:0001")
	n,_ :=net.DialTCP("tcp",nil,conn)
	// 发送数据
	n.Write(src)
	//n.Write(len)
}
func main(){
	// 获的签名结果
	sg :=signature("zhengshaobing","PrivateKey.pem")
	// 发送数据
	file, _ :=os.Open("duanzi.json")
	info , _ :=file.Stat()
	recbuf := make([]byte, info.Size())
	file.Read(recbuf)
	var data = make([]byte,len(recbuf)+len(sg))
	copy(data[:len(recbuf)],[]byte(recbuf))
	copy(data[len(recbuf):],sg)
	// 发送的数据为数据+签名结果
	//SendTcp(data,[]byte{byte(len(recbuf))})
	SendTcp(data)
	}