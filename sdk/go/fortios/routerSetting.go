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

type RouterSetting struct {
	pulumi.CustomResourceState

	BgpDebugFlags            pulumi.StringOutput    `pulumi:"bgpDebugFlags"`
	Hostname                 pulumi.StringOutput    `pulumi:"hostname"`
	IgmpDebugFlags           pulumi.StringOutput    `pulumi:"igmpDebugFlags"`
	ImiDebugFlags            pulumi.StringOutput    `pulumi:"imiDebugFlags"`
	IsisDebugFlags           pulumi.StringOutput    `pulumi:"isisDebugFlags"`
	Ospf6DebugEventsFlags    pulumi.StringOutput    `pulumi:"ospf6DebugEventsFlags"`
	Ospf6DebugIfsmFlags      pulumi.StringOutput    `pulumi:"ospf6DebugIfsmFlags"`
	Ospf6DebugLsaFlags       pulumi.StringOutput    `pulumi:"ospf6DebugLsaFlags"`
	Ospf6DebugNfsmFlags      pulumi.StringOutput    `pulumi:"ospf6DebugNfsmFlags"`
	Ospf6DebugNsmFlags       pulumi.StringOutput    `pulumi:"ospf6DebugNsmFlags"`
	Ospf6DebugPacketFlags    pulumi.StringOutput    `pulumi:"ospf6DebugPacketFlags"`
	Ospf6DebugRouteFlags     pulumi.StringOutput    `pulumi:"ospf6DebugRouteFlags"`
	OspfDebugEventsFlags     pulumi.StringOutput    `pulumi:"ospfDebugEventsFlags"`
	OspfDebugIfsmFlags       pulumi.StringOutput    `pulumi:"ospfDebugIfsmFlags"`
	OspfDebugLsaFlags        pulumi.StringOutput    `pulumi:"ospfDebugLsaFlags"`
	OspfDebugNfsmFlags       pulumi.StringOutput    `pulumi:"ospfDebugNfsmFlags"`
	OspfDebugNsmFlags        pulumi.StringOutput    `pulumi:"ospfDebugNsmFlags"`
	OspfDebugPacketFlags     pulumi.StringOutput    `pulumi:"ospfDebugPacketFlags"`
	OspfDebugRouteFlags      pulumi.StringOutput    `pulumi:"ospfDebugRouteFlags"`
	PimdmDebugFlags          pulumi.StringOutput    `pulumi:"pimdmDebugFlags"`
	PimsmDebugJoinpruneFlags pulumi.StringOutput    `pulumi:"pimsmDebugJoinpruneFlags"`
	PimsmDebugSimpleFlags    pulumi.StringOutput    `pulumi:"pimsmDebugSimpleFlags"`
	PimsmDebugTimerFlags     pulumi.StringOutput    `pulumi:"pimsmDebugTimerFlags"`
	RipDebugFlags            pulumi.StringOutput    `pulumi:"ripDebugFlags"`
	RipngDebugFlags          pulumi.StringOutput    `pulumi:"ripngDebugFlags"`
	ShowFilter               pulumi.StringOutput    `pulumi:"showFilter"`
	Vdomparam                pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterSetting registers a new resource with the given unique name, arguments, and options.
func NewRouterSetting(ctx *pulumi.Context,
	name string, args *RouterSettingArgs, opts ...pulumi.ResourceOption) (*RouterSetting, error) {
	if args == nil {
		args = &RouterSettingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterSetting
	err := ctx.RegisterResource("fortios:index/routerSetting:RouterSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterSetting gets an existing RouterSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterSettingState, opts ...pulumi.ResourceOption) (*RouterSetting, error) {
	var resource RouterSetting
	err := ctx.ReadResource("fortios:index/routerSetting:RouterSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterSetting resources.
type routerSettingState struct {
	BgpDebugFlags            *string `pulumi:"bgpDebugFlags"`
	Hostname                 *string `pulumi:"hostname"`
	IgmpDebugFlags           *string `pulumi:"igmpDebugFlags"`
	ImiDebugFlags            *string `pulumi:"imiDebugFlags"`
	IsisDebugFlags           *string `pulumi:"isisDebugFlags"`
	Ospf6DebugEventsFlags    *string `pulumi:"ospf6DebugEventsFlags"`
	Ospf6DebugIfsmFlags      *string `pulumi:"ospf6DebugIfsmFlags"`
	Ospf6DebugLsaFlags       *string `pulumi:"ospf6DebugLsaFlags"`
	Ospf6DebugNfsmFlags      *string `pulumi:"ospf6DebugNfsmFlags"`
	Ospf6DebugNsmFlags       *string `pulumi:"ospf6DebugNsmFlags"`
	Ospf6DebugPacketFlags    *string `pulumi:"ospf6DebugPacketFlags"`
	Ospf6DebugRouteFlags     *string `pulumi:"ospf6DebugRouteFlags"`
	OspfDebugEventsFlags     *string `pulumi:"ospfDebugEventsFlags"`
	OspfDebugIfsmFlags       *string `pulumi:"ospfDebugIfsmFlags"`
	OspfDebugLsaFlags        *string `pulumi:"ospfDebugLsaFlags"`
	OspfDebugNfsmFlags       *string `pulumi:"ospfDebugNfsmFlags"`
	OspfDebugNsmFlags        *string `pulumi:"ospfDebugNsmFlags"`
	OspfDebugPacketFlags     *string `pulumi:"ospfDebugPacketFlags"`
	OspfDebugRouteFlags      *string `pulumi:"ospfDebugRouteFlags"`
	PimdmDebugFlags          *string `pulumi:"pimdmDebugFlags"`
	PimsmDebugJoinpruneFlags *string `pulumi:"pimsmDebugJoinpruneFlags"`
	PimsmDebugSimpleFlags    *string `pulumi:"pimsmDebugSimpleFlags"`
	PimsmDebugTimerFlags     *string `pulumi:"pimsmDebugTimerFlags"`
	RipDebugFlags            *string `pulumi:"ripDebugFlags"`
	RipngDebugFlags          *string `pulumi:"ripngDebugFlags"`
	ShowFilter               *string `pulumi:"showFilter"`
	Vdomparam                *string `pulumi:"vdomparam"`
}

type RouterSettingState struct {
	BgpDebugFlags            pulumi.StringPtrInput
	Hostname                 pulumi.StringPtrInput
	IgmpDebugFlags           pulumi.StringPtrInput
	ImiDebugFlags            pulumi.StringPtrInput
	IsisDebugFlags           pulumi.StringPtrInput
	Ospf6DebugEventsFlags    pulumi.StringPtrInput
	Ospf6DebugIfsmFlags      pulumi.StringPtrInput
	Ospf6DebugLsaFlags       pulumi.StringPtrInput
	Ospf6DebugNfsmFlags      pulumi.StringPtrInput
	Ospf6DebugNsmFlags       pulumi.StringPtrInput
	Ospf6DebugPacketFlags    pulumi.StringPtrInput
	Ospf6DebugRouteFlags     pulumi.StringPtrInput
	OspfDebugEventsFlags     pulumi.StringPtrInput
	OspfDebugIfsmFlags       pulumi.StringPtrInput
	OspfDebugLsaFlags        pulumi.StringPtrInput
	OspfDebugNfsmFlags       pulumi.StringPtrInput
	OspfDebugNsmFlags        pulumi.StringPtrInput
	OspfDebugPacketFlags     pulumi.StringPtrInput
	OspfDebugRouteFlags      pulumi.StringPtrInput
	PimdmDebugFlags          pulumi.StringPtrInput
	PimsmDebugJoinpruneFlags pulumi.StringPtrInput
	PimsmDebugSimpleFlags    pulumi.StringPtrInput
	PimsmDebugTimerFlags     pulumi.StringPtrInput
	RipDebugFlags            pulumi.StringPtrInput
	RipngDebugFlags          pulumi.StringPtrInput
	ShowFilter               pulumi.StringPtrInput
	Vdomparam                pulumi.StringPtrInput
}

func (RouterSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerSettingState)(nil)).Elem()
}

type routerSettingArgs struct {
	BgpDebugFlags            *string `pulumi:"bgpDebugFlags"`
	Hostname                 *string `pulumi:"hostname"`
	IgmpDebugFlags           *string `pulumi:"igmpDebugFlags"`
	ImiDebugFlags            *string `pulumi:"imiDebugFlags"`
	IsisDebugFlags           *string `pulumi:"isisDebugFlags"`
	Ospf6DebugEventsFlags    *string `pulumi:"ospf6DebugEventsFlags"`
	Ospf6DebugIfsmFlags      *string `pulumi:"ospf6DebugIfsmFlags"`
	Ospf6DebugLsaFlags       *string `pulumi:"ospf6DebugLsaFlags"`
	Ospf6DebugNfsmFlags      *string `pulumi:"ospf6DebugNfsmFlags"`
	Ospf6DebugNsmFlags       *string `pulumi:"ospf6DebugNsmFlags"`
	Ospf6DebugPacketFlags    *string `pulumi:"ospf6DebugPacketFlags"`
	Ospf6DebugRouteFlags     *string `pulumi:"ospf6DebugRouteFlags"`
	OspfDebugEventsFlags     *string `pulumi:"ospfDebugEventsFlags"`
	OspfDebugIfsmFlags       *string `pulumi:"ospfDebugIfsmFlags"`
	OspfDebugLsaFlags        *string `pulumi:"ospfDebugLsaFlags"`
	OspfDebugNfsmFlags       *string `pulumi:"ospfDebugNfsmFlags"`
	OspfDebugNsmFlags        *string `pulumi:"ospfDebugNsmFlags"`
	OspfDebugPacketFlags     *string `pulumi:"ospfDebugPacketFlags"`
	OspfDebugRouteFlags      *string `pulumi:"ospfDebugRouteFlags"`
	PimdmDebugFlags          *string `pulumi:"pimdmDebugFlags"`
	PimsmDebugJoinpruneFlags *string `pulumi:"pimsmDebugJoinpruneFlags"`
	PimsmDebugSimpleFlags    *string `pulumi:"pimsmDebugSimpleFlags"`
	PimsmDebugTimerFlags     *string `pulumi:"pimsmDebugTimerFlags"`
	RipDebugFlags            *string `pulumi:"ripDebugFlags"`
	RipngDebugFlags          *string `pulumi:"ripngDebugFlags"`
	ShowFilter               *string `pulumi:"showFilter"`
	Vdomparam                *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterSetting resource.
type RouterSettingArgs struct {
	BgpDebugFlags            pulumi.StringPtrInput
	Hostname                 pulumi.StringPtrInput
	IgmpDebugFlags           pulumi.StringPtrInput
	ImiDebugFlags            pulumi.StringPtrInput
	IsisDebugFlags           pulumi.StringPtrInput
	Ospf6DebugEventsFlags    pulumi.StringPtrInput
	Ospf6DebugIfsmFlags      pulumi.StringPtrInput
	Ospf6DebugLsaFlags       pulumi.StringPtrInput
	Ospf6DebugNfsmFlags      pulumi.StringPtrInput
	Ospf6DebugNsmFlags       pulumi.StringPtrInput
	Ospf6DebugPacketFlags    pulumi.StringPtrInput
	Ospf6DebugRouteFlags     pulumi.StringPtrInput
	OspfDebugEventsFlags     pulumi.StringPtrInput
	OspfDebugIfsmFlags       pulumi.StringPtrInput
	OspfDebugLsaFlags        pulumi.StringPtrInput
	OspfDebugNfsmFlags       pulumi.StringPtrInput
	OspfDebugNsmFlags        pulumi.StringPtrInput
	OspfDebugPacketFlags     pulumi.StringPtrInput
	OspfDebugRouteFlags      pulumi.StringPtrInput
	PimdmDebugFlags          pulumi.StringPtrInput
	PimsmDebugJoinpruneFlags pulumi.StringPtrInput
	PimsmDebugSimpleFlags    pulumi.StringPtrInput
	PimsmDebugTimerFlags     pulumi.StringPtrInput
	RipDebugFlags            pulumi.StringPtrInput
	RipngDebugFlags          pulumi.StringPtrInput
	ShowFilter               pulumi.StringPtrInput
	Vdomparam                pulumi.StringPtrInput
}

func (RouterSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerSettingArgs)(nil)).Elem()
}

type RouterSettingInput interface {
	pulumi.Input

	ToRouterSettingOutput() RouterSettingOutput
	ToRouterSettingOutputWithContext(ctx context.Context) RouterSettingOutput
}

func (*RouterSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterSetting)(nil)).Elem()
}

