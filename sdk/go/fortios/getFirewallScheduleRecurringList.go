// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFirewallScheduleRecurringList(ctx *pulumi.Context, args *GetFirewallScheduleRecurringListArgs, opts ...pulumi.InvokeOption) (*GetFirewallScheduleRecurringListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFirewallScheduleRecurringListResult
	err := ctx.Invoke("fortios:index/getFirewallScheduleRecurringList:GetFirewallScheduleRecurringList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallScheduleRecurringList.
type GetFirewallScheduleRecurringListArgs struct {
	Filter    *string `pulumi:"filter"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallScheduleRecurringList.
type GetFirewallScheduleRecurringListResult struct {
	Filter *string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Namelists []string `pulumi:"namelists"`
	Vdomparam *string  `pulumi:"vdomparam"`
}

func GetFirewallScheduleRecurringListOutput(ctx *pulumi.Context, args GetFirewallScheduleRecurringListOutputArgs, opts ...pulumi.InvokeOption) GetFirewallScheduleRecurringListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFirewallScheduleRecurringListResult, error) {
			args := v.(GetFirewallScheduleRecurringListArgs)
			r, err := GetFirewallScheduleRecurringList(ctx, &args, opts...)
			var s GetFirewallScheduleRecurringListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFirewallScheduleRecurringListResultOutput)
}

// A collection of arguments for invoking GetFirewallScheduleRecurringList.
type GetFirewallScheduleRecurringListOutputArgs struct {
	Filter    pulumi.StringPtrInput `pulumi:"filter"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (GetFirewallScheduleRecurringListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallScheduleRecurringListArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallScheduleRecurringList.
type GetFirewallScheduleRecurringListResultOutput struct{ *pulumi.OutputState }

func (GetFirewallScheduleRecurringListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallScheduleRecurringListResult)(nil)).Elem()
}

func (o GetFirewallScheduleRecurringListResultOutput) ToGetFirewallScheduleRecurringListResultOutput() GetFirewallScheduleRecurringListResultOutput {
	return o
}

func (o GetFirewallScheduleRecurringListResultOutput) ToGetFirewallScheduleRecurringListResultOutputWithContext(ctx context.Context) GetFirewallScheduleRecurringListResultOutput {
	return o
}

func (o GetFirewallScheduleRecurringListResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallScheduleRecurringListResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFirewallScheduleRecurringListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFirewallScheduleRecurringListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFirewallScheduleRecurringListResultOutput) Namelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFirewallScheduleRecurringListResult) []string { return v.Namelists }).(pulumi.StringArrayOutput)
}

func (o GetFirewallScheduleRecurringListResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallScheduleRecurringListResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallScheduleRecurringListResultOutput{})
}
