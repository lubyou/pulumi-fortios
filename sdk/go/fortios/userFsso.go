// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Fortinet Single Sign On (FSSO) agents.
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
// 		_, err := fortios.NewUserFsso(ctx, "trname", &fortios.UserFssoArgs{
// 			Port:      pulumi.Int(32381),
// 			Port2:     pulumi.Int(8000),
// 			Port3:     pulumi.Int(8000),
// 			Port4:     pulumi.Int(8000),
// 			Port5:     pulumi.Int(8000),
// 			Server:    pulumi.String("1.1.1.1"),
// 			SourceIp:  pulumi.String("0.0.0.0"),
// 			SourceIp6: pulumi.String("::"),
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
// User Fsso can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/userFsso:UserFsso labelname {{name}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/userFsso:UserFsso labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type UserFsso struct {
	pulumi.CustomResourceState

	// Interval in minutes within to fetch groups from FSSO server, or unset to disable.
	GroupPollInterval pulumi.IntOutput `pulumi:"groupPollInterval"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
	LdapPoll pulumi.StringOutput `pulumi:"ldapPoll"`
	// Filter used to fetch groups.
	LdapPollFilter pulumi.StringOutput `pulumi:"ldapPollFilter"`
	// Interval in minutes within to fetch groups from LDAP server.
	LdapPollInterval pulumi.IntOutput `pulumi:"ldapPollInterval"`
	// LDAP server to get group information.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// Interval in minutes to keep logons after FSSO server down.
	LogonTimeout pulumi.IntOutput `pulumi:"logonTimeout"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password of the first FSSO collector agent.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Password of the second FSSO collector agent.
	Password2 pulumi.StringPtrOutput `pulumi:"password2"`
	// Password of the third FSSO collector agent.
	Password3 pulumi.StringPtrOutput `pulumi:"password3"`
	// Password of the fourth FSSO collector agent.
	Password4 pulumi.StringPtrOutput `pulumi:"password4"`
	// Password of the fifth FSSO collector agent.
	Password5 pulumi.StringPtrOutput `pulumi:"password5"`
	// Port of the first FSSO collector agent.
	Port pulumi.IntOutput `pulumi:"port"`
	// Port of the second FSSO collector agent.
	Port2 pulumi.IntOutput `pulumi:"port2"`
	// Port of the third FSSO collector agent.
	Port3 pulumi.IntOutput `pulumi:"port3"`
	// Port of the fourth FSSO collector agent.
	Port4 pulumi.IntOutput `pulumi:"port4"`
	// Port of the fifth FSSO collector agent.
	Port5 pulumi.IntOutput `pulumi:"port5"`
	// Domain name or IP address of the first FSSO collector agent.
	Server pulumi.StringOutput `pulumi:"server"`
	// Domain name or IP address of the second FSSO collector agent.
	Server2 pulumi.StringOutput `pulumi:"server2"`
	// Domain name or IP address of the third FSSO collector agent.
	Server3 pulumi.StringOutput `pulumi:"server3"`
	// Domain name or IP address of the fourth FSSO collector agent.
	Server4 pulumi.StringOutput `pulumi:"server4"`
	// Domain name or IP address of the fifth FSSO collector agent.
	Server5 pulumi.StringOutput `pulumi:"server5"`
	// Source IP for communications to FSSO agent.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// IPv6 source for communications to FSSO agent.
	SourceIp6 pulumi.StringOutput `pulumi:"sourceIp6"`
	// Enable/disable use of SSL. Valid values: `enable`, `disable`.
	Ssl pulumi.StringOutput `pulumi:"ssl"`
	// Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
	SslServerHostIpCheck pulumi.StringOutput `pulumi:"sslServerHostIpCheck"`
	// Trusted server certificate or CA certificate.
	SslTrustedCert pulumi.StringOutput `pulumi:"sslTrustedCert"`
	// Server type.
	Type pulumi.StringOutput `pulumi:"type"`
	// LDAP server to get user information.
	UserInfoServer pulumi.StringOutput `pulumi:"userInfoServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserFsso registers a new resource with the given unique name, arguments, and options.
func NewUserFsso(ctx *pulumi.Context,
	name string, args *UserFssoArgs, opts ...pulumi.ResourceOption) (*UserFsso, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserFsso
	err := ctx.RegisterResource("fortios:index/userFsso:UserFsso", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserFsso gets an existing UserFsso resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserFsso(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserFssoState, opts ...pulumi.ResourceOption) (*UserFsso, error) {
	var resource UserFsso
	err := ctx.ReadResource("fortios:index/userFsso:UserFsso", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserFsso resources.
type userFssoState struct {
	// Interval in minutes within to fetch groups from FSSO server, or unset to disable.
	GroupPollInterval *int `pulumi:"groupPollInterval"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
	LdapPoll *string `pulumi:"ldapPoll"`
	// Filter used to fetch groups.
	LdapPollFilter *string `pulumi:"ldapPollFilter"`
	// Interval in minutes within to fetch groups from LDAP server.
	LdapPollInterval *int `pulumi:"ldapPollInterval"`
	// LDAP server to get group information.
	LdapServer *string `pulumi:"ldapServer"`
	// Interval in minutes to keep logons after FSSO server down.
	LogonTimeout *int `pulumi:"logonTimeout"`
	// Name.
	Name *string `pulumi:"name"`
	// Password of the first FSSO collector agent.
	Password *string `pulumi:"password"`
	// Password of the second FSSO collector agent.
	Password2 *string `pulumi:"password2"`
	// Password of the third FSSO collector agent.
	Password3 *string `pulumi:"password3"`
	// Password of the fourth FSSO collector agent.
	Password4 *string `pulumi:"password4"`
	// Password of the fifth FSSO collector agent.
	Password5 *string `pulumi:"password5"`
	// Port of the first FSSO collector agent.
	Port *int `pulumi:"port"`
	// Port of the second FSSO collector agent.
	Port2 *int `pulumi:"port2"`
	// Port of the third FSSO collector agent.
	Port3 *int `pulumi:"port3"`
	// Port of the fourth FSSO collector agent.
	Port4 *int `pulumi:"port4"`
	// Port of the fifth FSSO collector agent.
	Port5 *int `pulumi:"port5"`
	// Domain name or IP address of the first FSSO collector agent.
	Server *string `pulumi:"server"`
	// Domain name or IP address of the second FSSO collector agent.
	Server2 *string `pulumi:"server2"`
	// Domain name or IP address of the third FSSO collector agent.
	Server3 *string `pulumi:"server3"`
	// Domain name or IP address of the fourth FSSO collector agent.
	Server4 *string `pulumi:"server4"`
	// Domain name or IP address of the fifth FSSO collector agent.
	Server5 *string `pulumi:"server5"`
	// Source IP for communications to FSSO agent.
	SourceIp *string `pulumi:"sourceIp"`
	// IPv6 source for communications to FSSO agent.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// Enable/disable use of SSL. Valid values: `enable`, `disable`.
	Ssl *string `pulumi:"ssl"`
	// Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
	SslServerHostIpCheck *string `pulumi:"sslServerHostIpCheck"`
	// Trusted server certificate or CA certificate.
	SslTrustedCert *string `pulumi:"sslTrustedCert"`
	// Server type.
	Type *string `pulumi:"type"`
	// LDAP server to get user information.
	UserInfoServer *string `pulumi:"userInfoServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserFssoState struct {
	// Interval in minutes within to fetch groups from FSSO server, or unset to disable.
	GroupPollInterval pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
	LdapPoll pulumi.StringPtrInput
	// Filter used to fetch groups.
	LdapPollFilter pulumi.StringPtrInput
	// Interval in minutes within to fetch groups from LDAP server.
	LdapPollInterval pulumi.IntPtrInput
	// LDAP server to get group information.
	LdapServer pulumi.StringPtrInput
	// Interval in minutes to keep logons after FSSO server down.
	LogonTimeout pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Password of the first FSSO collector agent.
	Password pulumi.StringPtrInput
	// Password of the second FSSO collector agent.
	Password2 pulumi.StringPtrInput
	// Password of the third FSSO collector agent.
	Password3 pulumi.StringPtrInput
	// Password of the fourth FSSO collector agent.
	Password4 pulumi.StringPtrInput
	// Password of the fifth FSSO collector agent.
	Password5 pulumi.StringPtrInput
	// Port of the first FSSO collector agent.
	Port pulumi.IntPtrInput
	// Port of the second FSSO collector agent.
	Port2 pulumi.IntPtrInput
	// Port of the third FSSO collector agent.
	Port3 pulumi.IntPtrInput
	// Port of the fourth FSSO collector agent.
	Port4 pulumi.IntPtrInput
	// Port of the fifth FSSO collector agent.
	Port5 pulumi.IntPtrInput
	// Domain name or IP address of the first FSSO collector agent.
	Server pulumi.StringPtrInput
	// Domain name or IP address of the second FSSO collector agent.
	Server2 pulumi.StringPtrInput
	// Domain name or IP address of the third FSSO collector agent.
	Server3 pulumi.StringPtrInput
	// Domain name or IP address of the fourth FSSO collector agent.
	Server4 pulumi.StringPtrInput
	// Domain name or IP address of the fifth FSSO collector agent.
	Server5 pulumi.StringPtrInput
	// Source IP for communications to FSSO agent.
	SourceIp pulumi.StringPtrInput
	// IPv6 source for communications to FSSO agent.
	SourceIp6 pulumi.StringPtrInput
	// Enable/disable use of SSL. Valid values: `enable`, `disable`.
	Ssl pulumi.StringPtrInput
	// Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
	SslServerHostIpCheck pulumi.StringPtrInput
	// Trusted server certificate or CA certificate.
	SslTrustedCert pulumi.StringPtrInput
	// Server type.
	Type pulumi.StringPtrInput
	// LDAP server to get user information.
	UserInfoServer pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserFssoState) ElementType() reflect.Type {
	return reflect.TypeOf((*userFssoState)(nil)).Elem()
}

type userFssoArgs struct {
	// Interval in minutes within to fetch groups from FSSO server, or unset to disable.
	GroupPollInterval *int `pulumi:"groupPollInterval"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
	LdapPoll *string `pulumi:"ldapPoll"`
	// Filter used to fetch groups.
	LdapPollFilter *string `pulumi:"ldapPollFilter"`
	// Interval in minutes within to fetch groups from LDAP server.
	LdapPollInterval *int `pulumi:"ldapPollInterval"`
	// LDAP server to get group information.
	LdapServer *string `pulumi:"ldapServer"`
	// Interval in minutes to keep logons after FSSO server down.
	LogonTimeout *int `pulumi:"logonTimeout"`
	// Name.
	Name *string `pulumi:"name"`
	// Password of the first FSSO collector agent.
	Password *string `pulumi:"password"`
	// Password of the second FSSO collector agent.
	Password2 *string `pulumi:"password2"`
	// Password of the third FSSO collector agent.
	Password3 *string `pulumi:"password3"`
	// Password of the fourth FSSO collector agent.
	Password4 *string `pulumi:"password4"`
	// Password of the fifth FSSO collector agent.
	Password5 *string `pulumi:"password5"`
	// Port of the first FSSO collector agent.
	Port *int `pulumi:"port"`
	// Port of the second FSSO collector agent.
	Port2 *int `pulumi:"port2"`
	// Port of the third FSSO collector agent.
	Port3 *int `pulumi:"port3"`
	// Port of the fourth FSSO collector agent.
	Port4 *int `pulumi:"port4"`
	// Port of the fifth FSSO collector agent.
	Port5 *int `pulumi:"port5"`
	// Domain name or IP address of the first FSSO collector agent.
	Server string `pulumi:"server"`
	// Domain name or IP address of the second FSSO collector agent.
	Server2 *string `pulumi:"server2"`
	// Domain name or IP address of the third FSSO collector agent.
	Server3 *string `pulumi:"server3"`
	// Domain name or IP address of the fourth FSSO collector agent.
	Server4 *string `pulumi:"server4"`
	// Domain name or IP address of the fifth FSSO collector agent.
	Server5 *string `pulumi:"server5"`
	// Source IP for communications to FSSO agent.
	SourceIp *string `pulumi:"sourceIp"`
	// IPv6 source for communications to FSSO agent.
	SourceIp6 *string `pulumi:"sourceIp6"`
	// Enable/disable use of SSL. Valid values: `enable`, `disable`.
	Ssl *string `pulumi:"ssl"`
	// Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
	SslServerHostIpCheck *string `pulumi:"sslServerHostIpCheck"`
	// Trusted server certificate or CA certificate.
	SslTrustedCert *string `pulumi:"sslTrustedCert"`
	// Server type.
	Type *string `pulumi:"type"`
	// LDAP server to get user information.
	UserInfoServer *string `pulumi:"userInfoServer"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserFsso resource.
type UserFssoArgs struct {
	// Interval in minutes within to fetch groups from FSSO server, or unset to disable.
	GroupPollInterval pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
	LdapPoll pulumi.StringPtrInput
	// Filter used to fetch groups.
	LdapPollFilter pulumi.StringPtrInput
	// Interval in minutes within to fetch groups from LDAP server.
	LdapPollInterval pulumi.IntPtrInput
	// LDAP server to get group information.
	LdapServer pulumi.StringPtrInput
	// Interval in minutes to keep logons after FSSO server down.
	LogonTimeout pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Password of the first FSSO collector agent.
	Password pulumi.StringPtrInput
	// Password of the second FSSO collector agent.
	Password2 pulumi.StringPtrInput
	// Password of the third FSSO collector agent.
	Password3 pulumi.StringPtrInput
	// Password of the fourth FSSO collector agent.
	Password4 pulumi.StringPtrInput
	// Password of the fifth FSSO collector agent.
	Password5 pulumi.StringPtrInput
	// Port of the first FSSO collector agent.
	Port pulumi.IntPtrInput
	// Port of the second FSSO collector agent.
	Port2 pulumi.IntPtrInput
	// Port of the third FSSO collector agent.
	Port3 pulumi.IntPtrInput
	// Port of the fourth FSSO collector agent.
	Port4 pulumi.IntPtrInput
	// Port of the fifth FSSO collector agent.
	Port5 pulumi.IntPtrInput
	// Domain name or IP address of the first FSSO collector agent.
	Server pulumi.StringInput
	// Domain name or IP address of the second FSSO collector agent.
	Server2 pulumi.StringPtrInput
	// Domain name or IP address of the third FSSO collector agent.
	Server3 pulumi.StringPtrInput
	// Domain name or IP address of the fourth FSSO collector agent.
	Server4 pulumi.StringPtrInput
	// Domain name or IP address of the fifth FSSO collector agent.
	Server5 pulumi.StringPtrInput
	// Source IP for communications to FSSO agent.
	SourceIp pulumi.StringPtrInput
	// IPv6 source for communications to FSSO agent.
	SourceIp6 pulumi.StringPtrInput
	// Enable/disable use of SSL. Valid values: `enable`, `disable`.
	Ssl pulumi.StringPtrInput
	// Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
	SslServerHostIpCheck pulumi.StringPtrInput
	// Trusted server certificate or CA certificate.
	SslTrustedCert pulumi.StringPtrInput
	// Server type.
	Type pulumi.StringPtrInput
	// LDAP server to get user information.
	UserInfoServer pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserFssoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userFssoArgs)(nil)).Elem()
}

type UserFssoInput interface {
	pulumi.Input

	ToUserFssoOutput() UserFssoOutput
	ToUserFssoOutputWithContext(ctx context.Context) UserFssoOutput
}

func (*UserFsso) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFsso)(nil)).Elem()
}

func (i *UserFsso) ToUserFssoOutput() UserFssoOutput {
	return i.ToUserFssoOutputWithContext(context.Background())
}

func (i *UserFsso) ToUserFssoOutputWithContext(ctx context.Context) UserFssoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssoOutput)
}

// UserFssoArrayInput is an input type that accepts UserFssoArray and UserFssoArrayOutput values.
// You can construct a concrete instance of `UserFssoArrayInput` via:
//
//          UserFssoArray{ UserFssoArgs{...} }
type UserFssoArrayInput interface {
	pulumi.Input

	ToUserFssoArrayOutput() UserFssoArrayOutput
	ToUserFssoArrayOutputWithContext(context.Context) UserFssoArrayOutput
}

type UserFssoArray []UserFssoInput

func (UserFssoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFsso)(nil)).Elem()
}

func (i UserFssoArray) ToUserFssoArrayOutput() UserFssoArrayOutput {
	return i.ToUserFssoArrayOutputWithContext(context.Background())
}

func (i UserFssoArray) ToUserFssoArrayOutputWithContext(ctx context.Context) UserFssoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssoArrayOutput)
}

// UserFssoMapInput is an input type that accepts UserFssoMap and UserFssoMapOutput values.
// You can construct a concrete instance of `UserFssoMapInput` via:
//
//          UserFssoMap{ "key": UserFssoArgs{...} }
type UserFssoMapInput interface {
	pulumi.Input

	ToUserFssoMapOutput() UserFssoMapOutput
	ToUserFssoMapOutputWithContext(context.Context) UserFssoMapOutput
}

type UserFssoMap map[string]UserFssoInput

func (UserFssoMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFsso)(nil)).Elem()
}

func (i UserFssoMap) ToUserFssoMapOutput() UserFssoMapOutput {
	return i.ToUserFssoMapOutputWithContext(context.Background())
}

func (i UserFssoMap) ToUserFssoMapOutputWithContext(ctx context.Context) UserFssoMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssoMapOutput)
}

type UserFssoOutput struct{ *pulumi.OutputState }

func (UserFssoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFsso)(nil)).Elem()
}

func (o UserFssoOutput) ToUserFssoOutput() UserFssoOutput {
	return o
}

func (o UserFssoOutput) ToUserFssoOutputWithContext(ctx context.Context) UserFssoOutput {
	return o
}

type UserFssoArrayOutput struct{ *pulumi.OutputState }

func (UserFssoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserFsso)(nil)).Elem()
}

func (o UserFssoArrayOutput) ToUserFssoArrayOutput() UserFssoArrayOutput {
	return o
}

func (o UserFssoArrayOutput) ToUserFssoArrayOutputWithContext(ctx context.Context) UserFssoArrayOutput {
	return o
}

func (o UserFssoArrayOutput) Index(i pulumi.IntInput) UserFssoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserFsso {
		return vs[0].([]*UserFsso)[vs[1].(int)]
	}).(UserFssoOutput)
}

type UserFssoMapOutput struct{ *pulumi.OutputState }

func (UserFssoMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserFsso)(nil)).Elem()
}

func (o UserFssoMapOutput) ToUserFssoMapOutput() UserFssoMapOutput {
	return o
}

func (o UserFssoMapOutput) ToUserFssoMapOutputWithContext(ctx context.Context) UserFssoMapOutput {
	return o
}

func (o UserFssoMapOutput) MapIndex(k pulumi.StringInput) UserFssoOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserFsso {
		return vs[0].(map[string]*UserFsso)[vs[1].(string)]
	}).(UserFssoOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserFssoInput)(nil)).Elem(), &UserFsso{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFssoArrayInput)(nil)).Elem(), UserFssoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserFssoMapInput)(nil)).Elem(), UserFssoMap{})
	pulumi.RegisterOutputType(UserFssoOutput{})
	pulumi.RegisterOutputType(UserFssoArrayOutput{})
	pulumi.RegisterOutputType(UserFssoMapOutput{})
}
