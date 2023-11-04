// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type WirelessControllerHotspot20AnqpVenueName struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                                       `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                                       `pulumi:"getAllTables"`
	Name                pulumi.StringOutput                                          `pulumi:"name"`
	ValueLists          WirelessControllerHotspot20AnqpVenueNameValueListArrayOutput `pulumi:"valueLists"`
	Vdomparam           pulumi.StringPtrOutput                                       `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20AnqpVenueName registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20AnqpVenueName(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20AnqpVenueNameArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20AnqpVenueName, error) {
	if args == nil {
		args = &WirelessControllerHotspot20AnqpVenueNameArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerHotspot20AnqpVenueName
	err := ctx.RegisterResource("fortios:index/wirelessControllerHotspot20AnqpVenueName:WirelessControllerHotspot20AnqpVenueName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerHotspot20AnqpVenueName gets an existing WirelessControllerHotspot20AnqpVenueName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerHotspot20AnqpVenueName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerHotspot20AnqpVenueNameState, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20AnqpVenueName, error) {
	var resource WirelessControllerHotspot20AnqpVenueName
	err := ctx.ReadResource("fortios:index/wirelessControllerHotspot20AnqpVenueName:WirelessControllerHotspot20AnqpVenueName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerHotspot20AnqpVenueName resources.
type wirelessControllerHotspot20AnqpVenueNameState struct {
	DynamicSortSubtable *string                                             `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                             `pulumi:"getAllTables"`
	Name                *string                                             `pulumi:"name"`
	ValueLists          []WirelessControllerHotspot20AnqpVenueNameValueList `pulumi:"valueLists"`
	Vdomparam           *string                                             `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20AnqpVenueNameState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ValueLists          WirelessControllerHotspot20AnqpVenueNameValueListArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerHotspot20AnqpVenueNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20AnqpVenueNameState)(nil)).Elem()
}

type wirelessControllerHotspot20AnqpVenueNameArgs struct {
	DynamicSortSubtable *string                                             `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                             `pulumi:"getAllTables"`
	Name                *string                                             `pulumi:"name"`
	ValueLists          []WirelessControllerHotspot20AnqpVenueNameValueList `pulumi:"valueLists"`
	Vdomparam           *string                                             `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20AnqpVenueName resource.
type WirelessControllerHotspot20AnqpVenueNameArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	ValueLists          WirelessControllerHotspot20AnqpVenueNameValueListArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerHotspot20AnqpVenueNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20AnqpVenueNameArgs)(nil)).Elem()
}

type WirelessControllerHotspot20AnqpVenueNameInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpVenueNameOutput() WirelessControllerHotspot20AnqpVenueNameOutput
	ToWirelessControllerHotspot20AnqpVenueNameOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueNameOutput
}

func (*WirelessControllerHotspot20AnqpVenueName) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20AnqpVenueName)(nil)).Elem()
}

func (i *WirelessControllerHotspot20AnqpVenueName) ToWirelessControllerHotspot20AnqpVenueNameOutput() WirelessControllerHotspot20AnqpVenueNameOutput {
	return i.ToWirelessControllerHotspot20AnqpVenueNameOutputWithContext(context.Background())
}

func (i *WirelessControllerHotspot20AnqpVenueName) ToWirelessControllerHotspot20AnqpVenueNameOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpVenueNameOutput)
}

func (i *WirelessControllerHotspot20AnqpVenueName) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerHotspot20AnqpVenueName] {
	return pulumix.Output[*WirelessControllerHotspot20AnqpVenueName]{
		OutputState: i.ToWirelessControllerHotspot20AnqpVenueNameOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerHotspot20AnqpVenueNameArrayInput is an input type that accepts WirelessControllerHotspot20AnqpVenueNameArray and WirelessControllerHotspot20AnqpVenueNameArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20AnqpVenueNameArrayInput` via:
//
//	WirelessControllerHotspot20AnqpVenueNameArray{ WirelessControllerHotspot20AnqpVenueNameArgs{...} }
type WirelessControllerHotspot20AnqpVenueNameArrayInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpVenueNameArrayOutput() WirelessControllerHotspot20AnqpVenueNameArrayOutput
	ToWirelessControllerHotspot20AnqpVenueNameArrayOutputWithContext(context.Context) WirelessControllerHotspot20AnqpVenueNameArrayOutput
}

type WirelessControllerHotspot20AnqpVenueNameArray []WirelessControllerHotspot20AnqpVenueNameInput

func (WirelessControllerHotspot20AnqpVenueNameArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20AnqpVenueName)(nil)).Elem()
}

func (i WirelessControllerHotspot20AnqpVenueNameArray) ToWirelessControllerHotspot20AnqpVenueNameArrayOutput() WirelessControllerHotspot20AnqpVenueNameArrayOutput {
	return i.ToWirelessControllerHotspot20AnqpVenueNameArrayOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20AnqpVenueNameArray) ToWirelessControllerHotspot20AnqpVenueNameArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueNameArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpVenueNameArrayOutput)
}

