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

type FirewallWildcardFqdnGroup struct {
	pulumi.CustomResourceState

	Color               pulumi.IntOutput                           `pulumi:"color"`
	Comment             pulumi.StringPtrOutput                     `pulumi:"comment"`
	DynamicSortSubtable pulumi.StringPtrOutput                     `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                     `pulumi:"getAllTables"`
	Members             FirewallWildcardFqdnGroupMemberArrayOutput `pulumi:"members"`
	Name                pulumi.StringOutput                        `pulumi:"name"`
	Uuid                pulumi.StringOutput                        `pulumi:"uuid"`
	Vdomparam           pulumi.StringPtrOutput                     `pulumi:"vdomparam"`
	Visibility          pulumi.StringOutput                        `pulumi:"visibility"`
}

// NewFirewallWildcardFqdnGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallWildcardFqdnGroup(ctx *pulumi.Context,
	name string, args *FirewallWildcardFqdnGroupArgs, opts ...pulumi.ResourceOption) (*FirewallWildcardFqdnGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallWildcardFqdnGroup
	err := ctx.RegisterResource("fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallWildcardFqdnGroup gets an existing FirewallWildcardFqdnGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallWildcardFqdnGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallWildcardFqdnGroupState, opts ...pulumi.ResourceOption) (*FirewallWildcardFqdnGroup, error) {
	var resource FirewallWildcardFqdnGroup
	err := ctx.ReadResource("fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallWildcardFqdnGroup resources.
type firewallWildcardFqdnGroupState struct {
	Color               *int                              `pulumi:"color"`
	Comment             *string                           `pulumi:"comment"`
	DynamicSortSubtable *string                           `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                           `pulumi:"getAllTables"`
	Members             []FirewallWildcardFqdnGroupMember `pulumi:"members"`
	Name                *string                           `pulumi:"name"`
	Uuid                *string                           `pulumi:"uuid"`
	Vdomparam           *string                           `pulumi:"vdomparam"`
	Visibility          *string                           `pulumi:"visibility"`
}

type FirewallWildcardFqdnGroupState struct {
	Color               pulumi.IntPtrInput
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Members             FirewallWildcardFqdnGroupMemberArrayInput
	Name                pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Visibility          pulumi.StringPtrInput
}

func (FirewallWildcardFqdnGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallWildcardFqdnGroupState)(nil)).Elem()
}

type firewallWildcardFqdnGroupArgs struct {
	Color               *int                              `pulumi:"color"`
	Comment             *string                           `pulumi:"comment"`
	DynamicSortSubtable *string                           `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                           `pulumi:"getAllTables"`
	Members             []FirewallWildcardFqdnGroupMember `pulumi:"members"`
	Name                *string                           `pulumi:"name"`
	Uuid                *string                           `pulumi:"uuid"`
	Vdomparam           *string                           `pulumi:"vdomparam"`
	Visibility          *string                           `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallWildcardFqdnGroup resource.
type FirewallWildcardFqdnGroupArgs struct {
	Color               pulumi.IntPtrInput
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Members             FirewallWildcardFqdnGroupMemberArrayInput
	Name                pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Visibility          pulumi.StringPtrInput
}

func (FirewallWildcardFqdnGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallWildcardFqdnGroupArgs)(nil)).Elem()
}

type FirewallWildcardFqdnGroupInput interface {
	pulumi.Input

	ToFirewallWildcardFqdnGroupOutput() FirewallWildcardFqdnGroupOutput
	ToFirewallWildcardFqdnGroupOutputWithContext(ctx context.Context) FirewallWildcardFqdnGroupOutput
}

func (*FirewallWildcardFqdnGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallWildcardFqdnGroup)(nil)).Elem()
}

func (i *FirewallWildcardFqdnGroup) ToFirewallWildcardFqdnGroupOutput() FirewallWildcardFqdnGroupOutput {
	return i.ToFirewallWildcardFqdnGroupOutputWithContext(context.Background())
}

