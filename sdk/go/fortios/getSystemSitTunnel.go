// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system sittunnel
func LookupSystemSitTunnel(ctx *pulumi.Context, args *LookupSystemSitTunnelArgs, opts ...pulumi.InvokeOption) (*LookupSystemSitTunnelResult, error) {
	var rv LookupSystemSitTunnelResult
	err := ctx.Invoke("fortios:index/getSystemSitTunnel:GetSystemSitTunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemSitTunnel.
type LookupSystemSitTunnelArgs struct {
	// Specify the name of the desired system sittunnel.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemSitTunnel.
type LookupSystemSitTunnelResult struct {
	// Destination IP address of the tunnel.
	Destination string `pulumi:"destination"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface name.
	Interface string `pulumi:"interface"`
	// IPv6 address of the tunnel.
	Ip6 string `pulumi:"ip6"`
	// Tunnel name.
	Name string `pulumi:"name"`
	// Source IP address of the tunnel.
	Source    string  `pulumi:"source"`
	Vdomparam *string `pulumi:"vdomparam"`
}