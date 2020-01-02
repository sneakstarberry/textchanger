package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// C://Users/gihon/OneDrive/SYU_CODE/go/src/readfile.txt
func main() {
	//file1을 읽는다.
	file1, err := os.Open("./readfile1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	//file2를 읽는다.
	file2, err := os.Open("./readfile2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	fo, err := os.Create("./createfile1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()
	//읽은 파일 배열로 담아두기
	//file1
	b1, err := ioutil.ReadAll(file1)
	c1 := string(b1)
	d1 := strings.Split(c1, "\n")
	//file2
	b2, err := ioutil.ReadAll(file2)
	c2 := string(b2)
	d2 := strings.Split(c2, "\n")
	//가져온 내용을 프린트한다.
	for i := 0; i < len(d1); i++ {
		fmt.Println(d1[i])
	}
	//배열의 4번째 내용을 바꾼다.
	for i := 1; i < len(d1)/4+1; i++{
		d1[4*i-1]=d2[2*i-1][7:]
	}
	fmt.Println(d1)
	for i := 0; i < len(d1); i++ {
		_, err = fo.WriteString(d1[i])
		fmt.Fprintf(fo, "\n")
		if err != nil {
			panic(err)
		}
	}
}
