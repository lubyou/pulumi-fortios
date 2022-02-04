// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router static6
func LookupRouterStatic6(ctx *pulumi.Context, args *LookupRouterStatic6Args, opts ...pulumi.InvokeOption) (*LookupRouterStatic6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterStatic6Result
	err := ctx.Invoke("fortios:index/getRouterStatic6:GetRouterStatic6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterStatic6.
type LookupRouterStatic6Args struct {
	// Specify the seqNum of the desired router static6.
	SeqNum int `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterStatic6.
type LookupRouterStatic6Result struct {
	// Enable/disable Bidirectional Forwarding Detection (BFD).
	Bfd string `pulumi:"bfd"`
	// Enable/disable black hole.
	Blackhole string `pulumi:"blackhole"`
	// Optional comments.
	Comment string `pulumi:"comment"`
	// Gateway out interface or tunnel.
	Device string `pulumi:"device"`
	// Device index (0 - 4294967295).
	Devindex int `pulumi:"devindex"`
	// Administrative distance (1 - 255).
	Distance int `pulumi:"distance"`
	// Destination IPv6 prefix.
	Dst string `pulumi:"dst"`
	// Enable use of dynamic gateway retrieved from Router Advertisement (RA).
	DynamicGateway string `pulumi:"dynamicGateway"`
	// IPv6 address of the gateway.
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable withdrawal of this static route when link monitor or health check is down.
	LinkMonitorExempt string `pulumi:"linkMonitorExempt"`
	// Administrative priority (0 - 4294967295).
	Priority int `pulumi:"priority"`
	// Enable/disable egress through the SD-WAN.
	Sdwan string `pulumi:"sdwan"`
	// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
	SdwanZones []GetRouterStatic6SdwanZone `pulumi:"sdwanZones"`
	// Sequence number.
	SeqNum int `pulumi:"seqNum"`
	// Enable/disable this static route.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable egress through the virtual-wan-link.
	VirtualWanLink string `pulumi:"virtualWanLink"`
	// Virtual Routing Forwarding ID.
	Vrf int `pulumi:"vrf"`
}

func LookupRouterStatic6Output(ctx *pulumi.Context, args LookupRouterStatic6OutputArgs, opts ...pulumi.InvokeOption) LookupRouterStatic6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterStatic6Result, error) {
			args := v.(LookupRouterStatic6Args)
			r, err := LookupRouterStatic6(ctx, &args, opts...)
			return *r, err
		}).(LookupRouterStatic6ResultOutput)
}

// A collection of arguments for invoking GetRouterStatic6.
type LookupRouterStatic6OutputArgs struct {
	// Specify the seqNum of the desired router static6.
	SeqNum pulumi.IntInput `pulumi:"seqNum"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterStatic6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterStatic6Args)(nil)).Elem()
}

// A collection of values returned by GetRouterStatic6.
type LookupRouterStatic6ResultOutput struct{ *pulumi.OutputState }

func (LookupRouterStatic6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterStatic6Result)(nil)).Elem()
}

func (o LookupRouterStatic6ResultOutput) ToLookupRouterStatic6ResultOutput() LookupRouterStatic6ResultOutput {
	return o
}

func (o LookupRouterStatic6ResultOutput) ToLookupRouterStatic6ResultOutputWithContext(ctx context.Context) LookupRouterStatic6ResultOutput {
	return o
}

// Enable/disable Bidirectional Forwarding Detection (BFD).
func (o LookupRouterStatic6ResultOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Bfd }).(pulumi.StringOutput)
}

// Enable/disable black hole.
func (o LookupRouterStatic6ResultOutput) Blackhole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Blackhole }).(pulumi.StringOutput)
}

// Optional comments.
func (o LookupRouterStatic6ResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Comment }).(pulumi.StringOutput)
}

// Gateway out interface or tunnel.
func (o LookupRouterStatic6ResultOutput) Device() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Device }).(pulumi.StringOutput)
}

// Device index (0 - 4294967295).
func (o LookupRouterStatic6ResultOutput) Devindex() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) int { return v.Devindex }).(pulumi.IntOutput)
}

// Administrative distance (1 - 255).
func (o LookupRouterStatic6ResultOutput) Distance() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) int { return v.Distance }).(pulumi.IntOutput)
}

// Destination IPv6 prefix.
func (o LookupRouterStatic6ResultOutput) Dst() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Dst }).(pulumi.StringOutput)
}

// Enable use of dynamic gateway retrieved from Router Advertisement (RA).
func (o LookupRouterStatic6ResultOutput) DynamicGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.DynamicGateway }).(pulumi.StringOutput)
}

// IPv6 address of the gateway.
func (o LookupRouterStatic6ResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterStatic6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable withdrawal of this static route when link monitor or health check is down.
func (o LookupRouterStatic6ResultOutput) LinkMonitorExempt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.LinkMonitorExempt }).(pulumi.StringOutput)
}

// Administrative priority (0 - 4294967295).
func (o LookupRouterStatic6ResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) int { return v.Priority }).(pulumi.IntOutput)
}

// Enable/disable egress through the SD-WAN.
func (o LookupRouterStatic6ResultOutput) Sdwan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Sdwan }).(pulumi.StringOutput)
}

// Choose SD-WAN Zone. The structure of `sdwanZone` block is documented below.
func (o LookupRouterStatic6ResultOutput) SdwanZones() GetRouterStatic6SdwanZoneArrayOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) []GetRouterStatic6SdwanZone { return v.SdwanZones }).(GetRouterStatic6SdwanZoneArrayOutput)
}

// Sequence number.
func (o LookupRouterStatic6ResultOutput) SeqNum() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) int { return v.SeqNum }).(pulumi.IntOutput)
}

// Enable/disable this static route.
func (o LookupRouterStatic6ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupRouterStatic6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable egress through the virtual-wan-link.
func (o LookupRouterStatic6ResultOutput) VirtualWanLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) string { return v.VirtualWanLink }).(pulumi.StringOutput)
}

// Virtual Routing Forwarding ID.
func (o LookupRouterStatic6ResultOutput) Vrf() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterStatic6Result) int { return v.Vrf }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterStatic6ResultOutput{})
}
