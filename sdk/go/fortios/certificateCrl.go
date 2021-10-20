// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Certificate Revocation List as a PEM file.
//
// ## Import
//
// Certificate Crl can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"
//
// ```sh
//  $ pulumi import fortios:index/certificateCrl:CertificateCrl labelname {{name}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type CertificateCrl struct {
	pulumi.CustomResourceState

	// Certificate Revocation List as a PEM file.
	Crl pulumi.StringOutput `pulumi:"crl"`
	// HTTP server URL for CRL auto-update.
	HttpUrl pulumi.StringOutput `pulumi:"httpUrl"`
	// Time at which CRL was last updated.
	LastUpdated pulumi.IntOutput `pulumi:"lastUpdated"`
	// LDAP server user password.
	LdapPassword pulumi.StringPtrOutput `pulumi:"ldapPassword"`
	// LDAP server name for CRL auto-update.
	LdapServer pulumi.StringOutput `pulumi:"ldapServer"`
	// LDAP server user name.
	LdapUsername pulumi.StringOutput `pulumi:"ldapUsername"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringOutput `pulumi:"range"`
	// Local certificate for SCEP communication for CRL auto-update.
	ScepCert pulumi.StringOutput `pulumi:"scepCert"`
	// SCEP server URL for CRL auto-update.
	ScepUrl pulumi.StringOutput `pulumi:"scepUrl"`
	// Certificate source type.
	Source pulumi.StringOutput `pulumi:"source"`
	// Source IP address for communications to a HTTP or SCEP CA server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
	UpdateInterval pulumi.IntOutput `pulumi:"updateInterval"`
	// VDOM for CRL update.
	UpdateVdom pulumi.StringOutput `pulumi:"updateVdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCertificateCrl registers a new resource with the given unique name, arguments, and options.
