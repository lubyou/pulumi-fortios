// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system externalresource
func LookupSystemExternalResource(ctx *pulumi.Context, args *LookupSystemExternalResourceArgs, opts ...pulumi.InvokeOption) (*LookupSystemExternalResourceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemExternalResourceResult
	err := ctx.Invoke("fortios:index/getSystemExternalResource:GetSystemExternalResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemExternalResource.
type LookupSystemExternalResourceArgs struct {
	// Specify the name of the desired system externalresource.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemExternalResource.
type LookupSystemExternalResourceResult struct {
	// User resource category.
	Category int `pulumi:"category"`
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interface string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server.
	InterfaceSelectMethod string `pulumi:"interfaceSelectMethod"`
	// External resource name.
	Name string `pulumi:"name"`
	// HTTP basic authentication password.
	Password string `pulumi:"password"`
	// Time interval to refresh external resource (1 - 43200 min, default = 5 min).
	RefreshRate int `pulumi:"refreshRate"`
	// URI of external resource.
	Resource string `pulumi:"resource"`
	// Source IPv4 address used to communicate with server.
	SourceIp string `pulumi:"sourceIp"`
	// Enable/disable user resource.
	Status string `pulumi:"status"`
	// User resource type.
	Type string `pulumi:"type"`
	// Override HTTP User-Agent header used when retrieving this external resource.
	UserAgent string `pulumi:"userAgent"`
	// HTTP basic authentication user name.
	Username string `pulumi:"username"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemExternalResourceOutput(ctx *pulumi.Context, args LookupSystemExternalResourceOutputArgs, opts ...pulumi.InvokeOption) LookupSystemExternalResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemExternalResourceResult, error) {
			args := v.(LookupSystemExternalResourceArgs)
			r, err := LookupSystemExternalResource(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemExternalResourceResultOutput)
}

// A collection of arguments for invoking GetSystemExternalResource.
type LookupSystemExternalResourceOutputArgs struct {
	// Specify the name of the desired system externalresource.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemExternalResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemExternalResourceArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemExternalResource.
type LookupSystemExternalResourceResultOutput struct{ *pulumi.OutputState }

func (LookupSystemExternalResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemExternalResourceResult)(nil)).Elem()
}

func (o LookupSystemExternalResourceResultOutput) ToLookupSystemExternalResourceResultOutput() LookupSystemExternalResourceResultOutput {
	return o
}

func (o LookupSystemExternalResourceResultOutput) ToLookupSystemExternalResourceResultOutputWithContext(ctx context.Context) LookupSystemExternalResourceResultOutput {
	return o
}

// User resource category.
func (o LookupSystemExternalResourceResultOutput) Category() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) int { return v.Category }).(pulumi.IntOutput)
}

// Comment.
func (o LookupSystemExternalResourceResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemExternalResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o LookupSystemExternalResourceResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server.
func (o LookupSystemExternalResourceResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// External resource name.
func (o LookupSystemExternalResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// HTTP basic authentication password.
func (o LookupSystemExternalResourceResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Password }).(pulumi.StringOutput)
}

// Time interval to refresh external resource (1 - 43200 min, default = 5 min).
func (o LookupSystemExternalResourceResultOutput) RefreshRate() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) int { return v.RefreshRate }).(pulumi.IntOutput)
}

// URI of external resource.
func (o LookupSystemExternalResourceResultOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Resource }).(pulumi.StringOutput)
}

// Source IPv4 address used to communicate with server.
func (o LookupSystemExternalResourceResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable user resource.
func (o LookupSystemExternalResourceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Status }).(pulumi.StringOutput)
}

// User resource type.
func (o LookupSystemExternalResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

// Override HTTP User-Agent header used when retrieving this external resource.
func (o LookupSystemExternalResourceResultOutput) UserAgent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.UserAgent }).(pulumi.StringOutput)
}

// HTTP basic authentication user name.
func (o LookupSystemExternalResourceResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Username }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupSystemExternalResourceResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupSystemExternalResourceResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemExternalResourceResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemExternalResourceResultOutput{})
}
