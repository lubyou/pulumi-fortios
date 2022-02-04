// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure virtual Access Point (VAP) groups.
//
// ## Import
//
// WirelessController VapGroup can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerVapGroup:WirelessControllerVapGroup labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerVapGroup struct {
	pulumi.CustomResourceState

	// Comment.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// vap name
	Name pulumi.StringOutput `pulumi:"name"`
	// List of SSIDs to be included in the VAP group. The structure of `vaps` block is documented below.
	Vaps WirelessControllerVapGroupVapArrayOutput `pulumi:"vaps"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerVapGroup registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerVapGroup(ctx *pulumi.Context,
	name string, args *WirelessControllerVapGroupArgs, opts ...pulumi.ResourceOption) (*WirelessControllerVapGroup, error) {
	if args == nil {
		args = &WirelessControllerVapGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelessControllerVapGroup
	err := ctx.RegisterResource("fortios:index/wirelessControllerVapGroup:WirelessControllerVapGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerVapGroup gets an existing WirelessControllerVapGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerVapGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerVapGroupState, opts ...pulumi.ResourceOption) (*WirelessControllerVapGroup, error) {
	var resource WirelessControllerVapGroup
	err := ctx.ReadResource("fortios:index/wirelessControllerVapGroup:WirelessControllerVapGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerVapGroup resources.
type wirelessControllerVapGroupState struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// vap name
	Name *string `pulumi:"name"`
	// List of SSIDs to be included in the VAP group. The structure of `vaps` block is documented below.
	Vaps []WirelessControllerVapGroupVap `pulumi:"vaps"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerVapGroupState struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// vap name
	Name pulumi.StringPtrInput
	// List of SSIDs to be included in the VAP group. The structure of `vaps` block is documented below.
	Vaps WirelessControllerVapGroupVapArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerVapGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerVapGroupState)(nil)).Elem()
}

type wirelessControllerVapGroupArgs struct {
	// Comment.
	Comment *string `pulumi:"comment"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// vap name
	Name *string `pulumi:"name"`
	// List of SSIDs to be included in the VAP group. The structure of `vaps` block is documented below.
	Vaps []WirelessControllerVapGroupVap `pulumi:"vaps"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerVapGroup resource.
type WirelessControllerVapGroupArgs struct {
	// Comment.
	Comment pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// vap name
	Name pulumi.StringPtrInput
	// List of SSIDs to be included in the VAP group. The structure of `vaps` block is documented below.
	Vaps WirelessControllerVapGroupVapArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerVapGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerVapGroupArgs)(nil)).Elem()
}

type WirelessControllerVapGroupInput interface {
	pulumi.Input

	ToWirelessControllerVapGroupOutput() WirelessControllerVapGroupOutput
	ToWirelessControllerVapGroupOutputWithContext(ctx context.Context) WirelessControllerVapGroupOutput
}

func (*WirelessControllerVapGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerVapGroup)(nil)).Elem()
}

func (i *WirelessControllerVapGroup) ToWirelessControllerVapGroupOutput() WirelessControllerVapGroupOutput {
	return i.ToWirelessControllerVapGroupOutputWithContext(context.Background())
}

func (i *WirelessControllerVapGroup) ToWirelessControllerVapGroupOutputWithContext(ctx context.Context) WirelessControllerVapGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerVapGroupOutput)
}

// WirelessControllerVapGroupArrayInput is an input type that accepts WirelessControllerVapGroupArray and WirelessControllerVapGroupArrayOutput values.
// You can construct a concrete instance of `WirelessControllerVapGroupArrayInput` via:
//
//          WirelessControllerVapGroupArray{ WirelessControllerVapGroupArgs{...} }
type WirelessControllerVapGroupArrayInput interface {
	pulumi.Input

	ToWirelessControllerVapGroupArrayOutput() WirelessControllerVapGroupArrayOutput
	ToWirelessControllerVapGroupArrayOutputWithContext(context.Context) WirelessControllerVapGroupArrayOutput
}

type WirelessControllerVapGroupArray []WirelessControllerVapGroupInput

