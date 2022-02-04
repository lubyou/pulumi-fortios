// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SAML server entry configuration. Applies to FortiOS Version `>= 6.2.4`.
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
// 		_, err := fortios.NewUserSaml(ctx, "tr3", &fortios.UserSamlArgs{
// 			Cert:               pulumi.String("Fortinet_Factory"),
// 			EntityId:           pulumi.String("https://1.1.1.1"),
// 			IdpCert:            pulumi.String("cer11"),
// 			IdpEntityId:        pulumi.String("https://1.1.1.1/acc"),
// 			IdpSingleLogoutUrl: pulumi.String("https://1.1.1.1/lo"),
// 			IdpSingleSignOnUrl: pulumi.String("https://1.1.1.1/sou"),
// 			SingleLogoutUrl:    pulumi.String("https://1.1.1.1/logout"),
// 			SingleSignOnUrl:    pulumi.String("https://1.1.1.1/sign"),
// 			UserName:           pulumi.String("ad111"),
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
// User Saml can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/userSaml:UserSaml labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type UserSaml struct {
	pulumi.CustomResourceState

	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim pulumi.StringOutput `pulumi:"adfsClaim"`
	// Certificate to sign SAML messages.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance pulumi.IntOutput `pulumi:"clockTolerance"`
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod pulumi.StringOutput `pulumi:"digestMethod"`
	// SP entity ID.
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType pulumi.StringOutput `pulumi:"groupClaimType"`
	// Group name in assertion statement.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// IDP Certificate name.
	IdpCert pulumi.StringOutput `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId pulumi.StringOutput `pulumi:"idpEntityId"`
	// IDP single logout url.
	IdpSingleLogoutUrl pulumi.StringOutput `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringOutput `pulumi:"idpSingleSignOnUrl"`
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate pulumi.StringOutput `pulumi:"limitRelaystate"`
	// SAML server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringOutput `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringOutput `pulumi:"singleSignOnUrl"`
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType pulumi.StringOutput `pulumi:"userClaimType"`
	// User name in assertion statement.
	UserName pulumi.StringOutput `pulumi:"userName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserSaml registers a new resource with the given unique name, arguments, and options.
func NewUserSaml(ctx *pulumi.Context,
	name string, args *UserSamlArgs, opts ...pulumi.ResourceOption) (*UserSaml, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntityId == nil {
		return nil, errors.New("invalid value for required argument 'EntityId'")
	}
	if args.IdpCert == nil {
		return nil, errors.New("invalid value for required argument 'IdpCert'")
	}
	if args.IdpEntityId == nil {
		return nil, errors.New("invalid value for required argument 'IdpEntityId'")
	}
	if args.IdpSingleSignOnUrl == nil {
		return nil, errors.New("invalid value for required argument 'IdpSingleSignOnUrl'")
	}
	if args.SingleSignOnUrl == nil {
		return nil, errors.New("invalid value for required argument 'SingleSignOnUrl'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserSaml
	err := ctx.RegisterResource("fortios:index/userSaml:UserSaml", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserSaml gets an existing UserSaml resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserSaml(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSamlState, opts ...pulumi.ResourceOption) (*UserSaml, error) {
	var resource UserSaml
	err := ctx.ReadResource("fortios:index/userSaml:UserSaml", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserSaml resources.
type userSamlState struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim *string `pulumi:"adfsClaim"`
	// Certificate to sign SAML messages.
	Cert *string `pulumi:"cert"`
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance *int `pulumi:"clockTolerance"`
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod *string `pulumi:"digestMethod"`
	// SP entity ID.
	EntityId *string `pulumi:"entityId"`
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType *string `pulumi:"groupClaimType"`
	// Group name in assertion statement.
	GroupName *string `pulumi:"groupName"`
	// IDP Certificate name.
	IdpCert *string `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId *string `pulumi:"idpEntityId"`
	// IDP single logout url.
	IdpSingleLogoutUrl *string `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl *string `pulumi:"idpSingleSignOnUrl"`
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate *string `pulumi:"limitRelaystate"`
	// SAML server entry name.
	Name *string `pulumi:"name"`
	// SP single logout URL.
	SingleLogoutUrl *string `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl *string `pulumi:"singleSignOnUrl"`
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType *string `pulumi:"userClaimType"`
	// User name in assertion statement.
	UserName *string `pulumi:"userName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserSamlState struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim pulumi.StringPtrInput
	// Certificate to sign SAML messages.
	Cert pulumi.StringPtrInput
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance pulumi.IntPtrInput
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod pulumi.StringPtrInput
	// SP entity ID.
	EntityId pulumi.StringPtrInput
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType pulumi.StringPtrInput
	// Group name in assertion statement.
	GroupName pulumi.StringPtrInput
	// IDP Certificate name.
	IdpCert pulumi.StringPtrInput
	// IDP entity ID.
	IdpEntityId pulumi.StringPtrInput
	// IDP single logout url.
	IdpSingleLogoutUrl pulumi.StringPtrInput
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringPtrInput
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate pulumi.StringPtrInput
	// SAML server entry name.
	Name pulumi.StringPtrInput
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringPtrInput
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringPtrInput
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType pulumi.StringPtrInput
	// User name in assertion statement.
	UserName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserSamlState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSamlState)(nil)).Elem()
}

type userSamlArgs struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim *string `pulumi:"adfsClaim"`
	// Certificate to sign SAML messages.
	Cert *string `pulumi:"cert"`
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance *int `pulumi:"clockTolerance"`
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod *string `pulumi:"digestMethod"`
	// SP entity ID.
	EntityId string `pulumi:"entityId"`
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType *string `pulumi:"groupClaimType"`
	// Group name in assertion statement.
	GroupName *string `pulumi:"groupName"`
	// IDP Certificate name.
	IdpCert string `pulumi:"idpCert"`
	// IDP entity ID.
	IdpEntityId string `pulumi:"idpEntityId"`
	// IDP single logout url.
	IdpSingleLogoutUrl *string `pulumi:"idpSingleLogoutUrl"`
	// IDP single sign-on URL.
	IdpSingleSignOnUrl string `pulumi:"idpSingleSignOnUrl"`
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate *string `pulumi:"limitRelaystate"`
	// SAML server entry name.
	Name *string `pulumi:"name"`
	// SP single logout URL.
	SingleLogoutUrl *string `pulumi:"singleLogoutUrl"`
	// SP single sign-on URL.
	SingleSignOnUrl string `pulumi:"singleSignOnUrl"`
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType *string `pulumi:"userClaimType"`
	// User name in assertion statement.
	UserName *string `pulumi:"userName"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserSaml resource.
