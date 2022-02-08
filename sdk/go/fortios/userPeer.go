// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure peer users.
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
// 		_, err := fortios.NewUserPeer(ctx, "trname1", &fortios.UserPeerArgs{
// 			Ca:                pulumi.String("EC-ACC"),
// 			CnType:            pulumi.String("string"),
// 			LdapMode:          pulumi.String("password"),
// 			MandatoryCaVerify: pulumi.String("enable"),
// 			TwoFactor:         pulumi.String("disable"),
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
// User Peer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/userPeer:UserPeer labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type UserPeer struct {
	pulumi.CustomResourceState

	// Name of the CA certificate as returned by the execute vpn certificate ca list command.
	Ca pulumi.StringOutput `pulumi:"ca"`
	// Peer certificate common name.
	Cn pulumi.StringOutput `pulumi:"cn"`
	// Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
	CnType pulumi.StringOutput `pulumi:"cnType"`
	// Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
	LdapMode pulumi.StringOutput `pulumi:"ldapMode"`
	// Password for LDAP server bind.
	LdapPassword pulumi.StringPtrOutput `pulumi:"ldapPassword"`
	// Name of an LDAP server defined under the user ldap command. Performs client access rights check.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// Username for LDAP server bind.
	LdapUsername pulumi.StringOutput `pulumi:"ldapUsername"`
	// Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
	MandatoryCaVerify pulumi.StringOutput `pulumi:"mandatoryCaVerify"`
	// Peer name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Online Certificate Status Protocol (OCSP) server for certificate retrieval.
	OcspOverrideServer pulumi.StringOutput `pulumi:"ocspOverrideServer"`
	// Peer's password used for two-factor authentication.
	Passwd pulumi.StringPtrOutput `pulumi:"passwd"`
	// Peer certificate name constraints.
	Subject pulumi.StringOutput `pulumi:"subject"`
	// Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
	TwoFactor pulumi.StringOutput `pulumi:"twoFactor"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserPeer registers a new resource with the given unique name, arguments, and options.
func NewUserPeer(ctx *pulumi.Context,
	name string, args *UserPeerArgs, opts ...pulumi.ResourceOption) (*UserPeer, error) {
	if args == nil {
		args = &UserPeerArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource UserPeer
	err := ctx.RegisterResource("fortios:index/userPeer:UserPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPeer gets an existing UserPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPeerState, opts ...pulumi.ResourceOption) (*UserPeer, error) {
	var resource UserPeer
	err := ctx.ReadResource("fortios:index/userPeer:UserPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPeer resources.
type userPeerState struct {
	// Name of the CA certificate as returned by the execute vpn certificate ca list command.
	Ca *string `pulumi:"ca"`
	// Peer certificate common name.
	Cn *string `pulumi:"cn"`
	// Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
	CnType *string `pulumi:"cnType"`
	// Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
	LdapMode *string `pulumi:"ldapMode"`
	// Password for LDAP server bind.
	LdapPassword *string `pulumi:"ldapPassword"`
	// Name of an LDAP server defined under the user ldap command. Performs client access rights check.
	LdapServer *string `pulumi:"ldapServer"`
	// Username for LDAP server bind.
	LdapUsername *string `pulumi:"ldapUsername"`
	// Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
	MandatoryCaVerify *string `pulumi:"mandatoryCaVerify"`
	// Peer name.
	Name *string `pulumi:"name"`
	// Online Certificate Status Protocol (OCSP) server for certificate retrieval.
	OcspOverrideServer *string `pulumi:"ocspOverrideServer"`
	// Peer's password used for two-factor authentication.
	Passwd *string `pulumi:"passwd"`
	// Peer certificate name constraints.
	Subject *string `pulumi:"subject"`
	// Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
	TwoFactor *string `pulumi:"twoFactor"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserPeerState struct {
	// Name of the CA certificate as returned by the execute vpn certificate ca list command.
	Ca pulumi.StringPtrInput
	// Peer certificate common name.
	Cn pulumi.StringPtrInput
	// Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
	CnType pulumi.StringPtrInput
	// Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
	LdapMode pulumi.StringPtrInput
	// Password for LDAP server bind.
	LdapPassword pulumi.StringPtrInput
	// Name of an LDAP server defined under the user ldap command. Performs client access rights check.
	LdapServer pulumi.StringPtrInput
	// Username for LDAP server bind.
	LdapUsername pulumi.StringPtrInput
	// Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
	MandatoryCaVerify pulumi.StringPtrInput
	// Peer name.
	Name pulumi.StringPtrInput
	// Online Certificate Status Protocol (OCSP) server for certificate retrieval.
	OcspOverrideServer pulumi.StringPtrInput
	// Peer's password used for two-factor authentication.
	Passwd pulumi.StringPtrInput
	// Peer certificate name constraints.
	Subject pulumi.StringPtrInput
	// Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
	TwoFactor pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPeerState)(nil)).Elem()
}

