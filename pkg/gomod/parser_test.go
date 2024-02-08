package gomod

import (
	"reflect"
	"testing"

	"github.com/siroa/go-osv-scanner/pkg/api"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

func TestNewGoMod(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *GoMod
	}{
		{
			name: "New_GoMod",
			args: args{
				name: "test01",
			},
			want: &GoMod{
				Name:    "test01",
				Modules: []Module{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGoMod(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoMod_setModules(t *testing.T) {
	type fields struct {
		Name    string
		Modules []Module
	}
	type args struct {
		ms []Module
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set_Modules",
			fields: fields{
				Name: "test01",
				Modules: []Module{
					{
						Name:        "test01",
						Version:     "test01",
						AdvisoryIDs: []string{},
					},
					{
						Name:        "test02",
						Version:     "test02",
						AdvisoryIDs: []string{},
					},
				},
			},
			args: args{
				ms: []Module{
					{
						Name:        "test01",
						Version:     "test01",
						AdvisoryIDs: []string{},
					},
					{
						Name:        "test02",
						Version:     "test02",
						AdvisoryIDs: []string{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GoMod{
				Name:    tt.fields.Name,
				Modules: tt.fields.Modules,
			}
			g.setModules(tt.args.ms)
		})
	}
}

func TestNewModule(t *testing.T) {
	type args struct {
		name string
		ver  string
	}
	tests := []struct {
		name string
		args args
		want *Module
	}{
		{
			name: "New_Module",
			args: args{
				name: "test01",
				ver:  "test01",
			},
			want: &Module{
				Name:    "test01",
				Version: "test01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewModule(tt.args.name, tt.args.ver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewModule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_SetAdvisoryKeys(t *testing.T) {
	type fields struct {
		Name        string
		Version     string
		AdvisoryIDs []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "set advisoryKeys",
			fields: fields{
				Name:        "test01",
				Version:     "v1.0.0",
				AdvisoryIDs: []string{},
			},
			want: []string{"Key1", "Key2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				Name:        tt.fields.Name,
				Version:     tt.fields.Version,
				AdvisoryIDs: tt.fields.AdvisoryIDs,
			}
			m.SetAdvisoryKeys(api.DepsdevRepositoryStub{})
			if !reflect.DeepEqual(tt.want, m.AdvisoryIDs) {
				t.Error("AdvisoryID is different")
			}
		})
	}

}

func TestModule_PrintModule(t *testing.T) {
	type fields struct {
		Name        string
		Version     string
		AdvisoryIDs []string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "No error print module",
			fields: fields{
				Name:        "test",
				Version:     "test",
				AdvisoryIDs: []string{"test01", "test02"},
			},
		},
		{
			name: "No advisory ids",
			fields: fields{
				Name:        "test",
				Version:     "test",
				AdvisoryIDs: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Module{
				Name:        tt.fields.Name,
				Version:     tt.fields.Version,
				AdvisoryIDs: tt.fields.AdvisoryIDs,
			}
			m.PrintModule()
		})
	}
}

func Test_isReplaceModule(t *testing.T) {
	type args struct {
		replace []*modfile.Replace
		path    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contain replace module",
			args: args{
				replace: []*modfile.Replace{{Old: module.Version{Path: "replacemod"}}},
				path:    "replacemod",
			},
			want: true,
		},
		{
			name: "Doesn't contain replace module",
			args: args{
				replace: []*modfile.Replace{{Old: module.Version{Path: "replacemod"}}},
				path:    "module",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReplaceModule(tt.args.replace, tt.args.path); got != tt.want {
				t.Errorf("isReplaceModule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseGoMod(t *testing.T) {
	gomod := `
	module scanner

	go 1.21.4

	replace project1 => ../project1
	
	require (
		project1 v0.0.0
		github.com/edoardottt/depsdev v0.0.8
	)
	
	require (
		github.com/avast/retry-go v3.0.0+incompatible // indirect
	)`

	type args struct {
		file []byte
	}
	tests := []struct {
		name string
		args args
		want *GoMod
	}{
		{
			name: "parse GoMod",
			args: args{
				file: []byte(gomod),
			},
			want: &GoMod{
				Name: "scanner",
				Modules: []Module{
					{
						Name:    "github.com/edoardottt/depsdev",
						Version: "v0.0.8",
					},
					{
						Name:    "github.com/avast/retry-go",
						Version: "v3.0.0+incompatible",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseGoMod(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseGoMod() = %v, want %v", got, tt.want)
			}
		})
	}
}