func (i *RouterSetting) ToRouterSettingOutput() RouterSettingOutput {
	return i.ToRouterSettingOutputWithContext(context.Background())
}

func (i *RouterSetting) ToRouterSettingOutputWithContext(ctx context.Context) RouterSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterSettingOutput)
}

func (i *RouterSetting) ToOutput(ctx context.Context) pulumix.Output[*RouterSetting] {
	return pulumix.Output[*RouterSetting]{
		OutputState: i.ToRouterSettingOutputWithContext(ctx).OutputState,
	}
}

// RouterSettingArrayInput is an input type that accepts RouterSettingArray and RouterSettingArrayOutput values.
// You can construct a concrete instance of `RouterSettingArrayInput` via:
//
//	RouterSettingArray{ RouterSettingArgs{...} }
type RouterSettingArrayInput interface {
	pulumi.Input

	ToRouterSettingArrayOutput() RouterSettingArrayOutput
	ToRouterSettingArrayOutputWithContext(context.Context) RouterSettingArrayOutput
}

type RouterSettingArray []RouterSettingInput

func (RouterSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterSetting)(nil)).Elem()
}

func (i RouterSettingArray) ToRouterSettingArrayOutput() RouterSettingArrayOutput {
	return i.ToRouterSettingArrayOutputWithContext(context.Background())
}

