// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure sniffer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/lubyou/pulumi-fortios/sdk/go/fortios"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := fortios.NewFirewallSniffer(ctx, "trname", &fortios.FirewallSnifferArgs{
// 			ApplicationListStatus:   pulumi.String("disable"),
// 			AvProfileStatus:         pulumi.String("disable"),
// 			DlpSensorStatus:         pulumi.String("disable"),
// 			Dsri:                    pulumi.String("disable"),
// 			Fosid:                   pulumi.Int(1),
// 			Interface:               pulumi.String("port4"),
// 			IpsDosStatus:            pulumi.String("disable"),
// 			IpsSensorStatus:         pulumi.String("disable"),
// 			Ipv6:                    pulumi.String("disable"),
// 			Logtraffic:              pulumi.String("utm"),
// 			MaxPacketCount:          pulumi.Int(4000),
// 			NonIp:                   pulumi.String("enable"),
// 			ScanBotnetConnections:   pulumi.String("disable"),
// 			SpamfilterProfileStatus: pulumi.String("disable"),
// 			Status:                  pulumi.String("enable"),
// 			WebfilterProfileStatus:  pulumi.String("disable"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Firewall Sniffer can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import fortios:index/firewallSniffer:FirewallSniffer labelname {{fosid}}
// ```
//
//  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//  $ pulumi import fortios:index/firewallSniffer:FirewallSniffer labelname {{fosid}}
// ```
//
//  $ unset "FORTIOS_IMPORT_TABLE"
type FirewallSniffer struct {
	pulumi.CustomResourceState

	// Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
	Anomalies FirewallSnifferAnomalyArrayOutput `pulumi:"anomalies"`
	// Name of an existing application list.
	ApplicationList pulumi.StringOutput `pulumi:"applicationList"`
	// Enable/disable application control profile. Valid values: `enable`, `disable`.
	ApplicationListStatus pulumi.StringOutput `pulumi:"applicationListStatus"`
	// Name of an existing antivirus profile.
	AvProfile pulumi.StringOutput `pulumi:"avProfile"`
	// Enable/disable antivirus profile. Valid values: `enable`, `disable`.
	AvProfileStatus pulumi.StringOutput `pulumi:"avProfileStatus"`
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringOutput `pulumi:"dlpSensor"`
	// Enable/disable DLP sensor. Valid values: `enable`, `disable`.
	DlpSensorStatus pulumi.StringOutput `pulumi:"dlpSensorStatus"`
	// Enable/disable DSRI. Valid values: `enable`, `disable`.
	Dsri pulumi.StringOutput `pulumi:"dsri"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringOutput `pulumi:"emailfilterProfile"`
	// Enable/disable emailfilter. Valid values: `enable`, `disable`.
	EmailfilterProfileStatus pulumi.StringOutput `pulumi:"emailfilterProfileStatus"`
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringOutput `pulumi:"fileFilterProfile"`
	// Enable/disable file filter. Valid values: `enable`, `disable`.
	FileFilterProfileStatus pulumi.StringOutput `pulumi:"fileFilterProfileStatus"`
	// Sniffer ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
	Host pulumi.StringOutput `pulumi:"host"`
	// Interface name that traffic sniffing will take place on.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Enable/disable IP threat feed. Valid values: `enable`, `disable`.
	IpThreatfeedStatus pulumi.StringOutput `pulumi:"ipThreatfeedStatus"`
	// Name of an existing IP threat feed. The structure of `ipThreatfeed` block is documented below.
	IpThreatfeeds FirewallSnifferIpThreatfeedArrayOutput `pulumi:"ipThreatfeeds"`
	// Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
	IpsDosStatus pulumi.StringOutput `pulumi:"ipsDosStatus"`
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringOutput `pulumi:"ipsSensor"`
	// Enable/disable IPS sensor. Valid values: `enable`, `disable`.
	IpsSensorStatus pulumi.StringOutput `pulumi:"ipsSensorStatus"`
	// Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
	Ipv6 pulumi.StringOutput `pulumi:"ipv6"`
	// Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
	Logtraffic pulumi.StringOutput `pulumi:"logtraffic"`
	// Maximum packet count (1 - 1000000, default = 10000).
	MaxPacketCount pulumi.IntOutput `pulumi:"maxPacketCount"`
	// Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
	NonIp pulumi.StringOutput `pulumi:"nonIp"`
	// Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
	Port pulumi.StringOutput `pulumi:"port"`
	// Integer value for the protocol type as defined by IANA (0 - 255).
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections pulumi.StringOutput `pulumi:"scanBotnetConnections"`
	// Name of an existing spam filter profile.
	SpamfilterProfile pulumi.StringOutput `pulumi:"spamfilterProfile"`
	// Enable/disable spam filter. Valid values: `enable`, `disable`.
	SpamfilterProfileStatus pulumi.StringOutput `pulumi:"spamfilterProfileStatus"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// List of VLANs to sniff.
	Vlan pulumi.StringOutput `pulumi:"vlan"`
	// Name of an existing web filter profile.
	WebfilterProfile pulumi.StringOutput `pulumi:"webfilterProfile"`
	// Enable/disable web filter profile. Valid values: `enable`, `disable`.
	WebfilterProfileStatus pulumi.StringOutput `pulumi:"webfilterProfileStatus"`
}

