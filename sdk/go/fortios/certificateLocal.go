// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CertificateLocal struct {
	pulumi.CustomResourceState

	AcmeCaUrl                 pulumi.StringOutput    `pulumi:"acmeCaUrl"`
	AcmeDomain                pulumi.StringOutput    `pulumi:"acmeDomain"`
	AcmeEmail                 pulumi.StringOutput    `pulumi:"acmeEmail"`
	AcmeRenewWindow           pulumi.IntOutput       `pulumi:"acmeRenewWindow"`
	AcmeRsaKeySize            pulumi.IntOutput       `pulumi:"acmeRsaKeySize"`
	AutoRegenerateDays        pulumi.IntOutput       `pulumi:"autoRegenerateDays"`
	AutoRegenerateDaysWarning pulumi.IntOutput       `pulumi:"autoRegenerateDaysWarning"`
	CaIdentifier              pulumi.StringOutput    `pulumi:"caIdentifier"`
	Certificate               pulumi.StringOutput    `pulumi:"certificate"`
	CmpPath                   pulumi.StringOutput    `pulumi:"cmpPath"`
	CmpRegenerationMethod     pulumi.StringOutput    `pulumi:"cmpRegenerationMethod"`
	CmpServer                 pulumi.StringOutput    `pulumi:"cmpServer"`
	CmpServerCert             pulumi.StringOutput    `pulumi:"cmpServerCert"`
	Comments                  pulumi.StringOutput    `pulumi:"comments"`
	Csr                       pulumi.StringOutput    `pulumi:"csr"`
	EnrollProtocol            pulumi.StringOutput    `pulumi:"enrollProtocol"`
	EstCaId                   pulumi.StringOutput    `pulumi:"estCaId"`
	EstClientCert             pulumi.StringOutput    `pulumi:"estClientCert"`
	EstHttpPassword           pulumi.StringOutput    `pulumi:"estHttpPassword"`
	EstHttpUsername           pulumi.StringOutput    `pulumi:"estHttpUsername"`
	EstServer                 pulumi.StringOutput    `pulumi:"estServer"`
	EstServerCert             pulumi.StringOutput    `pulumi:"estServerCert"`
	EstSrpPassword            pulumi.StringOutput    `pulumi:"estSrpPassword"`
	EstSrpUsername            pulumi.StringOutput    `pulumi:"estSrpUsername"`
	IkeLocalid                pulumi.StringOutput    `pulumi:"ikeLocalid"`
	IkeLocalidType            pulumi.StringOutput    `pulumi:"ikeLocalidType"`
	LastUpdated               pulumi.IntOutput       `pulumi:"lastUpdated"`
	Name                      pulumi.StringOutput    `pulumi:"name"`
	NameEncoding              pulumi.StringOutput    `pulumi:"nameEncoding"`
	Password                  pulumi.StringPtrOutput `pulumi:"password"`
	PrivateKey                pulumi.StringOutput    `pulumi:"privateKey"`
	PrivateKeyRetain          pulumi.StringOutput    `pulumi:"privateKeyRetain"`
	Range                     pulumi.StringOutput    `pulumi:"range"`
	ScepPassword              pulumi.StringPtrOutput `pulumi:"scepPassword"`
	ScepUrl                   pulumi.StringOutput    `pulumi:"scepUrl"`
	Source                    pulumi.StringOutput    `pulumi:"source"`
	SourceIp                  pulumi.StringOutput    `pulumi:"sourceIp"`
	State                     pulumi.StringOutput    `pulumi:"state"`
	Vdomparam                 pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCertificateLocal registers a new resource with the given unique name, arguments, and options.
