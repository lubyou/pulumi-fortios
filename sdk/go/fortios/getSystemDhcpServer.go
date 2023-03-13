// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemDhcpServer(ctx *pulumi.Context, args *LookupSystemDhcpServerArgs, opts ...pulumi.InvokeOption) (*LookupSystemDhcpServerResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemDhcpServerResult
	err := ctx.Invoke("fortios:index/getSystemDhcpServer:GetSystemDhcpServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemDhcpServer.
type LookupSystemDhcpServerArgs struct {
	Fosid     int     `pulumi:"fosid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemDhcpServer.
type LookupSystemDhcpServerResult struct {
	AutoConfiguration         string                            `pulumi:"autoConfiguration"`
	AutoManagedStatus         string                            `pulumi:"autoManagedStatus"`
	ConflictedIpTimeout       int                               `pulumi:"conflictedIpTimeout"`
	DdnsAuth                  string                            `pulumi:"ddnsAuth"`
	DdnsKey                   string                            `pulumi:"ddnsKey"`
	DdnsKeyname               string                            `pulumi:"ddnsKeyname"`
	DdnsServerIp              string                            `pulumi:"ddnsServerIp"`
	DdnsTtl                   int                               `pulumi:"ddnsTtl"`
	DdnsUpdate                string                            `pulumi:"ddnsUpdate"`
	DdnsUpdateOverride        string                            `pulumi:"ddnsUpdateOverride"`
	DdnsZone                  string                            `pulumi:"ddnsZone"`
	DefaultGateway            string                            `pulumi:"defaultGateway"`
	DhcpSettingsFromFortiipam string                            `pulumi:"dhcpSettingsFromFortiipam"`
	DnsServer1                string                            `pulumi:"dnsServer1"`
	DnsServer2                string                            `pulumi:"dnsServer2"`
	DnsServer3                string                            `pulumi:"dnsServer3"`
	DnsServer4                string                            `pulumi:"dnsServer4"`
	DnsService                string                            `pulumi:"dnsService"`
	Domain                    string                            `pulumi:"domain"`
	ExcludeRanges             []GetSystemDhcpServerExcludeRange `pulumi:"excludeRanges"`
	Filename                  string                            `pulumi:"filename"`
	ForticlientOnNetStatus    string                            `pulumi:"forticlientOnNetStatus"`
	Fosid                     int                               `pulumi:"fosid"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                               `pulumi:"id"`
	Interface           string                               `pulumi:"interface"`
	IpMode              string                               `pulumi:"ipMode"`
	IpRanges            []GetSystemDhcpServerIpRange         `pulumi:"ipRanges"`
	IpsecLeaseHold      int                                  `pulumi:"ipsecLeaseHold"`
	LeaseTime           int                                  `pulumi:"leaseTime"`
	MacAclDefaultAction string                               `pulumi:"macAclDefaultAction"`
	Netmask             string                               `pulumi:"netmask"`
	NextServer          string                               `pulumi:"nextServer"`
	NtpServer1          string                               `pulumi:"ntpServer1"`
	NtpServer2          string                               `pulumi:"ntpServer2"`
	NtpServer3          string                               `pulumi:"ntpServer3"`
	NtpService          string                               `pulumi:"ntpService"`
	Options             []GetSystemDhcpServerOption          `pulumi:"options"`
	ReservedAddresses   []GetSystemDhcpServerReservedAddress `pulumi:"reservedAddresses"`
	ServerType          string                               `pulumi:"serverType"`
	Status              string                               `pulumi:"status"`
	TftpServers         []GetSystemDhcpServerTftpServer      `pulumi:"tftpServers"`
	Timezone            string                               `pulumi:"timezone"`
	TimezoneOption      string                               `pulumi:"timezoneOption"`
	VciMatch            string                               `pulumi:"vciMatch"`
	VciStrings          []GetSystemDhcpServerVciString       `pulumi:"vciStrings"`
	Vdomparam           *string                              `pulumi:"vdomparam"`
	WifiAc1             string                               `pulumi:"wifiAc1"`
	WifiAc2             string                               `pulumi:"wifiAc2"`
	WifiAc3             string                               `pulumi:"wifiAc3"`
	WifiAcService       string                               `pulumi:"wifiAcService"`
	WinsServer1         string                               `pulumi:"winsServer1"`
	WinsServer2         string                               `pulumi:"winsServer2"`
}

func LookupSystemDhcpServerOutput(ctx *pulumi.Context, args LookupSystemDhcpServerOutputArgs, opts ...pulumi.InvokeOption) LookupSystemDhcpServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemDhcpServerResult, error) {
			args := v.(LookupSystemDhcpServerArgs)
			r, err := LookupSystemDhcpServer(ctx, &args, opts...)
			var s LookupSystemDhcpServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemDhcpServerResultOutput)
}

