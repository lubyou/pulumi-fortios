// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WirelessControllerHotspot20AnqpVenueUrl struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                                      `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                                      `pulumi:"getAllTables"`
	Name                pulumi.StringOutput                                         `pulumi:"name"`
	ValueLists          WirelessControllerHotspot20AnqpVenueUrlValueListArrayOutput `pulumi:"valueLists"`
	Vdomparam           pulumi.StringPtrOutput                                      `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20AnqpVenueUrl registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20AnqpVenueUrl(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20AnqpVenueUrlArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20AnqpVenueUrl, error) {
	if args == nil {
		args = &WirelessControllerHotspot20AnqpVenueUrlArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerHotspot20AnqpVenueUrl
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20AnqpVenueUrl:WirelessControllerHotspot20AnqpVenueUrl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20AnqpVenueUrl gets an existing WirelessControllerHotspot20AnqpVenueUrl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20AnqpVenueUrl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20AnqpVenueUrlState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20AnqpVenueUrl, error) {
	var resource WirelessControllerHotspot20AnqpVenueUrl
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20AnqpVenueUrl:WirelessControllerHotspot20AnqpVenueUrl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20AnqpVenueUrl resources.
type wirelessControllerHotspot20AnqpVenueUrlState struct {
	DynamicSortSubtable *string                                            `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                            `pulumi:"getAllTables"`
	Name                *string                                            `pulumi:"name"`
	ValueLists          []WirelessControllerHotspot20AnqpVenueUrlValueList `pulumi:"valueLists"`
	Vdomparam           *string                                            `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20AnqpVenueUrlState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ValueLists          WirelessControllerHotspot20AnqpVenueUrlValueListArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerHotspot20AnqpVenueUrlState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20AnqpVenueUrlState)(nil)).Elem()
}

type wirelessControllerHotspot20AnqpVenueUrlArgs struct {
	DynamicSortSubtable *string                                            `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                            `pulumi:"getAllTables"`
	Name                *string                                            `pulumi:"name"`
	ValueLists          []WirelessControllerHotspot20AnqpVenueUrlValueList `pulumi:"valueLists"`
	Vdomparam           *string                                            `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20AnqpVenueUrl resource.
type WirelessControllerHotspot20AnqpVenueUrlArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ValueLists          WirelessControllerHotspot20AnqpVenueUrlValueListArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerHotspot20AnqpVenueUrlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20AnqpVenueUrlArgs)(nil)).Elem()
}

type WirelessControllerHotspot20AnqpVenueUrlInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpVenueUrlOutput() WirelessControllerHotspot20AnqpVenueUrlOutput
	ToWirelessControllerHotspot20AnqpVenueUrlOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueUrlOutput
}

func (*WirelessControllerHotspot20AnqpVenueUrl) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20AnqpVenueUrl)(nil)).Elem()
}

func (i *WirelessControllerHotspot20AnqpVenueUrl) ToWirelessControllerHotspot20AnqpVenueUrlOutput() WirelessControllerHotspot20AnqpVenueUrlOutput {
	return i.ToWirelessControllerHotspot20AnqpVenueUrlOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20AnqpVenueUrl) ToWirelessControllerHotspot20AnqpVenueUrlOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueUrlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpVenueUrlOutput)
}

// WirelessControllerHotspot20AnqpVenueUrlArrayInput is an input type that accepts WirelessControllerHotspot20AnqpVenueUrlArray and WirelessControllerHotspot20AnqpVenueUrlArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20AnqpVenueUrlArrayInput` via:
//
//	WirelessControllerHotspot20AnqpVenueUrlArray{ WirelessControllerHotspot20AnqpVenueUrlArgs{...} }
type WirelessControllerHotspot20AnqpVenueUrlArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpVenueUrlArrayOutput() WirelessControllerHotspot20AnqpVenueUrlArrayOutput
	ToWirelessControllerHotspot20AnqpVenueUrlArrayOutputWithContext(context.Context) WirelessControllerHotspot20AnqpVenueUrlArrayOutput
}

type WirelessControllerHotspot20AnqpVenueUrlArray []WirelessControllerHotspot20AnqpVenueUrlInput

func (WirelessControllerHotspot20AnqpVenueUrlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20AnqpVenueUrl)(nil)).Elem()
}