func NewCertificateLocal(ctx *pulumi.Context,
	name string, args *CertificateLocalArgs, opts ...pulumi.ResourceOption) (*CertificateLocal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateKey == nil {
		return nil, errors.New("invalid value for required argument 'PrivateKey'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringInput)
	}
	if args.ScepPassword != nil {
		args.ScepPassword = pulumi.ToSecret(args.ScepPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
		"privateKey",
		"scepPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CertificateLocal
	err := ctx.RegisterResource("fortios:index/certificateLocal:CertificateLocal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateLocal gets an existing CertificateLocal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateLocal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateLocalState, opts ...pulumi.ResourceOption) (*CertificateLocal, error) {
	var resource CertificateLocal
	err := ctx.ReadResource("fortios:index/certificateLocal:CertificateLocal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateLocal resources.
type certificateLocalState struct {
	AcmeCaUrl                 *string `pulumi:"acmeCaUrl"`
	AcmeDomain                *string `pulumi:"acmeDomain"`
	AcmeEmail                 *string `pulumi:"acmeEmail"`
	AcmeRenewWindow           *int    `pulumi:"acmeRenewWindow"`
	AcmeRsaKeySize            *int    `pulumi:"acmeRsaKeySize"`
	AutoRegenerateDays        *int    `pulumi:"autoRegenerateDays"`
	AutoRegenerateDaysWarning *int    `pulumi:"autoRegenerateDaysWarning"`
	CaIdentifier              *string `pulumi:"caIdentifier"`
	Certificate               *string `pulumi:"certificate"`
	CmpPath                   *string `pulumi:"cmpPath"`
	CmpRegenerationMethod     *string `pulumi:"cmpRegenerationMethod"`
	CmpServer                 *string `pulumi:"cmpServer"`
	CmpServerCert             *string `pulumi:"cmpServerCert"`
	Comments                  *string `pulumi:"comments"`
	Csr                       *string `pulumi:"csr"`
	EnrollProtocol            *string `pulumi:"enrollProtocol"`
	EstCaId                   *string `pulumi:"estCaId"`
	EstClientCert             *string `pulumi:"estClientCert"`
	EstHttpPassword           *string `pulumi:"estHttpPassword"`
	EstHttpUsername           *string `pulumi:"estHttpUsername"`
	EstServer                 *string `pulumi:"estServer"`
	EstServerCert             *string `pulumi:"estServerCert"`
	EstSrpPassword            *string `pulumi:"estSrpPassword"`
	EstSrpUsername            *string `pulumi:"estSrpUsername"`
	IkeLocalid                *string `pulumi:"ikeLocalid"`
	IkeLocalidType            *string `pulumi:"ikeLocalidType"`
	LastUpdated               *int    `pulumi:"lastUpdated"`
	Name                      *string `pulumi:"name"`
	NameEncoding              *string `pulumi:"nameEncoding"`
	Password                  *string `pulumi:"password"`
	PrivateKey                *string `pulumi:"privateKey"`
	PrivateKeyRetain          *string `pulumi:"privateKeyRetain"`
	Range                     *string `pulumi:"range"`
	ScepPassword              *string `pulumi:"scepPassword"`
	ScepUrl                   *string `pulumi:"scepUrl"`
	Source                    *string `pulumi:"source"`
	SourceIp                  *string `pulumi:"sourceIp"`
	State                     *string `pulumi:"state"`
	Vdomparam                 *string `pulumi:"vdomparam"`
}

type CertificateLocalState struct {
	AcmeCaUrl                 pulumi.StringPtrInput
	AcmeDomain                pulumi.StringPtrInput
	AcmeEmail                 pulumi.StringPtrInput
	AcmeRenewWindow           pulumi.IntPtrInput
	AcmeRsaKeySize            pulumi.IntPtrInput
	AutoRegenerateDays        pulumi.IntPtrInput
	AutoRegenerateDaysWarning pulumi.IntPtrInput
	CaIdentifier              pulumi.StringPtrInput
	Certificate               pulumi.StringPtrInput
	CmpPath                   pulumi.StringPtrInput
	CmpRegenerationMethod     pulumi.StringPtrInput
	CmpServer                 pulumi.StringPtrInput
	CmpServerCert             pulumi.StringPtrInput
	Comments                  pulumi.StringPtrInput
	Csr                       pulumi.StringPtrInput
	EnrollProtocol            pulumi.StringPtrInput
	EstCaId                   pulumi.StringPtrInput
	EstClientCert             pulumi.StringPtrInput
	EstHttpPassword           pulumi.StringPtrInput
	EstHttpUsername           pulumi.StringPtrInput
	EstServer                 pulumi.StringPtrInput
	EstServerCert             pulumi.StringPtrInput
	EstSrpPassword            pulumi.StringPtrInput
	EstSrpUsername            pulumi.StringPtrInput
	IkeLocalid                pulumi.StringPtrInput
	IkeLocalidType            pulumi.StringPtrInput
	LastUpdated               pulumi.IntPtrInput
	Name                      pulumi.StringPtrInput
	NameEncoding              pulumi.StringPtrInput
	Password                  pulumi.StringPtrInput
	PrivateKey                pulumi.StringPtrInput
	PrivateKeyRetain          pulumi.StringPtrInput
	Range                     pulumi.StringPtrInput
	ScepPassword              pulumi.StringPtrInput
	ScepUrl                   pulumi.StringPtrInput
	Source                    pulumi.StringPtrInput
	SourceIp                  pulumi.StringPtrInput
	State                     pulumi.StringPtrInput
	Vdomparam                 pulumi.StringPtrInput
}

func (CertificateLocalState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateLocalState)(nil)).Elem()
}

type certificateLocalArgs struct {
	AcmeCaUrl                 *string `pulumi:"acmeCaUrl"`
	AcmeDomain                *string `pulumi:"acmeDomain"`
	AcmeEmail                 *string `pulumi:"acmeEmail"`
	AcmeRenewWindow           *int    `pulumi:"acmeRenewWindow"`
	AcmeRsaKeySize            *int    `pulumi:"acmeRsaKeySize"`
	AutoRegenerateDays        *int    `pulumi:"autoRegenerateDays"`
	AutoRegenerateDaysWarning *int    `pulumi:"autoRegenerateDaysWarning"`
	CaIdentifier              *string `pulumi:"caIdentifier"`
	Certificate               *string `pulumi:"certificate"`
	CmpPath                   *string `pulumi:"cmpPath"`
	CmpRegenerationMethod     *string `pulumi:"cmpRegenerationMethod"`
	CmpServer                 *string `pulumi:"cmpServer"`
	CmpServerCert             *string `pulumi:"cmpServerCert"`
	Comments                  *string `pulumi:"comments"`
	Csr                       *string `pulumi:"csr"`
	EnrollProtocol            *string `pulumi:"enrollProtocol"`
	EstCaId                   *string `pulumi:"estCaId"`
	EstClientCert             *string `pulumi:"estClientCert"`
	EstHttpPassword           *string `pulumi:"estHttpPassword"`
	EstHttpUsername           *string `pulumi:"estHttpUsername"`
	EstServer                 *string `pulumi:"estServer"`
	EstServerCert             *string `pulumi:"estServerCert"`
	EstSrpPassword            *string `pulumi:"estSrpPassword"`
	EstSrpUsername            *string `pulumi:"estSrpUsername"`
	IkeLocalid                *string `pulumi:"ikeLocalid"`
	IkeLocalidType            *string `pulumi:"ikeLocalidType"`
	LastUpdated               *int    `pulumi:"lastUpdated"`
	Name                      *string `pulumi:"name"`
	NameEncoding              *string `pulumi:"nameEncoding"`
	Password                  *string `pulumi:"password"`
	PrivateKey                string  `pulumi:"privateKey"`
	PrivateKeyRetain          *string `pulumi:"privateKeyRetain"`
	Range                     *string `pulumi:"range"`
	ScepPassword              *string `pulumi:"scepPassword"`
	ScepUrl                   *string `pulumi:"scepUrl"`
	Source                    *string `pulumi:"source"`
	SourceIp                  *string `pulumi:"sourceIp"`
	State                     *string `pulumi:"state"`
	Vdomparam                 *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a CertificateLocal resource.
type CertificateLocalArgs struct {
	AcmeCaUrl                 pulumi.StringPtrInput
	AcmeDomain                pulumi.StringPtrInput
	AcmeEmail                 pulumi.StringPtrInput
	AcmeRenewWindow           pulumi.IntPtrInput
	AcmeRsaKeySize            pulumi.IntPtrInput
	AutoRegenerateDays        pulumi.IntPtrInput
	AutoRegenerateDaysWarning pulumi.IntPtrInput
	CaIdentifier              pulumi.StringPtrInput
	Certificate               pulumi.StringPtrInput
	CmpPath                   pulumi.StringPtrInput
	CmpRegenerationMethod     pulumi.StringPtrInput
	CmpServer                 pulumi.StringPtrInput
	CmpServerCert             pulumi.StringPtrInput
	Comments                  pulumi.StringPtrInput
	Csr                       pulumi.StringPtrInput
	EnrollProtocol            pulumi.StringPtrInput
	EstCaId                   pulumi.StringPtrInput
	EstClientCert             pulumi.StringPtrInput
	EstHttpPassword           pulumi.StringPtrInput
	EstHttpUsername           pulumi.StringPtrInput
	EstServer                 pulumi.StringPtrInput
	EstServerCert             pulumi.StringPtrInput
	EstSrpPassword            pulumi.StringPtrInput
	EstSrpUsername            pulumi.StringPtrInput
	IkeLocalid                pulumi.StringPtrInput
	IkeLocalidType            pulumi.StringPtrInput
	LastUpdated               pulumi.IntPtrInput
	Name                      pulumi.StringPtrInput
	NameEncoding              pulumi.StringPtrInput
	Password                  pulumi.StringPtrInput
	PrivateKey                pulumi.StringInput
	PrivateKeyRetain          pulumi.StringPtrInput
	Range                     pulumi.StringPtrInput
	ScepPassword              pulumi.StringPtrInput
	ScepUrl                   pulumi.StringPtrInput
	Source                    pulumi.StringPtrInput
	SourceIp                  pulumi.StringPtrInput
	State                     pulumi.StringPtrInput
	Vdomparam                 pulumi.StringPtrInput
}

func (CertificateLocalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateLocalArgs)(nil)).Elem()
}

type CertificateLocalInput interface {
	pulumi.Input

	ToCertificateLocalOutput() CertificateLocalOutput
	ToCertificateLocalOutputWithContext(ctx context.Context) CertificateLocalOutput
}

func (*CertificateLocal) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateLocal)(nil)).Elem()
}

