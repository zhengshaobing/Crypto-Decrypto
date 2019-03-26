package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"os"
)

func Receive() []byte {
	netListen,_ :=net.Listen("tcp","222.28.46.98:0001")
	 //延时关闭listen
	defer netListen.Close()
	// 循环监听端口
	for {
		conn,_ := netListen.Accept()
		// 设置接收数据缓存
		buf := make([]byte,2048)
		for{
			n,_ := conn.Read(buf)
			// 返回接收数据
			return buf[:n]
		}
	}
}
func VerifySig(src []byte,filename string)  {
	file,err := os.Open(filename)
	if err !=nil{
		panic(err)
	}
	// 1.1得到文件属性信息，通过属性信息对象得到文件大小
	info, err := file.Stat()
	if err!=nil{
		panic(err)
	}
	recvBuf := make([]byte, info.Size())
	file.Read(recvBuf)
	// 2.将得到的字符串解码
	block, _ := pem.Decode(recvBuf)
	// 3.使用x509将编码之后的公钥解析
	pubInter, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	pubkey := pubInter.(*rsa.PublicKey)
	h :=md5.New()
	h.Write([]byte("zhengshaobin"))
	hashed :=h.Sum(nil)
	err =rsa.VerifyPSS(pubkey,crypto.MD5,hashed,src,nil)
	if err !=nil{
		fmt.Println("数据可能被篡改，请谨慎使用")
	} else{
		fmt.Println("验签成功")
	}
}
func main()  {
	// 接收数据
	data:=Receive()
	// 拆分数据
	plaintxt := data[:674]
	fmt.Println("接受的明文是：",string(plaintxt))
	// 获取数字签名结果
	sig := data[674:]
	// 公钥验签
	VerifySig(sig,"Public.pem")
}