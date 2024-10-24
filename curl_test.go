package curl2

import (
	"fmt"
	"testing"
)

func TestPostJson(t *testing.T) {
	out, err := Define[any, any]().Get()
	if err != nil {
		return
	}
	fmt.Println(out)
}