func (i *FirewallWildcardFqdnGroup) ToFirewallWildcardFqdnGroupOutputWithContext(ctx context.Context) FirewallWildcardFqdnGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallWildcardFqdnGroupOutput)
}

func (i *FirewallWildcardFqdnGroup) ToOutput(ctx context.Context) pulumix.Output[*FirewallWildcardFqdnGroup] {
	return pulumix.Output[*FirewallWildcardFqdnGroup]{
		OutputState: i.ToFirewallWildcardFqdnGroupOutputWithContext(ctx).OutputState,
	}
}

// FirewallWildcardFqdnGroupArrayInput is an input type that accepts FirewallWildcardFqdnGroupArray and FirewallWildcardFqdnGroupArrayOutput values.
// You can construct a concrete instance of `FirewallWildcardFqdnGroupArrayInput` via:
//
//	FirewallWildcardFqdnGroupArray{ FirewallWildcardFqdnGroupArgs{...} }
type FirewallWildcardFqdnGroupArrayInput interface {
	pulumi.Input

	ToFirewallWildcardFqdnGroupArrayOutput() FirewallWildcardFqdnGroupArrayOutput
	ToFirewallWildcardFqdnGroupArrayOutputWithContext(context.Context) FirewallWildcardFqdnGroupArrayOutput
}

type FirewallWildcardFqdnGroupArray []FirewallWildcardFqdnGroupInput

func (FirewallWildcardFqdnGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallWildcardFqdnGroup)(nil)).Elem()
}

func (i FirewallWildcardFqdnGroupArray) ToFirewallWildcardFqdnGroupArrayOutput() FirewallWildcardFqdnGroupArrayOutput {
	return i.ToFirewallWildcardFqdnGroupArrayOutputWithContext(context.Background())
}

func (i FirewallWildcardFqdnGroupArray) ToFirewallWildcardFqdnGroupArrayOutputWithContext(ctx context.Context) FirewallWildcardFqdnGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallWildcardFqdnGroupArrayOutput)
}

func (i FirewallWildcardFqdnGroupArray) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallWildcardFqdnGroup] {
	return pulumix.Output[[]*FirewallWildcardFqdnGroup]{
		OutputState: i.ToFirewallWildcardFqdnGroupArrayOutputWithContext(ctx).OutputState,
	}
}

// FirewallWildcardFqdnGroupMapInput is an input type that accepts FirewallWildcardFqdnGroupMap and FirewallWildcardFqdnGroupMapOutput values.
// You can construct a concrete instance of `FirewallWildcardFqdnGroupMapInput` via:
//
//	FirewallWildcardFqdnGroupMap{ "key": FirewallWildcardFqdnGroupArgs{...} }
type FirewallWildcardFqdnGroupMapInput interface {
	pulumi.Input

	ToFirewallWildcardFqdnGroupMapOutput() FirewallWildcardFqdnGroupMapOutput
	ToFirewallWildcardFqdnGroupMapOutputWithContext(context.Context) FirewallWildcardFqdnGroupMapOutput
}

type FirewallWildcardFqdnGroupMap map[string]FirewallWildcardFqdnGroupInput

func (FirewallWildcardFqdnGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallWildcardFqdnGroup)(nil)).Elem()
}

func (i FirewallWildcardFqdnGroupMap) ToFirewallWildcardFqdnGroupMapOutput() FirewallWildcardFqdnGroupMapOutput {
	return i.ToFirewallWildcardFqdnGroupMapOutputWithContext(context.Background())
}

func (i FirewallWildcardFqdnGroupMap) ToFirewallWildcardFqdnGroupMapOutputWithContext(ctx context.Context) FirewallWildcardFqdnGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallWildcardFqdnGroupMapOutput)
}

func (i FirewallWildcardFqdnGroupMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallWildcardFqdnGroup] {
	return pulumix.Output[map[string]*FirewallWildcardFqdnGroup]{
		OutputState: i.ToFirewallWildcardFqdnGroupMapOutputWithContext(ctx).OutputState,
	}
}

type FirewallWildcardFqdnGroupOutput struct{ *pulumi.OutputState }

func (FirewallWildcardFqdnGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallWildcardFqdnGroup)(nil)).Elem()
}

func (o FirewallWildcardFqdnGroupOutput) ToFirewallWildcardFqdnGroupOutput() FirewallWildcardFqdnGroupOutput {
	return o
}

func (o FirewallWildcardFqdnGroupOutput) ToFirewallWildcardFqdnGroupOutputWithContext(ctx context.Context) FirewallWildcardFqdnGroupOutput {
	return o
}

func (o FirewallWildcardFqdnGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*FirewallWildcardFqdnGroup] {
	return pulumix.Output[*FirewallWildcardFqdnGroup]{
		OutputState: o.OutputState,
	}
}

func (o FirewallWildcardFqdnGroupOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

func (o FirewallWildcardFqdnGroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallWildcardFqdnGroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallWildcardFqdnGroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallWildcardFqdnGroupOutput) Members() FirewallWildcardFqdnGroupMemberArrayOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) FirewallWildcardFqdnGroupMemberArrayOutput { return v.Members }).(FirewallWildcardFqdnGroupMemberArrayOutput)
}

func (o FirewallWildcardFqdnGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallWildcardFqdnGroupOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o FirewallWildcardFqdnGroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o FirewallWildcardFqdnGroupOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallWildcardFqdnGroup) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type FirewallWildcardFqdnGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallWildcardFqdnGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallWildcardFqdnGroup)(nil)).Elem()
}

func (o FirewallWildcardFqdnGroupArrayOutput) ToFirewallWildcardFqdnGroupArrayOutput() FirewallWildcardFqdnGroupArrayOutput {
	return o
}

func (o FirewallWildcardFqdnGroupArrayOutput) ToFirewallWildcardFqdnGroupArrayOutputWithContext(ctx context.Context) FirewallWildcardFqdnGroupArrayOutput {
	return o
}

func (o FirewallWildcardFqdnGroupArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FirewallWildcardFqdnGroup] {
	return pulumix.Output[[]*FirewallWildcardFqdnGroup]{
		OutputState: o.OutputState,
	}
}

func (o FirewallWildcardFqdnGroupArrayOutput) Index(i pulumi.IntInput) FirewallWildcardFqdnGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallWildcardFqdnGroup {
		return vs[0].([]*FirewallWildcardFqdnGroup)[vs[1].(int)]
	}).(FirewallWildcardFqdnGroupOutput)
}

type FirewallWildcardFqdnGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallWildcardFqdnGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallWildcardFqdnGroup)(nil)).Elem()
}

func (o FirewallWildcardFqdnGroupMapOutput) ToFirewallWildcardFqdnGroupMapOutput() FirewallWildcardFqdnGroupMapOutput {
	return o
}

func (o FirewallWildcardFqdnGroupMapOutput) ToFirewallWildcardFqdnGroupMapOutputWithContext(ctx context.Context) FirewallWildcardFqdnGroupMapOutput {
	return o
}

func (o FirewallWildcardFqdnGroupMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FirewallWildcardFqdnGroup] {
	return pulumix.Output[map[string]*FirewallWildcardFqdnGroup]{
		OutputState: o.OutputState,
	}
}

func (o FirewallWildcardFqdnGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallWildcardFqdnGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallWildcardFqdnGroup {
		return vs[0].(map[string]*FirewallWildcardFqdnGroup)[vs[1].(string)]
	}).(FirewallWildcardFqdnGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallWildcardFqdnGroupInput)(nil)).Elem(), &FirewallWildcardFqdnGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallWildcardFqdnGroupArrayInput)(nil)).Elem(), FirewallWildcardFqdnGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallWildcardFqdnGroupMapInput)(nil)).Elem(), FirewallWildcardFqdnGroupMap{})
	pulumi.RegisterOutputType(FirewallWildcardFqdnGroupOutput{})
	pulumi.RegisterOutputType(FirewallWildcardFqdnGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallWildcardFqdnGroupMapOutput{})
}
