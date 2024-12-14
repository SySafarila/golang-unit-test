package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Print("AFTER UNIT TEST")
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name      string
		request   string
		exprected string
	}{
		{
			name:      "Syahrul",
			request:   "Syahrul",
			exprected: "Hello Syahrul",
		},
		{
			name:      "Safarila",
			request:   "Safarila",
			exprected: "Hello Safarila",
		},
		{
			name:      "SySafarila",
			request:   "SySafarila",
			exprected: "Hello SySafarila",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.exprected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Syahrul", func(t *testing.T) {
		result := HelloWorld("Syahrul")
		require.Equal(t, "Hello Syahrul", result, "Result must be 'Hello Syahrul'")
	})

	t.Run("Safarila", func(t *testing.T) {
		result := HelloWorld("Safarila")
		require.Equal(t, "Hello Safarila", result, "Result must be 'Hello Safarila'")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Will not run on Windows")
	}

	result := HelloWorld("Syahrul")
	require.Equal(t, "Hello Syahrul", result, "Result must be 'Hello Syahrul'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Syahrul")
	require.Equal(t, "Hello Syahrul", result, "Result must be 'Hello Syahrul'")
	fmt.Println("TestHelloWorld with Assert done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Syahrul")
	assert.Equal(t, "Hello Syahrul", result, "Result must be 'Hello Syahrul'")
	fmt.Println("TestHelloWorld with Assert done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Syahrul")

	if result != "Hello Syahrul" {
		// error
		t.Error("Result must be Hello Syahrul")
	}

	fmt.Println("TestHelloWorldSyahrul Done")
}

func TestHelloWorldSafarila(t *testing.T) {
	result := HelloWorld("Safarila")

	if result != "Hello Safarila" {
		// error
		t.Fatal("Result must be Hello Safarila")
	}
	fmt.Println("TestHelloWorld Done")
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Syahrul")
	}
}

func BenchmarkHelloWorldSafarila(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Safarila")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Syahrul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Syahrul")
		}
	})

	b.Run("Safarila", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Safarila")
		}
	})
}
