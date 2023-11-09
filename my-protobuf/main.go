package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Println(time.Now().Format("15:04:05" + " " + string(bytes)))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	//basic.BasicHello()
	basic.BasicUser()
	//basic.BasicUserGroup()
	//basic.ProtoToJsonUser()
	//basic.JsonToProto()
	//basic.BasicUnmarshalAnyKnow()
	//basic.BasicUnmarshalAnyIs()
	//basic.BasicOneOf()
	//basic.WriteToFileSample()
	//basic.ReadFromFileSample()
	//basic.WriteToJSONSample()
	//basic.ReadFromJSONSample()
	//basic.BasicWriteUserContentV1()
	//basic.BasicReadUserContentV1()
	//basic.BasicWriteUserContentV2()
	//basic.BasicReadUserContentV2()
	//basic.BasicReadUserPayment()
}
