package main

import (
	"fmt"
)

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	ServiceID uint64
	Status    string
}

func checkEven(i int) string {
	if i%2 == 0 && i != 0 {
		return PassStatus
	} else {
		return FailStatus
	}
}

func GenerateCheck() []HealthCheck {
	var check []HealthCheck
	for i := 0; i <= 5; i++ {
		check = append(check, HealthCheck{ServiceID: uint64(i), Status: checkEven(i)})
	}
	return check
}

func main() {
	slice := GenerateCheck()
	fmt.Println("Тут будет выведен идентификатор")
	for _, v := range slice {
		if v.Status == PassStatus {
			fmt.Println(v.ServiceID)
		}
	}
}
