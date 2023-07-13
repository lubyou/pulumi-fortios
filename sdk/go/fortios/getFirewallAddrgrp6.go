// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallAddrgrp6(ctx *pulumi.Context, args *LookupFirewallAddrgrp6Args, opts ...pulumi.InvokeOption) (*LookupFirewallAddrgrp6Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallAddrgrp6Result
	err := ctx.Invoke("fortios:index/getFirewallAddrgrp6:GetFirewallAddrgrp6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallAddrgrp6.
type LookupFirewallAddrgrp6Args struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallAddrgrp6.
type LookupFirewallAddrgrp6Result struct {
	Color          int                                `pulumi:"color"`
	Comment        string                             `pulumi:"comment"`
	Exclude        string                             `pulumi:"exclude"`
	ExcludeMembers []GetFirewallAddrgrp6ExcludeMember `pulumi:"excludeMembers"`
	FabricObject   string                             `pulumi:"fabricObject"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                       `pulumi:"id"`
	Members    []GetFirewallAddrgrp6Member  `pulumi:"members"`
	Name       string                       `pulumi:"name"`
	Taggings   []GetFirewallAddrgrp6Tagging `pulumi:"taggings"`
	Uuid       string                       `pulumi:"uuid"`
	Vdomparam  *string                      `pulumi:"vdomparam"`
	Visibility string                       `pulumi:"visibility"`
}

func LookupFirewallAddrgrp6Output(ctx *pulumi.Context, args LookupFirewallAddrgrp6OutputArgs, opts ...pulumi.InvokeOption) LookupFirewallAddrgrp6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallAddrgrp6Result, error) {
			args := v.(LookupFirewallAddrgrp6Args)
			r, err := LookupFirewallAddrgrp6(ctx, &args, opts...)
			var s LookupFirewallAddrgrp6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallAddrgrp6ResultOutput)
}

// A collection of arguments for invoking GetFirewallAddrgrp6.
type LookupFirewallAddrgrp6OutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallAddrgrp6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallAddrgrp6Args)(nil)).Elem()
}

// A collection of values returned by GetFirewallAddrgrp6.
type LookupFirewallAddrgrp6ResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallAddrgrp6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallAddrgrp6Result)(nil)).Elem()
}

func (o LookupFirewallAddrgrp6ResultOutput) ToLookupFirewallAddrgrp6ResultOutput() LookupFirewallAddrgrp6ResultOutput {
	return o
}

func (o LookupFirewallAddrgrp6ResultOutput) ToLookupFirewallAddrgrp6ResultOutputWithContext(ctx context.Context) LookupFirewallAddrgrp6ResultOutput {
	return o
}

func (o LookupFirewallAddrgrp6ResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) int { return v.Color }).(pulumi.IntOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) string { return v.Comment }).(pulumi.StringOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) Exclude() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) string { return v.Exclude }).(pulumi.StringOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) ExcludeMembers() GetFirewallAddrgrp6ExcludeMemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) []GetFirewallAddrgrp6ExcludeMember { return v.ExcludeMembers }).(GetFirewallAddrgrp6ExcludeMemberArrayOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) string { return v.FabricObject }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallAddrgrp6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) Members() GetFirewallAddrgrp6MemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) []GetFirewallAddrgrp6Member { return v.Members }).(GetFirewallAddrgrp6MemberArrayOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) Taggings() GetFirewallAddrgrp6TaggingArrayOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) []GetFirewallAddrgrp6Tagging { return v.Taggings }).(GetFirewallAddrgrp6TaggingArrayOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallAddrgrp6ResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallAddrgrp6Result) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallAddrgrp6ResultOutput{})
}
