// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Internet Services Extension.
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
// 		_, err := fortios.NewFirewallInternetServiceExtension(ctx, "trname", &fortios.FirewallInternetServiceExtensionArgs{
// 			Comment: pulumi.String("EIWE"),
// 			Fosid:   pulumi.Int(65536),
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
// Firewall InternetServiceExtension can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension labelname {{fosid}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetServiceExtension struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
	DisableEntries FirewallInternetServiceExtensionDisableEntryArrayOutput `pulumi:"disableEntries"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
	Entries FirewallInternetServiceExtensionEntryArrayOutput `pulumi:"entries"`
	// Internet Service ID in the Internet Service database.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceExtension registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceExtension(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceExtensionArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceExtension, error) {
	if args == nil {
		args = &FirewallInternetServiceExtensionArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetServiceExtension
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceExtension gets an existing FirewallInternetServiceExtension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceExtensionState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceExtension, error) {
	var resource FirewallInternetServiceExtension
	err := ctx.ReadResource("fortios:index/firewallInternetServiceExtension:FirewallInternetServiceExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceExtension resources.
type firewallInternetServiceExtensionState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
	DisableEntries []FirewallInternetServiceExtensionDisableEntry `pulumi:"disableEntries"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
	Entries []FirewallInternetServiceExtensionEntry `pulumi:"entries"`
	// Internet Service ID in the Internet Service database.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetServiceExtensionState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
	DisableEntries FirewallInternetServiceExtensionDisableEntryArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
	Entries FirewallInternetServiceExtensionEntryArrayInput
	// Internet Service ID in the Internet Service database.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceExtensionState)(nil)).Elem()
}

type firewallInternetServiceExtensionArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
	DisableEntries []FirewallInternetServiceExtensionDisableEntry `pulumi:"disableEntries"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
	Entries []FirewallInternetServiceExtensionEntry `pulumi:"entries"`
	// Internet Service ID in the Internet Service database.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceExtension resource.
type FirewallInternetServiceExtensionArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// Disable entries in the Internet Service database. The structure of `disableEntry` block is documented below.
	DisableEntries FirewallInternetServiceExtensionDisableEntryArrayInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
	Entries FirewallInternetServiceExtensionEntryArrayInput
	// Internet Service ID in the Internet Service database.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetServiceExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceExtensionArgs)(nil)).Elem()
}

type FirewallInternetServiceExtensionInput interface {
	pulumi.Input

	ToFirewallInternetServiceExtensionOutput() FirewallInternetServiceExtensionOutput
	ToFirewallInternetServiceExtensionOutputWithContext(ctx context.Context) FirewallInternetServiceExtensionOutput
}

func (*FirewallInternetServiceExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceExtension)(nil)).Elem()
}

func (i *FirewallInternetServiceExtension) ToFirewallInternetServiceExtensionOutput() FirewallInternetServiceExtensionOutput {
	return i.ToFirewallInternetServiceExtensionOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceExtension) ToFirewallInternetServiceExtensionOutputWithContext(ctx context.Context) FirewallInternetServiceExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceExtensionOutput)
}

// FirewallInternetServiceExtensionArrayInput is an input type that accepts FirewallInternetServiceExtensionArray and FirewallInternetServiceExtensionArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceExtensionArrayInput` via:
//
//          FirewallInternetServiceExtensionArray{ FirewallInternetServiceExtensionArgs{...} }
type FirewallInternetServiceExtensionArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceExtensionArrayOutput() FirewallInternetServiceExtensionArrayOutput
	ToFirewallInternetServiceExtensionArrayOutputWithContext(context.Context) FirewallInternetServiceExtensionArrayOutput
}

type FirewallInternetServiceExtensionArray []FirewallInternetServiceExtensionInput

func (FirewallInternetServiceExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceExtension)(nil)).Elem()
}

func (i FirewallInternetServiceExtensionArray) ToFirewallInternetServiceExtensionArrayOutput() FirewallInternetServiceExtensionArrayOutput {
	return i.ToFirewallInternetServiceExtensionArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceExtensionArray) ToFirewallInternetServiceExtensionArrayOutputWithContext(ctx context.Context) FirewallInternetServiceExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceExtensionArrayOutput)
}

// FirewallInternetServiceExtensionMapInput is an input type that accepts FirewallInternetServiceExtensionMap and FirewallInternetServiceExtensionMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceExtensionMapInput` via:
//
//          FirewallInternetServiceExtensionMap{ "key": FirewallInternetServiceExtensionArgs{...} }
type FirewallInternetServiceExtensionMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceExtensionMapOutput() FirewallInternetServiceExtensionMapOutput
	ToFirewallInternetServiceExtensionMapOutputWithContext(context.Context) FirewallInternetServiceExtensionMapOutput
}

type FirewallInternetServiceExtensionMap map[string]FirewallInternetServiceExtensionInput

func (FirewallInternetServiceExtensionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceExtension)(nil)).Elem()
}

func (i FirewallInternetServiceExtensionMap) ToFirewallInternetServiceExtensionMapOutput() FirewallInternetServiceExtensionMapOutput {
	return i.ToFirewallInternetServiceExtensionMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceExtensionMap) ToFirewallInternetServiceExtensionMapOutputWithContext(ctx context.Context) FirewallInternetServiceExtensionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceExtensionMapOutput)
}

type FirewallInternetServiceExtensionOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceExtension)(nil)).Elem()
}

func (o FirewallInternetServiceExtensionOutput) ToFirewallInternetServiceExtensionOutput() FirewallInternetServiceExtensionOutput {
	return o
}

func (o FirewallInternetServiceExtensionOutput) ToFirewallInternetServiceExtensionOutputWithContext(ctx context.Context) FirewallInternetServiceExtensionOutput {
	return o
}

type FirewallInternetServiceExtensionArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceExtension)(nil)).Elem()
}

func (o FirewallInternetServiceExtensionArrayOutput) ToFirewallInternetServiceExtensionArrayOutput() FirewallInternetServiceExtensionArrayOutput {
	return o
}

func (o FirewallInternetServiceExtensionArrayOutput) ToFirewallInternetServiceExtensionArrayOutputWithContext(ctx context.Context) FirewallInternetServiceExtensionArrayOutput {
	return o
}

func (o FirewallInternetServiceExtensionArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetServiceExtension {
		return vs[0].([]*FirewallInternetServiceExtension)[vs[1].(int)]
	}).(FirewallInternetServiceExtensionOutput)
}

type FirewallInternetServiceExtensionMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceExtensionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceExtension)(nil)).Elem()
}

func (o FirewallInternetServiceExtensionMapOutput) ToFirewallInternetServiceExtensionMapOutput() FirewallInternetServiceExtensionMapOutput {
	return o
}

func (o FirewallInternetServiceExtensionMapOutput) ToFirewallInternetServiceExtensionMapOutputWithContext(ctx context.Context) FirewallInternetServiceExtensionMapOutput {
	return o
}

func (o FirewallInternetServiceExtensionMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceExtensionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetServiceExtension {
		return vs[0].(map[string]*FirewallInternetServiceExtension)[vs[1].(string)]
	}).(FirewallInternetServiceExtensionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceExtensionInput)(nil)).Elem(), &FirewallInternetServiceExtension{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceExtensionArrayInput)(nil)).Elem(), FirewallInternetServiceExtensionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceExtensionMapInput)(nil)).Elem(), FirewallInternetServiceExtensionMap{})
	pulumi.RegisterOutputType(FirewallInternetServiceExtensionOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceExtensionArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceExtensionMapOutput{})
}
