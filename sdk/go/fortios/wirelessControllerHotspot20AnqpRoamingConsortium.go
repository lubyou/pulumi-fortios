// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure roaming consortium.
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
// 		_, err := fortios.NewWirelessControllerHotspot20AnqpRoamingConsortium(ctx, "trname", nil)
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
// WirelessControllerHotspot20 AnqpRoamingConsortium can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerHotspot20AnqpRoamingConsortium struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Roaming consortium name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Organization identifier list. The structure of `oiList` block is documented below.
	OiLists WirelessControllerHotspot20AnqpRoamingConsortiumOiListArrayOutput `pulumi:"oiLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20AnqpRoamingConsortium registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20AnqpRoamingConsortium(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20AnqpRoamingConsortiumArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20AnqpRoamingConsortium, error) {
	if args == nil {
		args = &WirelessControllerHotspot20AnqpRoamingConsortiumArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerHotspot20AnqpRoamingConsortium
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20AnqpRoamingConsortium gets an existing WirelessControllerHotspot20AnqpRoamingConsortium resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20AnqpRoamingConsortium(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20AnqpRoamingConsortiumState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20AnqpRoamingConsortium, error) {
	var resource WirelessControllerHotspot20AnqpRoamingConsortium
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20AnqpRoamingConsortium resources.
type wirelessControllerHotspot20AnqpRoamingConsortiumState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Roaming consortium name.
	Name *string `pulumi:"name"`
	// Organization identifier list. The structure of `oiList` block is documented below.
	OiLists []WirelessControllerHotspot20AnqpRoamingConsortiumOiList `pulumi:"oiLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20AnqpRoamingConsortiumState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Roaming consortium name.
	Name pulumi.StringPtrInput
	// Organization identifier list. The structure of `oiList` block is documented below.
	OiLists WirelessControllerHotspot20AnqpRoamingConsortiumOiListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20AnqpRoamingConsortiumState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20AnqpRoamingConsortiumState)(nil)).Elem()
}

type wirelessControllerHotspot20AnqpRoamingConsortiumArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Roaming consortium name.
	Name *string `pulumi:"name"`
	// Organization identifier list. The structure of `oiList` block is documented below.
	OiLists []WirelessControllerHotspot20AnqpRoamingConsortiumOiList `pulumi:"oiLists"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20AnqpRoamingConsortium resource.
type WirelessControllerHotspot20AnqpRoamingConsortiumArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Roaming consortium name.
	Name pulumi.StringPtrInput
	// Organization identifier list. The structure of `oiList` block is documented below.
	OiLists WirelessControllerHotspot20AnqpRoamingConsortiumOiListArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerHotspot20AnqpRoamingConsortiumArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20AnqpRoamingConsortiumArgs)(nil)).Elem()
}

type WirelessControllerHotspot20AnqpRoamingConsortiumInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpRoamingConsortiumOutput() WirelessControllerHotspot20AnqpRoamingConsortiumOutput
	ToWirelessControllerHotspot20AnqpRoamingConsortiumOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumOutput
}

func (*WirelessControllerHotspot20AnqpRoamingConsortium) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20AnqpRoamingConsortium)(nil)).Elem()
}

func (i *WirelessControllerHotspot20AnqpRoamingConsortium) ToWirelessControllerHotspot20AnqpRoamingConsortiumOutput() WirelessControllerHotspot20AnqpRoamingConsortiumOutput {
	return i.ToWirelessControllerHotspot20AnqpRoamingConsortiumOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20AnqpRoamingConsortium) ToWirelessControllerHotspot20AnqpRoamingConsortiumOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpRoamingConsortiumOutput)
}

// WirelessControllerHotspot20AnqpRoamingConsortiumArrayInput is an input type that accepts WirelessControllerHotspot20AnqpRoamingConsortiumArray and WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20AnqpRoamingConsortiumArrayInput` via:
//
//          WirelessControllerHotspot20AnqpRoamingConsortiumArray{ WirelessControllerHotspot20AnqpRoamingConsortiumArgs{...} }
type WirelessControllerHotspot20AnqpRoamingConsortiumArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput() WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput
	ToWirelessControllerHotspot20AnqpRoamingConsortiumArrayOutputWithContext(context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput
}

