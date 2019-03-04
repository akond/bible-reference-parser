package parser

import "testing"

func Test_cleanText(t *testing.T) {
	type args struct {
		string string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	{"Empty string", args{""},""},
	{"Regular text", args{"В последние часы Своей земной жизни Иисус утешал учеников словами"},"В последние часы Своей земной жизни Иисус утешал учеников словами"},
	{"Text with angle brackets", args{"Выходит дух его, и он возвращается в землю свою: в тот день исчезают <все> помышления его."},"Выходит дух его, и он возвращается в землю свою: в тот день исчезают [все] помышления его."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanText(tt.args.string); got != tt.want {
				t.Errorf("cleanText() = %v, want %v", got, tt.want)
			}
		})
	}
}
