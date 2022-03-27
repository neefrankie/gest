package html

import (
	"testing"
)

func TestNewHTML(t *testing.T) {
	tests := []struct {
		name string
		args Params
		want *HTML
	}{
		{
			name: "HTML",
			args: Params{
				Lang: "en",
				Head: nil,
				Body: nil,
			},
			want: &HTML{
				Tag{
					Token:       NewToken("html"),
					SelfClosing: false,
					Attr:        NewAttr().Set("lang", "en"),
					Children: []Node{
						&Head{},
						&Body{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewHTML(tt.args)

			if got.String() != tt.want.String() {
				t.Errorf("NewHTML() = %v, want %v", got, tt.want)
				return
			}

			t.Logf("%s", got)
		})
	}
}
