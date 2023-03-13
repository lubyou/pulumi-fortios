// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FirewallSslSshProfile struct {
	pulumi.CustomResourceState

	Allowlist                    pulumi.StringOutput                       `pulumi:"allowlist"`
	BlockBlacklistedCertificates pulumi.StringOutput                       `pulumi:"blockBlacklistedCertificates"`
	BlockBlocklistedCertificates pulumi.StringOutput                       `pulumi:"blockBlocklistedCertificates"`
	Caname                       pulumi.StringOutput                       `pulumi:"caname"`
	Comment                      pulumi.StringPtrOutput                    `pulumi:"comment"`
	Dot                          FirewallSslSshProfileDotOutput            `pulumi:"dot"`
	DynamicSortSubtable          pulumi.StringPtrOutput                    `pulumi:"dynamicSortSubtable"`
	Ftps                         FirewallSslSshProfileFtpsOutput           `pulumi:"ftps"`
	Https                        FirewallSslSshProfileHttpsOutput          `pulumi:"https"`
	Imaps                        FirewallSslSshProfileImapsOutput          `pulumi:"imaps"`
	MapiOverHttps                pulumi.StringOutput                       `pulumi:"mapiOverHttps"`
	Name                         pulumi.StringOutput                       `pulumi:"name"`
	Pop3s                        FirewallSslSshProfilePop3sOutput          `pulumi:"pop3s"`
	RpcOverHttps                 pulumi.StringOutput                       `pulumi:"rpcOverHttps"`
	ServerCert                   pulumi.StringOutput                       `pulumi:"serverCert"`
	ServerCertMode               pulumi.StringOutput                       `pulumi:"serverCertMode"`
	Smtps                        FirewallSslSshProfileSmtpsOutput          `pulumi:"smtps"`
	Ssh                          FirewallSslSshProfileSshOutput            `pulumi:"ssh"`
	Ssl                          FirewallSslSshProfileSslOutput            `pulumi:"ssl"`
	SslAnomaliesLog              pulumi.StringOutput                       `pulumi:"sslAnomaliesLog"`
	SslAnomalyLog                pulumi.StringOutput                       `pulumi:"sslAnomalyLog"`
	SslExemptionIpRating         pulumi.StringOutput                       `pulumi:"sslExemptionIpRating"`
	SslExemptionLog              pulumi.StringOutput                       `pulumi:"sslExemptionLog"`
	SslExemptionsLog             pulumi.StringOutput                       `pulumi:"sslExemptionsLog"`
	SslExempts                   FirewallSslSshProfileSslExemptArrayOutput `pulumi:"sslExempts"`
	SslHandshakeLog              pulumi.StringOutput                       `pulumi:"sslHandshakeLog"`
	SslNegotiationLog            pulumi.StringOutput                       `pulumi:"sslNegotiationLog"`
	SslServerCertLog             pulumi.StringOutput                       `pulumi:"sslServerCertLog"`
	SslServers                   FirewallSslSshProfileSslServerArrayOutput `pulumi:"sslServers"`
	SupportedAlpn                pulumi.StringOutput                       `pulumi:"supportedAlpn"`
	UntrustedCaname              pulumi.StringOutput                       `pulumi:"untrustedCaname"`
	UseSslServer                 pulumi.StringOutput                       `pulumi:"useSslServer"`
	Vdomparam                    pulumi.StringPtrOutput                    `pulumi:"vdomparam"`
	Whitelist                    pulumi.StringOutput                       `pulumi:"whitelist"`
}

