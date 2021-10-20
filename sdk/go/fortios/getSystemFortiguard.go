// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system fortiguard
func LookupSystemFortiguard(ctx *pulumi.Context, args *LookupSystemFortiguardArgs, opts ...pulumi.InvokeOption) (*LookupSystemFortiguardResult, error) {
	var rv LookupSystemFortiguardResult
	err := ctx.Invoke("fortios:index/getSystemFortiguard:GetSystemFortiguard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemFortiguard.
type LookupSystemFortiguardArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemFortiguard.
type LookupSystemFortiguardResult struct {
	// Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance.
	AntispamCache string `pulumi:"antispamCache"`
	// Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).
	AntispamCacheMpercent int `pulumi:"antispamCacheMpercent"`
	// Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
	AntispamCacheTtl int `pulumi:"antispamCacheTtl"`
	// Expiration date of the FortiGuard antispam contract.
	AntispamExpiration int `pulumi:"antispamExpiration"`
	// Enable/disable turning off the FortiGuard antispam service.
	AntispamForceOff string `pulumi:"antispamForceOff"`
	// Interval of time between license checks for the FortiGuard antispam contract.
	AntispamLicense int `pulumi:"antispamLicense"`
	// Antispam query time out (1 - 30 sec, default = 7).
	AntispamTimeout int `pulumi:"antispamTimeout"`
	// IP address of the FortiGuard anycast DNS rating server.
	AnycastSdnsServerIp string `pulumi:"anycastSdnsServerIp"`
	// Port to connect to on the FortiGuard anycast DNS rating server.
	AnycastSdnsServerPort int `pulumi:"anycastSdnsServerPort"`
	// Automatically connect to and login to FortiCloud.
	AutoJoinForticloud string `pulumi:"autoJoinForticloud"`
	// IP address of the FortiDDNS server.
	DdnsServerIp string `pulumi:"ddnsServerIp"`
	// Port used to communicate with FortiDDNS servers.
	DdnsServerPort int `pulumi:"ddnsServerPort"`
	// Enable/disable use of FortiGuard's anycast network.
	FortiguardAnycast string `pulumi:"fortiguardAnycast"`
	// Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet.
	FortiguardAnycastSource string `pulumi:"fortiguardAnycastSource"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interface string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server.
	InterfaceSelectMethod string `pulumi:"interfaceSelectMethod"`
	// Number of servers to alternate between as first FortiGuard option.
	LoadBalanceServers int `pulumi:"loadBalanceServers"`
	// Enable/disable FortiGuard Virus Outbreak Prevention cache.
	OutbreakPreventionCache string `pulumi:"outbreakPreventionCache"`
	// Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
	OutbreakPreventionCacheMpercent int `pulumi:"outbreakPreventionCacheMpercent"`
	// Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
	OutbreakPreventionCacheTtl int `pulumi:"outbreakPreventionCacheTtl"`
	// Expiration date of FortiGuard Virus Outbreak Prevention contract.
	OutbreakPreventionExpiration int `pulumi:"outbreakPreventionExpiration"`
	// Turn off FortiGuard Virus Outbreak Prevention service.
	OutbreakPreventionForceOff string `pulumi:"outbreakPreventionForceOff"`
	// Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
	OutbreakPreventionLicense int `pulumi:"outbreakPreventionLicense"`
	// FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
	OutbreakPreventionTimeout int `pulumi:"outbreakPreventionTimeout"`
	// Port used to communicate with the FortiGuard servers.
	Port string `pulumi:"port"`
	// Protocol used to communicate with the FortiGuard servers.
	Protocol string `pulumi:"protocol"`
	// Proxy user password.
	ProxyPassword string `pulumi:"proxyPassword"`
	// IP address of the proxy server.
	ProxyServerIp string `pulumi:"proxyServerIp"`
	// Port used to communicate with the proxy server.
	ProxyServerPort int `pulumi:"proxyServerPort"`
	// Proxy user name.
	ProxyUsername string `pulumi:"proxyUsername"`
	// Cloud sandbox region.
	SandboxRegion string `pulumi:"sandboxRegion"`
	// Customization options for the FortiGuard DNS service.
	SdnsOptions string `pulumi:"sdnsOptions"`
	// IP address of the FortiDNS server.
	SdnsServerIp string `pulumi:"sdnsServerIp"`
	// Port used to communicate with FortiDNS servers.
	SdnsServerPort int `pulumi:"sdnsServerPort"`
	// Service account ID.
	ServiceAccountId string `pulumi:"serviceAccountId"`
	// Source IPv4 address used to communicate with FortiGuard.
	SourceIp string `pulumi:"sourceIp"`
	// Source IPv6 address used to communicate with FortiGuard.
	SourceIp6 string `pulumi:"sourceIp6"`
	// Signature update server location.
	UpdateServerLocation string  `pulumi:"updateServerLocation"`
	Vdomparam            *string `pulumi:"vdomparam"`
	// Enable/disable FortiGuard web filter caching.
	WebfilterCache string `pulumi:"webfilterCache"`
	// Time-to-live for web filter cache entries in seconds (300 - 86400).
	WebfilterCacheTtl int `pulumi:"webfilterCacheTtl"`
	// Expiration date of the FortiGuard web filter contract.
	WebfilterExpiration int `pulumi:"webfilterExpiration"`
	// Enable/disable turning off the FortiGuard web filtering service.
	WebfilterForceOff string `pulumi:"webfilterForceOff"`
	// Interval of time between license checks for the FortiGuard web filter contract.
	WebfilterLicense int `pulumi:"webfilterLicense"`
	// Web filter query time out (1 - 30 sec, default = 7).
	WebfilterTimeout int `pulumi:"webfilterTimeout"`
}