func (i RouterSettingArray) ToRouterSettingArrayOutputWithContext(ctx context.Context) RouterSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterSettingArrayOutput)
}

func (i RouterSettingArray) ToOutput(ctx context.Context) pulumix.Output[[]*RouterSetting] {
	return pulumix.Output[[]*RouterSetting]{
		OutputState: i.ToRouterSettingArrayOutputWithContext(ctx).OutputState,
	}
}

// RouterSettingMapInput is an input type that accepts RouterSettingMap and RouterSettingMapOutput values.
// You can construct a concrete instance of `RouterSettingMapInput` via:
//
//	RouterSettingMap{ "key": RouterSettingArgs{...} }
type RouterSettingMapInput interface {
	pulumi.Input

	ToRouterSettingMapOutput() RouterSettingMapOutput
	ToRouterSettingMapOutputWithContext(context.Context) RouterSettingMapOutput
}

type RouterSettingMap map[string]RouterSettingInput

func (RouterSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterSetting)(nil)).Elem()
}

func (i RouterSettingMap) ToRouterSettingMapOutput() RouterSettingMapOutput {
	return i.ToRouterSettingMapOutputWithContext(context.Background())
}

func (i RouterSettingMap) ToRouterSettingMapOutputWithContext(ctx context.Context) RouterSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterSettingMapOutput)
}

