// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallInternetServiceList(ctx *pulumi.Context, args *LookupFirewallInternetServiceListArgs, opts ...pulumi.InvokeOption) (*LookupFirewallInternetServiceListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallInternetServiceListResult
	err := ctx.Invoke("fortios:index/getFirewallInternetServiceList:GetFirewallInternetServiceList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallInternetServiceList.
type LookupFirewallInternetServiceListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallInternetServiceList.
type LookupFirewallInternetServiceListResult struct {
	Filter     *string `pulumi:"filter"`
	Fosidlists []int   `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupFirewallInternetServiceListOutput(ctx *pulumi.Context, args LookupFirewallInternetServiceListOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallInternetServiceListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallInternetServiceListResult, error) {
			args := v.(LookupFirewallInternetServiceListArgs)
			r, err := LookupFirewallInternetServiceList(ctx, &args, opts...)
			var s LookupFirewallInternetServiceListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallInternetServiceListResultOutput)
}

// A collection of arguments for invoking GetFirewallInternetServiceList.
type LookupFirewallInternetServiceListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallInternetServiceListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetServiceListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallInternetServiceList.
type LookupFirewallInternetServiceListResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallInternetServiceListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallInternetServiceListResult)(nil)).Elem()
}

func (o LookupFirewallInternetServiceListResultOutput) ToLookupFirewallInternetServiceListResultOutput() LookupFirewallInternetServiceListResultOutput {
	return o
}

func (o LookupFirewallInternetServiceListResultOutput) ToLookupFirewallInternetServiceListResultOutputWithContext(ctx context.Context) LookupFirewallInternetServiceListResultOutput {
	return o
}

func (o LookupFirewallInternetServiceListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallInternetServiceListResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceListResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallInternetServiceListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallInternetServiceListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallInternetServiceListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallInternetServiceListResultOutput{})
}
