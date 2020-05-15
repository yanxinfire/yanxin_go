package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"yxtool/common"
	"yxtool/util"
)

var (
	help             bool
	encryptPasswd string
	unicodeStr    string
	chineseStr    string
)

func init()  {
	flag.BoolVar(&help, "h",false,"Help information")
	flag.StringVar(&encryptPasswd, "dp", "", "Decrypt encrypted passwd")
	flag.StringVar(&unicodeStr, "uni", "", "Turn unicode string into chinese")
	flag.StringVar(&chineseStr, "ch", "", "Turn chinese into unicode string")
	flag.Usage=usage
}

func main() {
	parseArguments()
}

func parseArguments() {
	flag.Parse()
	if flag.NFlag()==0{
		help=true
	}
	if help{
		flag.Usage()
		return
	}
	if encryptPasswd!=""{
		decryptPasswd(strings.TrimSpace(encryptPasswd))
	}
	if unicodeStr!=""{
		uni2Ch(strings.TrimSpace(unicodeStr))
	}
	if chineseStr!=""{
		ch2Uni(strings.TrimSpace(chineseStr))
	}
}

func decryptPasswd(encryptPasswd string) {
	//encryptPasswd = `1ba0746be69dc92a24e5ffb086d834a905c47110faf02b35698eec15bbb437b150fa2dd40bcd7efadc2248d3a8be9daea7185193d1c6fc998220dfaa4bb2328afe9e6997bf504adb54e436f8355e7e6d7d7d995fb25d18c6b1c52ccd69438dede329e86b19744a66dd1f2914abe0c2d1f8144c1f35d120199054e54395fe6fb6`
	encryptType := "RSA"

	encryptRes, err := hex.DecodeString(encryptPasswd)
	if err != nil {
		fmt.Println("hex.DecodeString fail: value ", encryptPasswd, " err: ", err.Error())
	}
	privateKey := common.GetPrivateKey()
	decryRes, err := util.RsaDecrypt(encryptRes, []byte(privateKey))
	if err != nil {
		fmt.Println("RsaDecrypt fail, encryptType ", encryptType)
		fmt.Println(errors.New("Passwd RsaDecrypt fail, err: " + err.Error()))
		return
	}

	fmt.Println(string(decryRes))
}

func uni2Ch(unicodeStr string)  {
	sUnicodev := strings.Split(unicodeStr, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		context += fmt.Sprintf("%c", temp)
	}
	fmt.Println(context)
}

func ch2Uni(chineseStr string)  {
	textQuoted := strconv.QuoteToASCII(chineseStr)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	fmt.Println(textUnquoted)
}

func usage() {
	_,_=fmt.Fprintf(os.Stderr, `yxtool version: v1.0.0
Usage: yxtool [-h] [-dp enpasswd]

Options:
`)
	flag.PrintDefaults()
}

