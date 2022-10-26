package caddy_dns_solid

import (
	"github.com/bashopman/caddy-dns-solid/solid_provider"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"reflect"
	"testing"
)

func TestProvider_CaddyModule(t *testing.T) {
	type fields struct {
		Provider *solid_provider.Provider
	}
	tests := []struct {
		name   string
		fields fields
		want   caddy.ModuleInfo
	}{
		{
			name:   "successful ModuleInfo",
			fields: fields{},
			want: caddy.ModuleInfo{
				ID: "dns.providers.solid",
				New: func() caddy.Module {
					return &Provider{}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := Provider{
				Provider: tt.fields.Provider,
			}
			if got := pr.CaddyModule(); !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("CaddyModule() = %v, want %v", got.ID, tt.want.ID)
			}
		})
	}
}

func TestProvider_Provision(t *testing.T) {
	type fields struct {
		Provider *solid_provider.Provider
	}
	type args struct {
		ctx caddy.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "provision",
			fields: fields{},
			args: args{
				caddy.Context{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				Provider: tt.fields.Provider,
			}
			if err := p.Provision(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Provision() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProvider_UnmarshalCaddyfile(t *testing.T) {
	type fields struct {
		Provider *solid_provider.Provider
	}
	type args struct {
		d *caddyfile.Dispenser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				Provider: tt.fields.Provider,
			}
			if err := p.UnmarshalCaddyfile(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalCaddyfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
