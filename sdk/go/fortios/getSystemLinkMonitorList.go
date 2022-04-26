// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemLinkMonitor`.
func GetSystemLinkMonitorList(ctx *pulumi.Context, args *GetSystemLinkMonitorListArgs, opts ...pulumi.InvokeOption) (*GetSystemLinkMonitorListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemLinkMonitorListResult
	err := ctx.Invoke("fortios:index/getSystemLinkMonitorList:GetSystemLinkMonitorList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemLinkMonitorList.
type GetSystemLinkMonitorListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemLinkMonitorList.
type GetSystemLinkMonitorListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemLinkMonitor`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemLinkMonitorListOutput(ctx *pulumi.Context, args GetSystemLinkMonitorListOutputArgs, opts ...pulumi.InvokeOption) GetSystemLinkMonitorListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemLinkMonitorListResult, error) {
			args := v.(GetSystemLinkMonitorListArgs)
			r, err := GetSystemLinkMonitorList(ctx, &args, opts...)
			var s GetSystemLinkMonitorListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemLinkMonitorListResultOutput)
}

// A collection of arguments for invoking GetSystemLinkMonitorList.
type GetSystemLinkMonitorListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemLinkMonitorListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemLinkMonitorListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemLinkMonitorList.
type GetSystemLinkMonitorListResultOutput struct{ *pulumi.OutputState }

func (GetSystemLinkMonitorListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemLinkMonitorListResult)(nil)).Elem()
}

func (o GetSystemLinkMonitorListResultOutput) ToGetSystemLinkMonitorListResultOutput() GetSystemLinkMonitorListResultOutput {
	return o
}

func (o GetSystemLinkMonitorListResultOutput) ToGetSystemLinkMonitorListResultOutputWithContext(ctx context.Context) GetSystemLinkMonitorListResultOutput {
	return o
}

func (o GetSystemLinkMonitorListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemLinkMonitorListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemLinkMonitorListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemLinkMonitorListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemLinkMonitor`.
func (o GetSystemLinkMonitorListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemLinkMonitorListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemLinkMonitorListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemLinkMonitorListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemLinkMonitorListResultOutput{})
}
