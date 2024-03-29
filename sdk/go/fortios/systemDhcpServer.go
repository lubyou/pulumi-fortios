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

type SystemDhcpServer struct {
	pulumi.CustomResourceState

	AutoConfiguration         pulumi.StringOutput                        `pulumi:"autoConfiguration"`
	AutoManagedStatus         pulumi.StringOutput                        `pulumi:"autoManagedStatus"`
	ConflictedIpTimeout       pulumi.IntOutput                           `pulumi:"conflictedIpTimeout"`
	DdnsAuth                  pulumi.StringOutput                        `pulumi:"ddnsAuth"`
	DdnsKey                   pulumi.StringOutput                        `pulumi:"ddnsKey"`
	DdnsKeyname               pulumi.StringOutput                        `pulumi:"ddnsKeyname"`
	DdnsServerIp              pulumi.StringOutput                        `pulumi:"ddnsServerIp"`
	DdnsTtl                   pulumi.IntOutput                           `pulumi:"ddnsTtl"`
	DdnsUpdate                pulumi.StringOutput                        `pulumi:"ddnsUpdate"`
	DdnsUpdateOverride        pulumi.StringOutput                        `pulumi:"ddnsUpdateOverride"`
	DdnsZone                  pulumi.StringOutput                        `pulumi:"ddnsZone"`
	DefaultGateway            pulumi.StringOutput                        `pulumi:"defaultGateway"`
	DhcpSettingsFromFortiipam pulumi.StringOutput                        `pulumi:"dhcpSettingsFromFortiipam"`
	DnsServer1                pulumi.StringOutput                        `pulumi:"dnsServer1"`
	DnsServer2                pulumi.StringOutput                        `pulumi:"dnsServer2"`
	DnsServer3                pulumi.StringOutput                        `pulumi:"dnsServer3"`
	DnsServer4                pulumi.StringOutput                        `pulumi:"dnsServer4"`
	DnsService                pulumi.StringOutput                        `pulumi:"dnsService"`
	Domain                    pulumi.StringOutput                        `pulumi:"domain"`
	DynamicSortSubtable       pulumi.StringPtrOutput                     `pulumi:"dynamicSortSubtable"`
	ExcludeRanges             SystemDhcpServerExcludeRangeArrayOutput    `pulumi:"excludeRanges"`
	Filename                  pulumi.StringOutput                        `pulumi:"filename"`
	ForticlientOnNetStatus    pulumi.StringOutput                        `pulumi:"forticlientOnNetStatus"`
	Fosid                     pulumi.IntOutput                           `pulumi:"fosid"`
	GetAllTables              pulumi.StringPtrOutput                     `pulumi:"getAllTables"`
	Interface                 pulumi.StringOutput                        `pulumi:"interface"`
	IpMode                    pulumi.StringOutput                        `pulumi:"ipMode"`
	IpRanges                  SystemDhcpServerIpRangeArrayOutput         `pulumi:"ipRanges"`
	IpsecLeaseHold            pulumi.IntOutput                           `pulumi:"ipsecLeaseHold"`
	LeaseTime                 pulumi.IntOutput                           `pulumi:"leaseTime"`
	MacAclDefaultAction       pulumi.StringOutput                        `pulumi:"macAclDefaultAction"`
	Netmask                   pulumi.StringOutput                        `pulumi:"netmask"`
	NextServer                pulumi.StringOutput                        `pulumi:"nextServer"`
	NtpServer1                pulumi.StringOutput                        `pulumi:"ntpServer1"`
	NtpServer2                pulumi.StringOutput                        `pulumi:"ntpServer2"`
	NtpServer3                pulumi.StringOutput                        `pulumi:"ntpServer3"`
	NtpService                pulumi.StringOutput                        `pulumi:"ntpService"`
	Options                   SystemDhcpServerOptionArrayOutput          `pulumi:"options"`
	RelayAgent                pulumi.StringOutput                        `pulumi:"relayAgent"`
	ReservedAddresses         SystemDhcpServerReservedAddressArrayOutput `pulumi:"reservedAddresses"`
	ServerType                pulumi.StringOutput                        `pulumi:"serverType"`
	SharedSubnet              pulumi.StringOutput                        `pulumi:"sharedSubnet"`
	Status                    pulumi.StringOutput                        `pulumi:"status"`
	TftpServers               SystemDhcpServerTftpServerArrayOutput      `pulumi:"tftpServers"`
	Timezone                  pulumi.StringOutput                        `pulumi:"timezone"`
	TimezoneOption            pulumi.StringOutput                        `pulumi:"timezoneOption"`
	VciMatch                  pulumi.StringOutput                        `pulumi:"vciMatch"`
	VciStrings                SystemDhcpServerVciStringArrayOutput       `pulumi:"vciStrings"`
	Vdomparam                 pulumi.StringPtrOutput                     `pulumi:"vdomparam"`
	WifiAc1                   pulumi.StringOutput                        `pulumi:"wifiAc1"`
	WifiAc2                   pulumi.StringOutput                        `pulumi:"wifiAc2"`
	WifiAc3                   pulumi.StringOutput                        `pulumi:"wifiAc3"`
	WifiAcService             pulumi.StringOutput                        `pulumi:"wifiAcService"`
	WinsServer1               pulumi.StringOutput                        `pulumi:"winsServer1"`
	WinsServer2               pulumi.StringOutput                        `pulumi:"winsServer2"`
}

