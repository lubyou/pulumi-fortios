// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define country table. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// Firewall Country can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallCountry:FirewallCountry labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallCountry struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Country ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Country name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Region ID list. The structure of `region` block is documented below.
	Regions FirewallCountryRegionArrayOutput `pulumi:"regions"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallCountry registers a new resource with the given unique name, arguments, and options.
func NewFirewallCountry(ctx *pulumi.Context,
	name string, args *FirewallCountryArgs, opts ...pulumi.ResourceOption) (*FirewallCountry, error) {
	if args == nil {
		args = &FirewallCountryArgs{}
	}

	var resource FirewallCountry
	err := ctx.RegisterResource("fortios:index/firewallCountry:FirewallCountry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallCountry gets an existing FirewallCountry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallCountry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallCountryState, opts ...pulumi.ResourceOption) (*FirewallCountry, error) {
	var resource FirewallCountry
	err := ctx.ReadResource("fortios:index/firewallCountry:FirewallCountry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallCountry resources.
type firewallCountryState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Country ID.
	Fosid *int `pulumi:"fosid"`
	// Country name.
	Name *string `pulumi:"name"`
	// Region ID list. The structure of `region` block is documented below.
	Regions []FirewallCountryRegion `pulumi:"regions"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallCountryState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Country ID.
	Fosid pulumi.IntPtrInput
	// Country name.
	Name pulumi.StringPtrInput
	// Region ID list. The structure of `region` block is documented below.
	Regions FirewallCountryRegionArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallCountryState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallCountryState)(nil)).Elem()
}

type firewallCountryArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Country ID.
	Fosid *int `pulumi:"fosid"`
	// Country name.
	Name *string `pulumi:"name"`
	// Region ID list. The structure of `region` block is documented below.
	Regions []FirewallCountryRegion `pulumi:"regions"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallCountry resource.
type FirewallCountryArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Country ID.
	Fosid pulumi.IntPtrInput
	// Country name.
	Name pulumi.StringPtrInput
	// Region ID list. The structure of `region` block is documented below.
	Regions FirewallCountryRegionArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallCountryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallCountryArgs)(nil)).Elem()
}

type FirewallCountryInput interface {
	pulumi.Input

	ToFirewallCountryOutput() FirewallCountryOutput
	ToFirewallCountryOutputWithContext(ctx context.Context) FirewallCountryOutput
}

func (*FirewallCountry) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallCountry)(nil))
}

func (i *FirewallCountry) ToFirewallCountryOutput() FirewallCountryOutput {
	return i.ToFirewallCountryOutputWithContext(context.Background())
}

func (i *FirewallCountry) ToFirewallCountryOutputWithContext(ctx context.Context) FirewallCountryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCountryOutput)
}

func (i *FirewallCountry) ToFirewallCountryPtrOutput() FirewallCountryPtrOutput {
	return i.ToFirewallCountryPtrOutputWithContext(context.Background())
}

func (i *FirewallCountry) ToFirewallCountryPtrOutputWithContext(ctx context.Context) FirewallCountryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCountryPtrOutput)
}

type FirewallCountryPtrInput interface {
	pulumi.Input

	ToFirewallCountryPtrOutput() FirewallCountryPtrOutput
	ToFirewallCountryPtrOutputWithContext(ctx context.Context) FirewallCountryPtrOutput
}

type firewallCountryPtrType FirewallCountryArgs

func (*firewallCountryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallCountry)(nil))
}

func (i *firewallCountryPtrType) ToFirewallCountryPtrOutput() FirewallCountryPtrOutput {
	return i.ToFirewallCountryPtrOutputWithContext(context.Background())
}

func (i *firewallCountryPtrType) ToFirewallCountryPtrOutputWithContext(ctx context.Context) FirewallCountryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCountryPtrOutput)
}

// FirewallCountryArrayInput is an input type that accepts FirewallCountryArray and FirewallCountryArrayOutput values.
// You can construct a concrete instance of `FirewallCountryArrayInput` via:
//
//          FirewallCountryArray{ FirewallCountryArgs{...} }
type FirewallCountryArrayInput interface {
	pulumi.Input

	ToFirewallCountryArrayOutput() FirewallCountryArrayOutput
	ToFirewallCountryArrayOutputWithContext(context.Context) FirewallCountryArrayOutput
}

type FirewallCountryArray []FirewallCountryInput

func (FirewallCountryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallCountry)(nil))
}

func (i FirewallCountryArray) ToFirewallCountryArrayOutput() FirewallCountryArrayOutput {
	return i.ToFirewallCountryArrayOutputWithContext(context.Background())
}

func (i FirewallCountryArray) ToFirewallCountryArrayOutputWithContext(ctx context.Context) FirewallCountryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCountryArrayOutput)
}

// FirewallCountryMapInput is an input type that accepts FirewallCountryMap and FirewallCountryMapOutput values.
// You can construct a concrete instance of `FirewallCountryMapInput` via:
//
//          FirewallCountryMap{ "key": FirewallCountryArgs{...} }
type FirewallCountryMapInput interface {
	pulumi.Input

	ToFirewallCountryMapOutput() FirewallCountryMapOutput
	ToFirewallCountryMapOutputWithContext(context.Context) FirewallCountryMapOutput
}

type FirewallCountryMap map[string]FirewallCountryInput

func (FirewallCountryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallCountry)(nil))
}

func (i FirewallCountryMap) ToFirewallCountryMapOutput() FirewallCountryMapOutput {
	return i.ToFirewallCountryMapOutputWithContext(context.Background())
}

func (i FirewallCountryMap) ToFirewallCountryMapOutputWithContext(ctx context.Context) FirewallCountryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCountryMapOutput)
}

type FirewallCountryOutput struct {
	*pulumi.OutputState
}

func (FirewallCountryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallCountry)(nil))
}

func (o FirewallCountryOutput) ToFirewallCountryOutput() FirewallCountryOutput {
	return o
}

func (o FirewallCountryOutput) ToFirewallCountryOutputWithContext(ctx context.Context) FirewallCountryOutput {
	return o
}

func (o FirewallCountryOutput) ToFirewallCountryPtrOutput() FirewallCountryPtrOutput {
	return o.ToFirewallCountryPtrOutputWithContext(context.Background())
}

func (o FirewallCountryOutput) ToFirewallCountryPtrOutputWithContext(ctx context.Context) FirewallCountryPtrOutput {
	return o.ApplyT(func(v FirewallCountry) *FirewallCountry {
		return &v
	}).(FirewallCountryPtrOutput)
}

type FirewallCountryPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallCountryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallCountry)(nil))
}

func (o FirewallCountryPtrOutput) ToFirewallCountryPtrOutput() FirewallCountryPtrOutput {
	return o
}

func (o FirewallCountryPtrOutput) ToFirewallCountryPtrOutputWithContext(ctx context.Context) FirewallCountryPtrOutput {
	return o
}

type FirewallCountryArrayOutput struct{ *pulumi.OutputState }

func (FirewallCountryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallCountry)(nil))
}

func (o FirewallCountryArrayOutput) ToFirewallCountryArrayOutput() FirewallCountryArrayOutput {
	return o
}

func (o FirewallCountryArrayOutput) ToFirewallCountryArrayOutputWithContext(ctx context.Context) FirewallCountryArrayOutput {
	return o
}

func (o FirewallCountryArrayOutput) Index(i pulumi.IntInput) FirewallCountryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallCountry {
		return vs[0].([]FirewallCountry)[vs[1].(int)]
	}).(FirewallCountryOutput)
}

type FirewallCountryMapOutput struct{ *pulumi.OutputState }

func (FirewallCountryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallCountry)(nil))
}

func (o FirewallCountryMapOutput) ToFirewallCountryMapOutput() FirewallCountryMapOutput {
	return o
}

func (o FirewallCountryMapOutput) ToFirewallCountryMapOutputWithContext(ctx context.Context) FirewallCountryMapOutput {
	return o
}

func (o FirewallCountryMapOutput) MapIndex(k pulumi.StringInput) FirewallCountryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallCountry {
		return vs[0].(map[string]FirewallCountry)[vs[1].(string)]
	}).(FirewallCountryOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallCountryOutput{})
	pulumi.RegisterOutputType(FirewallCountryPtrOutput{})
	pulumi.RegisterOutputType(FirewallCountryArrayOutput{})
	pulumi.RegisterOutputType(FirewallCountryMapOutput{})
}