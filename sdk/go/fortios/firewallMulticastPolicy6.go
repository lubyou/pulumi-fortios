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

type FirewallMulticastPolicy6 struct {
	pulumi.CustomResourceState

	Action              pulumi.StringOutput                        `pulumi:"action"`
	AutoAsicOffload     pulumi.StringOutput                        `pulumi:"autoAsicOffload"`
	Comments            pulumi.StringPtrOutput                     `pulumi:"comments"`
	Dstaddrs            FirewallMulticastPolicy6DstaddrArrayOutput `pulumi:"dstaddrs"`
	Dstintf             pulumi.StringOutput                        `pulumi:"dstintf"`
	DynamicSortSubtable pulumi.StringPtrOutput                     `pulumi:"dynamicSortSubtable"`
	EndPort             pulumi.IntOutput                           `pulumi:"endPort"`
	Fosid               pulumi.IntOutput                           `pulumi:"fosid"`
	GetAllTables        pulumi.StringPtrOutput                     `pulumi:"getAllTables"`
	Logtraffic          pulumi.StringOutput                        `pulumi:"logtraffic"`
	Name                pulumi.StringOutput                        `pulumi:"name"`
	Protocol            pulumi.IntOutput                           `pulumi:"protocol"`
	Srcaddrs            FirewallMulticastPolicy6SrcaddrArrayOutput `pulumi:"srcaddrs"`
	Srcintf             pulumi.StringOutput                        `pulumi:"srcintf"`
	StartPort           pulumi.IntOutput                           `pulumi:"startPort"`
	Status              pulumi.StringOutput                        `pulumi:"status"`
	Uuid                pulumi.StringOutput                        `pulumi:"uuid"`
	Vdomparam           pulumi.StringPtrOutput                     `pulumi:"vdomparam"`
}

