// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type FirewallVip46 struct {
	pulumi.CustomResourceState

	ArpReply            pulumi.StringOutput                   `pulumi:"arpReply"`
	Color               pulumi.IntOutput                      `pulumi:"color"`
	Comment             pulumi.StringPtrOutput                `pulumi:"comment"`
	DynamicSortSubtable pulumi.StringPtrOutput                `pulumi:"dynamicSortSubtable"`
	Extip               pulumi.StringOutput                   `pulumi:"extip"`
	Extport             pulumi.StringOutput                   `pulumi:"extport"`
	Fosid               pulumi.IntOutput                      `pulumi:"fosid"`
	GetAllTables        pulumi.StringPtrOutput                `pulumi:"getAllTables"`
	LdbMethod           pulumi.StringOutput                   `pulumi:"ldbMethod"`
	Mappedip            pulumi.StringOutput                   `pulumi:"mappedip"`
	Mappedport          pulumi.StringOutput                   `pulumi:"mappedport"`
	Monitors            FirewallVip46MonitorArrayOutput       `pulumi:"monitors"`
	Name                pulumi.StringOutput                   `pulumi:"name"`
	Portforward         pulumi.StringOutput                   `pulumi:"portforward"`
	Protocol            pulumi.StringOutput                   `pulumi:"protocol"`
	Realservers         FirewallVip46RealserverArrayOutput    `pulumi:"realservers"`
	ServerType          pulumi.StringOutput                   `pulumi:"serverType"`
	SrcFilters          FirewallVip46SrcFilterArrayOutput     `pulumi:"srcFilters"`
	SrcintfFilters      FirewallVip46SrcintfFilterArrayOutput `pulumi:"srcintfFilters"`
	Type                pulumi.StringOutput                   `pulumi:"type"`
	Uuid                pulumi.StringOutput                   `pulumi:"uuid"`
	Vdomparam           pulumi.StringPtrOutput                `pulumi:"vdomparam"`
}

// NewFirewallVip46 registers a new resource with the given unique name, arguments, and options.
func NewFirewallVip46(ctx *pulumi.Context,
	name string, args *FirewallVip46Args, opts ...pulumi.ResourceOption) (*FirewallVip46, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Extip == nil {
		return nil, errors.New("invalid value for required argument 'Extip'")
	}
	if args.Mappedip == nil {
		return nil, errors.New("invalid value for required argument 'Mappedip'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallVip46
	err := ctx.RegisterResource("fortios:index/firewallVip46:FirewallVip46", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallVip46 gets an existing FirewallVip46 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallVip46(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallVip46State, opts ...pulumi.ResourceOption) (*FirewallVip46, error) {
	var resource FirewallVip46
	err := ctx.ReadResource("fortios:index/firewallVip46:FirewallVip46", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallVip46 resources.
type firewallVip46State struct {
	ArpReply            *string                      `pulumi:"arpReply"`
	Color               *int                         `pulumi:"color"`
	Comment             *string                      `pulumi:"comment"`
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	Extip               *string                      `pulumi:"extip"`
	Extport             *string                      `pulumi:"extport"`
	Fosid               *int                         `pulumi:"fosid"`
	GetAllTables        *string                      `pulumi:"getAllTables"`
	LdbMethod           *string                      `pulumi:"ldbMethod"`
	Mappedip            *string                      `pulumi:"mappedip"`
	Mappedport          *string                      `pulumi:"mappedport"`
	Monitors            []FirewallVip46Monitor       `pulumi:"monitors"`
	Name                *string                      `pulumi:"name"`
	Portforward         *string                      `pulumi:"portforward"`
	Protocol            *string                      `pulumi:"protocol"`
	Realservers         []FirewallVip46Realserver    `pulumi:"realservers"`
	ServerType          *string                      `pulumi:"serverType"`
	SrcFilters          []FirewallVip46SrcFilter     `pulumi:"srcFilters"`
	SrcintfFilters      []FirewallVip46SrcintfFilter `pulumi:"srcintfFilters"`
	Type                *string                      `pulumi:"type"`
	Uuid                *string                      `pulumi:"uuid"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

type FirewallVip46State struct {
	ArpReply            pulumi.StringPtrInput
	Color               pulumi.IntPtrInput
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Extip               pulumi.StringPtrInput
	Extport             pulumi.StringPtrInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	LdbMethod           pulumi.StringPtrInput
	Mappedip            pulumi.StringPtrInput
	Mappedport          pulumi.StringPtrInput
	Monitors            FirewallVip46MonitorArrayInput
	Name                pulumi.StringPtrInput
	Portforward         pulumi.StringPtrInput
	Protocol            pulumi.StringPtrInput
	Realservers         FirewallVip46RealserverArrayInput
	ServerType          pulumi.StringPtrInput
	SrcFilters          FirewallVip46SrcFilterArrayInput
	SrcintfFilters      FirewallVip46SrcintfFilterArrayInput
	Type                pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallVip46State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVip46State)(nil)).Elem()
}

type firewallVip46Args struct {
	ArpReply            *string                      `pulumi:"arpReply"`
	Color               *int                         `pulumi:"color"`
	Comment             *string                      `pulumi:"comment"`
	DynamicSortSubtable *string                      `pulumi:"dynamicSortSubtable"`
	Extip               string                       `pulumi:"extip"`
	Extport             *string                      `pulumi:"extport"`
	Fosid               *int                         `pulumi:"fosid"`
	GetAllTables        *string                      `pulumi:"getAllTables"`
	LdbMethod           *string                      `pulumi:"ldbMethod"`
	Mappedip            string                       `pulumi:"mappedip"`
	Mappedport          *string                      `pulumi:"mappedport"`
	Monitors            []FirewallVip46Monitor       `pulumi:"monitors"`
	Name                *string                      `pulumi:"name"`
	Portforward         *string                      `pulumi:"portforward"`
	Protocol            *string                      `pulumi:"protocol"`
	Realservers         []FirewallVip46Realserver    `pulumi:"realservers"`
	ServerType          *string                      `pulumi:"serverType"`
	SrcFilters          []FirewallVip46SrcFilter     `pulumi:"srcFilters"`
	SrcintfFilters      []FirewallVip46SrcintfFilter `pulumi:"srcintfFilters"`
	Type                *string                      `pulumi:"type"`
	Uuid                *string                      `pulumi:"uuid"`
	Vdomparam           *string                      `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallVip46 resource.
type FirewallVip46Args struct {
	ArpReply            pulumi.StringPtrInput
	Color               pulumi.IntPtrInput
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Extip               pulumi.StringInput
	Extport             pulumi.StringPtrInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	LdbMethod           pulumi.StringPtrInput
	Mappedip            pulumi.StringInput
	Mappedport          pulumi.StringPtrInput
	Monitors            FirewallVip46MonitorArrayInput
	Name                pulumi.StringPtrInput
	Portforward         pulumi.StringPtrInput
	Protocol            pulumi.StringPtrInput
	Realservers         FirewallVip46RealserverArrayInput
	ServerType          pulumi.StringPtrInput
	SrcFilters          FirewallVip46SrcFilterArrayInput
	SrcintfFilters      FirewallVip46SrcintfFilterArrayInput
	Type                pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallVip46Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallVip46Args)(nil)).Elem()
}

type FirewallVip46Input interface {
	pulumi.Input

	ToFirewallVip46Output() FirewallVip46Output
	ToFirewallVip46OutputWithContext(ctx context.Context) FirewallVip46Output
}

func (*FirewallVip46) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVip46)(nil)).Elem()
}

func (i *FirewallVip46) ToFirewallVip46Output() FirewallVip46Output {
	return i.ToFirewallVip46OutputWithContext(context.Background())
}

func (i *FirewallVip46) ToFirewallVip46OutputWithContext(ctx context.Context) FirewallVip46Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip46Output)
}