type userPeerArgs struct {
	// Name of the CA certificate as returned by the execute vpn certificate ca list command.
	Ca *string `pulumi:"ca"`
	// Peer certificate common name.
	Cn *string `pulumi:"cn"`
	// Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
	CnType *string `pulumi:"cnType"`
	// Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
	LdapMode *string `pulumi:"ldapMode"`
	// Password for LDAP server bind.
	LdapPassword *string `pulumi:"ldapPassword"`
	// Name of an LDAP server defined under the user ldap command. Performs client access rights check.
	LdapServer *string `pulumi:"ldapServer"`
	// Username for LDAP server bind.
	LdapUsername *string `pulumi:"ldapUsername"`
	// Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
	MandatoryCaVerify *string `pulumi:"mandatoryCaVerify"`
	// Peer name.
	Name *string `pulumi:"name"`
	// Online Certificate Status Protocol (OCSP) server for certificate retrieval.
	OcspOverrideServer *string `pulumi:"ocspOverrideServer"`
	// Peer's password used for two-factor authentication.
	Passwd *string `pulumi:"passwd"`
	// Peer certificate name constraints.
	Subject *string `pulumi:"subject"`
	// Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
	TwoFactor *string `pulumi:"twoFactor"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserPeer resource.
type UserPeerArgs struct {
	// Name of the CA certificate as returned by the execute vpn certificate ca list command.
	Ca pulumi.StringPtrInput
	// Peer certificate common name.
	Cn pulumi.StringPtrInput
	// Peer certificate common name type. Valid values: `string`, `email`, `FQDN`, `ipv4`, `ipv6`.
	CnType pulumi.StringPtrInput
	// Mode for LDAP peer authentication. Valid values: `password`, `principal-name`.
	LdapMode pulumi.StringPtrInput
	// Password for LDAP server bind.
	LdapPassword pulumi.StringPtrInput
	// Name of an LDAP server defined under the user ldap command. Performs client access rights check.
	LdapServer pulumi.StringPtrInput
	// Username for LDAP server bind.
	LdapUsername pulumi.StringPtrInput
	// Determine what happens to the peer if the CA certificate is not installed. Disable to automatically consider the peer certificate as valid. Valid values: `enable`, `disable`.
	MandatoryCaVerify pulumi.StringPtrInput
	// Peer name.
	Name pulumi.StringPtrInput
	// Online Certificate Status Protocol (OCSP) server for certificate retrieval.
	OcspOverrideServer pulumi.StringPtrInput
	// Peer's password used for two-factor authentication.
	Passwd pulumi.StringPtrInput
	// Peer certificate name constraints.
	Subject pulumi.StringPtrInput
	// Enable/disable two-factor authentication, applying certificate and password-based authentication. Valid values: `enable`, `disable`.
	TwoFactor pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPeerArgs)(nil)).Elem()
}

type UserPeerInput interface {
	pulumi.Input

	ToUserPeerOutput() UserPeerOutput
	ToUserPeerOutputWithContext(ctx context.Context) UserPeerOutput
}

func (*UserPeer) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPeer)(nil)).Elem()
}

func (i *UserPeer) ToUserPeerOutput() UserPeerOutput {
	return i.ToUserPeerOutputWithContext(context.Background())
}

func (i *UserPeer) ToUserPeerOutputWithContext(ctx context.Context) UserPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPeerOutput)
}

// UserPeerArrayInput is an input type that accepts UserPeerArray and UserPeerArrayOutput values.
// You can construct a concrete instance of `UserPeerArrayInput` via:
//
//          UserPeerArray{ UserPeerArgs{...} }
type UserPeerArrayInput interface {
	pulumi.Input

	ToUserPeerArrayOutput() UserPeerArrayOutput
	ToUserPeerArrayOutputWithContext(context.Context) UserPeerArrayOutput
}

type UserPeerArray []UserPeerInput

func (UserPeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPeer)(nil)).Elem()
}

func (i UserPeerArray) ToUserPeerArrayOutput() UserPeerArrayOutput {
	return i.ToUserPeerArrayOutputWithContext(context.Background())
}

func (i UserPeerArray) ToUserPeerArrayOutputWithContext(ctx context.Context) UserPeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPeerArrayOutput)
}

// UserPeerMapInput is an input type that accepts UserPeerMap and UserPeerMapOutput values.
// You can construct a concrete instance of `UserPeerMapInput` via:
//
//          UserPeerMap{ "key": UserPeerArgs{...} }
type UserPeerMapInput interface {
	pulumi.Input

	ToUserPeerMapOutput() UserPeerMapOutput
	ToUserPeerMapOutputWithContext(context.Context) UserPeerMapOutput
}

type UserPeerMap map[string]UserPeerInput

func (UserPeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPeer)(nil)).Elem()
}

func (i UserPeerMap) ToUserPeerMapOutput() UserPeerMapOutput {
	return i.ToUserPeerMapOutputWithContext(context.Background())
}

func (i UserPeerMap) ToUserPeerMapOutputWithContext(ctx context.Context) UserPeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPeerMapOutput)
}

type UserPeerOutput struct{ *pulumi.OutputState }

func (UserPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPeer)(nil)).Elem()
}

func (o UserPeerOutput) ToUserPeerOutput() UserPeerOutput {
	return o
}

func (o UserPeerOutput) ToUserPeerOutputWithContext(ctx context.Context) UserPeerOutput {
	return o
}

type UserPeerArrayOutput struct{ *pulumi.OutputState }

func (UserPeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPeer)(nil)).Elem()
}

func (o UserPeerArrayOutput) ToUserPeerArrayOutput() UserPeerArrayOutput {
	return o
}

func (o UserPeerArrayOutput) ToUserPeerArrayOutputWithContext(ctx context.Context) UserPeerArrayOutput {
	return o
}

func (o UserPeerArrayOutput) Index(i pulumi.IntInput) UserPeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPeer {
		return vs[0].([]*UserPeer)[vs[1].(int)]
	}).(UserPeerOutput)
}

type UserPeerMapOutput struct{ *pulumi.OutputState }

func (UserPeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPeer)(nil)).Elem()
}

func (o UserPeerMapOutput) ToUserPeerMapOutput() UserPeerMapOutput {
	return o
}

func (o UserPeerMapOutput) ToUserPeerMapOutputWithContext(ctx context.Context) UserPeerMapOutput {
	return o
}

func (o UserPeerMapOutput) MapIndex(k pulumi.StringInput) UserPeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPeer {
		return vs[0].(map[string]*UserPeer)[vs[1].(string)]
	}).(UserPeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPeerInput)(nil)).Elem(), &UserPeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPeerArrayInput)(nil)).Elem(), UserPeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPeerMapInput)(nil)).Elem(), UserPeerMap{})
	pulumi.RegisterOutputType(UserPeerOutput{})
	pulumi.RegisterOutputType(UserPeerArrayOutput{})
	pulumi.RegisterOutputType(UserPeerMapOutput{})
}
