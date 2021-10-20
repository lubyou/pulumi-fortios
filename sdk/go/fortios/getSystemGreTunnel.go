// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system gretunnel
func LookupSystemGreTunnel(ctx *pulumi.Context, args *LookupSystemGreTunnelArgs, opts ...pulumi.InvokeOption) (*LookupSystemGreTunnelResult, error) {
	var rv LookupSystemGreTunnelResult
	err := ctx.Invoke("fortios:index/getSystemGreTunnel:GetSystemGreTunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemGreTunnel.
type LookupSystemGreTunnelArgs struct {
	// Specify the name of the desired system gretunnel.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemGreTunnel.
type LookupSystemGreTunnelResult struct {
	// Enable/disable validating checksums in received GRE packets.
	ChecksumReception string `pulumi:"checksumReception"`
	// Enable/disable including checksums in transmitted GRE packets.
	ChecksumTransmission string `pulumi:"checksumTransmission"`
	// DiffServ setting to be applied to GRE tunnel outer IP header.
	Diffservcode string `pulumi:"diffservcode"`
	// Enable/disable DSCP copying.
	DscpCopying string `pulumi:"dscpCopying"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface name.
	Interface string `pulumi:"interface"`
	// IP version to use for VPN interface.
	IpVersion string `pulumi:"ipVersion"`
	// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
	KeepaliveFailtimes int `pulumi:"keepaliveFailtimes"`
	// Keepalive message interval (0 - 32767, 0 = disabled).
	KeepaliveInterval int `pulumi:"keepaliveInterval"`
	// Require received GRE packets contain this key (0 - 4294967295).
	KeyInbound int `pulumi:"keyInbound"`
	// Include this key in transmitted GRE packets (0 - 4294967295).
	KeyOutbound int `pulumi:"keyOutbound"`
	// IP address of the local gateway.
	LocalGw string `pulumi:"localGw"`
	// IPv6 address of the local gateway.
	LocalGw6 string `pulumi:"localGw6"`
	// Tunnel name.
	Name string `pulumi:"name"`
	// IP address of the remote gateway.
	RemoteGw string `pulumi:"remoteGw"`
	// IPv6 address of the remote gateway.
	RemoteGw6 string `pulumi:"remoteGw6"`
	// Enable/disable validating sequence numbers in received GRE packets.
	SequenceNumberReception string `pulumi:"sequenceNumberReception"`
	// Enable/disable including of sequence numbers in transmitted GRE packets.
	SequenceNumberTransmission string  `pulumi:"sequenceNumberTransmission"`
	Vdomparam                  *string `pulumi:"vdomparam"`
}