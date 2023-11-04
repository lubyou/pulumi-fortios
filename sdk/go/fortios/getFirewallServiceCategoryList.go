// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func GetFirewallServiceCategoryList(ctx *pulumi.Context, args *GetFirewallServiceCategoryListArgs, opts ...pulumi.InvokeOption) (*GetFirewallServiceCategoryListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFirewallServiceCategoryListResult
	err := ctx.Invoke("fortios:index/getFirewallServiceCategoryList:GetFirewallServiceCategoryList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallServiceCategoryList.
type GetFirewallServiceCategoryListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallServiceCategoryList.
type GetFirewallServiceCategoryListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallServiceCategoryListOutput(ctx *pulumi.Context, args GetFirewallServiceCategoryListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallServiceCategoryListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallServiceCategoryListResult, error) {
			args := v.(GetFirewallServiceCategoryListArgs)
			r, err := GetFirewallServiceCategoryList(ctx, &args, opts...)
			var s GetFirewallServiceCategoryListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallServiceCategoryListResultOutput)
}

// A collection of arguments for invoking GetFirewallServiceCategoryList.
type GetFirewallServiceCategoryListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallServiceCategoryListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallServiceCategoryListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallServiceCategoryList.
type GetFirewallServiceCategoryListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallServiceCategoryListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallServiceCategoryListResult)(nil)).Elem()
}

func (o GetFirewallServiceCategoryListResultOutput) ToGetFirewallServiceCategoryListResultOutput() GetFirewallServiceCategoryListResultOutput {
	return o
}

func (o GetFirewallServiceCategoryListResultOutput) ToGetFirewallServiceCategoryListResultOutputWithContext(ctx context.Context) GetFirewallServiceCategoryListResultOutput {
	return o
}

func (o GetFirewallServiceCategoryListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetFirewallServiceCategoryListResult] {
	return pulumix.Output[GetFirewallServiceCategoryListResult]{
		OutputState: o.OutputState,
	}
}

func (o GetFirewallServiceCategoryListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallServiceCategoryListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallServiceCategoryListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallServiceCategoryListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFirewallServiceCategoryListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallServiceCategoryListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallServiceCategoryListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallServiceCategoryListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallServiceCategoryListResultOutput{})
}
