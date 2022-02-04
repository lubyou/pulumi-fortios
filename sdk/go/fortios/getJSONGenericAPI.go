// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a FortiAPI Generic Interface data source.
func LookupJSONGenericAPI(ctx *pulumi.Context, args *LookupJSONGenericAPIArgs, opts ...pulumi.InvokeOption) (*LookupJSONGenericAPIResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupJSONGenericAPIResult
	err := ctx.Invoke("fortios:index/getJSONGenericAPI:GetJSONGenericAPI", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetJSONGenericAPI.
type LookupJSONGenericAPIArgs struct {
	// FortiAPI URL path.
	Path string `pulumi:"path"`
	// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
	Specialparams *string `pulumi:"specialparams"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetJSONGenericAPI.
type LookupJSONGenericAPIResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// FortiAPI URL path.
	Path string `pulumi:"path"`
	// FortiAPI returns results.
	Response string `pulumi:"response"`
	// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
	Specialparams *string `pulumi:"specialparams"`
	Vdomparam     *string `pulumi:"vdomparam"`
}

func LookupJSONGenericAPIOutput(ctx *pulumi.Context, args LookupJSONGenericAPIOutputArgs, opts ...pulumi.InvokeOption) LookupJSONGenericAPIResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJSONGenericAPIResult, error) {
			args := v.(LookupJSONGenericAPIArgs)
			r, err := LookupJSONGenericAPI(ctx, &args, opts...)
			return *r, err
		}).(LookupJSONGenericAPIResultOutput)
}

// A collection of arguments for invoking GetJSONGenericAPI.
type LookupJSONGenericAPIOutputArgs struct {
	// FortiAPI URL path.
	Path pulumi.StringInput `pulumi:"path"`
	// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
	Specialparams pulumi.StringPtrInput `pulumi:"specialparams"`
	Vdomparam     pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupJSONGenericAPIOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJSONGenericAPIArgs)(nil)).Elem()
}

// A collection of values returned by GetJSONGenericAPI.
type LookupJSONGenericAPIResultOutput struct{ *pulumi.OutputState }

func (LookupJSONGenericAPIResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJSONGenericAPIResult)(nil)).Elem()
}

func (o LookupJSONGenericAPIResultOutput) ToLookupJSONGenericAPIResultOutput() LookupJSONGenericAPIResultOutput {
	return o
}

func (o LookupJSONGenericAPIResultOutput) ToLookupJSONGenericAPIResultOutputWithContext(ctx context.Context) LookupJSONGenericAPIResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupJSONGenericAPIResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJSONGenericAPIResult) string { return v.Id }).(pulumi.StringOutput)
}

// FortiAPI URL path.
func (o LookupJSONGenericAPIResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJSONGenericAPIResult) string { return v.Path }).(pulumi.StringOutput)
}

// FortiAPI returns results.
func (o LookupJSONGenericAPIResultOutput) Response() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJSONGenericAPIResult) string { return v.Response }).(pulumi.StringOutput)
}

// URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
func (o LookupJSONGenericAPIResultOutput) Specialparams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJSONGenericAPIResult) *string { return v.Specialparams }).(pulumi.StringPtrOutput)
}

func (o LookupJSONGenericAPIResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJSONGenericAPIResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJSONGenericAPIResultOutput{})
}
