// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnCertificateLocal struct {
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

// NewVpnCertificateLocal registers a new resource with the given unique name, arguments, and options.
func NewVpnCertificateLocal(ctx *pulumi.Context,
	name string, args *VpnCertificateLocalArgs, opts ...pulumi.ResourceOption) (*VpnCertificateLocal, error) {
	if args == nil {
		args = &VpnCertificateLocalArgs{}
	}

	if args.Certificate != nil {
		args.Certificate = pulumi.ToSecret(args.Certificate).(pulumi.StringPtrInput)
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	if args.PrivateKey != nil {
		args.PrivateKey = pulumi.ToSecret(args.PrivateKey).(pulumi.StringPtrInput)
	}
	if args.ScepPassword != nil {
		args.ScepPassword = pulumi.ToSecret(args.ScepPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"certificate",
		"password",
		"privateKey",
		"scepPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnCertificateLocal
	err := ctx.RegisterResource("fortios:index/vpnCertificateLocal:VpnCertificateLocal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnCertificateLocal gets an existing VpnCertificateLocal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnCertificateLocal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnCertificateLocalState, opts ...pulumi.ResourceOption) (*VpnCertificateLocal, error) {
	var resource VpnCertificateLocal
	err := ctx.ReadResource("fortios:index/vpnCertificateLocal:VpnCertificateLocal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnCertificateLocal resources.
type vpnCertificateLocalState struct {
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

type VpnCertificateLocalState struct {
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

func (VpnCertificateLocalState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCertificateLocalState)(nil)).Elem()
}

type vpnCertificateLocalArgs struct {
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

// The set of arguments for constructing a VpnCertificateLocal resource.
type VpnCertificateLocalArgs struct {
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

func (VpnCertificateLocalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnCertificateLocalArgs)(nil)).Elem()
}

type VpnCertificateLocalInput interface {
	pulumi.Input

	ToVpnCertificateLocalOutput() VpnCertificateLocalOutput
	ToVpnCertificateLocalOutputWithContext(ctx context.Context) VpnCertificateLocalOutput
}

func (*VpnCertificateLocal) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCertificateLocal)(nil)).Elem()
}

func (i *VpnCertificateLocal) ToVpnCertificateLocalOutput() VpnCertificateLocalOutput {
	return i.ToVpnCertificateLocalOutputWithContext(context.Background())
}

func (i *VpnCertificateLocal) ToVpnCertificateLocalOutputWithContext(ctx context.Context) VpnCertificateLocalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateLocalOutput)
}

// VpnCertificateLocalArrayInput is an input type that accepts VpnCertificateLocalArray and VpnCertificateLocalArrayOutput values.
// You can construct a concrete instance of `VpnCertificateLocalArrayInput` via:
//
//	VpnCertificateLocalArray{ VpnCertificateLocalArgs{...} }
type VpnCertificateLocalArrayInput interface {
	pulumi.Input

	ToVpnCertificateLocalArrayOutput() VpnCertificateLocalArrayOutput
	ToVpnCertificateLocalArrayOutputWithContext(context.Context) VpnCertificateLocalArrayOutput
}

type VpnCertificateLocalArray []VpnCertificateLocalInput

func (VpnCertificateLocalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnCertificateLocal)(nil)).Elem()
}

func (i VpnCertificateLocalArray) ToVpnCertificateLocalArrayOutput() VpnCertificateLocalArrayOutput {
	return i.ToVpnCertificateLocalArrayOutputWithContext(context.Background())
}

func (i VpnCertificateLocalArray) ToVpnCertificateLocalArrayOutputWithContext(ctx context.Context) VpnCertificateLocalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateLocalArrayOutput)
}

// VpnCertificateLocalMapInput is an input type that accepts VpnCertificateLocalMap and VpnCertificateLocalMapOutput values.
// You can construct a concrete instance of `VpnCertificateLocalMapInput` via:
//
//	VpnCertificateLocalMap{ "key": VpnCertificateLocalArgs{...} }
type VpnCertificateLocalMapInput interface {
	pulumi.Input

	ToVpnCertificateLocalMapOutput() VpnCertificateLocalMapOutput
	ToVpnCertificateLocalMapOutputWithContext(context.Context) VpnCertificateLocalMapOutput
}

type VpnCertificateLocalMap map[string]VpnCertificateLocalInput

func (VpnCertificateLocalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnCertificateLocal)(nil)).Elem()
}

func (i VpnCertificateLocalMap) ToVpnCertificateLocalMapOutput() VpnCertificateLocalMapOutput {
	return i.ToVpnCertificateLocalMapOutputWithContext(context.Background())
}

func (i VpnCertificateLocalMap) ToVpnCertificateLocalMapOutputWithContext(ctx context.Context) VpnCertificateLocalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnCertificateLocalMapOutput)
}

