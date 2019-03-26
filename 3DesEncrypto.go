package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"io/ioutil"
	"fmt"
)
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
// 使用3des加密
func Encrypt3DES(src, key[]byte) []byte{
	// 1.创建并返回一个使用3des算法的cipher.block接口
	block, err := des.NewTripleDESCipher(key)
	if err !=nil {
		panic(err)
	}
	// 2.对最后一个明文分组进行填充
	src = paddingText(src, block.BlockSize())
	// 3.创建一个密码分组为链接模式的，底层使用3DES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block,key[:block.BlockSize()])
	// 4.加密连续的数据块
	blockMode.CryptBlocks(src, src)
	// 5.返回结果
	return src
}
// 使用3DES解密
func Decrypt3DES(src, key []byte) []byte{
	// 1.创建并返回一个使用3des算法的cipher.block接口
	block , err := des.NewTripleDESCipher(key)
	if err != nil{
		panic(err)
	}
	// 2.创建一个密码分组为链接模式的，底层使用3DES解密的BlockMode接口
	blockModel := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	// 3.对数据进行解密
	blockModel.CryptBlocks(src, src)
	// 4.去除填充数据
	src = unpaddingTex(src)
	// 5.返回
	return src
}
func TDEStest()  {
	fmt.Println("=========3DES加解密操作==========")

	src :=Read()
	fmt.Println("src" + string(src))
	key := []byte("abcdefgh12345678!@#$%^&*")
	str1 := Encrypt3DES(src,key)
	fmt.Println("加密结果:",string(str1))
	name := "crypto.txt"
	content := string(str1)
	WriteWithIoutil(name,content)
	str2 := Decrypt3DES(str1,key)
	fmt.Println("解密结果:",string(str2))
}
func main(){
	TDEStest()
}