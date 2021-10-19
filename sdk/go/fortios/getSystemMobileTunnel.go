// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system mobiletunnel
func LookupSystemMobileTunnel(ctx *pulumi.Context, args *LookupSystemMobileTunnelArgs, opts ...pulumi.InvokeOption) (*LookupSystemMobileTunnelResult, error) {
	var rv LookupSystemMobileTunnelResult
	err := ctx.Invoke("fortios:index/getSystemMobileTunnel:GetSystemMobileTunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemMobileTunnel.
type LookupSystemMobileTunnelArgs struct {
	// Specify the name of the desired system mobiletunnel.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemMobileTunnel.
type LookupSystemMobileTunnelResult struct {
	// Hash Algorithm (Keyed MD5).
	HashAlgorithm string `pulumi:"hashAlgorithm"`
	// Home IP address (Format: xxx.xxx.xxx.xxx).
	HomeAddress string `pulumi:"homeAddress"`
	// IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
	HomeAgent string `pulumi:"homeAgent"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
	Lifetime int `pulumi:"lifetime"`
	// NEMO authentication key.
	NMhaeKey string `pulumi:"nMhaeKey"`
	// NEMO authentication key type (ascii or base64).
	NMhaeKeyType string `pulumi:"nMhaeKeyType"`
	// NEMO authentication SPI (default: 256).
	NMhaeSpi int `pulumi:"nMhaeSpi"`
	// Tunnel name.
	Name string `pulumi:"name"`
	// NEMO network configuration. The structure of `network` block is documented below.
	Networks []GetSystemMobileTunnelNetwork `pulumi:"networks"`
	// NMMO HA registration interval (5 - 300, default = 5).
	RegInterval int `pulumi:"regInterval"`
	// Maximum number of NMMO HA registration retries (1 to 30, default = 3).
	RegRetry int `pulumi:"regRetry"`
	// Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
	RenewInterval int `pulumi:"renewInterval"`
	// Select the associated interface name from available options.
	RoamingInterface string `pulumi:"roamingInterface"`
	// Enable/disable this mobile tunnel.
	Status string `pulumi:"status"`
	// NEMO tunnnel mode (GRE tunnel).
	TunnelMode string  `pulumi:"tunnelMode"`
	Vdomparam  *string `pulumi:"vdomparam"`
}
