// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router static
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		sample1, err := fortios.GetRouterStatic(ctx, &fortios.LookupRouterStaticArgs{
// 			SeqNum: 1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("output1", sample1)
// 		return nil
// 	})
// }
// ```
func LookupRouterStatic(ctx *pulumi.Context, args *LookupRouterStaticArgs, opts ...pulumi.InvokeOption) (*LookupRouterStaticResult, error) {
	var rv LookupRouterStaticResult
	err := ctx.Invoke("fortios:index/getRouterStatic:GetRouterStatic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterStatic.
type LookupRouterStaticArgs struct {
	// Specify the seqNum of the desired router static.
	SeqNum int `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterStatic.
type LookupRouterStaticResult struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD).
	Bfd string `pulumi:"bfd"`
	// Enable/disable black hole.
	Blackhole string `pulumi:"blackhole"`
	// Optional comments.
	Comment string `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device string `pulumi:"device"`
	// Administrative distance (1 - 255).
	Distance int `pulumi:"distance"`
	// Destination IP and mask for this route.
	Dst string `pulumi:"dst"`
	// Name of firewall address or address group.
	Dstaddr string `pulumi:"dstaddr"`
	// Enable use of dynamic gateway retrieved from a DHCP or PPP server.
	DynamicGateway string `pulumi:"dynamicGateway"`
	// Gateway IP for this route.
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Application ID in the Internet service database.
	InternetService int `pulumi:"internetService"`
	// Application name in the Internet service custom database.
	InternetServiceCustom string `pulumi:"internetServiceCustom"`
	// Enable/disable withdrawing this route when link monitor or health check is down.
	LinkMonitorExempt string `pulumi:"linkMonitorExempt"`
	// Administrative priority (0 - 4294967295).
	Priority int `pulumi:"priority"`
	// Enable/disable egress through SD-WAN.
	Sdwan string `pulumi:"sdwan"`
	// Sequence number.
	SeqNum int `pulumi:"seqNum"`
	// Source prefix for this route.
	Src string `pulumi:"src"`
	// Enable/disable this static route.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link.
	VirtualWanLink string `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf int `pulumi:"vrf"`
	// Administrative weight (0 - 255).
	Weight int `pulumi:"weight"`
}