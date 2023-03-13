// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthenticationScheme struct {
	pulumi.CustomResourceState

	DomainController    pulumi.StringOutput                         `pulumi:"domainController"`
	DynamicSortSubtable pulumi.StringPtrOutput                      `pulumi:"dynamicSortSubtable"`
	FssoAgentForNtlm    pulumi.StringOutput                         `pulumi:"fssoAgentForNtlm"`
	FssoGuest           pulumi.StringOutput                         `pulumi:"fssoGuest"`
	KerberosKeytab      pulumi.StringOutput                         `pulumi:"kerberosKeytab"`
	Method              pulumi.StringOutput                         `pulumi:"method"`
	Name                pulumi.StringOutput                         `pulumi:"name"`
	NegotiateNtlm       pulumi.StringOutput                         `pulumi:"negotiateNtlm"`
	RequireTfa          pulumi.StringOutput                         `pulumi:"requireTfa"`
	SamlServer          pulumi.StringOutput                         `pulumi:"samlServer"`
	SamlTimeout         pulumi.IntOutput                            `pulumi:"samlTimeout"`
	SshCa               pulumi.StringOutput                         `pulumi:"sshCa"`
	UserCert            pulumi.StringOutput                         `pulumi:"userCert"`
	UserDatabases       AuthenticationSchemeUserDatabaseArrayOutput `pulumi:"userDatabases"`
	Vdomparam           pulumi.StringPtrOutput                      `pulumi:"vdomparam"`
}

// NewAuthenticationScheme registers a new resource with the given unique name, arguments, and options.
func NewAuthenticationScheme(ctx *pulumi.Context,
	name string, args *AuthenticationSchemeArgs, opts ...pulumi.ResourceOption) (*AuthenticationScheme, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Method == nil {
		return nil, errors.New("invalid value for required argument 'Method'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AuthenticationScheme
	err := ctx.RegisterResource("fortios:index/authenticationScheme:AuthenticationScheme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthenticationScheme gets an existing AuthenticationScheme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthenticationScheme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthenticationSchemeState, opts ...pulumi.ResourceOption) (*AuthenticationScheme, error) {
	var resource AuthenticationScheme
	err := ctx.ReadResource("fortios:index/authenticationScheme:AuthenticationScheme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthenticationScheme resources.
type authenticationSchemeState struct {
	DomainController    *string                            `pulumi:"domainController"`
	DynamicSortSubtable *string                            `pulumi:"dynamicSortSubtable"`
	FssoAgentForNtlm    *string                            `pulumi:"fssoAgentForNtlm"`
	FssoGuest           *string                            `pulumi:"fssoGuest"`
	KerberosKeytab      *string                            `pulumi:"kerberosKeytab"`
	Method              *string                            `pulumi:"method"`
	Name                *string                            `pulumi:"name"`
	NegotiateNtlm       *string                            `pulumi:"negotiateNtlm"`
	RequireTfa          *string                            `pulumi:"requireTfa"`
	SamlServer          *string                            `pulumi:"samlServer"`
	SamlTimeout         *int                               `pulumi:"samlTimeout"`
	SshCa               *string                            `pulumi:"sshCa"`
	UserCert            *string                            `pulumi:"userCert"`
	UserDatabases       []AuthenticationSchemeUserDatabase `pulumi:"userDatabases"`
	Vdomparam           *string                            `pulumi:"vdomparam"`
}

type AuthenticationSchemeState struct {
	DomainController    pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	FssoAgentForNtlm    pulumi.StringPtrInput
	FssoGuest           pulumi.StringPtrInput
	KerberosKeytab      pulumi.StringPtrInput
	Method              pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	NegotiateNtlm       pulumi.StringPtrInput
	RequireTfa          pulumi.StringPtrInput
	SamlServer          pulumi.StringPtrInput
	SamlTimeout         pulumi.IntPtrInput
	SshCa               pulumi.StringPtrInput
	UserCert            pulumi.StringPtrInput
	UserDatabases       AuthenticationSchemeUserDatabaseArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (AuthenticationSchemeState) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationSchemeState)(nil)).Elem()
}

type authenticationSchemeArgs struct {
	DomainController    *string                            `pulumi:"domainController"`
	DynamicSortSubtable *string                            `pulumi:"dynamicSortSubtable"`
	FssoAgentForNtlm    *string                            `pulumi:"fssoAgentForNtlm"`
	FssoGuest           *string                            `pulumi:"fssoGuest"`
	KerberosKeytab      *string                            `pulumi:"kerberosKeytab"`
	Method              string                             `pulumi:"method"`
	Name                *string                            `pulumi:"name"`
	NegotiateNtlm       *string                            `pulumi:"negotiateNtlm"`
	RequireTfa          *string                            `pulumi:"requireTfa"`
	SamlServer          *string                            `pulumi:"samlServer"`
	SamlTimeout         *int                               `pulumi:"samlTimeout"`
	SshCa               *string                            `pulumi:"sshCa"`
	UserCert            *string                            `pulumi:"userCert"`
	UserDatabases       []AuthenticationSchemeUserDatabase `pulumi:"userDatabases"`
	Vdomparam           *string                            `pulumi:"vdomparam"`
}

// The set of arguments for constructing a AuthenticationScheme resource.
type AuthenticationSchemeArgs struct {
	DomainController    pulumi.StringPtrInput
	DynamicSortSubtable pulumi.StringPtrInput
	FssoAgentForNtlm    pulumi.StringPtrInput
	FssoGuest           pulumi.StringPtrInput
	KerberosKeytab      pulumi.StringPtrInput
	Method              pulumi.StringInput
	Name                pulumi.StringPtrInput
	NegotiateNtlm       pulumi.StringPtrInput
	RequireTfa          pulumi.StringPtrInput
	SamlServer          pulumi.StringPtrInput
	SamlTimeout         pulumi.IntPtrInput
	SshCa               pulumi.StringPtrInput
	UserCert            pulumi.StringPtrInput
	UserDatabases       AuthenticationSchemeUserDatabaseArrayInput
	Vdomparam           pulumi.StringPtrInput
}

func (AuthenticationSchemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationSchemeArgs)(nil)).Elem()
}

type AuthenticationSchemeInput interface {
	pulumi.Input

	ToAuthenticationSchemeOutput() AuthenticationSchemeOutput
	ToAuthenticationSchemeOutputWithContext(ctx context.Context) AuthenticationSchemeOutput
}

func (*AuthenticationScheme) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationScheme)(nil)).Elem()
}

