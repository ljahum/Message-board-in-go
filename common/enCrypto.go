package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"unsafe"
)

//高级加密标准(Adevanced Encryption Standard,AES)
//双向加密  在目前认知范围内是不可能被破解的

//16或者24或者32字符串的话 分表对应不同的加密标准 AES-128  AES-192 AES-256加密方法
var PwdKey = []byte("KDJDKJJFJ*LKJSD)")  //这个Pwdkey是一定不能泄露的哈！很重要！

//PKCS7 填充模式

func PKCS7Padding(ciphertext []byte,blockSize int) []byte {

	padding := blockSize-len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个,然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(ciphertext,padtext...)
}

//PKCS7填充的反向操作，删除填充的字符串
func PKCS7UnPadding(origData []byte)([]byte,error){
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil,errors.New("加密字符串错误！")
	}else{
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文 [:8]截取切片 这个在之前是学过的哈
		return origData[:(length-unpadding)],nil
	}
}

//aes加密操作
func AesEcrypt(origData []byte,key []byte) ([]byte,error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	//采用AES加密方法中的CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block,key[:blockSize])
	crypted := make([]byte,len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted,origData)
	return crypted,nil
}

//aes解密操作
func AesDeCrypt(cypted []byte,key []byte)([]byte,error){
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//采用AES加密方法中的CBC加密模式 创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte,len(cypted))
	//这个函数还可以用来解密
	blockMode.CryptBlocks(origData,cypted)
	//去除填充字符串
	origData,err = PKCS7UnPadding(origData)
	if err != nil {
		return nil,err
	}
	return origData,err
}

//加密base64
func EnPwdCode(pwdByte []byte)(string,error){
	//pwdByte := str2bytes(pwd)
	result, err := AesEcrypt(pwdByte, PwdKey)
	if err != nil {
		return "",err
	}
	return base64.StdEncoding.EncodeToString(result),err
}

//解密base64
func DePwdCode(pwd string)(string,error){
	//解密base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return "Cookie error",err
	}
	//执行aes解密
	pwdByte, err = AesDeCrypt(pwdByte, PwdKey)
	pwd = bytes2str(pwdByte)

	return pwd,err
}
//======================================
func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
