// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 DoS policies.
//
// ## Import
//
// Firewall DosPolicy6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallDosPolicy6:FirewallDosPolicy6 labelname {{policyid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallDosPolicy6 struct {
	pulumi.CustomResourceState

	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDosPolicy6AnomalyArrayOutput `pulumi:"anomalies"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDosPolicy6DstaddrArrayOutput `pulumi:"dstaddrs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Anomaly name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy ID.
	Policyid pulumi.IntOutput `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDosPolicy6ServiceArrayOutput `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDosPolicy6SrcaddrArrayOutput `pulumi:"srcaddrs"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallDosPolicy6 registers a new resource with the given unique name, arguments, and options.
func NewFirewallDosPolicy6(ctx *pulumi.Context,
	name string, args *FirewallDosPolicy6Args, opts ...pulumi.ResourceOption) (*FirewallDosPolicy6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dstaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddrs'")
	}
	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Srcaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddrs'")
	}
	var resource FirewallDosPolicy6
	err := ctx.RegisterResource("fortios:index/firewallDosPolicy6:FirewallDosPolicy6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallDosPolicy6 gets an existing FirewallDosPolicy6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallDosPolicy6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallDosPolicy6State, opts ...pulumi.ResourceOption) (*FirewallDosPolicy6, error) {
	var resource FirewallDosPolicy6
	err := ctx.ReadResource("fortios:index/firewallDosPolicy6:FirewallDosPolicy6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallDosPolicy6 resources.
type firewallDosPolicy6State struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []FirewallDosPolicy6Anomaly `pulumi:"anomalies"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallDosPolicy6Dstaddr `pulumi:"dstaddrs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface *string `pulumi:"interface"`
	// Anomaly name.
	Name *string `pulumi:"name"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []FirewallDosPolicy6Service `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallDosPolicy6Srcaddr `pulumi:"srcaddrs"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallDosPolicy6State struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDosPolicy6AnomalyArrayInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDosPolicy6DstaddrArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Incoming interface name from available interfaces.
	Interface pulumi.StringPtrInput
	// Anomaly name.
	Name pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDosPolicy6ServiceArrayInput
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDosPolicy6SrcaddrArrayInput
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallDosPolicy6State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallDosPolicy6State)(nil)).Elem()
}

