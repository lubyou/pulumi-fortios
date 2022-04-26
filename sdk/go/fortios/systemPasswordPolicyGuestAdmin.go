// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure the password policy for guest administrators.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewSystemPasswordPolicyGuestAdmin(ctx, "trname", &fortios.SystemPasswordPolicyGuestAdminArgs{
// 			ApplyTo:            pulumi.String("guest-admin-password"),
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
// System PasswordPolicyGuestAdmin can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/systemPasswordPolicyGuestAdmin:SystemPasswordPolicyGuestAdmin labelname SystemPasswordPolicyGuestAdmin
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/systemPasswordPolicyGuestAdmin:SystemPasswordPolicyGuestAdmin labelname SystemPasswordPolicyGuestAdmin
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type SystemPasswordPolicyGuestAdmin struct {
	pulumi.CustomResourceState

	// Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
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

// NewSystemPasswordPolicyGuestAdmin registers a new resource with the given unique name, arguments, and options.
func NewSystemPasswordPolicyGuestAdmin(ctx *pulumi.Context,
	name string, args *SystemPasswordPolicyGuestAdminArgs, opts ...pulumi.ResourceOption) (*SystemPasswordPolicyGuestAdmin, error) {
	if args == nil {
		args = &SystemPasswordPolicyGuestAdminArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemPasswordPolicyGuestAdmin
	err := ctx.RegisterResource("fortios:index/systemPasswordPolicyGuestAdmin:SystemPasswordPolicyGuestAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemPasswordPolicyGuestAdmin gets an existing SystemPasswordPolicyGuestAdmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemPasswordPolicyGuestAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemPasswordPolicyGuestAdminState, opts ...pulumi.ResourceOption) (*SystemPasswordPolicyGuestAdmin, error) {
	var resource SystemPasswordPolicyGuestAdmin
	err := ctx.ReadResource("fortios:index/systemPasswordPolicyGuestAdmin:SystemPasswordPolicyGuestAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemPasswordPolicyGuestAdmin resources.
type systemPasswordPolicyGuestAdminState struct {
	// Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
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

type SystemPasswordPolicyGuestAdminState struct {
	// Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
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

func (SystemPasswordPolicyGuestAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemPasswordPolicyGuestAdminState)(nil)).Elem()
}

type systemPasswordPolicyGuestAdminArgs struct {
	// Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
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

// The set of arguments for constructing a SystemPasswordPolicyGuestAdmin resource.
type SystemPasswordPolicyGuestAdminArgs struct {
	// Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
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

func (SystemPasswordPolicyGuestAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemPasswordPolicyGuestAdminArgs)(nil)).Elem()
}

type SystemPasswordPolicyGuestAdminInput interface {
	pulumi.Input

	ToSystemPasswordPolicyGuestAdminOutput() SystemPasswordPolicyGuestAdminOutput
	ToSystemPasswordPolicyGuestAdminOutputWithContext(ctx context.Context) SystemPasswordPolicyGuestAdminOutput
}

func (*SystemPasswordPolicyGuestAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemPasswordPolicyGuestAdmin)(nil)).Elem()
}

func (i *SystemPasswordPolicyGuestAdmin) ToSystemPasswordPolicyGuestAdminOutput() SystemPasswordPolicyGuestAdminOutput {
	return i.ToSystemPasswordPolicyGuestAdminOutputWithContext(context.Background())
}

func (i *SystemPasswordPolicyGuestAdmin) ToSystemPasswordPolicyGuestAdminOutputWithContext(ctx context.Context) SystemPasswordPolicyGuestAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordPolicyGuestAdminOutput)
}

// SystemPasswordPolicyGuestAdminArrayInput is an input type that accepts SystemPasswordPolicyGuestAdminArray and SystemPasswordPolicyGuestAdminArrayOutput values.
// You can construct a concrete instance of `SystemPasswordPolicyGuestAdminArrayInput` via:
//
//          SystemPasswordPolicyGuestAdminArray{ SystemPasswordPolicyGuestAdminArgs{...} }
type SystemPasswordPolicyGuestAdminArrayInput interface {
	pulumi.Input

	ToSystemPasswordPolicyGuestAdminArrayOutput() SystemPasswordPolicyGuestAdminArrayOutput
	ToSystemPasswordPolicyGuestAdminArrayOutputWithContext(context.Context) SystemPasswordPolicyGuestAdminArrayOutput
}

type SystemPasswordPolicyGuestAdminArray []SystemPasswordPolicyGuestAdminInput

func (SystemPasswordPolicyGuestAdminArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemPasswordPolicyGuestAdmin)(nil)).Elem()
}

func (i SystemPasswordPolicyGuestAdminArray) ToSystemPasswordPolicyGuestAdminArrayOutput() SystemPasswordPolicyGuestAdminArrayOutput {
	return i.ToSystemPasswordPolicyGuestAdminArrayOutputWithContext(context.Background())
}

func (i SystemPasswordPolicyGuestAdminArray) ToSystemPasswordPolicyGuestAdminArrayOutputWithContext(ctx context.Context) SystemPasswordPolicyGuestAdminArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordPolicyGuestAdminArrayOutput)
}

