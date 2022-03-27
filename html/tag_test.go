package html

import "testing"

func TestTag_String(t1 *testing.T) {

	tests := []struct {
		name   string
		fields Tag
		want   string
	}{
		{
			name: "Tag string",
			fields: Tag{
				Token:       Token{"div"},
				SelfClosing: false,
				Attr:        NewAttr().Set("class", "d-flex"),
				Children:    nil,
			},
			want: `<div class="d-flex"></div>`,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tag{
				Token:       tt.fields.Token,
				SelfClosing: tt.fields.SelfClosing,
				Attr:        tt.fields.Attr,
				Children:    tt.fields.Children,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
