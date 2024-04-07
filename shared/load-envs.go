package shared

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
}
