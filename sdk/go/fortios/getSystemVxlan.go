// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system vxlan
func LookupSystemVxlan(ctx *pulumi.Context, args *LookupSystemVxlanArgs, opts ...pulumi.InvokeOption) (*LookupSystemVxlanResult, error) {
	var rv LookupSystemVxlanResult
	err := ctx.Invoke("fortios:index/getSystemVxlan:GetSystemVxlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemVxlan.
type LookupSystemVxlanArgs struct {
	// Specify the name of the desired system vxlan.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemVxlan.
type LookupSystemVxlanResult struct {
	// VXLAN destination port (1 - 65535, default = 4789).
	Dstport int `pulumi:"dstport"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Outgoing interface for VXLAN encapsulated traffic.
	Interface string `pulumi:"interface"`
	// IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.
	IpVersion string `pulumi:"ipVersion"`
	// VXLAN multicast TTL (1-255, default = 0).
	MulticastTtl int `pulumi:"multicastTtl"`
	// VXLAN device or interface name. Must be a unique interface name.
	Name string `pulumi:"name"`
	// IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp6` block is documented below.
	RemoteIp6s []GetSystemVxlanRemoteIp6 `pulumi:"remoteIp6s"`
	// IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remoteIp` block is documented below.
	RemoteIps []GetSystemVxlanRemoteIp `pulumi:"remoteIps"`
	Vdomparam *string                  `pulumi:"vdomparam"`
	// VXLAN network ID.
	Vni int `pulumi:"vni"`
}