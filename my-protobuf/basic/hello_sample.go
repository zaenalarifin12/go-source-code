package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{Name: "zainal"}

	log.Println(&h)
}