// NewFirewallSslSshProfile registers a new resource with the given unique name, arguments, and options.
func NewFirewallSslSshProfile(ctx *pulumi.Context,
	name string, args *FirewallSslSshProfileArgs, opts ...pulumi.ResourceOption) (*FirewallSslSshProfile, error) {
	if args == nil {
		args = &FirewallSslSshProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallSslSshProfile
	err := ctx.RegisterResource("fortios:index/firewallSslSshProfile:FirewallSslSshProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSslSshProfile gets an existing FirewallSslSshProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSslSshProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSslSshProfileState, opts ...pulumi.ResourceOption) (*FirewallSslSshProfile, error) {
	var resource FirewallSslSshProfile
	err := ctx.ReadResource("fortios:index/firewallSslSshProfile:FirewallSslSshProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSslSshProfile resources.
type firewallSslSshProfileState struct {
	Allowlist                    *string                          `pulumi:"allowlist"`
	BlockBlacklistedCertificates *string                          `pulumi:"blockBlacklistedCertificates"`
	BlockBlocklistedCertificates *string                          `pulumi:"blockBlocklistedCertificates"`
	Caname                       *string                          `pulumi:"caname"`
	Comment                      *string                          `pulumi:"comment"`
	Dot                          *FirewallSslSshProfileDot        `pulumi:"dot"`
	DynamicSortSubtable          *string                          `pulumi:"dynamicSortSubtable"`
	Ftps                         *FirewallSslSshProfileFtps       `pulumi:"ftps"`
	Https                        *FirewallSslSshProfileHttps      `pulumi:"https"`
	Imaps                        *FirewallSslSshProfileImaps      `pulumi:"imaps"`
	MapiOverHttps                *string                          `pulumi:"mapiOverHttps"`
	Name                         *string                          `pulumi:"name"`
	Pop3s                        *FirewallSslSshProfilePop3s      `pulumi:"pop3s"`
	RpcOverHttps                 *string                          `pulumi:"rpcOverHttps"`
	ServerCert                   *string                          `pulumi:"serverCert"`
	ServerCertMode               *string                          `pulumi:"serverCertMode"`
	Smtps                        *FirewallSslSshProfileSmtps      `pulumi:"smtps"`
	Ssh                          *FirewallSslSshProfileSsh        `pulumi:"ssh"`
	Ssl                          *FirewallSslSshProfileSsl        `pulumi:"ssl"`
	SslAnomaliesLog              *string                          `pulumi:"sslAnomaliesLog"`
	SslAnomalyLog                *string                          `pulumi:"sslAnomalyLog"`
	SslExemptionIpRating         *string                          `pulumi:"sslExemptionIpRating"`
	SslExemptionLog              *string                          `pulumi:"sslExemptionLog"`
	SslExemptionsLog             *string                          `pulumi:"sslExemptionsLog"`
	SslExempts                   []FirewallSslSshProfileSslExempt `pulumi:"sslExempts"`
	SslHandshakeLog              *string                          `pulumi:"sslHandshakeLog"`
	SslNegotiationLog            *string                          `pulumi:"sslNegotiationLog"`
	SslServerCertLog             *string                          `pulumi:"sslServerCertLog"`
	SslServers                   []FirewallSslSshProfileSslServer `pulumi:"sslServers"`
	SupportedAlpn                *string                          `pulumi:"supportedAlpn"`
	UntrustedCaname              *string                          `pulumi:"untrustedCaname"`
	UseSslServer                 *string                          `pulumi:"useSslServer"`
	Vdomparam                    *string                          `pulumi:"vdomparam"`
	Whitelist                    *string                          `pulumi:"whitelist"`
}

type FirewallSslSshProfileState struct {
	Allowlist                    pulumi.StringPtrInput
	BlockBlacklistedCertificates pulumi.StringPtrInput
	BlockBlocklistedCertificates pulumi.StringPtrInput
	Caname                       pulumi.StringPtrInput
	Comment                      pulumi.StringPtrInput
	Dot                          FirewallSslSshProfileDotPtrInput
	DynamicSortSubtable          pulumi.StringPtrInput
	Ftps                         FirewallSslSshProfileFtpsPtrInput
	Https                        FirewallSslSshProfileHttpsPtrInput
	Imaps                        FirewallSslSshProfileImapsPtrInput
	MapiOverHttps                pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	Pop3s                        FirewallSslSshProfilePop3sPtrInput
	RpcOverHttps                 pulumi.StringPtrInput
	ServerCert                   pulumi.StringPtrInput
	ServerCertMode               pulumi.StringPtrInput
	Smtps                        FirewallSslSshProfileSmtpsPtrInput
	Ssh                          FirewallSslSshProfileSshPtrInput
	Ssl                          FirewallSslSshProfileSslPtrInput
	SslAnomaliesLog              pulumi.StringPtrInput
	SslAnomalyLog                pulumi.StringPtrInput
	SslExemptionIpRating         pulumi.StringPtrInput
	SslExemptionLog              pulumi.StringPtrInput
	SslExemptionsLog             pulumi.StringPtrInput
	SslExempts                   FirewallSslSshProfileSslExemptArrayInput
	SslHandshakeLog              pulumi.StringPtrInput
	SslNegotiationLog            pulumi.StringPtrInput
	SslServerCertLog             pulumi.StringPtrInput
	SslServers                   FirewallSslSshProfileSslServerArrayInput
	SupportedAlpn                pulumi.StringPtrInput
	UntrustedCaname              pulumi.StringPtrInput
	UseSslServer                 pulumi.StringPtrInput
	Vdomparam                    pulumi.StringPtrInput
	Whitelist                    pulumi.StringPtrInput
}

func (FirewallSslSshProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSslSshProfileState)(nil)).Elem()
}

type firewallSslSshProfileArgs struct {
	Allowlist                    *string                          `pulumi:"allowlist"`
	BlockBlacklistedCertificates *string                          `pulumi:"blockBlacklistedCertificates"`
	BlockBlocklistedCertificates *string                          `pulumi:"blockBlocklistedCertificates"`
	Caname                       *string                          `pulumi:"caname"`
	Comment                      *string                          `pulumi:"comment"`
	Dot                          *FirewallSslSshProfileDot        `pulumi:"dot"`
	DynamicSortSubtable          *string                          `pulumi:"dynamicSortSubtable"`
	Ftps                         *FirewallSslSshProfileFtps       `pulumi:"ftps"`
	Https                        *FirewallSslSshProfileHttps      `pulumi:"https"`
	Imaps                        *FirewallSslSshProfileImaps      `pulumi:"imaps"`
	MapiOverHttps                *string                          `pulumi:"mapiOverHttps"`
	Name                         *string                          `pulumi:"name"`
	Pop3s                        *FirewallSslSshProfilePop3s      `pulumi:"pop3s"`
	RpcOverHttps                 *string                          `pulumi:"rpcOverHttps"`
	ServerCert                   *string                          `pulumi:"serverCert"`
	ServerCertMode               *string                          `pulumi:"serverCertMode"`
	Smtps                        *FirewallSslSshProfileSmtps      `pulumi:"smtps"`
	Ssh                          *FirewallSslSshProfileSsh        `pulumi:"ssh"`
	Ssl                          *FirewallSslSshProfileSsl        `pulumi:"ssl"`
	SslAnomaliesLog              *string                          `pulumi:"sslAnomaliesLog"`
	SslAnomalyLog                *string                          `pulumi:"sslAnomalyLog"`
	SslExemptionIpRating         *string                          `pulumi:"sslExemptionIpRating"`
	SslExemptionLog              *string                          `pulumi:"sslExemptionLog"`
	SslExemptionsLog             *string                          `pulumi:"sslExemptionsLog"`
	SslExempts                   []FirewallSslSshProfileSslExempt `pulumi:"sslExempts"`
	SslHandshakeLog              *string                          `pulumi:"sslHandshakeLog"`
	SslNegotiationLog            *string                          `pulumi:"sslNegotiationLog"`
	SslServerCertLog             *string                          `pulumi:"sslServerCertLog"`
	SslServers                   []FirewallSslSshProfileSslServer `pulumi:"sslServers"`
	SupportedAlpn                *string                          `pulumi:"supportedAlpn"`
	UntrustedCaname              *string                          `pulumi:"untrustedCaname"`
	UseSslServer                 *string                          `pulumi:"useSslServer"`
	Vdomparam                    *string                          `pulumi:"vdomparam"`
	Whitelist                    *string                          `pulumi:"whitelist"`
}

// The set of arguments for constructing a FirewallSslSshProfile resource.
type FirewallSslSshProfileArgs struct {
	Allowlist                    pulumi.StringPtrInput
	BlockBlacklistedCertificates pulumi.StringPtrInput
	BlockBlocklistedCertificates pulumi.StringPtrInput
	Caname                       pulumi.StringPtrInput
	Comment                      pulumi.StringPtrInput
	Dot                          FirewallSslSshProfileDotPtrInput
	DynamicSortSubtable          pulumi.StringPtrInput
	Ftps                         FirewallSslSshProfileFtpsPtrInput
	Https                        FirewallSslSshProfileHttpsPtrInput
	Imaps                        FirewallSslSshProfileImapsPtrInput
	MapiOverHttps                pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	Pop3s                        FirewallSslSshProfilePop3sPtrInput
	RpcOverHttps                 pulumi.StringPtrInput
	ServerCert                   pulumi.StringPtrInput
	ServerCertMode               pulumi.StringPtrInput
	Smtps                        FirewallSslSshProfileSmtpsPtrInput
	Ssh                          FirewallSslSshProfileSshPtrInput
	Ssl                          FirewallSslSshProfileSslPtrInput
	SslAnomaliesLog              pulumi.StringPtrInput
	SslAnomalyLog                pulumi.StringPtrInput
	SslExemptionIpRating         pulumi.StringPtrInput
	SslExemptionLog              pulumi.StringPtrInput
	SslExemptionsLog             pulumi.StringPtrInput
	SslExempts                   FirewallSslSshProfileSslExemptArrayInput
	SslHandshakeLog              pulumi.StringPtrInput
	SslNegotiationLog            pulumi.StringPtrInput
	SslServerCertLog             pulumi.StringPtrInput
	SslServers                   FirewallSslSshProfileSslServerArrayInput
	SupportedAlpn                pulumi.StringPtrInput
	UntrustedCaname              pulumi.StringPtrInput
	UseSslServer                 pulumi.StringPtrInput
	Vdomparam                    pulumi.StringPtrInput
	Whitelist                    pulumi.StringPtrInput
}

func (FirewallSslSshProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSslSshProfileArgs)(nil)).Elem()
}

