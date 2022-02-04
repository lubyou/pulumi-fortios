// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemObjectTagging`.
func GetSystemObjectTaggingList(ctx *pulumi.Context, args *GetSystemObjectTaggingListArgs, opts ...pulumi.InvokeOption) (*GetSystemObjectTaggingListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemObjectTaggingListResult
	err := ctx.Invoke("fortios:index/getSystemObjectTaggingList:GetSystemObjectTaggingList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemObjectTaggingList.
type GetSystemObjectTaggingListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemObjectTaggingList.
type GetSystemObjectTaggingListResult struct {
	// A list of the `SystemObjectTagging`.
	Categorylists []string `pulumi:"categorylists"`
	Filter        *string  `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func GetSystemObjectTaggingListOutput(ctx *pulumi.Context, args GetSystemObjectTaggingListOutputArgs, opts ...pulumi.InvokeOption) GetSystemObjectTaggingListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemObjectTaggingListResult, error) {
			args := v.(GetSystemObjectTaggingListArgs)
			r, err := GetSystemObjectTaggingList(ctx, &args, opts...)
			return *r, err
		}).(GetSystemObjectTaggingListResultOutput)
}

// A collection of arguments for invoking GetSystemObjectTaggingList.
type GetSystemObjectTaggingListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemObjectTaggingListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemObjectTaggingListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemObjectTaggingList.
type GetSystemObjectTaggingListResultOutput struct{ *pulumi.OutputState }

func (GetSystemObjectTaggingListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemObjectTaggingListResult)(nil)).Elem()
}

func (o GetSystemObjectTaggingListResultOutput) ToGetSystemObjectTaggingListResultOutput() GetSystemObjectTaggingListResultOutput {
	return o
}

func (o GetSystemObjectTaggingListResultOutput) ToGetSystemObjectTaggingListResultOutputWithContext(ctx context.Context) GetSystemObjectTaggingListResultOutput {
	return o
}

// A list of the `SystemObjectTagging`.
func (o GetSystemObjectTaggingListResultOutput) Categorylists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemObjectTaggingListResult) []string { return v.Categorylists }).(pulumi.StringArrayOutput)
}

func (o GetSystemObjectTaggingListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemObjectTaggingListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemObjectTaggingListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemObjectTaggingListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSystemObjectTaggingListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemObjectTaggingListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemObjectTaggingListResultOutput{})
}