type VpnCertificateLocalOutput struct{ *pulumi.OutputState }

func (VpnCertificateLocalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnCertificateLocal)(nil)).Elem()
}

func (o VpnCertificateLocalOutput) ToVpnCertificateLocalOutput() VpnCertificateLocalOutput {
	return o
}

func (o VpnCertificateLocalOutput) ToVpnCertificateLocalOutputWithContext(ctx context.Context) VpnCertificateLocalOutput {
	return o
}

func (o VpnCertificateLocalOutput) AcmeCaUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.AcmeCaUrl }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) AcmeDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.AcmeDomain }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) AcmeEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.AcmeEmail }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) AcmeRenewWindow() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.IntOutput { return v.AcmeRenewWindow }).(pulumi.IntOutput)
}

func (o VpnCertificateLocalOutput) AcmeRsaKeySize() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.IntOutput { return v.AcmeRsaKeySize }).(pulumi.IntOutput)
}

func (o VpnCertificateLocalOutput) AutoRegenerateDays() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.IntOutput { return v.AutoRegenerateDays }).(pulumi.IntOutput)
}

func (o VpnCertificateLocalOutput) AutoRegenerateDaysWarning() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.IntOutput { return v.AutoRegenerateDaysWarning }).(pulumi.IntOutput)
}

func (o VpnCertificateLocalOutput) CaIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.CaIdentifier }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) CmpPath() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.CmpPath }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) CmpRegenerationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.CmpRegenerationMethod }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) CmpServer() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.CmpServer }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) CmpServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.CmpServerCert }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.Comments }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) Csr() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.Csr }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EnrollProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EnrollProtocol }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EstCaId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EstCaId }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EstClientCert() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EstClientCert }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EstHttpPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EstHttpPassword }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EstHttpUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EstHttpUsername }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EstServer() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EstServer }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EstServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EstServerCert }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EstSrpPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EstSrpPassword }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) EstSrpUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.EstSrpUsername }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) IkeLocalid() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.IkeLocalid }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) IkeLocalidType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.IkeLocalidType }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) LastUpdated() pulumi.IntOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.IntOutput { return v.LastUpdated }).(pulumi.IntOutput)
}

func (o VpnCertificateLocalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) NameEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.NameEncoding }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VpnCertificateLocalOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) PrivateKeyRetain() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.PrivateKeyRetain }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) Range() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.Range }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) ScepPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringPtrOutput { return v.ScepPassword }).(pulumi.StringPtrOutput)
}

func (o VpnCertificateLocalOutput) ScepUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.ScepUrl }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o VpnCertificateLocalOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnCertificateLocal) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpnCertificateLocalArrayOutput struct{ *pulumi.OutputState }

func (VpnCertificateLocalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnCertificateLocal)(nil)).Elem()
}

func (o VpnCertificateLocalArrayOutput) ToVpnCertificateLocalArrayOutput() VpnCertificateLocalArrayOutput {
	return o
}

func (o VpnCertificateLocalArrayOutput) ToVpnCertificateLocalArrayOutputWithContext(ctx context.Context) VpnCertificateLocalArrayOutput {
	return o
}

func (o VpnCertificateLocalArrayOutput) Index(i pulumi.IntInput) VpnCertificateLocalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnCertificateLocal {
		return vs[0].([]*VpnCertificateLocal)[vs[1].(int)]
	}).(VpnCertificateLocalOutput)
}

type VpnCertificateLocalMapOutput struct{ *pulumi.OutputState }

func (VpnCertificateLocalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnCertificateLocal)(nil)).Elem()
}

func (o VpnCertificateLocalMapOutput) ToVpnCertificateLocalMapOutput() VpnCertificateLocalMapOutput {
	return o
}

func (o VpnCertificateLocalMapOutput) ToVpnCertificateLocalMapOutputWithContext(ctx context.Context) VpnCertificateLocalMapOutput {
	return o
}

func (o VpnCertificateLocalMapOutput) MapIndex(k pulumi.StringInput) VpnCertificateLocalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnCertificateLocal {
		return vs[0].(map[string]*VpnCertificateLocal)[vs[1].(string)]
	}).(VpnCertificateLocalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateLocalInput)(nil)).Elem(), &VpnCertificateLocal{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateLocalArrayInput)(nil)).Elem(), VpnCertificateLocalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnCertificateLocalMapInput)(nil)).Elem(), VpnCertificateLocalMap{})
	pulumi.RegisterOutputType(VpnCertificateLocalOutput{})
	pulumi.RegisterOutputType(VpnCertificateLocalArrayOutput{})
	pulumi.RegisterOutputType(VpnCertificateLocalMapOutput{})
}
