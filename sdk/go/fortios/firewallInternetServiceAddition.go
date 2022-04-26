// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Internet Services Addition. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Firewall InternetServiceAddition can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceAddition:FirewallInternetServiceAddition labelname {{fosid}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceAddition:FirewallInternetServiceAddition labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetServiceAddition struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Entries added to the Internet Service addition database. The structure of `entry` block is documented below.
	Entries FirewallInternetServiceAdditionEntryArrayOutput `pulumi:"entries"`
	// Internet Service ID in the Internet Service database.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceAddition registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceAddition(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceAdditionArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceAddition, error) {
	if args == nil {
		args = &FirewallInternetServiceAdditionArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetServiceAddition
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceAddition:FirewallInternetServiceAddition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceAddition gets an existing FirewallInternetServiceAddition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceAddition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceAdditionState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceAddition, error) {
	var resource FirewallInternetServiceAddition
	err := ctx.ReadResource("fortios:index/firewallInternetServiceAddition:FirewallInternetServiceAddition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceAddition resources.
type firewallInternetServiceAdditionState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Entries added to the Internet Service addition database. The structure of `entry` block is documented below.
	Entries []FirewallInternetServiceAdditionEntry `pulumi:"entries"`
	// Internet Service ID in the Internet Service database.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetServiceAdditionState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Entries added to the Internet Service addition database. The structure of `entry` block is documented below.
	Entries FirewallInternetServiceAdditionEntryArrayInput
	// Internet Service ID in the Internet Service database.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceAdditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceAdditionState)(nil)).Elem()
}

type firewallInternetServiceAdditionArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Entries added to the Internet Service addition database. The structure of `entry` block is documented below.
	Entries []FirewallInternetServiceAdditionEntry `pulumi:"entries"`
	// Internet Service ID in the Internet Service database.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceAddition resource.
type FirewallInternetServiceAdditionArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Entries added to the Internet Service addition database. The structure of `entry` block is documented below.
	Entries FirewallInternetServiceAdditionEntryArrayInput
	// Internet Service ID in the Internet Service database.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceAdditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceAdditionArgs)(nil)).Elem()
}

type FirewallInternetServiceAdditionInput interface {
	pulumi.Input

	ToFirewallInternetServiceAdditionOutput() FirewallInternetServiceAdditionOutput
	ToFirewallInternetServiceAdditionOutputWithContext(ctx context.Context) FirewallInternetServiceAdditionOutput
}

func (*FirewallInternetServiceAddition) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceAddition)(nil)).Elem()
}

func (i *FirewallInternetServiceAddition) ToFirewallInternetServiceAdditionOutput() FirewallInternetServiceAdditionOutput {
	return i.ToFirewallInternetServiceAdditionOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceAddition) ToFirewallInternetServiceAdditionOutputWithContext(ctx context.Context) FirewallInternetServiceAdditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceAdditionOutput)
}

// FirewallInternetServiceAdditionArrayInput is an input type that accepts FirewallInternetServiceAdditionArray and FirewallInternetServiceAdditionArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceAdditionArrayInput` via:
//
//          FirewallInternetServiceAdditionArray{ FirewallInternetServiceAdditionArgs{...} }
type FirewallInternetServiceAdditionArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceAdditionArrayOutput() FirewallInternetServiceAdditionArrayOutput
	ToFirewallInternetServiceAdditionArrayOutputWithContext(context.Context) FirewallInternetServiceAdditionArrayOutput
}

type FirewallInternetServiceAdditionArray []FirewallInternetServiceAdditionInput

func (FirewallInternetServiceAdditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceAddition)(nil)).Elem()
}

func (i FirewallInternetServiceAdditionArray) ToFirewallInternetServiceAdditionArrayOutput() FirewallInternetServiceAdditionArrayOutput {
	return i.ToFirewallInternetServiceAdditionArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceAdditionArray) ToFirewallInternetServiceAdditionArrayOutputWithContext(ctx context.Context) FirewallInternetServiceAdditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceAdditionArrayOutput)
}

// FirewallInternetServiceAdditionMapInput is an input type that accepts FirewallInternetServiceAdditionMap and FirewallInternetServiceAdditionMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceAdditionMapInput` via:
//
//          FirewallInternetServiceAdditionMap{ "key": FirewallInternetServiceAdditionArgs{...} }
type FirewallInternetServiceAdditionMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceAdditionMapOutput() FirewallInternetServiceAdditionMapOutput
	ToFirewallInternetServiceAdditionMapOutputWithContext(context.Context) FirewallInternetServiceAdditionMapOutput
}

type FirewallInternetServiceAdditionMap map[string]FirewallInternetServiceAdditionInput

func (FirewallInternetServiceAdditionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceAddition)(nil)).Elem()
}

func (i FirewallInternetServiceAdditionMap) ToFirewallInternetServiceAdditionMapOutput() FirewallInternetServiceAdditionMapOutput {
	return i.ToFirewallInternetServiceAdditionMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceAdditionMap) ToFirewallInternetServiceAdditionMapOutputWithContext(ctx context.Context) FirewallInternetServiceAdditionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceAdditionMapOutput)
}

type FirewallInternetServiceAdditionOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceAdditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceAddition)(nil)).Elem()
}

func (o FirewallInternetServiceAdditionOutput) ToFirewallInternetServiceAdditionOutput() FirewallInternetServiceAdditionOutput {
	return o
}

func (o FirewallInternetServiceAdditionOutput) ToFirewallInternetServiceAdditionOutputWithContext(ctx context.Context) FirewallInternetServiceAdditionOutput {
	return o
}

type FirewallInternetServiceAdditionArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceAdditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceAddition)(nil)).Elem()
}

func (o FirewallInternetServiceAdditionArrayOutput) ToFirewallInternetServiceAdditionArrayOutput() FirewallInternetServiceAdditionArrayOutput {
	return o
}

func (o FirewallInternetServiceAdditionArrayOutput) ToFirewallInternetServiceAdditionArrayOutputWithContext(ctx context.Context) FirewallInternetServiceAdditionArrayOutput {
	return o
}

func (o FirewallInternetServiceAdditionArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceAdditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetServiceAddition {
		return vs[0].([]*FirewallInternetServiceAddition)[vs[1].(int)]
	}).(FirewallInternetServiceAdditionOutput)
}

type FirewallInternetServiceAdditionMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceAdditionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceAddition)(nil)).Elem()
}

func (o FirewallInternetServiceAdditionMapOutput) ToFirewallInternetServiceAdditionMapOutput() FirewallInternetServiceAdditionMapOutput {
	return o
}

func (o FirewallInternetServiceAdditionMapOutput) ToFirewallInternetServiceAdditionMapOutputWithContext(ctx context.Context) FirewallInternetServiceAdditionMapOutput {
	return o
}

func (o FirewallInternetServiceAdditionMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceAdditionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetServiceAddition {
		return vs[0].(map[string]*FirewallInternetServiceAddition)[vs[1].(string)]
	}).(FirewallInternetServiceAdditionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceAdditionInput)(nil)).Elem(), &FirewallInternetServiceAddition{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceAdditionArrayInput)(nil)).Elem(), FirewallInternetServiceAdditionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceAdditionMapInput)(nil)).Elem(), FirewallInternetServiceAdditionMap{})
	pulumi.RegisterOutputType(FirewallInternetServiceAdditionOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceAdditionArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceAdditionMapOutput{})
}
