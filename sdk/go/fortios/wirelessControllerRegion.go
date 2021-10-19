// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiAP regions (for floor plans and maps).
//
// ## Import
//
// WirelessController Region can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/wirelessControllerRegion:WirelessControllerRegion labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WirelessControllerRegion struct {
	pulumi.CustomResourceState

	// Comments.
	Comments pulumi.StringOutput `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringOutput `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringOutput `pulumi:"imageType"`
	// FortiAP region name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity pulumi.IntOutput `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelessControllerRegion registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerRegion(ctx *pulumi.Context,
	name string, args *WirelessControllerRegionArgs, opts ...pulumi.ResourceOption) (*WirelessControllerRegion, error) {
	if args == nil {
		args = &WirelessControllerRegionArgs{}
	}

	var resource WirelessControllerRegion
	err := ctx.RegisterResource("fortios:index/wirelessControllerRegion:WirelessControllerRegion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerRegion gets an existing WirelessControllerRegion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerRegion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerRegionState, opts ...pulumi.ResourceOption) (*WirelessControllerRegion, error) {
	var resource WirelessControllerRegion
	err := ctx.ReadResource("fortios:index/wirelessControllerRegion:WirelessControllerRegion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerRegion resources.
type wirelessControllerRegionState struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale *string `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType *string `pulumi:"imageType"`
	// FortiAP region name.
	Name *string `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity *int `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelessControllerRegionState struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringPtrInput
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringPtrInput
	// FortiAP region name.
	Name pulumi.StringPtrInput
	// Region image opacity (0 - 100).
	Opacity pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerRegionState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerRegionState)(nil)).Elem()
}

type wirelessControllerRegionArgs struct {
	// Comments.
	Comments *string `pulumi:"comments"`
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale *string `pulumi:"grayscale"`
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType *string `pulumi:"imageType"`
	// FortiAP region name.
	Name *string `pulumi:"name"`
	// Region image opacity (0 - 100).
	Opacity *int `pulumi:"opacity"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerRegion resource.
type WirelessControllerRegionArgs struct {
	// Comments.
	Comments pulumi.StringPtrInput
	// Region image grayscale. Valid values: `enable`, `disable`.
	Grayscale pulumi.StringPtrInput
	// FortiAP region image type (png|jpeg|gif). Valid values: `png`, `jpeg`, `gif`.
	ImageType pulumi.StringPtrInput
	// FortiAP region name.
	Name pulumi.StringPtrInput
	// Region image opacity (0 - 100).
	Opacity pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelessControllerRegionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerRegionArgs)(nil)).Elem()
}

type WirelessControllerRegionInput interface {
	pulumi.Input

	ToWirelessControllerRegionOutput() WirelessControllerRegionOutput
	ToWirelessControllerRegionOutputWithContext(ctx context.Context) WirelessControllerRegionOutput
}

func (*WirelessControllerRegion) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerRegion)(nil))
}

func (i *WirelessControllerRegion) ToWirelessControllerRegionOutput() WirelessControllerRegionOutput {
	return i.ToWirelessControllerRegionOutputWithContext(context.Background())
}

func (i *WirelessControllerRegion) ToWirelessControllerRegionOutputWithContext(ctx context.Context) WirelessControllerRegionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerRegionOutput)
}

func (i *WirelessControllerRegion) ToWirelessControllerRegionPtrOutput() WirelessControllerRegionPtrOutput {
	return i.ToWirelessControllerRegionPtrOutputWithContext(context.Background())
}

func (i *WirelessControllerRegion) ToWirelessControllerRegionPtrOutputWithContext(ctx context.Context) WirelessControllerRegionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerRegionPtrOutput)
}

type WirelessControllerRegionPtrInput interface {
	pulumi.Input

	ToWirelessControllerRegionPtrOutput() WirelessControllerRegionPtrOutput
	ToWirelessControllerRegionPtrOutputWithContext(ctx context.Context) WirelessControllerRegionPtrOutput
}

type wirelessControllerRegionPtrType WirelessControllerRegionArgs

func (*wirelessControllerRegionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerRegion)(nil))
}

func (i *wirelessControllerRegionPtrType) ToWirelessControllerRegionPtrOutput() WirelessControllerRegionPtrOutput {
	return i.ToWirelessControllerRegionPtrOutputWithContext(context.Background())
}

func (i *wirelessControllerRegionPtrType) ToWirelessControllerRegionPtrOutputWithContext(ctx context.Context) WirelessControllerRegionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerRegionPtrOutput)
}

// WirelessControllerRegionArrayInput is an input type that accepts WirelessControllerRegionArray and WirelessControllerRegionArrayOutput values.
// You can construct a concrete instance of `WirelessControllerRegionArrayInput` via:
//
//          WirelessControllerRegionArray{ WirelessControllerRegionArgs{...} }
type WirelessControllerRegionArrayInput interface {
	pulumi.Input

	ToWirelessControllerRegionArrayOutput() WirelessControllerRegionArrayOutput
	ToWirelessControllerRegionArrayOutputWithContext(context.Context) WirelessControllerRegionArrayOutput
}

