package basic

import (
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"
)

func writeProtoToFile(message proto.Message, fname string) {
	b, err := proto.Marshal(message)

	if err != nil {
		log.Fatalln("Can't marshal message", err)
	}

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
	}
}

func ReadProtoFromFile(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can't read file", err)
	}

	if err := proto.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can't unmarshal", err)
	}
}

func dummyUser() basic.User {
	addr := basic.Address{
		Street:  "Daily Planet",
		City:    "Metropolis",
		Country: "US",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  40.04040404435345,
			Longitude: 05.34534534534534,
		},
	}

	comm := randomCommunicationChannel()
	sr := map[string]uint32{"fly": 5, "speed": 3, "durability": 4}

	return basic.User{
		Id:                    22,
		Username:              "asdasd",
		IsActive:              true,
		Password:              []byte("asdasasdasdasd"),
		Email:                 []string{"asdasasdasdasd@gmail.com"},
		Gender:                0,
		Address:               &addr,
		CommunicationChannel:  &comm,
		ElectronicCommChannel: nil,
		SkillRating:           sr,
	}
}

func WriteToFileSample() {
	u := dummyUser()
	writeProtoToFile(&u, "superman_file.bin")
}

func ReadFromFileSample() {
	dest := basic.User{}

	ReadProtoFromFile("superman_file.bin", &dest)

	log.Println(&dest)
}