func (i *CertificateLocal) ToCertificateLocalOutput() CertificateLocalOutput {
	return i.ToCertificateLocalOutputWithContext(context.Background())
}

func (i *CertificateLocal) ToCertificateLocalOutputWithContext(ctx context.Context) CertificateLocalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateLocalOutput)
}

// CertificateLocalArrayInput is an input type that accepts CertificateLocalArray and CertificateLocalArrayOutput values.
// You can construct a concrete instance of `CertificateLocalArrayInput` via:
//
//	CertificateLocalArray{ CertificateLocalArgs{...} }
type CertificateLocalArrayInput interface {
	pulumi.Input

	ToCertificateLocalArrayOutput() CertificateLocalArrayOutput
	ToCertificateLocalArrayOutputWithContext(context.Context) CertificateLocalArrayOutput
}

type CertificateLocalArray []CertificateLocalInput

func (CertificateLocalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateLocal)(nil)).Elem()
}

func (i CertificateLocalArray) ToCertificateLocalArrayOutput() CertificateLocalArrayOutput {
	return i.ToCertificateLocalArrayOutputWithContext(context.Background())
}

func (i CertificateLocalArray) ToCertificateLocalArrayOutputWithContext(ctx context.Context) CertificateLocalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateLocalArrayOutput)
}

