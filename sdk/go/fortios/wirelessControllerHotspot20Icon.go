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

type WirelessControllerHotspot20Icon struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                             `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                             `pulumi:"getAllTables"`
	IconLists           WirelessControllerHotspot20IconIconListArrayOutput `pulumi:"iconLists"`
	Name                pulumi.StringOutput                                `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput                             `pulumi:"vdomparam"`
}

// NewWirelessControllerHotspot20Icon registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerHotspot20Icon(ctx *pulumi.Context,
	name string, args *WirelessControllerHotspot20IconArgs, opts ...pulumi.ResourceOption) (*WirelessControllerHotspot20Icon, error) {
	if args == nil {
		args = &WirelessControllerHotspot20IconArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	DynamicSortSubtable *string                                   `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                   `pulumi:"getAllTables"`
	IconLists           []WirelessControllerHotspot20IconIconList `pulumi:"iconLists"`
	Name                *string                                   `pulumi:"name"`
	Vdomparam           *string                                   `pulumi:"vdomparam"`
}

type WirelessControllerHotspot20IconState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	IconLists           WirelessControllerHotspot20IconIconListArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (WirelessControllerHotspot20IconState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerHotspot20IconState)(nil)).Elem()
}

type wirelessControllerHotspot20IconArgs struct {
	DynamicSortSubtable *string                                   `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                                   `pulumi:"getAllTables"`
	IconLists           []WirelessControllerHotspot20IconIconList `pulumi:"iconLists"`
	Name                *string                                   `pulumi:"name"`
	Vdomparam           *string                                   `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelessControllerHotspot20Icon resource.
type WirelessControllerHotspot20IconArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	IconLists           WirelessControllerHotspot20IconIconListArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
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

func (i *WirelessControllerHotspot20Icon) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerHotspot20Icon] {
	return pulumix.Output[*WirelessControllerHotspot20Icon]{
		OutputState: i.ToWirelessControllerHotspot20IconOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerHotspot20IconArrayInput is an input type that accepts WirelessControllerHotspot20IconArray and WirelessControllerHotspot20IconArrayOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20IconArrayInput` via:
//
//	WirelessControllerHotspot20IconArray{ WirelessControllerHotspot20IconArgs{...} }
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

func (i WirelessControllerHotspot20IconArray) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerHotspot20Icon] {
	return pulumix.Output[[]*WirelessControllerHotspot20Icon]{
		OutputState: i.ToWirelessControllerHotspot20IconArrayOutputWithContext(ctx).OutputState,
	}
}

// WirelessControllerHotspot20IconMapInput is an input type that accepts WirelessControllerHotspot20IconMap and WirelessControllerHotspot20IconMapOutput values.
// You can construct a concrete instance of `WirelessControllerHotspot20IconMapInput` via:
//
//	WirelessControllerHotspot20IconMap{ "key": WirelessControllerHotspot20IconArgs{...} }
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

func (i WirelessControllerHotspot20IconMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerHotspot20Icon] {
	return pulumix.Output[map[string]*WirelessControllerHotspot20Icon]{
		OutputState: i.ToWirelessControllerHotspot20IconMapOutputWithContext(ctx).OutputState,
	}
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

func (o WirelessControllerHotspot20IconOutput) ToOutput(ctx context.Context) pulumix.Output[*WirelessControllerHotspot20Icon] {
	return pulumix.Output[*WirelessControllerHotspot20Icon]{
		OutputState: o.OutputState,
	}
}

func (o WirelessControllerHotspot20IconOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20Icon) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerHotspot20IconOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20Icon) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerHotspot20IconOutput) IconLists() WirelessControllerHotspot20IconIconListArrayOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20Icon) WirelessControllerHotspot20IconIconListArrayOutput {
		return v.IconLists
	}).(WirelessControllerHotspot20IconIconListArrayOutput)
}

func (o WirelessControllerHotspot20IconOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20Icon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerHotspot20IconOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerHotspot20Icon) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
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

func (o WirelessControllerHotspot20IconArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WirelessControllerHotspot20Icon] {
	return pulumix.Output[[]*WirelessControllerHotspot20Icon]{
		OutputState: o.OutputState,
	}
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

func (o WirelessControllerHotspot20IconMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WirelessControllerHotspot20Icon] {
	return pulumix.Output[map[string]*WirelessControllerHotspot20Icon]{
		OutputState: o.OutputState,
	}
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
