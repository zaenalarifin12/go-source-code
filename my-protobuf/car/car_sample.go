package car

import (
	"github.com/google/uuid"
	"log"
	car2 "my-protobuf/protogen/car"
)

func ValidateCar() {
	car := car2.Car{
		CarId:           uuid.NewString(),
		Brand:           "Honda",
		Model:           "Sport",
		Price:           20000000,
		ManufactureYear: 2024,
	}

	err := car.ValidateAll()

	if err != nil {
		log.Fatalln("Validation failed", err)
	}

	log.Println(&car)
}