func (i *FirewallVip46) ToOutput(ctx context.Context) pulumix.Output[*FirewallVip46] {
	return pulumix.Output[*FirewallVip46]{
		OutputState: i.ToFirewallVip46OutputWithContext(ctx).OutputState,
	}
}

// FirewallVip46ArrayInput is an input type that accepts FirewallVip46Array and FirewallVip46ArrayOutput values.
// You can construct a concrete instance of `FirewallVip46ArrayInput` via:
//
//	FirewallVip46Array{ FirewallVip46Args{...} }
type FirewallVip46ArrayInput interface {
	pulumi.Input

	ToFirewallVip46ArrayOutput() FirewallVip46ArrayOutput
	ToFirewallVip46ArrayOutputWithContext(context.Context) FirewallVip46ArrayOutput
}

type FirewallVip46Array []FirewallVip46Input

func (FirewallVip46Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVip46)(nil)).Elem()
}

func (i FirewallVip46Array) ToFirewallVip46ArrayOutput() FirewallVip46ArrayOutput {
	return i.ToFirewallVip46ArrayOutputWithContext(context.Background())
}

func (i FirewallVip46Array) ToFirewallVip46ArrayOutputWithContext(ctx context.Context) FirewallVip46ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip46ArrayOutput)
}

func (i FirewallVip46Array) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallVip46] {
	return pulumix.Output[[]*FirewallVip46]{
		OutputState: i.ToFirewallVip46ArrayOutputWithContext(ctx).OutputState,
	}
}

// FirewallVip46MapInput is an input type that accepts FirewallVip46Map and FirewallVip46MapOutput values.
// You can construct a concrete instance of `FirewallVip46MapInput` via:
//
//	FirewallVip46Map{ "key": FirewallVip46Args{...} }
type FirewallVip46MapInput interface {
	pulumi.Input

	ToFirewallVip46MapOutput() FirewallVip46MapOutput
	ToFirewallVip46MapOutputWithContext(context.Context) FirewallVip46MapOutput
}

