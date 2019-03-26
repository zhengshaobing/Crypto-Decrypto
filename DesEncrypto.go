package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"io/ioutil"
)
// 填充最后一个分组函数
// src--原始数据
// blocksize --分组长度
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
// 使用DES加密
func CryptoDEC(src, key []byte) []byte {
	// 1.创建并返回一个使用DES算法的cipher.Block接口
	block , err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2.对最后一个明文分组进行数据填充
	srt := paddingText(src,block.BlockSize())
	// 3.创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	iv := []byte("zhengzxc") 
	BlockMode := cipher.NewCBCEncrypter(block,iv)
	// 4.加密连续的数据快
	dst := make([]byte, len(srt))
	BlockMode.CryptBlocks(dst, srt)
	return dst
}
// 使用DES解密
func DecryptoDEC(src,key []byte) []byte {
	//1.创建并返回一个使用DES算法的cipher.Block接口
	block,err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2.创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	iv := []byte("zhengzxc") 
	BlockMode := cipher.NewCBCDecrypter(block,iv)
	// 3.数据块解密
	BlockMode.CryptBlocks(src, src)
	// 4.去除填充数据 
	newTex := unpaddingTex(src)
	// 5.返回文本byte类型
	return newTex
}
// 读取文件
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
func Test() {
	fmt.Println("=======使用DES解密=========")
	a := string(Read())
	src :=[]byte(a)
	fmt.Println("src" +string(src))
	key :=[]byte("12345678")
	str1 := CryptoDEC(src, key)
	name := "crypto.txt"
	content := string(str1)
	WriteWithIoutil(name,content)
	fmt.Println("str1:" + string(str1))
	str2 := DecryptoDEC(str1, key)
	fmt.Println("str2:" + string(str2))
}
 func main() {
 	Test()
}