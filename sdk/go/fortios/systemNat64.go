// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure NAT64.
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
// 		_, err := fortios.NewSystemNat64(ctx, "trname", &fortios.SystemNat64Args{
// 			AlwaysSynthesizeAaaaRecord:     pulumi.String("enable"),
// 			GenerateIpv6FragmentHeader:     pulumi.String("disable"),
// 			Nat46ForceIpv4PacketForwarding: pulumi.String("disable"),
// 			Nat64Prefix:                    pulumi.String("2001:1:2:3::/96"),
// 			SecondaryPrefixStatus:          pulumi.String("disable"),
// 			Status:                         pulumi.String("disable"),
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
// System Nat64 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/systemNat64:SystemNat64 labelname SystemNat64
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemNat64 struct {
	pulumi.CustomResourceState

	// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
	AlwaysSynthesizeAaaaRecord pulumi.StringOutput `pulumi:"alwaysSynthesizeAaaaRecord"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
	GenerateIpv6FragmentHeader pulumi.StringOutput `pulumi:"generateIpv6FragmentHeader"`
	// Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
	Nat46ForceIpv4PacketForwarding pulumi.StringOutput `pulumi:"nat46ForceIpv4PacketForwarding"`
	// NAT64 prefix.
	Nat64Prefix pulumi.StringOutput `pulumi:"nat64Prefix"`
	// Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
	SecondaryPrefixStatus pulumi.StringOutput `pulumi:"secondaryPrefixStatus"`
	// Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
	SecondaryPrefixes SystemNat64SecondaryPrefixArrayOutput `pulumi:"secondaryPrefixes"`
	// Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemNat64 registers a new resource with the given unique name, arguments, and options.
func NewSystemNat64(ctx *pulumi.Context,
	name string, args *SystemNat64Args, opts ...pulumi.ResourceOption) (*SystemNat64, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Nat64Prefix == nil {
		return nil, errors.New("invalid value for required argument 'Nat64Prefix'")
	}
	var resource SystemNat64
	err := ctx.RegisterResource("fortios:index/systemNat64:SystemNat64", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemNat64 gets an existing SystemNat64 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemNat64(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemNat64State, opts ...pulumi.ResourceOption) (*SystemNat64, error) {
	var resource SystemNat64
	err := ctx.ReadResource("fortios:index/systemNat64:SystemNat64", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemNat64 resources.
type systemNat64State struct {
	// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
	AlwaysSynthesizeAaaaRecord *string `pulumi:"alwaysSynthesizeAaaaRecord"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
	GenerateIpv6FragmentHeader *string `pulumi:"generateIpv6FragmentHeader"`
	// Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
	Nat46ForceIpv4PacketForwarding *string `pulumi:"nat46ForceIpv4PacketForwarding"`
	// NAT64 prefix.
	Nat64Prefix *string `pulumi:"nat64Prefix"`
	// Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
	SecondaryPrefixStatus *string `pulumi:"secondaryPrefixStatus"`
	// Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
	SecondaryPrefixes []SystemNat64SecondaryPrefix `pulumi:"secondaryPrefixes"`
	// Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemNat64State struct {
	// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
	AlwaysSynthesizeAaaaRecord pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
	GenerateIpv6FragmentHeader pulumi.StringPtrInput
	// Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
	Nat46ForceIpv4PacketForwarding pulumi.StringPtrInput
	// NAT64 prefix.
	Nat64Prefix pulumi.StringPtrInput
	// Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
	SecondaryPrefixStatus pulumi.StringPtrInput
	// Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
	SecondaryPrefixes SystemNat64SecondaryPrefixArrayInput
	// Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemNat64State) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNat64State)(nil)).Elem()
}

