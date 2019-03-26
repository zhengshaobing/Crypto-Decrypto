package main
import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"os"
)
func sign(src string, bits int)  {
	// 创建私钥
	privatekey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil{
		panic(err)
	}
	// 创建公钥
	pub := &privatekey.PublicKey
	// 签名
	// 1.取hash值
	h := md5.New()
	h.Write([]byte(src))
	hashed := h.Sum(nil)
	fmt.Println("hash摘要是：",hex.EncodeToString(hashed))
	// 2.数字签名
	opts := rsa.PSSOptions{rsa.PSSSaltLengthAuto,crypto.MD5}// 添加杂质
	sig, err :=rsa.SignPSS(rand.Reader,privatekey,crypto.MD5,hashed,&opts)
	if err !=nil{
		panic(err)
	}
	fmt.Println("签名的结果是：",string(sig))
	// 公钥验签
	error := rsa.VerifyPSS(pub,crypto.MD5,hashed,sig,&opts)
	if error == nil{
		fmt.Println("验证成功")
	}
}
func main()  {
	file, _ :=os.Open("duanzi.json")
	info , _ :=file.Stat()
	recbuf := make([]byte, info.Size())
	file.Read(recbuf)
	sign(string(recbuf),1024)
}