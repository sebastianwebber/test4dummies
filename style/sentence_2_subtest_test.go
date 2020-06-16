package style

import (
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
)

func Test_subtest_Sentence(t *testing.T) {
	t.Run("vai funcionar", func(t *testing.T) {
		out, _ := Sentence(input)

		if !reflect.DeepEqual(out, expected) {
			t.Errorf("Result not as expected:\n%v", diff.CharacterDiff(expected, out))
		}
	})

	t.Run("vai dar erro quando for em branco", func(t *testing.T) {
		_, err := Sentence("")

		if err == nil {
			t.Errorf("it should throw an error for empty strings")
		}
	})
}
