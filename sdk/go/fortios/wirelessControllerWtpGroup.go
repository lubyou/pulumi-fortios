// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WirelessControllerWtpGroup struct {
	pulumi.CustomResourceState

	DynamicSortSubtable pulumi.StringPtrOutput                   `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                   `pulumi:"getAllTables"`
	Name                pulumi.StringOutput                      `pulumi:"name"`
	PlatformType        pulumi.StringOutput                      `pulumi:"platformType"`
	Vdomparam           pulumi.StringPtrOutput                   `pulumi:"vdomparam"`
	Wtps                WirelessControllerWtpGroupWtpArrayOutput `pulumi:"wtps"`
}

// NewWirelessControllerWtpGroup registers a new resource with the given unique name, arguments, and options.
func NewWirelessControllerWtpGroup(ctx *pulumi.Context,
	name string, args *WirelessControllerWtpGroupArgs, opts ...pulumi.ResourceOption) (*WirelessControllerWtpGroup, error) {
	if args == nil {
		args = &WirelessControllerWtpGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WirelessControllerWtpGroup
	err := ctx.RegisterResource("fortios:index/wirelessControllerWtpGroup:WirelessControllerWtpGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelessControllerWtpGroup gets an existing WirelessControllerWtpGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelessControllerWtpGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelessControllerWtpGroupState, opts ...pulumi.ResourceOption) (*WirelessControllerWtpGroup, error) {
	var resource WirelessControllerWtpGroup
	err := ctx.ReadResource("fortios:index/wirelessControllerWtpGroup:WirelessControllerWtpGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelessControllerWtpGroup resources.
type wirelessControllerWtpGroupState struct {
	DynamicSortSubtable *string                         `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                         `pulumi:"getAllTables"`
	Name                *string                         `pulumi:"name"`
	PlatformType        *string                         `pulumi:"platformType"`
	Vdomparam           *string                         `pulumi:"vdomparam"`
	Wtps                []WirelessControllerWtpGroupWtp `pulumi:"wtps"`
}

type WirelessControllerWtpGroupState struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	PlatformType        pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Wtps                WirelessControllerWtpGroupWtpArrayInput
}

func (WirelessControllerWtpGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerWtpGroupState)(nil)).Elem()
}

