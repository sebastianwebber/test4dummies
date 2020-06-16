package style

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
)

var (
	input    = `primeira frase. segunda frase com acento. é...só que não.`
	expected = `Primeira frase. Segunda frase com acento. É...Só que não.`
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

func Test_basic_Sentence(t *testing.T) {

	out, err := Sentence(input)

	if err != nil {
		t.Errorf("error running Title: %w", err)
	}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Result not as expected:\n%v", diff.CharacterDiff(expected, out))
	}

}