type FirewallSslSshProfileInput interface {
	pulumi.Input

	ToFirewallSslSshProfileOutput() FirewallSslSshProfileOutput
	ToFirewallSslSshProfileOutputWithContext(ctx context.Context) FirewallSslSshProfileOutput
}

func (*FirewallSslSshProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSslSshProfile)(nil)).Elem()
}

func (i *FirewallSslSshProfile) ToFirewallSslSshProfileOutput() FirewallSslSshProfileOutput {
	return i.ToFirewallSslSshProfileOutputWithContext(context.Background())
}

func (i *FirewallSslSshProfile) ToFirewallSslSshProfileOutputWithContext(ctx context.Context) FirewallSslSshProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslSshProfileOutput)
}

// FirewallSslSshProfileArrayInput is an input type that accepts FirewallSslSshProfileArray and FirewallSslSshProfileArrayOutput values.
// You can construct a concrete instance of `FirewallSslSshProfileArrayInput` via:
//
//	FirewallSslSshProfileArray{ FirewallSslSshProfileArgs{...} }
type FirewallSslSshProfileArrayInput interface {
	pulumi.Input

	ToFirewallSslSshProfileArrayOutput() FirewallSslSshProfileArrayOutput
	ToFirewallSslSshProfileArrayOutputWithContext(context.Context) FirewallSslSshProfileArrayOutput
}

