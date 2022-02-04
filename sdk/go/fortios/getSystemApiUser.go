// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system apiuser
func LookupSystemApiUser(ctx *pulumi.Context, args *LookupSystemApiUserArgs, opts ...pulumi.InvokeOption) (*LookupSystemApiUserResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemApiUserResult
	err := ctx.Invoke("fortios:index/getSystemApiUser:GetSystemApiUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemApiUser.
type LookupSystemApiUserArgs struct {
	// Specify the name of the desired system apiuser.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemApiUser.
type LookupSystemApiUserResult struct {
	// Admin user access profile.
	Accprofile string `pulumi:"accprofile"`
	// Admin user password.
	ApiKey string `pulumi:"apiKey"`
	// Comment.
	Comments string `pulumi:"comments"`
	// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
	CorsAllowOrigin string `pulumi:"corsAllowOrigin"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Virtual domain name.
	Name string `pulumi:"name"`
	// Enable/disable peer authentication.
	PeerAuth string `pulumi:"peerAuth"`
	// Peer group name.
	PeerGroup string `pulumi:"peerGroup"`
	// Schedule name.
	Schedule string `pulumi:"schedule"`
	// Trusthost. The structure of `trusthost` block is documented below.
	Trusthosts []GetSystemApiUserTrusthost `pulumi:"trusthosts"`
	Vdomparam  *string                     `pulumi:"vdomparam"`
	// Virtual domains. The structure of `vdom` block is documented below.
	Vdoms []GetSystemApiUserVdom `pulumi:"vdoms"`
}

func LookupSystemApiUserOutput(ctx *pulumi.Context, args LookupSystemApiUserOutputArgs, opts ...pulumi.InvokeOption) LookupSystemApiUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemApiUserResult, error) {
			args := v.(LookupSystemApiUserArgs)
			r, err := LookupSystemApiUser(ctx, &args, opts...)
			return *r, err
		}).(LookupSystemApiUserResultOutput)
}

// A collection of arguments for invoking GetSystemApiUser.
type LookupSystemApiUserOutputArgs struct {
	// Specify the name of the desired system apiuser.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemApiUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemApiUserArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemApiUser.
type LookupSystemApiUserResultOutput struct{ *pulumi.OutputState }

func (LookupSystemApiUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemApiUserResult)(nil)).Elem()
}

func (o LookupSystemApiUserResultOutput) ToLookupSystemApiUserResultOutput() LookupSystemApiUserResultOutput {
	return o
}

func (o LookupSystemApiUserResultOutput) ToLookupSystemApiUserResultOutputWithContext(ctx context.Context) LookupSystemApiUserResultOutput {
	return o
}

// Admin user access profile.
func (o LookupSystemApiUserResultOutput) Accprofile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.Accprofile }).(pulumi.StringOutput)
}

// Admin user password.
func (o LookupSystemApiUserResultOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.ApiKey }).(pulumi.StringOutput)
}

// Comment.
func (o LookupSystemApiUserResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.Comments }).(pulumi.StringOutput)
}

// Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
func (o LookupSystemApiUserResultOutput) CorsAllowOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.CorsAllowOrigin }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemApiUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Virtual domain name.
func (o LookupSystemApiUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable peer authentication.
func (o LookupSystemApiUserResultOutput) PeerAuth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.PeerAuth }).(pulumi.StringOutput)
}

// Peer group name.
func (o LookupSystemApiUserResultOutput) PeerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.PeerGroup }).(pulumi.StringOutput)
}

// Schedule name.
func (o LookupSystemApiUserResultOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) string { return v.Schedule }).(pulumi.StringOutput)
}

// Trusthost. The structure of `trusthost` block is documented below.
func (o LookupSystemApiUserResultOutput) Trusthosts() GetSystemApiUserTrusthostArrayOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) []GetSystemApiUserTrusthost { return v.Trusthosts }).(GetSystemApiUserTrusthostArrayOutput)
}

func (o LookupSystemApiUserResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Virtual domains. The structure of `vdom` block is documented below.
func (o LookupSystemApiUserResultOutput) Vdoms() GetSystemApiUserVdomArrayOutput {
	return o.ApplyT(func(v LookupSystemApiUserResult) []GetSystemApiUserVdom { return v.Vdoms }).(GetSystemApiUserVdomArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemApiUserResultOutput{})
}
