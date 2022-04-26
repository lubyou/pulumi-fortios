// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSystemPasswordPolicy(ctx, "trname", &fortios.SystemPasswordPolicyArgs{
// 			ApplyTo:            pulumi.String("admin-password"),
// 			Change4Characters:  pulumi.String("disable"),
// 			ExpireDay:          pulumi.Int(90),
// 			ExpireStatus:       pulumi.String("disable"),
// 			MinLowerCaseLetter: pulumi.Int(0),
// 			MinNonAlphanumeric: pulumi.Int(0),
// 			MinNumber:          pulumi.Int(0),
// 			MinUpperCaseLetter: pulumi.Int(0),
// 			MinimumLength:      pulumi.Int(8),
// 			ReusePassword:      pulumi.String("enable"),
// 			Status:             pulumi.String("disable"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// System PasswordPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemPasswordPolicy:SystemPasswordPolicy labelname SystemPasswordPolicy
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemPasswordPolicy:SystemPasswordPolicy labelname SystemPasswordPolicy
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemPasswordPolicy struct {
	pulumi.CustomResourceState

	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo pulumi.StringOutput `pulumi:"applyTo"`
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters pulumi.StringOutput `pulumi:"change4Characters"`
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay pulumi.IntOutput `pulumi:"expireDay"`
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus pulumi.StringOutput `pulumi:"expireStatus"`
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters pulumi.IntOutput `pulumi:"minChangeCharacters"`
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter pulumi.IntOutput `pulumi:"minLowerCaseLetter"`
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric pulumi.IntOutput `pulumi:"minNonAlphanumeric"`
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber pulumi.IntOutput `pulumi:"minNumber"`
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter pulumi.IntOutput `pulumi:"minUpperCaseLetter"`
	// Minimum password length (8 - 128, default = 8).
	MinimumLength pulumi.IntOutput `pulumi:"minimumLength"`
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword pulumi.StringOutput `pulumi:"reusePassword"`
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemPasswordPolicy registers a new resource with the given unique name, arguments, and options.
func NewSystemPasswordPolicy(ctx *pulumi.Context,
	name string, args *SystemPasswordPolicyArgs, opts ...pulumi.ResourceOption) (*SystemPasswordPolicy, error) {
	if args == nil {
		args = &SystemPasswordPolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemPasswordPolicy
	err := ctx.RegisterResource("fortios:index/systemPasswordPolicy:SystemPasswordPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemPasswordPolicy gets an existing SystemPasswordPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemPasswordPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemPasswordPolicyState, opts ...pulumi.ResourceOption) (*SystemPasswordPolicy, error) {
	var resource SystemPasswordPolicy
	err := ctx.ReadResource("fortios:index/systemPasswordPolicy:SystemPasswordPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemPasswordPolicy resources.
type systemPasswordPolicyState struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo *string `pulumi:"applyTo"`
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters *string `pulumi:"change4Characters"`
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay *int `pulumi:"expireDay"`
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus *string `pulumi:"expireStatus"`
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters *int `pulumi:"minChangeCharacters"`
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter *int `pulumi:"minLowerCaseLetter"`
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric *int `pulumi:"minNonAlphanumeric"`
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber *int `pulumi:"minNumber"`
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter *int `pulumi:"minUpperCaseLetter"`
	// Minimum password length (8 - 128, default = 8).
	MinimumLength *int `pulumi:"minimumLength"`
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword *string `pulumi:"reusePassword"`
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemPasswordPolicyState struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo pulumi.StringPtrInput
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters pulumi.StringPtrInput
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay pulumi.IntPtrInput
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus pulumi.StringPtrInput
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters pulumi.IntPtrInput
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter pulumi.IntPtrInput
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric pulumi.IntPtrInput
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber pulumi.IntPtrInput
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter pulumi.IntPtrInput
	// Minimum password length (8 - 128, default = 8).
	MinimumLength pulumi.IntPtrInput
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword pulumi.StringPtrInput
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemPasswordPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemPasswordPolicyState)(nil)).Elem()
}

type systemPasswordPolicyArgs struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo *string `pulumi:"applyTo"`
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters *string `pulumi:"change4Characters"`
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay *int `pulumi:"expireDay"`
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus *string `pulumi:"expireStatus"`
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters *int `pulumi:"minChangeCharacters"`
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter *int `pulumi:"minLowerCaseLetter"`
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric *int `pulumi:"minNonAlphanumeric"`
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber *int `pulumi:"minNumber"`
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter *int `pulumi:"minUpperCaseLetter"`
	// Minimum password length (8 - 128, default = 8).
	MinimumLength *int `pulumi:"minimumLength"`
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword *string `pulumi:"reusePassword"`
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemPasswordPolicy resource.
type SystemPasswordPolicyArgs struct {
	// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space. Valid values: `admin-password`, `ipsec-preshared-key`.
	ApplyTo pulumi.StringPtrInput
	// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
	Change4Characters pulumi.StringPtrInput
	// Number of days after which passwords expire (1 - 999 days, default = 90).
	ExpireDay pulumi.IntPtrInput
	// Enable/disable password expiration. Valid values: `enable`, `disable`.
	ExpireStatus pulumi.StringPtrInput
	// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
	MinChangeCharacters pulumi.IntPtrInput
	// Minimum number of lowercase characters in password (0 - 128, default = 0).
	MinLowerCaseLetter pulumi.IntPtrInput
	// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
	MinNonAlphanumeric pulumi.IntPtrInput
	// Minimum number of numeric characters in password (0 - 128, default = 0).
	MinNumber pulumi.IntPtrInput
	// Minimum number of uppercase characters in password (0 - 128, default = 0).
	MinUpperCaseLetter pulumi.IntPtrInput
	// Minimum password length (8 - 128, default = 8).
	MinimumLength pulumi.IntPtrInput
	// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
	ReusePassword pulumi.StringPtrInput
	// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemPasswordPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemPasswordPolicyArgs)(nil)).Elem()
}

type SystemPasswordPolicyInput interface {
	pulumi.Input

	ToSystemPasswordPolicyOutput() SystemPasswordPolicyOutput
	ToSystemPasswordPolicyOutputWithContext(ctx context.Context) SystemPasswordPolicyOutput
}

func (*SystemPasswordPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemPasswordPolicy)(nil)).Elem()
}

