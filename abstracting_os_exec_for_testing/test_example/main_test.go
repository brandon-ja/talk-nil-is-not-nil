package main

import (
	"github.com/weave-lab/puffin"
	"testing"
)

// START_TEST OMIT
func Test_branchIsClean(t *testing.T) {
	type args struct {
		exec puffin.Exec
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := branchIsClean(tt.args.exec)
			if (err != nil) != tt.wantErr {
				t.Fatalf("branchIsClean() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("branchIsClean() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// END_TEST OMIT
