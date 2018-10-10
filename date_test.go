package monent

import (
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	for _,x := range Split(Now(Year), Duration(3*time.Hour)) {
		fmt.Println(x)
	}
}
