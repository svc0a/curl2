package curl2

import (
	"fmt"
	"testing"
)

func TestPostJson(t *testing.T) {
	out, err := PostJson[any, any]("https://example.com/data", nil, nil)
	if err != nil {
		return
	}
	fmt.Println(out)
}