func (i WirelessControllerHotspot20AnqpVenueUrlArray) ToWirelessControllerHotspot20AnqpVenueUrlArrayOutput() WirelessControllerHotspot20AnqpVenueUrlArrayOutput {
	return i.ToWirelessControllerHotspot20AnqpVenueUrlArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20AnqpVenueUrlArray) ToWirelessControllerHotspot20AnqpVenueUrlArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueUrlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpVenueUrlArrayOutput)
}

// WirelessControllerHotspot20AnqpVenueUrlMapInput is an input type that accepts WirelessControllerHotspot20AnqpVenueUrlMap and WirelessControllerHotspot20AnqpVenueUrlMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20AnqpVenueUrlMapInput` via:
//
//	WirelessControllerHotspot20AnqpVenueUrlMap{ "key": WirelessControllerHotspot20AnqpVenueUrlArgs{...} }
type WirelessControllerHotspot20AnqpVenueUrlMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpVenueUrlMapOutput() WirelessControllerHotspot20AnqpVenueUrlMapOutput
	ToWirelessControllerHotspot20AnqpVenueUrlMapOutputWithContext(context.Context) WirelessControllerHotspot20AnqpVenueUrlMapOutput
}

type WirelessControllerHotspot20AnqpVenueUrlMap map[string]WirelessControllerHotspot20AnqpVenueUrlInput

func (WirelessControllerHotspot20AnqpVenueUrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20AnqpVenueUrl)(nil)).Elem()
}

func (i WirelessControllerHotspot20AnqpVenueUrlMap) ToWirelessControllerHotspot20AnqpVenueUrlMapOutput() WirelessControllerHotspot20AnqpVenueUrlMapOutput {
	return i.ToWirelessControllerHotspot20AnqpVenueUrlMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20AnqpVenueUrlMap) ToWirelessControllerHotspot20AnqpVenueUrlMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueUrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpVenueUrlMapOutput)
}

type WirelessControllerHotspot20AnqpVenueUrlOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpVenueUrlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20AnqpVenueUrl)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpVenueUrlOutput) ToWirelessControllerHotspot20AnqpVenueUrlOutput() WirelessControllerHotspot20AnqpVenueUrlOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueUrlOutput) ToWirelessControllerHotspot20AnqpVenueUrlOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueUrlOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueUrlOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueUrl) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerHotspot20AnqpVenueUrlOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueUrl) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerHotspot20AnqpVenueUrlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueUrl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerHotspot20AnqpVenueUrlOutput) ValueLists() WirelessControllerHotspot20AnqpVenueUrlValueListArrayOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueUrl) WirelessControllerHotspot20AnqpVenueUrlValueListArrayOutput {
		return v.ValueLists
	}).(WirelessControllerHotspot20AnqpVenueUrlValueListArrayOutput)
}

func (o WirelessControllerHotspot20AnqpVenueUrlOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueUrl) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerHotspot20AnqpVenueUrlArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpVenueUrlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20AnqpVenueUrl)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpVenueUrlArrayOutput) ToWirelessControllerHotspot20AnqpVenueUrlArrayOutput() WirelessControllerHotspot20AnqpVenueUrlArrayOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueUrlArrayOutput) ToWirelessControllerHotspot20AnqpVenueUrlArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueUrlArrayOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueUrlArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20AnqpVenueUrlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20AnqpVenueUrl {
		return vs[0].([]*WirelessControllerHotspot20AnqpVenueUrl)[vs[1].(int)]
	}).(WirelessControllerHotspot20AnqpVenueUrlOutput)
}

type WirelessControllerHotspot20AnqpVenueUrlMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpVenueUrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20AnqpVenueUrl)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpVenueUrlMapOutput) ToWirelessControllerHotspot20AnqpVenueUrlMapOutput() WirelessControllerHotspot20AnqpVenueUrlMapOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueUrlMapOutput) ToWirelessControllerHotspot20AnqpVenueUrlMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueUrlMapOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueUrlMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20AnqpVenueUrlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20AnqpVenueUrl {
		return vs[0].(map[string]*WirelessControllerHotspot20AnqpVenueUrl)[vs[1].(string)]
	}).(WirelessControllerHotspot20AnqpVenueUrlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpVenueUrlInput)(nil)).Elem(), &WirelessControllerHotspot20AnqpVenueUrl{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpVenueUrlArrayInput)(nil)).Elem(), WirelessControllerHotspot20AnqpVenueUrlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpVenueUrlMapInput)(nil)).Elem(), WirelessControllerHotspot20AnqpVenueUrlMap{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpVenueUrlOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpVenueUrlArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpVenueUrlMapOutput{})
}