func (WirelessControllerVapGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerVapGroup)(nil)).Elem()
}

func (i WirelessControllerVapGroupArray) ToWirelessControllerVapGroupArrayOutput() WirelessControllerVapGroupArrayOutput {
	return i.ToWirelessControllerVapGroupArrayOutputWithContext(context.Background())
}

func (i WirelessControllerVapGroupArray) ToWirelessControllerVapGroupArrayOutputWithContext(ctx context.Context) WirelessControllerVapGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerVapGroupArrayOutput)
}

// WirelessControllerVapGroupMapInput is an input type that accepts WirelessControllerVapGroupMap and WirelessControllerVapGroupMapOutput values.
// You can construct a concrete instance of `WirelessControllerVapGroupMapInput` via:
//
//          WirelessControllerVapGroupMap{ "key": WirelessControllerVapGroupArgs{...} }
type WirelessControllerVapGroupMapInput interface {
	pulumi.Input

	ToWirelessControllerVapGroupMapOutput() WirelessControllerVapGroupMapOutput
	ToWirelessControllerVapGroupMapOutputWithContext(context.Context) WirelessControllerVapGroupMapOutput
}

type WirelessControllerVapGroupMap map[string]WirelessControllerVapGroupInput

func (WirelessControllerVapGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerVapGroup)(nil)).Elem()
}

func (i WirelessControllerVapGroupMap) ToWirelessControllerVapGroupMapOutput() WirelessControllerVapGroupMapOutput {
	return i.ToWirelessControllerVapGroupMapOutputWithContext(context.Background())
}

func (i WirelessControllerVapGroupMap) ToWirelessControllerVapGroupMapOutputWithContext(ctx context.Context) WirelessControllerVapGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerVapGroupMapOutput)
}

type WirelessControllerVapGroupOutput struct{ *pulumi.OutputState }

func (WirelessControllerVapGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerVapGroup)(nil)).Elem()
}

func (o WirelessControllerVapGroupOutput) ToWirelessControllerVapGroupOutput() WirelessControllerVapGroupOutput {
	return o
}

func (o WirelessControllerVapGroupOutput) ToWirelessControllerVapGroupOutputWithContext(ctx context.Context) WirelessControllerVapGroupOutput {
	return o
}

type WirelessControllerVapGroupArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerVapGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerVapGroup)(nil)).Elem()
}

func (o WirelessControllerVapGroupArrayOutput) ToWirelessControllerVapGroupArrayOutput() WirelessControllerVapGroupArrayOutput {
	return o
}

func (o WirelessControllerVapGroupArrayOutput) ToWirelessControllerVapGroupArrayOutputWithContext(ctx context.Context) WirelessControllerVapGroupArrayOutput {
	return o
}

func (o WirelessControllerVapGroupArrayOutput) Index(i pulumi.IntInput) WirelessControllerVapGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerVapGroup {
		return vs[0].([]*WirelessControllerVapGroup)[vs[1].(int)]
	}).(WirelessControllerVapGroupOutput)
}

type WirelessControllerVapGroupMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerVapGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerVapGroup)(nil)).Elem()
}

func (o WirelessControllerVapGroupMapOutput) ToWirelessControllerVapGroupMapOutput() WirelessControllerVapGroupMapOutput {
	return o
}

func (o WirelessControllerVapGroupMapOutput) ToWirelessControllerVapGroupMapOutputWithContext(ctx context.Context) WirelessControllerVapGroupMapOutput {
	return o
}

func (o WirelessControllerVapGroupMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerVapGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerVapGroup {
		return vs[0].(map[string]*WirelessControllerVapGroup)[vs[1].(string)]
	}).(WirelessControllerVapGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerVapGroupInput)(nil)).Elem(), &WirelessControllerVapGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerVapGroupArrayInput)(nil)).Elem(), WirelessControllerVapGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerVapGroupMapInput)(nil)).Elem(), WirelessControllerVapGroupMap{})
	pulumi.RegisterOutputType(WirelessControllerVapGroupOutput{})
	pulumi.RegisterOutputType(WirelessControllerVapGroupArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerVapGroupMapOutput{})
}
