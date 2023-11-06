package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
)

func BasicUserGroup() {
	batman := basic.User{
		Id:       1,
		Username: "batman",
		IsActive: true,
		Password: []byte("batmanpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	suparman := basic.User{
		Id:       2,
		Username: "suparman",
		IsActive: true,
		Password: []byte("supamanpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	sapiderman := basic.User{
		Id:       3,
		Username: "sapiderman",
		IsActive: true,
		Password: []byte("sapidermanpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	batFamily := basic.UserGroup{
		GroupId:     999,
		GroupName:   "Bat Family",
		Roles:       nil,
		Users:       []*basic.User{&batman, &suparman, &sapiderman},
		Description: "",
	}

	jsonBytes, _ := protojson.Marshal(&batFamily)
	log.Println(string(jsonBytes))
}
