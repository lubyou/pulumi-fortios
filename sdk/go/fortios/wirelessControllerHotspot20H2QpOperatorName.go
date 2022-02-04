// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure operator friendly name.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewWirelessControllerHotspot20H2QpOperatorName(ctx, "trname", nil)
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
// WirelessControllerHotspot20 H2QpOperatorName can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20H2QpOperatorName:WirelessControllerHotspot20H2QpOperatorName labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerHotspot20H2QpOperatorName struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Friendly name ID.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name list. The structure of `valueList` block is documented below.
	ValueLists WirelessControllerHotspot20H2QpOperatorNameValueListArrayOutput `pulumi:"valueLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20H2QpOperatorName registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20H2QpOperatorName(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20H2QpOperatorNameArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20H2QpOperatorName, error) {
	if args == nil {
		args = &WirelessControllerHotspot20H2QpOperatorNameArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerHotspot20H2QpOperatorName
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20H2QpOperatorName:WirelessControllerHotspot20H2QpOperatorName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20H2QpOperatorName gets an existing WirelessControllerHotspot20H2QpOperatorName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20H2QpOperatorName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20H2QpOperatorNameState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20H2QpOperatorName, error) {
	var resource WirelessControllerHotspot20H2QpOperatorName
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20H2QpOperatorName:WirelessControllerHotspot20H2QpOperatorName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20H2QpOperatorName resources.
type wirelessControllerHotspot20H2QpOperatorNameState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Friendly name ID.
	Name *string `pulumi:"name"`
	// Name list. The structure of `valueList` block is documented below.
	ValueLists []WirelessControllerHotspot20H2QpOperatorNameValueList `pulumi:"valueLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20H2QpOperatorNameState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Friendly name ID.
	Name pulumi.StringPtrInput
	// Name list. The structure of `valueList` block is documented below.
	ValueLists WirelessControllerHotspot20H2QpOperatorNameValueListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20H2QpOperatorNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20H2QpOperatorNameState)(nil)).Elem()
}

type wirelessControllerHotspot20H2QpOperatorNameArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Friendly name ID.
	Name *string `pulumi:"name"`
	// Name list. The structure of `valueList` block is documented below.
	ValueLists []WirelessControllerHotspot20H2QpOperatorNameValueList `pulumi:"valueLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20H2QpOperatorName resource.
type WirelessControllerHotspot20H2QpOperatorNameArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Friendly name ID.
	Name pulumi.StringPtrInput
	// Name list. The structure of `valueList` block is documented below.
	ValueLists WirelessControllerHotspot20H2QpOperatorNameValueListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20H2QpOperatorNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20H2QpOperatorNameArgs)(nil)).Elem()
}

type WirelessControllerHotspot20H2QpOperatorNameInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpOperatorNameOutput() WirelessControllerHotspot20H2QpOperatorNameOutput
	ToWirelessControllerHotspot20H2QpOperatorNameOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOperatorNameOutput
}

func (*WirelessControllerHotspot20H2QpOperatorName) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20H2QpOperatorName)(nil)).Elem()
}

func (i *WirelessControllerHotspot20H2QpOperatorName) ToWirelessControllerHotspot20H2QpOperatorNameOutput() WirelessControllerHotspot20H2QpOperatorNameOutput {
	return i.ToWirelessControllerHotspot20H2QpOperatorNameOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20H2QpOperatorName) ToWirelessControllerHotspot20H2QpOperatorNameOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOperatorNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpOperatorNameOutput)
}

// WirelessControllerHotspot20H2QpOperatorNameArrayInput is an input type that accepts WirelessControllerHotspot20H2QpOperatorNameArray and WirelessControllerHotspot20H2QpOperatorNameArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20H2QpOperatorNameArrayInput` via:
//
//          WirelessControllerHotspot20H2QpOperatorNameArray{ WirelessControllerHotspot20H2QpOperatorNameArgs{...} }
type WirelessControllerHotspot20H2QpOperatorNameArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpOperatorNameArrayOutput() WirelessControllerHotspot20H2QpOperatorNameArrayOutput
	ToWirelessControllerHotspot20H2QpOperatorNameArrayOutputWithContext(context.Context) WirelessControllerHotspot20H2QpOperatorNameArrayOutput
}

type WirelessControllerHotspot20H2QpOperatorNameArray []WirelessControllerHotspot20H2QpOperatorNameInput

