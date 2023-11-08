package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"
)

func writeProtoToJSON(message proto.Message, fname string) {
	b, err := protojson.Marshal(message)

	if err != nil {
		log.Fatalln("Can't marshal message", err)
	}

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
	}
}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read file", err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can't unmarshal", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()
	writeProtoToJSON(&u, "superman_file.json")
}

func ReadFromJSONSample() {
	dest := basic.User{}

	ReadProtoFromJSON("superman_file.json", &dest)

	log.Println(&dest)
}