// NewSystemDhcpServer registers a new resource with the given unique name, arguments, and options.
func NewSystemDhcpServer(ctx *pulumi.Context,
	name string, args *SystemDhcpServerArgs, opts ...pulumi.ResourceOption) (*SystemDhcpServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.Netmask == nil {
		return nil, errors.New("invalid value for required argument 'Netmask'")
	}
	if args.DdnsKey != nil {
		args.DdnsKey = pulumi.ToSecret(args.DdnsKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"ddnsKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SystemDhcpServer
	err := ctx.RegisterResource("fortios:index/systemDhcpServer:SystemDhcpServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemDhcpServer gets an existing SystemDhcpServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemDhcpServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemDhcpServerState, opts ...pulumi.ResourceOption) (*SystemDhcpServer, error) {
	var resource SystemDhcpServer
	err := ctx.ReadResource("fortios:index/systemDhcpServer:SystemDhcpServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemDhcpServer resources.
type systemDhcpServerState struct {
	AutoConfiguration         *string                           `pulumi:"autoConfiguration"`
	AutoManagedStatus         *string                           `pulumi:"autoManagedStatus"`
	ConflictedIpTimeout       *int                              `pulumi:"conflictedIpTimeout"`
	DdnsAuth                  *string                           `pulumi:"ddnsAuth"`
	DdnsKey                   *string                           `pulumi:"ddnsKey"`
	DdnsKeyname               *string                           `pulumi:"ddnsKeyname"`
	DdnsServerIp              *string                           `pulumi:"ddnsServerIp"`
	DdnsTtl                   *int                              `pulumi:"ddnsTtl"`
	DdnsUpdate                *string                           `pulumi:"ddnsUpdate"`
	DdnsUpdateOverride        *string                           `pulumi:"ddnsUpdateOverride"`
	DdnsZone                  *string                           `pulumi:"ddnsZone"`
	DefaultGateway            *string                           `pulumi:"defaultGateway"`
	DhcpSettingsFromFortiipam *string                           `pulumi:"dhcpSettingsFromFortiipam"`
	DnsServer1                *string                           `pulumi:"dnsServer1"`
	DnsServer2                *string                           `pulumi:"dnsServer2"`
	DnsServer3                *string                           `pulumi:"dnsServer3"`
	DnsServer4                *string                           `pulumi:"dnsServer4"`
	DnsService                *string                           `pulumi:"dnsService"`
	Domain                    *string                           `pulumi:"domain"`
	DynamicSortSubtable       *string                           `pulumi:"dynamicSortSubtable"`
	ExcludeRanges             []SystemDhcpServerExcludeRange    `pulumi:"excludeRanges"`
	Filename                  *string                           `pulumi:"filename"`
	ForticlientOnNetStatus    *string                           `pulumi:"forticlientOnNetStatus"`
	Fosid                     *int                              `pulumi:"fosid"`
	GetAllTables              *string                           `pulumi:"getAllTables"`
	Interface                 *string                           `pulumi:"interface"`
	IpMode                    *string                           `pulumi:"ipMode"`
	IpRanges                  []SystemDhcpServerIpRange         `pulumi:"ipRanges"`
	IpsecLeaseHold            *int                              `pulumi:"ipsecLeaseHold"`
	LeaseTime                 *int                              `pulumi:"leaseTime"`
	MacAclDefaultAction       *string                           `pulumi:"macAclDefaultAction"`
	Netmask                   *string                           `pulumi:"netmask"`
	NextServer                *string                           `pulumi:"nextServer"`
	NtpServer1                *string                           `pulumi:"ntpServer1"`
	NtpServer2                *string                           `pulumi:"ntpServer2"`
	NtpServer3                *string                           `pulumi:"ntpServer3"`
	NtpService                *string                           `pulumi:"ntpService"`
	Options                   []SystemDhcpServerOption          `pulumi:"options"`
	RelayAgent                *string                           `pulumi:"relayAgent"`
	ReservedAddresses         []SystemDhcpServerReservedAddress `pulumi:"reservedAddresses"`
	ServerType                *string                           `pulumi:"serverType"`
	SharedSubnet              *string                           `pulumi:"sharedSubnet"`
	Status                    *string                           `pulumi:"status"`
	TftpServers               []SystemDhcpServerTftpServer      `pulumi:"tftpServers"`
	Timezone                  *string                           `pulumi:"timezone"`
	TimezoneOption            *string                           `pulumi:"timezoneOption"`
	VciMatch                  *string                           `pulumi:"vciMatch"`
	VciStrings                []SystemDhcpServerVciString       `pulumi:"vciStrings"`
	Vdomparam                 *string                           `pulumi:"vdomparam"`
	WifiAc1                   *string                           `pulumi:"wifiAc1"`
	WifiAc2                   *string                           `pulumi:"wifiAc2"`
	WifiAc3                   *string                           `pulumi:"wifiAc3"`
	WifiAcService             *string                           `pulumi:"wifiAcService"`
	WinsServer1               *string                           `pulumi:"winsServer1"`
	WinsServer2               *string                           `pulumi:"winsServer2"`
}

type SystemDhcpServerState struct {
	AutoConfiguration         pulumi.StringPtrInput
	AutoManagedStatus         pulumi.StringPtrInput
	ConflictedIpTimeout       pulumi.IntPtrInput
	DdnsAuth                  pulumi.StringPtrInput
	DdnsKey                   pulumi.StringPtrInput
	DdnsKeyname               pulumi.StringPtrInput
	DdnsServerIp              pulumi.StringPtrInput
	DdnsTtl                   pulumi.IntPtrInput
	DdnsUpdate                pulumi.StringPtrInput
	DdnsUpdateOverride        pulumi.StringPtrInput
	DdnsZone                  pulumi.StringPtrInput
	DefaultGateway            pulumi.StringPtrInput
	DhcpSettingsFromFortiipam pulumi.StringPtrInput
	DnsServer1                pulumi.StringPtrInput
	DnsServer2                pulumi.StringPtrInput
	DnsServer3                pulumi.StringPtrInput
	DnsServer4                pulumi.StringPtrInput
	DnsService                pulumi.StringPtrInput
	Domain                    pulumi.StringPtrInput
	DynamicSortSubtable       pulumi.StringPtrInput
	ExcludeRanges             SystemDhcpServerExcludeRangeArrayInput
	Filename                  pulumi.StringPtrInput
	ForticlientOnNetStatus    pulumi.StringPtrInput
	Fosid                     pulumi.IntPtrInput
	GetAllTables              pulumi.StringPtrInput
	Interface                 pulumi.StringPtrInput
	IpMode                    pulumi.StringPtrInput
	IpRanges                  SystemDhcpServerIpRangeArrayInput
	IpsecLeaseHold            pulumi.IntPtrInput
	LeaseTime                 pulumi.IntPtrInput
	MacAclDefaultAction       pulumi.StringPtrInput
	Netmask                   pulumi.StringPtrInput
	NextServer                pulumi.StringPtrInput
	NtpServer1                pulumi.StringPtrInput
	NtpServer2                pulumi.StringPtrInput
	NtpServer3                pulumi.StringPtrInput
	NtpService                pulumi.StringPtrInput
	Options                   SystemDhcpServerOptionArrayInput
	RelayAgent                pulumi.StringPtrInput
	ReservedAddresses         SystemDhcpServerReservedAddressArrayInput
	ServerType                pulumi.StringPtrInput
	SharedSubnet              pulumi.StringPtrInput
	Status                    pulumi.StringPtrInput
	TftpServers               SystemDhcpServerTftpServerArrayInput
	Timezone                  pulumi.StringPtrInput
	TimezoneOption            pulumi.StringPtrInput
	VciMatch                  pulumi.StringPtrInput
	VciStrings                SystemDhcpServerVciStringArrayInput
	Vdomparam                 pulumi.StringPtrInput
	WifiAc1                   pulumi.StringPtrInput
	WifiAc2                   pulumi.StringPtrInput
	WifiAc3                   pulumi.StringPtrInput
	WifiAcService             pulumi.StringPtrInput
	WinsServer1               pulumi.StringPtrInput
	WinsServer2               pulumi.StringPtrInput
}

func (SystemDhcpServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDhcpServerState)(nil)).Elem()
}

type systemDhcpServerArgs struct {
	AutoConfiguration         *string                           `pulumi:"autoConfiguration"`
	AutoManagedStatus         *string                           `pulumi:"autoManagedStatus"`
	ConflictedIpTimeout       *int                              `pulumi:"conflictedIpTimeout"`
	DdnsAuth                  *string                           `pulumi:"ddnsAuth"`
	DdnsKey                   *string                           `pulumi:"ddnsKey"`
	DdnsKeyname               *string                           `pulumi:"ddnsKeyname"`
	DdnsServerIp              *string                           `pulumi:"ddnsServerIp"`
	DdnsTtl                   *int                              `pulumi:"ddnsTtl"`
	DdnsUpdate                *string                           `pulumi:"ddnsUpdate"`
	DdnsUpdateOverride        *string                           `pulumi:"ddnsUpdateOverride"`
	DdnsZone                  *string                           `pulumi:"ddnsZone"`
	DefaultGateway            *string                           `pulumi:"defaultGateway"`
	DhcpSettingsFromFortiipam *string                           `pulumi:"dhcpSettingsFromFortiipam"`
	DnsServer1                *string                           `pulumi:"dnsServer1"`
	DnsServer2                *string                           `pulumi:"dnsServer2"`
	DnsServer3                *string                           `pulumi:"dnsServer3"`
	DnsServer4                *string                           `pulumi:"dnsServer4"`
	DnsService                *string                           `pulumi:"dnsService"`
	Domain                    *string                           `pulumi:"domain"`
	DynamicSortSubtable       *string                           `pulumi:"dynamicSortSubtable"`
	ExcludeRanges             []SystemDhcpServerExcludeRange    `pulumi:"excludeRanges"`
	Filename                  *string                           `pulumi:"filename"`
	ForticlientOnNetStatus    *string                           `pulumi:"forticlientOnNetStatus"`
	Fosid                     *int                              `pulumi:"fosid"`
	GetAllTables              *string                           `pulumi:"getAllTables"`
	Interface                 string                            `pulumi:"interface"`
	IpMode                    *string                           `pulumi:"ipMode"`
	IpRanges                  []SystemDhcpServerIpRange         `pulumi:"ipRanges"`
	IpsecLeaseHold            *int                              `pulumi:"ipsecLeaseHold"`
	LeaseTime                 *int                              `pulumi:"leaseTime"`
	MacAclDefaultAction       *string                           `pulumi:"macAclDefaultAction"`
	Netmask                   string                            `pulumi:"netmask"`
	NextServer                *string                           `pulumi:"nextServer"`
	NtpServer1                *string                           `pulumi:"ntpServer1"`
	NtpServer2                *string                           `pulumi:"ntpServer2"`
	NtpServer3                *string                           `pulumi:"ntpServer3"`
	NtpService                *string                           `pulumi:"ntpService"`
	Options                   []SystemDhcpServerOption          `pulumi:"options"`
	RelayAgent                *string                           `pulumi:"relayAgent"`
	ReservedAddresses         []SystemDhcpServerReservedAddress `pulumi:"reservedAddresses"`
	ServerType                *string                           `pulumi:"serverType"`
	SharedSubnet              *string                           `pulumi:"sharedSubnet"`
	Status                    *string                           `pulumi:"status"`
	TftpServers               []SystemDhcpServerTftpServer      `pulumi:"tftpServers"`
	Timezone                  *string                           `pulumi:"timezone"`
	TimezoneOption            *string                           `pulumi:"timezoneOption"`
	VciMatch                  *string                           `pulumi:"vciMatch"`
	VciStrings                []SystemDhcpServerVciString       `pulumi:"vciStrings"`
	Vdomparam                 *string                           `pulumi:"vdomparam"`
	WifiAc1                   *string                           `pulumi:"wifiAc1"`
	WifiAc2                   *string                           `pulumi:"wifiAc2"`
	WifiAc3                   *string                           `pulumi:"wifiAc3"`
	WifiAcService             *string                           `pulumi:"wifiAcService"`
	WinsServer1               *string                           `pulumi:"winsServer1"`
	WinsServer2               *string                           `pulumi:"winsServer2"`
}

// The set of arguments for constructing a SystemDhcpServer resource.
type SystemDhcpServerArgs struct {
	AutoConfiguration         pulumi.StringPtrInput
	AutoManagedStatus         pulumi.StringPtrInput
	ConflictedIpTimeout       pulumi.IntPtrInput
	DdnsAuth                  pulumi.StringPtrInput
	DdnsKey                   pulumi.StringPtrInput
	DdnsKeyname               pulumi.StringPtrInput
	DdnsServerIp              pulumi.StringPtrInput
	DdnsTtl                   pulumi.IntPtrInput
	DdnsUpdate                pulumi.StringPtrInput
	DdnsUpdateOverride        pulumi.StringPtrInput
	DdnsZone                  pulumi.StringPtrInput
	DefaultGateway            pulumi.StringPtrInput
	DhcpSettingsFromFortiipam pulumi.StringPtrInput
	DnsServer1                pulumi.StringPtrInput
	DnsServer2                pulumi.StringPtrInput
	DnsServer3                pulumi.StringPtrInput
	DnsServer4                pulumi.StringPtrInput
	DnsService                pulumi.StringPtrInput
	Domain                    pulumi.StringPtrInput
	DynamicSortSubtable       pulumi.StringPtrInput
	ExcludeRanges             SystemDhcpServerExcludeRangeArrayInput
	Filename                  pulumi.StringPtrInput
	ForticlientOnNetStatus    pulumi.StringPtrInput
	Fosid                     pulumi.IntPtrInput
	GetAllTables              pulumi.StringPtrInput
	Interface                 pulumi.StringInput
	IpMode                    pulumi.StringPtrInput
	IpRanges                  SystemDhcpServerIpRangeArrayInput
	IpsecLeaseHold            pulumi.IntPtrInput
	LeaseTime                 pulumi.IntPtrInput
	MacAclDefaultAction       pulumi.StringPtrInput
	Netmask                   pulumi.StringInput
	NextServer                pulumi.StringPtrInput
	NtpServer1                pulumi.StringPtrInput
	NtpServer2                pulumi.StringPtrInput
	NtpServer3                pulumi.StringPtrInput
	NtpService                pulumi.StringPtrInput
	Options                   SystemDhcpServerOptionArrayInput
	RelayAgent                pulumi.StringPtrInput
	ReservedAddresses         SystemDhcpServerReservedAddressArrayInput
	ServerType                pulumi.StringPtrInput
	SharedSubnet              pulumi.StringPtrInput
	Status                    pulumi.StringPtrInput
	TftpServers               SystemDhcpServerTftpServerArrayInput
	Timezone                  pulumi.StringPtrInput
	TimezoneOption            pulumi.StringPtrInput
	VciMatch                  pulumi.StringPtrInput
	VciStrings                SystemDhcpServerVciStringArrayInput
	Vdomparam                 pulumi.StringPtrInput
	WifiAc1                   pulumi.StringPtrInput
	WifiAc2                   pulumi.StringPtrInput
	WifiAc3                   pulumi.StringPtrInput
	WifiAcService             pulumi.StringPtrInput
	WinsServer1               pulumi.StringPtrInput
	WinsServer2               pulumi.StringPtrInput
}

func (SystemDhcpServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDhcpServerArgs)(nil)).Elem()
}

type SystemDhcpServerInput interface {
	pulumi.Input

	ToSystemDhcpServerOutput() SystemDhcpServerOutput
	ToSystemDhcpServerOutputWithContext(ctx context.Context) SystemDhcpServerOutput
}

func (*SystemDhcpServer) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDhcpServer)(nil)).Elem()
}

