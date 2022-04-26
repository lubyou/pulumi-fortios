// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `RouterPolicy6`.
func GetRouterPolicy6List(ctx *pulumi.Context, args *GetRouterPolicy6ListArgs, opts ...pulumi.InvokeOption) (*GetRouterPolicy6ListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRouterPolicy6ListResult
	err := ctx.Invoke("fortios:index/getRouterPolicy6List:GetRouterPolicy6List", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterPolicy6List.
type GetRouterPolicy6ListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterPolicy6List.
type GetRouterPolicy6ListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `RouterPolicy6`.
	SeqNumlists []int   `pulumi:"seqNumlists"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func GetRouterPolicy6ListOutput(ctx *pulumi.Context, args GetRouterPolicy6ListOutputArgs, opts ...pulumi.InvokeOption) GetRouterPolicy6ListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterPolicy6ListResult, error) {
			args := v.(GetRouterPolicy6ListArgs)
			r, err := GetRouterPolicy6List(ctx, &args, opts...)
			var s GetRouterPolicy6ListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterPolicy6ListResultOutput)
}

// A collection of arguments for invoking GetRouterPolicy6List.
type GetRouterPolicy6ListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterPolicy6ListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterPolicy6ListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterPolicy6List.
type GetRouterPolicy6ListResultOutput struct{ *pulumi.OutputState }

func (GetRouterPolicy6ListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterPolicy6ListResult)(nil)).Elem()
}

func (o GetRouterPolicy6ListResultOutput) ToGetRouterPolicy6ListResultOutput() GetRouterPolicy6ListResultOutput {
	return o
}

func (o GetRouterPolicy6ListResultOutput) ToGetRouterPolicy6ListResultOutputWithContext(ctx context.Context) GetRouterPolicy6ListResultOutput {
	return o
}

func (o GetRouterPolicy6ListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterPolicy6ListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterPolicy6ListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterPolicy6ListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `RouterPolicy6`.
func (o GetRouterPolicy6ListResultOutput) SeqNumlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetRouterPolicy6ListResult) []int { return v.SeqNumlists }).(pulumi.IntArrayOutput)
}

func (o GetRouterPolicy6ListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterPolicy6ListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterPolicy6ListResultOutput{})
}
