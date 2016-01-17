package tags

import (
	"fmt"
	"testing"
)

func TestTagCalculation(t *testing.T) {
	txt := "lorem ipsum. Bla text test lol! Ümläuts!"
	tags := Calculate(txt)
	fmt.Println(tags)
}
