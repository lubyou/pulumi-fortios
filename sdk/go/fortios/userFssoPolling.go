// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FSSO active directory servers for polling mode.
//
// ## Import
//
// User FssoPolling can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/userFssoPolling:UserFssoPolling labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type UserFssoPolling struct {
	pulumi.CustomResourceState

	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps UserFssoPollingAdgrpArrayOutput `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringOutput `pulumi:"defaultDomain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntOutput `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntOutput `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port pulumi.IntOutput `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringOutput `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringOutput `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringOutput `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User pulumi.StringOutput `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewUserFssoPolling registers a new resource with the given unique name, arguments, and options.
func NewUserFssoPolling(ctx *pulumi.Context,
	name string, args *UserFssoPollingArgs, opts ...pulumi.ResourceOption) (*UserFssoPolling, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapServer == nil {
		return nil, errors.New("invalid value for required argument 'LdapServer'")
	}
	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	var resource UserFssoPolling
	err := ctx.RegisterResource("fortios:index/userFssoPolling:UserFssoPolling", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserFssoPolling gets an existing UserFssoPolling resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserFssoPolling(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserFssoPollingState, opts ...pulumi.ResourceOption) (*UserFssoPolling, error) {
	var resource UserFssoPolling
	err := ctx.ReadResource("fortios:index/userFssoPolling:UserFssoPolling", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserFssoPolling resources.
type userFssoPollingState struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps []UserFssoPollingAdgrp `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid *int `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer *string `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory *int `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password *string `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency *int `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port *int `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server *string `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth *string `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 *string `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User *string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type UserFssoPollingState struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps UserFssoPollingAdgrpArrayInput
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Active Directory server ID.
	Fosid pulumi.IntPtrInput
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringPtrInput
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntPtrInput
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrInput
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntPtrInput
	// Port to communicate with this Active Directory server.
	Port pulumi.IntPtrInput
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringPtrInput
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringPtrInput
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringPtrInput
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User name required to log into this Active Directory server.
	User pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserFssoPollingState) ElementType() reflect.Type {
	return reflect.TypeOf((*userFssoPollingState)(nil)).Elem()
}

type userFssoPollingArgs struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps []UserFssoPollingAdgrp `pulumi:"adgrps"`
	// Default domain managed by this Active Directory server.
	DefaultDomain *string `pulumi:"defaultDomain"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Active Directory server ID.
	Fosid *int `pulumi:"fosid"`
	// LDAP server name used in LDAP connection strings.
	LdapServer string `pulumi:"ldapServer"`
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory *int `pulumi:"logonHistory"`
	// Password required to log into this Active Directory server
	Password *string `pulumi:"password"`
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency *int `pulumi:"pollingFrequency"`
	// Port to communicate with this Active Directory server.
	Port *int `pulumi:"port"`
	// Host name or IP address of the Active Directory server.
	Server string `pulumi:"server"`
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth *string `pulumi:"smbNtlmv1Auth"`
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 *string `pulumi:"smbv1"`
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// User name required to log into this Active Directory server.
	User string `pulumi:"user"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserFssoPolling resource.
type UserFssoPollingArgs struct {
	// LDAP Group Info. The structure of `adgrp` block is documented below.
	Adgrps UserFssoPollingAdgrpArrayInput
	// Default domain managed by this Active Directory server.
	DefaultDomain pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Active Directory server ID.
	Fosid pulumi.IntPtrInput
	// LDAP server name used in LDAP connection strings.
	LdapServer pulumi.StringInput
	// Number of hours of logon history to keep, 0 means keep all history.
	LogonHistory pulumi.IntPtrInput
	// Password required to log into this Active Directory server
	Password pulumi.StringPtrInput
	// Polling frequency (every 1 to 30 seconds).
	PollingFrequency pulumi.IntPtrInput
	// Port to communicate with this Active Directory server.
	Port pulumi.IntPtrInput
	// Host name or IP address of the Active Directory server.
	Server pulumi.StringInput
	// Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
	SmbNtlmv1Auth pulumi.StringPtrInput
	// Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
	Smbv1 pulumi.StringPtrInput
	// Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// User name required to log into this Active Directory server.
	User pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (UserFssoPollingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userFssoPollingArgs)(nil)).Elem()
}

type UserFssoPollingInput interface {
	pulumi.Input

	ToUserFssoPollingOutput() UserFssoPollingOutput
	ToUserFssoPollingOutputWithContext(ctx context.Context) UserFssoPollingOutput
}

func (*UserFssoPolling) ElementType() reflect.Type {
	return reflect.TypeOf((*UserFssoPolling)(nil))
}

func (i *UserFssoPolling) ToUserFssoPollingOutput() UserFssoPollingOutput {
	return i.ToUserFssoPollingOutputWithContext(context.Background())
}

func (i *UserFssoPolling) ToUserFssoPollingOutputWithContext(ctx context.Context) UserFssoPollingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssoPollingOutput)
}

func (i *UserFssoPolling) ToUserFssoPollingPtrOutput() UserFssoPollingPtrOutput {
	return i.ToUserFssoPollingPtrOutputWithContext(context.Background())
}

func (i *UserFssoPolling) ToUserFssoPollingPtrOutputWithContext(ctx context.Context) UserFssoPollingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssoPollingPtrOutput)
}

type UserFssoPollingPtrInput interface {
	pulumi.Input

	ToUserFssoPollingPtrOutput() UserFssoPollingPtrOutput
	ToUserFssoPollingPtrOutputWithContext(ctx context.Context) UserFssoPollingPtrOutput
}

type userFssoPollingPtrType UserFssoPollingArgs

