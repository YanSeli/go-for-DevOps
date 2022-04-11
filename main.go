package main

import (
	"fmt"
)

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	ServiceID string
	Status    string
}

func GenerateCheck() HealthCheck {
	return HealthCheck{
		ServiceID: "",
		Status:    PassStatus,
	}
}

func main() {
	fmt.Println("Тут будет выведен идентификатор")
	fmt.Println(GenerateCheck())
}
