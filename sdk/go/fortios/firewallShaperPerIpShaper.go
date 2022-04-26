// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure per-IP traffic shaper.
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
// 		_, err := fortios.NewFirewallShaperPerIpShaper(ctx, "trname", &fortios.FirewallShaperPerIpShaperArgs{
// 			BandwidthUnit:        pulumi.String("kbps"),
// 			DiffservForward:      pulumi.String("disable"),
// 			DiffservReverse:      pulumi.String("disable"),
// 			DiffservcodeForward:  pulumi.String("000000"),
// 			DiffservcodeRev:      pulumi.String("000000"),
// 			MaxBandwidth:         pulumi.Int(1024),
// 			MaxConcurrentSession: pulumi.Int(33),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// FirewallShaper PerIpShaper can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallShaperPerIpShaper:FirewallShaperPerIpShaper labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallShaperPerIpShaper:FirewallShaperPerIpShaper labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallShaperPerIpShaper struct {
	pulumi.CustomResourceState

	// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
	BandwidthUnit pulumi.StringOutput `pulumi:"bandwidthUnit"`
	// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservForward pulumi.StringOutput `pulumi:"diffservForward"`
	// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservReverse pulumi.StringOutput `pulumi:"diffservReverse"`
	// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeForward pulumi.StringOutput `pulumi:"diffservcodeForward"`
	// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeRev pulumi.StringOutput `pulumi:"diffservcodeRev"`
	// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
	MaxBandwidth pulumi.IntOutput `pulumi:"maxBandwidth"`
	// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentSession pulumi.IntOutput `pulumi:"maxConcurrentSession"`
	// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentTcpSession pulumi.IntOutput `pulumi:"maxConcurrentTcpSession"`
	// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentUdpSession pulumi.IntOutput `pulumi:"maxConcurrentUdpSession"`
	// Traffic shaper name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallShaperPerIpShaper registers a new resource with the given unique name, arguments, and options.
