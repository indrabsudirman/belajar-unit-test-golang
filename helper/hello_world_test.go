package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Indra")
	require.Equal(t, "Hello Indra", result, "result must be 'Hello Indra'")
	//require call FailNow, jadi tidak akan sampai ke bawah ini
	fmt.Println("Test HelloWorldRequire done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Indra")
	assert.Equal(t, "Hello Indra", result, "result must be 'Hello Indra'")
	fmt.Println("TestHelloWorldAssert done")
}

func TestHelloWorldIndra(t *testing.T) {
	result := HelloWorld("Indra")

	if result != "Hello Indra" {
		//Error
		t.Error("result must be 'Hello Indra'")
	}
	fmt.Println("TestHelloWorldIndra done")
}
func TestHelloWorldHaby(t *testing.T) {
	result := HelloWorld("Haby")

	if result != "Hello Haby" {
		//Error
		t.Fatal("result must be 'Hello Haby'")
	}
	fmt.Println("TestHelloWorldIndra done")
}