func (i *SystemDhcpServer) ToSystemDhcpServerOutput() SystemDhcpServerOutput {
	return i.ToSystemDhcpServerOutputWithContext(context.Background())
}

func (i *SystemDhcpServer) ToSystemDhcpServerOutputWithContext(ctx context.Context) SystemDhcpServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDhcpServerOutput)
}

// SystemDhcpServerArrayInput is an input type that accepts SystemDhcpServerArray and SystemDhcpServerArrayOutput values.
// You can construct a concrete instance of `SystemDhcpServerArrayInput` via:
//
//	SystemDhcpServerArray{ SystemDhcpServerArgs{...} }
type SystemDhcpServerArrayInput interface {
	pulumi.Input

	ToSystemDhcpServerArrayOutput() SystemDhcpServerArrayOutput
	ToSystemDhcpServerArrayOutputWithContext(context.Context) SystemDhcpServerArrayOutput
}

type SystemDhcpServerArray []SystemDhcpServerInput

func (SystemDhcpServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDhcpServer)(nil)).Elem()
}

func (i SystemDhcpServerArray) ToSystemDhcpServerArrayOutput() SystemDhcpServerArrayOutput {
	return i.ToSystemDhcpServerArrayOutputWithContext(context.Background())
}

func (i SystemDhcpServerArray) ToSystemDhcpServerArrayOutputWithContext(ctx context.Context) SystemDhcpServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDhcpServerArrayOutput)
}

