// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserDomainController struct {
	pulumi.CustomResourceState

	AdMode                pulumi.StringOutput                        `pulumi:"adMode"`
	AdldsDn               pulumi.StringOutput                        `pulumi:"adldsDn"`
	AdldsIp6              pulumi.StringOutput                        `pulumi:"adldsIp6"`
	AdldsIpAddress        pulumi.StringOutput                        `pulumi:"adldsIpAddress"`
	AdldsPort             pulumi.IntOutput                           `pulumi:"adldsPort"`
	DnsSrvLookup          pulumi.StringOutput                        `pulumi:"dnsSrvLookup"`
	DomainName            pulumi.StringOutput                        `pulumi:"domainName"`
	DynamicSortSubtable   pulumi.StringPtrOutput                     `pulumi:"dynamicSortSubtable"`
	ExtraServers          UserDomainControllerExtraServerArrayOutput `pulumi:"extraServers"`
	Hostname              pulumi.StringOutput                        `pulumi:"hostname"`
	Interface             pulumi.StringOutput                        `pulumi:"interface"`
	InterfaceSelectMethod pulumi.StringOutput                        `pulumi:"interfaceSelectMethod"`
	Ip6                   pulumi.StringOutput                        `pulumi:"ip6"`
	IpAddress             pulumi.StringOutput                        `pulumi:"ipAddress"`
	LdapServer            pulumi.StringOutput                        `pulumi:"ldapServer"`
	Name                  pulumi.StringOutput                        `pulumi:"name"`
	Password              pulumi.StringPtrOutput                     `pulumi:"password"`
	Port                  pulumi.IntOutput                           `pulumi:"port"`
	ReplicationPort       pulumi.IntOutput                           `pulumi:"replicationPort"`
	SourceIp6             pulumi.StringOutput                        `pulumi:"sourceIp6"`
	SourceIpAddress       pulumi.StringOutput                        `pulumi:"sourceIpAddress"`
	SourcePort            pulumi.IntOutput                           `pulumi:"sourcePort"`
	Username              pulumi.StringOutput                        `pulumi:"username"`
	Vdomparam             pulumi.StringPtrOutput                     `pulumi:"vdomparam"`
}

