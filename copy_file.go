/*package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	a := "C:/Users/wind/Desktop/2006400013饶安.jpg"
	b := "D:/go-language/go_program/1.jpg"
	pp, err := os.Open(a)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err = pp.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	bufa := bufio.NewReader(pp)
	fw, err := os.Create(b)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		err = fw.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	bufb := bufio.NewWriter(fw)
	count, err := io.Copy(bufb, bufa)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(count)
	fmt.Println("over")
}
*/
