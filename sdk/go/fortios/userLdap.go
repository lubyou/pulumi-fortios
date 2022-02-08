// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure LDAP server entries.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewUserLdap(ctx, "trname", &fortios.UserLdapArgs{
// 			AccountKeyFilter:      pulumi.String(fmt.Sprintf("%v%v%v", "(&(userPrincipalName=", "%", "s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))")),
// 			AccountKeyProcessing:  pulumi.String("same"),
// 			Cnid:                  pulumi.String("cn"),
// 			Dn:                    pulumi.String("EIWNCIEW"),
// 			GroupMemberCheck:      pulumi.String("user-attr"),
// 			GroupObjectFilter:     pulumi.String("(&(objectcategory=group)(member=*))"),
// 			MemberAttr:            pulumi.String("memberOf"),
// 			PasswordExpiryWarning: pulumi.String("disable"),
// 			PasswordRenewal:       pulumi.String("disable"),
// 			Port:                  pulumi.Int(389),
// 			Secure:                pulumi.String("disable"),
// 			Server:                pulumi.String("1.1.1.1"),
// 			ServerIdentityCheck:   pulumi.String("disable"),
// 			SourceIp:              pulumi.String("0.0.0.0"),
// 			SslMinProtoVersion:    pulumi.String("default"),
// 			Type:                  pulumi.String("simple"),
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
// User Ldap can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/userLdap:UserLdap labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type UserLdap struct {
	pulumi.CustomResourceState

	// Account key filter, using the UPN as the search filter.
	AccountKeyFilter pulumi.StringOutput `pulumi:"accountKeyFilter"`
	// Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
	AccountKeyProcessing pulumi.StringOutput `pulumi:"accountKeyProcessing"`
	// Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
	Antiphish pulumi.StringOutput `pulumi:"antiphish"`
	// CA certificate name.
	CaCert pulumi.StringOutput `pulumi:"caCert"`
	// Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
	Cnid pulumi.StringOutput `pulumi:"cnid"`
	// Distinguished name used to look up entries on the LDAP server.
	Dn pulumi.StringOutput `pulumi:"dn"`
	// Filter used for group matching.
	GroupFilter pulumi.StringOutput `pulumi:"groupFilter"`
	// Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
	GroupMemberCheck pulumi.StringOutput `pulumi:"groupMemberCheck"`
	// Filter used for group searching.
	GroupObjectFilter pulumi.StringOutput `pulumi:"groupObjectFilter"`
	// Search base used for group searching.
	GroupSearchBase pulumi.StringOutput `pulumi:"groupSearchBase"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Name of attribute from which to get group membership.
	MemberAttr pulumi.StringOutput `pulumi:"memberAttr"`
	// LDAP server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
	ObtainUserInfo pulumi.StringOutput `pulumi:"obtainUserInfo"`
	// Password for initial binding.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Name of attribute to get password hash.
	PasswordAttr pulumi.StringOutput `pulumi:"passwordAttr"`
	// Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
	PasswordExpiryWarning pulumi.StringOutput `pulumi:"passwordExpiryWarning"`
	// Enable/disable online password renewal. Valid values: `enable`, `disable`.
	PasswordRenewal pulumi.StringOutput `pulumi:"passwordRenewal"`
	// Port to be used for communication with the LDAP server (default = 389).
	Port pulumi.IntOutput `pulumi:"port"`
	// Search type. Valid values: `recursive`.
	SearchType pulumi.StringOutput `pulumi:"searchType"`
	// Secondary LDAP server CN domain name or IP.
	SecondaryServer pulumi.StringOutput `pulumi:"secondaryServer"`
	// Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
	Secure pulumi.StringOutput `pulumi:"secure"`
	// LDAP server CN domain name or IP.
	Server pulumi.StringOutput `pulumi:"server"`
	// Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
	ServerIdentityCheck pulumi.StringOutput `pulumi:"serverIdentityCheck"`
	// Source IP for communications to LDAP server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Source port to be used for communication with the LDAP server.
	SourcePort pulumi.IntOutput `pulumi:"sourcePort"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringOutput `pulumi:"sslMinProtoVersion"`
	// Tertiary LDAP server CN domain name or IP.
	TertiaryServer pulumi.StringOutput `pulumi:"tertiaryServer"`
	// Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
	TwoFactor pulumi.StringOutput `pulumi:"twoFactor"`
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication pulumi.StringOutput `pulumi:"twoFactorAuthentication"`
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification pulumi.StringOutput `pulumi:"twoFactorNotification"`
	// Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
	Type pulumi.StringOutput `pulumi:"type"`
	// MS Exchange server from which to fetch user information.
	UserInfoExchangeServer pulumi.StringOutput `pulumi:"userInfoExchangeServer"`
	// Username (full DN) for initial binding.
	Username pulumi.StringOutput `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserLdap registers a new resource with the given unique name, arguments, and options.
func NewUserLdap(ctx *pulumi.Context,
	name string, args *UserLdapArgs, opts ...pulumi.ResourceOption) (*UserLdap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dn == nil {
		return nil, errors.New("invalid value for required argument 'Dn'")
	}
	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserLdap
	err := ctx.RegisterResource("fortios:index/userLdap:UserLdap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserLdap gets an existing UserLdap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserLdap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserLdapState, opts ...pulumi.ResourceOption) (*UserLdap, error) {
	var resource UserLdap
	err := ctx.ReadResource("fortios:index/userLdap:UserLdap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserLdap resources.
type userLdapState struct {
	// Account key filter, using the UPN as the search filter.
	AccountKeyFilter *string `pulumi:"accountKeyFilter"`
	// Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
	AccountKeyProcessing *string `pulumi:"accountKeyProcessing"`
	// Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
	Antiphish *string `pulumi:"antiphish"`
	// CA certificate name.
	CaCert *string `pulumi:"caCert"`
	// Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
	Cnid *string `pulumi:"cnid"`
	// Distinguished name used to look up entries on the LDAP server.
	Dn *string `pulumi:"dn"`
	// Filter used for group matching.
	GroupFilter *string `pulumi:"groupFilter"`
	// Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
	GroupMemberCheck *string `pulumi:"groupMemberCheck"`
	// Filter used for group searching.
	GroupObjectFilter *string `pulumi:"groupObjectFilter"`
	// Search base used for group searching.
	GroupSearchBase *string `pulumi:"groupSearchBase"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Name of attribute from which to get group membership.
	MemberAttr *string `pulumi:"memberAttr"`
	// LDAP server entry name.
	Name *string `pulumi:"name"`
	// Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
	ObtainUserInfo *string `pulumi:"obtainUserInfo"`
	// Password for initial binding.
	Password *string `pulumi:"password"`
	// Name of attribute to get password hash.
	PasswordAttr *string `pulumi:"passwordAttr"`
	// Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
	PasswordExpiryWarning *string `pulumi:"passwordExpiryWarning"`
	// Enable/disable online password renewal. Valid values: `enable`, `disable`.
	PasswordRenewal *string `pulumi:"passwordRenewal"`
	// Port to be used for communication with the LDAP server (default = 389).
	Port *int `pulumi:"port"`
	// Search type. Valid values: `recursive`.
	SearchType *string `pulumi:"searchType"`
	// Secondary LDAP server CN domain name or IP.
	SecondaryServer *string `pulumi:"secondaryServer"`
	// Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
	Secure *string `pulumi:"secure"`
	// LDAP server CN domain name or IP.
	Server *string `pulumi:"server"`
	// Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
	ServerIdentityCheck *string `pulumi:"serverIdentityCheck"`
	// Source IP for communications to LDAP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Source port to be used for communication with the LDAP server.
	SourcePort *int `pulumi:"sourcePort"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Tertiary LDAP server CN domain name or IP.
	TertiaryServer *string `pulumi:"tertiaryServer"`
	// Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
	TwoFactor *string `pulumi:"twoFactor"`
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication *string `pulumi:"twoFactorAuthentication"`
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification *string `pulumi:"twoFactorNotification"`
	// Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
	Type *string `pulumi:"type"`
	// MS Exchange server from which to fetch user information.
	UserInfoExchangeServer *string `pulumi:"userInfoExchangeServer"`
	// Username (full DN) for initial binding.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserLdapState struct {
	// Account key filter, using the UPN as the search filter.
	AccountKeyFilter pulumi.StringPtrInput
	// Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
	AccountKeyProcessing pulumi.StringPtrInput
	// Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
	Antiphish pulumi.StringPtrInput
	// CA certificate name.
	CaCert pulumi.StringPtrInput
	// Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
	Cnid pulumi.StringPtrInput
	// Distinguished name used to look up entries on the LDAP server.
	Dn pulumi.StringPtrInput
	// Filter used for group matching.
	GroupFilter pulumi.StringPtrInput
	// Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
	GroupMemberCheck pulumi.StringPtrInput
	// Filter used for group searching.
	GroupObjectFilter pulumi.StringPtrInput
	// Search base used for group searching.
	GroupSearchBase pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Name of attribute from which to get group membership.
	MemberAttr pulumi.StringPtrInput
	// LDAP server entry name.
	Name pulumi.StringPtrInput
	// Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
	ObtainUserInfo pulumi.StringPtrInput
	// Password for initial binding.
	Password pulumi.StringPtrInput
	// Name of attribute to get password hash.
	PasswordAttr pulumi.StringPtrInput
	// Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
	PasswordExpiryWarning pulumi.StringPtrInput
	// Enable/disable online password renewal. Valid values: `enable`, `disable`.
	PasswordRenewal pulumi.StringPtrInput
	// Port to be used for communication with the LDAP server (default = 389).
	Port pulumi.IntPtrInput
	// Search type. Valid values: `recursive`.
	SearchType pulumi.StringPtrInput
	// Secondary LDAP server CN domain name or IP.
	SecondaryServer pulumi.StringPtrInput
	// Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
	Secure pulumi.StringPtrInput
	// LDAP server CN domain name or IP.
	Server pulumi.StringPtrInput
	// Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
	ServerIdentityCheck pulumi.StringPtrInput
	// Source IP for communications to LDAP server.
	SourceIp pulumi.StringPtrInput
	// Source port to be used for communication with the LDAP server.
	SourcePort pulumi.IntPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Tertiary LDAP server CN domain name or IP.
	TertiaryServer pulumi.StringPtrInput
	// Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
	TwoFactor pulumi.StringPtrInput
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication pulumi.StringPtrInput
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification pulumi.StringPtrInput
	// Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
	Type pulumi.StringPtrInput
	// MS Exchange server from which to fetch user information.
	UserInfoExchangeServer pulumi.StringPtrInput
	// Username (full DN) for initial binding.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserLdapState) ElementType() reflect.Type {
	return reflect.TypeOf((*userLdapState)(nil)).Elem()
}

type userLdapArgs struct {
	// Account key filter, using the UPN as the search filter.
	AccountKeyFilter *string `pulumi:"accountKeyFilter"`
	// Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
	AccountKeyProcessing *string `pulumi:"accountKeyProcessing"`
	// Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
	Antiphish *string `pulumi:"antiphish"`
	// CA certificate name.
	CaCert *string `pulumi:"caCert"`
	// Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
	Cnid *string `pulumi:"cnid"`
	// Distinguished name used to look up entries on the LDAP server.
	Dn string `pulumi:"dn"`
	// Filter used for group matching.
	GroupFilter *string `pulumi:"groupFilter"`
	// Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
	GroupMemberCheck *string `pulumi:"groupMemberCheck"`
	// Filter used for group searching.
	GroupObjectFilter *string `pulumi:"groupObjectFilter"`
	// Search base used for group searching.
	GroupSearchBase *string `pulumi:"groupSearchBase"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Name of attribute from which to get group membership.
	MemberAttr *string `pulumi:"memberAttr"`
	// LDAP server entry name.
	Name *string `pulumi:"name"`
	// Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
	ObtainUserInfo *string `pulumi:"obtainUserInfo"`
	// Password for initial binding.
	Password *string `pulumi:"password"`
	// Name of attribute to get password hash.
	PasswordAttr *string `pulumi:"passwordAttr"`
	// Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
	PasswordExpiryWarning *string `pulumi:"passwordExpiryWarning"`
	// Enable/disable online password renewal. Valid values: `enable`, `disable`.
	PasswordRenewal *string `pulumi:"passwordRenewal"`
	// Port to be used for communication with the LDAP server (default = 389).
	Port *int `pulumi:"port"`
	// Search type. Valid values: `recursive`.
	SearchType *string `pulumi:"searchType"`
	// Secondary LDAP server CN domain name or IP.
	SecondaryServer *string `pulumi:"secondaryServer"`
	// Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
	Secure *string `pulumi:"secure"`
	// LDAP server CN domain name or IP.
	Server string `pulumi:"server"`
	// Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
	ServerIdentityCheck *string `pulumi:"serverIdentityCheck"`
	// Source IP for communications to LDAP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Source port to be used for communication with the LDAP server.
	SourcePort *int `pulumi:"sourcePort"`
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion *string `pulumi:"sslMinProtoVersion"`
	// Tertiary LDAP server CN domain name or IP.
	TertiaryServer *string `pulumi:"tertiaryServer"`
	// Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
	TwoFactor *string `pulumi:"twoFactor"`
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication *string `pulumi:"twoFactorAuthentication"`
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification *string `pulumi:"twoFactorNotification"`
	// Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
	Type *string `pulumi:"type"`
	// MS Exchange server from which to fetch user information.
	UserInfoExchangeServer *string `pulumi:"userInfoExchangeServer"`
	// Username (full DN) for initial binding.
	Username *string `pulumi:"username"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserLdap resource.
type UserLdapArgs struct {
	// Account key filter, using the UPN as the search filter.
	AccountKeyFilter pulumi.StringPtrInput
	// Account key processing operation, either keep or strip domain string of UPN in the token. Valid values: `same`, `strip`.
	AccountKeyProcessing pulumi.StringPtrInput
	// Enable/disable AntiPhishing credential backend. Valid values: `enable`, `disable`.
	Antiphish pulumi.StringPtrInput
	// CA certificate name.
	CaCert pulumi.StringPtrInput
	// Common name identifier for the LDAP server. The common name identifier for most LDAP servers is "cn".
	Cnid pulumi.StringPtrInput
	// Distinguished name used to look up entries on the LDAP server.
	Dn pulumi.StringInput
	// Filter used for group matching.
	GroupFilter pulumi.StringPtrInput
	// Group member checking methods. Valid values: `user-attr`, `group-object`, `posix-group-object`.
	GroupMemberCheck pulumi.StringPtrInput
	// Filter used for group searching.
	GroupObjectFilter pulumi.StringPtrInput
	// Search base used for group searching.
	GroupSearchBase pulumi.StringPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Name of attribute from which to get group membership.
	MemberAttr pulumi.StringPtrInput
	// LDAP server entry name.
	Name pulumi.StringPtrInput
	// Enable/disable obtaining of user information. Valid values: `enable`, `disable`.
	ObtainUserInfo pulumi.StringPtrInput
	// Password for initial binding.
	Password pulumi.StringPtrInput
	// Name of attribute to get password hash.
	PasswordAttr pulumi.StringPtrInput
	// Enable/disable password expiry warnings. Valid values: `enable`, `disable`.
	PasswordExpiryWarning pulumi.StringPtrInput
	// Enable/disable online password renewal. Valid values: `enable`, `disable`.
	PasswordRenewal pulumi.StringPtrInput
	// Port to be used for communication with the LDAP server (default = 389).
	Port pulumi.IntPtrInput
	// Search type. Valid values: `recursive`.
	SearchType pulumi.StringPtrInput
	// Secondary LDAP server CN domain name or IP.
	SecondaryServer pulumi.StringPtrInput
	// Port to be used for authentication. Valid values: `disable`, `starttls`, `ldaps`.
	Secure pulumi.StringPtrInput
	// LDAP server CN domain name or IP.
	Server pulumi.StringInput
	// Enable/disable LDAP server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
	ServerIdentityCheck pulumi.StringPtrInput
	// Source IP for communications to LDAP server.
	SourceIp pulumi.StringPtrInput
	// Source port to be used for communication with the LDAP server.
	SourcePort pulumi.IntPtrInput
	// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
	SslMinProtoVersion pulumi.StringPtrInput
	// Tertiary LDAP server CN domain name or IP.
	TertiaryServer pulumi.StringPtrInput
	// Enable/disable two-factor authentication. Valid values: `disable`, `fortitoken-cloud`.
	TwoFactor pulumi.StringPtrInput
	// Authentication method by FortiToken Cloud. Valid values: `fortitoken`, `email`, `sms`.
	TwoFactorAuthentication pulumi.StringPtrInput
	// Notification method for user activation by FortiToken Cloud. Valid values: `email`, `sms`.
	TwoFactorNotification pulumi.StringPtrInput
	// Authentication type for LDAP searches. Valid values: `simple`, `anonymous`, `regular`.
	Type pulumi.StringPtrInput
	// MS Exchange server from which to fetch user information.
	UserInfoExchangeServer pulumi.StringPtrInput
	// Username (full DN) for initial binding.
	Username pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserLdapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userLdapArgs)(nil)).Elem()
}

type UserLdapInput interface {
	pulumi.Input

	ToUserLdapOutput() UserLdapOutput
	ToUserLdapOutputWithContext(ctx context.Context) UserLdapOutput
}

func (*UserLdap) ElementType() reflect.Type {
	return reflect.TypeOf((**UserLdap)(nil)).Elem()
}

func (i *UserLdap) ToUserLdapOutput() UserLdapOutput {
	return i.ToUserLdapOutputWithContext(context.Background())
}

func (i *UserLdap) ToUserLdapOutputWithContext(ctx context.Context) UserLdapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLdapOutput)
}

// UserLdapArrayInput is an input type that accepts UserLdapArray and UserLdapArrayOutput values.
// You can construct a concrete instance of `UserLdapArrayInput` via:
//
//          UserLdapArray{ UserLdapArgs{...} }
type UserLdapArrayInput interface {
	pulumi.Input

	ToUserLdapArrayOutput() UserLdapArrayOutput
	ToUserLdapArrayOutputWithContext(context.Context) UserLdapArrayOutput
}

type UserLdapArray []UserLdapInput

func (UserLdapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserLdap)(nil)).Elem()
}

func (i UserLdapArray) ToUserLdapArrayOutput() UserLdapArrayOutput {
	return i.ToUserLdapArrayOutputWithContext(context.Background())
}

func (i UserLdapArray) ToUserLdapArrayOutputWithContext(ctx context.Context) UserLdapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLdapArrayOutput)
}

// UserLdapMapInput is an input type that accepts UserLdapMap and UserLdapMapOutput values.
// You can construct a concrete instance of `UserLdapMapInput` via:
//
//          UserLdapMap{ "key": UserLdapArgs{...} }
type UserLdapMapInput interface {
	pulumi.Input

	ToUserLdapMapOutput() UserLdapMapOutput
	ToUserLdapMapOutputWithContext(context.Context) UserLdapMapOutput
}

type UserLdapMap map[string]UserLdapInput

func (UserLdapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserLdap)(nil)).Elem()
}

func (i UserLdapMap) ToUserLdapMapOutput() UserLdapMapOutput {
	return i.ToUserLdapMapOutputWithContext(context.Background())
}

func (i UserLdapMap) ToUserLdapMapOutputWithContext(ctx context.Context) UserLdapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLdapMapOutput)
}

type UserLdapOutput struct{ *pulumi.OutputState }

func (UserLdapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserLdap)(nil)).Elem()
}

func (o UserLdapOutput) ToUserLdapOutput() UserLdapOutput {
	return o
}

func (o UserLdapOutput) ToUserLdapOutputWithContext(ctx context.Context) UserLdapOutput {
	return o
}

type UserLdapArrayOutput struct{ *pulumi.OutputState }

func (UserLdapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserLdap)(nil)).Elem()
}

func (o UserLdapArrayOutput) ToUserLdapArrayOutput() UserLdapArrayOutput {
	return o
}

func (o UserLdapArrayOutput) ToUserLdapArrayOutputWithContext(ctx context.Context) UserLdapArrayOutput {
	return o
}

func (o UserLdapArrayOutput) Index(i pulumi.IntInput) UserLdapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserLdap {
		return vs[0].([]*UserLdap)[vs[1].(int)]
	}).(UserLdapOutput)
}

type UserLdapMapOutput struct{ *pulumi.OutputState }

func (UserLdapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserLdap)(nil)).Elem()
}

func (o UserLdapMapOutput) ToUserLdapMapOutput() UserLdapMapOutput {
	return o
}

func (o UserLdapMapOutput) ToUserLdapMapOutputWithContext(ctx context.Context) UserLdapMapOutput {
	return o
}

func (o UserLdapMapOutput) MapIndex(k pulumi.StringInput) UserLdapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserLdap {
		return vs[0].(map[string]*UserLdap)[vs[1].(string)]
	}).(UserLdapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserLdapInput)(nil)).Elem(), &UserLdap{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserLdapArrayInput)(nil)).Elem(), UserLdapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserLdapMapInput)(nil)).Elem(), UserLdapMap{})
	pulumi.RegisterOutputType(UserLdapOutput{})
	pulumi.RegisterOutputType(UserLdapArrayOutput{})
	pulumi.RegisterOutputType(UserLdapMapOutput{})
}
