package config


//AppConf struct
type AppConf struct{
	KafkaConf `ini:"kafka"`
	TailLog `ini:"taillog"`

}
//KafkaConf struct
type KafkaConf struct{
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}
//TailLog struct
type TailLog struct{
	FileName string `ini:"path"`
}

