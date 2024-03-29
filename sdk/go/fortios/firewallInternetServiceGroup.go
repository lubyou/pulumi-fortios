// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallInternetServiceGroup struct {
	pulumi.CustomResourceState

	Comment             pulumi.StringPtrOutput                        `pulumi:"comment"`
	Direction           pulumi.StringOutput                           `pulumi:"direction"`
	DynamicSortSubtable pulumi.StringPtrOutput                        `pulumi:"dynamicSortSubtable"`
	GetAllTables        pulumi.StringPtrOutput                        `pulumi:"getAllTables"`
	Members             FirewallInternetServiceGroupMemberArrayOutput `pulumi:"members"`
	Name                pulumi.StringOutput                           `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput                        `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceGroup(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceGroupArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceGroup, error) {
	if args == nil {
		args = &FirewallInternetServiceGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallInternetServiceGroup
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceGroup:FirewallInternetServiceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceGroup gets an existing FirewallInternetServiceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceGroupState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceGroup, error) {
	var resource FirewallInternetServiceGroup
	err := ctx.ReadResource("fortios:index/firewallInternetServiceGroup:FirewallInternetServiceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceGroup resources.
type firewallInternetServiceGroupState struct {
	Comment             *string                              `pulumi:"comment"`
	Direction           *string                              `pulumi:"direction"`
	DynamicSortSubtable *string                              `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                              `pulumi:"getAllTables"`
	Members             []FirewallInternetServiceGroupMember `pulumi:"members"`
	Name                *string                              `pulumi:"name"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
}

type FirewallInternetServiceGroupState struct {
	Comment             pulumi.StringPtrInput
	Direction           pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Members             FirewallInternetServiceGroupMemberArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallInternetServiceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceGroupState)(nil)).Elem()
}

type firewallInternetServiceGroupArgs struct {
	Comment             *string                              `pulumi:"comment"`
	Direction           *string                              `pulumi:"direction"`
	DynamicSortSubtable *string                              `pulumi:"dynamicSortSubtable"`
	GetAllTables        *string                              `pulumi:"getAllTables"`
	Members             []FirewallInternetServiceGroupMember `pulumi:"members"`
	Name                *string                              `pulumi:"name"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceGroup resource.
type FirewallInternetServiceGroupArgs struct {
	Comment             pulumi.StringPtrInput
	Direction           pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Members             FirewallInternetServiceGroupMemberArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallInternetServiceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceGroupArgs)(nil)).Elem()
}

type FirewallInternetServiceGroupInput interface {
	pulumi.Input

	ToFirewallInternetServiceGroupOutput() FirewallInternetServiceGroupOutput
	ToFirewallInternetServiceGroupOutputWithContext(ctx context.Context) FirewallInternetServiceGroupOutput
}

func (*FirewallInternetServiceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceGroup)(nil)).Elem()
}

func (i *FirewallInternetServiceGroup) ToFirewallInternetServiceGroupOutput() FirewallInternetServiceGroupOutput {
	return i.ToFirewallInternetServiceGroupOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceGroup) ToFirewallInternetServiceGroupOutputWithContext(ctx context.Context) FirewallInternetServiceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceGroupOutput)
}

// FirewallInternetServiceGroupArrayInput is an input type that accepts FirewallInternetServiceGroupArray and FirewallInternetServiceGroupArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceGroupArrayInput` via:
//
//	FirewallInternetServiceGroupArray{ FirewallInternetServiceGroupArgs{...} }
type FirewallInternetServiceGroupArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceGroupArrayOutput() FirewallInternetServiceGroupArrayOutput
	ToFirewallInternetServiceGroupArrayOutputWithContext(context.Context) FirewallInternetServiceGroupArrayOutput
}

type FirewallInternetServiceGroupArray []FirewallInternetServiceGroupInput

func (FirewallInternetServiceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceGroup)(nil)).Elem()
}

