// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebProxyExplicit struct {
	pulumi.CustomResourceState

	DynamicSortSubtable     pulumi.StringPtrOutput                        `pulumi:"dynamicSortSubtable"`
	FtpIncomingPort         pulumi.StringOutput                           `pulumi:"ftpIncomingPort"`
	FtpOverHttp             pulumi.StringOutput                           `pulumi:"ftpOverHttp"`
	GetAllTables            pulumi.StringPtrOutput                        `pulumi:"getAllTables"`
	HttpConnectionMode      pulumi.StringOutput                           `pulumi:"httpConnectionMode"`
	HttpIncomingPort        pulumi.StringOutput                           `pulumi:"httpIncomingPort"`
	HttpsIncomingPort       pulumi.StringOutput                           `pulumi:"httpsIncomingPort"`
	HttpsReplacementMessage pulumi.StringOutput                           `pulumi:"httpsReplacementMessage"`
	IncomingIp              pulumi.StringOutput                           `pulumi:"incomingIp"`
	IncomingIp6             pulumi.StringOutput                           `pulumi:"incomingIp6"`
	Ipv6Status              pulumi.StringOutput                           `pulumi:"ipv6Status"`
	MessageUponServerError  pulumi.StringOutput                           `pulumi:"messageUponServerError"`
	OutgoingIp              pulumi.StringOutput                           `pulumi:"outgoingIp"`
	OutgoingIp6             pulumi.StringOutput                           `pulumi:"outgoingIp6"`
	PacFileData             pulumi.StringOutput                           `pulumi:"pacFileData"`
	PacFileName             pulumi.StringOutput                           `pulumi:"pacFileName"`
	PacFileServerPort       pulumi.StringOutput                           `pulumi:"pacFileServerPort"`
	PacFileServerStatus     pulumi.StringOutput                           `pulumi:"pacFileServerStatus"`
	PacFileThroughHttps     pulumi.StringOutput                           `pulumi:"pacFileThroughHttps"`
	PacFileUrl              pulumi.StringOutput                           `pulumi:"pacFileUrl"`
	PacPolicies             WebProxyExplicitPacPolicyArrayOutput          `pulumi:"pacPolicies"`
	PrefDnsResult           pulumi.StringOutput                           `pulumi:"prefDnsResult"`
	Realm                   pulumi.StringOutput                           `pulumi:"realm"`
	SecDefaultAction        pulumi.StringOutput                           `pulumi:"secDefaultAction"`
	SecureWebProxy          pulumi.StringOutput                           `pulumi:"secureWebProxy"`
	SecureWebProxyCerts     WebProxyExplicitSecureWebProxyCertArrayOutput `pulumi:"secureWebProxyCerts"`
	Socks                   pulumi.StringOutput                           `pulumi:"socks"`
	SocksIncomingPort       pulumi.StringOutput                           `pulumi:"socksIncomingPort"`
	SslAlgorithm            pulumi.StringOutput                           `pulumi:"sslAlgorithm"`
	SslDhBits               pulumi.StringOutput                           `pulumi:"sslDhBits"`
	Status                  pulumi.StringOutput                           `pulumi:"status"`
	StrictGuest             pulumi.StringOutput                           `pulumi:"strictGuest"`
	TraceAuthNoRsp          pulumi.StringOutput                           `pulumi:"traceAuthNoRsp"`
	UnknownHttpVersion      pulumi.StringOutput                           `pulumi:"unknownHttpVersion"`
	Vdomparam               pulumi.StringPtrOutput                        `pulumi:"vdomparam"`
}

