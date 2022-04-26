// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `SystemDnsDatabase`.
func GetSystemDnsDatabaseList(ctx *pulumi.Context, args *GetSystemDnsDatabaseListArgs, opts ...pulumi.InvokeOption) (*GetSystemDnsDatabaseListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSystemDnsDatabaseListResult
	err := ctx.Invoke("fortios:index/getSystemDnsDatabaseList:GetSystemDnsDatabaseList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDnsDatabaseList.
type GetSystemDnsDatabaseListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDnsDatabaseList.
type GetSystemDnsDatabaseListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `SystemDnsDatabase`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetSystemDnsDatabaseListOutput(ctx *pulumi.Context, args GetSystemDnsDatabaseListOutputArgs, opts ...pulumi.InvokeOption) GetSystemDnsDatabaseListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSystemDnsDatabaseListResult, error) {
			args := v.(GetSystemDnsDatabaseListArgs)
			r, err := GetSystemDnsDatabaseList(ctx, &args, opts...)
			var s GetSystemDnsDatabaseListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSystemDnsDatabaseListResultOutput)
}

// A collection of arguments for invoking GetSystemDnsDatabaseList.
type GetSystemDnsDatabaseListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetSystemDnsDatabaseListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemDnsDatabaseListArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemDnsDatabaseList.
type GetSystemDnsDatabaseListResultOutput struct{ *pulumi.OutputState }

func (GetSystemDnsDatabaseListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSystemDnsDatabaseListResult)(nil)).Elem()
}

func (o GetSystemDnsDatabaseListResultOutput) ToGetSystemDnsDatabaseListResultOutput() GetSystemDnsDatabaseListResultOutput {
	return o
}

func (o GetSystemDnsDatabaseListResultOutput) ToGetSystemDnsDatabaseListResultOutputWithContext(ctx context.Context) GetSystemDnsDatabaseListResultOutput {
	return o
}

func (o GetSystemDnsDatabaseListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemDnsDatabaseListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSystemDnsDatabaseListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSystemDnsDatabaseListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `SystemDnsDatabase`.
func (o GetSystemDnsDatabaseListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSystemDnsDatabaseListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetSystemDnsDatabaseListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSystemDnsDatabaseListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSystemDnsDatabaseListResultOutput{})
}
