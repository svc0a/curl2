package curl2

import (
	"github.com/svc0a/enum/gen"
	"testing"
)

func TestEnum(t *testing.T) {
	gen.Generate("vo.go")
}