// SystemPasswordPolicyGuestAdminMapInput is an input type that accepts SystemPasswordPolicyGuestAdminMap and SystemPasswordPolicyGuestAdminMapOutput values.
// You can construct a concrete instance of `SystemPasswordPolicyGuestAdminMapInput` via:
//
//          SystemPasswordPolicyGuestAdminMap{ "key": SystemPasswordPolicyGuestAdminArgs{...} }
type SystemPasswordPolicyGuestAdminMapInput interface {
	pulumi.Input

	ToSystemPasswordPolicyGuestAdminMapOutput() SystemPasswordPolicyGuestAdminMapOutput
	ToSystemPasswordPolicyGuestAdminMapOutputWithContext(context.Context) SystemPasswordPolicyGuestAdminMapOutput
}

type SystemPasswordPolicyGuestAdminMap map[string]SystemPasswordPolicyGuestAdminInput

func (SystemPasswordPolicyGuestAdminMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemPasswordPolicyGuestAdmin)(nil)).Elem()
}

func (i SystemPasswordPolicyGuestAdminMap) ToSystemPasswordPolicyGuestAdminMapOutput() SystemPasswordPolicyGuestAdminMapOutput {
	return i.ToSystemPasswordPolicyGuestAdminMapOutputWithContext(context.Background())
}

func (i SystemPasswordPolicyGuestAdminMap) ToSystemPasswordPolicyGuestAdminMapOutputWithContext(ctx context.Context) SystemPasswordPolicyGuestAdminMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPasswordPolicyGuestAdminMapOutput)
}

type SystemPasswordPolicyGuestAdminOutput struct{ *pulumi.OutputState }

func (SystemPasswordPolicyGuestAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemPasswordPolicyGuestAdmin)(nil)).Elem()
}

func (o SystemPasswordPolicyGuestAdminOutput) ToSystemPasswordPolicyGuestAdminOutput() SystemPasswordPolicyGuestAdminOutput {
	return o
}

func (o SystemPasswordPolicyGuestAdminOutput) ToSystemPasswordPolicyGuestAdminOutputWithContext(ctx context.Context) SystemPasswordPolicyGuestAdminOutput {
	return o
}

type SystemPasswordPolicyGuestAdminArrayOutput struct{ *pulumi.OutputState }

func (SystemPasswordPolicyGuestAdminArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemPasswordPolicyGuestAdmin)(nil)).Elem()
}

func (o SystemPasswordPolicyGuestAdminArrayOutput) ToSystemPasswordPolicyGuestAdminArrayOutput() SystemPasswordPolicyGuestAdminArrayOutput {
	return o
}

func (o SystemPasswordPolicyGuestAdminArrayOutput) ToSystemPasswordPolicyGuestAdminArrayOutputWithContext(ctx context.Context) SystemPasswordPolicyGuestAdminArrayOutput {
	return o
}

func (o SystemPasswordPolicyGuestAdminArrayOutput) Index(i pulumi.IntInput) SystemPasswordPolicyGuestAdminOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemPasswordPolicyGuestAdmin {
		return vs[0].([]*SystemPasswordPolicyGuestAdmin)[vs[1].(int)]
	}).(SystemPasswordPolicyGuestAdminOutput)
}

type SystemPasswordPolicyGuestAdminMapOutput struct{ *pulumi.OutputState }

func (SystemPasswordPolicyGuestAdminMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemPasswordPolicyGuestAdmin)(nil)).Elem()
}

func (o SystemPasswordPolicyGuestAdminMapOutput) ToSystemPasswordPolicyGuestAdminMapOutput() SystemPasswordPolicyGuestAdminMapOutput {
	return o
}

func (o SystemPasswordPolicyGuestAdminMapOutput) ToSystemPasswordPolicyGuestAdminMapOutputWithContext(ctx context.Context) SystemPasswordPolicyGuestAdminMapOutput {
	return o
}

func (o SystemPasswordPolicyGuestAdminMapOutput) MapIndex(k pulumi.StringInput) SystemPasswordPolicyGuestAdminOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemPasswordPolicyGuestAdmin {
		return vs[0].(map[string]*SystemPasswordPolicyGuestAdmin)[vs[1].(string)]
	}).(SystemPasswordPolicyGuestAdminOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordPolicyGuestAdminInput)(nil)).Elem(), &SystemPasswordPolicyGuestAdmin{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordPolicyGuestAdminArrayInput)(nil)).Elem(), SystemPasswordPolicyGuestAdminArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPasswordPolicyGuestAdminMapInput)(nil)).Elem(), SystemPasswordPolicyGuestAdminMap{})
	pulumi.RegisterOutputType(SystemPasswordPolicyGuestAdminOutput{})
	pulumi.RegisterOutputType(SystemPasswordPolicyGuestAdminArrayOutput{})
	pulumi.RegisterOutputType(SystemPasswordPolicyGuestAdminMapOutput{})
}
