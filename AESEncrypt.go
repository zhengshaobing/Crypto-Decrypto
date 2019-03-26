package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func Read() []byte {
	data, err := ioutil.ReadFile("duanzi.json")
	if err != nil {
		panic(err)
	}
	return data
}
// 写入加密文本
func WriteWithIoutil(name,content string) {
    data :=  []byte(content)
    if ioutil.WriteFile(name,data,0644) == nil {
        fmt.Println("写入文件成功:")
        }
    }
func paddingText(src []byte, blocksize int) []byte {
	// 1. 求出最后一个分组所要填充的字节
	padding := blocksize - len(src)%blocksize
	// 2. 创建新的切片，并初始化，切片字节数应为padding，每个字节值为padding
	padTex := bytes.Repeat([]byte {byte(padding)},padding)
	// 3.将切片和原始数据进行链接
	newTex := append(src,padTex...)
	// 4.返回新的数据
	return newTex
}
// 删除末尾填充字节数
func unpaddingTex(src []byte) []byte {
	// 1.求出要处理字符串的长度
	len := len(src)
	// 2. 取出最后一个字节，并得到其整型数
	num := int(src[len-1])
	// 3.将末尾的num个字节删除
	newTex := src[:len-num]
	// 4.返回数据
	return newTex
}
// 使用AeS对称加密
func AESEncrypt(src, key []byte) []byte {
	// 1.创建并返回一个使用AES算法的cipher.block接口对象
	block ,err := aes.NewCipher(key)
	if err !=nil{
		panic(err)
	}
	// 2. 对最后一个明文分组进行数据填充
	src =  paddingText(src,block.BlockSize())
	// 3.创建一个密码分组为链接模式的，底层使用AES加密的BlockMode接口
	blockModel :=cipher.NewCBCEncrypter(block, key)
	// 4.数据加密
	blockModel.CryptBlocks(src,src)
	return src
}
// 数据解密
func AESDecrypt(src,key []byte) []byte {
	// 1.创建并返回一个使用AES算法的cipher.block接口对象
	block ,err := aes.NewCipher(key)
	if err !=nil{
		panic(err)
	}
	// 2.创建一个密码分组为链接模式的，底层使用AES解密的BlockMode接口
	blockModel := cipher.NewCBCDecrypter(block,key)
	// 3.数据解密
	blockModel.CryptBlocks(src,src)
	// 4.去除添加的数据量
	src = unpaddingTex(src)
	// 5.返回
	return src
}
func TestAes()  {
	fmt.Println("=======AES加解秘测试=========")
	src :=Read()
	fmt.Println("明文是：",string(src))
	key :=[]byte("12345678abcdefgh")
	str := AESEncrypt(src,key)
	name := "加密结果.txt"
	content := string(str)
	WriteWithIoutil(name,content)
	fmt.Println("加密后密文：",string(str))
	str1 := AESDecrypt(str,key)
	fmt.Println("解密后明文：",string(str1))
}
func main()  {
	TestAes()
}