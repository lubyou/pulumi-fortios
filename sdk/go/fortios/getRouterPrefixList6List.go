// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `RouterPrefixList6`.
func GetRouterPrefixList6List(ctx *pulumi.Context, args *GetRouterPrefixList6ListArgs, opts ...pulumi.InvokeOption) (*GetRouterPrefixList6ListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRouterPrefixList6ListResult
	err := ctx.Invoke("fortios:index/getRouterPrefixList6List:GetRouterPrefixList6List", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRouterPrefixList6List.
type GetRouterPrefixList6ListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetRouterPrefixList6List.
type GetRouterPrefixList6ListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `RouterPrefixList6`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetRouterPrefixList6ListOutput(ctx *pulumi.Context, args GetRouterPrefixList6ListOutputArgs, opts ...pulumi.InvokeOption) GetRouterPrefixList6ListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouterPrefixList6ListResult, error) {
			args := v.(GetRouterPrefixList6ListArgs)
			r, err := GetRouterPrefixList6List(ctx, &args, opts...)
			var s GetRouterPrefixList6ListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouterPrefixList6ListResultOutput)
}

// A collection of arguments for invoking GetRouterPrefixList6List.
type GetRouterPrefixList6ListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetRouterPrefixList6ListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterPrefixList6ListArgs)(nil)).Elem()
}

// A collection of values returned by GetRouterPrefixList6List.
type GetRouterPrefixList6ListResultOutput struct{ *pulumi.OutputState }

func (GetRouterPrefixList6ListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouterPrefixList6ListResult)(nil)).Elem()
}

func (o GetRouterPrefixList6ListResultOutput) ToGetRouterPrefixList6ListResultOutput() GetRouterPrefixList6ListResultOutput {
	return o
}

func (o GetRouterPrefixList6ListResultOutput) ToGetRouterPrefixList6ListResultOutputWithContext(ctx context.Context) GetRouterPrefixList6ListResultOutput {
	return o
}

func (o GetRouterPrefixList6ListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterPrefixList6ListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouterPrefixList6ListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouterPrefixList6ListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `RouterPrefixList6`.
func (o GetRouterPrefixList6ListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouterPrefixList6ListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetRouterPrefixList6ListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouterPrefixList6ListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouterPrefixList6ListResultOutput{})
}