type UserSamlArgs struct {
	// Enable/disable ADFS Claim for user/group attribute in assertion statement (default = disable). Valid values: `enable`, `disable`.
	AdfsClaim pulumi.StringPtrInput
	// Certificate to sign SAML messages.
	Cert pulumi.StringPtrInput
	// Clock skew tolerance in seconds (0 - 300, default = 15, 0 = no tolerance).
	ClockTolerance pulumi.IntPtrInput
	// Digest Method Algorithm. (default = sha1). Valid values: `sha1`, `sha256`.
	DigestMethod pulumi.StringPtrInput
	// SP entity ID.
	EntityId pulumi.StringInput
	// Group claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	GroupClaimType pulumi.StringPtrInput
	// Group name in assertion statement.
	GroupName pulumi.StringPtrInput
	// IDP Certificate name.
	IdpCert pulumi.StringInput
	// IDP entity ID.
	IdpEntityId pulumi.StringInput
	// IDP single logout url.
	IdpSingleLogoutUrl pulumi.StringPtrInput
	// IDP single sign-on URL.
	IdpSingleSignOnUrl pulumi.StringInput
	// Enable/disable limiting of relay-state parameter when it exceeds SAML 2.0 specification limits (80 bytes). Valid values: `enable`, `disable`.
	LimitRelaystate pulumi.StringPtrInput
	// SAML server entry name.
	Name pulumi.StringPtrInput
	// SP single logout URL.
	SingleLogoutUrl pulumi.StringPtrInput
	// SP single sign-on URL.
	SingleSignOnUrl pulumi.StringInput
	// User name claim in assertion statement. Valid values: `email`, `given-name`, `name`, `upn`, `common-name`, `email-adfs-1x`, `group`, `upn-adfs-1x`, `role`, `sur-name`, `ppid`, `name-identifier`, `authentication-method`, `deny-only-group-sid`, `deny-only-primary-sid`, `deny-only-primary-group-sid`, `group-sid`, `primary-group-sid`, `primary-sid`, `windows-account-name`.
	UserClaimType pulumi.StringPtrInput
	// User name in assertion statement.
	UserName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserSamlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSamlArgs)(nil)).Elem()
}

