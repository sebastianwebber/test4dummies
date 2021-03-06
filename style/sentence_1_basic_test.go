package style

import (
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
)

var (
	input    = `primeira frase. segunda frase com acento. é...só que não.`
	expected = `Primeira frase. Segunda frase com acento. É...Só que não.`
)

func Test_basic_Sentence(t *testing.T) {

	out, _ := Sentence(input)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Result not as expected:\n%v", diff.CharacterDiff(expected, out))
	}
}
