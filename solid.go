package caddy_dns_solid

import (
	"fmt"
	"github.com/bashopman/caddy-dns-solid/solid_provider"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

type Provider struct{ *solid_provider.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

func (Provider) CaddyModule() caddy.ModuleInfo {
	fmt.Println("solid.go:CaddyModule")
	return caddy.ModuleInfo{
		ID:  "dns.providers.solid",
		New: func() caddy.Module { return &Provider{new(solid_provider.Provider)} },
	}
}

func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.SolidApiUrl = caddy.NewReplacer().ReplaceAll(p.Provider.SolidApiUrl, "")
	return nil
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.SolidApiUrl = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "solidapiurl":
				if p.Provider.SolidApiUrl != "" {
					return d.Err("Account name already set")
				}
				if d.NextArg() {
					p.Provider.SolidApiUrl = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.SolidApiUrl == "" {
		return d.Err("missing account name")
	}
	return nil
}

var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