type FirewallSslSshProfileArray []FirewallSslSshProfileInput

func (FirewallSslSshProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSslSshProfile)(nil)).Elem()
}

func (i FirewallSslSshProfileArray) ToFirewallSslSshProfileArrayOutput() FirewallSslSshProfileArrayOutput {
	return i.ToFirewallSslSshProfileArrayOutputWithContext(context.Background())
}

func (i FirewallSslSshProfileArray) ToFirewallSslSshProfileArrayOutputWithContext(ctx context.Context) FirewallSslSshProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslSshProfileArrayOutput)
}

// FirewallSslSshProfileMapInput is an input type that accepts FirewallSslSshProfileMap and FirewallSslSshProfileMapOutput values.
// You can construct a concrete instance of `FirewallSslSshProfileMapInput` via:
//
//	FirewallSslSshProfileMap{ "key": FirewallSslSshProfileArgs{...} }
type FirewallSslSshProfileMapInput interface {
	pulumi.Input

	ToFirewallSslSshProfileMapOutput() FirewallSslSshProfileMapOutput
	ToFirewallSslSshProfileMapOutputWithContext(context.Context) FirewallSslSshProfileMapOutput
}

type FirewallSslSshProfileMap map[string]FirewallSslSshProfileInput

func (FirewallSslSshProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSslSshProfile)(nil)).Elem()
}

