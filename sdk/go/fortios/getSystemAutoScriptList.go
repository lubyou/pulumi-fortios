// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemAutoScript`.
func GetSystemAutoScriptList(ctx *pulumi.Context, args *GetSystemAutoScriptListArgs, opts ...pulumi.InvokeOption) (*GetSystemAutoScriptListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemAutoScriptListResult
	err := ctx.Invoke("fortios:index/getSystemAutoScriptList:GetSystemAutoScriptList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutoScriptList.
type GetSystemAutoScriptListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutoScriptList.
type GetSystemAutoScriptListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemAutoScript`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemAutoScriptListOutput(ctx *pulumi.Context, args GetSystemAutoScriptListOutputArgs, opts ...pulumi.InvokeOption) GetSystemAutoScriptListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemAutoScriptListResult, error) {
			args := v.(GetSystemAutoScriptListArgs)
			r, err := GetSystemAutoScriptList(ctx, &args, opts...)
			var s GetSystemAutoScriptListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemAutoScriptListResultOutput)
}

// A collection of arguments for invoking GetSystemAutoScriptList.
type GetSystemAutoScriptListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemAutoScriptListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutoScriptListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemAutoScriptList.
type GetSystemAutoScriptListResultOutput struct{ *pulumi.OutputState }

func (GetSystemAutoScriptListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutoScriptListResult)(nil)).Elem()
}

func (o GetSystemAutoScriptListResultOutput) ToGetSystemAutoScriptListResultOutput() GetSystemAutoScriptListResultOutput {
	return o
}

func (o GetSystemAutoScriptListResultOutput) ToGetSystemAutoScriptListResultOutputWithContext(ctx context.Context) GetSystemAutoScriptListResultOutput {
	return o
}

func (o GetSystemAutoScriptListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutoScriptListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemAutoScriptListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemAutoScriptListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemAutoScript`.
func (o GetSystemAutoScriptListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemAutoScriptListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemAutoScriptListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutoScriptListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemAutoScriptListResultOutput{})
}