// SystemDhcpServerMapInput is an input type that accepts SystemDhcpServerMap and SystemDhcpServerMapOutput values.
// You can construct a concrete instance of `SystemDhcpServerMapInput` via:
//
//	SystemDhcpServerMap{ "key": SystemDhcpServerArgs{...} }
type SystemDhcpServerMapInput interface {
	pulumi.Input

	ToSystemDhcpServerMapOutput() SystemDhcpServerMapOutput
	ToSystemDhcpServerMapOutputWithContext(context.Context) SystemDhcpServerMapOutput
}

type SystemDhcpServerMap map[string]SystemDhcpServerInput

func (SystemDhcpServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDhcpServer)(nil)).Elem()
}

func (i SystemDhcpServerMap) ToSystemDhcpServerMapOutput() SystemDhcpServerMapOutput {
	return i.ToSystemDhcpServerMapOutputWithContext(context.Background())
}

func (i SystemDhcpServerMap) ToSystemDhcpServerMapOutputWithContext(ctx context.Context) SystemDhcpServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDhcpServerMapOutput)
}

type SystemDhcpServerOutput struct{ *pulumi.OutputState }

func (SystemDhcpServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDhcpServer)(nil)).Elem()
}

func (o SystemDhcpServerOutput) ToSystemDhcpServerOutput() SystemDhcpServerOutput {
	return o
}

