package main
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	. "encoding/pem"
	"fmt"
	"os"
)
func GenerateKey(bits int) error{
	// 1.使用rsa中的GenerateKey方法生成私钥
	privatekey, err :=rsa.GenerateKey(rand.Reader,bits)
	if err !=nil {
		panic(err)
	}
	// 2.通过x509标准将得到rsa私钥序列化为ASN.1的DER编码字符串
	privStream := x509.MarshalPKCS1PrivateKey(privatekey)
	// 3.将编码字符串设置到pem格式快中
	block := Block{
		Type:"RSA Private Key",
		Bytes:privStream,
	}
	// 4.通过pem将设置好的数据进行编码，并写入磁盘文件中
	PrivFile, err := os.Create("PrivateKey.pem")
	if err !=nil {
		panic(err)
	}
	defer PrivFile.Close()
	err = Encode(PrivFile, &block)
	if err !=nil {
		panic(err)
	}
	//return nil
	// 1.从得到的私钥对象中将公钥信息取出
	pubkey := privatekey.PublicKey
	// 2.通过x509标准将得到rsa公钥序列化为字符串
	pubStream,err := x509.MarshalPKIXPublicKey(&pubkey)
	if err != nil {
		panic(err)
	}
	// 3.将编码字符串设置到pem格式快中
	block = Block{
		Type:"RSA Public Key",
		Bytes:pubStream,
	}
	// 4.通过pem将设置好的数据进行编码，并写入磁盘文件中
	PubFile, err := os.Create("Public.pem")
	if err !=nil{
		panic(err)
	}
	err = Encode(PubFile,&block)
	if err != nil {
		panic(err)
	}
	PubFile.Close()
	return nil
}
// src待加密数据，filename公/私钥文件路径
func RsaEncryptPubkey(src []byte,filename string) []byte {
	// 1.将公钥文件读出，得到使用pem编码的字符串
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
	block, _ := Decode(recvBuf)
	// 3.使用x509将编码之后的公钥解析
	pubInter, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	pubkey := pubInter.(*rsa.PublicKey)
	fmt.Println("pubkey is:",pubkey)
	// 4.使用解析后的公钥进行RSA数据加密
	message ,err :=rsa.EncryptPKCS1v15(rand.Reader, pubkey, src)
	if err!=nil{
		panic(err)
	}
	fmt.Println("密文是：",string(message))
	return message
}
func RsaDecryptSecckey(src []byte, filename string) []byte {
	// 1.将私钥文件读出，得到使用pem编码的字符串
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	// 1.1 得到文件属性信息
	info , _ :=file.Stat()
	recbuf := make([]byte, info.Size())
	file.Read(recbuf)
	// 2.将得到的字符串解码
	block, _ := Decode(recbuf)
	// 3.使用x509将编码之后的私钥解析
	privkey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil{
		panic(err)
	}
	fmt.Println("private key is ：",privkey)
	// 4.使用解析后的私钥进行RSA数据解密
	message, err :=rsa.DecryptPKCS1v15(rand.Reader,privkey,src)
		if err != nil{
		panic(err)
	}
	fmt.Println("明文是：",string(message))
	return message
}
func TestGenKey()  {
	err :=GenerateKey(4096)
	fmt.Println("错误信息：",err)
	// 加密
	src :=[]byte( "少壮不努力，hdshfud4561313￥%……&*（（&&&^&")
	str := RsaEncryptPubkey(src,"Public.pem")
	RsaDecryptSecckey(str,"PrivateKey.pem")

}
func main()  {
	TestGenKey()
}