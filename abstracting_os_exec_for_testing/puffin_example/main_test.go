package main

import (
	"github.com/weave-lab/puffin"
	"testing"
)

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
		// START_CASES OMIT
		{"is clean",
			args{
				puffin.NewFuncExec(puffin.NewHandlerMux(
					func(cmd *puffin.FuncCmd) int {
						return 0
					},
				)),
			},
			true,
			false,
		},
		{"is dirty",
			args{
				puffin.NewFuncExec(puffin.NewHandlerMux(
					func(cmd *puffin.FuncCmd) int {
						cmd.Stdout().Write([]byte("M README.md"))
						return 0
					},
				)),
			},
			false,
			false,
		},
		// END_CASES OMIT
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
