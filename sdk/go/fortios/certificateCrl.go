// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type CertificateCrl struct {
	pulumi.CustomResourceState

	Crl            pulumi.StringOutput    `pulumi:"crl"`
	HttpUrl        pulumi.StringOutput    `pulumi:"httpUrl"`
	LastUpdated    pulumi.IntOutput       `pulumi:"lastUpdated"`
	LdapPassword   pulumi.StringPtrOutput `pulumi:"ldapPassword"`
	LdapServer     pulumi.StringOutput    `pulumi:"ldapServer"`
	LdapUsername   pulumi.StringOutput    `pulumi:"ldapUsername"`
	Name           pulumi.StringOutput    `pulumi:"name"`
	Range          pulumi.StringOutput    `pulumi:"range"`
	ScepCert       pulumi.StringOutput    `pulumi:"scepCert"`
	ScepUrl        pulumi.StringOutput    `pulumi:"scepUrl"`
	Source         pulumi.StringOutput    `pulumi:"source"`
	SourceIp       pulumi.StringOutput    `pulumi:"sourceIp"`
	UpdateInterval pulumi.IntOutput       `pulumi:"updateInterval"`
	UpdateVdom     pulumi.StringOutput    `pulumi:"updateVdom"`
	Vdomparam      pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCertificateCrl registers a new resource with the given unique name, arguments, and options.
