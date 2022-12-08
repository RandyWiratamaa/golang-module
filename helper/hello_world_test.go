package helper

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
* Membuat unit test
**/

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Randy")

	if result != "Hello Randy" {
		/**
		* Error
		**/

		// t.Fail()                                // Unit test error, proses akan tetap berjalan
		// t.FailNow()                             // Unit test error, proses unit test akan langsung berhenti
		// t.Error("Result is not 'Hello Randy' ") // Sama seperti Log Error, dan secara otomatis memanggil function t.Fail()
		t.Fatal("result is not 'Hello Randy'") // Sama spt t.Error(), proses unit test akan langsung berhenti
	}
	fmt.Println("String ini bisa diakses setelah 't.Fail() dan t.Error()'")
	fmt.Println("String ini tidak akan bisa diakses setelah 't.FailNow() dan t.Fatal()'")
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

/**
GOROUTINE
**/

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestGoroutine(t *testing.T) {
	go RunHelloWorld() //Goroutine diawali dengan kata kunci 'go', akan running secara asynchronous. Tidak perlu menunggu proses selesai
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Randy Wiratama"
		fmt.Println("Selesai mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	defer close(channel)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Randy Wiratama"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Randy Wiratama"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// Anonymous Function
	go func() {
		channel <- "Randy"
		channel <- "Wiratama"
		channel <- "Golang"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// Anonymous Function
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {

}