func (i WirelessControllerHotspot20AnqpVenueNameArray) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerHotspot20AnqpVenueName] {
	return pulumix.Output[[]*WirelessControllerHotspot20AnqpVenueName]{
		OutputState: i.ToWirelessControllerHotspot20AnqpVenueNameArrayOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerHotspot20AnqpVenueNameMapInput is an input type that accepts WirelessControllerHotspot20AnqpVenueNameMap and WirelessControllerHotspot20AnqpVenueNameMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20AnqpVenueNameMapInput` via:
//
//	WirelessControllerHotspot20AnqpVenueNameMap{ "key": WirelessControllerHotspot20AnqpVenueNameArgs{...} }
type WirelessControllerHotspot20AnqpVenueNameMapInput interface {
	pulumi.Input

	ToWirelessControllerHotspot20AnqpVenueNameMapOutput() WirelessControllerHotspot20AnqpVenueNameMapOutput
	ToWirelessControllerHotspot20AnqpVenueNameMapOutputWithContext(context.Context) WirelessControllerHotspot20AnqpVenueNameMapOutput
}

type WirelessControllerHotspot20AnqpVenueNameMap map[string]WirelessControllerHotspot20AnqpVenueNameInput

func (WirelessControllerHotspot20AnqpVenueNameMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20AnqpVenueName)(nil)).Elem()
}

func (i WirelessControllerHotspot20AnqpVenueNameMap) ToWirelessControllerHotspot20AnqpVenueNameMapOutput() WirelessControllerHotspot20AnqpVenueNameMapOutput {
	return i.ToWirelessControllerHotspot20AnqpVenueNameMapOutputWithContext(context.Background())
}

func (i WirelessControllerHotspot20AnqpVenueNameMap) ToWirelessControllerHotspot20AnqpVenueNameMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueNameMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerHotspot20AnqpVenueNameMapOutput)
}

func (i WirelessControllerHotspot20AnqpVenueNameMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerHotspot20AnqpVenueName] {
	return pulumix.Output[map[string]*WirelessControllerHotspot20AnqpVenueName]{
		OutputState: i.ToWirelessControllerHotspot20AnqpVenueNameMapOutputWithContext(ctx).OutputState,
	}
}

type WirelessControllerHotspot20AnqpVenueNameOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpVenueNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerHotspot20AnqpVenueName)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpVenueNameOutput) ToWirelessControllerHotspot20AnqpVenueNameOutput() WirelessControllerHotspot20AnqpVenueNameOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueNameOutput) ToWirelessControllerHotspot20AnqpVenueNameOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueNameOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueNameOutput) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerHotspot20AnqpVenueName] {
	return pulumix.Output[*WirelessControllerHotspot20AnqpVenueName]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerHotspot20AnqpVenueNameOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueName) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerHotspot20AnqpVenueNameOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueName) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerHotspot20AnqpVenueNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerHotspot20AnqpVenueNameOutput) ValueLists() WirelessControllerHotspot20AnqpVenueNameValueListArrayOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueName) WirelessControllerHotspot20AnqpVenueNameValueListArrayOutput {
		return v.ValueLists
	}).(WirelessControllerHotspot20AnqpVenueNameValueListArrayOutput)
}

func (o WirelessControllerHotspot20AnqpVenueNameOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20AnqpVenueName) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelessControllerHotspot20AnqpVenueNameArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpVenueNameArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerHotspot20AnqpVenueName)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpVenueNameArrayOutput) ToWirelessControllerHotspot20AnqpVenueNameArrayOutput() WirelessControllerHotspot20AnqpVenueNameArrayOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueNameArrayOutput) ToWirelessControllerHotspot20AnqpVenueNameArrayOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueNameArrayOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueNameArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerHotspot20AnqpVenueName] {
	return pulumix.Output[[]*WirelessControllerHotspot20AnqpVenueName]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerHotspot20AnqpVenueNameArrayOutput) Index(i pulumi.IntInput) WirelessControllerHotspot20AnqpVenueNameOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20AnqpVenueName {
		return vs[0].([]*WirelessControllerHotspot20AnqpVenueName)[vs[1].(int)]
	}).(WirelessControllerHotspot20AnqpVenueNameOutput)
}

type WirelessControllerHotspot20AnqpVenueNameMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerHotspot20AnqpVenueNameMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerHotspot20AnqpVenueName)(nil)).Elem()
}

func (o WirelessControllerHotspot20AnqpVenueNameMapOutput) ToWirelessControllerHotspot20AnqpVenueNameMapOutput() WirelessControllerHotspot20AnqpVenueNameMapOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueNameMapOutput) ToWirelessControllerHotspot20AnqpVenueNameMapOutputWithContext(ctx context.Context) WirelessControllerHotspot20AnqpVenueNameMapOutput {
	return o
}

func (o WirelessControllerHotspot20AnqpVenueNameMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerHotspot20AnqpVenueName] {
	return pulumix.Output[map[string]*WirelessControllerHotspot20AnqpVenueName]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerHotspot20AnqpVenueNameMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerHotspot20AnqpVenueNameOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerHotspot20AnqpVenueName {
		return vs[0].(map[string]*WirelessControllerHotspot20AnqpVenueName)[vs[1].(string)]
	}).(WirelessControllerHotspot20AnqpVenueNameOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpVenueNameInput)(nil)).Elem(), &WirelessControllerHotspot20AnqpVenueName{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpVenueNameArrayInput)(nil)).Elem(), WirelessControllerHotspot20AnqpVenueNameArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerHotspot20AnqpVenueNameMapInput)(nil)).Elem(), WirelessControllerHotspot20AnqpVenueNameMap{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpVenueNameOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpVenueNameArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerHotspot20AnqpVenueNameMapOutput{})
}
