// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemArpTable`.
func GetSystemArpTableList(ctx *pulumi.Context, args *GetSystemArpTableListArgs, opts ...pulumi.InvokeOption) (*GetSystemArpTableListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemArpTableListResult
	err := ctx.Invoke("fortios:index/getSystemArpTableList:GetSystemArpTableList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemArpTableList.
type GetSystemArpTableListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemArpTableList.
type GetSystemArpTableListResult struct {
	Filter *string `pulumi:"filter"`
	// A list of the `SystemArpTable`.
	Fosidlists []int `pulumi:"fosidlists"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemArpTableListOutput(ctx *pulumi.Context, args GetSystemArpTableListOutputArgs, opts ...pulumi.InvokeOption) GetSystemArpTableListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemArpTableListResult, error) {
			args := v.(GetSystemArpTableListArgs)
			r, err := GetSystemArpTableList(ctx, &args, opts...)
			var s GetSystemArpTableListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemArpTableListResultOutput)
}

// A collection of arguments for invoking GetSystemArpTableList.
type GetSystemArpTableListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemArpTableListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemArpTableListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemArpTableList.
type GetSystemArpTableListResultOutput struct{ *pulumi.OutputState }

func (GetSystemArpTableListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemArpTableListResult)(nil)).Elem()
}

func (o GetSystemArpTableListResultOutput) ToGetSystemArpTableListResultOutput() GetSystemArpTableListResultOutput {
	return o
}

func (o GetSystemArpTableListResultOutput) ToGetSystemArpTableListResultOutputWithContext(ctx context.Context) GetSystemArpTableListResultOutput {
	return o
}

func (o GetSystemArpTableListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemArpTableListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// A list of the `SystemArpTable`.
func (o GetSystemArpTableListResultOutput) Fosidlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemArpTableListResult) []int { return v.Fosidlists }).(pulumi.IntArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemArpTableListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemArpTableListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemArpTableListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemArpTableListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemArpTableListResultOutput{})
}
