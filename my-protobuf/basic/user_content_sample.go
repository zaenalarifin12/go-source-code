package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug:          "/blog-ver1",
		Title:         "10 strong man",
		HtmlContent:   "<p>dummy blog strong man </p>",
		AuthorId:      99,
	}

	writeProtoToFile(&uc, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Read V1")

	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v1.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 1002,
		Slug:          "/blog-ver2",
		Title:         "20 strong man",
		//HtmlContent:   "<p>dummy blog 20 strong man </p>",
		//AuthorId:      99,
		//Category: "NEWS",
	}

	writeProtoToFile(&uc, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Read V2")

	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v2.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}