func (i *SystemPasswordPolicy) ToSystemPasswordPolicyOutput() SystemPasswordPolicyOutput {
	return i.ToSystemPasswordPolicyOutputWithContext(context.Background())
}

func (i *SystemPasswordPolicy) ToSystemPasswordPolicyOutputWithContext(ctx context.Context) SystemPasswordPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordPolicyOutput)
}

// SystemPasswordPolicyArrayInput is an input type that accepts SystemPasswordPolicyArray and SystemPasswordPolicyArrayOutput values.
// You can construct a concrete instance of `SystemPasswordPolicyArrayInput` via:
//
//          SystemPasswordPolicyArray{ SystemPasswordPolicyArgs{...} }
type SystemPasswordPolicyArrayInput interface {
	pulumi.Input

	ToSystemPasswordPolicyArrayOutput() SystemPasswordPolicyArrayOutput
	ToSystemPasswordPolicyArrayOutputWithContext(context.Context) SystemPasswordPolicyArrayOutput
}

type SystemPasswordPolicyArray []SystemPasswordPolicyInput

func (SystemPasswordPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemPasswordPolicy)(nil)).Elem()
}

func (i SystemPasswordPolicyArray) ToSystemPasswordPolicyArrayOutput() SystemPasswordPolicyArrayOutput {
	return i.ToSystemPasswordPolicyArrayOutputWithContext(context.Background())
}

