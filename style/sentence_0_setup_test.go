package style

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	beforeTests()
	out := m.Run()
	afterTests()
	os.Exit(out)
}

func beforeTests() {
	fmt.Println("Running tests Setup - TEAR UP")
}

func afterTests() {
	fmt.Println("Ending tests - TEAR DOWN")
}