// CertificateLocalMapInput is an input type that accepts CertificateLocalMap and CertificateLocalMapOutput values.
// You can construct a concrete instance of `CertificateLocalMapInput` via:
//
//	CertificateLocalMap{ "key": CertificateLocalArgs{...} }
type CertificateLocalMapInput interface {
	pulumi.Input

	ToCertificateLocalMapOutput() CertificateLocalMapOutput
	ToCertificateLocalMapOutputWithContext(context.Context) CertificateLocalMapOutput
}

type CertificateLocalMap map[string]CertificateLocalInput

func (CertificateLocalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateLocal)(nil)).Elem()
}

func (i CertificateLocalMap) ToCertificateLocalMapOutput() CertificateLocalMapOutput {
	return i.ToCertificateLocalMapOutputWithContext(context.Background())
}

func (i CertificateLocalMap) ToCertificateLocalMapOutputWithContext(ctx context.Context) CertificateLocalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateLocalMapOutput)
}

type CertificateLocalOutput struct{ *pulumi.OutputState }

func (CertificateLocalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateLocal)(nil)).Elem()
}

func (o CertificateLocalOutput) ToCertificateLocalOutput() CertificateLocalOutput {
	return o
}

func (o CertificateLocalOutput) ToCertificateLocalOutputWithContext(ctx context.Context) CertificateLocalOutput {
	return o
}

