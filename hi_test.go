package modtest

import "testing"


func TestHi(t *testing.T) {
	type args struct {
		greeting string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name:"ok",
			args:args{greeting:"xx"},
			want:"echo hi xx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hi(tt.args.greeting); got != tt.want {
				t.Errorf("Hi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReflectLaw(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name:"ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReflectLaw()
		})
	}
}
