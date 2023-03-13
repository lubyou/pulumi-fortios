// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRouterAccessListList(ctx *pulumi.Context, args *GetRouterAccessListListArgs, opts ...pulumi.InvokeOption) (*GetRouterAccessListListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRouterAccessListListResult
	err := ctx.Invoke("fortios:index/getRouterAccessListList:GetRouterAccessListList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterAccessListList.
type GetRouterAccessListListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterAccessListList.
type GetRouterAccessListListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetRouterAccessListListOutput(ctx *pulumi.Context, args GetRouterAccessListListOutputArgs, opts ...pulumi.InvokeOption) GetRouterAccessListListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterAccessListListResult, error) {
			args := v.(GetRouterAccessListListArgs)
			r, err := GetRouterAccessListList(ctx, &args, opts...)
			var s GetRouterAccessListListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterAccessListListResultOutput)
}

// A collection of arguments for invoking GetRouterAccessListList.
type GetRouterAccessListListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterAccessListListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterAccessListListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterAccessListList.
type GetRouterAccessListListResultOutput struct{ *pulumi.OutputState }

func (GetRouterAccessListListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterAccessListListResult)(nil)).Elem()
}

func (o GetRouterAccessListListResultOutput) ToGetRouterAccessListListResultOutput() GetRouterAccessListListResultOutput {
	return o
}

func (o GetRouterAccessListListResultOutput) ToGetRouterAccessListListResultOutputWithContext(ctx context.Context) GetRouterAccessListListResultOutput {
	return o
}

func (o GetRouterAccessListListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterAccessListListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterAccessListListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterAccessListListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRouterAccessListListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouterAccessListListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetRouterAccessListListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterAccessListListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterAccessListListResultOutput{})
}
