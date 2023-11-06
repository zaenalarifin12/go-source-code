package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"
	"time"
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

	comm := randomCommunicationChannel()

	u := basic.User{
		Id:                   1,
		Username:             "Zainal",
		IsActive:             true,
		Password:             []byte("suparmanheh"),
		Email:                []string{"supermanmovie@gmail.com", "suparmandc@gmail.com"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		CommunicationChannel: &comm,
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

func randomCommunicationChannel() anypb.Any {
	rand.Seed(time.Now().UnixNano())

	paperMail := basic.PaperMail{PaperMailAddress: "some people mail address"}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "facebook",
		SocialMediaUsername: "babahahah",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "whatsup",
		InstantMessagingUsername: "wawawawa",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	case 2:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a

}

func BasicUnmarshalAnyKnow() {
	sociaMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &sociaMedia, proto.MarshalOptions{})

	// know type social media
	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}

func BasicUnsmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage
	unmarshaled, err := a.UnmarshalNew()

	if err != nil {
		return
	}

	log.Println("Unmarshal as ", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, but : ", a.TypeUrl)
	}
}

func BasicOneOf() {
	//socialMedia := basic.SocialMedia{
	//	SocialMediaPlatform: "byteme",
	//	SocialMediaUsername: "aquaman",
	//}
	//
	//ecomm := basic.User_SocialMedia{SocialMedia: &socialMedia}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "wa",
		InstantMessagingUsername: "aquaman",
	}

	ecomm := basic.User_InstantMessaging{InstantMessaging: &instantMessaging}

	u := basic.User{
		Id:                    96,
		Username:              "sdsadasd",
		IsActive:              false,
		Password:              nil,
		Email:                 nil,
		Gender:                0,
		Address:               nil,
		CommunicationChannel:  nil,
		ElectronicCommChannel: &ecomm,
	}

	bytes, _ := protojson.Marshal(&u)
	log.Println(string(bytes))
}