func NewCertificateCrl(ctx *pulumi.Context,
	name string, args *CertificateCrlArgs, opts ...pulumi.ResourceOption) (*CertificateCrl, error) {
	if args == nil {
		args = &CertificateCrlArgs{}
	}

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
	// Certificate Revocation List as a PEM file.
	Crl *string `pulumi:"crl"`
	// HTTP server URL for CRL auto-update.
	HttpUrl *string `pulumi:"httpUrl"`
	// Time at which CRL was last updated.
	LastUpdated *int `pulumi:"lastUpdated"`
	// LDAP server user password.
	LdapPassword *string `pulumi:"ldapPassword"`
	// LDAP server name for CRL auto-update.
	LdapServer *string `pulumi:"ldapServer"`
	// LDAP server user name.
	LdapUsername *string `pulumi:"ldapUsername"`
	// Name.
	Name *string `pulumi:"name"`
	// Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// Local certificate for SCEP communication for CRL auto-update.
	ScepCert *string `pulumi:"scepCert"`
	// SCEP server URL for CRL auto-update.
	ScepUrl *string `pulumi:"scepUrl"`
	// Certificate source type.
	Source *string `pulumi:"source"`
	// Source IP address for communications to a HTTP or SCEP CA server.
	SourceIp *string `pulumi:"sourceIp"`
	// Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
	UpdateInterval *int `pulumi:"updateInterval"`
	// VDOM for CRL update.
	UpdateVdom *string `pulumi:"updateVdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CertificateCrlState struct {
	// Certificate Revocation List as a PEM file.
	Crl pulumi.StringPtrInput
	// HTTP server URL for CRL auto-update.
	HttpUrl pulumi.StringPtrInput
	// Time at which CRL was last updated.
	LastUpdated pulumi.IntPtrInput
	// LDAP server user password.
	LdapPassword pulumi.StringPtrInput
	// LDAP server name for CRL auto-update.
	LdapServer pulumi.StringPtrInput
	// LDAP server user name.
	LdapUsername pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// Local certificate for SCEP communication for CRL auto-update.
	ScepCert pulumi.StringPtrInput
	// SCEP server URL for CRL auto-update.
	ScepUrl pulumi.StringPtrInput
	// Certificate source type.
	Source pulumi.StringPtrInput
	// Source IP address for communications to a HTTP or SCEP CA server.
	SourceIp pulumi.StringPtrInput
	// Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
	UpdateInterval pulumi.IntPtrInput
	// VDOM for CRL update.
	UpdateVdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CertificateCrlState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateCrlState)(nil)).Elem()
}

type certificateCrlArgs struct {
	// Certificate Revocation List as a PEM file.
	Crl *string `pulumi:"crl"`
	// HTTP server URL for CRL auto-update.
	HttpUrl *string `pulumi:"httpUrl"`
	// Time at which CRL was last updated.
	LastUpdated *int `pulumi:"lastUpdated"`
	// LDAP server user password.
	LdapPassword *string `pulumi:"ldapPassword"`
	// LDAP server name for CRL auto-update.
	LdapServer *string `pulumi:"ldapServer"`
	// LDAP server user name.
	LdapUsername *string `pulumi:"ldapUsername"`
	// Name.
	Name *string `pulumi:"name"`
	// Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// Local certificate for SCEP communication for CRL auto-update.
	ScepCert *string `pulumi:"scepCert"`
	// SCEP server URL for CRL auto-update.
	ScepUrl *string `pulumi:"scepUrl"`
	// Certificate source type.
	Source *string `pulumi:"source"`
	// Source IP address for communications to a HTTP or SCEP CA server.
	SourceIp *string `pulumi:"sourceIp"`
	// Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
	UpdateInterval *int `pulumi:"updateInterval"`
	// VDOM for CRL update.
	UpdateVdom *string `pulumi:"updateVdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a CertificateCrl resource.
type CertificateCrlArgs struct {
	// Certificate Revocation List as a PEM file.
	Crl pulumi.StringPtrInput
	// HTTP server URL for CRL auto-update.
	HttpUrl pulumi.StringPtrInput
	// Time at which CRL was last updated.
	LastUpdated pulumi.IntPtrInput
	// LDAP server user password.
	LdapPassword pulumi.StringPtrInput
	// LDAP server name for CRL auto-update.
	LdapServer pulumi.StringPtrInput
	// LDAP server user name.
	LdapUsername pulumi.StringPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Either global or VDOM IP address range for the certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// Local certificate for SCEP communication for CRL auto-update.
	ScepCert pulumi.StringPtrInput
	// SCEP server URL for CRL auto-update.
	ScepUrl pulumi.StringPtrInput
	// Certificate source type.
	Source pulumi.StringPtrInput
	// Source IP address for communications to a HTTP or SCEP CA server.
	SourceIp pulumi.StringPtrInput
	// Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.
	UpdateInterval pulumi.IntPtrInput
	// VDOM for CRL update.
	UpdateVdom pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
	return reflect.TypeOf((*CertificateCrl)(nil))
}

func (i *CertificateCrl) ToCertificateCrlOutput() CertificateCrlOutput {
	return i.ToCertificateCrlOutputWithContext(context.Background())
}

func (i *CertificateCrl) ToCertificateCrlOutputWithContext(ctx context.Context) CertificateCrlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCrlOutput)
}

func (i *CertificateCrl) ToCertificateCrlPtrOutput() CertificateCrlPtrOutput {
	return i.ToCertificateCrlPtrOutputWithContext(context.Background())
}

func (i *CertificateCrl) ToCertificateCrlPtrOutputWithContext(ctx context.Context) CertificateCrlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCrlPtrOutput)
}

type CertificateCrlPtrInput interface {
	pulumi.Input

	ToCertificateCrlPtrOutput() CertificateCrlPtrOutput
	ToCertificateCrlPtrOutputWithContext(ctx context.Context) CertificateCrlPtrOutput
}

type certificateCrlPtrType CertificateCrlArgs

func (*certificateCrlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateCrl)(nil))
}

func (i *certificateCrlPtrType) ToCertificateCrlPtrOutput() CertificateCrlPtrOutput {
	return i.ToCertificateCrlPtrOutputWithContext(context.Background())
}

func (i *certificateCrlPtrType) ToCertificateCrlPtrOutputWithContext(ctx context.Context) CertificateCrlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCrlPtrOutput)
}

// CertificateCrlArrayInput is an input type that accepts CertificateCrlArray and CertificateCrlArrayOutput values.
// You can construct a concrete instance of `CertificateCrlArrayInput` via:
//
//          CertificateCrlArray{ CertificateCrlArgs{...} }
type CertificateCrlArrayInput interface {
	pulumi.Input

	ToCertificateCrlArrayOutput() CertificateCrlArrayOutput
	ToCertificateCrlArrayOutputWithContext(context.Context) CertificateCrlArrayOutput
}

