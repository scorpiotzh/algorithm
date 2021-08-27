package option

import (
	"fmt"
	"testing"
)

func TestOption(t *testing.T) {
	c := NewClient(WithOption1("1"), WithOption2(1))
	fmt.Println(c)
}
