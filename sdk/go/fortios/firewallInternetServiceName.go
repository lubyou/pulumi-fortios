// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define internet service names. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// Firewall InternetServiceName can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceName:FirewallInternetServiceName labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceName:FirewallInternetServiceName labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetServiceName struct {
	pulumi.CustomResourceState

	// City ID.
	CityId pulumi.IntOutput `pulumi:"cityId"`
	// Country or Area ID.
	CountryId pulumi.IntOutput `pulumi:"countryId"`
	// Internet Service ID.
	InternetServiceId pulumi.IntOutput `pulumi:"internetServiceId"`
	// Internet Service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Region ID.
	RegionId pulumi.IntOutput `pulumi:"regionId"`
	// Internet Service name type. Valid values: `default`, `location`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceName registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceName(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceNameArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceName, error) {
	if args == nil {
		args = &FirewallInternetServiceNameArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetServiceName
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceName:FirewallInternetServiceName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceName gets an existing FirewallInternetServiceName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceNameState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceName, error) {
	var resource FirewallInternetServiceName
	err := ctx.ReadResource("fortios:index/firewallInternetServiceName:FirewallInternetServiceName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceName resources.
type firewallInternetServiceNameState struct {
	// City ID.
	CityId *int `pulumi:"cityId"`
	// Country or Area ID.
	CountryId *int `pulumi:"countryId"`
	// Internet Service ID.
	InternetServiceId *int `pulumi:"internetServiceId"`
	// Internet Service name.
	Name *string `pulumi:"name"`
	// Region ID.
	RegionId *int `pulumi:"regionId"`
	// Internet Service name type. Valid values: `default`, `location`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetServiceNameState struct {
	// City ID.
	CityId pulumi.IntPtrInput
	// Country or Area ID.
	CountryId pulumi.IntPtrInput
	// Internet Service ID.
	InternetServiceId pulumi.IntPtrInput
	// Internet Service name.
	Name pulumi.StringPtrInput
	// Region ID.
	RegionId pulumi.IntPtrInput
	// Internet Service name type. Valid values: `default`, `location`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceNameState)(nil)).Elem()
}

type firewallInternetServiceNameArgs struct {
	// City ID.
	CityId *int `pulumi:"cityId"`
	// Country or Area ID.
	CountryId *int `pulumi:"countryId"`
	// Internet Service ID.
	InternetServiceId *int `pulumi:"internetServiceId"`
	// Internet Service name.
	Name *string `pulumi:"name"`
	// Region ID.
	RegionId *int `pulumi:"regionId"`
	// Internet Service name type. Valid values: `default`, `location`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceName resource.
type FirewallInternetServiceNameArgs struct {
	// City ID.
	CityId pulumi.IntPtrInput
	// Country or Area ID.
	CountryId pulumi.IntPtrInput
	// Internet Service ID.
	InternetServiceId pulumi.IntPtrInput
	// Internet Service name.
	Name pulumi.StringPtrInput
	// Region ID.
	RegionId pulumi.IntPtrInput
	// Internet Service name type. Valid values: `default`, `location`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceNameArgs)(nil)).Elem()
}

type FirewallInternetServiceNameInput interface {
	pulumi.Input

	ToFirewallInternetServiceNameOutput() FirewallInternetServiceNameOutput
	ToFirewallInternetServiceNameOutputWithContext(ctx context.Context) FirewallInternetServiceNameOutput
}

func (*FirewallInternetServiceName) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceName)(nil)).Elem()
}

func (i *FirewallInternetServiceName) ToFirewallInternetServiceNameOutput() FirewallInternetServiceNameOutput {
	return i.ToFirewallInternetServiceNameOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceName) ToFirewallInternetServiceNameOutputWithContext(ctx context.Context) FirewallInternetServiceNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceNameOutput)
}

