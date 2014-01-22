package foo

import(
	"testing"
)

func Test_foo1(t *testing.T){
	if foo1(10){
		t.Log(">0")
	}
}

