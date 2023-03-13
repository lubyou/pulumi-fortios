// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallInternetServiceGroup(ctx *pulumi.Context, args *LookupFirewallInternetServiceGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetServiceGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallInternetServiceGroupResult
	err := ctx.Invoke("fortios:index/getFirewallInternetServiceGroup:GetFirewallInternetServiceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallInternetServiceGroup.
type LookupFirewallInternetServiceGroupArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallInternetServiceGroup.
type LookupFirewallInternetServiceGroupResult struct {
	Comment   string `pulumi:"comment"`
	Direction string `pulumi:"direction"`
	// The provider-assigned unique ID for this managed resource.
	Id        string                                  `pulumi:"id"`
	Members   []GetFirewallInternetServiceGroupMember `pulumi:"members"`
	Name      string                                  `pulumi:"name"`
	Vdomparam *string                                 `pulumi:"vdomparam"`
}

func LookupFirewallInternetServiceGroupOutput(ctx *pulumi.Context, args LookupFirewallInternetServiceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallInternetServiceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallInternetServiceGroupResult, error) {
			args := v.(LookupFirewallInternetServiceGroupArgs)
			r, err := LookupFirewallInternetServiceGroup(ctx, &args, opts...)
			var s LookupFirewallInternetServiceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallInternetServiceGroupResultOutput)
}

// A collection of arguments for invoking GetFirewallInternetServiceGroup.
type LookupFirewallInternetServiceGroupOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallInternetServiceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetServiceGroupArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallInternetServiceGroup.
type LookupFirewallInternetServiceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallInternetServiceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetServiceGroupResult)(nil)).Elem()
}

func (o LookupFirewallInternetServiceGroupResultOutput) ToLookupFirewallInternetServiceGroupResultOutput() LookupFirewallInternetServiceGroupResultOutput {
	return o
}

func (o LookupFirewallInternetServiceGroupResultOutput) ToLookupFirewallInternetServiceGroupResultOutputWithContext(ctx context.Context) LookupFirewallInternetServiceGroupResultOutput {
	return o
}

func (o LookupFirewallInternetServiceGroupResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceGroupResult) string { return v.Comment }).(pulumi.StringOutput)
}

func (o LookupFirewallInternetServiceGroupResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceGroupResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallInternetServiceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallInternetServiceGroupResultOutput) Members() GetFirewallInternetServiceGroupMemberArrayOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceGroupResult) []GetFirewallInternetServiceGroupMember {
		return v.Members
	}).(GetFirewallInternetServiceGroupMemberArrayOutput)
}

func (o LookupFirewallInternetServiceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallInternetServiceGroupResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceGroupResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallInternetServiceGroupResultOutput{})
}