func (o SystemDhcpServerOutput) ToSystemDhcpServerOutputWithContext(ctx context.Context) SystemDhcpServerOutput {
	return o
}

func (o SystemDhcpServerOutput) AutoConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.AutoConfiguration }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) AutoManagedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.AutoManagedStatus }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) ConflictedIpTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.IntOutput { return v.ConflictedIpTimeout }).(pulumi.IntOutput)
}

func (o SystemDhcpServerOutput) DdnsAuth() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DdnsAuth }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DdnsKey() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DdnsKey }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DdnsKeyname() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DdnsKeyname }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DdnsServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DdnsServerIp }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DdnsTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.IntOutput { return v.DdnsTtl }).(pulumi.IntOutput)
}

func (o SystemDhcpServerOutput) DdnsUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DdnsUpdate }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DdnsUpdateOverride() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DdnsUpdateOverride }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DdnsZone() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DdnsZone }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DefaultGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DefaultGateway }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DhcpSettingsFromFortiipam() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DhcpSettingsFromFortiipam }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DnsServer1() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DnsServer1 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DnsServer2() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DnsServer2 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DnsServer3() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DnsServer3 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DnsServer4() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DnsServer4 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DnsService() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.DnsService }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

func (o SystemDhcpServerOutput) ExcludeRanges() SystemDhcpServerExcludeRangeArrayOutput {
	return o.ApplyT(func(v *SystemDhcpServer) SystemDhcpServerExcludeRangeArrayOutput { return v.ExcludeRanges }).(SystemDhcpServerExcludeRangeArrayOutput)
}