// NewFirewallSniffer registers a new resource with the given unique name, arguments, and options.
func NewFirewallSniffer(ctx *pulumi.Context,
	name string, args *FirewallSnifferArgs, opts ...pulumi.ResourceOption) (*FirewallSniffer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallSniffer
	err := ctx.RegisterResource("fortios:index/firewallSniffer:FirewallSniffer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSniffer gets an existing FirewallSniffer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSniffer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSnifferState, opts ...pulumi.ResourceOption) (*FirewallSniffer, error) {
	var resource FirewallSniffer
	err := ctx.ReadResource("fortios:index/firewallSniffer:FirewallSniffer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSniffer resources.
type firewallSnifferState struct {
	// Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
	Anomalies []FirewallSnifferAnomaly `pulumi:"anomalies"`
	// Name of an existing application list.
	ApplicationList *string `pulumi:"applicationList"`
	// Enable/disable application control profile. Valid values: `enable`, `disable`.
	ApplicationListStatus *string `pulumi:"applicationListStatus"`
	// Name of an existing antivirus profile.
	AvProfile *string `pulumi:"avProfile"`
	// Enable/disable antivirus profile. Valid values: `enable`, `disable`.
	AvProfileStatus *string `pulumi:"avProfileStatus"`
	// Name of an existing DLP sensor.
	DlpSensor *string `pulumi:"dlpSensor"`
	// Enable/disable DLP sensor. Valid values: `enable`, `disable`.
	DlpSensorStatus *string `pulumi:"dlpSensorStatus"`
	// Enable/disable DSRI. Valid values: `enable`, `disable`.
	Dsri *string `pulumi:"dsri"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name of an existing email filter profile.
	EmailfilterProfile *string `pulumi:"emailfilterProfile"`
	// Enable/disable emailfilter. Valid values: `enable`, `disable`.
	EmailfilterProfileStatus *string `pulumi:"emailfilterProfileStatus"`
	// Name of an existing file-filter profile.
	FileFilterProfile *string `pulumi:"fileFilterProfile"`
	// Enable/disable file filter. Valid values: `enable`, `disable`.
	FileFilterProfileStatus *string `pulumi:"fileFilterProfileStatus"`
	// Sniffer ID.
	Fosid *int `pulumi:"fosid"`
	// Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
	Host *string `pulumi:"host"`
	// Interface name that traffic sniffing will take place on.
	Interface *string `pulumi:"interface"`
	// Enable/disable IP threat feed. Valid values: `enable`, `disable`.
	IpThreatfeedStatus *string `pulumi:"ipThreatfeedStatus"`
	// Name of an existing IP threat feed. The structure of `ipThreatfeed` block is documented below.
	IpThreatfeeds []FirewallSnifferIpThreatfeed `pulumi:"ipThreatfeeds"`
	// Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
	IpsDosStatus *string `pulumi:"ipsDosStatus"`
	// Name of an existing IPS sensor.
	IpsSensor *string `pulumi:"ipsSensor"`
	// Enable/disable IPS sensor. Valid values: `enable`, `disable`.
	IpsSensorStatus *string `pulumi:"ipsSensorStatus"`
	// Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
	Ipv6 *string `pulumi:"ipv6"`
	// Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
	Logtraffic *string `pulumi:"logtraffic"`
	// Maximum packet count (1 - 1000000, default = 10000).
	MaxPacketCount *int `pulumi:"maxPacketCount"`
	// Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
	NonIp *string `pulumi:"nonIp"`
	// Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
	Port *string `pulumi:"port"`
	// Integer value for the protocol type as defined by IANA (0 - 255).
	Protocol *string `pulumi:"protocol"`
	// Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections *string `pulumi:"scanBotnetConnections"`
	// Name of an existing spam filter profile.
	SpamfilterProfile *string `pulumi:"spamfilterProfile"`
	// Enable/disable spam filter. Valid values: `enable`, `disable`.
	SpamfilterProfileStatus *string `pulumi:"spamfilterProfileStatus"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// List of VLANs to sniff.
	Vlan *string `pulumi:"vlan"`
	// Name of an existing web filter profile.
	WebfilterProfile *string `pulumi:"webfilterProfile"`
	// Enable/disable web filter profile. Valid values: `enable`, `disable`.
	WebfilterProfileStatus *string `pulumi:"webfilterProfileStatus"`
}

type FirewallSnifferState struct {
	// Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
	Anomalies FirewallSnifferAnomalyArrayInput
	// Name of an existing application list.
	ApplicationList pulumi.StringPtrInput
	// Enable/disable application control profile. Valid values: `enable`, `disable`.
	ApplicationListStatus pulumi.StringPtrInput
	// Name of an existing antivirus profile.
	AvProfile pulumi.StringPtrInput
	// Enable/disable antivirus profile. Valid values: `enable`, `disable`.
	AvProfileStatus pulumi.StringPtrInput
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringPtrInput
	// Enable/disable DLP sensor. Valid values: `enable`, `disable`.
	DlpSensorStatus pulumi.StringPtrInput
	// Enable/disable DSRI. Valid values: `enable`, `disable`.
	Dsri pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringPtrInput
	// Enable/disable emailfilter. Valid values: `enable`, `disable`.
	EmailfilterProfileStatus pulumi.StringPtrInput
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringPtrInput
	// Enable/disable file filter. Valid values: `enable`, `disable`.
	FileFilterProfileStatus pulumi.StringPtrInput
	// Sniffer ID.
	Fosid pulumi.IntPtrInput
	// Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
	Host pulumi.StringPtrInput
	// Interface name that traffic sniffing will take place on.
	Interface pulumi.StringPtrInput
	// Enable/disable IP threat feed. Valid values: `enable`, `disable`.
	IpThreatfeedStatus pulumi.StringPtrInput
	// Name of an existing IP threat feed. The structure of `ipThreatfeed` block is documented below.
	IpThreatfeeds FirewallSnifferIpThreatfeedArrayInput
	// Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
	IpsDosStatus pulumi.StringPtrInput
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringPtrInput
	// Enable/disable IPS sensor. Valid values: `enable`, `disable`.
	IpsSensorStatus pulumi.StringPtrInput
	// Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
	Ipv6 pulumi.StringPtrInput
	// Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
	Logtraffic pulumi.StringPtrInput
	// Maximum packet count (1 - 1000000, default = 10000).
	MaxPacketCount pulumi.IntPtrInput
	// Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
	NonIp pulumi.StringPtrInput
	// Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
	Port pulumi.StringPtrInput
	// Integer value for the protocol type as defined by IANA (0 - 255).
	Protocol pulumi.StringPtrInput
	// Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections pulumi.StringPtrInput
	// Name of an existing spam filter profile.
	SpamfilterProfile pulumi.StringPtrInput
	// Enable/disable spam filter. Valid values: `enable`, `disable`.
	SpamfilterProfileStatus pulumi.StringPtrInput
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// List of VLANs to sniff.
	Vlan pulumi.StringPtrInput
	// Name of an existing web filter profile.
	WebfilterProfile pulumi.StringPtrInput
	// Enable/disable web filter profile. Valid values: `enable`, `disable`.
	WebfilterProfileStatus pulumi.StringPtrInput
}

func (FirewallSnifferState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSnifferState)(nil)).Elem()
}

type firewallSnifferArgs struct {
	// Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
	Anomalies []FirewallSnifferAnomaly `pulumi:"anomalies"`
	// Name of an existing application list.
	ApplicationList *string `pulumi:"applicationList"`
	// Enable/disable application control profile. Valid values: `enable`, `disable`.
	ApplicationListStatus *string `pulumi:"applicationListStatus"`
	// Name of an existing antivirus profile.
	AvProfile *string `pulumi:"avProfile"`
	// Enable/disable antivirus profile. Valid values: `enable`, `disable`.
	AvProfileStatus *string `pulumi:"avProfileStatus"`
	// Name of an existing DLP sensor.
	DlpSensor *string `pulumi:"dlpSensor"`
	// Enable/disable DLP sensor. Valid values: `enable`, `disable`.
	DlpSensorStatus *string `pulumi:"dlpSensorStatus"`
	// Enable/disable DSRI. Valid values: `enable`, `disable`.
	Dsri *string `pulumi:"dsri"`
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Name of an existing email filter profile.
	EmailfilterProfile *string `pulumi:"emailfilterProfile"`
	// Enable/disable emailfilter. Valid values: `enable`, `disable`.
	EmailfilterProfileStatus *string `pulumi:"emailfilterProfileStatus"`
	// Name of an existing file-filter profile.
	FileFilterProfile *string `pulumi:"fileFilterProfile"`
	// Enable/disable file filter. Valid values: `enable`, `disable`.
	FileFilterProfileStatus *string `pulumi:"fileFilterProfileStatus"`
	// Sniffer ID.
	Fosid *int `pulumi:"fosid"`
	// Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
	Host *string `pulumi:"host"`
	// Interface name that traffic sniffing will take place on.
	Interface string `pulumi:"interface"`
	// Enable/disable IP threat feed. Valid values: `enable`, `disable`.
	IpThreatfeedStatus *string `pulumi:"ipThreatfeedStatus"`
	// Name of an existing IP threat feed. The structure of `ipThreatfeed` block is documented below.
	IpThreatfeeds []FirewallSnifferIpThreatfeed `pulumi:"ipThreatfeeds"`
	// Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
	IpsDosStatus *string `pulumi:"ipsDosStatus"`
	// Name of an existing IPS sensor.
	IpsSensor *string `pulumi:"ipsSensor"`
	// Enable/disable IPS sensor. Valid values: `enable`, `disable`.
	IpsSensorStatus *string `pulumi:"ipsSensorStatus"`
	// Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
	Ipv6 *string `pulumi:"ipv6"`
	// Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
	Logtraffic *string `pulumi:"logtraffic"`
	// Maximum packet count (1 - 1000000, default = 10000).
	MaxPacketCount *int `pulumi:"maxPacketCount"`
	// Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
	NonIp *string `pulumi:"nonIp"`
	// Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
	Port *string `pulumi:"port"`
	// Integer value for the protocol type as defined by IANA (0 - 255).
	Protocol *string `pulumi:"protocol"`
	// Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections *string `pulumi:"scanBotnetConnections"`
	// Name of an existing spam filter profile.
	SpamfilterProfile *string `pulumi:"spamfilterProfile"`
	// Enable/disable spam filter. Valid values: `enable`, `disable`.
	SpamfilterProfileStatus *string `pulumi:"spamfilterProfileStatus"`
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// List of VLANs to sniff.
	Vlan *string `pulumi:"vlan"`
	// Name of an existing web filter profile.
	WebfilterProfile *string `pulumi:"webfilterProfile"`
	// Enable/disable web filter profile. Valid values: `enable`, `disable`.
	WebfilterProfileStatus *string `pulumi:"webfilterProfileStatus"`
}

// The set of arguments for constructing a FirewallSniffer resource.
type FirewallSnifferArgs struct {
	// Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
	Anomalies FirewallSnifferAnomalyArrayInput
	// Name of an existing application list.
	ApplicationList pulumi.StringPtrInput
	// Enable/disable application control profile. Valid values: `enable`, `disable`.
	ApplicationListStatus pulumi.StringPtrInput
	// Name of an existing antivirus profile.
	AvProfile pulumi.StringPtrInput
	// Enable/disable antivirus profile. Valid values: `enable`, `disable`.
	AvProfileStatus pulumi.StringPtrInput
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringPtrInput
	// Enable/disable DLP sensor. Valid values: `enable`, `disable`.
	DlpSensorStatus pulumi.StringPtrInput
	// Enable/disable DSRI. Valid values: `enable`, `disable`.
	Dsri pulumi.StringPtrInput
	// true or false, set this parameter to true when using dynamic forEach + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
	DynamicSortSubtable pulumi.StringPtrInput
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringPtrInput
	// Enable/disable emailfilter. Valid values: `enable`, `disable`.
	EmailfilterProfileStatus pulumi.StringPtrInput
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringPtrInput
	// Enable/disable file filter. Valid values: `enable`, `disable`.
	FileFilterProfileStatus pulumi.StringPtrInput
	// Sniffer ID.
	Fosid pulumi.IntPtrInput
	// Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
	Host pulumi.StringPtrInput
	// Interface name that traffic sniffing will take place on.
	Interface pulumi.StringInput
	// Enable/disable IP threat feed. Valid values: `enable`, `disable`.
	IpThreatfeedStatus pulumi.StringPtrInput
	// Name of an existing IP threat feed. The structure of `ipThreatfeed` block is documented below.
	IpThreatfeeds FirewallSnifferIpThreatfeedArrayInput
	// Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
	IpsDosStatus pulumi.StringPtrInput
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringPtrInput
	// Enable/disable IPS sensor. Valid values: `enable`, `disable`.
	IpsSensorStatus pulumi.StringPtrInput
	// Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
	Ipv6 pulumi.StringPtrInput
	// Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
	Logtraffic pulumi.StringPtrInput
	// Maximum packet count (1 - 1000000, default = 10000).
	MaxPacketCount pulumi.IntPtrInput
	// Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
	NonIp pulumi.StringPtrInput
	// Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
	Port pulumi.StringPtrInput
	// Integer value for the protocol type as defined by IANA (0 - 255).
	Protocol pulumi.StringPtrInput
	// Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
	ScanBotnetConnections pulumi.StringPtrInput
	// Name of an existing spam filter profile.
	SpamfilterProfile pulumi.StringPtrInput
	// Enable/disable spam filter. Valid values: `enable`, `disable`.
	SpamfilterProfileStatus pulumi.StringPtrInput
	// Enable/disable this anomaly. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// List of VLANs to sniff.
	Vlan pulumi.StringPtrInput
	// Name of an existing web filter profile.
	WebfilterProfile pulumi.StringPtrInput
	// Enable/disable web filter profile. Valid values: `enable`, `disable`.
	WebfilterProfileStatus pulumi.StringPtrInput
}

func (FirewallSnifferArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSnifferArgs)(nil)).Elem()
}

