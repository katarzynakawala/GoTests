package forms

import (
	"html/template"
	"testing"
)

var (
	tplTypeNameValue = template.Must(template.New("").Parse(`<input type="{{.Type}}" name="{{.Name}}"{{with .Value}} value="{{.}}"{{end}}>`))
)

func TestHTML(t *testing.T) {
	tests := map[string]struct {
		tpl     *template.Template
		strct   interface{}
		want    template.HTML
		wantErr error
	}{
		"A simple form with values": {
			tpl: tplTypeNameValue,
			strct: struct {
				Name  string
				Email string
			}{
				Name:  "Kajetan Ka",
				Email: "kajetanka@gmail.com",
			},
			want: `<input type="text" name="Name" value="Kajetan Ka">` +
				`<input type="text" name="Email" value="kajetanka@gmail.com">`,
		},
	}

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