type WirelessControllerHotspot20AnqpRoamingConsortiumArray []WirelessControllerHotspot20AnqpRoamingConsortiumInput

func (WirelessControllerHotspot20AnqpRoamingConsortiumArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20AnqpRoamingConsortium)(nil)).Elem()
}

func (i WirelessControllerHotspot20AnqpRoamingConsortiumArray) ToWirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput() WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput {
	return i.ToWirelessControllerHotspot20AnqpRoamingConsortiumArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20AnqpRoamingConsortiumArray) ToWirelessControllerHotspot20AnqpRoamingConsortiumArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput)
}

// WirelessControllerHotspot20AnqpRoamingConsortiumMapInput is an input type that accepts WirelessControllerHotspot20AnqpRoamingConsortiumMap and WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20AnqpRoamingConsortiumMapInput` via:
//
//          WirelessControllerHotspot20AnqpRoamingConsortiumMap{ "key": WirelessControllerHotspot20AnqpRoamingConsortiumArgs{...} }
type WirelessControllerHotspot20AnqpRoamingConsortiumMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpRoamingConsortiumMapOutput() WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput
	ToWirelessControllerHotspot20AnqpRoamingConsortiumMapOutputWithContext(context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput
}

type WirelessControllerHotspot20AnqpRoamingConsortiumMap map[string]WirelessControllerHotspot20AnqpRoamingConsortiumInput

func (WirelessControllerHotspot20AnqpRoamingConsortiumMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20AnqpRoamingConsortium)(nil)).Elem()
}

func (i WirelessControllerHotspot20AnqpRoamingConsortiumMap) ToWirelessControllerHotspot20AnqpRoamingConsortiumMapOutput() WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput {
	return i.ToWirelessControllerHotspot20AnqpRoamingConsortiumMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20AnqpRoamingConsortiumMap) ToWirelessControllerHotspot20AnqpRoamingConsortiumMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput)
}

type WirelessControllerHotspot20AnqpRoamingConsortiumOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpRoamingConsortiumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20AnqpRoamingConsortium)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpRoamingConsortiumOutput) ToWirelessControllerHotspot20AnqpRoamingConsortiumOutput() WirelessControllerHotspot20AnqpRoamingConsortiumOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpRoamingConsortiumOutput) ToWirelessControllerHotspot20AnqpRoamingConsortiumOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumOutput {
	return o
}

type WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20AnqpRoamingConsortium)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput) ToWirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput() WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput) ToWirelessControllerHotspot20AnqpRoamingConsortiumArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20AnqpRoamingConsortiumOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20AnqpRoamingConsortium {
		return vs[0].([]*WirelessControllerHotspot20AnqpRoamingConsortium)[vs[1].(int)]
	}).(WirelessControllerHotspot20AnqpRoamingConsortiumOutput)
}

type WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20AnqpRoamingConsortium)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput) ToWirelessControllerHotspot20AnqpRoamingConsortiumMapOutput() WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput) ToWirelessControllerHotspot20AnqpRoamingConsortiumMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20AnqpRoamingConsortiumOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20AnqpRoamingConsortium {
		return vs[0].(map[string]*WirelessControllerHotspot20AnqpRoamingConsortium)[vs[1].(string)]
	}).(WirelessControllerHotspot20AnqpRoamingConsortiumOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpRoamingConsortiumInput)(nil)).Elem(), &WirelessControllerHotspot20AnqpRoamingConsortium{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpRoamingConsortiumArrayInput)(nil)).Elem(), WirelessControllerHotspot20AnqpRoamingConsortiumArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpRoamingConsortiumMapInput)(nil)).Elem(), WirelessControllerHotspot20AnqpRoamingConsortiumMap{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpRoamingConsortiumOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpRoamingConsortiumArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpRoamingConsortiumMapOutput{})
}
