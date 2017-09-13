package bind

import (
	"github.com/WedgeServer/wedge"
	"github.com/WedgeServer/wedge/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("bind", caddy.Plugin{
		ServerType: "http",
		Action:     setupBind,
	})
}

func setupBind(c *caddy.Controller) error {
	config := httpserver.GetConfig(c)
	for c.Next() {
		if !c.Args(&config.ListenHost) {
			return c.ArgErr()
		}
		config.TLS.ListenHost = config.ListenHost // necessary for ACME challenges, see issue #309
	}
	return nil
}
