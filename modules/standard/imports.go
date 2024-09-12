package standard

import (
	// standard Caddy modules
	_ "github.com/InnovacationSRL/caddy/v2/caddyconfig/caddyfile"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddyevents"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddyevents/eventsconfig"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddyfs"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddyhttp/standard"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddypki"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddypki/acmeserver"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddytls"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddytls/distributedstek"
	_ "github.com/InnovacationSRL/caddy/v2/modules/caddytls/standardstek"
	_ "github.com/InnovacationSRL/caddy/v2/modules/filestorage"
	_ "github.com/InnovacationSRL/caddy/v2/modules/logging"
	_ "github.com/InnovacationSRL/caddy/v2/modules/metrics"
)
