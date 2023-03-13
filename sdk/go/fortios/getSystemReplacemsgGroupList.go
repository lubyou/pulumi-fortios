// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSystemReplacemsgGroupList(ctx *pulumi.Context, args *GetSystemReplacemsgGroupListArgs, opts ...pulumi.InvokeOption) (*GetSystemReplacemsgGroupListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemReplacemsgGroupListResult
	err := ctx.Invoke("fortios:index/getSystemReplacemsgGroupList:GetSystemReplacemsgGroupList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemReplacemsgGroupList.
type GetSystemReplacemsgGroupListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemReplacemsgGroupList.
type GetSystemReplacemsgGroupListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemReplacemsgGroupListOutput(ctx *pulumi.Context, args GetSystemReplacemsgGroupListOutputArgs, opts ...pulumi.InvokeOption) GetSystemReplacemsgGroupListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemReplacemsgGroupListResult, error) {
			args := v.(GetSystemReplacemsgGroupListArgs)
			r, err := GetSystemReplacemsgGroupList(ctx, &args, opts...)
			var s GetSystemReplacemsgGroupListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemReplacemsgGroupListResultOutput)
}

// A collection of arguments for invoking GetSystemReplacemsgGroupList.
type GetSystemReplacemsgGroupListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemReplacemsgGroupListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemReplacemsgGroupListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemReplacemsgGroupList.
type GetSystemReplacemsgGroupListResultOutput struct{ *pulumi.OutputState }

func (GetSystemReplacemsgGroupListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemReplacemsgGroupListResult)(nil)).Elem()
}

func (o GetSystemReplacemsgGroupListResultOutput) ToGetSystemReplacemsgGroupListResultOutput() GetSystemReplacemsgGroupListResultOutput {
	return o
}

func (o GetSystemReplacemsgGroupListResultOutput) ToGetSystemReplacemsgGroupListResultOutputWithContext(ctx context.Context) GetSystemReplacemsgGroupListResultOutput {
	return o
}

func (o GetSystemReplacemsgGroupListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemReplacemsgGroupListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemReplacemsgGroupListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemReplacemsgGroupListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemReplacemsgGroupListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemReplacemsgGroupListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemReplacemsgGroupListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemReplacemsgGroupListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemReplacemsgGroupListResultOutput{})
}