type wirelessControllerWtpGroupArgs struct {
	DynamicSortSubtable *string                         `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                         `pulumi:"getAllTables"`
	Name                *string                         `pulumi:"name"`
	PlatformType        *string                         `pulumi:"platformType"`
	Vdomparam           *string                         `pulumi:"vdomparam"`
	Wtps                []WirelessControllerWtpGroupWtp `pulumi:"wtps"`
}

// The set of arguments for constructing a WirelessControllerWtpGroup resource.
type WirelessControllerWtpGroupArgs struct {
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	PlatformType        pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Wtps                WirelessControllerWtpGroupWtpArrayInput
}

func (WirelessControllerWtpGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelessControllerWtpGroupArgs)(nil)).Elem()
}

type WirelessControllerWtpGroupInput interface {
	pulumi.Input

	ToWirelessControllerWtpGroupOutput() WirelessControllerWtpGroupOutput
	ToWirelessControllerWtpGroupOutputWithContext(ctx context.Context) WirelessControllerWtpGroupOutput
}

func (*WirelessControllerWtpGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerWtpGroup)(nil)).Elem()
}

func (i *WirelessControllerWtpGroup) ToWirelessControllerWtpGroupOutput() WirelessControllerWtpGroupOutput {
	return i.ToWirelessControllerWtpGroupOutputWithContext(context.Background())
}

func (i *WirelessControllerWtpGroup) ToWirelessControllerWtpGroupOutputWithContext(ctx context.Context) WirelessControllerWtpGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerWtpGroupOutput)
}

// WirelessControllerWtpGroupArrayInput is an input type that accepts WirelessControllerWtpGroupArray and WirelessControllerWtpGroupArrayOutput values.
// You can construct a concrete instance of `WirelessControllerWtpGroupArrayInput` via:
//
//	WirelessControllerWtpGroupArray{ WirelessControllerWtpGroupArgs{...} }
type WirelessControllerWtpGroupArrayInput interface {
	pulumi.Input

	ToWirelessControllerWtpGroupArrayOutput() WirelessControllerWtpGroupArrayOutput
	ToWirelessControllerWtpGroupArrayOutputWithContext(context.Context) WirelessControllerWtpGroupArrayOutput
}

type WirelessControllerWtpGroupArray []WirelessControllerWtpGroupInput

func (WirelessControllerWtpGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerWtpGroup)(nil)).Elem()
}

func (i WirelessControllerWtpGroupArray) ToWirelessControllerWtpGroupArrayOutput() WirelessControllerWtpGroupArrayOutput {
	return i.ToWirelessControllerWtpGroupArrayOutputWithContext(context.Background())
}

func (i WirelessControllerWtpGroupArray) ToWirelessControllerWtpGroupArrayOutputWithContext(ctx context.Context) WirelessControllerWtpGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerWtpGroupArrayOutput)
}

// WirelessControllerWtpGroupMapInput is an input type that accepts WirelessControllerWtpGroupMap and WirelessControllerWtpGroupMapOutput values.
// You can construct a concrete instance of `WirelessControllerWtpGroupMapInput` via:
//
//	WirelessControllerWtpGroupMap{ "key": WirelessControllerWtpGroupArgs{...} }
type WirelessControllerWtpGroupMapInput interface {
	pulumi.Input

	ToWirelessControllerWtpGroupMapOutput() WirelessControllerWtpGroupMapOutput
	ToWirelessControllerWtpGroupMapOutputWithContext(context.Context) WirelessControllerWtpGroupMapOutput
}

type WirelessControllerWtpGroupMap map[string]WirelessControllerWtpGroupInput

func (WirelessControllerWtpGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerWtpGroup)(nil)).Elem()
}

func (i WirelessControllerWtpGroupMap) ToWirelessControllerWtpGroupMapOutput() WirelessControllerWtpGroupMapOutput {
	return i.ToWirelessControllerWtpGroupMapOutputWithContext(context.Background())
}

func (i WirelessControllerWtpGroupMap) ToWirelessControllerWtpGroupMapOutputWithContext(ctx context.Context) WirelessControllerWtpGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelessControllerWtpGroupMapOutput)
}

type WirelessControllerWtpGroupOutput struct{ *pulumi.OutputState }

func (WirelessControllerWtpGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelessControllerWtpGroup)(nil)).Elem()
}

func (o WirelessControllerWtpGroupOutput) ToWirelessControllerWtpGroupOutput() WirelessControllerWtpGroupOutput {
	return o
}

func (o WirelessControllerWtpGroupOutput) ToWirelessControllerWtpGroupOutputWithContext(ctx context.Context) WirelessControllerWtpGroupOutput {
	return o
}

func (o WirelessControllerWtpGroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerWtpGroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerWtpGroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerWtpGroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerWtpGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtpGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpGroupOutput) PlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelessControllerWtpGroup) pulumi.StringOutput { return v.PlatformType }).(pulumi.StringOutput)
}

func (o WirelessControllerWtpGroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelessControllerWtpGroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o WirelessControllerWtpGroupOutput) Wtps() WirelessControllerWtpGroupWtpArrayOutput {
	return o.ApplyT(func(v *WirelessControllerWtpGroup) WirelessControllerWtpGroupWtpArrayOutput { return v.Wtps }).(WirelessControllerWtpGroupWtpArrayOutput)
}

type WirelessControllerWtpGroupArrayOutput struct{ *pulumi.OutputState }

func (WirelessControllerWtpGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelessControllerWtpGroup)(nil)).Elem()
}

func (o WirelessControllerWtpGroupArrayOutput) ToWirelessControllerWtpGroupArrayOutput() WirelessControllerWtpGroupArrayOutput {
	return o
}

func (o WirelessControllerWtpGroupArrayOutput) ToWirelessControllerWtpGroupArrayOutputWithContext(ctx context.Context) WirelessControllerWtpGroupArrayOutput {
	return o
}

func (o WirelessControllerWtpGroupArrayOutput) Index(i pulumi.IntInput) WirelessControllerWtpGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelessControllerWtpGroup {
		return vs[0].([]*WirelessControllerWtpGroup)[vs[1].(int)]
	}).(WirelessControllerWtpGroupOutput)
}

type WirelessControllerWtpGroupMapOutput struct{ *pulumi.OutputState }

func (WirelessControllerWtpGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelessControllerWtpGroup)(nil)).Elem()
}

func (o WirelessControllerWtpGroupMapOutput) ToWirelessControllerWtpGroupMapOutput() WirelessControllerWtpGroupMapOutput {
	return o
}

func (o WirelessControllerWtpGroupMapOutput) ToWirelessControllerWtpGroupMapOutputWithContext(ctx context.Context) WirelessControllerWtpGroupMapOutput {
	return o
}

func (o WirelessControllerWtpGroupMapOutput) MapIndex(k pulumi.StringInput) WirelessControllerWtpGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelessControllerWtpGroup {
		return vs[0].(map[string]*WirelessControllerWtpGroup)[vs[1].(string)]
	}).(WirelessControllerWtpGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerWtpGroupInput)(nil)).Elem(), &WirelessControllerWtpGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerWtpGroupArrayInput)(nil)).Elem(), WirelessControllerWtpGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelessControllerWtpGroupMapInput)(nil)).Elem(), WirelessControllerWtpGroupMap{})
	pulumi.RegisterOutputType(WirelessControllerWtpGroupOutput{})
	pulumi.RegisterOutputType(WirelessControllerWtpGroupArrayOutput{})
	pulumi.RegisterOutputType(WirelessControllerWtpGroupMapOutput{})
}