// NewWebProxyExplicit registers a new resource with the given unique name, arguments, and options.
func NewWebProxyExplicit(ctx *pulumi.Context,
	name string, args *WebProxyExplicitArgs, opts ...pulumi.ResourceOption) (*WebProxyExplicit, error) {
	if args == nil {
		args = &WebProxyExplicitArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebProxyExplicit
	err := ctx.RegisterResource("fortios:index/webProxyExplicit:WebProxyExplicit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebProxyExplicit gets an existing WebProxyExplicit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebProxyExplicit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebProxyExplicitState, opts ...pulumi.ResourceOption) (*WebProxyExplicit, error) {
	var resource WebProxyExplicit
	err := ctx.ReadResource("fortios:index/webProxyExplicit:WebProxyExplicit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebProxyExplicit resources.
type webProxyExplicitState struct {
	DynamicSortSubtable     *string                              `pulumi:"dynamicSortSubtable"`
	FtpIncomingPort         *string                              `pulumi:"ftpIncomingPort"`
	FtpOverHttp             *string                              `pulumi:"ftpOverHttp"`
	GetAllTables            *string                              `pulumi:"getAllTables"`
	HttpConnectionMode      *string                              `pulumi:"httpConnectionMode"`
	HttpIncomingPort        *string                              `pulumi:"httpIncomingPort"`
	HttpsIncomingPort       *string                              `pulumi:"httpsIncomingPort"`
	HttpsReplacementMessage *string                              `pulumi:"httpsReplacementMessage"`
	IncomingIp              *string                              `pulumi:"incomingIp"`
	IncomingIp6             *string                              `pulumi:"incomingIp6"`
	Ipv6Status              *string                              `pulumi:"ipv6Status"`
	MessageUponServerError  *string                              `pulumi:"messageUponServerError"`
	OutgoingIp              *string                              `pulumi:"outgoingIp"`
	OutgoingIp6             *string                              `pulumi:"outgoingIp6"`
	PacFileData             *string                              `pulumi:"pacFileData"`
	PacFileName             *string                              `pulumi:"pacFileName"`
	PacFileServerPort       *string                              `pulumi:"pacFileServerPort"`
	PacFileServerStatus     *string                              `pulumi:"pacFileServerStatus"`
	PacFileThroughHttps     *string                              `pulumi:"pacFileThroughHttps"`
	PacFileUrl              *string                              `pulumi:"pacFileUrl"`
	PacPolicies             []WebProxyExplicitPacPolicy          `pulumi:"pacPolicies"`
	PrefDnsResult           *string                              `pulumi:"prefDnsResult"`
	Realm                   *string                              `pulumi:"realm"`
	SecDefaultAction        *string                              `pulumi:"secDefaultAction"`
	SecureWebProxy          *string                              `pulumi:"secureWebProxy"`
	SecureWebProxyCerts     []WebProxyExplicitSecureWebProxyCert `pulumi:"secureWebProxyCerts"`
	Socks                   *string                              `pulumi:"socks"`
	SocksIncomingPort       *string                              `pulumi:"socksIncomingPort"`
	SslAlgorithm            *string                              `pulumi:"sslAlgorithm"`
	SslDhBits               *string                              `pulumi:"sslDhBits"`
	Status                  *string                              `pulumi:"status"`
	StrictGuest             *string                              `pulumi:"strictGuest"`
	TraceAuthNoRsp          *string                              `pulumi:"traceAuthNoRsp"`
	UnknownHttpVersion      *string                              `pulumi:"unknownHttpVersion"`
	Vdomparam               *string                              `pulumi:"vdomparam"`
}

type WebProxyExplicitState struct {
	DynamicSortSubtable     pulumi.StringPtrInput
	FtpIncomingPort         pulumi.StringPtrInput
	FtpOverHttp             pulumi.StringPtrInput
	GetAllTables            pulumi.StringPtrInput
	HttpConnectionMode      pulumi.StringPtrInput
	HttpIncomingPort        pulumi.StringPtrInput
	HttpsIncomingPort       pulumi.StringPtrInput
	HttpsReplacementMessage pulumi.StringPtrInput
	IncomingIp              pulumi.StringPtrInput
	IncomingIp6             pulumi.StringPtrInput
	Ipv6Status              pulumi.StringPtrInput
	MessageUponServerError  pulumi.StringPtrInput
	OutgoingIp              pulumi.StringPtrInput
	OutgoingIp6             pulumi.StringPtrInput
	PacFileData             pulumi.StringPtrInput
	PacFileName             pulumi.StringPtrInput
	PacFileServerPort       pulumi.StringPtrInput
	PacFileServerStatus     pulumi.StringPtrInput
	PacFileThroughHttps     pulumi.StringPtrInput
	PacFileUrl              pulumi.StringPtrInput
	PacPolicies             WebProxyExplicitPacPolicyArrayInput
	PrefDnsResult           pulumi.StringPtrInput
	Realm                   pulumi.StringPtrInput
	SecDefaultAction        pulumi.StringPtrInput
	SecureWebProxy          pulumi.StringPtrInput
	SecureWebProxyCerts     WebProxyExplicitSecureWebProxyCertArrayInput
	Socks                   pulumi.StringPtrInput
	SocksIncomingPort       pulumi.StringPtrInput
	SslAlgorithm            pulumi.StringPtrInput
	SslDhBits               pulumi.StringPtrInput
	Status                  pulumi.StringPtrInput
	StrictGuest             pulumi.StringPtrInput
	TraceAuthNoRsp          pulumi.StringPtrInput
	UnknownHttpVersion      pulumi.StringPtrInput
	Vdomparam               pulumi.StringPtrInput
}

func (WebProxyExplicitState) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyExplicitState)(nil)).Elem()
}

type webProxyExplicitArgs struct {
	DynamicSortSubtable     *string                              `pulumi:"dynamicSortSubtable"`
	FtpIncomingPort         *string                              `pulumi:"ftpIncomingPort"`
	FtpOverHttp             *string                              `pulumi:"ftpOverHttp"`
	GetAllTables            *string                              `pulumi:"getAllTables"`
	HttpConnectionMode      *string                              `pulumi:"httpConnectionMode"`
	HttpIncomingPort        *string                              `pulumi:"httpIncomingPort"`
	HttpsIncomingPort       *string                              `pulumi:"httpsIncomingPort"`
	HttpsReplacementMessage *string                              `pulumi:"httpsReplacementMessage"`
	IncomingIp              *string                              `pulumi:"incomingIp"`
	IncomingIp6             *string                              `pulumi:"incomingIp6"`
	Ipv6Status              *string                              `pulumi:"ipv6Status"`
	MessageUponServerError  *string                              `pulumi:"messageUponServerError"`
	OutgoingIp              *string                              `pulumi:"outgoingIp"`
	OutgoingIp6             *string                              `pulumi:"outgoingIp6"`
	PacFileData             *string                              `pulumi:"pacFileData"`
	PacFileName             *string                              `pulumi:"pacFileName"`
	PacFileServerPort       *string                              `pulumi:"pacFileServerPort"`
	PacFileServerStatus     *string                              `pulumi:"pacFileServerStatus"`
	PacFileThroughHttps     *string                              `pulumi:"pacFileThroughHttps"`
	PacFileUrl              *string                              `pulumi:"pacFileUrl"`
	PacPolicies             []WebProxyExplicitPacPolicy          `pulumi:"pacPolicies"`
	PrefDnsResult           *string                              `pulumi:"prefDnsResult"`
	Realm                   *string                              `pulumi:"realm"`
	SecDefaultAction        *string                              `pulumi:"secDefaultAction"`
	SecureWebProxy          *string                              `pulumi:"secureWebProxy"`
	SecureWebProxyCerts     []WebProxyExplicitSecureWebProxyCert `pulumi:"secureWebProxyCerts"`
	Socks                   *string                              `pulumi:"socks"`
	SocksIncomingPort       *string                              `pulumi:"socksIncomingPort"`
	SslAlgorithm            *string                              `pulumi:"sslAlgorithm"`
	SslDhBits               *string                              `pulumi:"sslDhBits"`
	Status                  *string                              `pulumi:"status"`
	StrictGuest             *string                              `pulumi:"strictGuest"`
	TraceAuthNoRsp          *string                              `pulumi:"traceAuthNoRsp"`
	UnknownHttpVersion      *string                              `pulumi:"unknownHttpVersion"`
	Vdomparam               *string                              `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebProxyExplicit resource.
type WebProxyExplicitArgs struct {
	DynamicSortSubtable     pulumi.StringPtrInput
	FtpIncomingPort         pulumi.StringPtrInput
	FtpOverHttp             pulumi.StringPtrInput
	GetAllTables            pulumi.StringPtrInput
	HttpConnectionMode      pulumi.StringPtrInput
	HttpIncomingPort        pulumi.StringPtrInput
	HttpsIncomingPort       pulumi.StringPtrInput
	HttpsReplacementMessage pulumi.StringPtrInput
	IncomingIp              pulumi.StringPtrInput
	IncomingIp6             pulumi.StringPtrInput
	Ipv6Status              pulumi.StringPtrInput
	MessageUponServerError  pulumi.StringPtrInput
	OutgoingIp              pulumi.StringPtrInput
	OutgoingIp6             pulumi.StringPtrInput
	PacFileData             pulumi.StringPtrInput
	PacFileName             pulumi.StringPtrInput
	PacFileServerPort       pulumi.StringPtrInput
	PacFileServerStatus     pulumi.StringPtrInput
	PacFileThroughHttps     pulumi.StringPtrInput
	PacFileUrl              pulumi.StringPtrInput
	PacPolicies             WebProxyExplicitPacPolicyArrayInput
	PrefDnsResult           pulumi.StringPtrInput
	Realm                   pulumi.StringPtrInput
	SecDefaultAction        pulumi.StringPtrInput
	SecureWebProxy          pulumi.StringPtrInput
	SecureWebProxyCerts     WebProxyExplicitSecureWebProxyCertArrayInput
	Socks                   pulumi.StringPtrInput
	SocksIncomingPort       pulumi.StringPtrInput
	SslAlgorithm            pulumi.StringPtrInput
	SslDhBits               pulumi.StringPtrInput
	Status                  pulumi.StringPtrInput
	StrictGuest             pulumi.StringPtrInput
	TraceAuthNoRsp          pulumi.StringPtrInput
	UnknownHttpVersion      pulumi.StringPtrInput
	Vdomparam               pulumi.StringPtrInput
}

func (WebProxyExplicitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyExplicitArgs)(nil)).Elem()
}

type WebProxyExplicitInput interface {
	pulumi.Input

	ToWebProxyExplicitOutput() WebProxyExplicitOutput
	ToWebProxyExplicitOutputWithContext(ctx context.Context) WebProxyExplicitOutput
}

func (*WebProxyExplicit) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyExplicit)(nil)).Elem()
}

func (i *WebProxyExplicit) ToWebProxyExplicitOutput() WebProxyExplicitOutput {
	return i.ToWebProxyExplicitOutputWithContext(context.Background())
}

func (i *WebProxyExplicit) ToWebProxyExplicitOutputWithContext(ctx context.Context) WebProxyExplicitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyExplicitOutput)
}

// WebProxyExplicitArrayInput is an input type that accepts WebProxyExplicitArray and WebProxyExplicitArrayOutput values.
// You can construct a concrete instance of `WebProxyExplicitArrayInput` via:
//
//	WebProxyExplicitArray{ WebProxyExplicitArgs{...} }
type WebProxyExplicitArrayInput interface {
	pulumi.Input

	ToWebProxyExplicitArrayOutput() WebProxyExplicitArrayOutput
	ToWebProxyExplicitArrayOutputWithContext(context.Context) WebProxyExplicitArrayOutput
}

type WebProxyExplicitArray []WebProxyExplicitInput

func (WebProxyExplicitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyExplicit)(nil)).Elem()
}

func (i WebProxyExplicitArray) ToWebProxyExplicitArrayOutput() WebProxyExplicitArrayOutput {
	return i.ToWebProxyExplicitArrayOutputWithContext(context.Background())
}

func (i WebProxyExplicitArray) ToWebProxyExplicitArrayOutputWithContext(ctx context.Context) WebProxyExplicitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyExplicitArrayOutput)
}

// WebProxyExplicitMapInput is an input type that accepts WebProxyExplicitMap and WebProxyExplicitMapOutput values.
// You can construct a concrete instance of `WebProxyExplicitMapInput` via:
//
//	WebProxyExplicitMap{ "key": WebProxyExplicitArgs{...} }
type WebProxyExplicitMapInput interface {
	pulumi.Input

	ToWebProxyExplicitMapOutput() WebProxyExplicitMapOutput
	ToWebProxyExplicitMapOutputWithContext(context.Context) WebProxyExplicitMapOutput
}

type WebProxyExplicitMap map[string]WebProxyExplicitInput

func (WebProxyExplicitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyExplicit)(nil)).Elem()
}

func (i WebProxyExplicitMap) ToWebProxyExplicitMapOutput() WebProxyExplicitMapOutput {
	return i.ToWebProxyExplicitMapOutputWithContext(context.Background())
}

func (i WebProxyExplicitMap) ToWebProxyExplicitMapOutputWithContext(ctx context.Context) WebProxyExplicitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebProxyExplicitMapOutput)
}

type WebProxyExplicitOutput struct{ *pulumi.OutputState }

func (WebProxyExplicitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebProxyExplicit)(nil)).Elem()
}

func (o WebProxyExplicitOutput) ToWebProxyExplicitOutput() WebProxyExplicitOutput {
	return o
}

func (o WebProxyExplicitOutput) ToWebProxyExplicitOutputWithContext(ctx context.Context) WebProxyExplicitOutput {
	return o
}

func (o WebProxyExplicitOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o WebProxyExplicitOutput) FtpIncomingPort() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.FtpIncomingPort }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) FtpOverHttp() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.FtpOverHttp }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o WebProxyExplicitOutput) HttpConnectionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.HttpConnectionMode }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) HttpIncomingPort() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.HttpIncomingPort }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) HttpsIncomingPort() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.HttpsIncomingPort }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) HttpsReplacementMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.HttpsReplacementMessage }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) IncomingIp() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.IncomingIp }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) IncomingIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.IncomingIp6 }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) Ipv6Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.Ipv6Status }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) MessageUponServerError() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.MessageUponServerError }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) OutgoingIp() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.OutgoingIp }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) OutgoingIp6() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.OutgoingIp6 }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) PacFileData() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.PacFileData }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) PacFileName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.PacFileName }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) PacFileServerPort() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.PacFileServerPort }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) PacFileServerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.PacFileServerStatus }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) PacFileThroughHttps() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.PacFileThroughHttps }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) PacFileUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.PacFileUrl }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) PacPolicies() WebProxyExplicitPacPolicyArrayOutput {
	return o.ApplyT(func(v *WebProxyExplicit) WebProxyExplicitPacPolicyArrayOutput { return v.PacPolicies }).(WebProxyExplicitPacPolicyArrayOutput)
}

func (o WebProxyExplicitOutput) PrefDnsResult() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.PrefDnsResult }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) Realm() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.Realm }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) SecDefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.SecDefaultAction }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) SecureWebProxy() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.SecureWebProxy }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) SecureWebProxyCerts() WebProxyExplicitSecureWebProxyCertArrayOutput {
	return o.ApplyT(func(v *WebProxyExplicit) WebProxyExplicitSecureWebProxyCertArrayOutput { return v.SecureWebProxyCerts }).(WebProxyExplicitSecureWebProxyCertArrayOutput)
}

func (o WebProxyExplicitOutput) Socks() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.Socks }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) SocksIncomingPort() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.SocksIncomingPort }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) SslAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.SslAlgorithm }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) SslDhBits() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.SslDhBits }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) StrictGuest() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.StrictGuest }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) TraceAuthNoRsp() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.TraceAuthNoRsp }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) UnknownHttpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringOutput { return v.UnknownHttpVersion }).(pulumi.StringOutput)
}

func (o WebProxyExplicitOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebProxyExplicit) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebProxyExplicitArrayOutput struct{ *pulumi.OutputState }

func (WebProxyExplicitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebProxyExplicit)(nil)).Elem()
}

func (o WebProxyExplicitArrayOutput) ToWebProxyExplicitArrayOutput() WebProxyExplicitArrayOutput {
	return o
}

func (o WebProxyExplicitArrayOutput) ToWebProxyExplicitArrayOutputWithContext(ctx context.Context) WebProxyExplicitArrayOutput {
	return o
}

func (o WebProxyExplicitArrayOutput) Index(i pulumi.IntInput) WebProxyExplicitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebProxyExplicit {
		return vs[0].([]*WebProxyExplicit)[vs[1].(int)]
	}).(WebProxyExplicitOutput)
}

type WebProxyExplicitMapOutput struct{ *pulumi.OutputState }

func (WebProxyExplicitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebProxyExplicit)(nil)).Elem()
}

func (o WebProxyExplicitMapOutput) ToWebProxyExplicitMapOutput() WebProxyExplicitMapOutput {
	return o
}

func (o WebProxyExplicitMapOutput) ToWebProxyExplicitMapOutputWithContext(ctx context.Context) WebProxyExplicitMapOutput {
	return o
}

func (o WebProxyExplicitMapOutput) MapIndex(k pulumi.StringInput) WebProxyExplicitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebProxyExplicit {
		return vs[0].(map[string]*WebProxyExplicit)[vs[1].(string)]
	}).(WebProxyExplicitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyExplicitInput)(nil)).Elem(), &WebProxyExplicit{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyExplicitArrayInput)(nil)).Elem(), WebProxyExplicitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebProxyExplicitMapInput)(nil)).Elem(), WebProxyExplicitMap{})
	pulumi.RegisterOutputType(WebProxyExplicitOutput{})
	pulumi.RegisterOutputType(WebProxyExplicitArrayOutput{})
	pulumi.RegisterOutputType(WebProxyExplicitMapOutput{})
}