func (i SystemPasswordPolicyArray) ToSystemPasswordPolicyArrayOutputWithContext(ctx context.Context) SystemPasswordPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordPolicyArrayOutput)
}

// SystemPasswordPolicyMapInput is an input type that accepts SystemPasswordPolicyMap and SystemPasswordPolicyMapOutput values.
// You can construct a concrete instance of `SystemPasswordPolicyMapInput` via:
//
//          SystemPasswordPolicyMap{ "key": SystemPasswordPolicyArgs{...} }
type SystemPasswordPolicyMapInput interface {
	pulumi.Input

	ToSystemPasswordPolicyMapOutput() SystemPasswordPolicyMapOutput
	ToSystemPasswordPolicyMapOutputWithContext(context.Context) SystemPasswordPolicyMapOutput
}

type SystemPasswordPolicyMap map[string]SystemPasswordPolicyInput

func (SystemPasswordPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemPasswordPolicy)(nil)).Elem()
}

func (i SystemPasswordPolicyMap) ToSystemPasswordPolicyMapOutput() SystemPasswordPolicyMapOutput {
	return i.ToSystemPasswordPolicyMapOutputWithContext(context.Background())
}

func (i SystemPasswordPolicyMap) ToSystemPasswordPolicyMapOutputWithContext(ctx context.Context) SystemPasswordPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordPolicyMapOutput)
}

type SystemPasswordPolicyOutput struct{ *pulumi.OutputState }

func (SystemPasswordPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemPasswordPolicy)(nil)).Elem()
}

func (o SystemPasswordPolicyOutput) ToSystemPasswordPolicyOutput() SystemPasswordPolicyOutput {
	return o
}

func (o SystemPasswordPolicyOutput) ToSystemPasswordPolicyOutputWithContext(ctx context.Context) SystemPasswordPolicyOutput {
	return o
}

type SystemPasswordPolicyArrayOutput struct{ *pulumi.OutputState }

func (SystemPasswordPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemPasswordPolicy)(nil)).Elem()
}

func (o SystemPasswordPolicyArrayOutput) ToSystemPasswordPolicyArrayOutput() SystemPasswordPolicyArrayOutput {
	return o
}

func (o SystemPasswordPolicyArrayOutput) ToSystemPasswordPolicyArrayOutputWithContext(ctx context.Context) SystemPasswordPolicyArrayOutput {
	return o
}

func (o SystemPasswordPolicyArrayOutput) Index(i pulumi.IntInput) SystemPasswordPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemPasswordPolicy {
		return vs[0].([]*SystemPasswordPolicy)[vs[1].(int)]
	}).(SystemPasswordPolicyOutput)
}

type SystemPasswordPolicyMapOutput struct{ *pulumi.OutputState }

func (SystemPasswordPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemPasswordPolicy)(nil)).Elem()
}

func (o SystemPasswordPolicyMapOutput) ToSystemPasswordPolicyMapOutput() SystemPasswordPolicyMapOutput {
	return o
}

func (o SystemPasswordPolicyMapOutput) ToSystemPasswordPolicyMapOutputWithContext(ctx context.Context) SystemPasswordPolicyMapOutput {
	return o
}

func (o SystemPasswordPolicyMapOutput) MapIndex(k pulumi.StringInput) SystemPasswordPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemPasswordPolicy {
		return vs[0].(map[string]*SystemPasswordPolicy)[vs[1].(string)]
	}).(SystemPasswordPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordPolicyInput)(nil)).Elem(), &SystemPasswordPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordPolicyArrayInput)(nil)).Elem(), SystemPasswordPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordPolicyMapInput)(nil)).Elem(), SystemPasswordPolicyMap{})
	pulumi.RegisterOutputType(SystemPasswordPolicyOutput{})
	pulumi.RegisterOutputType(SystemPasswordPolicyArrayOutput{})
	pulumi.RegisterOutputType(SystemPasswordPolicyMapOutput{})
}