func (i RouterSettingMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouterSetting] {
	return pulumix.Output[map[string]*RouterSetting]{
		OutputState: i.ToRouterSettingMapOutputWithContext(ctx).OutputState,
	}
}

type RouterSettingOutput struct{ *pulumi.OutputState }

func (RouterSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterSetting)(nil)).Elem()
}

func (o RouterSettingOutput) ToRouterSettingOutput() RouterSettingOutput {
	return o
}

func (o RouterSettingOutput) ToRouterSettingOutputWithContext(ctx context.Context) RouterSettingOutput {
	return o
}

func (o RouterSettingOutput) ToOutput(ctx context.Context) pulumix.Output[*RouterSetting] {
	return pulumix.Output[*RouterSetting]{
		OutputState: o.OutputState,
	}
}

func (o RouterSettingOutput) BgpDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.BgpDebugFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) IgmpDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.IgmpDebugFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) ImiDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.ImiDebugFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) IsisDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.IsisDebugFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Ospf6DebugEventsFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugEventsFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Ospf6DebugIfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugIfsmFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Ospf6DebugLsaFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugLsaFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Ospf6DebugNfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugNfsmFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Ospf6DebugNsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugNsmFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Ospf6DebugPacketFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugPacketFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Ospf6DebugRouteFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.Ospf6DebugRouteFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) OspfDebugEventsFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugEventsFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) OspfDebugIfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugIfsmFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) OspfDebugLsaFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugLsaFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) OspfDebugNfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugNfsmFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) OspfDebugNsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugNsmFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) OspfDebugPacketFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugPacketFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) OspfDebugRouteFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.OspfDebugRouteFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) PimdmDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.PimdmDebugFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) PimsmDebugJoinpruneFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.PimsmDebugJoinpruneFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) PimsmDebugSimpleFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.PimsmDebugSimpleFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) PimsmDebugTimerFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.PimsmDebugTimerFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) RipDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.RipDebugFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) RipngDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.RipngDebugFlags }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) ShowFilter() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringOutput { return v.ShowFilter }).(pulumi.StringOutput)
}

