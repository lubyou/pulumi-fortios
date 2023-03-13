// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFirewallWildcardFqdnCustomList(ctx *pulumi.Context, args *GetFirewallWildcardFqdnCustomListArgs, opts ...pulumi.InvokeOption) (*GetFirewallWildcardFqdnCustomListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallWildcardFqdnCustomListResult
	err := ctx.Invoke("fortios:index/getFirewallWildcardFqdnCustomList:GetFirewallWildcardFqdnCustomList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallWildcardFqdnCustomList.
type GetFirewallWildcardFqdnCustomListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallWildcardFqdnCustomList.
type GetFirewallWildcardFqdnCustomListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallWildcardFqdnCustomListOutput(ctx *pulumi.Context, args GetFirewallWildcardFqdnCustomListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallWildcardFqdnCustomListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallWildcardFqdnCustomListResult, error) {
			args := v.(GetFirewallWildcardFqdnCustomListArgs)
			r, err := GetFirewallWildcardFqdnCustomList(ctx, &args, opts...)
			var s GetFirewallWildcardFqdnCustomListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallWildcardFqdnCustomListResultOutput)
}

// A collection of arguments for invoking GetFirewallWildcardFqdnCustomList.
type GetFirewallWildcardFqdnCustomListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallWildcardFqdnCustomListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallWildcardFqdnCustomListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallWildcardFqdnCustomList.
type GetFirewallWildcardFqdnCustomListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallWildcardFqdnCustomListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallWildcardFqdnCustomListResult)(nil)).Elem()
}

func (o GetFirewallWildcardFqdnCustomListResultOutput) ToGetFirewallWildcardFqdnCustomListResultOutput() GetFirewallWildcardFqdnCustomListResultOutput {
	return o
}

func (o GetFirewallWildcardFqdnCustomListResultOutput) ToGetFirewallWildcardFqdnCustomListResultOutputWithContext(ctx context.Context) GetFirewallWildcardFqdnCustomListResultOutput {
	return o
}

func (o GetFirewallWildcardFqdnCustomListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallWildcardFqdnCustomListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallWildcardFqdnCustomListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallWildcardFqdnCustomListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFirewallWildcardFqdnCustomListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallWildcardFqdnCustomListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallWildcardFqdnCustomListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallWildcardFqdnCustomListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallWildcardFqdnCustomListResultOutput{})
}
