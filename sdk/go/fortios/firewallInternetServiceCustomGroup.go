// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallInternetServiceCustomGroup struct {
	pulumi.CustomResourceState

	Comment             pulumi.StringPtrOutput                              `pulumi:"comment"`
	DynamicSortSubtable pulumi.StringPtrOutput                              `pulumi:"dynamicSortSubtable"`
	Members             FirewallInternetServiceCustomGroupMemberArrayOutput `pulumi:"members"`
	Name                pulumi.StringOutput                                 `pulumi:"name"`
	Vdomparam           pulumi.StringPtrOutput                              `pulumi:"vdomparam"`
}

// NewFirewallInternetServiceCustomGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetServiceCustomGroup(ctx *pulumi.Context,
	name string, args *FirewallInternetServiceCustomGroupArgs, opts ...pulumi.ResourceOption) (*FirewallInternetServiceCustomGroup, error) {
	if args == nil {
		args = &FirewallInternetServiceCustomGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetServiceCustomGroup
	err := ctx.RegisterResource("fortios:index/firewallInternetServiceCustomGroup:FirewallInternetServiceCustomGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetServiceCustomGroup gets an existing FirewallInternetServiceCustomGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetServiceCustomGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetServiceCustomGroupState, opts ...pulumi.ResourceOption) (*FirewallInternetServiceCustomGroup, error) {
	var resource FirewallInternetServiceCustomGroup
	err := ctx.ReadResource("fortios:index/firewallInternetServiceCustomGroup:FirewallInternetServiceCustomGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetServiceCustomGroup resources.
type firewallInternetServiceCustomGroupState struct {
	Comment             *string                                    `pulumi:"comment"`
	DynamicSortSubtable *string                                    `pulumi:"dynamicSortSubtable"`
	Members             []FirewallInternetServiceCustomGroupMember `pulumi:"members"`
	Name                *string                                    `pulumi:"name"`
	Vdomparam           *string                                    `pulumi:"vdomparam"`
}

type FirewallInternetServiceCustomGroupState struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Members             FirewallInternetServiceCustomGroupMemberArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallInternetServiceCustomGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceCustomGroupState)(nil)).Elem()
}

type firewallInternetServiceCustomGroupArgs struct {
	Comment             *string                                    `pulumi:"comment"`
	DynamicSortSubtable *string                                    `pulumi:"dynamicSortSubtable"`
	Members             []FirewallInternetServiceCustomGroupMember `pulumi:"members"`
	Name                *string                                    `pulumi:"name"`
	Vdomparam           *string                                    `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetServiceCustomGroup resource.
type FirewallInternetServiceCustomGroupArgs struct {
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Members             FirewallInternetServiceCustomGroupMemberArrayInput
	Name                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
}

func (FirewallInternetServiceCustomGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetServiceCustomGroupArgs)(nil)).Elem()
}

type FirewallInternetServiceCustomGroupInput interface {
	pulumi.Input

	ToFirewallInternetServiceCustomGroupOutput() FirewallInternetServiceCustomGroupOutput
	ToFirewallInternetServiceCustomGroupOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupOutput
}

func (*FirewallInternetServiceCustomGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceCustomGroup)(nil)).Elem()
}

func (i *FirewallInternetServiceCustomGroup) ToFirewallInternetServiceCustomGroupOutput() FirewallInternetServiceCustomGroupOutput {
	return i.ToFirewallInternetServiceCustomGroupOutputWithContext(context.Background())
}

func (i *FirewallInternetServiceCustomGroup) ToFirewallInternetServiceCustomGroupOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceCustomGroupOutput)
}

// FirewallInternetServiceCustomGroupArrayInput is an input type that accepts FirewallInternetServiceCustomGroupArray and FirewallInternetServiceCustomGroupArrayOutput values.
// You can construct a concrete instance of `FirewallInternetServiceCustomGroupArrayInput` via:
//
//	FirewallInternetServiceCustomGroupArray{ FirewallInternetServiceCustomGroupArgs{...} }
type FirewallInternetServiceCustomGroupArrayInput interface {
	pulumi.Input

	ToFirewallInternetServiceCustomGroupArrayOutput() FirewallInternetServiceCustomGroupArrayOutput
	ToFirewallInternetServiceCustomGroupArrayOutputWithContext(context.Context) FirewallInternetServiceCustomGroupArrayOutput
}

type FirewallInternetServiceCustomGroupArray []FirewallInternetServiceCustomGroupInput

func (FirewallInternetServiceCustomGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceCustomGroup)(nil)).Elem()
}

func (i FirewallInternetServiceCustomGroupArray) ToFirewallInternetServiceCustomGroupArrayOutput() FirewallInternetServiceCustomGroupArrayOutput {
	return i.ToFirewallInternetServiceCustomGroupArrayOutputWithContext(context.Background())
}

func (i FirewallInternetServiceCustomGroupArray) ToFirewallInternetServiceCustomGroupArrayOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceCustomGroupArrayOutput)
}

// FirewallInternetServiceCustomGroupMapInput is an input type that accepts FirewallInternetServiceCustomGroupMap and FirewallInternetServiceCustomGroupMapOutput values.
// You can construct a concrete instance of `FirewallInternetServiceCustomGroupMapInput` via:
//
//	FirewallInternetServiceCustomGroupMap{ "key": FirewallInternetServiceCustomGroupArgs{...} }
type FirewallInternetServiceCustomGroupMapInput interface {
	pulumi.Input

	ToFirewallInternetServiceCustomGroupMapOutput() FirewallInternetServiceCustomGroupMapOutput
	ToFirewallInternetServiceCustomGroupMapOutputWithContext(context.Context) FirewallInternetServiceCustomGroupMapOutput
}

type FirewallInternetServiceCustomGroupMap map[string]FirewallInternetServiceCustomGroupInput

func (FirewallInternetServiceCustomGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceCustomGroup)(nil)).Elem()
}

func (i FirewallInternetServiceCustomGroupMap) ToFirewallInternetServiceCustomGroupMapOutput() FirewallInternetServiceCustomGroupMapOutput {
	return i.ToFirewallInternetServiceCustomGroupMapOutputWithContext(context.Background())
}

func (i FirewallInternetServiceCustomGroupMap) ToFirewallInternetServiceCustomGroupMapOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetServiceCustomGroupMapOutput)
}

type FirewallInternetServiceCustomGroupOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceCustomGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetServiceCustomGroup)(nil)).Elem()
}

func (o FirewallInternetServiceCustomGroupOutput) ToFirewallInternetServiceCustomGroupOutput() FirewallInternetServiceCustomGroupOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupOutput) ToFirewallInternetServiceCustomGroupOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetServiceCustomGroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallInternetServiceCustomGroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetServiceCustomGroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallInternetServiceCustomGroupOutput) Members() FirewallInternetServiceCustomGroupMemberArrayOutput {
	return o.ApplyT(func(v *FirewallInternetServiceCustomGroup) FirewallInternetServiceCustomGroupMemberArrayOutput {
		return v.Members
	}).(FirewallInternetServiceCustomGroupMemberArrayOutput)
}

func (o FirewallInternetServiceCustomGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInternetServiceCustomGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallInternetServiceCustomGroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetServiceCustomGroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallInternetServiceCustomGroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceCustomGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetServiceCustomGroup)(nil)).Elem()
}

func (o FirewallInternetServiceCustomGroupArrayOutput) ToFirewallInternetServiceCustomGroupArrayOutput() FirewallInternetServiceCustomGroupArrayOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupArrayOutput) ToFirewallInternetServiceCustomGroupArrayOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupArrayOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupArrayOutput) Index(i pulumi.IntInput) FirewallInternetServiceCustomGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetServiceCustomGroup {
		return vs[0].([]*FirewallInternetServiceCustomGroup)[vs[1].(int)]
	}).(FirewallInternetServiceCustomGroupOutput)
}

type FirewallInternetServiceCustomGroupMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetServiceCustomGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetServiceCustomGroup)(nil)).Elem()
}

func (o FirewallInternetServiceCustomGroupMapOutput) ToFirewallInternetServiceCustomGroupMapOutput() FirewallInternetServiceCustomGroupMapOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupMapOutput) ToFirewallInternetServiceCustomGroupMapOutputWithContext(ctx context.Context) FirewallInternetServiceCustomGroupMapOutput {
	return o
}

func (o FirewallInternetServiceCustomGroupMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetServiceCustomGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetServiceCustomGroup {
		return vs[0].(map[string]*FirewallInternetServiceCustomGroup)[vs[1].(string)]
	}).(FirewallInternetServiceCustomGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceCustomGroupInput)(nil)).Elem(), &FirewallInternetServiceCustomGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceCustomGroupArrayInput)(nil)).Elem(), FirewallInternetServiceCustomGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetServiceCustomGroupMapInput)(nil)).Elem(), FirewallInternetServiceCustomGroupMap{})
	pulumi.RegisterOutputType(FirewallInternetServiceCustomGroupOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceCustomGroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetServiceCustomGroupMapOutput{})
}
