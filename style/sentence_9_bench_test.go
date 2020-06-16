package style

import "testing"

func BenchmarkSentenceBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sentence("super fake string to put here")
	}
}

const (
	oneSentence    = "apenas uma frase."
	threeSentences = "apenas uma frase. mentira! essa é a segunda."
	lorem          = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin vitae libero vel lacus aliquet scelerisque et nec lorem. Proin ultricies ex et eros condimentum pellentesque. In at malesuada ligula, eget malesuada ipsum. Aliquam eget eros vitae orci vehicula auctor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Vivamus nec bibendum felis, in imperdiet diam. Fusce sollicitudin commodo nisi malesuada pulvinar. In aliquam, dui eu dapibus vehicula, metus dui mattis arcu, quis porta sem eros in purus. Curabitur vel tempus lectus. In faucibus odio id justo tincidunt, eu vulputate purus scelerisque. Pellentesque varius leo scelerisque, porttitor nulla in, convallis ex. Morbi tempus est sed aliquam blandit. Cras accumsan odio et metus aliquam, vitae posuere tellus finibus. Sed a purus nec massa congue rutrum. Praesent in ante quis ligula blandit auctor a id metus.`
)

func BenchmarkSentenceSub(b *testing.B) {

	b.Run("with 1 sentence", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Sentence(oneSentence)
		}
	})
	b.Run("with 3 sentences", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Sentence(threeSentences)
		}
	})
}

func BenchmarkSentenceTable(b *testing.B) {

	tests := []struct {
		name string
		args string
	}{
		{name: "with 1 sentence", args: oneSentence},
		{name: "with 3 sentences", args: threeSentences},
		{name: "Lorem Ipsum", args: lorem},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sentence(tt.args)
			}
		})
	}
}
