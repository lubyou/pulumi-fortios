// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemClusterSync`.
func GetSystemClusterSyncList(ctx *pulumi.Context, args *GetSystemClusterSyncListArgs, opts ...pulumi.InvokeOption) (*GetSystemClusterSyncListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemClusterSyncListResult
	err := ctx.Invoke("fortios:index/getSystemClusterSyncList:GetSystemClusterSyncList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemClusterSyncList.
type GetSystemClusterSyncListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemClusterSyncList.
type GetSystemClusterSyncListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemClusterSync`.
	SyncIdlists []int   `pulumi:"syncIdlists"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func GetSystemClusterSyncListOutput(ctx *pulumi.Context, args GetSystemClusterSyncListOutputArgs, opts ...pulumi.InvokeOption) GetSystemClusterSyncListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemClusterSyncListResult, error) {
			args := v.(GetSystemClusterSyncListArgs)
			r, err := GetSystemClusterSyncList(ctx, &args, opts...)
			var s GetSystemClusterSyncListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemClusterSyncListResultOutput)
}

// A collection of arguments for invoking GetSystemClusterSyncList.
type GetSystemClusterSyncListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemClusterSyncListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemClusterSyncListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemClusterSyncList.
type GetSystemClusterSyncListResultOutput struct{ *pulumi.OutputState }

func (GetSystemClusterSyncListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemClusterSyncListResult)(nil)).Elem()
}

func (o GetSystemClusterSyncListResultOutput) ToGetSystemClusterSyncListResultOutput() GetSystemClusterSyncListResultOutput {
	return o
}

func (o GetSystemClusterSyncListResultOutput) ToGetSystemClusterSyncListResultOutputWithContext(ctx context.Context) GetSystemClusterSyncListResultOutput {
	return o
}

func (o GetSystemClusterSyncListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemClusterSyncListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemClusterSyncListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemClusterSyncListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemClusterSync`.
func (o GetSystemClusterSyncListResultOutput) SyncIdlists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSystemClusterSyncListResult) []int { return v.SyncIdlists }).(pulumi.IntArrayOutput)
}

func (o GetSystemClusterSyncListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemClusterSyncListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemClusterSyncListResultOutput{})
}