type FirewallSnifferInput interface {
	pulumi.Input

	ToFirewallSnifferOutput() FirewallSnifferOutput
	ToFirewallSnifferOutputWithContext(ctx context.Context) FirewallSnifferOutput
}

func (*FirewallSniffer) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSniffer)(nil)).Elem()
}

func (i *FirewallSniffer) ToFirewallSnifferOutput() FirewallSnifferOutput {
	return i.ToFirewallSnifferOutputWithContext(context.Background())
}

func (i *FirewallSniffer) ToFirewallSnifferOutputWithContext(ctx context.Context) FirewallSnifferOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSnifferOutput)
}

// FirewallSnifferArrayInput is an input type that accepts FirewallSnifferArray and FirewallSnifferArrayOutput values.
// You can construct a concrete instance of `FirewallSnifferArrayInput` via:
//
//          FirewallSnifferArray{ FirewallSnifferArgs{...} }
type FirewallSnifferArrayInput interface {
	pulumi.Input

	ToFirewallSnifferArrayOutput() FirewallSnifferArrayOutput
	ToFirewallSnifferArrayOutputWithContext(context.Context) FirewallSnifferArrayOutput
}

type FirewallSnifferArray []FirewallSnifferInput

func (FirewallSnifferArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSniffer)(nil)).Elem()
}