func (i FirewallSslSshProfileMap) ToFirewallSslSshProfileMapOutput() FirewallSslSshProfileMapOutput {
	return i.ToFirewallSslSshProfileMapOutputWithContext(context.Background())
}

func (i FirewallSslSshProfileMap) ToFirewallSslSshProfileMapOutputWithContext(ctx context.Context) FirewallSslSshProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslSshProfileMapOutput)
}

type FirewallSslSshProfileOutput struct{ *pulumi.OutputState }

func (FirewallSslSshProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSslSshProfile)(nil)).Elem()
}

func (o FirewallSslSshProfileOutput) ToFirewallSslSshProfileOutput() FirewallSslSshProfileOutput {
	return o
}

func (o FirewallSslSshProfileOutput) ToFirewallSslSshProfileOutputWithContext(ctx context.Context) FirewallSslSshProfileOutput {
	return o
}

func (o FirewallSslSshProfileOutput) Allowlist() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.Allowlist }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) BlockBlacklistedCertificates() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.BlockBlacklistedCertificates }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) BlockBlocklistedCertificates() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.BlockBlocklistedCertificates }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) Caname() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.Caname }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o FirewallSslSshProfileOutput) Dot() FirewallSslSshProfileDotOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileDotOutput { return v.Dot }).(FirewallSslSshProfileDotOutput)
}

func (o FirewallSslSshProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o FirewallSslSshProfileOutput) Ftps() FirewallSslSshProfileFtpsOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileFtpsOutput { return v.Ftps }).(FirewallSslSshProfileFtpsOutput)
}

func (o FirewallSslSshProfileOutput) Https() FirewallSslSshProfileHttpsOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileHttpsOutput { return v.Https }).(FirewallSslSshProfileHttpsOutput)
}

func (o FirewallSslSshProfileOutput) Imaps() FirewallSslSshProfileImapsOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileImapsOutput { return v.Imaps }).(FirewallSslSshProfileImapsOutput)
}

func (o FirewallSslSshProfileOutput) MapiOverHttps() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.MapiOverHttps }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) Pop3s() FirewallSslSshProfilePop3sOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfilePop3sOutput { return v.Pop3s }).(FirewallSslSshProfilePop3sOutput)
}

func (o FirewallSslSshProfileOutput) RpcOverHttps() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.RpcOverHttps }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) ServerCert() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.ServerCert }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) ServerCertMode() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.ServerCertMode }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) Smtps() FirewallSslSshProfileSmtpsOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileSmtpsOutput { return v.Smtps }).(FirewallSslSshProfileSmtpsOutput)
}

