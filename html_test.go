package forms

import (
	"html/template"
	"testing"
)

func TestHTML(t *testing.T) {
	tests := map[string]struct {
		tpl     *template.Template
		strct   interface{}
		want    template.HTML
		wantErr error
	}{}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := HTML(tc.tpl, tc.strct)
			if err != tc.wantErr {
				t.Fatalf("HTML() err = %v; want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("HTML() = %q; want %q", got, tc.want)
			}
		})
	}
}
