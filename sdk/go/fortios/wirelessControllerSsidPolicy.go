// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure WiFi SSID policies. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// WirelessController SsidPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerSsidPolicy:WirelessControllerSsidPolicy labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerSsidPolicy:WirelessControllerSsidPolicy labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerSsidPolicy struct {
	pulumi.CustomResourceState

	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// VLAN interface name.
	Vlan pulumi.StringOutput `pulumi:"vlan"`
}

// NewWirelessControllerSsidPolicy registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerSsidPolicy(ctx *pulumi.Context,
	name string, args *WirelessControllerSsidPolicyArgs, opts ...pulumi.ResourceOption) (*WirelessControllerSsidPolicy, error) {
	if args == nil {
		args = &WirelessControllerSsidPolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerSsidPolicy
	err := ctx.RegisterResource("fortios:index/wirelessControllerSsidPolicy:WirelessControllerSsidPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerSsidPolicy gets an existing WirelessControllerSsidPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerSsidPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerSsidPolicyState, opts ...pulumi.ResourceOption) (*WirelessControllerSsidPolicy, error) {
	var resource WirelessControllerSsidPolicy
	err := ctx.ReadResource("fortios:index/wirelessControllerSsidPolicy:WirelessControllerSsidPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerSsidPolicy resources.
type wirelessControllerSsidPolicyState struct {
	// Description.
	Description *string `pulumi:"description"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN interface name.
	Vlan *string `pulumi:"vlan"`
}

type WirelessControllerSsidPolicyState struct {
	// Description.
	Description pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN interface name.
	Vlan pulumi.StringPtrInput
}

func (WirelessControllerSsidPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerSsidPolicyState)(nil)).Elem()
}

type wirelessControllerSsidPolicyArgs struct {
	// Description.
	Description *string `pulumi:"description"`
	// Name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN interface name.
	Vlan *string `pulumi:"vlan"`
}

// The set of arguments for constructing a WirelessControllerSsidPolicy resource.
type WirelessControllerSsidPolicyArgs struct {
	// Description.
	Description pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// VLAN interface name.
	Vlan pulumi.StringPtrInput
}

func (WirelessControllerSsidPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerSsidPolicyArgs)(nil)).Elem()
}

type WirelessControllerSsidPolicyInput interface {
	pulumi.Input

	ToWirelessControllerSsidPolicyOutput() WirelessControllerSsidPolicyOutput
	ToWirelessControllerSsidPolicyOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyOutput
}

func (*WirelessControllerSsidPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerSsidPolicy)(nil)).Elem()
}

func (i *WirelessControllerSsidPolicy) ToWirelessControllerSsidPolicyOutput() WirelessControllerSsidPolicyOutput {
	return i.ToWirelessControllerSsidPolicyOutputWithContext(context.Background())
}

func (i *WirelessControllerSsidPolicy) ToWirelessControllerSsidPolicyOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSsidPolicyOutput)
}

// WirelessControllerSsidPolicyArrayInput is an input type that accepts WirelessControllerSsidPolicyArray and WirelessControllerSsidPolicyArrayOutput values.
// You can construct a concrete instance of `WirelessControllerSsidPolicyArrayInput` via:
//
//          WirelessControllerSsidPolicyArray{ WirelessControllerSsidPolicyArgs{...} }
type WirelessControllerSsidPolicyArrayInput interface {
	pulumi.Input

	ToWirelessControllerSsidPolicyArrayOutput() WirelessControllerSsidPolicyArrayOutput
	ToWirelessControllerSsidPolicyArrayOutputWithContext(context.Context) WirelessControllerSsidPolicyArrayOutput
}

type WirelessControllerSsidPolicyArray []WirelessControllerSsidPolicyInput

func (WirelessControllerSsidPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerSsidPolicy)(nil)).Elem()
}

func (i WirelessControllerSsidPolicyArray) ToWirelessControllerSsidPolicyArrayOutput() WirelessControllerSsidPolicyArrayOutput {
	return i.ToWirelessControllerSsidPolicyArrayOutputWithContext(context.Background())
}

func (i WirelessControllerSsidPolicyArray) ToWirelessControllerSsidPolicyArrayOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSsidPolicyArrayOutput)
}

// WirelessControllerSsidPolicyMapInput is an input type that accepts WirelessControllerSsidPolicyMap and WirelessControllerSsidPolicyMapOutput values.
// You can construct a concrete instance of `WirelessControllerSsidPolicyMapInput` via:
//
//          WirelessControllerSsidPolicyMap{ "key": WirelessControllerSsidPolicyArgs{...} }
type WirelessControllerSsidPolicyMapInput interface {
	pulumi.Input

	ToWirelessControllerSsidPolicyMapOutput() WirelessControllerSsidPolicyMapOutput
	ToWirelessControllerSsidPolicyMapOutputWithContext(context.Context) WirelessControllerSsidPolicyMapOutput
}

type WirelessControllerSsidPolicyMap map[string]WirelessControllerSsidPolicyInput

func (WirelessControllerSsidPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerSsidPolicy)(nil)).Elem()
}

func (i WirelessControllerSsidPolicyMap) ToWirelessControllerSsidPolicyMapOutput() WirelessControllerSsidPolicyMapOutput {
	return i.ToWirelessControllerSsidPolicyMapOutputWithContext(context.Background())
}

func (i WirelessControllerSsidPolicyMap) ToWirelessControllerSsidPolicyMapOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerSsidPolicyMapOutput)
}

type WirelessControllerSsidPolicyOutput struct{ *pulumi.OutputState }

func (WirelessControllerSsidPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerSsidPolicy)(nil)).Elem()
}

func (o WirelessControllerSsidPolicyOutput) ToWirelessControllerSsidPolicyOutput() WirelessControllerSsidPolicyOutput {
	return o
}

func (o WirelessControllerSsidPolicyOutput) ToWirelessControllerSsidPolicyOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyOutput {
	return o
}

type WirelessControllerSsidPolicyArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerSsidPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerSsidPolicy)(nil)).Elem()
}

func (o WirelessControllerSsidPolicyArrayOutput) ToWirelessControllerSsidPolicyArrayOutput() WirelessControllerSsidPolicyArrayOutput {
	return o
}

func (o WirelessControllerSsidPolicyArrayOutput) ToWirelessControllerSsidPolicyArrayOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyArrayOutput {
	return o
}

func (o WirelessControllerSsidPolicyArrayOutput) Index(i pulumi.IntInput) WirelessControllerSsidPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerSsidPolicy {
		return vs[0].([]*WirelessControllerSsidPolicy)[vs[1].(int)]
	}).(WirelessControllerSsidPolicyOutput)
}

type WirelessControllerSsidPolicyMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerSsidPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerSsidPolicy)(nil)).Elem()
}

func (o WirelessControllerSsidPolicyMapOutput) ToWirelessControllerSsidPolicyMapOutput() WirelessControllerSsidPolicyMapOutput {
	return o
}

func (o WirelessControllerSsidPolicyMapOutput) ToWirelessControllerSsidPolicyMapOutputWithContext(ctx context.Context) WirelessControllerSsidPolicyMapOutput {
	return o
}

func (o WirelessControllerSsidPolicyMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerSsidPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerSsidPolicy {
		return vs[0].(map[string]*WirelessControllerSsidPolicy)[vs[1].(string)]
	}).(WirelessControllerSsidPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerSsidPolicyInput)(nil)).Elem(), &WirelessControllerSsidPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerSsidPolicyArrayInput)(nil)).Elem(), WirelessControllerSsidPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerSsidPolicyMapInput)(nil)).Elem(), WirelessControllerSsidPolicyMap{})
	pulumi.RegisterOutputType(WirelessControllerSsidPolicyOutput{})
	pulumi.RegisterOutputType(WirelessControllerSsidPolicyArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerSsidPolicyMapOutput{})
}
