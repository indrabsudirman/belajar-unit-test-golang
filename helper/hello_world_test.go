package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Indra",
			request: "Indra",
		},
		{
			name:    "Bayu",
			request: "Bayu",
		},
		{
			name:    "Sudirman",
			request: "Sudirman",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})

	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Indra", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Indra")
		}
	})
	b.Run("Bayu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bayu")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Indra")
	}
}
func BenchmarkHelloWorldBayu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bayu")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Indra",
			request:  "Indra",
			expected: "Hello Indra",
		},
		{
			name:     "Bayu",
			request:  "Bayu",
			expected: "Hello Bayu",
		},
		{
			name:     "Sudirman",
			request:  "Sudirman",
			expected: "Hello Sudirman",
		},
		{
			name:     "Lubna",
			request:  "Lubna",
			expected: "Hello Lubna",
		},
		{
			name:     "Haby",
			request:  "Haby",
			expected: "Hello Haby",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	//Membuat sub test didalam function test
	t.Run("Indra", func(t *testing.T) {
		result := HelloWorld("Indra")
		assert.Equal(t, "Hello Indra", result, "result must be 'Hello Indra'")
	})
	t.Run("Sudirman", func(t *testing.T) {
		result := HelloWorld("Sudirman")
		require.Equal(t, "Hello Sudirman", result, "result must be 'Hello Sudirman'")
	})
}

func TestMain(m *testing.M) {
	//Before, disini bisa eksekusi code before, seperti koneksi ke database dll
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	//After, disini bisa eksekusi code after
	fmt.Println("AFTER UNIT TEST")
}
func TestHelloWorldSkipTest(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Test skip for windows OS")
	}
	result := HelloWorld("Indra")
	require.Equal(t, "Hello Indra", result, "result must be 'Hello Indra'")
	//require call FailNow, jadi tidak akan sampai ke bawah ini
	fmt.Println("Test HelloWorldRequire done")
}

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
