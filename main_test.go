package main

import (
	"fmt"
	"os"
	"testing"
)

// Setup do teste vai aqui!
func TestMain(m *testing.M) {
	antesDoTeste()
	out := m.Run()
	depoisDoTeste()
	os.Exit(out)
}

func antesDoTeste() {
	fmt.Println("Running test Setup")
}

func depoisDoTeste() {
	fmt.Println("Running test tear down")
}

func Test_superUpper(t *testing.T) {
	genValue, err := superUpper("golang-meetup")
	expectedValue := "GOLANG-MEETUP"

	if err != nil {
		t.Errorf("Could not call superUpper: %v", err)
	}
	if genValue != expectedValue {
		t.Errorf("Expected '%s', got '%s'",
			expectedValue,
			genValue)
	}
}

func Test_superUpper_subTests(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		wantOut string
		wantErr bool
	}{
		{name: "error when is empty", args: "", wantErr: true},
		{name: "ok when send a valid value", args: "golang-meetup", wantOut: "GOLANG-MEETUP", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := superUpper(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("superUpper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("superUpper() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
