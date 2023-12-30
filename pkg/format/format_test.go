package format

import (
	"reflect"
	"testing"
)

func TestNewOsv(t *testing.T) {
	type args struct {
		key   string
		title string
		url   string
		score string
		vec   string
	}
	tests := []struct {
		name string
		args args
		want Osv
	}{
		{
			name: "new osv",
			args: args{
				key:   "test001",
				title: "test002",
				url:   "http://test003",
				score: "1.0",
				vec:   "test005",
			},
			want: Osv{
				AdvisoryKey: "test001",
				Title:       "test002",
				Url:         "http://test003",
				Score:       "1.0",
				Vector:      "test005",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOsv(tt.args.key, tt.args.title, tt.args.url, tt.args.score, tt.args.vec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOsv_AddOsvs(t *testing.T) {
	type fields struct {
		AdvisoryKey string
		Title       string
		Url         string
		Score       string
		Vector      string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "add osvs",
			fields: fields{
				AdvisoryKey: "test001",
				Title:       "test002",
				Url:         "http://test003",
				Score:       "test004",
				Vector:      "test005",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Osv{
				AdvisoryKey: tt.fields.AdvisoryKey,
				Title:       tt.fields.Title,
				Url:         tt.fields.Url,
				Score:       tt.fields.Score,
				Vector:      tt.fields.Vector,
			}
			o.AddOsvs()
			if len(Osvs) != 1 {
				t.Error("Array length is not 1")
			}
		})
	}
}

func TestGetOsvs(t *testing.T) {
	tests := []struct {
		name string
		want [][]string
	}{
		{
			name: "et Osvs",
			want: [][]string{
				{"test001", "test002", "http://test003", "test004", "test005"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOsvs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOsvs() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPrintTable(t *testing.T) {
	type args struct {
		data [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "print table",
			args: args{
				data: [][]string{
					{"test001", "test002", "http://test003", "test004", "test005"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintTable(tt.args.data)
		})
	}
}