type firewallDosPolicy6Args struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies []FirewallDosPolicy6Anomaly `pulumi:"anomalies"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs []FirewallDosPolicy6Dstaddr `pulumi:"dstaddrs"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Incoming interface name from available interfaces.
	Interface string `pulumi:"interface"`
	// Anomaly name.
	Name *string `pulumi:"name"`
	// Policy ID.
	Policyid *int `pulumi:"policyid"`
	// Service object from available options. The structure of `service` block is documented below.
	Services []FirewallDosPolicy6Service `pulumi:"services"`
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs []FirewallDosPolicy6Srcaddr `pulumi:"srcaddrs"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallDosPolicy6 resource.
type FirewallDosPolicy6Args struct {
	// Anomaly name. The structure of `anomaly` block is documented below.
	Anomalies FirewallDosPolicy6AnomalyArrayInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
	Dstaddrs FirewallDosPolicy6DstaddrArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Incoming interface name from available interfaces.
	Interface pulumi.StringInput
	// Anomaly name.
	Name pulumi.StringPtrInput
	// Policy ID.
	Policyid pulumi.IntPtrInput
	// Service object from available options. The structure of `service` block is documented below.
	Services FirewallDosPolicy6ServiceArrayInput
	// Source address name from available addresses. The structure of `srcaddr` block is documented below.
	Srcaddrs FirewallDosPolicy6SrcaddrArrayInput
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallDosPolicy6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallDosPolicy6Args)(nil)).Elem()
}

type FirewallDosPolicy6Input interface {
	pulumi.Input

	ToFirewallDosPolicy6Output() FirewallDosPolicy6Output
	ToFirewallDosPolicy6OutputWithContext(ctx context.Context) FirewallDosPolicy6Output
}

func (*FirewallDosPolicy6) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallDosPolicy6)(nil))
}

func (i *FirewallDosPolicy6) ToFirewallDosPolicy6Output() FirewallDosPolicy6Output {
	return i.ToFirewallDosPolicy6OutputWithContext(context.Background())
}

func (i *FirewallDosPolicy6) ToFirewallDosPolicy6OutputWithContext(ctx context.Context) FirewallDosPolicy6Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDosPolicy6Output)
}

func (i *FirewallDosPolicy6) ToFirewallDosPolicy6PtrOutput() FirewallDosPolicy6PtrOutput {
	return i.ToFirewallDosPolicy6PtrOutputWithContext(context.Background())
}

func (i *FirewallDosPolicy6) ToFirewallDosPolicy6PtrOutputWithContext(ctx context.Context) FirewallDosPolicy6PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDosPolicy6PtrOutput)
}

type FirewallDosPolicy6PtrInput interface {
	pulumi.Input

	ToFirewallDosPolicy6PtrOutput() FirewallDosPolicy6PtrOutput
	ToFirewallDosPolicy6PtrOutputWithContext(ctx context.Context) FirewallDosPolicy6PtrOutput
}

type firewallDosPolicy6PtrType FirewallDosPolicy6Args

func (*firewallDosPolicy6PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallDosPolicy6)(nil))
}

func (i *firewallDosPolicy6PtrType) ToFirewallDosPolicy6PtrOutput() FirewallDosPolicy6PtrOutput {
	return i.ToFirewallDosPolicy6PtrOutputWithContext(context.Background())
}

func (i *firewallDosPolicy6PtrType) ToFirewallDosPolicy6PtrOutputWithContext(ctx context.Context) FirewallDosPolicy6PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDosPolicy6PtrOutput)
}

// FirewallDosPolicy6ArrayInput is an input type that accepts FirewallDosPolicy6Array and FirewallDosPolicy6ArrayOutput values.
// You can construct a concrete instance of `FirewallDosPolicy6ArrayInput` via:
//
//          FirewallDosPolicy6Array{ FirewallDosPolicy6Args{...} }
type FirewallDosPolicy6ArrayInput interface {
	pulumi.Input

	ToFirewallDosPolicy6ArrayOutput() FirewallDosPolicy6ArrayOutput
	ToFirewallDosPolicy6ArrayOutputWithContext(context.Context) FirewallDosPolicy6ArrayOutput
}

type FirewallDosPolicy6Array []FirewallDosPolicy6Input

func (FirewallDosPolicy6Array) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallDosPolicy6)(nil))
}

func (i FirewallDosPolicy6Array) ToFirewallDosPolicy6ArrayOutput() FirewallDosPolicy6ArrayOutput {
	return i.ToFirewallDosPolicy6ArrayOutputWithContext(context.Background())
}

func (i FirewallDosPolicy6Array) ToFirewallDosPolicy6ArrayOutputWithContext(ctx context.Context) FirewallDosPolicy6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDosPolicy6ArrayOutput)
}

// FirewallDosPolicy6MapInput is an input type that accepts FirewallDosPolicy6Map and FirewallDosPolicy6MapOutput values.
// You can construct a concrete instance of `FirewallDosPolicy6MapInput` via:
//
//          FirewallDosPolicy6Map{ "key": FirewallDosPolicy6Args{...} }
type FirewallDosPolicy6MapInput interface {
	pulumi.Input

	ToFirewallDosPolicy6MapOutput() FirewallDosPolicy6MapOutput
	ToFirewallDosPolicy6MapOutputWithContext(context.Context) FirewallDosPolicy6MapOutput
}

type FirewallDosPolicy6Map map[string]FirewallDosPolicy6Input

func (FirewallDosPolicy6Map) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallDosPolicy6)(nil))
}

func (i FirewallDosPolicy6Map) ToFirewallDosPolicy6MapOutput() FirewallDosPolicy6MapOutput {
	return i.ToFirewallDosPolicy6MapOutputWithContext(context.Background())
}

func (i FirewallDosPolicy6Map) ToFirewallDosPolicy6MapOutputWithContext(ctx context.Context) FirewallDosPolicy6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallDosPolicy6MapOutput)
}

type FirewallDosPolicy6Output struct {
	*pulumi.OutputState
}

func (FirewallDosPolicy6Output) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallDosPolicy6)(nil))
}

func (o FirewallDosPolicy6Output) ToFirewallDosPolicy6Output() FirewallDosPolicy6Output {
	return o
}

func (o FirewallDosPolicy6Output) ToFirewallDosPolicy6OutputWithContext(ctx context.Context) FirewallDosPolicy6Output {
	return o
}

func (o FirewallDosPolicy6Output) ToFirewallDosPolicy6PtrOutput() FirewallDosPolicy6PtrOutput {
	return o.ToFirewallDosPolicy6PtrOutputWithContext(context.Background())
}

func (o FirewallDosPolicy6Output) ToFirewallDosPolicy6PtrOutputWithContext(ctx context.Context) FirewallDosPolicy6PtrOutput {
	return o.ApplyT(func(v FirewallDosPolicy6) *FirewallDosPolicy6 {
		return &v
	}).(FirewallDosPolicy6PtrOutput)
}

type FirewallDosPolicy6PtrOutput struct {
	*pulumi.OutputState
}

func (FirewallDosPolicy6PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallDosPolicy6)(nil))
}

func (o FirewallDosPolicy6PtrOutput) ToFirewallDosPolicy6PtrOutput() FirewallDosPolicy6PtrOutput {
	return o
}

func (o FirewallDosPolicy6PtrOutput) ToFirewallDosPolicy6PtrOutputWithContext(ctx context.Context) FirewallDosPolicy6PtrOutput {
	return o
}

type FirewallDosPolicy6ArrayOutput struct{ *pulumi.OutputState }

func (FirewallDosPolicy6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallDosPolicy6)(nil))
}

func (o FirewallDosPolicy6ArrayOutput) ToFirewallDosPolicy6ArrayOutput() FirewallDosPolicy6ArrayOutput {
	return o
}

func (o FirewallDosPolicy6ArrayOutput) ToFirewallDosPolicy6ArrayOutputWithContext(ctx context.Context) FirewallDosPolicy6ArrayOutput {
	return o
}

func (o FirewallDosPolicy6ArrayOutput) Index(i pulumi.IntInput) FirewallDosPolicy6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallDosPolicy6 {
		return vs[0].([]FirewallDosPolicy6)[vs[1].(int)]
	}).(FirewallDosPolicy6Output)
}

type FirewallDosPolicy6MapOutput struct{ *pulumi.OutputState }

func (FirewallDosPolicy6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallDosPolicy6)(nil))
}

func (o FirewallDosPolicy6MapOutput) ToFirewallDosPolicy6MapOutput() FirewallDosPolicy6MapOutput {
	return o
}

func (o FirewallDosPolicy6MapOutput) ToFirewallDosPolicy6MapOutputWithContext(ctx context.Context) FirewallDosPolicy6MapOutput {
	return o
}

func (o FirewallDosPolicy6MapOutput) MapIndex(k pulumi.StringInput) FirewallDosPolicy6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallDosPolicy6 {
		return vs[0].(map[string]FirewallDosPolicy6)[vs[1].(string)]
	}).(FirewallDosPolicy6Output)
}

func init() {
	pulumi.RegisterOutputType(FirewallDosPolicy6Output{})
	pulumi.RegisterOutputType(FirewallDosPolicy6PtrOutput{})
	pulumi.RegisterOutputType(FirewallDosPolicy6ArrayOutput{})
	pulumi.RegisterOutputType(FirewallDosPolicy6MapOutput{})
}
