package forms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {
	tests := []struct {
		strct interface{}
		want  []field
	}{
		{
			strct: struct {
				Name string
			}{},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "",
				},
			},
		},
		{
			strct: struct {
				FullName string
			}{},
			want: []field{
				{
					Label:       "FullName",
					Name:        "FullName",
					Type:        "text",
					Placeholder: "FullName",
					Value:       "",
				},
			},
		},
		{
			strct: struct {
				Name  string
				Email string
				Age   int
			}{},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "",
				},
				{
					Label:       "Email",
					Name:        "Email",
					Type:        "text",
					Placeholder: "Email",
					Value:       "",
				},
				{
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       0,
				},
			},
		},
		{
			strct: struct {
				Name  string
				Email string
				Age   int
			}{
				Name:  "Kasia Kawala",
				Email: "kat.kawala@gmail.com",
				Age:   123,
			},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "Kasia Kawala",
				},
				{
					Label:       "Email",
					Name:        "Email",
					Type:        "text",
					Placeholder: "Email",
					Value:       "kat.kawala@gmail.com",
				},
				{
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       123,
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%T", tc.strct), func(t *testing.T) {
			got := fields(tc.strct)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("fields() = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestFields_labels(t *testing.T) {
	hasLabels := func(labels ...string) func(*testing.T, []field) {
		return func(t *testing.T, fields []field) {
			if len(fields) != len(labels) {
				t.Fatalf("fields() len = %d; want %d", len(fields), len(labels))
			}
			for i := 0; i < len(fields); i++ {
				if fields[i].Label != labels[i] {
					t.Errorf("fields()[%d].Label = %s; want %s", i, fields[i].Label, labels[i])
				}
			}
		}
	}
	hasValues := func(values ...interface{}) func(*testing.T, []field) {
		return func(t *testing.T, fields []field) {
			if len(fields) != len(values) {
				t.Fatalf("fields() len = %d; want %d", len(fields), len(values))
			}
			for i := 0; i < len(fields); i++ {
				if fields[i].Value != values[i] {
					t.Errorf("fields()[%d].Value = %v; want %v", i, fields[i].Value, values[i])
				}
			}
		}
	}
	check := func(checks ...func(*testing.T, []field)) []func(*testing.T, []field) {
		return checks
	}

	tests := map[string]struct {
		strct  interface{}
		checks []func(*testing.T, []field)
	}{
		"No values": {
			strct: struct {
				Name string
			}{},
			checks: check(hasLabels("Name")),
		},
		"Multiple fields with values": {
			strct: struct {
				Name  string
				Email string
				Age   int
			}{
				Name:  "Kasia Kawala",
				Email: "kat.kawala@gmail.com",
				Age:   123,
			},
			checks: check(
				hasLabels("Name", "Email", "Age"),
				hasValues("Kasia Kawala", "kat.kawala@gmail.com", 123),
			),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := fields(tc.strct)
			for _, check := range tc.checks {
				check(t, got)
			}
		})
	}
}
