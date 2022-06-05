package funtion

import "fmt"

func ToJson(body string) {
	fmt.Println("To Json")
}

func FromJson() {
	fmt.Println("From Json")
}

func sliceContain(value string, slice []string) bool {
	exist := false
	for _, v := range slice {
		if value == v {
			exist = true
		}
	}
	return exist
}
