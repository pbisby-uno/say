package say_test

import (
	"fmt"
	"github.com/pbisby-uno/say"
	"testing"
)

func ExampleHello() {
	greeting := say.Hello("Bob")
	fmt.Println(greeting)
}

func TestHello(t *testing.T) {
	n := "Bob"
	expected := "Hello Bob."
	actual := say.Hello(n)

	if expected != actual {
		t.Logf("Hello: expected [%s] got [%s]", expected, actual)
		t.Fail()
	}
}