func (i FirewallInternetServiceGroupArray) ToFirewallInternetServiceGroupArrayOutput() FirewallInternetServiceGroupArrayOutput {
	return i.ToFirewallInternetServiceGroupArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceGroupArray) ToFirewallInternetServiceGroupArrayOutputWithContext(ctx context.Context) FirewallInternetServiceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceGroupArrayOutput)
}

// FirewallInternetServiceGroupMapInput is an input type that accepts FirewallInternetServiceGroupMap and FirewallInternetServiceGroupMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceGroupMapInput` via:
//
//	FirewallInternetServiceGroupMap{ "key": FirewallInternetServiceGroupArgs{...} }
type FirewallInternetServiceGroupMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceGroupMapOutput() FirewallInternetServiceGroupMapOutput
	ToFirewallInternetServiceGroupMapOutputWithContext(context.Context) FirewallInternetServiceGroupMapOutput
}

type FirewallInternetServiceGroupMap map[string]FirewallInternetServiceGroupInput

func (FirewallInternetServiceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceGroup)(nil)).Elem()
}

func (i FirewallInternetServiceGroupMap) ToFirewallInternetServiceGroupMapOutput() FirewallInternetServiceGroupMapOutput {
	return i.ToFirewallInternetServiceGroupMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceGroupMap) ToFirewallInternetServiceGroupMapOutputWithContext(ctx context.Context) FirewallInternetServiceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceGroupMapOutput)
}

type FirewallInternetServiceGroupOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceGroup)(nil)).Elem()
}

func (o FirewallInternetServiceGroupOutput) ToFirewallInternetServiceGroupOutput() FirewallInternetServiceGroupOutput {
	return o
}

func (o FirewallInternetServiceGroupOutput) ToFirewallInternetServiceGroupOutputWithContext(ctx context.Context) FirewallInternetServiceGroupOutput {
	return o
}

func (o FirewallInternetServiceGroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetServiceGroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallInternetServiceGroupOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInternetServiceGroup) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

func (o FirewallInternetServiceGroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetServiceGroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallInternetServiceGroupOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetServiceGroup) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallInternetServiceGroupOutput) Members() FirewallInternetServiceGroupMemberArrayOutput {
	return o.ApplyT(func(v *FirewallInternetServiceGroup) FirewallInternetServiceGroupMemberArrayOutput { return v.Members }).(FirewallInternetServiceGroupMemberArrayOutput)
}

func (o FirewallInternetServiceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInternetServiceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallInternetServiceGroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetServiceGroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallInternetServiceGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceGroup)(nil)).Elem()
}

func (o FirewallInternetServiceGroupArrayOutput) ToFirewallInternetServiceGroupArrayOutput() FirewallInternetServiceGroupArrayOutput {
	return o
}

func (o FirewallInternetServiceGroupArrayOutput) ToFirewallInternetServiceGroupArrayOutputWithContext(ctx context.Context) FirewallInternetServiceGroupArrayOutput {
	return o
}

func (o FirewallInternetServiceGroupArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetServiceGroup {
		return vs[0].([]*FirewallInternetServiceGroup)[vs[1].(int)]
	}).(FirewallInternetServiceGroupOutput)
}

type FirewallInternetServiceGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceGroup)(nil)).Elem()
}

func (o FirewallInternetServiceGroupMapOutput) ToFirewallInternetServiceGroupMapOutput() FirewallInternetServiceGroupMapOutput {
	return o
}

func (o FirewallInternetServiceGroupMapOutput) ToFirewallInternetServiceGroupMapOutputWithContext(ctx context.Context) FirewallInternetServiceGroupMapOutput {
	return o
}

func (o FirewallInternetServiceGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetServiceGroup {
		return vs[0].(map[string]*FirewallInternetServiceGroup)[vs[1].(string)]
	}).(FirewallInternetServiceGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceGroupInput)(nil)).Elem(), &FirewallInternetServiceGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceGroupArrayInput)(nil)).Elem(), FirewallInternetServiceGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceGroupMapInput)(nil)).Elem(), FirewallInternetServiceGroupMap{})
	pulumi.RegisterOutputType(FirewallInternetServiceGroupOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceGroupMapOutput{})
}