// A collection of arguments for invoking GetSystemDhcpServer.
type LookupSystemDhcpServerOutputArgs struct {
	Fosid     pulumi.IntInput       `pulumi:"fosid"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemDhcpServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDhcpServerArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemDhcpServer.
type LookupSystemDhcpServerResultOutput struct{ *pulumi.OutputState }

func (LookupSystemDhcpServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDhcpServerResult)(nil)).Elem()
}

func (o LookupSystemDhcpServerResultOutput) ToLookupSystemDhcpServerResultOutput() LookupSystemDhcpServerResultOutput {
	return o
}

func (o LookupSystemDhcpServerResultOutput) ToLookupSystemDhcpServerResultOutputWithContext(ctx context.Context) LookupSystemDhcpServerResultOutput {
	return o
}

func (o LookupSystemDhcpServerResultOutput) AutoConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.AutoConfiguration }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) AutoManagedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.AutoManagedStatus }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) ConflictedIpTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) int { return v.ConflictedIpTimeout }).(pulumi.IntOutput)
}

func (o LookupSystemDhcpServerResultOutput) DdnsAuth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DdnsAuth }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DdnsKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DdnsKey }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DdnsKeyname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DdnsKeyname }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DdnsServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DdnsServerIp }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DdnsTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) int { return v.DdnsTtl }).(pulumi.IntOutput)
}

func (o LookupSystemDhcpServerResultOutput) DdnsUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DdnsUpdate }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DdnsUpdateOverride() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DdnsUpdateOverride }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DdnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DdnsZone }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DefaultGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DefaultGateway }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DhcpSettingsFromFortiipam() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DhcpSettingsFromFortiipam }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DnsServer1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DnsServer1 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DnsServer2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DnsServer2 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DnsServer3() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DnsServer3 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DnsServer4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DnsServer4 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) DnsService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.DnsService }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) ExcludeRanges() GetSystemDhcpServerExcludeRangeArrayOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) []GetSystemDhcpServerExcludeRange { return v.ExcludeRanges }).(GetSystemDhcpServerExcludeRangeArrayOutput)
}

func (o LookupSystemDhcpServerResultOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.Filename }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) ForticlientOnNetStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.ForticlientOnNetStatus }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) int { return v.Fosid }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemDhcpServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.Interface }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) IpMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.IpMode }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) IpRanges() GetSystemDhcpServerIpRangeArrayOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) []GetSystemDhcpServerIpRange { return v.IpRanges }).(GetSystemDhcpServerIpRangeArrayOutput)
}

func (o LookupSystemDhcpServerResultOutput) IpsecLeaseHold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) int { return v.IpsecLeaseHold }).(pulumi.IntOutput)
}

func (o LookupSystemDhcpServerResultOutput) LeaseTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) int { return v.LeaseTime }).(pulumi.IntOutput)
}

func (o LookupSystemDhcpServerResultOutput) MacAclDefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.MacAclDefaultAction }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.Netmask }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) NextServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.NextServer }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) NtpServer1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.NtpServer1 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) NtpServer2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.NtpServer2 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) NtpServer3() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.NtpServer3 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) NtpService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.NtpService }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) Options() GetSystemDhcpServerOptionArrayOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) []GetSystemDhcpServerOption { return v.Options }).(GetSystemDhcpServerOptionArrayOutput)
}

func (o LookupSystemDhcpServerResultOutput) ReservedAddresses() GetSystemDhcpServerReservedAddressArrayOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) []GetSystemDhcpServerReservedAddress { return v.ReservedAddresses }).(GetSystemDhcpServerReservedAddressArrayOutput)
}

func (o LookupSystemDhcpServerResultOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.ServerType }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) TftpServers() GetSystemDhcpServerTftpServerArrayOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) []GetSystemDhcpServerTftpServer { return v.TftpServers }).(GetSystemDhcpServerTftpServerArrayOutput)
}

func (o LookupSystemDhcpServerResultOutput) Timezone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.Timezone }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) TimezoneOption() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.TimezoneOption }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) VciMatch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.VciMatch }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) VciStrings() GetSystemDhcpServerVciStringArrayOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) []GetSystemDhcpServerVciString { return v.VciStrings }).(GetSystemDhcpServerVciStringArrayOutput)
}

func (o LookupSystemDhcpServerResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupSystemDhcpServerResultOutput) WifiAc1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.WifiAc1 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) WifiAc2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.WifiAc2 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) WifiAc3() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.WifiAc3 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) WifiAcService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.WifiAcService }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) WinsServer1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.WinsServer1 }).(pulumi.StringOutput)
}

func (o LookupSystemDhcpServerResultOutput) WinsServer2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDhcpServerResult) string { return v.WinsServer2 }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemDhcpServerResultOutput{})
}
