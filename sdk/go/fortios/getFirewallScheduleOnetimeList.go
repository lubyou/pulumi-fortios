// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a list of `FirewallScheduleOnetime`.
func GetFirewallScheduleOnetimeList(ctx *pulumi.Context, args *GetFirewallScheduleOnetimeListArgs, opts ...pulumi.InvokeOption) (*GetFirewallScheduleOnetimeListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallScheduleOnetimeListResult
	err := ctx.Invoke("fortios:index/getFirewallScheduleOnetimeList:GetFirewallScheduleOnetimeList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallScheduleOnetimeList.
type GetFirewallScheduleOnetimeListArgs struct {
	Filter *string `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallScheduleOnetimeList.
type GetFirewallScheduleOnetimeListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the `FirewallScheduleOnetime`.
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallScheduleOnetimeListOutput(ctx *pulumi.Context, args GetFirewallScheduleOnetimeListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallScheduleOnetimeListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallScheduleOnetimeListResult, error) {
			args := v.(GetFirewallScheduleOnetimeListArgs)
			r, err := GetFirewallScheduleOnetimeList(ctx, &args, opts...)
			return *r, err
		}).(GetFirewallScheduleOnetimeListResultOutput)
}

// A collection of arguments for invoking GetFirewallScheduleOnetimeList.
type GetFirewallScheduleOnetimeListOutputArgs struct {
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallScheduleOnetimeListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallScheduleOnetimeListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallScheduleOnetimeList.
type GetFirewallScheduleOnetimeListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallScheduleOnetimeListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallScheduleOnetimeListResult)(nil)).Elem()
}

func (o GetFirewallScheduleOnetimeListResultOutput) ToGetFirewallScheduleOnetimeListResultOutput() GetFirewallScheduleOnetimeListResultOutput {
	return o
}

func (o GetFirewallScheduleOnetimeListResultOutput) ToGetFirewallScheduleOnetimeListResultOutputWithContext(ctx context.Context) GetFirewallScheduleOnetimeListResultOutput {
	return o
}

func (o GetFirewallScheduleOnetimeListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallScheduleOnetimeListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallScheduleOnetimeListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallScheduleOnetimeListResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of the `FirewallScheduleOnetime`.
func (o GetFirewallScheduleOnetimeListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallScheduleOnetimeListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallScheduleOnetimeListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallScheduleOnetimeListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallScheduleOnetimeListResultOutput{})
}
