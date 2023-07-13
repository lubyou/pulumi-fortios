// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRouterPrefixListList(ctx *pulumi.Context, args *GetRouterPrefixListListArgs, opts ...pulumi.InvokeOption) (*GetRouterPrefixListListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRouterPrefixListListResult
	err := ctx.Invoke("fortios:index/getRouterPrefixListList:GetRouterPrefixListList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterPrefixListList.
type GetRouterPrefixListListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterPrefixListList.
type GetRouterPrefixListListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetRouterPrefixListListOutput(ctx *pulumi.Context, args GetRouterPrefixListListOutputArgs, opts ...pulumi.InvokeOption) GetRouterPrefixListListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterPrefixListListResult, error) {
			args := v.(GetRouterPrefixListListArgs)
			r, err := GetRouterPrefixListList(ctx, &args, opts...)
			var s GetRouterPrefixListListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterPrefixListListResultOutput)
}

// A collection of arguments for invoking GetRouterPrefixListList.
type GetRouterPrefixListListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterPrefixListListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterPrefixListListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterPrefixListList.
type GetRouterPrefixListListResultOutput struct{ *pulumi.OutputState }

func (GetRouterPrefixListListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterPrefixListListResult)(nil)).Elem()
}

func (o GetRouterPrefixListListResultOutput) ToGetRouterPrefixListListResultOutput() GetRouterPrefixListListResultOutput {
	return o
}

func (o GetRouterPrefixListListResultOutput) ToGetRouterPrefixListListResultOutputWithContext(ctx context.Context) GetRouterPrefixListListResultOutput {
	return o
}

func (o GetRouterPrefixListListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterPrefixListListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterPrefixListListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterPrefixListListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRouterPrefixListListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouterPrefixListListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetRouterPrefixListListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterPrefixListListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterPrefixListListResultOutput{})
}
