package main

import (
	"fmt"
	"time"
	// "io"
	// "os"

	"example.com/m/v2/GoTest/day2/logagent/config"
	"example.com/m/v2/GoTest/day2/logagent/kafka"
	"example.com/m/v2/GoTest/day2/logagent/taillog"
	"gopkg.in/ini.v1"
)



// func main(){
	
// }

var (
	cfg =new(config.AppConf)
)


func run()  {
	for {
		select{
		case line:=<-taillog.ReadChan():
			print(line.Text)
			kafka.SendTokafka(cfg.FileName,line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}

func main(){
	
	path:="./config/config.ini"
	// f,err:=os.Open(path)
	// if err !=nil{
	// 	panic(err)
	// }
	// defer f.Close()
	// var chunks []byte
	// buf := make([]byte, 1024)
    // for {
    //     n, err := f.Read(buf)
    //     if err != nil &&err!=io.EOF{
    //         panic(err)
    //     }
    //     if 0 == n {
    //         break
    //     }
    //     //fmt.Println(string(buf))
    //     chunks = append(chunks, buf...)
    // }
   	// fmt.Print(string(chunks)) 
	err :=ini.MapTo(cfg,path)	
	if err !=nil{
		panic(err)
	}
	// a:=cfg.Address
	fmt.Println(cfg.KafkaConf.Address) 
	err=kafka.Init([]string{cfg.KafkaConf.Address})
	if err !=nil{
		panic(err)
	}
	err=taillog.Init(cfg.FileName)
	if err !=nil{
		panic(err)
	}
	run()
}