func (i *AuthenticationScheme) ToAuthenticationSchemeOutput() AuthenticationSchemeOutput {
	return i.ToAuthenticationSchemeOutputWithContext(context.Background())
}

func (i *AuthenticationScheme) ToAuthenticationSchemeOutputWithContext(ctx context.Context) AuthenticationSchemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSchemeOutput)
}

// AuthenticationSchemeArrayInput is an input type that accepts AuthenticationSchemeArray and AuthenticationSchemeArrayOutput values.
// You can construct a concrete instance of `AuthenticationSchemeArrayInput` via:
//
//	AuthenticationSchemeArray{ AuthenticationSchemeArgs{...} }
type AuthenticationSchemeArrayInput interface {
	pulumi.Input

	ToAuthenticationSchemeArrayOutput() AuthenticationSchemeArrayOutput
	ToAuthenticationSchemeArrayOutputWithContext(context.Context) AuthenticationSchemeArrayOutput
}

type AuthenticationSchemeArray []AuthenticationSchemeInput

func (AuthenticationSchemeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationScheme)(nil)).Elem()
}

func (i AuthenticationSchemeArray) ToAuthenticationSchemeArrayOutput() AuthenticationSchemeArrayOutput {
	return i.ToAuthenticationSchemeArrayOutputWithContext(context.Background())
}

func (i AuthenticationSchemeArray) ToAuthenticationSchemeArrayOutputWithContext(ctx context.Context) AuthenticationSchemeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSchemeArrayOutput)
}

// AuthenticationSchemeMapInput is an input type that accepts AuthenticationSchemeMap and AuthenticationSchemeMapOutput values.
// You can construct a concrete instance of `AuthenticationSchemeMapInput` via:
//
//	AuthenticationSchemeMap{ "key": AuthenticationSchemeArgs{...} }
type AuthenticationSchemeMapInput interface {
	pulumi.Input

	ToAuthenticationSchemeMapOutput() AuthenticationSchemeMapOutput
	ToAuthenticationSchemeMapOutputWithContext(context.Context) AuthenticationSchemeMapOutput
}

