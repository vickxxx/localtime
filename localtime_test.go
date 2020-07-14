package localtime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type S struct {
	T LocalTime `json:"t"`
}

func TestMain(t *testing.T) {
	s := S{
		T: LocalTime(time.Now()),
	}

	data, err := json.Marshal(&s)
	fmt.Println(string(data), err)
	fmt.Println(s)
}

func TestMain2(t *testing.T) {
	fmt.Println("Xxx")
}