type CertificateCrlArray []CertificateCrlInput

func (CertificateCrlArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*CertificateCrl)(nil))
}

func (i CertificateCrlArray) ToCertificateCrlArrayOutput() CertificateCrlArrayOutput {
	return i.ToCertificateCrlArrayOutputWithContext(context.Background())
}

func (i CertificateCrlArray) ToCertificateCrlArrayOutputWithContext(ctx context.Context) CertificateCrlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCrlArrayOutput)
}

// CertificateCrlMapInput is an input type that accepts CertificateCrlMap and CertificateCrlMapOutput values.
// You can construct a concrete instance of `CertificateCrlMapInput` via:
//
//          CertificateCrlMap{ "key": CertificateCrlArgs{...} }
type CertificateCrlMapInput interface {
	pulumi.Input

	ToCertificateCrlMapOutput() CertificateCrlMapOutput
	ToCertificateCrlMapOutputWithContext(context.Context) CertificateCrlMapOutput
}

type CertificateCrlMap map[string]CertificateCrlInput

func (CertificateCrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*CertificateCrl)(nil))
}

func (i CertificateCrlMap) ToCertificateCrlMapOutput() CertificateCrlMapOutput {
	return i.ToCertificateCrlMapOutputWithContext(context.Background())
}

func (i CertificateCrlMap) ToCertificateCrlMapOutputWithContext(ctx context.Context) CertificateCrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateCrlMapOutput)
}

type CertificateCrlOutput struct {
	*pulumi.OutputState
}

func (CertificateCrlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateCrl)(nil))
}

func (o CertificateCrlOutput) ToCertificateCrlOutput() CertificateCrlOutput {
	return o
}

func (o CertificateCrlOutput) ToCertificateCrlOutputWithContext(ctx context.Context) CertificateCrlOutput {
	return o
}

func (o CertificateCrlOutput) ToCertificateCrlPtrOutput() CertificateCrlPtrOutput {
	return o.ToCertificateCrlPtrOutputWithContext(context.Background())
}

func (o CertificateCrlOutput) ToCertificateCrlPtrOutputWithContext(ctx context.Context) CertificateCrlPtrOutput {
	return o.ApplyT(func(v CertificateCrl) *CertificateCrl {
		return &v
	}).(CertificateCrlPtrOutput)
}

type CertificateCrlPtrOutput struct {
	*pulumi.OutputState
}

func (CertificateCrlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateCrl)(nil))
}

func (o CertificateCrlPtrOutput) ToCertificateCrlPtrOutput() CertificateCrlPtrOutput {
	return o
}

func (o CertificateCrlPtrOutput) ToCertificateCrlPtrOutputWithContext(ctx context.Context) CertificateCrlPtrOutput {
	return o
}

type CertificateCrlArrayOutput struct{ *pulumi.OutputState }

func (CertificateCrlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateCrl)(nil))
}

func (o CertificateCrlArrayOutput) ToCertificateCrlArrayOutput() CertificateCrlArrayOutput {
	return o
}

func (o CertificateCrlArrayOutput) ToCertificateCrlArrayOutputWithContext(ctx context.Context) CertificateCrlArrayOutput {
	return o
}

func (o CertificateCrlArrayOutput) Index(i pulumi.IntInput) CertificateCrlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateCrl {
		return vs[0].([]CertificateCrl)[vs[1].(int)]
	}).(CertificateCrlOutput)
}

type CertificateCrlMapOutput struct{ *pulumi.OutputState }

func (CertificateCrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CertificateCrl)(nil))
}

func (o CertificateCrlMapOutput) ToCertificateCrlMapOutput() CertificateCrlMapOutput {
	return o
}

func (o CertificateCrlMapOutput) ToCertificateCrlMapOutputWithContext(ctx context.Context) CertificateCrlMapOutput {
	return o
}

func (o CertificateCrlMapOutput) MapIndex(k pulumi.StringInput) CertificateCrlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CertificateCrl {
		return vs[0].(map[string]CertificateCrl)[vs[1].(string)]
	}).(CertificateCrlOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateCrlOutput{})
	pulumi.RegisterOutputType(CertificateCrlPtrOutput{})
	pulumi.RegisterOutputType(CertificateCrlArrayOutput{})
	pulumi.RegisterOutputType(CertificateCrlMapOutput{})
}