// NewUserDomainController registers a new resource with the given unique name, arguments, and options.
func NewUserDomainController(ctx *pulumi.Context,
	name string, args *UserDomainControllerArgs, opts ...pulumi.ResourceOption) (*UserDomainController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	if args.LdapServer == nil {
		return nil, errors.New("invalid value for required argument 'LdapServer'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserDomainController
	err := ctx.RegisterResource("fortios:index/userDomainController:UserDomainController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserDomainController gets an existing UserDomainController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserDomainController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserDomainControllerState, opts ...pulumi.ResourceOption) (*UserDomainController, error) {
	var resource UserDomainController
	err := ctx.ReadResource("fortios:index/userDomainController:UserDomainController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserDomainController resources.
type userDomainControllerState struct {
	AdMode                *string                           `pulumi:"adMode"`
	AdldsDn               *string                           `pulumi:"adldsDn"`
	AdldsIp6              *string                           `pulumi:"adldsIp6"`
	AdldsIpAddress        *string                           `pulumi:"adldsIpAddress"`
	AdldsPort             *int                              `pulumi:"adldsPort"`
	DnsSrvLookup          *string                           `pulumi:"dnsSrvLookup"`
	DomainName            *string                           `pulumi:"domainName"`
	DynamicSortSubtable   *string                           `pulumi:"dynamicSortSubtable"`
	ExtraServers          []UserDomainControllerExtraServer `pulumi:"extraServers"`
	Hostname              *string                           `pulumi:"hostname"`
	Interface             *string                           `pulumi:"interface"`
	InterfaceSelectMethod *string                           `pulumi:"interfaceSelectMethod"`
	Ip6                   *string                           `pulumi:"ip6"`
	IpAddress             *string                           `pulumi:"ipAddress"`
	LdapServer            *string                           `pulumi:"ldapServer"`
	Name                  *string                           `pulumi:"name"`
	Password              *string                           `pulumi:"password"`
	Port                  *int                              `pulumi:"port"`
	ReplicationPort       *int                              `pulumi:"replicationPort"`
	SourceIp6             *string                           `pulumi:"sourceIp6"`
	SourceIpAddress       *string                           `pulumi:"sourceIpAddress"`
	SourcePort            *int                              `pulumi:"sourcePort"`
	Username              *string                           `pulumi:"username"`
	Vdomparam             *string                           `pulumi:"vdomparam"`
}

type UserDomainControllerState struct {
	AdMode                pulumi.StringPtrInput
	AdldsDn               pulumi.StringPtrInput
	AdldsIp6              pulumi.StringPtrInput
	AdldsIpAddress        pulumi.StringPtrInput
	AdldsPort             pulumi.IntPtrInput
	DnsSrvLookup          pulumi.StringPtrInput
	DomainName            pulumi.StringPtrInput
	DynamicSortSubtable   pulumi.StringPtrInput
	ExtraServers          UserDomainControllerExtraServerArrayInput
	Hostname              pulumi.StringPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	Ip6                   pulumi.StringPtrInput
	IpAddress             pulumi.StringPtrInput
	LdapServer            pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	Password              pulumi.StringPtrInput
	Port                  pulumi.IntPtrInput
	ReplicationPort       pulumi.IntPtrInput
	SourceIp6             pulumi.StringPtrInput
	SourceIpAddress       pulumi.StringPtrInput
	SourcePort            pulumi.IntPtrInput
	Username              pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (UserDomainControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*userDomainControllerState)(nil)).Elem()
}

type userDomainControllerArgs struct {
	AdMode                *string                           `pulumi:"adMode"`
	AdldsDn               *string                           `pulumi:"adldsDn"`
	AdldsIp6              *string                           `pulumi:"adldsIp6"`
	AdldsIpAddress        *string                           `pulumi:"adldsIpAddress"`
	AdldsPort             *int                              `pulumi:"adldsPort"`
	DnsSrvLookup          *string                           `pulumi:"dnsSrvLookup"`
	DomainName            *string                           `pulumi:"domainName"`
	DynamicSortSubtable   *string                           `pulumi:"dynamicSortSubtable"`
	ExtraServers          []UserDomainControllerExtraServer `pulumi:"extraServers"`
	Hostname              *string                           `pulumi:"hostname"`
	Interface             *string                           `pulumi:"interface"`
	InterfaceSelectMethod *string                           `pulumi:"interfaceSelectMethod"`
	Ip6                   *string                           `pulumi:"ip6"`
	IpAddress             string                            `pulumi:"ipAddress"`
	LdapServer            string                            `pulumi:"ldapServer"`
	Name                  *string                           `pulumi:"name"`
	Password              *string                           `pulumi:"password"`
	Port                  *int                              `pulumi:"port"`
	ReplicationPort       *int                              `pulumi:"replicationPort"`
	SourceIp6             *string                           `pulumi:"sourceIp6"`
	SourceIpAddress       *string                           `pulumi:"sourceIpAddress"`
	SourcePort            *int                              `pulumi:"sourcePort"`
	Username              *string                           `pulumi:"username"`
	Vdomparam             *string                           `pulumi:"vdomparam"`
}

// The set of arguments for constructing a UserDomainController resource.
type UserDomainControllerArgs struct {
	AdMode                pulumi.StringPtrInput
	AdldsDn               pulumi.StringPtrInput
	AdldsIp6              pulumi.StringPtrInput
	AdldsIpAddress        pulumi.StringPtrInput
	AdldsPort             pulumi.IntPtrInput
	DnsSrvLookup          pulumi.StringPtrInput
	DomainName            pulumi.StringPtrInput
	DynamicSortSubtable   pulumi.StringPtrInput
	ExtraServers          UserDomainControllerExtraServerArrayInput
	Hostname              pulumi.StringPtrInput
	Interface             pulumi.StringPtrInput
	InterfaceSelectMethod pulumi.StringPtrInput
	Ip6                   pulumi.StringPtrInput
	IpAddress             pulumi.StringInput
	LdapServer            pulumi.StringInput
	Name                  pulumi.StringPtrInput
	Password              pulumi.StringPtrInput
	Port                  pulumi.IntPtrInput
	ReplicationPort       pulumi.IntPtrInput
	SourceIp6             pulumi.StringPtrInput
	SourceIpAddress       pulumi.StringPtrInput
	SourcePort            pulumi.IntPtrInput
	Username              pulumi.StringPtrInput
	Vdomparam             pulumi.StringPtrInput
}

func (UserDomainControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userDomainControllerArgs)(nil)).Elem()
}

type UserDomainControllerInput interface {
	pulumi.Input

	ToUserDomainControllerOutput() UserDomainControllerOutput
	ToUserDomainControllerOutputWithContext(ctx context.Context) UserDomainControllerOutput
}

func (*UserDomainController) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDomainController)(nil)).Elem()
}

func (i *UserDomainController) ToUserDomainControllerOutput() UserDomainControllerOutput {
	return i.ToUserDomainControllerOutputWithContext(context.Background())
}

func (i *UserDomainController) ToUserDomainControllerOutputWithContext(ctx context.Context) UserDomainControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDomainControllerOutput)
}

// UserDomainControllerArrayInput is an input type that accepts UserDomainControllerArray and UserDomainControllerArrayOutput values.
// You can construct a concrete instance of `UserDomainControllerArrayInput` via:
//
//	UserDomainControllerArray{ UserDomainControllerArgs{...} }
type UserDomainControllerArrayInput interface {
	pulumi.Input

	ToUserDomainControllerArrayOutput() UserDomainControllerArrayOutput
	ToUserDomainControllerArrayOutputWithContext(context.Context) UserDomainControllerArrayOutput
}

type UserDomainControllerArray []UserDomainControllerInput

