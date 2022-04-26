// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemAutomationDestination`.
func GetSystemAutomationDestinationList(ctx *pulumi.Context, args *GetSystemAutomationDestinationListArgs, opts ...pulumi.InvokeOption) (*GetSystemAutomationDestinationListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemAutomationDestinationListResult
	err := ctx.Invoke("fortios:index/getSystemAutomationDestinationList:GetSystemAutomationDestinationList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAutomationDestinationList.
type GetSystemAutomationDestinationListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAutomationDestinationList.
type GetSystemAutomationDestinationListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemAutomationDestination`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemAutomationDestinationListOutput(ctx *pulumi.Context, args GetSystemAutomationDestinationListOutputArgs, opts ...pulumi.InvokeOption) GetSystemAutomationDestinationListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemAutomationDestinationListResult, error) {
			args := v.(GetSystemAutomationDestinationListArgs)
			r, err := GetSystemAutomationDestinationList(ctx, &args, opts...)
			var s GetSystemAutomationDestinationListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemAutomationDestinationListResultOutput)
}

// A collection of arguments for invoking GetSystemAutomationDestinationList.
type GetSystemAutomationDestinationListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemAutomationDestinationListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutomationDestinationListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemAutomationDestinationList.
type GetSystemAutomationDestinationListResultOutput struct{ *pulumi.OutputState }

func (GetSystemAutomationDestinationListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemAutomationDestinationListResult)(nil)).Elem()
}

func (o GetSystemAutomationDestinationListResultOutput) ToGetSystemAutomationDestinationListResultOutput() GetSystemAutomationDestinationListResultOutput {
	return o
}

func (o GetSystemAutomationDestinationListResultOutput) ToGetSystemAutomationDestinationListResultOutputWithContext(ctx context.Context) GetSystemAutomationDestinationListResultOutput {
	return o
}

func (o GetSystemAutomationDestinationListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutomationDestinationListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemAutomationDestinationListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemAutomationDestinationListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemAutomationDestination`.
func (o GetSystemAutomationDestinationListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemAutomationDestinationListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemAutomationDestinationListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemAutomationDestinationListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemAutomationDestinationListResultOutput{})
}
