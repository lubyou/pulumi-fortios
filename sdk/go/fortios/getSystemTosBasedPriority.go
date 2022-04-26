// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system tosbasedpriority
func LookupSystemTosBasedPriority(ctx *pulumi.Context, args *LookupSystemTosBasedPriorityArgs, opts ...pulumi.InvokeOption) (*LookupSystemTosBasedPriorityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemTosBasedPriorityResult
	err := ctx.Invoke("fortios:index/getSystemTosBasedPriority:GetSystemTosBasedPriority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemTosBasedPriority.
type LookupSystemTosBasedPriorityArgs struct {
	// Specify the fosid of the desired system tosbasedpriority.
	Fosid int `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemTosBasedPriority.
type LookupSystemTosBasedPriorityResult struct {
	// Item ID.
	Fosid int `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium).
	Priority string `pulumi:"priority"`
	// Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
	Tos       int     `pulumi:"tos"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemTosBasedPriorityOutput(ctx *pulumi.Context, args LookupSystemTosBasedPriorityOutputArgs, opts ...pulumi.InvokeOption) LookupSystemTosBasedPriorityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemTosBasedPriorityResult, error) {
			args := v.(LookupSystemTosBasedPriorityArgs)
			r, err := LookupSystemTosBasedPriority(ctx, &args, opts...)
			var s LookupSystemTosBasedPriorityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemTosBasedPriorityResultOutput)
}

// A collection of arguments for invoking GetSystemTosBasedPriority.
type LookupSystemTosBasedPriorityOutputArgs struct {
	// Specify the fosid of the desired system tosbasedpriority.
	Fosid pulumi.IntInput `pulumi:"fosid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemTosBasedPriorityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTosBasedPriorityArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemTosBasedPriority.
type LookupSystemTosBasedPriorityResultOutput struct{ *pulumi.OutputState }

func (LookupSystemTosBasedPriorityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTosBasedPriorityResult)(nil)).Elem()
}

func (o LookupSystemTosBasedPriorityResultOutput) ToLookupSystemTosBasedPriorityResultOutput() LookupSystemTosBasedPriorityResultOutput {
	return o
}

func (o LookupSystemTosBasedPriorityResultOutput) ToLookupSystemTosBasedPriorityResultOutputWithContext(ctx context.Context) LookupSystemTosBasedPriorityResultOutput {
	return o
}

// Item ID.
func (o LookupSystemTosBasedPriorityResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemTosBasedPriorityResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemTosBasedPriorityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTosBasedPriorityResult) string { return v.Id }).(pulumi.StringOutput)
}

// ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium).
func (o LookupSystemTosBasedPriorityResultOutput) Priority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTosBasedPriorityResult) string { return v.Priority }).(pulumi.StringOutput)
}

// Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
func (o LookupSystemTosBasedPriorityResultOutput) Tos() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemTosBasedPriorityResult) int { return v.Tos }).(pulumi.IntOutput)
}

func (o LookupSystemTosBasedPriorityResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemTosBasedPriorityResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemTosBasedPriorityResultOutput{})
}