func NewFirewallShaperPerIpShaper(ctx *pulumi.Context,
	name string, args *FirewallShaperPerIpShaperArgs, opts ...pulumi.ResourceOption) (*FirewallShaperPerIpShaper, error) {
	if args == nil {
		args = &FirewallShaperPerIpShaperArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallShaperPerIpShaper
	err := ctx.RegisterResource("fortios:index/firewallShaperPerIpShaper:FirewallShaperPerIpShaper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallShaperPerIpShaper gets an existing FirewallShaperPerIpShaper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallShaperPerIpShaper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallShaperPerIpShaperState, opts ...pulumi.ResourceOption) (*FirewallShaperPerIpShaper, error) {
	var resource FirewallShaperPerIpShaper
	err := ctx.ReadResource("fortios:index/firewallShaperPerIpShaper:FirewallShaperPerIpShaper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallShaperPerIpShaper resources.
type firewallShaperPerIpShaperState struct {
	// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
	BandwidthUnit *string `pulumi:"bandwidthUnit"`
	// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservForward *string `pulumi:"diffservForward"`
	// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservReverse *string `pulumi:"diffservReverse"`
	// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeForward *string `pulumi:"diffservcodeForward"`
	// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeRev *string `pulumi:"diffservcodeRev"`
	// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
	MaxBandwidth *int `pulumi:"maxBandwidth"`
	// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentSession *int `pulumi:"maxConcurrentSession"`
	// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentTcpSession *int `pulumi:"maxConcurrentTcpSession"`
	// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentUdpSession *int `pulumi:"maxConcurrentUdpSession"`
	// Traffic shaper name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallShaperPerIpShaperState struct {
	// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
	BandwidthUnit pulumi.StringPtrInput
	// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservForward pulumi.StringPtrInput
	// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservReverse pulumi.StringPtrInput
	// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeForward pulumi.StringPtrInput
	// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeRev pulumi.StringPtrInput
	// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
	MaxBandwidth pulumi.IntPtrInput
	// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentSession pulumi.IntPtrInput
	// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentTcpSession pulumi.IntPtrInput
	// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentUdpSession pulumi.IntPtrInput
	// Traffic shaper name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallShaperPerIpShaperState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShaperPerIpShaperState)(nil)).Elem()
}

type firewallShaperPerIpShaperArgs struct {
	// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
	BandwidthUnit *string `pulumi:"bandwidthUnit"`
	// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservForward *string `pulumi:"diffservForward"`
	// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservReverse *string `pulumi:"diffservReverse"`
	// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeForward *string `pulumi:"diffservcodeForward"`
	// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeRev *string `pulumi:"diffservcodeRev"`
	// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
	MaxBandwidth *int `pulumi:"maxBandwidth"`
	// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentSession *int `pulumi:"maxConcurrentSession"`
	// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentTcpSession *int `pulumi:"maxConcurrentTcpSession"`
	// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentUdpSession *int `pulumi:"maxConcurrentUdpSession"`
	// Traffic shaper name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallShaperPerIpShaper resource.
type FirewallShaperPerIpShaperArgs struct {
	// Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
	BandwidthUnit pulumi.StringPtrInput
	// Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservForward pulumi.StringPtrInput
	// Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
	DiffservReverse pulumi.StringPtrInput
	// Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeForward pulumi.StringPtrInput
	// Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
	DiffservcodeRev pulumi.StringPtrInput
	// Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
	MaxBandwidth pulumi.IntPtrInput
	// Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentSession pulumi.IntPtrInput
	// Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentTcpSession pulumi.IntPtrInput
	// Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
	MaxConcurrentUdpSession pulumi.IntPtrInput
	// Traffic shaper name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallShaperPerIpShaperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallShaperPerIpShaperArgs)(nil)).Elem()
}

type FirewallShaperPerIpShaperInput interface {
	pulumi.Input

	ToFirewallShaperPerIpShaperOutput() FirewallShaperPerIpShaperOutput
	ToFirewallShaperPerIpShaperOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperOutput
}

func (*FirewallShaperPerIpShaper) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShaperPerIpShaper)(nil)).Elem()
}

func (i *FirewallShaperPerIpShaper) ToFirewallShaperPerIpShaperOutput() FirewallShaperPerIpShaperOutput {
	return i.ToFirewallShaperPerIpShaperOutputWithContext(context.Background())
}

func (i *FirewallShaperPerIpShaper) ToFirewallShaperPerIpShaperOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperPerIpShaperOutput)
}

// FirewallShaperPerIpShaperArrayInput is an input type that accepts FirewallShaperPerIpShaperArray and FirewallShaperPerIpShaperArrayOutput values.
// You can construct a concrete instance of `FirewallShaperPerIpShaperArrayInput` via:
//
//          FirewallShaperPerIpShaperArray{ FirewallShaperPerIpShaperArgs{...} }
type FirewallShaperPerIpShaperArrayInput interface {
	pulumi.Input

	ToFirewallShaperPerIpShaperArrayOutput() FirewallShaperPerIpShaperArrayOutput
	ToFirewallShaperPerIpShaperArrayOutputWithContext(context.Context) FirewallShaperPerIpShaperArrayOutput
}

type FirewallShaperPerIpShaperArray []FirewallShaperPerIpShaperInput

func (FirewallShaperPerIpShaperArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShaperPerIpShaper)(nil)).Elem()
}

func (i FirewallShaperPerIpShaperArray) ToFirewallShaperPerIpShaperArrayOutput() FirewallShaperPerIpShaperArrayOutput {
	return i.ToFirewallShaperPerIpShaperArrayOutputWithContext(context.Background())
}

func (i FirewallShaperPerIpShaperArray) ToFirewallShaperPerIpShaperArrayOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperPerIpShaperArrayOutput)
}

