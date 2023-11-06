package jobsearch

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/dummy"
	"my-protobuf/protogen/jobsearch"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 777,
		Application: &basic.Application{
			Version:  "1.0.0",
			Name:     "The amazing app",
			Platform: []string{"Mac", "Windows"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("Software : " + string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 888,
		Application: &dummy.Application{
			ApplicationId:       887,
			ApplicationFullName: "shazam",
			Phone:               "097098098",
			Email:               "shazam@gmail.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Println("Shazam : " + string(jsonBytes))
}
