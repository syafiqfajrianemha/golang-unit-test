package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Syafiq")
	if result != "Hello Syafiq" {
		// error
		panic("Result is not 'Hello Syafiq'")
	}
}