func (o CertificateLocalOutput) AcmeCaUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.AcmeCaUrl }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) AcmeDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.AcmeDomain }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) AcmeEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.AcmeEmail }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) AcmeRenewWindow() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.IntOutput { return v.AcmeRenewWindow }).(pulumi.IntOutput)
}

func (o CertificateLocalOutput) AcmeRsaKeySize() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.IntOutput { return v.AcmeRsaKeySize }).(pulumi.IntOutput)
}

func (o CertificateLocalOutput) AutoRegenerateDays() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.IntOutput { return v.AutoRegenerateDays }).(pulumi.IntOutput)
}

func (o CertificateLocalOutput) AutoRegenerateDaysWarning() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.IntOutput { return v.AutoRegenerateDaysWarning }).(pulumi.IntOutput)
}

func (o CertificateLocalOutput) CaIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.CaIdentifier }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) CmpPath() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.CmpPath }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) CmpRegenerationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.CmpRegenerationMethod }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) CmpServer() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.CmpServer }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) CmpServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.CmpServerCert }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) Csr() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.Csr }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EnrollProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EnrollProtocol }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EstCaId() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EstCaId }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EstClientCert() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EstClientCert }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EstHttpPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EstHttpPassword }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EstHttpUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EstHttpUsername }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EstServer() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EstServer }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EstServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EstServerCert }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EstSrpPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EstSrpPassword }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) EstSrpUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.EstSrpUsername }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) IkeLocalid() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.IkeLocalid }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) IkeLocalidType() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.IkeLocalidType }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) LastUpdated() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.IntOutput { return v.LastUpdated }).(pulumi.IntOutput)
}

func (o CertificateLocalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) NameEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.NameEncoding }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o CertificateLocalOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) PrivateKeyRetain() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.PrivateKeyRetain }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) Range() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.Range }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) ScepPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringPtrOutput { return v.ScepPassword }).(pulumi.StringPtrOutput)
}

func (o CertificateLocalOutput) ScepUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.ScepUrl }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o CertificateLocalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateLocal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CertificateLocalArrayOutput struct{ *pulumi.OutputState }

func (CertificateLocalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateLocal)(nil)).Elem()
}

func (o CertificateLocalArrayOutput) ToCertificateLocalArrayOutput() CertificateLocalArrayOutput {
	return o
}

func (o CertificateLocalArrayOutput) ToCertificateLocalArrayOutputWithContext(ctx context.Context) CertificateLocalArrayOutput {
	return o
}

func (o CertificateLocalArrayOutput) Index(i pulumi.IntInput) CertificateLocalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateLocal {
		return vs[0].([]*CertificateLocal)[vs[1].(int)]
	}).(CertificateLocalOutput)
}

type CertificateLocalMapOutput struct{ *pulumi.OutputState }

func (CertificateLocalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateLocal)(nil)).Elem()
}

func (o CertificateLocalMapOutput) ToCertificateLocalMapOutput() CertificateLocalMapOutput {
	return o
}

func (o CertificateLocalMapOutput) ToCertificateLocalMapOutputWithContext(ctx context.Context) CertificateLocalMapOutput {
	return o
}

func (o CertificateLocalMapOutput) MapIndex(k pulumi.StringInput) CertificateLocalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateLocal {
		return vs[0].(map[string]*CertificateLocal)[vs[1].(string)]
	}).(CertificateLocalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateLocalInput)(nil)).Elem(), &CertificateLocal{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateLocalArrayInput)(nil)).Elem(), CertificateLocalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateLocalMapInput)(nil)).Elem(), CertificateLocalMap{})
	pulumi.RegisterOutputType(CertificateLocalOutput{})
	pulumi.RegisterOutputType(CertificateLocalArrayOutput{})
	pulumi.RegisterOutputType(CertificateLocalMapOutput{})
}