func (o SystemDhcpServerOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.Filename }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) ForticlientOnNetStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.ForticlientOnNetStatus }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

func (o SystemDhcpServerOutput) GetAllTables() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringPtrOutput { return v.GetAllTables }).(pulumi.StringPtrOutput)
}

func (o SystemDhcpServerOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) IpMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.IpMode }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) IpRanges() SystemDhcpServerIpRangeArrayOutput {
	return o.ApplyT(func(v *SystemDhcpServer) SystemDhcpServerIpRangeArrayOutput { return v.IpRanges }).(SystemDhcpServerIpRangeArrayOutput)
}

func (o SystemDhcpServerOutput) IpsecLeaseHold() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.IntOutput { return v.IpsecLeaseHold }).(pulumi.IntOutput)
}

func (o SystemDhcpServerOutput) LeaseTime() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.IntOutput { return v.LeaseTime }).(pulumi.IntOutput)
}

func (o SystemDhcpServerOutput) MacAclDefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.MacAclDefaultAction }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.Netmask }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) NextServer() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.NextServer }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) NtpServer1() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.NtpServer1 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) NtpServer2() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.NtpServer2 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) NtpServer3() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.NtpServer3 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) NtpService() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.NtpService }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) Options() SystemDhcpServerOptionArrayOutput {
	return o.ApplyT(func(v *SystemDhcpServer) SystemDhcpServerOptionArrayOutput { return v.Options }).(SystemDhcpServerOptionArrayOutput)
}

