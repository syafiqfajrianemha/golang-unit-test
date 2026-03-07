package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// running semua sub test : go test -v -run=TestSubTest
// running hanya satu sub test : go test -v -run=TestSubTest/Syafiq
func TestSubTest(t *testing.T) {
	t.Run("Syafiq", func(t *testing.T) {
		result := HelloWorld("Syafiq")
		assert.Equal(t, "Hello Syafiq", result, "Result must be 'Hello Syafiq'")
	})
	t.Run("Emha", func(t *testing.T) {
		result := HelloWorld("Emha")
		assert.Equal(t, "Hello Emha", result, "Result must be 'Hello Emha'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac Os")
	}

	result := HelloWorld("Syafiq")
	assert.Equal(t, "Hello Syafiq", result, "Result must be 'Hello Syafiq'")
}

// gunakan library testify : go get github.com/stretchr/testify
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Syafiq")
	assert.Equal(t, "Hello Syafiq", result, "Result must be 'Hello Syafiq'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Syafiq")
	require.Equal(t, "Hello Syafiq", result, "Result must be 'Hello Syafiq'")
	fmt.Println("TestHelloWorld with Require Done")
}

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