type AuthenticationSchemeMap map[string]AuthenticationSchemeInput

func (AuthenticationSchemeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationScheme)(nil)).Elem()
}

func (i AuthenticationSchemeMap) ToAuthenticationSchemeMapOutput() AuthenticationSchemeMapOutput {
	return i.ToAuthenticationSchemeMapOutputWithContext(context.Background())
}

func (i AuthenticationSchemeMap) ToAuthenticationSchemeMapOutputWithContext(ctx context.Context) AuthenticationSchemeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationSchemeMapOutput)
}

type AuthenticationSchemeOutput struct{ *pulumi.OutputState }

func (AuthenticationSchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationScheme)(nil)).Elem()
}

func (o AuthenticationSchemeOutput) ToAuthenticationSchemeOutput() AuthenticationSchemeOutput {
	return o
}

func (o AuthenticationSchemeOutput) ToAuthenticationSchemeOutputWithContext(ctx context.Context) AuthenticationSchemeOutput {
	return o
}

func (o AuthenticationSchemeOutput) DomainController() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.DomainController }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o AuthenticationSchemeOutput) FssoAgentForNtlm() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.FssoAgentForNtlm }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) FssoGuest() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.FssoGuest }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) KerberosKeytab() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.KerberosKeytab }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.Method }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) NegotiateNtlm() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.NegotiateNtlm }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) RequireTfa() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.RequireTfa }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) SamlServer() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.SamlServer }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) SamlTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.IntOutput { return v.SamlTimeout }).(pulumi.IntOutput)
}

func (o AuthenticationSchemeOutput) SshCa() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.SshCa }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) UserCert() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringOutput { return v.UserCert }).(pulumi.StringOutput)
}

func (o AuthenticationSchemeOutput) UserDatabases() AuthenticationSchemeUserDatabaseArrayOutput {
	return o.ApplyT(func(v *AuthenticationScheme) AuthenticationSchemeUserDatabaseArrayOutput { return v.UserDatabases }).(AuthenticationSchemeUserDatabaseArrayOutput)
}

func (o AuthenticationSchemeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationScheme) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type AuthenticationSchemeArrayOutput struct{ *pulumi.OutputState }

func (AuthenticationSchemeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationScheme)(nil)).Elem()
}

func (o AuthenticationSchemeArrayOutput) ToAuthenticationSchemeArrayOutput() AuthenticationSchemeArrayOutput {
	return o
}

func (o AuthenticationSchemeArrayOutput) ToAuthenticationSchemeArrayOutputWithContext(ctx context.Context) AuthenticationSchemeArrayOutput {
	return o
}

func (o AuthenticationSchemeArrayOutput) Index(i pulumi.IntInput) AuthenticationSchemeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthenticationScheme {
		return vs[0].([]*AuthenticationScheme)[vs[1].(int)]
	}).(AuthenticationSchemeOutput)
}

type AuthenticationSchemeMapOutput struct{ *pulumi.OutputState }

func (AuthenticationSchemeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationScheme)(nil)).Elem()
}

func (o AuthenticationSchemeMapOutput) ToAuthenticationSchemeMapOutput() AuthenticationSchemeMapOutput {
	return o
}

func (o AuthenticationSchemeMapOutput) ToAuthenticationSchemeMapOutputWithContext(ctx context.Context) AuthenticationSchemeMapOutput {
	return o
}

func (o AuthenticationSchemeMapOutput) MapIndex(k pulumi.StringInput) AuthenticationSchemeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthenticationScheme {
		return vs[0].(map[string]*AuthenticationScheme)[vs[1].(string)]
	}).(AuthenticationSchemeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationSchemeInput)(nil)).Elem(), &AuthenticationScheme{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationSchemeArrayInput)(nil)).Elem(), AuthenticationSchemeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationSchemeMapInput)(nil)).Elem(), AuthenticationSchemeMap{})
	pulumi.RegisterOutputType(AuthenticationSchemeOutput{})
	pulumi.RegisterOutputType(AuthenticationSchemeArrayOutput{})
	pulumi.RegisterOutputType(AuthenticationSchemeMapOutput{})
}