func (i FirewallSnifferArray) ToFirewallSnifferArrayOutput() FirewallSnifferArrayOutput {
	return i.ToFirewallSnifferArrayOutputWithContext(context.Background())
}

func (i FirewallSnifferArray) ToFirewallSnifferArrayOutputWithContext(ctx context.Context) FirewallSnifferArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSnifferArrayOutput)
}

// FirewallSnifferMapInput is an input type that accepts FirewallSnifferMap and FirewallSnifferMapOutput values.
// You can construct a concrete instance of `FirewallSnifferMapInput` via:
//
//          FirewallSnifferMap{ "key": FirewallSnifferArgs{...} }
type FirewallSnifferMapInput interface {
	pulumi.Input

	ToFirewallSnifferMapOutput() FirewallSnifferMapOutput
	ToFirewallSnifferMapOutputWithContext(context.Context) FirewallSnifferMapOutput
}

type FirewallSnifferMap map[string]FirewallSnifferInput

func (FirewallSnifferMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSniffer)(nil)).Elem()
}

func (i FirewallSnifferMap) ToFirewallSnifferMapOutput() FirewallSnifferMapOutput {
	return i.ToFirewallSnifferMapOutputWithContext(context.Background())
}

func (i FirewallSnifferMap) ToFirewallSnifferMapOutputWithContext(ctx context.Context) FirewallSnifferMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSnifferMapOutput)
}