func (o FirewallSslSshProfileOutput) Ssh() FirewallSslSshProfileSshOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileSshOutput { return v.Ssh }).(FirewallSslSshProfileSshOutput)
}

func (o FirewallSslSshProfileOutput) Ssl() FirewallSslSshProfileSslOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileSslOutput { return v.Ssl }).(FirewallSslSshProfileSslOutput)
}

func (o FirewallSslSshProfileOutput) SslAnomaliesLog() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SslAnomaliesLog }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) SslAnomalyLog() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SslAnomalyLog }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) SslExemptionIpRating() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SslExemptionIpRating }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) SslExemptionLog() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SslExemptionLog }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) SslExemptionsLog() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SslExemptionsLog }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) SslExempts() FirewallSslSshProfileSslExemptArrayOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileSslExemptArrayOutput { return v.SslExempts }).(FirewallSslSshProfileSslExemptArrayOutput)
}

func (o FirewallSslSshProfileOutput) SslHandshakeLog() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SslHandshakeLog }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) SslNegotiationLog() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SslNegotiationLog }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) SslServerCertLog() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SslServerCertLog }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) SslServers() FirewallSslSshProfileSslServerArrayOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) FirewallSslSshProfileSslServerArrayOutput { return v.SslServers }).(FirewallSslSshProfileSslServerArrayOutput)
}

func (o FirewallSslSshProfileOutput) SupportedAlpn() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.SupportedAlpn }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) UntrustedCaname() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.UntrustedCaname }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) UseSslServer() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.UseSslServer }).(pulumi.StringOutput)
}

func (o FirewallSslSshProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o FirewallSslSshProfileOutput) Whitelist() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslSshProfile) pulumi.StringOutput { return v.Whitelist }).(pulumi.StringOutput)
}

type FirewallSslSshProfileArrayOutput struct{ *pulumi.OutputState }

func (FirewallSslSshProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSslSshProfile)(nil)).Elem()
}

func (o FirewallSslSshProfileArrayOutput) ToFirewallSslSshProfileArrayOutput() FirewallSslSshProfileArrayOutput {
	return o
}

func (o FirewallSslSshProfileArrayOutput) ToFirewallSslSshProfileArrayOutputWithContext(ctx context.Context) FirewallSslSshProfileArrayOutput {
	return o
}

func (o FirewallSslSshProfileArrayOutput) Index(i pulumi.IntInput) FirewallSslSshProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSslSshProfile {
		return vs[0].([]*FirewallSslSshProfile)[vs[1].(int)]
	}).(FirewallSslSshProfileOutput)
}

type FirewallSslSshProfileMapOutput struct{ *pulumi.OutputState }

func (FirewallSslSshProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSslSshProfile)(nil)).Elem()
}

func (o FirewallSslSshProfileMapOutput) ToFirewallSslSshProfileMapOutput() FirewallSslSshProfileMapOutput {
	return o
}

func (o FirewallSslSshProfileMapOutput) ToFirewallSslSshProfileMapOutputWithContext(ctx context.Context) FirewallSslSshProfileMapOutput {
	return o
}

func (o FirewallSslSshProfileMapOutput) MapIndex(k pulumi.StringInput) FirewallSslSshProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSslSshProfile {
		return vs[0].(map[string]*FirewallSslSshProfile)[vs[1].(string)]
	}).(FirewallSslSshProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslSshProfileInput)(nil)).Elem(), &FirewallSslSshProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslSshProfileArrayInput)(nil)).Elem(), FirewallSslSshProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslSshProfileMapInput)(nil)).Elem(), FirewallSslSshProfileMap{})
	pulumi.RegisterOutputType(FirewallSslSshProfileOutput{})
	pulumi.RegisterOutputType(FirewallSslSshProfileArrayOutput{})
	pulumi.RegisterOutputType(FirewallSslSshProfileMapOutput{})
}
