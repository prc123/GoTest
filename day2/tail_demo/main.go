package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)
func main()  {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset:0,Whence:2},
		MustExist: false,
		Poll: true,
	}	
	tails,err:=tail.TailFile(fileName,config)
	if err !=nil{
		fmt.Println("tail file failed ,err:",err)
	}
	var(
		line *tail.Line
		ok bool
	)
	
	for {
		// fmt.Println(tails.fi)

		// line,ok = <-tails.Lines
		line,ok=<-tails.Lines
		//fmt.Println("yes")
		if !ok{
			fmt.Printf("tail file close reopen,filename:%s\n",tails.Filename)
			time.Sleep(time.Second)
			continue
		}

		fmt.Println("line:",line.Text)
	}
}