type FirewallVip46Map map[string]FirewallVip46Input

func (FirewallVip46Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVip46)(nil)).Elem()
}

func (i FirewallVip46Map) ToFirewallVip46MapOutput() FirewallVip46MapOutput {
	return i.ToFirewallVip46MapOutputWithContext(context.Background())
}

func (i FirewallVip46Map) ToFirewallVip46MapOutputWithContext(ctx context.Context) FirewallVip46MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallVip46MapOutput)
}

func (i FirewallVip46Map) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallVip46] {
	return pulumix.Output[map[string]*FirewallVip46]{
		OutputState: i.ToFirewallVip46MapOutputWithContext(ctx).OutputState,
	}
}

type FirewallVip46Output struct{ *pulumi.OutputState }

func (FirewallVip46Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallVip46)(nil)).Elem()
}

func (o FirewallVip46Output) ToFirewallVip46Output() FirewallVip46Output {
	return o
}

func (o FirewallVip46Output) ToFirewallVip46OutputWithContext(ctx context.Context) FirewallVip46Output {
	return o
}

func (o FirewallVip46Output) ToOutput(ctx context.Context) pulumix.Output[*FirewallVip46] {
	return pulumix.Output[*FirewallVip46]{
		OutputState: o.OutputState,
	}
}

func (o FirewallVip46Output) ArpReply() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.ArpReply }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

func (o FirewallVip46Output) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallVip46Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallVip46Output) Extip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Extip }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Extport() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Extport }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o FirewallVip46Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallVip46Output) LdbMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.LdbMethod }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Mappedip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Mappedip }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Mappedport() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Mappedport }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Monitors() FirewallVip46MonitorArrayOutput {
	return o.ApplyT(func(v *FirewallVip46) FirewallVip46MonitorArrayOutput { return v.Monitors }).(FirewallVip46MonitorArrayOutput)
}

func (o FirewallVip46Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Portforward() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Portforward }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Realservers() FirewallVip46RealserverArrayOutput {
	return o.ApplyT(func(v *FirewallVip46) FirewallVip46RealserverArrayOutput { return v.Realservers }).(FirewallVip46RealserverArrayOutput)
}

func (o FirewallVip46Output) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) SrcFilters() FirewallVip46SrcFilterArrayOutput {
	return o.ApplyT(func(v *FirewallVip46) FirewallVip46SrcFilterArrayOutput { return v.SrcFilters }).(FirewallVip46SrcFilterArrayOutput)
}

func (o FirewallVip46Output) SrcintfFilters() FirewallVip46SrcintfFilterArrayOutput {
	return o.ApplyT(func(v *FirewallVip46) FirewallVip46SrcintfFilterArrayOutput { return v.SrcintfFilters }).(FirewallVip46SrcintfFilterArrayOutput)
}

func (o FirewallVip46Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o FirewallVip46Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallVip46) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallVip46ArrayOutput struct{ *pulumi.OutputState }

func (FirewallVip46ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallVip46)(nil)).Elem()
}

func (o FirewallVip46ArrayOutput) ToFirewallVip46ArrayOutput() FirewallVip46ArrayOutput {
	return o
}

func (o FirewallVip46ArrayOutput) ToFirewallVip46ArrayOutputWithContext(ctx context.Context) FirewallVip46ArrayOutput {
	return o
}

func (o FirewallVip46ArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallVip46] {
	return pulumix.Output[[]*FirewallVip46]{
		OutputState: o.OutputState,
	}
}

func (o FirewallVip46ArrayOutput) Index(i pulumi.IntInput) FirewallVip46Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallVip46 {
		return vs[0].([]*FirewallVip46)[vs[1].(int)]
	}).(FirewallVip46Output)
}

type FirewallVip46MapOutput struct{ *pulumi.OutputState }

func (FirewallVip46MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallVip46)(nil)).Elem()
}

func (o FirewallVip46MapOutput) ToFirewallVip46MapOutput() FirewallVip46MapOutput {
	return o
}

func (o FirewallVip46MapOutput) ToFirewallVip46MapOutputWithContext(ctx context.Context) FirewallVip46MapOutput {
	return o
}

func (o FirewallVip46MapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallVip46] {
	return pulumix.Output[map[string]*FirewallVip46]{
		OutputState: o.OutputState,
	}
}

func (o FirewallVip46MapOutput) MapIndex(k pulumi.StringInput) FirewallVip46Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallVip46 {
		return vs[0].(map[string]*FirewallVip46)[vs[1].(string)]
	}).(FirewallVip46Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVip46Input)(nil)).Elem(), &FirewallVip46{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVip46ArrayInput)(nil)).Elem(), FirewallVip46Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallVip46MapInput)(nil)).Elem(), FirewallVip46Map{})
	pulumi.RegisterOutputType(FirewallVip46Output{})
	pulumi.RegisterOutputType(FirewallVip46ArrayOutput{})
	pulumi.RegisterOutputType(FirewallVip46MapOutput{})
}