type WirelessControllerRegionArray []WirelessControllerRegionInput

func (WirelessControllerRegionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WirelessControllerRegion)(nil))
}

func (i WirelessControllerRegionArray) ToWirelessControllerRegionArrayOutput() WirelessControllerRegionArrayOutput {
	return i.ToWirelessControllerRegionArrayOutputWithContext(context.Background())
}

func (i WirelessControllerRegionArray) ToWirelessControllerRegionArrayOutputWithContext(ctx context.Context) WirelessControllerRegionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerRegionArrayOutput)
}

// WirelessControllerRegionMapInput is an input type that accepts WirelessControllerRegionMap and WirelessControllerRegionMapOutput values.
// You can construct a concrete instance of `WirelessControllerRegionMapInput` via:
//
//          WirelessControllerRegionMap{ "key": WirelessControllerRegionArgs{...} }
type WirelessControllerRegionMapInput interface {
	pulumi.Input

	ToWirelessControllerRegionMapOutput() WirelessControllerRegionMapOutput
	ToWirelessControllerRegionMapOutputWithContext(context.Context) WirelessControllerRegionMapOutput
}

type WirelessControllerRegionMap map[string]WirelessControllerRegionInput

func (WirelessControllerRegionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WirelessControllerRegion)(nil))
}

func (i WirelessControllerRegionMap) ToWirelessControllerRegionMapOutput() WirelessControllerRegionMapOutput {
	return i.ToWirelessControllerRegionMapOutputWithContext(context.Background())
}

func (i WirelessControllerRegionMap) ToWirelessControllerRegionMapOutputWithContext(ctx context.Context) WirelessControllerRegionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerRegionMapOutput)
}

type WirelessControllerRegionOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerRegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WirelessControllerRegion)(nil))
}

func (o WirelessControllerRegionOutput) ToWirelessControllerRegionOutput() WirelessControllerRegionOutput {
	return o
}

func (o WirelessControllerRegionOutput) ToWirelessControllerRegionOutputWithContext(ctx context.Context) WirelessControllerRegionOutput {
	return o
}

func (o WirelessControllerRegionOutput) ToWirelessControllerRegionPtrOutput() WirelessControllerRegionPtrOutput {
	return o.ToWirelessControllerRegionPtrOutputWithContext(context.Background())
}

func (o WirelessControllerRegionOutput) ToWirelessControllerRegionPtrOutputWithContext(ctx context.Context) WirelessControllerRegionPtrOutput {
	return o.ApplyT(func(v WirelessControllerRegion) *WirelessControllerRegion {
		return &v
	}).(WirelessControllerRegionPtrOutput)
}

type WirelessControllerRegionPtrOutput struct {
	*pulumi.OutputState
}

func (WirelessControllerRegionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerRegion)(nil))
}

func (o WirelessControllerRegionPtrOutput) ToWirelessControllerRegionPtrOutput() WirelessControllerRegionPtrOutput {
	return o
}

func (o WirelessControllerRegionPtrOutput) ToWirelessControllerRegionPtrOutputWithContext(ctx context.Context) WirelessControllerRegionPtrOutput {
	return o
}

type WirelessControllerRegionArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerRegionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WirelessControllerRegion)(nil))
}

func (o WirelessControllerRegionArrayOutput) ToWirelessControllerRegionArrayOutput() WirelessControllerRegionArrayOutput {
	return o
}

func (o WirelessControllerRegionArrayOutput) ToWirelessControllerRegionArrayOutputWithContext(ctx context.Context) WirelessControllerRegionArrayOutput {
	return o
}

func (o WirelessControllerRegionArrayOutput) Index(i pulumi.IntInput) WirelessControllerRegionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WirelessControllerRegion {
		return vs[0].([]WirelessControllerRegion)[vs[1].(int)]
	}).(WirelessControllerRegionOutput)
}

type WirelessControllerRegionMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerRegionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WirelessControllerRegion)(nil))
}

func (o WirelessControllerRegionMapOutput) ToWirelessControllerRegionMapOutput() WirelessControllerRegionMapOutput {
	return o
}

func (o WirelessControllerRegionMapOutput) ToWirelessControllerRegionMapOutputWithContext(ctx context.Context) WirelessControllerRegionMapOutput {
	return o
}

func (o WirelessControllerRegionMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerRegionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WirelessControllerRegion {
		return vs[0].(map[string]WirelessControllerRegion)[vs[1].(string)]
	}).(WirelessControllerRegionOutput)
}

func init() {
	pulumi.RegisterOutputType(WirelessControllerRegionOutput{})
	pulumi.RegisterOutputType(WirelessControllerRegionPtrOutput{})
	pulumi.RegisterOutputType(WirelessControllerRegionArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerRegionMapOutput{})
}
