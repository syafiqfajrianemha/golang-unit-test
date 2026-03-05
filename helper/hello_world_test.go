package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldSyafiq(t *testing.T) {
	result := HelloWorld("Syafiq")
	if result != "Hello Syafiq" {
		// error
		// t.Fail()
		t.Error("Result must be 'Hello Syafiq'")
	}

	fmt.Println("TestHelloWorldSyafiq Done")
}

func TestHelloWorldEmha(t *testing.T) {
	result := HelloWorld("Emha")
	if result != "Hello Emha" {
		// error
		// t.FailNow()
		t.Error("Result must be 'Hello Emha'")
	}

	fmt.Println("TestHelloWorldEmha Done")
}

// go test -v menjalakan semua unit test
// go test -v -run=TestHelloWorldSyafiq menjalankan spesifik test
