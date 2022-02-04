// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemVxlan`.
func GetSystemVxlanList(ctx *pulumi.Context, args *GetSystemVxlanListArgs, opts ...pulumi.InvokeOption) (*GetSystemVxlanListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemVxlanListResult
	err := ctx.Invoke("fortios:index/getSystemVxlanList:GetSystemVxlanList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemVxlanList.
type GetSystemVxlanListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemVxlanList.
type GetSystemVxlanListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemVxlan`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemVxlanListOutput(ctx *pulumi.Context, args GetSystemVxlanListOutputArgs, opts ...pulumi.InvokeOption) GetSystemVxlanListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemVxlanListResult, error) {
			args := v.(GetSystemVxlanListArgs)
			r, err := GetSystemVxlanList(ctx, &args, opts...)
			return *r, err
		}).(GetSystemVxlanListResultOutput)
}

// A collection of arguments for invoking GetSystemVxlanList.
type GetSystemVxlanListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemVxlanListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemVxlanListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemVxlanList.
type GetSystemVxlanListResultOutput struct{ *pulumi.OutputState }

func (GetSystemVxlanListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemVxlanListResult)(nil)).Elem()
}

func (o GetSystemVxlanListResultOutput) ToGetSystemVxlanListResultOutput() GetSystemVxlanListResultOutput {
	return o
}

func (o GetSystemVxlanListResultOutput) ToGetSystemVxlanListResultOutputWithContext(ctx context.Context) GetSystemVxlanListResultOutput {
	return o
}

func (o GetSystemVxlanListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemVxlanListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemVxlanListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemVxlanListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemVxlan`.
func (o GetSystemVxlanListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemVxlanListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemVxlanListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemVxlanListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemVxlanListResultOutput{})
}
