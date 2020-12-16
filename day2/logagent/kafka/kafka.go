package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)
var(
	client sarama.SyncProducer//申明一个全局的连接kafka的生产者client
)
//Init kafka
func Init(addrs []string) (err error){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Return.Successes=true
	
	client,err=sarama.NewSyncProducer(addrs,config)
	
	if err!=nil{
		fmt.Println("111closed , err",err)
		return
	}
	fmt.Print(config.Producer)
	return
}
//SendTokafka senMessage to kafa 
func SendTokafka(topic,data string)  {
	msg:=&sarama.ProducerMessage{}
	msg.Topic=topic
	msg.Value=sarama.StringEncoder(data)
	fmt.Println(msg)
	pid,offset,err:=client.SendMessage(msg)
	if err!=nil{
		fmt.Printf("pid:%v offset:%v\n",pid,offset)
		fmt.Println("12345closed , err",err)
		return
	}
	fmt.Printf("pid:%v offset:%v/n",pid,offset)
}
//CloseClient !
func CloseClient()error{
	return client.Close()
}