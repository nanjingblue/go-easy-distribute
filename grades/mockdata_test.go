package grades

import (
	"testing"
)

func TestFaker(t *testing.T) {
	students := GetFaker(1)
	t.Errorf("%v", students)
}
