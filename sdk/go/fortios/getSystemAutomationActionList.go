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

func GetSystemAutomationActionList(ctx *pulumi.Context, args *GetSystemAutomationActionListArgs, opts ...pulumi.InvokeOption) (*GetSystemAutomationActionListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSystemAutomationActionListResult
	err := ctx.Invoke("fortios:index/getSystemAutomationActionList:GetSystemAutomationActionList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutomationActionList.
type GetSystemAutomationActionListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutomationActionList.
type GetSystemAutomationActionListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemAutomationActionListOutput(ctx *pulumi.Context, args GetSystemAutomationActionListOutputArgs, opts ...pulumi.InvokeOption) GetSystemAutomationActionListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemAutomationActionListResult, error) {
			args := v.(GetSystemAutomationActionListArgs)
			r, err := GetSystemAutomationActionList(ctx, &args, opts...)
			var s GetSystemAutomationActionListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemAutomationActionListResultOutput)
}

// A collection of arguments for invoking GetSystemAutomationActionList.
type GetSystemAutomationActionListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemAutomationActionListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutomationActionListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemAutomationActionList.
type GetSystemAutomationActionListResultOutput struct{ *pulumi.OutputState }

func (GetSystemAutomationActionListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutomationActionListResult)(nil)).Elem()
}

func (o GetSystemAutomationActionListResultOutput) ToGetSystemAutomationActionListResultOutput() GetSystemAutomationActionListResultOutput {
	return o
}

func (o GetSystemAutomationActionListResultOutput) ToGetSystemAutomationActionListResultOutputWithContext(ctx context.Context) GetSystemAutomationActionListResultOutput {
	return o
}

func (o GetSystemAutomationActionListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSystemAutomationActionListResult] {
	return pulumix.Output[GetSystemAutomationActionListResult]{
		OutputState: o.OutputState,
	}
}

func (o GetSystemAutomationActionListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutomationActionListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemAutomationActionListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemAutomationActionListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemAutomationActionListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemAutomationActionListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemAutomationActionListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutomationActionListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemAutomationActionListResultOutput{})
}