func (WirelessControllerHotspot20H2QpOperatorNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20H2QpOperatorName)(nil)).Elem()
}

func (i WirelessControllerHotspot20H2QpOperatorNameArray) ToWirelessControllerHotspot20H2QpOperatorNameArrayOutput() WirelessControllerHotspot20H2QpOperatorNameArrayOutput {
	return i.ToWirelessControllerHotspot20H2QpOperatorNameArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20H2QpOperatorNameArray) ToWirelessControllerHotspot20H2QpOperatorNameArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOperatorNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpOperatorNameArrayOutput)
}

// WirelessControllerHotspot20H2QpOperatorNameMapInput is an input type that accepts WirelessControllerHotspot20H2QpOperatorNameMap and WirelessControllerHotspot20H2QpOperatorNameMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20H2QpOperatorNameMapInput` via:
//
//          WirelessControllerHotspot20H2QpOperatorNameMap{ "key": WirelessControllerHotspot20H2QpOperatorNameArgs{...} }
type WirelessControllerHotspot20H2QpOperatorNameMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20H2QpOperatorNameMapOutput() WirelessControllerHotspot20H2QpOperatorNameMapOutput
	ToWirelessControllerHotspot20H2QpOperatorNameMapOutputWithContext(context.Context) WirelessControllerHotspot20H2QpOperatorNameMapOutput
}

type WirelessControllerHotspot20H2QpOperatorNameMap map[string]WirelessControllerHotspot20H2QpOperatorNameInput

func (WirelessControllerHotspot20H2QpOperatorNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20H2QpOperatorName)(nil)).Elem()
}

func (i WirelessControllerHotspot20H2QpOperatorNameMap) ToWirelessControllerHotspot20H2QpOperatorNameMapOutput() WirelessControllerHotspot20H2QpOperatorNameMapOutput {
	return i.ToWirelessControllerHotspot20H2QpOperatorNameMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20H2QpOperatorNameMap) ToWirelessControllerHotspot20H2QpOperatorNameMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOperatorNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20H2QpOperatorNameMapOutput)
}

type WirelessControllerHotspot20H2QpOperatorNameOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpOperatorNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20H2QpOperatorName)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpOperatorNameOutput) ToWirelessControllerHotspot20H2QpOperatorNameOutput() WirelessControllerHotspot20H2QpOperatorNameOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOperatorNameOutput) ToWirelessControllerHotspot20H2QpOperatorNameOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOperatorNameOutput {
	return o
}

type WirelessControllerHotspot20H2QpOperatorNameArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpOperatorNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20H2QpOperatorName)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpOperatorNameArrayOutput) ToWirelessControllerHotspot20H2QpOperatorNameArrayOutput() WirelessControllerHotspot20H2QpOperatorNameArrayOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOperatorNameArrayOutput) ToWirelessControllerHotspot20H2QpOperatorNameArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOperatorNameArrayOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOperatorNameArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20H2QpOperatorNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20H2QpOperatorName {
		return vs[0].([]*WirelessControllerHotspot20H2QpOperatorName)[vs[1].(int)]
	}).(WirelessControllerHotspot20H2QpOperatorNameOutput)
}

type WirelessControllerHotspot20H2QpOperatorNameMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20H2QpOperatorNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20H2QpOperatorName)(nil)).Elem()
}

func (o WirelessControllerHotspot20H2QpOperatorNameMapOutput) ToWirelessControllerHotspot20H2QpOperatorNameMapOutput() WirelessControllerHotspot20H2QpOperatorNameMapOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOperatorNameMapOutput) ToWirelessControllerHotspot20H2QpOperatorNameMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20H2QpOperatorNameMapOutput {
	return o
}

func (o WirelessControllerHotspot20H2QpOperatorNameMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20H2QpOperatorNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20H2QpOperatorName {
		return vs[0].(map[string]*WirelessControllerHotspot20H2QpOperatorName)[vs[1].(string)]
	}).(WirelessControllerHotspot20H2QpOperatorNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpOperatorNameInput)(nil)).Elem(), &WirelessControllerHotspot20H2QpOperatorName{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpOperatorNameArrayInput)(nil)).Elem(), WirelessControllerHotspot20H2QpOperatorNameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20H2QpOperatorNameMapInput)(nil)).Elem(), WirelessControllerHotspot20H2QpOperatorNameMap{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpOperatorNameOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpOperatorNameArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20H2QpOperatorNameMapOutput{})
}
