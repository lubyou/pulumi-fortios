// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define city table. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// Firewall City can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/firewallCity:FirewallCity labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallCity struct {
	pulumi.CustomResourceState

	// City ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// City name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallCity registers a new resource with the given unique name, arguments, and options.
func NewFirewallCity(ctx *pulumi.Context,
	name string, args *FirewallCityArgs, opts ...pulumi.ResourceOption) (*FirewallCity, error) {
	if args == nil {
		args = &FirewallCityArgs{}
	}

	var resource FirewallCity
	err := ctx.RegisterResource("fortios:index/firewallCity:FirewallCity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallCity gets an existing FirewallCity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallCity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallCityState, opts ...pulumi.ResourceOption) (*FirewallCity, error) {
	var resource FirewallCity
	err := ctx.ReadResource("fortios:index/firewallCity:FirewallCity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallCity resources.
type firewallCityState struct {
	// City ID.
	Fosid *int `pulumi:"fosid"`
	// City name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallCityState struct {
	// City ID.
	Fosid pulumi.IntPtrInput
	// City name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallCityState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallCityState)(nil)).Elem()
}

type firewallCityArgs struct {
	// City ID.
	Fosid *int `pulumi:"fosid"`
	// City name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallCity resource.
type FirewallCityArgs struct {
	// City ID.
	Fosid pulumi.IntPtrInput
	// City name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallCityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallCityArgs)(nil)).Elem()
}

type FirewallCityInput interface {
	pulumi.Input

	ToFirewallCityOutput() FirewallCityOutput
	ToFirewallCityOutputWithContext(ctx context.Context) FirewallCityOutput
}

func (*FirewallCity) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallCity)(nil))
}

func (i *FirewallCity) ToFirewallCityOutput() FirewallCityOutput {
	return i.ToFirewallCityOutputWithContext(context.Background())
}

func (i *FirewallCity) ToFirewallCityOutputWithContext(ctx context.Context) FirewallCityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCityOutput)
}

func (i *FirewallCity) ToFirewallCityPtrOutput() FirewallCityPtrOutput {
	return i.ToFirewallCityPtrOutputWithContext(context.Background())
}

func (i *FirewallCity) ToFirewallCityPtrOutputWithContext(ctx context.Context) FirewallCityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCityPtrOutput)
}

type FirewallCityPtrInput interface {
	pulumi.Input

	ToFirewallCityPtrOutput() FirewallCityPtrOutput
	ToFirewallCityPtrOutputWithContext(ctx context.Context) FirewallCityPtrOutput
}

type firewallCityPtrType FirewallCityArgs

func (*firewallCityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallCity)(nil))
}

func (i *firewallCityPtrType) ToFirewallCityPtrOutput() FirewallCityPtrOutput {
	return i.ToFirewallCityPtrOutputWithContext(context.Background())
}

func (i *firewallCityPtrType) ToFirewallCityPtrOutputWithContext(ctx context.Context) FirewallCityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCityPtrOutput)
}

// FirewallCityArrayInput is an input type that accepts FirewallCityArray and FirewallCityArrayOutput values.
// You can construct a concrete instance of `FirewallCityArrayInput` via:
//
//          FirewallCityArray{ FirewallCityArgs{...} }
type FirewallCityArrayInput interface {
	pulumi.Input

	ToFirewallCityArrayOutput() FirewallCityArrayOutput
	ToFirewallCityArrayOutputWithContext(context.Context) FirewallCityArrayOutput
}

type FirewallCityArray []FirewallCityInput

func (FirewallCityArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallCity)(nil))
}

func (i FirewallCityArray) ToFirewallCityArrayOutput() FirewallCityArrayOutput {
	return i.ToFirewallCityArrayOutputWithContext(context.Background())
}

func (i FirewallCityArray) ToFirewallCityArrayOutputWithContext(ctx context.Context) FirewallCityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCityArrayOutput)
}

// FirewallCityMapInput is an input type that accepts FirewallCityMap and FirewallCityMapOutput values.
// You can construct a concrete instance of `FirewallCityMapInput` via:
//
//          FirewallCityMap{ "key": FirewallCityArgs{...} }
type FirewallCityMapInput interface {
	pulumi.Input

	ToFirewallCityMapOutput() FirewallCityMapOutput
	ToFirewallCityMapOutputWithContext(context.Context) FirewallCityMapOutput
}

type FirewallCityMap map[string]FirewallCityInput

func (FirewallCityMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallCity)(nil))
}

func (i FirewallCityMap) ToFirewallCityMapOutput() FirewallCityMapOutput {
	return i.ToFirewallCityMapOutputWithContext(context.Background())
}

func (i FirewallCityMap) ToFirewallCityMapOutputWithContext(ctx context.Context) FirewallCityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallCityMapOutput)
}

type FirewallCityOutput struct {
	*pulumi.OutputState
}

func (FirewallCityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallCity)(nil))
}

func (o FirewallCityOutput) ToFirewallCityOutput() FirewallCityOutput {
	return o
}

func (o FirewallCityOutput) ToFirewallCityOutputWithContext(ctx context.Context) FirewallCityOutput {
	return o
}

func (o FirewallCityOutput) ToFirewallCityPtrOutput() FirewallCityPtrOutput {
	return o.ToFirewallCityPtrOutputWithContext(context.Background())
}

func (o FirewallCityOutput) ToFirewallCityPtrOutputWithContext(ctx context.Context) FirewallCityPtrOutput {
	return o.ApplyT(func(v FirewallCity) *FirewallCity {
		return &v
	}).(FirewallCityPtrOutput)
}

type FirewallCityPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallCityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallCity)(nil))
}

func (o FirewallCityPtrOutput) ToFirewallCityPtrOutput() FirewallCityPtrOutput {
	return o
}

func (o FirewallCityPtrOutput) ToFirewallCityPtrOutputWithContext(ctx context.Context) FirewallCityPtrOutput {
	return o
}

type FirewallCityArrayOutput struct{ *pulumi.OutputState }

func (FirewallCityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallCity)(nil))
}

func (o FirewallCityArrayOutput) ToFirewallCityArrayOutput() FirewallCityArrayOutput {
	return o
}

func (o FirewallCityArrayOutput) ToFirewallCityArrayOutputWithContext(ctx context.Context) FirewallCityArrayOutput {
	return o
}

func (o FirewallCityArrayOutput) Index(i pulumi.IntInput) FirewallCityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallCity {
		return vs[0].([]FirewallCity)[vs[1].(int)]
	}).(FirewallCityOutput)
}

type FirewallCityMapOutput struct{ *pulumi.OutputState }

func (FirewallCityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallCity)(nil))
}

func (o FirewallCityMapOutput) ToFirewallCityMapOutput() FirewallCityMapOutput {
	return o
}

func (o FirewallCityMapOutput) ToFirewallCityMapOutputWithContext(ctx context.Context) FirewallCityMapOutput {
	return o
}

func (o FirewallCityMapOutput) MapIndex(k pulumi.StringInput) FirewallCityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallCity {
		return vs[0].(map[string]FirewallCity)[vs[1].(string)]
	}).(FirewallCityOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallCityOutput{})
	pulumi.RegisterOutputType(FirewallCityPtrOutput{})
	pulumi.RegisterOutputType(FirewallCityArrayOutput{})
	pulumi.RegisterOutputType(FirewallCityMapOutput{})
}
