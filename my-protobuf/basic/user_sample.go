package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUser() {

	addr := basic.Address{
		Street:  "JL kusuma",
		City:    "Saturnus",
		Country: "Jupiter",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  40.00000000000,
			Longitude: -76.0000000000,
		},
	}

	u := basic.User{
		Id:       1,
		Username: "Zainal",
		IsActive: true,
		Password: []byte("suparmanheh"),
		Email:    []string{"supermanmovie@gmail.com", "suparmandc@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
		Address:  &addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       2,
		Username: "Babayaga",
		IsActive: true,
		Password: []byte("babayaga"),
		Email:    []string{"babayaga@gmail.com", "babayaga22222@gmail.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProto() {
	json := `
		{
		  "id": 6,
		  "username": "Babayaga",
		  "is_active": true,
		  "password": "YmFiYXlhZ6E=",
		  "email": [
			"babayaga@gmail.com",
			"babayaga66666@gmail.com"
		  ],
		  "gender": "GENDER_FEMALE"
		}
	`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln("Got Error: ", err)
	}

	log.Println(&p)
}
