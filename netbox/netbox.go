package netbox

import (
	"fmt"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/client"
)

// NewNetboxAt returns a client which will connect to the given
// hostname, which can optionally include a port, e.g. localhost:8000
func NewNetboxAt(host string) *client.NetBox {
	t := client.DefaultTransportConfig().WithHost(host)
	return client.NewHTTPClientWithConfig(strfmt.Default, t)
}

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

// NewNetboxWithAPIKey returna client which will connect to the given
// hostname (and optionally port), and will set the expected Authorization
// header on each request
func NewNetboxWithAPIKey(host string, apiToken string) *client.NetBox {
	t := runtimeclient.New(host, client.DefaultBasePath, client.DefaultSchemes)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	return client.New(t, strfmt.Default)
}