func (o SystemDhcpServerOutput) RelayAgent() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.RelayAgent }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) ReservedAddresses() SystemDhcpServerReservedAddressArrayOutput {
	return o.ApplyT(func(v *SystemDhcpServer) SystemDhcpServerReservedAddressArrayOutput { return v.ReservedAddresses }).(SystemDhcpServerReservedAddressArrayOutput)
}

func (o SystemDhcpServerOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) SharedSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.SharedSubnet }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) TftpServers() SystemDhcpServerTftpServerArrayOutput {
	return o.ApplyT(func(v *SystemDhcpServer) SystemDhcpServerTftpServerArrayOutput { return v.TftpServers }).(SystemDhcpServerTftpServerArrayOutput)
}

func (o SystemDhcpServerOutput) Timezone() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.Timezone }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) TimezoneOption() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.TimezoneOption }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) VciMatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.VciMatch }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) VciStrings() SystemDhcpServerVciStringArrayOutput {
	return o.ApplyT(func(v *SystemDhcpServer) SystemDhcpServerVciStringArrayOutput { return v.VciStrings }).(SystemDhcpServerVciStringArrayOutput)
}

func (o SystemDhcpServerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o SystemDhcpServerOutput) WifiAc1() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.WifiAc1 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) WifiAc2() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.WifiAc2 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) WifiAc3() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.WifiAc3 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) WifiAcService() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.WifiAcService }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) WinsServer1() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.WinsServer1 }).(pulumi.StringOutput)
}