type FirewallSnifferOutput struct{ *pulumi.OutputState }

func (FirewallSnifferOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSniffer)(nil)).Elem()
}

func (o FirewallSnifferOutput) ToFirewallSnifferOutput() FirewallSnifferOutput {
	return o
}

func (o FirewallSnifferOutput) ToFirewallSnifferOutputWithContext(ctx context.Context) FirewallSnifferOutput {
	return o
}

type FirewallSnifferArrayOutput struct{ *pulumi.OutputState }

func (FirewallSnifferArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSniffer)(nil)).Elem()
}

func (o FirewallSnifferArrayOutput) ToFirewallSnifferArrayOutput() FirewallSnifferArrayOutput {
	return o
}

func (o FirewallSnifferArrayOutput) ToFirewallSnifferArrayOutputWithContext(ctx context.Context) FirewallSnifferArrayOutput {
	return o
}

func (o FirewallSnifferArrayOutput) Index(i pulumi.IntInput) FirewallSnifferOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSniffer {
		return vs[0].([]*FirewallSniffer)[vs[1].(int)]
	}).(FirewallSnifferOutput)
}

type FirewallSnifferMapOutput struct{ *pulumi.OutputState }

func (FirewallSnifferMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSniffer)(nil)).Elem()
}

func (o FirewallSnifferMapOutput) ToFirewallSnifferMapOutput() FirewallSnifferMapOutput {
	return o
}

func (o FirewallSnifferMapOutput) ToFirewallSnifferMapOutputWithContext(ctx context.Context) FirewallSnifferMapOutput {
	return o
}

func (o FirewallSnifferMapOutput) MapIndex(k pulumi.StringInput) FirewallSnifferOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSniffer {
		return vs[0].(map[string]*FirewallSniffer)[vs[1].(string)]
	}).(FirewallSnifferOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSnifferInput)(nil)).Elem(), &FirewallSniffer{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSnifferArrayInput)(nil)).Elem(), FirewallSnifferArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSnifferMapInput)(nil)).Elem(), FirewallSnifferMap{})
	pulumi.RegisterOutputType(FirewallSnifferOutput{})
	pulumi.RegisterOutputType(FirewallSnifferArrayOutput{})
	pulumi.RegisterOutputType(FirewallSnifferMapOutput{})
}
