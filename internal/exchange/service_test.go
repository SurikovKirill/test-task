package exchange

import (
	"testing"
)

func Test_setValidName(t *testing.T) {
	type args struct {
		rem int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Единица в остатке",
			args: args{rem: 21},
			want: "рублю",
		},
		{
			name: "В остатке не единица",
			args: args{rem: 22},
			want: "рублям",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setValidName(tt.args.rem); got != tt.want {
				t.Errorf("setValidName() = %v, want %v", got, tt.want)
			}
		})
	}
}
