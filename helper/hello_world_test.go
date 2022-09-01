package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Membuat unit test

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Randy")

	if result != "Hello Randy" {
		//error
		t.Fatal("result is not 'Hello Randy'")
	}
	fmt.Println("String ini tidak akan bisa diakses setelah 't.Fatal'")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Randy")
	assert.Equal(t, "Hello Randy", result, "Result must be 'Hello Randy'")
	fmt.Println("Test Hello World with Assertion")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Randy")
	require.Equal(t, "Hello Randy", result, "Result must be 'Hello Randy'")
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Randy", func(t *testing.T) {
		result := HelloWorld("Randy")
		require.Equal(t, "Hello Randy", result, "Result must be 'Hello Randy'")
	})
	t.Run("Wiratama", func(t *testing.T) {
		result := HelloWorld("Wiratama")
		require.Equal(t, "Hello Wiratama", result, "Result must be 'Hello Wiratama'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Randy",
			request:  "Randy",
			expected: "Hello Randy",
		},
		{
			name:     "Wiratama",
			request:  "Wiratama",
			expected: "Hello Wiratama",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Randy")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Randy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Randy")
		}
	})

	b.Run("Wiratama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Wiratama")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Randy",
			request: "Randy",
		},
		{
			name:    "Wiratama",
			request: "Wiratama",
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