// NewFirewallMulticastPolicy6 registers a new resource with the given unique name, arguments, and options.
func NewFirewallMulticastPolicy6(ctx *pulumi.Context,
	name string, args *FirewallMulticastPolicy6Args, opts ...pulumi.ResourceOption) (*FirewallMulticastPolicy6, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dstaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Dstaddrs'")
	}
	if args.Dstintf == nil {
		return nil, errors.New("invalid value for required argument 'Dstintf'")
	}
	if args.Srcaddrs == nil {
		return nil, errors.New("invalid value for required argument 'Srcaddrs'")
	}
	if args.Srcintf == nil {
		return nil, errors.New("invalid value for required argument 'Srcintf'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallMulticastPolicy6
	err := ctx.RegisterResource("fortios:index/firewallMulticastPolicy6:FirewallMulticastPolicy6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallMulticastPolicy6 gets an existing FirewallMulticastPolicy6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallMulticastPolicy6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallMulticastPolicy6State, opts ...pulumi.ResourceOption) (*FirewallMulticastPolicy6, error) {
	var resource FirewallMulticastPolicy6
	err := ctx.ReadResource("fortios:index/firewallMulticastPolicy6:FirewallMulticastPolicy6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallMulticastPolicy6 resources.
type firewallMulticastPolicy6State struct {
	Action              *string                           `pulumi:"action"`
	AutoAsicOffload     *string                           `pulumi:"autoAsicOffload"`
	Comments            *string                           `pulumi:"comments"`
	Dstaddrs            []FirewallMulticastPolicy6Dstaddr `pulumi:"dstaddrs"`
	Dstintf             *string                           `pulumi:"dstintf"`
	DynamicSortSubtable *string                           `pulumi:"dynamicSortSubtable"`
	EndPort             *int                              `pulumi:"endPort"`
	Fosid               *int                              `pulumi:"fosid"`
	GetAllTables        *string                           `pulumi:"getAllTables"`
	Logtraffic          *string                           `pulumi:"logtraffic"`
	Name                *string                           `pulumi:"name"`
	Protocol            *int                              `pulumi:"protocol"`
	Srcaddrs            []FirewallMulticastPolicy6Srcaddr `pulumi:"srcaddrs"`
	Srcintf             *string                           `pulumi:"srcintf"`
	StartPort           *int                              `pulumi:"startPort"`
	Status              *string                           `pulumi:"status"`
	Uuid                *string                           `pulumi:"uuid"`
	Vdomparam           *string                           `pulumi:"vdomparam"`
}

type FirewallMulticastPolicy6State struct {
	Action              pulumi.StringPtrInput
	AutoAsicOffload     pulumi.StringPtrInput
	Comments            pulumi.StringPtrInput
	Dstaddrs            FirewallMulticastPolicy6DstaddrArrayInput
	Dstintf             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	EndPort             pulumi.IntPtrInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	Logtraffic          pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Protocol            pulumi.IntPtrInput
	Srcaddrs            FirewallMulticastPolicy6SrcaddrArrayInput
	Srcintf             pulumi.StringPtrInput
	StartPort           pulumi.IntPtrInput
	Status              pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallMulticastPolicy6State) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallMulticastPolicy6State)(nil)).Elem()
}

type firewallMulticastPolicy6Args struct {
	Action              *string                           `pulumi:"action"`
	AutoAsicOffload     *string                           `pulumi:"autoAsicOffload"`
	Comments            *string                           `pulumi:"comments"`
	Dstaddrs            []FirewallMulticastPolicy6Dstaddr `pulumi:"dstaddrs"`
	Dstintf             string                            `pulumi:"dstintf"`
	DynamicSortSubtable *string                           `pulumi:"dynamicSortSubtable"`
	EndPort             *int                              `pulumi:"endPort"`
	Fosid               *int                              `pulumi:"fosid"`
	GetAllTables        *string                           `pulumi:"getAllTables"`
	Logtraffic          *string                           `pulumi:"logtraffic"`
	Name                *string                           `pulumi:"name"`
	Protocol            *int                              `pulumi:"protocol"`
	Srcaddrs            []FirewallMulticastPolicy6Srcaddr `pulumi:"srcaddrs"`
	Srcintf             string                            `pulumi:"srcintf"`
	StartPort           *int                              `pulumi:"startPort"`
	Status              *string                           `pulumi:"status"`
	Uuid                *string                           `pulumi:"uuid"`
	Vdomparam           *string                           `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallMulticastPolicy6 resource.
type FirewallMulticastPolicy6Args struct {
	Action              pulumi.StringPtrInput
	AutoAsicOffload     pulumi.StringPtrInput
	Comments            pulumi.StringPtrInput
	Dstaddrs            FirewallMulticastPolicy6DstaddrArrayInput
	Dstintf             pulumi.StringInput
	DynamicSortSubtable pulumi.StringPtrInput
	EndPort             pulumi.IntPtrInput
	Fosid               pulumi.IntPtrInput
	GetAllTables        pulumi.StringPtrInput
	Logtraffic          pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	Protocol            pulumi.IntPtrInput
	Srcaddrs            FirewallMulticastPolicy6SrcaddrArrayInput
	Srcintf             pulumi.StringInput
	StartPort           pulumi.IntPtrInput
	Status              pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallMulticastPolicy6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallMulticastPolicy6Args)(nil)).Elem()
}

type FirewallMulticastPolicy6Input interface {
	pulumi.Input

	ToFirewallMulticastPolicy6Output() FirewallMulticastPolicy6Output
	ToFirewallMulticastPolicy6OutputWithContext(ctx context.Context) FirewallMulticastPolicy6Output
}

func (*FirewallMulticastPolicy6) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallMulticastPolicy6)(nil)).Elem()
}

func (i *FirewallMulticastPolicy6) ToFirewallMulticastPolicy6Output() FirewallMulticastPolicy6Output {
	return i.ToFirewallMulticastPolicy6OutputWithContext(context.Background())
}

func (i *FirewallMulticastPolicy6) ToFirewallMulticastPolicy6OutputWithContext(ctx context.Context) FirewallMulticastPolicy6Output {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMulticastPolicy6Output)
}

func (i *FirewallMulticastPolicy6) ToOutput(ctx context.Context) pulumix.Output[*FirewallMulticastPolicy6] {
	return pulumix.Output[*FirewallMulticastPolicy6]{
		OutputState: i.ToFirewallMulticastPolicy6OutputWithContext(ctx).OutputState,
	}
}

// FirewallMulticastPolicy6ArrayInput is an input type that accepts FirewallMulticastPolicy6Array and FirewallMulticastPolicy6ArrayOutput values.
// You can construct a concrete instance of `FirewallMulticastPolicy6ArrayInput` via:
//
//	FirewallMulticastPolicy6Array{ FirewallMulticastPolicy6Args{...} }
type FirewallMulticastPolicy6ArrayInput interface {
	pulumi.Input

	ToFirewallMulticastPolicy6ArrayOutput() FirewallMulticastPolicy6ArrayOutput
	ToFirewallMulticastPolicy6ArrayOutputWithContext(context.Context) FirewallMulticastPolicy6ArrayOutput
}

type FirewallMulticastPolicy6Array []FirewallMulticastPolicy6Input

func (FirewallMulticastPolicy6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallMulticastPolicy6)(nil)).Elem()
}

func (i FirewallMulticastPolicy6Array) ToFirewallMulticastPolicy6ArrayOutput() FirewallMulticastPolicy6ArrayOutput {
	return i.ToFirewallMulticastPolicy6ArrayOutputWithContext(context.Background())
}

func (i FirewallMulticastPolicy6Array) ToFirewallMulticastPolicy6ArrayOutputWithContext(ctx context.Context) FirewallMulticastPolicy6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMulticastPolicy6ArrayOutput)
}

func (i FirewallMulticastPolicy6Array) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallMulticastPolicy6] {
	return pulumix.Output[[]*FirewallMulticastPolicy6]{
		OutputState: i.ToFirewallMulticastPolicy6ArrayOutputWithContext(ctx).OutputState,
	}
}

// FirewallMulticastPolicy6MapInput is an input type that accepts FirewallMulticastPolicy6Map and FirewallMulticastPolicy6MapOutput values.
// You can construct a concrete instance of `FirewallMulticastPolicy6MapInput` via:
//
//	FirewallMulticastPolicy6Map{ "key": FirewallMulticastPolicy6Args{...} }
type FirewallMulticastPolicy6MapInput interface {
	pulumi.Input

	ToFirewallMulticastPolicy6MapOutput() FirewallMulticastPolicy6MapOutput
	ToFirewallMulticastPolicy6MapOutputWithContext(context.Context) FirewallMulticastPolicy6MapOutput
}

type FirewallMulticastPolicy6Map map[string]FirewallMulticastPolicy6Input

func (FirewallMulticastPolicy6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallMulticastPolicy6)(nil)).Elem()
}

func (i FirewallMulticastPolicy6Map) ToFirewallMulticastPolicy6MapOutput() FirewallMulticastPolicy6MapOutput {
	return i.ToFirewallMulticastPolicy6MapOutputWithContext(context.Background())
}

func (i FirewallMulticastPolicy6Map) ToFirewallMulticastPolicy6MapOutputWithContext(ctx context.Context) FirewallMulticastPolicy6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMulticastPolicy6MapOutput)
}

func (i FirewallMulticastPolicy6Map) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallMulticastPolicy6] {
	return pulumix.Output[map[string]*FirewallMulticastPolicy6]{
		OutputState: i.ToFirewallMulticastPolicy6MapOutputWithContext(ctx).OutputState,
	}
}

type FirewallMulticastPolicy6Output struct{ *pulumi.OutputState }

func (FirewallMulticastPolicy6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallMulticastPolicy6)(nil)).Elem()
}

func (o FirewallMulticastPolicy6Output) ToFirewallMulticastPolicy6Output() FirewallMulticastPolicy6Output {
	return o
}

func (o FirewallMulticastPolicy6Output) ToFirewallMulticastPolicy6OutputWithContext(ctx context.Context) FirewallMulticastPolicy6Output {
	return o
}

func (o FirewallMulticastPolicy6Output) ToOutput(ctx context.Context) pulumix.Output[*FirewallMulticastPolicy6] {
	return pulumix.Output[*FirewallMulticastPolicy6]{
		OutputState: o.OutputState,
	}
}

func (o FirewallMulticastPolicy6Output) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

func (o FirewallMulticastPolicy6Output) AutoAsicOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringOutput { return v.AutoAsicOffload }).(pulumi.StringOutput)
}

func (o FirewallMulticastPolicy6Output) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o FirewallMulticastPolicy6Output) Dstaddrs() FirewallMulticastPolicy6DstaddrArrayOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) FirewallMulticastPolicy6DstaddrArrayOutput { return v.Dstaddrs }).(FirewallMulticastPolicy6DstaddrArrayOutput)
}

func (o FirewallMulticastPolicy6Output) Dstintf() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringOutput { return v.Dstintf }).(pulumi.StringOutput)
}

func (o FirewallMulticastPolicy6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallMulticastPolicy6Output) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.IntOutput { return v.EndPort }).(pulumi.IntOutput)
}

func (o FirewallMulticastPolicy6Output) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o FirewallMulticastPolicy6Output) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallMulticastPolicy6Output) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringOutput { return v.Logtraffic }).(pulumi.StringOutput)
}

func (o FirewallMulticastPolicy6Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallMulticastPolicy6Output) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.IntOutput { return v.Protocol }).(pulumi.IntOutput)
}

func (o FirewallMulticastPolicy6Output) Srcaddrs() FirewallMulticastPolicy6SrcaddrArrayOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) FirewallMulticastPolicy6SrcaddrArrayOutput { return v.Srcaddrs }).(FirewallMulticastPolicy6SrcaddrArrayOutput)
}

func (o FirewallMulticastPolicy6Output) Srcintf() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringOutput { return v.Srcintf }).(pulumi.StringOutput)
}

func (o FirewallMulticastPolicy6Output) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.IntOutput { return v.StartPort }).(pulumi.IntOutput)
}

func (o FirewallMulticastPolicy6Output) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o FirewallMulticastPolicy6Output) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o FirewallMulticastPolicy6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallMulticastPolicy6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallMulticastPolicy6ArrayOutput struct{ *pulumi.OutputState }

func (FirewallMulticastPolicy6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallMulticastPolicy6)(nil)).Elem()
}

func (o FirewallMulticastPolicy6ArrayOutput) ToFirewallMulticastPolicy6ArrayOutput() FirewallMulticastPolicy6ArrayOutput {
	return o
}

func (o FirewallMulticastPolicy6ArrayOutput) ToFirewallMulticastPolicy6ArrayOutputWithContext(ctx context.Context) FirewallMulticastPolicy6ArrayOutput {
	return o
}

func (o FirewallMulticastPolicy6ArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallMulticastPolicy6] {
	return pulumix.Output[[]*FirewallMulticastPolicy6]{
		OutputState: o.OutputState,
	}
}

func (o FirewallMulticastPolicy6ArrayOutput) Index(i pulumi.IntInput) FirewallMulticastPolicy6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallMulticastPolicy6 {
		return vs[0].([]*FirewallMulticastPolicy6)[vs[1].(int)]
	}).(FirewallMulticastPolicy6Output)
}

type FirewallMulticastPolicy6MapOutput struct{ *pulumi.OutputState }

func (FirewallMulticastPolicy6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallMulticastPolicy6)(nil)).Elem()
}

func (o FirewallMulticastPolicy6MapOutput) ToFirewallMulticastPolicy6MapOutput() FirewallMulticastPolicy6MapOutput {
	return o
}

func (o FirewallMulticastPolicy6MapOutput) ToFirewallMulticastPolicy6MapOutputWithContext(ctx context.Context) FirewallMulticastPolicy6MapOutput {
	return o
}

func (o FirewallMulticastPolicy6MapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallMulticastPolicy6] {
	return pulumix.Output[map[string]*FirewallMulticastPolicy6]{
		OutputState: o.OutputState,
	}
}

func (o FirewallMulticastPolicy6MapOutput) MapIndex(k pulumi.StringInput) FirewallMulticastPolicy6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallMulticastPolicy6 {
		return vs[0].(map[string]*FirewallMulticastPolicy6)[vs[1].(string)]
	}).(FirewallMulticastPolicy6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMulticastPolicy6Input)(nil)).Elem(), &FirewallMulticastPolicy6{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMulticastPolicy6ArrayInput)(nil)).Elem(), FirewallMulticastPolicy6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMulticastPolicy6MapInput)(nil)).Elem(), FirewallMulticastPolicy6Map{})
	pulumi.RegisterOutputType(FirewallMulticastPolicy6Output{})
	pulumi.RegisterOutputType(FirewallMulticastPolicy6ArrayOutput{})
	pulumi.RegisterOutputType(FirewallMulticastPolicy6MapOutput{})
}
