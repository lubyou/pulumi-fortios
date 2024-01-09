// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemNetflow(ctx *pulumi.Context, args *LookupSystemNetflowArgs, opts ...pulumi.InvokeOption) (*LookupSystemNetflowResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemNetflowResult
	err := ctx.Invoke("fortios:index/getSystemNetflow:GetSystemNetflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemNetflow.
type LookupSystemNetflowArgs struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemNetflow.
type LookupSystemNetflowResult struct {
	ActiveFlowTimeout int    `pulumi:"activeFlowTimeout"`
	CollectorIp       string `pulumi:"collectorIp"`
	CollectorPort     int    `pulumi:"collectorPort"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string  `pulumi:"id"`
	InactiveFlowTimeout   int     `pulumi:"inactiveFlowTimeout"`
	Interface             string  `pulumi:"interface"`
	InterfaceSelectMethod string  `pulumi:"interfaceSelectMethod"`
	SourceIp              string  `pulumi:"sourceIp"`
	TemplateTxCounter     int     `pulumi:"templateTxCounter"`
	TemplateTxTimeout     int     `pulumi:"templateTxTimeout"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

func LookupSystemNetflowOutput(ctx *pulumi.Context, args LookupSystemNetflowOutputArgs, opts ...pulumi.InvokeOption) LookupSystemNetflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemNetflowResult, error) {
			args := v.(LookupSystemNetflowArgs)
			r, err := LookupSystemNetflow(ctx, &args, opts...)
			var s LookupSystemNetflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemNetflowResultOutput)
}

// A collection of arguments for invoking GetSystemNetflow.
type LookupSystemNetflowOutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemNetflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemNetflowArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemNetflow.
type LookupSystemNetflowResultOutput struct{ *pulumi.OutputState }

func (LookupSystemNetflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemNetflowResult)(nil)).Elem()
}

func (o LookupSystemNetflowResultOutput) ToLookupSystemNetflowResultOutput() LookupSystemNetflowResultOutput {
	return o
}

func (o LookupSystemNetflowResultOutput) ToLookupSystemNetflowResultOutputWithContext(ctx context.Context) LookupSystemNetflowResultOutput {
	return o
}

func (o LookupSystemNetflowResultOutput) ActiveFlowTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) int { return v.ActiveFlowTimeout }).(pulumi.IntOutput)
}

func (o LookupSystemNetflowResultOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) string { return v.CollectorIp }).(pulumi.StringOutput)
}

func (o LookupSystemNetflowResultOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) int { return v.CollectorPort }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemNetflowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemNetflowResultOutput) InactiveFlowTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) int { return v.InactiveFlowTimeout }).(pulumi.IntOutput)
}

func (o LookupSystemNetflowResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) string { return v.Interface }).(pulumi.StringOutput)
}

func (o LookupSystemNetflowResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o LookupSystemNetflowResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LookupSystemNetflowResultOutput) TemplateTxCounter() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) int { return v.TemplateTxCounter }).(pulumi.IntOutput)
}

func (o LookupSystemNetflowResultOutput) TemplateTxTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) int { return v.TemplateTxTimeout }).(pulumi.IntOutput)
}

func (o LookupSystemNetflowResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemNetflowResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemNetflowResultOutput{})
}