func NewCertificateCrl(ctx *pulumi.Context,
	name string, args *CertificateCrlArgs, opts ...pulumi.ResourceOption) (*CertificateCrl, error) {
	if args == nil {
		args = &CertificateCrlArgs{}
	}

	if args.LdapPassword != nil {
		args.LdapPassword = pulumi.ToSecret(args.LdapPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"ldapPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CertificateCrl
	err := ctx.RegisterResource("fortios:index/certificateCrl:CertificateCrl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateCrl gets an existing CertificateCrl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateCrl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateCrlState, opts ...pulumi.ResourceOption) (*CertificateCrl, error) {
	var resource CertificateCrl
	err := ctx.ReadResource("fortios:index/certificateCrl:CertificateCrl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateCrl resources.
type certificateCrlState struct {
	Crl            *string `pulumi:"crl"`
	HttpUrl        *string `pulumi:"httpUrl"`
	LastUpdated    *int    `pulumi:"lastUpdated"`
	LdapPassword   *string `pulumi:"ldapPassword"`
	LdapServer     *string `pulumi:"ldapServer"`
	LdapUsername   *string `pulumi:"ldapUsername"`
	Name           *string `pulumi:"name"`
	Range          *string `pulumi:"range"`
	ScepCert       *string `pulumi:"scepCert"`
	ScepUrl        *string `pulumi:"scepUrl"`
	Source         *string `pulumi:"source"`
	SourceIp       *string `pulumi:"sourceIp"`
	UpdateInterval *int    `pulumi:"updateInterval"`
	UpdateVdom     *string `pulumi:"updateVdom"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

type CertificateCrlState struct {
	Crl            pulumi.StringPtrInput
	HttpUrl        pulumi.StringPtrInput
	LastUpdated    pulumi.IntPtrInput
	LdapPassword   pulumi.StringPtrInput
	LdapServer     pulumi.StringPtrInput
	LdapUsername   pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	Range          pulumi.StringPtrInput
	ScepCert       pulumi.StringPtrInput
	ScepUrl        pulumi.StringPtrInput
	Source         pulumi.StringPtrInput
	SourceIp       pulumi.StringPtrInput
	UpdateInterval pulumi.IntPtrInput
	UpdateVdom     pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (CertificateCrlState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateCrlState)(nil)).Elem()
}

type certificateCrlArgs struct {
	Crl            *string `pulumi:"crl"`
	HttpUrl        *string `pulumi:"httpUrl"`
	LastUpdated    *int    `pulumi:"lastUpdated"`
	LdapPassword   *string `pulumi:"ldapPassword"`
	LdapServer     *string `pulumi:"ldapServer"`
	LdapUsername   *string `pulumi:"ldapUsername"`
	Name           *string `pulumi:"name"`
	Range          *string `pulumi:"range"`
	ScepCert       *string `pulumi:"scepCert"`
	ScepUrl        *string `pulumi:"scepUrl"`
	Source         *string `pulumi:"source"`
	SourceIp       *string `pulumi:"sourceIp"`
	UpdateInterval *int    `pulumi:"updateInterval"`
	UpdateVdom     *string `pulumi:"updateVdom"`
	Vdomparam      *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a CertificateCrl resource.
type CertificateCrlArgs struct {
	Crl            pulumi.StringPtrInput
	HttpUrl        pulumi.StringPtrInput
	LastUpdated    pulumi.IntPtrInput
	LdapPassword   pulumi.StringPtrInput
	LdapServer     pulumi.StringPtrInput
	LdapUsername   pulumi.StringPtrInput
	Name           pulumi.StringPtrInput
	Range          pulumi.StringPtrInput
	ScepCert       pulumi.StringPtrInput
	ScepUrl        pulumi.StringPtrInput
	Source         pulumi.StringPtrInput
	SourceIp       pulumi.StringPtrInput
	UpdateInterval pulumi.IntPtrInput
	UpdateVdom     pulumi.StringPtrInput
	Vdomparam      pulumi.StringPtrInput
}

func (CertificateCrlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateCrlArgs)(nil)).Elem()
}

type CertificateCrlInput interface {
	pulumi.Input

	ToCertificateCrlOutput() CertificateCrlOutput
	ToCertificateCrlOutputWithContext(ctx context.Context) CertificateCrlOutput
}

func (*CertificateCrl) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateCrl)(nil)).Elem()
}

func (i *CertificateCrl) ToCertificateCrlOutput() CertificateCrlOutput {
	return i.ToCertificateCrlOutputWithContext(context.Background())
}

func (i *CertificateCrl) ToCertificateCrlOutputWithContext(ctx context.Context) CertificateCrlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCrlOutput)
}

func (i *CertificateCrl) ToOutput(ctx context.Context) pulumix.Output[*CertificateCrl] {
	return pulumix.Output[*CertificateCrl]{
		OutputState: i.ToCertificateCrlOutputWithContext(ctx).OutputState,
	}
}

// CertificateCrlArrayInput is an input type that accepts CertificateCrlArray and CertificateCrlArrayOutput values.
// You can construct a concrete instance of `CertificateCrlArrayInput` via:
//
//	CertificateCrlArray{ CertificateCrlArgs{...} }
type CertificateCrlArrayInput interface {
	pulumi.Input

	ToCertificateCrlArrayOutput() CertificateCrlArrayOutput
	ToCertificateCrlArrayOutputWithContext(context.Context) CertificateCrlArrayOutput
}

type CertificateCrlArray []CertificateCrlInput

func (CertificateCrlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateCrl)(nil)).Elem()
}

func (i CertificateCrlArray) ToCertificateCrlArrayOutput() CertificateCrlArrayOutput {
	return i.ToCertificateCrlArrayOutputWithContext(context.Background())
}

func (i CertificateCrlArray) ToCertificateCrlArrayOutputWithContext(ctx context.Context) CertificateCrlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCrlArrayOutput)
}

func (i CertificateCrlArray) ToOutput(ctx context.Context) pulumix.Output[[]*CertificateCrl] {
	return pulumix.Output[[]*CertificateCrl]{
		OutputState: i.ToCertificateCrlArrayOutputWithContext(ctx).OutputState,
	}
}

// CertificateCrlMapInput is an input type that accepts CertificateCrlMap and CertificateCrlMapOutput values.
// You can construct a concrete instance of `CertificateCrlMapInput` via:
//
//	CertificateCrlMap{ "key": CertificateCrlArgs{...} }
type CertificateCrlMapInput interface {
	pulumi.Input

	ToCertificateCrlMapOutput() CertificateCrlMapOutput
	ToCertificateCrlMapOutputWithContext(context.Context) CertificateCrlMapOutput
}

type CertificateCrlMap map[string]CertificateCrlInput

func (CertificateCrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateCrl)(nil)).Elem()
}

func (i CertificateCrlMap) ToCertificateCrlMapOutput() CertificateCrlMapOutput {
	return i.ToCertificateCrlMapOutputWithContext(context.Background())
}

func (i CertificateCrlMap) ToCertificateCrlMapOutputWithContext(ctx context.Context) CertificateCrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCrlMapOutput)
}

func (i CertificateCrlMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CertificateCrl] {
	return pulumix.Output[map[string]*CertificateCrl]{
		OutputState: i.ToCertificateCrlMapOutputWithContext(ctx).OutputState,
	}
}

type CertificateCrlOutput struct{ *pulumi.OutputState }

func (CertificateCrlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateCrl)(nil)).Elem()
}

func (o CertificateCrlOutput) ToCertificateCrlOutput() CertificateCrlOutput {
	return o
}

func (o CertificateCrlOutput) ToCertificateCrlOutputWithContext(ctx context.Context) CertificateCrlOutput {
	return o
}

func (o CertificateCrlOutput) ToOutput(ctx context.Context) pulumix.Output[*CertificateCrl] {
	return pulumix.Output[*CertificateCrl]{
		OutputState: o.OutputState,
	}
}

func (o CertificateCrlOutput) Crl() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.Crl }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) HttpUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.HttpUrl }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) LastUpdated() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.IntOutput { return v.LastUpdated }).(pulumi.IntOutput)
}

func (o CertificateCrlOutput) LdapPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringPtrOutput { return v.LdapPassword }).(pulumi.StringPtrOutput)
}

func (o CertificateCrlOutput) LdapServer() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.LdapServer }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) LdapUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.LdapUsername }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) Range() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.Range }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) ScepCert() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.ScepCert }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) ScepUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.ScepUrl }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) UpdateInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.IntOutput { return v.UpdateInterval }).(pulumi.IntOutput)
}

func (o CertificateCrlOutput) UpdateVdom() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringOutput { return v.UpdateVdom }).(pulumi.StringOutput)
}

func (o CertificateCrlOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateCrl) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CertificateCrlArrayOutput struct{ *pulumi.OutputState }

func (CertificateCrlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateCrl)(nil)).Elem()
}

func (o CertificateCrlArrayOutput) ToCertificateCrlArrayOutput() CertificateCrlArrayOutput {
	return o
}

func (o CertificateCrlArrayOutput) ToCertificateCrlArrayOutputWithContext(ctx context.Context) CertificateCrlArrayOutput {
	return o
}

func (o CertificateCrlArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CertificateCrl] {
	return pulumix.Output[[]*CertificateCrl]{
		OutputState: o.OutputState,
	}
}

func (o CertificateCrlArrayOutput) Index(i pulumi.IntInput) CertificateCrlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateCrl {
		return vs[0].([]*CertificateCrl)[vs[1].(int)]
	}).(CertificateCrlOutput)
}

type CertificateCrlMapOutput struct{ *pulumi.OutputState }

func (CertificateCrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateCrl)(nil)).Elem()
}

func (o CertificateCrlMapOutput) ToCertificateCrlMapOutput() CertificateCrlMapOutput {
	return o
}

func (o CertificateCrlMapOutput) ToCertificateCrlMapOutputWithContext(ctx context.Context) CertificateCrlMapOutput {
	return o
}

func (o CertificateCrlMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CertificateCrl] {
	return pulumix.Output[map[string]*CertificateCrl]{
		OutputState: o.OutputState,
	}
}

func (o CertificateCrlMapOutput) MapIndex(k pulumi.StringInput) CertificateCrlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateCrl {
		return vs[0].(map[string]*CertificateCrl)[vs[1].(string)]
	}).(CertificateCrlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateCrlInput)(nil)).Elem(), &CertificateCrl{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateCrlArrayInput)(nil)).Elem(), CertificateCrlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateCrlMapInput)(nil)).Elem(), CertificateCrlMap{})
	pulumi.RegisterOutputType(CertificateCrlOutput{})
	pulumi.RegisterOutputType(CertificateCrlArrayOutput{})
	pulumi.RegisterOutputType(CertificateCrlMapOutput{})
}
