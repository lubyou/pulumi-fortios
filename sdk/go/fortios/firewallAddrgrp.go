// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallAddrgrp struct {
	pulumi.CustomResourceState

	AllowRouting        pulumi.StringOutput                     `pulumi:"allowRouting"`
	Category            pulumi.StringOutput                     `pulumi:"category"`
	Color               pulumi.IntOutput                        `pulumi:"color"`
	Comment             pulumi.StringPtrOutput                  `pulumi:"comment"`
	DynamicSortSubtable pulumi.StringPtrOutput                  `pulumi:"dynamicSortSubtable"`
	Exclude             pulumi.StringOutput                     `pulumi:"exclude"`
	ExcludeMembers      FirewallAddrgrpExcludeMemberArrayOutput `pulumi:"excludeMembers"`
	FabricObject        pulumi.StringOutput                     `pulumi:"fabricObject"`
	GetAllTables        pulumi.StringPtrOutput                  `pulumi:"getAllTables"`
	Members             FirewallAddrgrpMemberArrayOutput        `pulumi:"members"`
	Name                pulumi.StringOutput                     `pulumi:"name"`
	Taggings            FirewallAddrgrpTaggingArrayOutput       `pulumi:"taggings"`
	Type                pulumi.StringOutput                     `pulumi:"type"`
	Uuid                pulumi.StringOutput                     `pulumi:"uuid"`
	Vdomparam           pulumi.StringPtrOutput                  `pulumi:"vdomparam"`
	Visibility          pulumi.StringOutput                     `pulumi:"visibility"`
}

