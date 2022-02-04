// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewallservice custom
func LookupFirewallServiceCustom(ctx *pulumi.Context, args *LookupFirewallServiceCustomArgs, opts ...pulumi.InvokeOption) (*LookupFirewallServiceCustomResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallServiceCustomResult
	err := ctx.Invoke("fortios:index/getFirewallServiceCustom:GetFirewallServiceCustom", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallServiceCustom.
type LookupFirewallServiceCustomArgs struct {
	// Specify the name of the desired firewallservice custom.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallServiceCustom.
type LookupFirewallServiceCustomResult struct {
	// Application category ID. The structure of `appCategory` block is documented below.
	AppCategories []GetFirewallServiceCustomAppCategory `pulumi:"appCategories"`
	// Application service type.
	AppServiceType string `pulumi:"appServiceType"`
	// Application ID. The structure of `application` block is documented below.
	Applications []GetFirewallServiceCustomApplication `pulumi:"applications"`
	// Service category.
	Category string `pulumi:"category"`
	// Configure the type of ICMP error message verification.
	CheckResetRange string `pulumi:"checkResetRange"`
	// Color of icon on the GUI.
	Color int `pulumi:"color"`
	// Comment.
	Comment string `pulumi:"comment"`
	// Security Fabric global object setting.
	FabricObject string `pulumi:"fabricObject"`
	// Fully qualified domain name.
	Fqdn string `pulumi:"fqdn"`
	// Helper name.
	Helper string `pulumi:"helper"`
	// ICMP code.
	Icmpcode int `pulumi:"icmpcode"`
	// ICMP type.
	Icmptype int `pulumi:"icmptype"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Start and end of the IP range associated with service.
	Iprange string `pulumi:"iprange"`
	// Custom service name.
	Name string `pulumi:"name"`
	// Protocol type based on IANA numbers.
	Protocol string `pulumi:"protocol"`
	// IP protocol number.
	ProtocolNumber int `pulumi:"protocolNumber"`
	// Enable/disable web proxy service.
	Proxy string `pulumi:"proxy"`
	// Multiple SCTP port ranges.
	SctpPortrange string `pulumi:"sctpPortrange"`
	// Session TTL (300 - 604800, 0 = default).
	SessionTtl int `pulumi:"sessionTtl"`
	// Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
	TcpHalfcloseTimer int `pulumi:"tcpHalfcloseTimer"`
	// Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
	TcpHalfopenTimer int `pulumi:"tcpHalfopenTimer"`
	// Multiple TCP port ranges.
	TcpPortrange string `pulumi:"tcpPortrange"`
	// Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
	TcpRstTimer int `pulumi:"tcpRstTimer"`
	// Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
	TcpTimewaitTimer int `pulumi:"tcpTimewaitTimer"`
	// UDP half close timeout (0 - 86400 sec, 0 = default).
	UdpIdleTimer int `pulumi:"udpIdleTimer"`
	// Multiple UDP port ranges.
	UdpPortrange string  `pulumi:"udpPortrange"`
	Vdomparam    *string `pulumi:"vdomparam"`
	// Enable/disable the visibility of the service on the GUI.
	Visibility string `pulumi:"visibility"`
}

func LookupFirewallServiceCustomOutput(ctx *pulumi.Context, args LookupFirewallServiceCustomOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallServiceCustomResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallServiceCustomResult, error) {
			args := v.(LookupFirewallServiceCustomArgs)
			r, err := LookupFirewallServiceCustom(ctx, &args, opts...)
			return *r, err
		}).(LookupFirewallServiceCustomResultOutput)
}

// A collection of arguments for invoking GetFirewallServiceCustom.
type LookupFirewallServiceCustomOutputArgs struct {
	// Specify the name of the desired firewallservice custom.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallServiceCustomOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallServiceCustomArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallServiceCustom.
type LookupFirewallServiceCustomResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallServiceCustomResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallServiceCustomResult)(nil)).Elem()
}

func (o LookupFirewallServiceCustomResultOutput) ToLookupFirewallServiceCustomResultOutput() LookupFirewallServiceCustomResultOutput {
	return o
}

func (o LookupFirewallServiceCustomResultOutput) ToLookupFirewallServiceCustomResultOutputWithContext(ctx context.Context) LookupFirewallServiceCustomResultOutput {
	return o
}

// Application category ID. The structure of `appCategory` block is documented below.
func (o LookupFirewallServiceCustomResultOutput) AppCategories() GetFirewallServiceCustomAppCategoryArrayOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) []GetFirewallServiceCustomAppCategory {
		return v.AppCategories
	}).(GetFirewallServiceCustomAppCategoryArrayOutput)
}

// Application service type.
func (o LookupFirewallServiceCustomResultOutput) AppServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.AppServiceType }).(pulumi.StringOutput)
}

// Application ID. The structure of `application` block is documented below.
func (o LookupFirewallServiceCustomResultOutput) Applications() GetFirewallServiceCustomApplicationArrayOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) []GetFirewallServiceCustomApplication { return v.Applications }).(GetFirewallServiceCustomApplicationArrayOutput)
}

// Service category.
func (o LookupFirewallServiceCustomResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Category }).(pulumi.StringOutput)
}

// Configure the type of ICMP error message verification.
func (o LookupFirewallServiceCustomResultOutput) CheckResetRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.CheckResetRange }).(pulumi.StringOutput)
}

// Color of icon on the GUI.
func (o LookupFirewallServiceCustomResultOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.Color }).(pulumi.IntOutput)
}

// Comment.
func (o LookupFirewallServiceCustomResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Comment }).(pulumi.StringOutput)
}

// Security Fabric global object setting.
func (o LookupFirewallServiceCustomResultOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.FabricObject }).(pulumi.StringOutput)
}

// Fully qualified domain name.
func (o LookupFirewallServiceCustomResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// Helper name.
func (o LookupFirewallServiceCustomResultOutput) Helper() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Helper }).(pulumi.StringOutput)
}

// ICMP code.
func (o LookupFirewallServiceCustomResultOutput) Icmpcode() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.Icmpcode }).(pulumi.IntOutput)
}

// ICMP type.
func (o LookupFirewallServiceCustomResultOutput) Icmptype() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.Icmptype }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallServiceCustomResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Id }).(pulumi.StringOutput)
}

// Start and end of the IP range associated with service.
func (o LookupFirewallServiceCustomResultOutput) Iprange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Iprange }).(pulumi.StringOutput)
}

// Custom service name.
func (o LookupFirewallServiceCustomResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Name }).(pulumi.StringOutput)
}

// Protocol type based on IANA numbers.
func (o LookupFirewallServiceCustomResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// IP protocol number.
func (o LookupFirewallServiceCustomResultOutput) ProtocolNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.ProtocolNumber }).(pulumi.IntOutput)
}

// Enable/disable web proxy service.
func (o LookupFirewallServiceCustomResultOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Proxy }).(pulumi.StringOutput)
}

// Multiple SCTP port ranges.
func (o LookupFirewallServiceCustomResultOutput) SctpPortrange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.SctpPortrange }).(pulumi.StringOutput)
}

// Session TTL (300 - 604800, 0 = default).
func (o LookupFirewallServiceCustomResultOutput) SessionTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.SessionTtl }).(pulumi.IntOutput)
}

// Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
func (o LookupFirewallServiceCustomResultOutput) TcpHalfcloseTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.TcpHalfcloseTimer }).(pulumi.IntOutput)
}

// Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
func (o LookupFirewallServiceCustomResultOutput) TcpHalfopenTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.TcpHalfopenTimer }).(pulumi.IntOutput)
}

// Multiple TCP port ranges.
func (o LookupFirewallServiceCustomResultOutput) TcpPortrange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.TcpPortrange }).(pulumi.StringOutput)
}

// Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
func (o LookupFirewallServiceCustomResultOutput) TcpRstTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.TcpRstTimer }).(pulumi.IntOutput)
}

// Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
func (o LookupFirewallServiceCustomResultOutput) TcpTimewaitTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.TcpTimewaitTimer }).(pulumi.IntOutput)
}

// UDP half close timeout (0 - 86400 sec, 0 = default).
func (o LookupFirewallServiceCustomResultOutput) UdpIdleTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) int { return v.UdpIdleTimer }).(pulumi.IntOutput)
}

// Multiple UDP port ranges.
func (o LookupFirewallServiceCustomResultOutput) UdpPortrange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.UdpPortrange }).(pulumi.StringOutput)
}

func (o LookupFirewallServiceCustomResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable the visibility of the service on the GUI.
func (o LookupFirewallServiceCustomResultOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallServiceCustomResult) string { return v.Visibility }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallServiceCustomResultOutput{})
}
