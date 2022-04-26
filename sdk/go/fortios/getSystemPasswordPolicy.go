// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system passwordpolicy
func LookupSystemPasswordPolicy(ctx *pulumi.Context, args *LookupSystemPasswordPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSystemPasswordPolicyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemPasswordPolicyResult
	err := ctx.Invoke("fortios:index/getSystemPasswordPolicy:GetSystemPasswordPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemPasswordPolicy.
type LookupSystemPasswordPolicyArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemPasswordPolicy.
type LookupSystemPasswordPolicyResult struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space.
	ApplyTo string `pulumi:"applyTo"`
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled).
	Change4Characters string `pulumi:"change4Characters"`
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay int `pulumi:"expireDay"`
	// Enable/disable password expiration.
	ExpireStatus string `pulumi:"expireStatus"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters int `pulumi:"minChangeCharacters"`
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter int `pulumi:"minLowerCaseLetter"`
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric int `pulumi:"minNonAlphanumeric"`
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber int `pulumi:"minNumber"`
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter int `pulumi:"minUpperCaseLetter"`
	// Minimum password length (8 - 128, default = 8).
	MinimumLength int `pulumi:"minimumLength"`
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides).
	ReusePassword string `pulumi:"reusePassword"`
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemPasswordPolicyOutput(ctx *pulumi.Context, args LookupSystemPasswordPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSystemPasswordPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemPasswordPolicyResult, error) {
			args := v.(LookupSystemPasswordPolicyArgs)
			r, err := LookupSystemPasswordPolicy(ctx, &args, opts...)
			var s LookupSystemPasswordPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemPasswordPolicyResultOutput)
}

// A collection of arguments for invoking GetSystemPasswordPolicy.
type LookupSystemPasswordPolicyOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemPasswordPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemPasswordPolicyArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemPasswordPolicy.
type LookupSystemPasswordPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSystemPasswordPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemPasswordPolicyResult)(nil)).Elem()
}

func (o LookupSystemPasswordPolicyResultOutput) ToLookupSystemPasswordPolicyResultOutput() LookupSystemPasswordPolicyResultOutput {
	return o
}

func (o LookupSystemPasswordPolicyResultOutput) ToLookupSystemPasswordPolicyResultOutputWithContext(ctx context.Context) LookupSystemPasswordPolicyResultOutput {
	return o
}

// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space.
func (o LookupSystemPasswordPolicyResultOutput) ApplyTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) string { return v.ApplyTo }).(pulumi.StringOutput)
}

// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled).
func (o LookupSystemPasswordPolicyResultOutput) Change4Characters() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) string { return v.Change4Characters }).(pulumi.StringOutput)
}

// Number of days after which passwords expire (1 - 999 days, default = 90).
func (o LookupSystemPasswordPolicyResultOutput) ExpireDay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) int { return v.ExpireDay }).(pulumi.IntOutput)
}

// Enable/disable password expiration.
func (o LookupSystemPasswordPolicyResultOutput) ExpireStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) string { return v.ExpireStatus }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemPasswordPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
func (o LookupSystemPasswordPolicyResultOutput) MinChangeCharacters() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) int { return v.MinChangeCharacters }).(pulumi.IntOutput)
}

// Minimum number of lowercase characters in password (0 - 128, default = 0).
func (o LookupSystemPasswordPolicyResultOutput) MinLowerCaseLetter() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) int { return v.MinLowerCaseLetter }).(pulumi.IntOutput)
}

// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
func (o LookupSystemPasswordPolicyResultOutput) MinNonAlphanumeric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) int { return v.MinNonAlphanumeric }).(pulumi.IntOutput)
}

// Minimum number of numeric characters in password (0 - 128, default = 0).
func (o LookupSystemPasswordPolicyResultOutput) MinNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) int { return v.MinNumber }).(pulumi.IntOutput)
}

// Minimum number of uppercase characters in password (0 - 128, default = 0).
func (o LookupSystemPasswordPolicyResultOutput) MinUpperCaseLetter() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) int { return v.MinUpperCaseLetter }).(pulumi.IntOutput)
}

// Minimum password length (8 - 128, default = 8).
func (o LookupSystemPasswordPolicyResultOutput) MinimumLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) int { return v.MinimumLength }).(pulumi.IntOutput)
}

// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides).
func (o LookupSystemPasswordPolicyResultOutput) ReusePassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) string { return v.ReusePassword }).(pulumi.StringOutput)
}

// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
func (o LookupSystemPasswordPolicyResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemPasswordPolicyResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemPasswordPolicyResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemPasswordPolicyResultOutput{})
}