type systemNat64Args struct {
	// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
	AlwaysSynthesizeAaaaRecord *string `pulumi:"alwaysSynthesizeAaaaRecord"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
	GenerateIpv6FragmentHeader *string `pulumi:"generateIpv6FragmentHeader"`
	// Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
	Nat46ForceIpv4PacketForwarding *string `pulumi:"nat46ForceIpv4PacketForwarding"`
	// NAT64 prefix.
	Nat64Prefix string `pulumi:"nat64Prefix"`
	// Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
	SecondaryPrefixStatus *string `pulumi:"secondaryPrefixStatus"`
	// Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
	SecondaryPrefixes []SystemNat64SecondaryPrefix `pulumi:"secondaryPrefixes"`
	// Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemNat64 resource.
type SystemNat64Args struct {
	// Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
	AlwaysSynthesizeAaaaRecord pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
	GenerateIpv6FragmentHeader pulumi.StringPtrInput
	// Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
	Nat46ForceIpv4PacketForwarding pulumi.StringPtrInput
	// NAT64 prefix.
	Nat64Prefix pulumi.StringInput
	// Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
	SecondaryPrefixStatus pulumi.StringPtrInput
	// Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
	SecondaryPrefixes SystemNat64SecondaryPrefixArrayInput
	// Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemNat64Args) ElementType() reflect.Type {
	return reflect.TypeOf((*systemNat64Args)(nil)).Elem()
}

type SystemNat64Input interface {
	pulumi.Input

	ToSystemNat64Output() SystemNat64Output
	ToSystemNat64OutputWithContext(ctx context.Context) SystemNat64Output
}

func (*SystemNat64) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemNat64)(nil))
}

func (i *SystemNat64) ToSystemNat64Output() SystemNat64Output {
	return i.ToSystemNat64OutputWithContext(context.Background())
}

func (i *SystemNat64) ToSystemNat64OutputWithContext(ctx context.Context) SystemNat64Output {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNat64Output)
}

func (i *SystemNat64) ToSystemNat64PtrOutput() SystemNat64PtrOutput {
	return i.ToSystemNat64PtrOutputWithContext(context.Background())
}

func (i *SystemNat64) ToSystemNat64PtrOutputWithContext(ctx context.Context) SystemNat64PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNat64PtrOutput)
}

type SystemNat64PtrInput interface {
	pulumi.Input

	ToSystemNat64PtrOutput() SystemNat64PtrOutput
	ToSystemNat64PtrOutputWithContext(ctx context.Context) SystemNat64PtrOutput
}

type systemNat64PtrType SystemNat64Args

func (*systemNat64PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNat64)(nil))
}

func (i *systemNat64PtrType) ToSystemNat64PtrOutput() SystemNat64PtrOutput {
	return i.ToSystemNat64PtrOutputWithContext(context.Background())
}

func (i *systemNat64PtrType) ToSystemNat64PtrOutputWithContext(ctx context.Context) SystemNat64PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNat64PtrOutput)
}

// SystemNat64ArrayInput is an input type that accepts SystemNat64Array and SystemNat64ArrayOutput values.
// You can construct a concrete instance of `SystemNat64ArrayInput` via:
//
//          SystemNat64Array{ SystemNat64Args{...} }
type SystemNat64ArrayInput interface {
	pulumi.Input

	ToSystemNat64ArrayOutput() SystemNat64ArrayOutput
	ToSystemNat64ArrayOutputWithContext(context.Context) SystemNat64ArrayOutput
}

type SystemNat64Array []SystemNat64Input

func (SystemNat64Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SystemNat64)(nil))
}

func (i SystemNat64Array) ToSystemNat64ArrayOutput() SystemNat64ArrayOutput {
	return i.ToSystemNat64ArrayOutputWithContext(context.Background())
}

func (i SystemNat64Array) ToSystemNat64ArrayOutputWithContext(ctx context.Context) SystemNat64ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNat64ArrayOutput)
}

// SystemNat64MapInput is an input type that accepts SystemNat64Map and SystemNat64MapOutput values.
// You can construct a concrete instance of `SystemNat64MapInput` via:
//
//          SystemNat64Map{ "key": SystemNat64Args{...} }
type SystemNat64MapInput interface {
	pulumi.Input

	ToSystemNat64MapOutput() SystemNat64MapOutput
	ToSystemNat64MapOutputWithContext(context.Context) SystemNat64MapOutput
}

type SystemNat64Map map[string]SystemNat64Input

func (SystemNat64Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SystemNat64)(nil))
}

func (i SystemNat64Map) ToSystemNat64MapOutput() SystemNat64MapOutput {
	return i.ToSystemNat64MapOutputWithContext(context.Background())
}

func (i SystemNat64Map) ToSystemNat64MapOutputWithContext(ctx context.Context) SystemNat64MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemNat64MapOutput)
}

type SystemNat64Output struct {
	*pulumi.OutputState
}

func (SystemNat64Output) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemNat64)(nil))
}

func (o SystemNat64Output) ToSystemNat64Output() SystemNat64Output {
	return o
}

func (o SystemNat64Output) ToSystemNat64OutputWithContext(ctx context.Context) SystemNat64Output {
	return o
}

func (o SystemNat64Output) ToSystemNat64PtrOutput() SystemNat64PtrOutput {
	return o.ToSystemNat64PtrOutputWithContext(context.Background())
}

func (o SystemNat64Output) ToSystemNat64PtrOutputWithContext(ctx context.Context) SystemNat64PtrOutput {
	return o.ApplyT(func(v SystemNat64) *SystemNat64 {
		return &v
	}).(SystemNat64PtrOutput)
}

type SystemNat64PtrOutput struct {
	*pulumi.OutputState
}

func (SystemNat64PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemNat64)(nil))
}

func (o SystemNat64PtrOutput) ToSystemNat64PtrOutput() SystemNat64PtrOutput {
	return o
}

func (o SystemNat64PtrOutput) ToSystemNat64PtrOutputWithContext(ctx context.Context) SystemNat64PtrOutput {
	return o
}

type SystemNat64ArrayOutput struct{ *pulumi.OutputState }

func (SystemNat64ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SystemNat64)(nil))
}

func (o SystemNat64ArrayOutput) ToSystemNat64ArrayOutput() SystemNat64ArrayOutput {
	return o
}

func (o SystemNat64ArrayOutput) ToSystemNat64ArrayOutputWithContext(ctx context.Context) SystemNat64ArrayOutput {
	return o
}

func (o SystemNat64ArrayOutput) Index(i pulumi.IntInput) SystemNat64Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SystemNat64 {
		return vs[0].([]SystemNat64)[vs[1].(int)]
	}).(SystemNat64Output)
}

type SystemNat64MapOutput struct{ *pulumi.OutputState }

func (SystemNat64MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SystemNat64)(nil))
}

func (o SystemNat64MapOutput) ToSystemNat64MapOutput() SystemNat64MapOutput {
	return o
}

func (o SystemNat64MapOutput) ToSystemNat64MapOutputWithContext(ctx context.Context) SystemNat64MapOutput {
	return o
}

func (o SystemNat64MapOutput) MapIndex(k pulumi.StringInput) SystemNat64Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SystemNat64 {
		return vs[0].(map[string]SystemNat64)[vs[1].(string)]
	}).(SystemNat64Output)
}

func init() {
	pulumi.RegisterOutputType(SystemNat64Output{})
	pulumi.RegisterOutputType(SystemNat64PtrOutput{})
	pulumi.RegisterOutputType(SystemNat64ArrayOutput{})
	pulumi.RegisterOutputType(SystemNat64MapOutput{})
}
