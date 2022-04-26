// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure explicit Web proxy settings.
//
// ## Import
//
// WebProxy Explicit can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/webProxyExplicit:WebProxyExplicit labelname WebProxyExplicit
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/webProxyExplicit:WebProxyExplicit labelname WebProxyExplicit
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type WebProxyExplicit struct {
	pulumi.CustomResourceState

	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	FtpIncomingPort pulumi.StringOutput `pulumi:"ftpIncomingPort"`
	// Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
	FtpOverHttp pulumi.StringOutput `pulumi:"ftpOverHttp"`
	// Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
	HttpIncomingPort pulumi.StringOutput `pulumi:"httpIncomingPort"`
	// Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
	HttpsIncomingPort pulumi.StringOutput `pulumi:"httpsIncomingPort"`
	// Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
	HttpsReplacementMessage pulumi.StringOutput `pulumi:"httpsReplacementMessage"`
	// Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringOutput `pulumi:"incomingIp"`
	// Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
	IncomingIp6 pulumi.StringOutput `pulumi:"incomingIp6"`
	// Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
	Ipv6Status pulumi.StringOutput `pulumi:"ipv6Status"`
	// Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
	MessageUponServerError pulumi.StringOutput `pulumi:"messageUponServerError"`
	// Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
	OutgoingIp pulumi.StringOutput `pulumi:"outgoingIp"`
	// Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
	OutgoingIp6 pulumi.StringOutput `pulumi:"outgoingIp6"`
	// PAC file contents enclosed in quotes (maximum of 256K bytes).
	PacFileData pulumi.StringOutput `pulumi:"pacFileData"`
	// Pac file name.
	PacFileName pulumi.StringOutput `pulumi:"pacFileName"`
	// Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
	PacFileServerPort pulumi.StringOutput `pulumi:"pacFileServerPort"`
	// Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
	PacFileServerStatus pulumi.StringOutput `pulumi:"pacFileServerStatus"`
	// PAC file access URL.
	PacFileUrl pulumi.StringOutput `pulumi:"pacFileUrl"`
	// PAC policies. The structure of `pacPolicy` block is documented below.
	PacPolicies WebProxyExplicitPacPolicyArrayOutput `pulumi:"pacPolicies"`
	// Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
	PrefDnsResult pulumi.StringOutput `pulumi:"prefDnsResult"`
	// Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringOutput `pulumi:"secDefaultAction"`
	// Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
	Socks pulumi.StringOutput `pulumi:"socks"`
	// Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	SocksIncomingPort pulumi.StringOutput `pulumi:"socksIncomingPort"`
	// Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringOutput `pulumi:"sslAlgorithm"`
	// Enable/disable policy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
	StrictGuest pulumi.StringOutput `pulumi:"strictGuest"`
	// Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
	TraceAuthNoRsp pulumi.StringOutput `pulumi:"traceAuthNoRsp"`
	// Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
	UnknownHttpVersion pulumi.StringOutput `pulumi:"unknownHttpVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebProxyExplicit registers a new resource with the given unique name, arguments, and options.
func NewWebProxyExplicit(ctx *pulumi.Context,
	name string, args *WebProxyExplicitArgs, opts ...pulumi.ResourceOption) (*WebProxyExplicit, error) {
	if args == nil {
		args = &WebProxyExplicitArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
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
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	FtpIncomingPort *string `pulumi:"ftpIncomingPort"`
	// Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
	FtpOverHttp *string `pulumi:"ftpOverHttp"`
	// Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
	HttpIncomingPort *string `pulumi:"httpIncomingPort"`
	// Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
	HttpsIncomingPort *string `pulumi:"httpsIncomingPort"`
	// Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
	HttpsReplacementMessage *string `pulumi:"httpsReplacementMessage"`
	// Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
	IncomingIp *string `pulumi:"incomingIp"`
	// Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
	IncomingIp6 *string `pulumi:"incomingIp6"`
	// Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
	Ipv6Status *string `pulumi:"ipv6Status"`
	// Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
	MessageUponServerError *string `pulumi:"messageUponServerError"`
	// Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
	OutgoingIp *string `pulumi:"outgoingIp"`
	// Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
	OutgoingIp6 *string `pulumi:"outgoingIp6"`
	// PAC file contents enclosed in quotes (maximum of 256K bytes).
	PacFileData *string `pulumi:"pacFileData"`
	// Pac file name.
	PacFileName *string `pulumi:"pacFileName"`
	// Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
	PacFileServerPort *string `pulumi:"pacFileServerPort"`
	// Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
	PacFileServerStatus *string `pulumi:"pacFileServerStatus"`
	// PAC file access URL.
	PacFileUrl *string `pulumi:"pacFileUrl"`
	// PAC policies. The structure of `pacPolicy` block is documented below.
	PacPolicies []WebProxyExplicitPacPolicy `pulumi:"pacPolicies"`
	// Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
	PrefDnsResult *string `pulumi:"prefDnsResult"`
	// Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
	Realm *string `pulumi:"realm"`
	// Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction *string `pulumi:"secDefaultAction"`
	// Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
	Socks *string `pulumi:"socks"`
	// Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	SocksIncomingPort *string `pulumi:"socksIncomingPort"`
	// Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Enable/disable policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
	StrictGuest *string `pulumi:"strictGuest"`
	// Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
	TraceAuthNoRsp *string `pulumi:"traceAuthNoRsp"`
	// Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
	UnknownHttpVersion *string `pulumi:"unknownHttpVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebProxyExplicitState struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	FtpIncomingPort pulumi.StringPtrInput
	// Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
	FtpOverHttp pulumi.StringPtrInput
	// Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
	HttpIncomingPort pulumi.StringPtrInput
	// Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
	HttpsIncomingPort pulumi.StringPtrInput
	// Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
	HttpsReplacementMessage pulumi.StringPtrInput
	// Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringPtrInput
	// Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
	IncomingIp6 pulumi.StringPtrInput
	// Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
	Ipv6Status pulumi.StringPtrInput
	// Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
	MessageUponServerError pulumi.StringPtrInput
	// Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
	OutgoingIp pulumi.StringPtrInput
	// Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
	OutgoingIp6 pulumi.StringPtrInput
	// PAC file contents enclosed in quotes (maximum of 256K bytes).
	PacFileData pulumi.StringPtrInput
	// Pac file name.
	PacFileName pulumi.StringPtrInput
	// Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
	PacFileServerPort pulumi.StringPtrInput
	// Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
	PacFileServerStatus pulumi.StringPtrInput
	// PAC file access URL.
	PacFileUrl pulumi.StringPtrInput
	// PAC policies. The structure of `pacPolicy` block is documented below.
	PacPolicies WebProxyExplicitPacPolicyArrayInput
	// Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
	PrefDnsResult pulumi.StringPtrInput
	// Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
	Realm pulumi.StringPtrInput
	// Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringPtrInput
	// Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
	Socks pulumi.StringPtrInput
	// Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	SocksIncomingPort pulumi.StringPtrInput
	// Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Enable/disable policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
	StrictGuest pulumi.StringPtrInput
	// Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
	TraceAuthNoRsp pulumi.StringPtrInput
	// Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
	UnknownHttpVersion pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebProxyExplicitState) ElementType() reflect.Type {
	return reflect.TypeOf((*webProxyExplicitState)(nil)).Elem()
}

type webProxyExplicitArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	FtpIncomingPort *string `pulumi:"ftpIncomingPort"`
	// Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
	FtpOverHttp *string `pulumi:"ftpOverHttp"`
	// Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
	HttpIncomingPort *string `pulumi:"httpIncomingPort"`
	// Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
	HttpsIncomingPort *string `pulumi:"httpsIncomingPort"`
	// Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
	HttpsReplacementMessage *string `pulumi:"httpsReplacementMessage"`
	// Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
	IncomingIp *string `pulumi:"incomingIp"`
	// Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
	IncomingIp6 *string `pulumi:"incomingIp6"`
	// Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
	Ipv6Status *string `pulumi:"ipv6Status"`
	// Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
	MessageUponServerError *string `pulumi:"messageUponServerError"`
	// Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
	OutgoingIp *string `pulumi:"outgoingIp"`
	// Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
	OutgoingIp6 *string `pulumi:"outgoingIp6"`
	// PAC file contents enclosed in quotes (maximum of 256K bytes).
	PacFileData *string `pulumi:"pacFileData"`
	// Pac file name.
	PacFileName *string `pulumi:"pacFileName"`
	// Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
	PacFileServerPort *string `pulumi:"pacFileServerPort"`
	// Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
	PacFileServerStatus *string `pulumi:"pacFileServerStatus"`
	// PAC file access URL.
	PacFileUrl *string `pulumi:"pacFileUrl"`
	// PAC policies. The structure of `pacPolicy` block is documented below.
	PacPolicies []WebProxyExplicitPacPolicy `pulumi:"pacPolicies"`
	// Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
	PrefDnsResult *string `pulumi:"prefDnsResult"`
	// Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
	Realm *string `pulumi:"realm"`
	// Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction *string `pulumi:"secDefaultAction"`
	// Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
	Socks *string `pulumi:"socks"`
	// Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	SocksIncomingPort *string `pulumi:"socksIncomingPort"`
	// Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Enable/disable policy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
	StrictGuest *string `pulumi:"strictGuest"`
	// Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
	TraceAuthNoRsp *string `pulumi:"traceAuthNoRsp"`
	// Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
	UnknownHttpVersion *string `pulumi:"unknownHttpVersion"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebProxyExplicit resource.
type WebProxyExplicitArgs struct {
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Accept incoming FTP-over-HTTP requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	FtpIncomingPort pulumi.StringPtrInput
	// Enable to proxy FTP-over-HTTP sessions sent from a web browser. Valid values: `enable`, `disable`.
	FtpOverHttp pulumi.StringPtrInput
	// Accept incoming HTTP requests on one or more ports (0 - 65535, default = 8080).
	HttpIncomingPort pulumi.StringPtrInput
	// Accept incoming HTTPS requests on one or more ports (0 - 65535, default = 0, use the same as HTTP).
	HttpsIncomingPort pulumi.StringPtrInput
	// Enable/disable sending the client a replacement message for HTTPS requests. Valid values: `enable`, `disable`.
	HttpsReplacementMessage pulumi.StringPtrInput
	// Restrict the explicit HTTP proxy to only accept sessions from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringPtrInput
	// Restrict the explicit web proxy to only accept sessions from this IPv6 address. An interface must have this IPv6 address.
	IncomingIp6 pulumi.StringPtrInput
	// Enable/disable allowing an IPv6 web proxy destination in policies and all IPv6 related entries in this command. Valid values: `enable`, `disable`.
	Ipv6Status pulumi.StringPtrInput
	// Enable/disable displaying a replacement message when a server error is detected. Valid values: `enable`, `disable`.
	MessageUponServerError pulumi.StringPtrInput
	// Outgoing HTTP requests will have this IP address as their source address. An interface must have this IP address.
	OutgoingIp pulumi.StringPtrInput
	// Outgoing HTTP requests will leave this IPv6. Multiple interfaces can be specified. Interfaces must have these IPv6 addresses.
	OutgoingIp6 pulumi.StringPtrInput
	// PAC file contents enclosed in quotes (maximum of 256K bytes).
	PacFileData pulumi.StringPtrInput
	// Pac file name.
	PacFileName pulumi.StringPtrInput
	// Port number that PAC traffic from client web browsers uses to connect to the explicit web proxy (0 - 65535, default = 0; use the same as HTTP).
	PacFileServerPort pulumi.StringPtrInput
	// Enable/disable Proxy Auto-Configuration (PAC) for users of this explicit proxy profile. Valid values: `enable`, `disable`.
	PacFileServerStatus pulumi.StringPtrInput
	// PAC file access URL.
	PacFileUrl pulumi.StringPtrInput
	// PAC policies. The structure of `pacPolicy` block is documented below.
	PacPolicies WebProxyExplicitPacPolicyArrayInput
	// Prefer resolving addresses using the configured IPv4 or IPv6 DNS server (default = ipv4). Valid values: `ipv4`, `ipv6`.
	PrefDnsResult pulumi.StringPtrInput
	// Authentication realm used to identify the explicit web proxy (maximum of 63 characters).
	Realm pulumi.StringPtrInput
	// Accept or deny explicit web proxy sessions when no web proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringPtrInput
	// Enable/disable the SOCKS proxy. Valid values: `enable`, `disable`.
	Socks pulumi.StringPtrInput
	// Accept incoming SOCKS proxy requests on one or more ports (0 - 65535, default = 0; use the same as HTTP).
	SocksIncomingPort pulumi.StringPtrInput
	// Relative strength of encryption algorithms accepted in HTTPS deep scan: high, medium, or low. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Enable/disable policy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable strict guest user checking by the explicit web proxy. Valid values: `enable`, `disable`.
	StrictGuest pulumi.StringPtrInput
	// Enable/disable logging timed-out authentication requests. Valid values: `enable`, `disable`.
	TraceAuthNoRsp pulumi.StringPtrInput
	// Either reject unknown HTTP traffic as malformed or handle unknown HTTP traffic as best as the proxy server can.
	UnknownHttpVersion pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
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
//          WebProxyExplicitArray{ WebProxyExplicitArgs{...} }
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
//          WebProxyExplicitMap{ "key": WebProxyExplicitArgs{...} }
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
