// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure OSU provider icon.
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
// 		_, err := fortios.NewWirelessControllerHotspot20Icon(ctx, "trname", nil)
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
// WirelessControllerHotspot20 Icon can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20Icon:WirelessControllerHotspot20Icon labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20Icon:WirelessControllerHotspot20Icon labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerHotspot20Icon struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Icon list. The structure of `iconList` block is documented below.
	IconLists WirelessControllerHotspot20IconIconListArrayOutput `pulumi:"iconLists"`
	// Icon name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20Icon registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20Icon(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20IconArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20Icon, error) {
	if args == nil {
		args = &WirelessControllerHotspot20IconArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerHotspot20Icon
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20Icon:WirelessControllerHotspot20Icon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20Icon gets an existing WirelessControllerHotspot20Icon resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20Icon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20IconState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20Icon, error) {
	var resource WirelessControllerHotspot20Icon
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20Icon:WirelessControllerHotspot20Icon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20Icon resources.
type wirelessControllerHotspot20IconState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Icon list. The structure of `iconList` block is documented below.
	IconLists []WirelessControllerHotspot20IconIconList `pulumi:"iconLists"`
	// Icon name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20IconState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Icon list. The structure of `iconList` block is documented below.
	IconLists WirelessControllerHotspot20IconIconListArrayInput
	// Icon name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20IconState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20IconState)(nil)).Elem()
}

type wirelessControllerHotspot20IconArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Icon list. The structure of `iconList` block is documented below.
	IconLists []WirelessControllerHotspot20IconIconList `pulumi:"iconLists"`
	// Icon name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20Icon resource.
type WirelessControllerHotspot20IconArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Icon list. The structure of `iconList` block is documented below.
	IconLists WirelessControllerHotspot20IconIconListArrayInput
	// Icon name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20IconArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20IconArgs)(nil)).Elem()
}

type WirelessControllerHotspot20IconInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20IconOutput() WirelessControllerHotspot20IconOutput
	ToWirelessControllerHotspot20IconOutputWithContext(ctx context.Context) WirelessControllerHotspot20IconOutput
}

func (*WirelessControllerHotspot20Icon) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20Icon)(nil)).Elem()
}

func (i *WirelessControllerHotspot20Icon) ToWirelessControllerHotspot20IconOutput() WirelessControllerHotspot20IconOutput {
	return i.ToWirelessControllerHotspot20IconOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20Icon) ToWirelessControllerHotspot20IconOutputWithContext(ctx context.Context) WirelessControllerHotspot20IconOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20IconOutput)
}

// WirelessControllerHotspot20IconArrayInput is an input type that accepts WirelessControllerHotspot20IconArray and WirelessControllerHotspot20IconArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20IconArrayInput` via:
//
//          WirelessControllerHotspot20IconArray{ WirelessControllerHotspot20IconArgs{...} }
type WirelessControllerHotspot20IconArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20IconArrayOutput() WirelessControllerHotspot20IconArrayOutput
	ToWirelessControllerHotspot20IconArrayOutputWithContext(context.Context) WirelessControllerHotspot20IconArrayOutput
}

type WirelessControllerHotspot20IconArray []WirelessControllerHotspot20IconInput

func (WirelessControllerHotspot20IconArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20Icon)(nil)).Elem()
}

func (i WirelessControllerHotspot20IconArray) ToWirelessControllerHotspot20IconArrayOutput() WirelessControllerHotspot20IconArrayOutput {
	return i.ToWirelessControllerHotspot20IconArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20IconArray) ToWirelessControllerHotspot20IconArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20IconArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20IconArrayOutput)
}

// WirelessControllerHotspot20IconMapInput is an input type that accepts WirelessControllerHotspot20IconMap and WirelessControllerHotspot20IconMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20IconMapInput` via:
//
//          WirelessControllerHotspot20IconMap{ "key": WirelessControllerHotspot20IconArgs{...} }
type WirelessControllerHotspot20IconMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20IconMapOutput() WirelessControllerHotspot20IconMapOutput
	ToWirelessControllerHotspot20IconMapOutputWithContext(context.Context) WirelessControllerHotspot20IconMapOutput
}

type WirelessControllerHotspot20IconMap map[string]WirelessControllerHotspot20IconInput

func (WirelessControllerHotspot20IconMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20Icon)(nil)).Elem()
}

func (i WirelessControllerHotspot20IconMap) ToWirelessControllerHotspot20IconMapOutput() WirelessControllerHotspot20IconMapOutput {
	return i.ToWirelessControllerHotspot20IconMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20IconMap) ToWirelessControllerHotspot20IconMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20IconMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20IconMapOutput)
}

type WirelessControllerHotspot20IconOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20IconOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20Icon)(nil)).Elem()
}

func (o WirelessControllerHotspot20IconOutput) ToWirelessControllerHotspot20IconOutput() WirelessControllerHotspot20IconOutput {
	return o
}

func (o WirelessControllerHotspot20IconOutput) ToWirelessControllerHotspot20IconOutputWithContext(ctx context.Context) WirelessControllerHotspot20IconOutput {
	return o
}

type WirelessControllerHotspot20IconArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20IconArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20Icon)(nil)).Elem()
}

func (o WirelessControllerHotspot20IconArrayOutput) ToWirelessControllerHotspot20IconArrayOutput() WirelessControllerHotspot20IconArrayOutput {
	return o
}

func (o WirelessControllerHotspot20IconArrayOutput) ToWirelessControllerHotspot20IconArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20IconArrayOutput {
	return o
}

func (o WirelessControllerHotspot20IconArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20IconOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20Icon {
		return vs[0].([]*WirelessControllerHotspot20Icon)[vs[1].(int)]
	}).(WirelessControllerHotspot20IconOutput)
}

type WirelessControllerHotspot20IconMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20IconMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20Icon)(nil)).Elem()
}

func (o WirelessControllerHotspot20IconMapOutput) ToWirelessControllerHotspot20IconMapOutput() WirelessControllerHotspot20IconMapOutput {
	return o
}

func (o WirelessControllerHotspot20IconMapOutput) ToWirelessControllerHotspot20IconMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20IconMapOutput {
	return o
}

func (o WirelessControllerHotspot20IconMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20IconOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20Icon {
		return vs[0].(map[string]*WirelessControllerHotspot20Icon)[vs[1].(string)]
	}).(WirelessControllerHotspot20IconOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20IconInput)(nil)).Elem(), &WirelessControllerHotspot20Icon{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20IconArrayInput)(nil)).Elem(), WirelessControllerHotspot20IconArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20IconMapInput)(nil)).Elem(), WirelessControllerHotspot20IconMap{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20IconOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20IconArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20IconMapOutput{})
}