// FirewallShaperPerIpShaperMapInput is an input type that accepts FirewallShaperPerIpShaperMap and FirewallShaperPerIpShaperMapOutput values.
// You can construct a concrete instance of `FirewallShaperPerIpShaperMapInput` via:
//
//          FirewallShaperPerIpShaperMap{ "key": FirewallShaperPerIpShaperArgs{...} }
type FirewallShaperPerIpShaperMapInput interface {
	pulumi.Input

	ToFirewallShaperPerIpShaperMapOutput() FirewallShaperPerIpShaperMapOutput
	ToFirewallShaperPerIpShaperMapOutputWithContext(context.Context) FirewallShaperPerIpShaperMapOutput
}

type FirewallShaperPerIpShaperMap map[string]FirewallShaperPerIpShaperInput

func (FirewallShaperPerIpShaperMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShaperPerIpShaper)(nil)).Elem()
}

func (i FirewallShaperPerIpShaperMap) ToFirewallShaperPerIpShaperMapOutput() FirewallShaperPerIpShaperMapOutput {
	return i.ToFirewallShaperPerIpShaperMapOutputWithContext(context.Background())
}

func (i FirewallShaperPerIpShaperMap) ToFirewallShaperPerIpShaperMapOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallShaperPerIpShaperMapOutput)
}

type FirewallShaperPerIpShaperOutput struct{ *pulumi.OutputState }

func (FirewallShaperPerIpShaperOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallShaperPerIpShaper)(nil)).Elem()
}

func (o FirewallShaperPerIpShaperOutput) ToFirewallShaperPerIpShaperOutput() FirewallShaperPerIpShaperOutput {
	return o
}

func (o FirewallShaperPerIpShaperOutput) ToFirewallShaperPerIpShaperOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperOutput {
	return o
}

type FirewallShaperPerIpShaperArrayOutput struct{ *pulumi.OutputState }

func (FirewallShaperPerIpShaperArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallShaperPerIpShaper)(nil)).Elem()
}

func (o FirewallShaperPerIpShaperArrayOutput) ToFirewallShaperPerIpShaperArrayOutput() FirewallShaperPerIpShaperArrayOutput {
	return o
}

func (o FirewallShaperPerIpShaperArrayOutput) ToFirewallShaperPerIpShaperArrayOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperArrayOutput {
	return o
}

func (o FirewallShaperPerIpShaperArrayOutput) Index(i pulumi.IntInput) FirewallShaperPerIpShaperOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallShaperPerIpShaper {
		return vs[0].([]*FirewallShaperPerIpShaper)[vs[1].(int)]
	}).(FirewallShaperPerIpShaperOutput)
}

type FirewallShaperPerIpShaperMapOutput struct{ *pulumi.OutputState }

func (FirewallShaperPerIpShaperMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallShaperPerIpShaper)(nil)).Elem()
}

func (o FirewallShaperPerIpShaperMapOutput) ToFirewallShaperPerIpShaperMapOutput() FirewallShaperPerIpShaperMapOutput {
	return o
}

func (o FirewallShaperPerIpShaperMapOutput) ToFirewallShaperPerIpShaperMapOutputWithContext(ctx context.Context) FirewallShaperPerIpShaperMapOutput {
	return o
}

func (o FirewallShaperPerIpShaperMapOutput) MapIndex(k pulumi.StringInput) FirewallShaperPerIpShaperOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallShaperPerIpShaper {
		return vs[0].(map[string]*FirewallShaperPerIpShaper)[vs[1].(string)]
	}).(FirewallShaperPerIpShaperOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperPerIpShaperInput)(nil)).Elem(), &FirewallShaperPerIpShaper{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperPerIpShaperArrayInput)(nil)).Elem(), FirewallShaperPerIpShaperArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallShaperPerIpShaperMapInput)(nil)).Elem(), FirewallShaperPerIpShaperMap{})
	pulumi.RegisterOutputType(FirewallShaperPerIpShaperOutput{})
	pulumi.RegisterOutputType(FirewallShaperPerIpShaperArrayOutput{})
	pulumi.RegisterOutputType(FirewallShaperPerIpShaperMapOutput{})
}