// FirewallInternetServiceNameArrayInput is an input type that accepts FirewallInternetServiceNameArray and FirewallInternetServiceNameArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceNameArrayInput` via:
//
//          FirewallInternetServiceNameArray{ FirewallInternetServiceNameArgs{...} }
type FirewallInternetServiceNameArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceNameArrayOutput() FirewallInternetServiceNameArrayOutput
	ToFirewallInternetServiceNameArrayOutputWithContext(context.Context) FirewallInternetServiceNameArrayOutput
}

type FirewallInternetServiceNameArray []FirewallInternetServiceNameInput

func (FirewallInternetServiceNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceName)(nil)).Elem()
}

func (i FirewallInternetServiceNameArray) ToFirewallInternetServiceNameArrayOutput() FirewallInternetServiceNameArrayOutput {
	return i.ToFirewallInternetServiceNameArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceNameArray) ToFirewallInternetServiceNameArrayOutputWithContext(ctx context.Context) FirewallInternetServiceNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceNameArrayOutput)
}

// FirewallInternetServiceNameMapInput is an input type that accepts FirewallInternetServiceNameMap and FirewallInternetServiceNameMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceNameMapInput` via:
//
//          FirewallInternetServiceNameMap{ "key": FirewallInternetServiceNameArgs{...} }
type FirewallInternetServiceNameMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceNameMapOutput() FirewallInternetServiceNameMapOutput
	ToFirewallInternetServiceNameMapOutputWithContext(context.Context) FirewallInternetServiceNameMapOutput
}

type FirewallInternetServiceNameMap map[string]FirewallInternetServiceNameInput

func (FirewallInternetServiceNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceName)(nil)).Elem()
}

func (i FirewallInternetServiceNameMap) ToFirewallInternetServiceNameMapOutput() FirewallInternetServiceNameMapOutput {
	return i.ToFirewallInternetServiceNameMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceNameMap) ToFirewallInternetServiceNameMapOutputWithContext(ctx context.Context) FirewallInternetServiceNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceNameMapOutput)
}

type FirewallInternetServiceNameOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceName)(nil)).Elem()
}

func (o FirewallInternetServiceNameOutput) ToFirewallInternetServiceNameOutput() FirewallInternetServiceNameOutput {
	return o
}

func (o FirewallInternetServiceNameOutput) ToFirewallInternetServiceNameOutputWithContext(ctx context.Context) FirewallInternetServiceNameOutput {
	return o
}

type FirewallInternetServiceNameArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceName)(nil)).Elem()
}

func (o FirewallInternetServiceNameArrayOutput) ToFirewallInternetServiceNameArrayOutput() FirewallInternetServiceNameArrayOutput {
	return o
}

func (o FirewallInternetServiceNameArrayOutput) ToFirewallInternetServiceNameArrayOutputWithContext(ctx context.Context) FirewallInternetServiceNameArrayOutput {
	return o
}

func (o FirewallInternetServiceNameArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetServiceName {
		return vs[0].([]*FirewallInternetServiceName)[vs[1].(int)]
	}).(FirewallInternetServiceNameOutput)
}

type FirewallInternetServiceNameMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceName)(nil)).Elem()
}

func (o FirewallInternetServiceNameMapOutput) ToFirewallInternetServiceNameMapOutput() FirewallInternetServiceNameMapOutput {
	return o
}

func (o FirewallInternetServiceNameMapOutput) ToFirewallInternetServiceNameMapOutputWithContext(ctx context.Context) FirewallInternetServiceNameMapOutput {
	return o
}

func (o FirewallInternetServiceNameMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetServiceName {
		return vs[0].(map[string]*FirewallInternetServiceName)[vs[1].(string)]
	}).(FirewallInternetServiceNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceNameInput)(nil)).Elem(), &FirewallInternetServiceName{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceNameArrayInput)(nil)).Elem(), FirewallInternetServiceNameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceNameMapInput)(nil)).Elem(), FirewallInternetServiceNameMap{})
	pulumi.RegisterOutputType(FirewallInternetServiceNameOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceNameArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceNameMapOutput{})
}