func (o SystemDhcpServerOutput) WinsServer2() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDhcpServer) pulumi.StringOutput { return v.WinsServer2 }).(pulumi.StringOutput)
}

type SystemDhcpServerArrayOutput struct{ *pulumi.OutputState }

func (SystemDhcpServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDhcpServer)(nil)).Elem()
}

func (o SystemDhcpServerArrayOutput) ToSystemDhcpServerArrayOutput() SystemDhcpServerArrayOutput {
	return o
}

func (o SystemDhcpServerArrayOutput) ToSystemDhcpServerArrayOutputWithContext(ctx context.Context) SystemDhcpServerArrayOutput {
	return o
}

func (o SystemDhcpServerArrayOutput) Index(i pulumi.IntInput) SystemDhcpServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemDhcpServer {
		return vs[0].([]*SystemDhcpServer)[vs[1].(int)]
	}).(SystemDhcpServerOutput)
}

type SystemDhcpServerMapOutput struct{ *pulumi.OutputState }

func (SystemDhcpServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDhcpServer)(nil)).Elem()
}

func (o SystemDhcpServerMapOutput) ToSystemDhcpServerMapOutput() SystemDhcpServerMapOutput {
	return o
}

func (o SystemDhcpServerMapOutput) ToSystemDhcpServerMapOutputWithContext(ctx context.Context) SystemDhcpServerMapOutput {
	return o
}

func (o SystemDhcpServerMapOutput) MapIndex(k pulumi.StringInput) SystemDhcpServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemDhcpServer {
		return vs[0].(map[string]*SystemDhcpServer)[vs[1].(string)]
	}).(SystemDhcpServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDhcpServerInput)(nil)).Elem(), &SystemDhcpServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDhcpServerArrayInput)(nil)).Elem(), SystemDhcpServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDhcpServerMapInput)(nil)).Elem(), SystemDhcpServerMap{})
	pulumi.RegisterOutputType(SystemDhcpServerOutput{})
	pulumi.RegisterOutputType(SystemDhcpServerArrayOutput{})
	pulumi.RegisterOutputType(SystemDhcpServerMapOutput{})
}
