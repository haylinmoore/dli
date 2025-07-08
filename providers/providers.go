package providers

import (
	"fmt"

	"github.com/libdns/libdns"
)

// Provider interface that wraps libdns interfaces
type Provider interface {
	libdns.RecordGetter
	libdns.RecordSetter
	libdns.RecordDeleter
	libdns.RecordAppender
}

// GetProvider returns a provider instance based on the provider name
func GetProvider(providerName string) (Provider, error) {
	switch providerName {
	case "bunny":
		return getBunnyProvider()
	case "porkbun":
		return getPorkbunProvider()
	case "desec":
		return getDesecProvider()
	case "cloudflare":
		return getCloudflareProvider()
	case "googleclouddns":
		return getGoogleclouddnsProvider()
	case "metaname":
		return getMetanameProvider()
	case "cloudns":
		return getCloudnsProvider()
	case "netcup":
		return getNetcupProvider()
	case "he":
		return getHeProvider()
	case "dnsimple":
		return getDnsimpleProvider()
	case "westcn":
		return getWestcnProvider()
	case "digitalocean":
		return getDigitaloceanProvider()
	case "mijnhost":
		return getMijnhostProvider()
	case "luadns":
		return getLuadnsProvider()
	case "scaleway":
		return getScalewayProvider()
	case "hetzner":
		return getHetznerProvider()
	case "nfsn":
		return getNfsnProvider()
	case "inwx":
		return getInwxProvider()
	case "ovh":
		return getOvhProvider()
	case "mailinabox":
		return getMailinaboxProvider()
	case "glesys":
		return getGlesysProvider()
	case "loopia":
		return getLoopiaProvider()
	case "azure":
		return getAzureProvider()
	case "duckdns":
		return getDuckdnsProvider()
	case "dynu":
		return getDynuProvider()
	case "vultr":
		return getVultrProvider()
	default:
		return nil, fmt.Errorf("unsupported provider: %s", providerName)
	}
}
