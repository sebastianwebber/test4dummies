package style

import "testing"

func Test_table_Sentence(t *testing.T) {

	// to make them parallel
	t.Parallel()

	tests := []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{name: "should throw an error when string is empty", args: "", want: "", wantErr: true},
		{name: "should convert string without errors", args: input, want: expected, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sentence(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sentence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
