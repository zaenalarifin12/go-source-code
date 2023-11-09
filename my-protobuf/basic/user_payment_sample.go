package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicReadUserPayment() {
	log.Println("read user payment")

	up := basic.UserPayment{}

	ReadProtoFromFile("user_content_v1.bin", &up)

	log.Println(&up)

	upJson, _ := protojson.Marshal(&up)
	log.Println(string(upJson))
	log.Println()
}