func (*userFssoPollingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFssoPolling)(nil))
}

func (i *userFssoPollingPtrType) ToUserFssoPollingPtrOutput() UserFssoPollingPtrOutput {
	return i.ToUserFssoPollingPtrOutputWithContext(context.Background())
}

func (i *userFssoPollingPtrType) ToUserFssoPollingPtrOutputWithContext(ctx context.Context) UserFssoPollingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssoPollingPtrOutput)
}

// UserFssoPollingArrayInput is an input type that accepts UserFssoPollingArray and UserFssoPollingArrayOutput values.
// You can construct a concrete instance of `UserFssoPollingArrayInput` via:
//
//          UserFssoPollingArray{ UserFssoPollingArgs{...} }
type UserFssoPollingArrayInput interface {
	pulumi.Input

	ToUserFssoPollingArrayOutput() UserFssoPollingArrayOutput
	ToUserFssoPollingArrayOutputWithContext(context.Context) UserFssoPollingArrayOutput
}

type UserFssoPollingArray []UserFssoPollingInput

func (UserFssoPollingArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*UserFssoPolling)(nil))
}

func (i UserFssoPollingArray) ToUserFssoPollingArrayOutput() UserFssoPollingArrayOutput {
	return i.ToUserFssoPollingArrayOutputWithContext(context.Background())
}

func (i UserFssoPollingArray) ToUserFssoPollingArrayOutputWithContext(ctx context.Context) UserFssoPollingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssoPollingArrayOutput)
}

// UserFssoPollingMapInput is an input type that accepts UserFssoPollingMap and UserFssoPollingMapOutput values.
// You can construct a concrete instance of `UserFssoPollingMapInput` via:
//
//          UserFssoPollingMap{ "key": UserFssoPollingArgs{...} }
type UserFssoPollingMapInput interface {
	pulumi.Input

	ToUserFssoPollingMapOutput() UserFssoPollingMapOutput
	ToUserFssoPollingMapOutputWithContext(context.Context) UserFssoPollingMapOutput
}

type UserFssoPollingMap map[string]UserFssoPollingInput

func (UserFssoPollingMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*UserFssoPolling)(nil))
}

func (i UserFssoPollingMap) ToUserFssoPollingMapOutput() UserFssoPollingMapOutput {
	return i.ToUserFssoPollingMapOutputWithContext(context.Background())
}

func (i UserFssoPollingMap) ToUserFssoPollingMapOutputWithContext(ctx context.Context) UserFssoPollingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFssoPollingMapOutput)
}

type UserFssoPollingOutput struct {
	*pulumi.OutputState
}

func (UserFssoPollingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserFssoPolling)(nil))
}

func (o UserFssoPollingOutput) ToUserFssoPollingOutput() UserFssoPollingOutput {
	return o
}

func (o UserFssoPollingOutput) ToUserFssoPollingOutputWithContext(ctx context.Context) UserFssoPollingOutput {
	return o
}

func (o UserFssoPollingOutput) ToUserFssoPollingPtrOutput() UserFssoPollingPtrOutput {
	return o.ToUserFssoPollingPtrOutputWithContext(context.Background())
}

func (o UserFssoPollingOutput) ToUserFssoPollingPtrOutputWithContext(ctx context.Context) UserFssoPollingPtrOutput {
	return o.ApplyT(func(v UserFssoPolling) *UserFssoPolling {
		return &v
	}).(UserFssoPollingPtrOutput)
}

type UserFssoPollingPtrOutput struct {
	*pulumi.OutputState
}

func (UserFssoPollingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFssoPolling)(nil))
}

func (o UserFssoPollingPtrOutput) ToUserFssoPollingPtrOutput() UserFssoPollingPtrOutput {
	return o
}

func (o UserFssoPollingPtrOutput) ToUserFssoPollingPtrOutputWithContext(ctx context.Context) UserFssoPollingPtrOutput {
	return o
}

type UserFssoPollingArrayOutput struct{ *pulumi.OutputState }

func (UserFssoPollingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserFssoPolling)(nil))
}

func (o UserFssoPollingArrayOutput) ToUserFssoPollingArrayOutput() UserFssoPollingArrayOutput {
	return o
}

func (o UserFssoPollingArrayOutput) ToUserFssoPollingArrayOutputWithContext(ctx context.Context) UserFssoPollingArrayOutput {
	return o
}

func (o UserFssoPollingArrayOutput) Index(i pulumi.IntInput) UserFssoPollingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserFssoPolling {
		return vs[0].([]UserFssoPolling)[vs[1].(int)]
	}).(UserFssoPollingOutput)
}

type UserFssoPollingMapOutput struct{ *pulumi.OutputState }

func (UserFssoPollingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserFssoPolling)(nil))
}

func (o UserFssoPollingMapOutput) ToUserFssoPollingMapOutput() UserFssoPollingMapOutput {
	return o
}

func (o UserFssoPollingMapOutput) ToUserFssoPollingMapOutputWithContext(ctx context.Context) UserFssoPollingMapOutput {
	return o
}

func (o UserFssoPollingMapOutput) MapIndex(k pulumi.StringInput) UserFssoPollingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserFssoPolling {
		return vs[0].(map[string]UserFssoPolling)[vs[1].(string)]
	}).(UserFssoPollingOutput)
}

func init() {
	pulumi.RegisterOutputType(UserFssoPollingOutput{})
	pulumi.RegisterOutputType(UserFssoPollingPtrOutput{})
	pulumi.RegisterOutputType(UserFssoPollingArrayOutput{})
	pulumi.RegisterOutputType(UserFssoPollingMapOutput{})
}