func (o RouterSettingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterSetting) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterSettingArrayOutput struct{ *pulumi.OutputState }

func (RouterSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterSetting)(nil)).Elem()
}

func (o RouterSettingArrayOutput) ToRouterSettingArrayOutput() RouterSettingArrayOutput {
	return o
}

func (o RouterSettingArrayOutput) ToRouterSettingArrayOutputWithContext(ctx context.Context) RouterSettingArrayOutput {
	return o
}

func (o RouterSettingArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RouterSetting] {
	return pulumix.Output[[]*RouterSetting]{
		OutputState: o.OutputState,
	}
}

func (o RouterSettingArrayOutput) Index(i pulumi.IntInput) RouterSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterSetting {
		return vs[0].([]*RouterSetting)[vs[1].(int)]
	}).(RouterSettingOutput)
}

type RouterSettingMapOutput struct{ *pulumi.OutputState }

func (RouterSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterSetting)(nil)).Elem()
}

func (o RouterSettingMapOutput) ToRouterSettingMapOutput() RouterSettingMapOutput {
	return o
}

func (o RouterSettingMapOutput) ToRouterSettingMapOutputWithContext(ctx context.Context) RouterSettingMapOutput {
	return o
}

func (o RouterSettingMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RouterSetting] {
	return pulumix.Output[map[string]*RouterSetting]{
		OutputState: o.OutputState,
	}
}

func (o RouterSettingMapOutput) MapIndex(k pulumi.StringInput) RouterSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterSetting {
		return vs[0].(map[string]*RouterSetting)[vs[1].(string)]
	}).(RouterSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterSettingInput)(nil)).Elem(), &RouterSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterSettingArrayInput)(nil)).Elem(), RouterSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterSettingMapInput)(nil)).Elem(), RouterSettingMap{})
	pulumi.RegisterOutputType(RouterSettingOutput{})
	pulumi.RegisterOutputType(RouterSettingArrayOutput{})
	pulumi.RegisterOutputType(RouterSettingMapOutput{})
}