// NewFirewallAddrgrp registers a new resource with the given unique name, arguments, and options.
func NewFirewallAddrgrp(ctx *pulumi.Context,
	name string, args *FirewallAddrgrpArgs, opts ...pulumi.ResourceOption) (*FirewallAddrgrp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FirewallAddrgrp
	err := ctx.RegisterResource("fortios:index/firewallAddrgrp:FirewallAddrgrp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallAddrgrp gets an existing FirewallAddrgrp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallAddrgrp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallAddrgrpState, opts ...pulumi.ResourceOption) (*FirewallAddrgrp, error) {
	var resource FirewallAddrgrp
	err := ctx.ReadResource("fortios:index/firewallAddrgrp:FirewallAddrgrp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallAddrgrp resources.
type firewallAddrgrpState struct {
	AllowRouting        *string                        `pulumi:"allowRouting"`
	Category            *string                        `pulumi:"category"`
	Color               *int                           `pulumi:"color"`
	Comment             *string                        `pulumi:"comment"`
	DynamicSortSubtable *string                        `pulumi:"dynamicSortSubtable"`
	Exclude             *string                        `pulumi:"exclude"`
	ExcludeMembers      []FirewallAddrgrpExcludeMember `pulumi:"excludeMembers"`
	FabricObject        *string                        `pulumi:"fabricObject"`
	GetAllTables        *string                        `pulumi:"getAllTables"`
	Members             []FirewallAddrgrpMember        `pulumi:"members"`
	Name                *string                        `pulumi:"name"`
	Taggings            []FirewallAddrgrpTagging       `pulumi:"taggings"`
	Type                *string                        `pulumi:"type"`
	Uuid                *string                        `pulumi:"uuid"`
	Vdomparam           *string                        `pulumi:"vdomparam"`
	Visibility          *string                        `pulumi:"visibility"`
}

type FirewallAddrgrpState struct {
	AllowRouting        pulumi.StringPtrInput
	Category            pulumi.StringPtrInput
	Color               pulumi.IntPtrInput
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Exclude             pulumi.StringPtrInput
	ExcludeMembers      FirewallAddrgrpExcludeMemberArrayInput
	FabricObject        pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Members             FirewallAddrgrpMemberArrayInput
	Name                pulumi.StringPtrInput
	Taggings            FirewallAddrgrpTaggingArrayInput
	Type                pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Visibility          pulumi.StringPtrInput
}

func (FirewallAddrgrpState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAddrgrpState)(nil)).Elem()
}

type firewallAddrgrpArgs struct {
	AllowRouting        *string                        `pulumi:"allowRouting"`
	Category            *string                        `pulumi:"category"`
	Color               *int                           `pulumi:"color"`
	Comment             *string                        `pulumi:"comment"`
	DynamicSortSubtable *string                        `pulumi:"dynamicSortSubtable"`
	Exclude             *string                        `pulumi:"exclude"`
	ExcludeMembers      []FirewallAddrgrpExcludeMember `pulumi:"excludeMembers"`
	FabricObject        *string                        `pulumi:"fabricObject"`
	GetAllTables        *string                        `pulumi:"getAllTables"`
	Members             []FirewallAddrgrpMember        `pulumi:"members"`
	Name                *string                        `pulumi:"name"`
	Taggings            []FirewallAddrgrpTagging       `pulumi:"taggings"`
	Type                *string                        `pulumi:"type"`
	Uuid                *string                        `pulumi:"uuid"`
	Vdomparam           *string                        `pulumi:"vdomparam"`
	Visibility          *string                        `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallAddrgrp resource.
type FirewallAddrgrpArgs struct {
	AllowRouting        pulumi.StringPtrInput
	Category            pulumi.StringPtrInput
	Color               pulumi.IntPtrInput
	Comment             pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	Exclude             pulumi.StringPtrInput
	ExcludeMembers      FirewallAddrgrpExcludeMemberArrayInput
	FabricObject        pulumi.StringPtrInput
	GetAllTables        pulumi.StringPtrInput
	Members             FirewallAddrgrpMemberArrayInput
	Name                pulumi.StringPtrInput
	Taggings            FirewallAddrgrpTaggingArrayInput
	Type                pulumi.StringPtrInput
	Uuid                pulumi.StringPtrInput
	Vdomparam           pulumi.StringPtrInput
	Visibility          pulumi.StringPtrInput
}

func (FirewallAddrgrpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallAddrgrpArgs)(nil)).Elem()
}

type FirewallAddrgrpInput interface {
	pulumi.Input

	ToFirewallAddrgrpOutput() FirewallAddrgrpOutput
	ToFirewallAddrgrpOutputWithContext(ctx context.Context) FirewallAddrgrpOutput
}

func (*FirewallAddrgrp) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAddrgrp)(nil)).Elem()
}

func (i *FirewallAddrgrp) ToFirewallAddrgrpOutput() FirewallAddrgrpOutput {
	return i.ToFirewallAddrgrpOutputWithContext(context.Background())
}

func (i *FirewallAddrgrp) ToFirewallAddrgrpOutputWithContext(ctx context.Context) FirewallAddrgrpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddrgrpOutput)
}

// FirewallAddrgrpArrayInput is an input type that accepts FirewallAddrgrpArray and FirewallAddrgrpArrayOutput values.
// You can construct a concrete instance of `FirewallAddrgrpArrayInput` via:
//
//	FirewallAddrgrpArray{ FirewallAddrgrpArgs{...} }
type FirewallAddrgrpArrayInput interface {
	pulumi.Input

	ToFirewallAddrgrpArrayOutput() FirewallAddrgrpArrayOutput
	ToFirewallAddrgrpArrayOutputWithContext(context.Context) FirewallAddrgrpArrayOutput
}

type FirewallAddrgrpArray []FirewallAddrgrpInput

func (FirewallAddrgrpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAddrgrp)(nil)).Elem()
}

func (i FirewallAddrgrpArray) ToFirewallAddrgrpArrayOutput() FirewallAddrgrpArrayOutput {
	return i.ToFirewallAddrgrpArrayOutputWithContext(context.Background())
}

func (i FirewallAddrgrpArray) ToFirewallAddrgrpArrayOutputWithContext(ctx context.Context) FirewallAddrgrpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddrgrpArrayOutput)
}

// FirewallAddrgrpMapInput is an input type that accepts FirewallAddrgrpMap and FirewallAddrgrpMapOutput values.
// You can construct a concrete instance of `FirewallAddrgrpMapInput` via:
//
//	FirewallAddrgrpMap{ "key": FirewallAddrgrpArgs{...} }
type FirewallAddrgrpMapInput interface {
	pulumi.Input

	ToFirewallAddrgrpMapOutput() FirewallAddrgrpMapOutput
	ToFirewallAddrgrpMapOutputWithContext(context.Context) FirewallAddrgrpMapOutput
}

type FirewallAddrgrpMap map[string]FirewallAddrgrpInput

func (FirewallAddrgrpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAddrgrp)(nil)).Elem()
}

func (i FirewallAddrgrpMap) ToFirewallAddrgrpMapOutput() FirewallAddrgrpMapOutput {
	return i.ToFirewallAddrgrpMapOutputWithContext(context.Background())
}

func (i FirewallAddrgrpMap) ToFirewallAddrgrpMapOutputWithContext(ctx context.Context) FirewallAddrgrpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallAddrgrpMapOutput)
}

type FirewallAddrgrpOutput struct{ *pulumi.OutputState }

func (FirewallAddrgrpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallAddrgrp)(nil)).Elem()
}

func (o FirewallAddrgrpOutput) ToFirewallAddrgrpOutput() FirewallAddrgrpOutput {
	return o
}

func (o FirewallAddrgrpOutput) ToFirewallAddrgrpOutputWithContext(ctx context.Context) FirewallAddrgrpOutput {
	return o
}

func (o FirewallAddrgrpOutput) AllowRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringOutput { return v.AllowRouting }).(pulumi.StringOutput)
}

func (o FirewallAddrgrpOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

func (o FirewallAddrgrpOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

func (o FirewallAddrgrpOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallAddrgrpOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallAddrgrpOutput) Exclude() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringOutput { return v.Exclude }).(pulumi.StringOutput)
}

func (o FirewallAddrgrpOutput) ExcludeMembers() FirewallAddrgrpExcludeMemberArrayOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) FirewallAddrgrpExcludeMemberArrayOutput { return v.ExcludeMembers }).(FirewallAddrgrpExcludeMemberArrayOutput)
}

func (o FirewallAddrgrpOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

func (o FirewallAddrgrpOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o FirewallAddrgrpOutput) Members() FirewallAddrgrpMemberArrayOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) FirewallAddrgrpMemberArrayOutput { return v.Members }).(FirewallAddrgrpMemberArrayOutput)
}

func (o FirewallAddrgrpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallAddrgrpOutput) Taggings() FirewallAddrgrpTaggingArrayOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) FirewallAddrgrpTaggingArrayOutput { return v.Taggings }).(FirewallAddrgrpTaggingArrayOutput)
}

func (o FirewallAddrgrpOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FirewallAddrgrpOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o FirewallAddrgrpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o FirewallAddrgrpOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallAddrgrp) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type FirewallAddrgrpArrayOutput struct{ *pulumi.OutputState }

func (FirewallAddrgrpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallAddrgrp)(nil)).Elem()
}

func (o FirewallAddrgrpArrayOutput) ToFirewallAddrgrpArrayOutput() FirewallAddrgrpArrayOutput {
	return o
}

func (o FirewallAddrgrpArrayOutput) ToFirewallAddrgrpArrayOutputWithContext(ctx context.Context) FirewallAddrgrpArrayOutput {
	return o
}

func (o FirewallAddrgrpArrayOutput) Index(i pulumi.IntInput) FirewallAddrgrpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallAddrgrp {
		return vs[0].([]*FirewallAddrgrp)[vs[1].(int)]
	}).(FirewallAddrgrpOutput)
}

type FirewallAddrgrpMapOutput struct{ *pulumi.OutputState }

func (FirewallAddrgrpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallAddrgrp)(nil)).Elem()
}

func (o FirewallAddrgrpMapOutput) ToFirewallAddrgrpMapOutput() FirewallAddrgrpMapOutput {
	return o
}

func (o FirewallAddrgrpMapOutput) ToFirewallAddrgrpMapOutputWithContext(ctx context.Context) FirewallAddrgrpMapOutput {
	return o
}

func (o FirewallAddrgrpMapOutput) MapIndex(k pulumi.StringInput) FirewallAddrgrpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallAddrgrp {
		return vs[0].(map[string]*FirewallAddrgrp)[vs[1].(string)]
	}).(FirewallAddrgrpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddrgrpInput)(nil)).Elem(), &FirewallAddrgrp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddrgrpArrayInput)(nil)).Elem(), FirewallAddrgrpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallAddrgrpMapInput)(nil)).Elem(), FirewallAddrgrpMap{})
	pulumi.RegisterOutputType(FirewallAddrgrpOutput{})
	pulumi.RegisterOutputType(FirewallAddrgrpArrayOutput{})
	pulumi.RegisterOutputType(FirewallAddrgrpMapOutput{})
}
