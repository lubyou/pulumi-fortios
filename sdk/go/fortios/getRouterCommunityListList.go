// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRouterCommunityListList(ctx *pulumi.Context, args *GetRouterCommunityListListArgs, opts ...pulumi.InvokeOption) (*GetRouterCommunityListListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRouterCommunityListListResult
	err := ctx.Invoke("fortios:index/getRouterCommunityListList:GetRouterCommunityListList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterCommunityListList.
type GetRouterCommunityListListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterCommunityListList.
type GetRouterCommunityListListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetRouterCommunityListListOutput(ctx *pulumi.Context, args GetRouterCommunityListListOutputArgs, opts ...pulumi.InvokeOption) GetRouterCommunityListListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterCommunityListListResult, error) {
			args := v.(GetRouterCommunityListListArgs)
			r, err := GetRouterCommunityListList(ctx, &args, opts...)
			var s GetRouterCommunityListListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterCommunityListListResultOutput)
}

// A collection of arguments for invoking GetRouterCommunityListList.
type GetRouterCommunityListListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterCommunityListListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterCommunityListListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterCommunityListList.
type GetRouterCommunityListListResultOutput struct{ *pulumi.OutputState }

func (GetRouterCommunityListListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterCommunityListListResult)(nil)).Elem()
}

func (o GetRouterCommunityListListResultOutput) ToGetRouterCommunityListListResultOutput() GetRouterCommunityListListResultOutput {
	return o
}

func (o GetRouterCommunityListListResultOutput) ToGetRouterCommunityListListResultOutputWithContext(ctx context.Context) GetRouterCommunityListListResultOutput {
	return o
}

func (o GetRouterCommunityListListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterCommunityListListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterCommunityListListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterCommunityListListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRouterCommunityListListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouterCommunityListListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetRouterCommunityListListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterCommunityListListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterCommunityListListResultOutput{})
}
