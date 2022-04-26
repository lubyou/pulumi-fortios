// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system dscpbasedpriority
func LookupSystemDscpBasedPriority(ctx *pulumi.Context, args *LookupSystemDscpBasedPriorityArgs, opts ...pulumi.InvokeOption) (*LookupSystemDscpBasedPriorityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemDscpBasedPriorityResult
	err := ctx.Invoke("fortios:index/getSystemDscpBasedPriority:GetSystemDscpBasedPriority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDscpBasedPriority.
type LookupSystemDscpBasedPriorityArgs struct {
	// Specify the fosid of the desired system dscpbasedpriority.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDscpBasedPriority.
type LookupSystemDscpBasedPriorityResult struct {
	// DSCP(DiffServ) DS value (0 - 63).
	Ds int `pulumi:"ds"`
	// Item ID.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// DSCP based priority level.
	Priority  string  `pulumi:"priority"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemDscpBasedPriorityOutput(ctx *pulumi.Context, args LookupSystemDscpBasedPriorityOutputArgs, opts ...pulumi.InvokeOption) LookupSystemDscpBasedPriorityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemDscpBasedPriorityResult, error) {
			args := v.(LookupSystemDscpBasedPriorityArgs)
			r, err := LookupSystemDscpBasedPriority(ctx, &args, opts...)
			var s LookupSystemDscpBasedPriorityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemDscpBasedPriorityResultOutput)
}

// A collection of arguments for invoking GetSystemDscpBasedPriority.
type LookupSystemDscpBasedPriorityOutputArgs struct {
	// Specify the fosid of the desired system dscpbasedpriority.
	Fosid pulumi.IntInput `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemDscpBasedPriorityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDscpBasedPriorityArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemDscpBasedPriority.
type LookupSystemDscpBasedPriorityResultOutput struct{ *pulumi.OutputState }

func (LookupSystemDscpBasedPriorityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDscpBasedPriorityResult)(nil)).Elem()
}

func (o LookupSystemDscpBasedPriorityResultOutput) ToLookupSystemDscpBasedPriorityResultOutput() LookupSystemDscpBasedPriorityResultOutput {
	return o
}

func (o LookupSystemDscpBasedPriorityResultOutput) ToLookupSystemDscpBasedPriorityResultOutputWithContext(ctx context.Context) LookupSystemDscpBasedPriorityResultOutput {
	return o
}

// DSCP(DiffServ) DS value (0 - 63).
func (o LookupSystemDscpBasedPriorityResultOutput) Ds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDscpBasedPriorityResult) int { return v.Ds }).(pulumi.IntOutput)
}

// Item ID.
func (o LookupSystemDscpBasedPriorityResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDscpBasedPriorityResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemDscpBasedPriorityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDscpBasedPriorityResult) string { return v.Id }).(pulumi.StringOutput)
}

// DSCP based priority level.
func (o LookupSystemDscpBasedPriorityResultOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDscpBasedPriorityResult) string { return v.Priority }).(pulumi.StringOutput)
}

func (o LookupSystemDscpBasedPriorityResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemDscpBasedPriorityResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemDscpBasedPriorityResultOutput{})
}
