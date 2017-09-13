package caddyhttp

import (
	// plug in the server
	_ "github.com/WedgeServer/wedge/caddyhttp/httpserver"

	// plug in the standard directives
	_ "github.com/WedgeServer/wedge/caddyhttp/basicauth"
	_ "github.com/WedgeServer/wedge/caddyhttp/bind"
	_ "github.com/WedgeServer/wedge/caddyhttp/browse"
	_ "github.com/WedgeServer/wedge/caddyhttp/errors"
	_ "github.com/WedgeServer/wedge/caddyhttp/expvar"
	_ "github.com/WedgeServer/wedge/caddyhttp/extensions"
	_ "github.com/WedgeServer/wedge/caddyhttp/fastcgi"
	_ "github.com/WedgeServer/wedge/caddyhttp/gzip"
	_ "github.com/WedgeServer/wedge/caddyhttp/header"
	_ "github.com/WedgeServer/wedge/caddyhttp/index"
	_ "github.com/WedgeServer/wedge/caddyhttp/internalsrv"
	_ "github.com/WedgeServer/wedge/caddyhttp/limits"
	_ "github.com/WedgeServer/wedge/caddyhttp/log"
	_ "github.com/WedgeServer/wedge/caddyhttp/markdown"
	_ "github.com/WedgeServer/wedge/caddyhttp/mime"
	_ "github.com/WedgeServer/wedge/caddyhttp/pprof"
	_ "github.com/WedgeServer/wedge/caddyhttp/proxy"
	_ "github.com/WedgeServer/wedge/caddyhttp/push"
	_ "github.com/WedgeServer/wedge/caddyhttp/redirect"
	_ "github.com/WedgeServer/wedge/caddyhttp/requestid"
	_ "github.com/WedgeServer/wedge/caddyhttp/rewrite"
	_ "github.com/WedgeServer/wedge/caddyhttp/root"
	_ "github.com/WedgeServer/wedge/caddyhttp/status"
	_ "github.com/WedgeServer/wedge/caddyhttp/templates"
	_ "github.com/WedgeServer/wedge/caddyhttp/timeouts"
	_ "github.com/WedgeServer/wedge/caddyhttp/websocket"
	_ "github.com/WedgeServer/wedge/startupshutdown"
)
