// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupSystemFortisandbox(ctx *pulumi.Context, args *LookupSystemFortisandboxArgs, opts ...pulumi.InvokeOption) (*LookupSystemFortisandboxResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemFortisandboxResult
	err := ctx.Invoke("fortios:index/getSystemFortisandbox:GetSystemFortisandbox", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemFortisandbox.
type LookupSystemFortisandboxArgs struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemFortisandbox.
type LookupSystemFortisandboxResult struct {
	Email        string `pulumi:"email"`
	EncAlgorithm string `pulumi:"encAlgorithm"`
	Forticloud   string `pulumi:"forticloud"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string  `pulumi:"id"`
	InlineScan            string  `pulumi:"inlineScan"`
	Interface             string  `pulumi:"interface"`
	InterfaceSelectMethod string  `pulumi:"interfaceSelectMethod"`
	Server                string  `pulumi:"server"`
	SourceIp              string  `pulumi:"sourceIp"`
	SslMinProtoVersion    string  `pulumi:"sslMinProtoVersion"`
	Status                string  `pulumi:"status"`
	Vdomparam             *string `pulumi:"vdomparam"`
}

func LookupSystemFortisandboxOutput(ctx *pulumi.Context, args LookupSystemFortisandboxOutputArgs, opts ...pulumi.InvokeOption) LookupSystemFortisandboxResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemFortisandboxResult, error) {
			args := v.(LookupSystemFortisandboxArgs)
			r, err := LookupSystemFortisandbox(ctx, &args, opts...)
			var s LookupSystemFortisandboxResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemFortisandboxResultOutput)
}

// A collection of arguments for invoking GetSystemFortisandbox.
type LookupSystemFortisandboxOutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemFortisandboxOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFortisandboxArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemFortisandbox.
type LookupSystemFortisandboxResultOutput struct{ *pulumi.OutputState }

func (LookupSystemFortisandboxResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFortisandboxResult)(nil)).Elem()
}

func (o LookupSystemFortisandboxResultOutput) ToLookupSystemFortisandboxResultOutput() LookupSystemFortisandboxResultOutput {
	return o
}

func (o LookupSystemFortisandboxResultOutput) ToLookupSystemFortisandboxResultOutputWithContext(ctx context.Context) LookupSystemFortisandboxResultOutput {
	return o
}

func (o LookupSystemFortisandboxResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSystemFortisandboxResult] {
	return pulumix.Output[LookupSystemFortisandboxResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupSystemFortisandboxResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) EncAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.EncAlgorithm }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) Forticloud() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Forticloud }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemFortisandboxResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) InlineScan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.InlineScan }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Interface }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Server }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) SslMinProtoVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.SslMinProtoVersion }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemFortisandboxResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemFortisandboxResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemFortisandboxResultOutput{})
}