type UserSamlInput interface {
	pulumi.Input

	ToUserSamlOutput() UserSamlOutput
	ToUserSamlOutputWithContext(ctx context.Context) UserSamlOutput
}

func (*UserSaml) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSaml)(nil)).Elem()
}

func (i *UserSaml) ToUserSamlOutput() UserSamlOutput {
	return i.ToUserSamlOutputWithContext(context.Background())
}

func (i *UserSaml) ToUserSamlOutputWithContext(ctx context.Context) UserSamlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSamlOutput)
}

// UserSamlArrayInput is an input type that accepts UserSamlArray and UserSamlArrayOutput values.
// You can construct a concrete instance of `UserSamlArrayInput` via:
//
//          UserSamlArray{ UserSamlArgs{...} }
type UserSamlArrayInput interface {
	pulumi.Input

	ToUserSamlArrayOutput() UserSamlArrayOutput
	ToUserSamlArrayOutputWithContext(context.Context) UserSamlArrayOutput
}

type UserSamlArray []UserSamlInput

func (UserSamlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSaml)(nil)).Elem()
}

func (i UserSamlArray) ToUserSamlArrayOutput() UserSamlArrayOutput {
	return i.ToUserSamlArrayOutputWithContext(context.Background())
}

func (i UserSamlArray) ToUserSamlArrayOutputWithContext(ctx context.Context) UserSamlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSamlArrayOutput)
}

// UserSamlMapInput is an input type that accepts UserSamlMap and UserSamlMapOutput values.
// You can construct a concrete instance of `UserSamlMapInput` via:
//
//          UserSamlMap{ "key": UserSamlArgs{...} }
type UserSamlMapInput interface {
	pulumi.Input

	ToUserSamlMapOutput() UserSamlMapOutput
	ToUserSamlMapOutputWithContext(context.Context) UserSamlMapOutput
}

type UserSamlMap map[string]UserSamlInput

func (UserSamlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSaml)(nil)).Elem()
}

func (i UserSamlMap) ToUserSamlMapOutput() UserSamlMapOutput {
	return i.ToUserSamlMapOutputWithContext(context.Background())
}

func (i UserSamlMap) ToUserSamlMapOutputWithContext(ctx context.Context) UserSamlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSamlMapOutput)
}

type UserSamlOutput struct{ *pulumi.OutputState }

func (UserSamlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSaml)(nil)).Elem()
}

func (o UserSamlOutput) ToUserSamlOutput() UserSamlOutput {
	return o
}

func (o UserSamlOutput) ToUserSamlOutputWithContext(ctx context.Context) UserSamlOutput {
	return o
}

type UserSamlArrayOutput struct{ *pulumi.OutputState }

func (UserSamlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserSaml)(nil)).Elem()
}

func (o UserSamlArrayOutput) ToUserSamlArrayOutput() UserSamlArrayOutput {
	return o
}

func (o UserSamlArrayOutput) ToUserSamlArrayOutputWithContext(ctx context.Context) UserSamlArrayOutput {
	return o
}

func (o UserSamlArrayOutput) Index(i pulumi.IntInput) UserSamlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserSaml {
		return vs[0].([]*UserSaml)[vs[1].(int)]
	}).(UserSamlOutput)
}

type UserSamlMapOutput struct{ *pulumi.OutputState }

func (UserSamlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserSaml)(nil)).Elem()
}

func (o UserSamlMapOutput) ToUserSamlMapOutput() UserSamlMapOutput {
	return o
}

func (o UserSamlMapOutput) ToUserSamlMapOutputWithContext(ctx context.Context) UserSamlMapOutput {
	return o
}

func (o UserSamlMapOutput) MapIndex(k pulumi.StringInput) UserSamlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserSaml {
		return vs[0].(map[string]*UserSaml)[vs[1].(string)]
	}).(UserSamlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserSamlInput)(nil)).Elem(), &UserSaml{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSamlArrayInput)(nil)).Elem(), UserSamlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserSamlMapInput)(nil)).Elem(), UserSamlMap{})
	pulumi.RegisterOutputType(UserSamlOutput{})
	pulumi.RegisterOutputType(UserSamlArrayOutput{})
	pulumi.RegisterOutputType(UserSamlMapOutput{})
}
