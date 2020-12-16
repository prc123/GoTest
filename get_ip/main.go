package main

import (
	"fmt"
	"net"
	"strings"
)

//GetOutbound get outbiund IP
func GetOutbound()(ip string,err error)  {
	conn,err:=net.Dial("udp","8.8.8.8:80")
	if err != nil {
		return 
	}
	defer conn.Close()
	fmt.Printf("%#v\n",conn.LocalAddr())
	fmt.Printf("%#v\n",conn.LocalAddr().(*net.UDPAddr))
	localAddr:=conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr)
	ip=strings.Split(localAddr.IP.String(),":")[0]
	return
}

func main()  {
	ip,err:=GetOutbound()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}