func (UserDomainControllerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDomainController)(nil)).Elem()
}

func (i UserDomainControllerArray) ToUserDomainControllerArrayOutput() UserDomainControllerArrayOutput {
	return i.ToUserDomainControllerArrayOutputWithContext(context.Background())
}

func (i UserDomainControllerArray) ToUserDomainControllerArrayOutputWithContext(ctx context.Context) UserDomainControllerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDomainControllerArrayOutput)
}

// UserDomainControllerMapInput is an input type that accepts UserDomainControllerMap and UserDomainControllerMapOutput values.
// You can construct a concrete instance of `UserDomainControllerMapInput` via:
//
//	UserDomainControllerMap{ "key": UserDomainControllerArgs{...} }
type UserDomainControllerMapInput interface {
	pulumi.Input

	ToUserDomainControllerMapOutput() UserDomainControllerMapOutput
	ToUserDomainControllerMapOutputWithContext(context.Context) UserDomainControllerMapOutput
}

type UserDomainControllerMap map[string]UserDomainControllerInput

func (UserDomainControllerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDomainController)(nil)).Elem()
}

func (i UserDomainControllerMap) ToUserDomainControllerMapOutput() UserDomainControllerMapOutput {
	return i.ToUserDomainControllerMapOutputWithContext(context.Background())
}

func (i UserDomainControllerMap) ToUserDomainControllerMapOutputWithContext(ctx context.Context) UserDomainControllerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserDomainControllerMapOutput)
}

type UserDomainControllerOutput struct{ *pulumi.OutputState }

func (UserDomainControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserDomainController)(nil)).Elem()
}

func (o UserDomainControllerOutput) ToUserDomainControllerOutput() UserDomainControllerOutput {
	return o
}

func (o UserDomainControllerOutput) ToUserDomainControllerOutputWithContext(ctx context.Context) UserDomainControllerOutput {
	return o
}

func (o UserDomainControllerOutput) AdMode() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.AdMode }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) AdldsDn() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.AdldsDn }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) AdldsIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.AdldsIp6 }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) AdldsIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.AdldsIpAddress }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) AdldsPort() pulumi.IntOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.IntOutput { return v.AdldsPort }).(pulumi.IntOutput)
}

func (o UserDomainControllerOutput) DnsSrvLookup() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.DnsSrvLookup }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o UserDomainControllerOutput) ExtraServers() UserDomainControllerExtraServerArrayOutput {
	return o.ApplyT(func(v *UserDomainController) UserDomainControllerExtraServerArrayOutput { return v.ExtraServers }).(UserDomainControllerExtraServerArrayOutput)
}

func (o UserDomainControllerOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) Ip6() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.Ip6 }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) LdapServer() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.LdapServer }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o UserDomainControllerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o UserDomainControllerOutput) ReplicationPort() pulumi.IntOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.IntOutput { return v.ReplicationPort }).(pulumi.IntOutput)
}

func (o UserDomainControllerOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.SourceIp6 }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) SourceIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.SourceIpAddress }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) SourcePort() pulumi.IntOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.IntOutput { return v.SourcePort }).(pulumi.IntOutput)
}

func (o UserDomainControllerOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func (o UserDomainControllerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserDomainController) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type UserDomainControllerArrayOutput struct{ *pulumi.OutputState }

func (UserDomainControllerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserDomainController)(nil)).Elem()
}

func (o UserDomainControllerArrayOutput) ToUserDomainControllerArrayOutput() UserDomainControllerArrayOutput {
	return o
}

func (o UserDomainControllerArrayOutput) ToUserDomainControllerArrayOutputWithContext(ctx context.Context) UserDomainControllerArrayOutput {
	return o
}

func (o UserDomainControllerArrayOutput) Index(i pulumi.IntInput) UserDomainControllerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserDomainController {
		return vs[0].([]*UserDomainController)[vs[1].(int)]
	}).(UserDomainControllerOutput)
}

type UserDomainControllerMapOutput struct{ *pulumi.OutputState }

func (UserDomainControllerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserDomainController)(nil)).Elem()
}

func (o UserDomainControllerMapOutput) ToUserDomainControllerMapOutput() UserDomainControllerMapOutput {
	return o
}

func (o UserDomainControllerMapOutput) ToUserDomainControllerMapOutputWithContext(ctx context.Context) UserDomainControllerMapOutput {
	return o
}

func (o UserDomainControllerMapOutput) MapIndex(k pulumi.StringInput) UserDomainControllerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserDomainController {
		return vs[0].(map[string]*UserDomainController)[vs[1].(string)]
	}).(UserDomainControllerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserDomainControllerInput)(nil)).Elem(), &UserDomainController{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDomainControllerArrayInput)(nil)).Elem(), UserDomainControllerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserDomainControllerMapInput)(nil)).Elem(), UserDomainControllerMap{})
	pulumi.RegisterOutputType(UserDomainControllerOutput{})
	pulumi.RegisterOutputType(UserDomainControllerArrayOutput{})
	pulumi.RegisterOutputType(UserDomainControllerMapOutput{})
}
