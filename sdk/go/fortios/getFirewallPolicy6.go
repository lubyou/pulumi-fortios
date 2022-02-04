// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall policy6
func LookupFirewallPolicy6(ctx *pulumi.Context, args *LookupFirewallPolicy6Args, opts ...pulumi.InvokeOption) (*LookupFirewallPolicy6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallPolicy6Result
	err := ctx.Invoke("fortios:index/getFirewallPolicy6:GetFirewallPolicy6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallPolicy6.
type LookupFirewallPolicy6Args struct {
	// Specify the policyid of the desired firewall policy6.
	Policyid int `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallPolicy6.
type LookupFirewallPolicy6Result struct {
	// Policy action (allow/deny/ipsec).
	Action string `pulumi:"action"`
	// Enable/disable anti-replay check.
	AntiReplay string `pulumi:"antiReplay"`
	// Application category ID list. The structure of `appCategory` block is documented below.
	AppCategories []GetFirewallPolicy6AppCategory `pulumi:"appCategories"`
	// Application group names. The structure of `appGroup` block is documented below.
	AppGroups []GetFirewallPolicy6AppGroup `pulumi:"appGroups"`
	// Name of an existing Application list.
	ApplicationList string `pulumi:"applicationList"`
	// Application ID list. The structure of `application` block is documented below.
	Applications []GetFirewallPolicy6Application `pulumi:"applications"`
	// Enable/disable policy traffic ASIC offloading.
	AutoAsicOffload string `pulumi:"autoAsicOffload"`
	// Name of an existing Antivirus profile.
	AvProfile string `pulumi:"avProfile"`
	// Name of an existing CIFS profile.
	CifsProfile string `pulumi:"cifsProfile"`
	// Comment.
	Comments string `pulumi:"comments"`
	// Log field index numbers to append custom log fields to log messages for this policy. The structure of `customLogFields` block is documented below.
	CustomLogFields []GetFirewallPolicy6CustomLogField `pulumi:"customLogFields"`
	// Decrypted traffic mirror.
	DecryptedTrafficMirror string `pulumi:"decryptedTrafficMirror"`
	// Names of devices or device groups that can be matched by the policy. The structure of `devices` block is documented below.
	Devices []GetFirewallPolicy6Device `pulumi:"devices"`
	// Enable to change packet's DiffServ values to the specified diffservcode-forward value.
	DiffservForward string `pulumi:"diffservForward"`
	// Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
	DiffservReverse string `pulumi:"diffservReverse"`
	// Change packet's DiffServ to this value.
	DiffservcodeForward string `pulumi:"diffservcodeForward"`
	// Change packet's reverse (reply) DiffServ to this value.
	DiffservcodeRev string `pulumi:"diffservcodeRev"`
	// Name of an existing DLP sensor.
	DlpSensor string `pulumi:"dlpSensor"`
	// Name of an existing DNS filter profile.
	DnsfilterProfile string `pulumi:"dnsfilterProfile"`
	// Enable DSRI to ignore HTTP server responses.
	Dsri string `pulumi:"dsri"`
	// When enabled dstaddr specifies what the destination address must NOT be.
	DstaddrNegate string `pulumi:"dstaddrNegate"`
	// Destination address and address group names. The structure of `dstaddr` block is documented below.
	Dstaddrs []GetFirewallPolicy6Dstaddr `pulumi:"dstaddrs"`
	// Outgoing (egress) interface. The structure of `dstintf` block is documented below.
	Dstintfs []GetFirewallPolicy6Dstintf `pulumi:"dstintfs"`
	// Name of an existing email filter profile.
	EmailfilterProfile string `pulumi:"emailfilterProfile"`
	// How to handle sessions if the configuration of this firewall policy changes.
	FirewallSessionDirty string `pulumi:"firewallSessionDirty"`
	// Enable to prevent source NAT from changing a session's source port.
	Fixedport string `pulumi:"fixedport"`
	// Names of FSSO groups. The structure of `fssoGroups` block is documented below.
	FssoGroups []GetFirewallPolicy6FssoGroup `pulumi:"fssoGroups"`
	// Label for the policy that appears when the GUI is in Global View mode.
	GlobalLabel string `pulumi:"globalLabel"`
	// Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
	Groups []GetFirewallPolicy6Group `pulumi:"groups"`
	// Redirect HTTP(S) traffic to matching transparent web proxy policy.
	HttpPolicyRedirect string `pulumi:"httpPolicyRedirect"`
	// Name of an existing ICAP profile.
	IcapProfile string `pulumi:"icapProfile"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
	Inbound string `pulumi:"inbound"`
	// Policy inspection mode (Flow/proxy). Default is Flow mode.
	InspectionMode string `pulumi:"inspectionMode"`
	// Enable to use IP Pools for source NAT.
	Ippool string `pulumi:"ippool"`
	// Name of an existing IPS sensor.
	IpsSensor string `pulumi:"ipsSensor"`
	// Label for the policy that appears when the GUI is in Section View mode.
	Label string `pulumi:"label"`
	// Enable or disable logging. Log all sessions or security profile sessions.
	Logtraffic string `pulumi:"logtraffic"`
	// Record logs when a session starts.
	LogtrafficStart string `pulumi:"logtrafficStart"`
	// Names of FSSO groups.
	Name string `pulumi:"name"`
	// Enable/disable source NAT.
	Nat string `pulumi:"nat"`
	// Policy-based IPsec VPN: apply destination NAT to inbound traffic.
	Natinbound string `pulumi:"natinbound"`
	// Policy-based IPsec VPN: apply source NAT to outbound traffic.
	Natoutbound string `pulumi:"natoutbound"`
	// Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
	Outbound string `pulumi:"outbound"`
	// Per-IP traffic shaper.
	PerIpShaper string `pulumi:"perIpShaper"`
	// Policy ID.
	Policyid int `pulumi:"policyid"`
	// IP Pool names. The structure of `poolname` block is documented below.
	Poolnames []GetFirewallPolicy6Poolname `pulumi:"poolnames"`
	// Name of profile group.
	ProfileGroup string `pulumi:"profileGroup"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions string `pulumi:"profileProtocolOptions"`
	// Determine whether the firewall policy allows security profile groups or single profiles only.
	ProfileType string `pulumi:"profileType"`
	// Override the default replacement message group for this policy.
	ReplacemsgOverrideGroup string `pulumi:"replacemsgOverrideGroup"`
	// Enable/disable RADIUS single sign-on (RSSO).
	Rsso string `pulumi:"rsso"`
	// Schedule name.
	Schedule string `pulumi:"schedule"`
	// Enable/disable return of deny-packet.
	SendDenyPacket string `pulumi:"sendDenyPacket"`
	// When enabled service specifies what the service must NOT be.
	ServiceNegate string `pulumi:"serviceNegate"`
	// Service and service group names. The structure of `service` block is documented below.
	Services []GetFirewallPolicy6Service `pulumi:"services"`
	// Session TTL in seconds for sessions accepted by this policy. 0 means use the system default session TTL.
	SessionTtl int `pulumi:"sessionTtl"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile string `pulumi:"spamfilterProfile"`
	// When enabled srcaddr specifies what the source address must NOT be.
	SrcaddrNegate string `pulumi:"srcaddrNegate"`
	// Source address and address group names. The structure of `srcaddr` block is documented below.
	Srcaddrs []GetFirewallPolicy6Srcaddr `pulumi:"srcaddrs"`
	// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
	Srcintfs []GetFirewallPolicy6Srcintf `pulumi:"srcintfs"`
	// Name of an existing SSH filter profile.
	SshFilterProfile string `pulumi:"sshFilterProfile"`
	// Redirect SSH traffic to matching transparent proxy policy.
	SshPolicyRedirect string `pulumi:"sshPolicyRedirect"`
	// Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring).
	SslMirror string `pulumi:"sslMirror"`
	// SSL mirror interface name. The structure of `sslMirrorIntf` block is documented below.
	SslMirrorIntfs []GetFirewallPolicy6SslMirrorIntf `pulumi:"sslMirrorIntfs"`
	// Name of an existing SSL SSH profile.
	SslSshProfile string `pulumi:"sslSshProfile"`
	// Enable or disable this policy.
	Status string `pulumi:"status"`
	// Receiver TCP maximum segment size (MSS).
	TcpMssReceiver int `pulumi:"tcpMssReceiver"`
	// Sender TCP maximum segment size (MSS).
	TcpMssSender int `pulumi:"tcpMssSender"`
	// Enable/disable creation of TCP session without SYN flag.
	TcpSessionWithoutSyn string `pulumi:"tcpSessionWithoutSyn"`
	// Enable/disable sending RST packets when TCP sessions expire.
	TimeoutSendRst string `pulumi:"timeoutSendRst"`
	// ToS (Type of Service) value used for comparison.
	Tos string `pulumi:"tos"`
	// Non-zero bit positions are used for comparison while zero bit positions are ignored.
	TosMask string `pulumi:"tosMask"`
	// Enable negated TOS match.
	TosNegate string `pulumi:"tosNegate"`
	// Reverse traffic shaper.
	TrafficShaper string `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse string `pulumi:"trafficShaperReverse"`
	// URL category ID list. The structure of `urlCategory` block is documented below.
	UrlCategories []GetFirewallPolicy6UrlCategory `pulumi:"urlCategories"`
	// Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
	Users []GetFirewallPolicy6User `pulumi:"users"`
	// Enable AV/web/ips protection profile.
	UtmStatus string `pulumi:"utmStatus"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest
	VlanCosFwd int `pulumi:"vlanCosFwd"`
	// VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest
	VlanCosRev int `pulumi:"vlanCosRev"`
	// Set VLAN filters.
	VlanFilter string `pulumi:"vlanFilter"`
	// Name of an existing VoIP profile.
	VoipProfile string `pulumi:"voipProfile"`
	// Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
	Vpntunnel string `pulumi:"vpntunnel"`
	// Name of an existing Web application firewall profile.
	WafProfile string `pulumi:"wafProfile"`
	// Enable/disable web cache.
	Webcache string `pulumi:"webcache"`
	// Enable/disable web cache for HTTPS.
	WebcacheHttps string `pulumi:"webcacheHttps"`
	// Name of an existing Web filter profile.
	WebfilterProfile string `pulumi:"webfilterProfile"`
	// Web proxy forward server name.
	WebproxyForwardServer string `pulumi:"webproxyForwardServer"`
	// Webproxy profile name.
	WebproxyProfile string `pulumi:"webproxyProfile"`
}

func LookupFirewallPolicy6Output(ctx *pulumi.Context, args LookupFirewallPolicy6OutputArgs, opts ...pulumi.InvokeOption) LookupFirewallPolicy6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallPolicy6Result, error) {
			args := v.(LookupFirewallPolicy6Args)
			r, err := LookupFirewallPolicy6(ctx, &args, opts...)
			return *r, err
		}).(LookupFirewallPolicy6ResultOutput)
}

// A collection of arguments for invoking GetFirewallPolicy6.
type LookupFirewallPolicy6OutputArgs struct {
	// Specify the policyid of the desired firewall policy6.
	Policyid pulumi.IntInput `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallPolicy6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicy6Args)(nil)).Elem()
}

// A collection of values returned by GetFirewallPolicy6.
type LookupFirewallPolicy6ResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallPolicy6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicy6Result)(nil)).Elem()
}

func (o LookupFirewallPolicy6ResultOutput) ToLookupFirewallPolicy6ResultOutput() LookupFirewallPolicy6ResultOutput {
	return o
}

func (o LookupFirewallPolicy6ResultOutput) ToLookupFirewallPolicy6ResultOutputWithContext(ctx context.Context) LookupFirewallPolicy6ResultOutput {
	return o
}

// Policy action (allow/deny/ipsec).
func (o LookupFirewallPolicy6ResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Action }).(pulumi.StringOutput)
}

// Enable/disable anti-replay check.
func (o LookupFirewallPolicy6ResultOutput) AntiReplay() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.AntiReplay }).(pulumi.StringOutput)
}

// Application category ID list. The structure of `appCategory` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) AppCategories() GetFirewallPolicy6AppCategoryArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6AppCategory { return v.AppCategories }).(GetFirewallPolicy6AppCategoryArrayOutput)
}

// Application group names. The structure of `appGroup` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) AppGroups() GetFirewallPolicy6AppGroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6AppGroup { return v.AppGroups }).(GetFirewallPolicy6AppGroupArrayOutput)
}

// Name of an existing Application list.
func (o LookupFirewallPolicy6ResultOutput) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ApplicationList }).(pulumi.StringOutput)
}

// Application ID list. The structure of `application` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Applications() GetFirewallPolicy6ApplicationArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Application { return v.Applications }).(GetFirewallPolicy6ApplicationArrayOutput)
}

// Enable/disable policy traffic ASIC offloading.
func (o LookupFirewallPolicy6ResultOutput) AutoAsicOffload() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.AutoAsicOffload }).(pulumi.StringOutput)
}

// Name of an existing Antivirus profile.
func (o LookupFirewallPolicy6ResultOutput) AvProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.AvProfile }).(pulumi.StringOutput)
}

// Name of an existing CIFS profile.
func (o LookupFirewallPolicy6ResultOutput) CifsProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.CifsProfile }).(pulumi.StringOutput)
}

// Comment.
func (o LookupFirewallPolicy6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Comments }).(pulumi.StringOutput)
}

// Log field index numbers to append custom log fields to log messages for this policy. The structure of `customLogFields` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) CustomLogFields() GetFirewallPolicy6CustomLogFieldArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6CustomLogField { return v.CustomLogFields }).(GetFirewallPolicy6CustomLogFieldArrayOutput)
}

// Decrypted traffic mirror.
func (o LookupFirewallPolicy6ResultOutput) DecryptedTrafficMirror() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DecryptedTrafficMirror }).(pulumi.StringOutput)
}

// Names of devices or device groups that can be matched by the policy. The structure of `devices` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Devices() GetFirewallPolicy6DeviceArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Device { return v.Devices }).(GetFirewallPolicy6DeviceArrayOutput)
}

// Enable to change packet's DiffServ values to the specified diffservcode-forward value.
func (o LookupFirewallPolicy6ResultOutput) DiffservForward() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DiffservForward }).(pulumi.StringOutput)
}

// Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
func (o LookupFirewallPolicy6ResultOutput) DiffservReverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DiffservReverse }).(pulumi.StringOutput)
}

// Change packet's DiffServ to this value.
func (o LookupFirewallPolicy6ResultOutput) DiffservcodeForward() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DiffservcodeForward }).(pulumi.StringOutput)
}

// Change packet's reverse (reply) DiffServ to this value.
func (o LookupFirewallPolicy6ResultOutput) DiffservcodeRev() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DiffservcodeRev }).(pulumi.StringOutput)
}

// Name of an existing DLP sensor.
func (o LookupFirewallPolicy6ResultOutput) DlpSensor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DlpSensor }).(pulumi.StringOutput)
}

// Name of an existing DNS filter profile.
func (o LookupFirewallPolicy6ResultOutput) DnsfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DnsfilterProfile }).(pulumi.StringOutput)
}

// Enable DSRI to ignore HTTP server responses.
func (o LookupFirewallPolicy6ResultOutput) Dsri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Dsri }).(pulumi.StringOutput)
}

// When enabled dstaddr specifies what the destination address must NOT be.
func (o LookupFirewallPolicy6ResultOutput) DstaddrNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DstaddrNegate }).(pulumi.StringOutput)
}

// Destination address and address group names. The structure of `dstaddr` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Dstaddrs() GetFirewallPolicy6DstaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Dstaddr { return v.Dstaddrs }).(GetFirewallPolicy6DstaddrArrayOutput)
}

// Outgoing (egress) interface. The structure of `dstintf` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Dstintfs() GetFirewallPolicy6DstintfArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Dstintf { return v.Dstintfs }).(GetFirewallPolicy6DstintfArrayOutput)
}

// Name of an existing email filter profile.
func (o LookupFirewallPolicy6ResultOutput) EmailfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.EmailfilterProfile }).(pulumi.StringOutput)
}

// How to handle sessions if the configuration of this firewall policy changes.
func (o LookupFirewallPolicy6ResultOutput) FirewallSessionDirty() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.FirewallSessionDirty }).(pulumi.StringOutput)
}

// Enable to prevent source NAT from changing a session's source port.
func (o LookupFirewallPolicy6ResultOutput) Fixedport() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Fixedport }).(pulumi.StringOutput)
}

// Names of FSSO groups. The structure of `fssoGroups` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) FssoGroups() GetFirewallPolicy6FssoGroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6FssoGroup { return v.FssoGroups }).(GetFirewallPolicy6FssoGroupArrayOutput)
}

// Label for the policy that appears when the GUI is in Global View mode.
func (o LookupFirewallPolicy6ResultOutput) GlobalLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.GlobalLabel }).(pulumi.StringOutput)
}

// Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Groups() GetFirewallPolicy6GroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Group { return v.Groups }).(GetFirewallPolicy6GroupArrayOutput)
}

// Redirect HTTP(S) traffic to matching transparent web proxy policy.
func (o LookupFirewallPolicy6ResultOutput) HttpPolicyRedirect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.HttpPolicyRedirect }).(pulumi.StringOutput)
}

// Name of an existing ICAP profile.
func (o LookupFirewallPolicy6ResultOutput) IcapProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.IcapProfile }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallPolicy6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
func (o LookupFirewallPolicy6ResultOutput) Inbound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Inbound }).(pulumi.StringOutput)
}

// Policy inspection mode (Flow/proxy). Default is Flow mode.
func (o LookupFirewallPolicy6ResultOutput) InspectionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.InspectionMode }).(pulumi.StringOutput)
}

// Enable to use IP Pools for source NAT.
func (o LookupFirewallPolicy6ResultOutput) Ippool() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Ippool }).(pulumi.StringOutput)
}

// Name of an existing IPS sensor.
func (o LookupFirewallPolicy6ResultOutput) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.IpsSensor }).(pulumi.StringOutput)
}

// Label for the policy that appears when the GUI is in Section View mode.
func (o LookupFirewallPolicy6ResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Label }).(pulumi.StringOutput)
}

// Enable or disable logging. Log all sessions or security profile sessions.
func (o LookupFirewallPolicy6ResultOutput) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Logtraffic }).(pulumi.StringOutput)
}

// Record logs when a session starts.
func (o LookupFirewallPolicy6ResultOutput) LogtrafficStart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.LogtrafficStart }).(pulumi.StringOutput)
}

// Names of FSSO groups.
func (o LookupFirewallPolicy6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable source NAT.
func (o LookupFirewallPolicy6ResultOutput) Nat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Nat }).(pulumi.StringOutput)
}

// Policy-based IPsec VPN: apply destination NAT to inbound traffic.
func (o LookupFirewallPolicy6ResultOutput) Natinbound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Natinbound }).(pulumi.StringOutput)
}

// Policy-based IPsec VPN: apply source NAT to outbound traffic.
func (o LookupFirewallPolicy6ResultOutput) Natoutbound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Natoutbound }).(pulumi.StringOutput)
}

// Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
func (o LookupFirewallPolicy6ResultOutput) Outbound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Outbound }).(pulumi.StringOutput)
}

// Per-IP traffic shaper.
func (o LookupFirewallPolicy6ResultOutput) PerIpShaper() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.PerIpShaper }).(pulumi.StringOutput)
}

// Policy ID.
func (o LookupFirewallPolicy6ResultOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.Policyid }).(pulumi.IntOutput)
}

// IP Pool names. The structure of `poolname` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Poolnames() GetFirewallPolicy6PoolnameArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Poolname { return v.Poolnames }).(GetFirewallPolicy6PoolnameArrayOutput)
}

// Name of profile group.
func (o LookupFirewallPolicy6ResultOutput) ProfileGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ProfileGroup }).(pulumi.StringOutput)
}

// Name of an existing Protocol options profile.
func (o LookupFirewallPolicy6ResultOutput) ProfileProtocolOptions() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ProfileProtocolOptions }).(pulumi.StringOutput)
}

// Determine whether the firewall policy allows security profile groups or single profiles only.
func (o LookupFirewallPolicy6ResultOutput) ProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ProfileType }).(pulumi.StringOutput)
}

// Override the default replacement message group for this policy.
func (o LookupFirewallPolicy6ResultOutput) ReplacemsgOverrideGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ReplacemsgOverrideGroup }).(pulumi.StringOutput)
}

// Enable/disable RADIUS single sign-on (RSSO).
func (o LookupFirewallPolicy6ResultOutput) Rsso() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Rsso }).(pulumi.StringOutput)
}

// Schedule name.
func (o LookupFirewallPolicy6ResultOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Schedule }).(pulumi.StringOutput)
}

// Enable/disable return of deny-packet.
func (o LookupFirewallPolicy6ResultOutput) SendDenyPacket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SendDenyPacket }).(pulumi.StringOutput)
}

// When enabled service specifies what the service must NOT be.
func (o LookupFirewallPolicy6ResultOutput) ServiceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ServiceNegate }).(pulumi.StringOutput)
}

// Service and service group names. The structure of `service` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Services() GetFirewallPolicy6ServiceArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Service { return v.Services }).(GetFirewallPolicy6ServiceArrayOutput)
}

// Session TTL in seconds for sessions accepted by this policy. 0 means use the system default session TTL.
func (o LookupFirewallPolicy6ResultOutput) SessionTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.SessionTtl }).(pulumi.IntOutput)
}

// Name of an existing Spam filter profile.
func (o LookupFirewallPolicy6ResultOutput) SpamfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SpamfilterProfile }).(pulumi.StringOutput)
}

// When enabled srcaddr specifies what the source address must NOT be.
func (o LookupFirewallPolicy6ResultOutput) SrcaddrNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SrcaddrNegate }).(pulumi.StringOutput)
}

// Source address and address group names. The structure of `srcaddr` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Srcaddrs() GetFirewallPolicy6SrcaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Srcaddr { return v.Srcaddrs }).(GetFirewallPolicy6SrcaddrArrayOutput)
}

// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Srcintfs() GetFirewallPolicy6SrcintfArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Srcintf { return v.Srcintfs }).(GetFirewallPolicy6SrcintfArrayOutput)
}

// Name of an existing SSH filter profile.
func (o LookupFirewallPolicy6ResultOutput) SshFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SshFilterProfile }).(pulumi.StringOutput)
}

// Redirect SSH traffic to matching transparent proxy policy.
func (o LookupFirewallPolicy6ResultOutput) SshPolicyRedirect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SshPolicyRedirect }).(pulumi.StringOutput)
}

// Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring).
func (o LookupFirewallPolicy6ResultOutput) SslMirror() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SslMirror }).(pulumi.StringOutput)
}

// SSL mirror interface name. The structure of `sslMirrorIntf` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) SslMirrorIntfs() GetFirewallPolicy6SslMirrorIntfArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6SslMirrorIntf { return v.SslMirrorIntfs }).(GetFirewallPolicy6SslMirrorIntfArrayOutput)
}

// Name of an existing SSL SSH profile.
func (o LookupFirewallPolicy6ResultOutput) SslSshProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SslSshProfile }).(pulumi.StringOutput)
}

// Enable or disable this policy.
func (o LookupFirewallPolicy6ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Status }).(pulumi.StringOutput)
}

// Receiver TCP maximum segment size (MSS).
func (o LookupFirewallPolicy6ResultOutput) TcpMssReceiver() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.TcpMssReceiver }).(pulumi.IntOutput)
}

// Sender TCP maximum segment size (MSS).
func (o LookupFirewallPolicy6ResultOutput) TcpMssSender() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.TcpMssSender }).(pulumi.IntOutput)
}

// Enable/disable creation of TCP session without SYN flag.
func (o LookupFirewallPolicy6ResultOutput) TcpSessionWithoutSyn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TcpSessionWithoutSyn }).(pulumi.StringOutput)
}

// Enable/disable sending RST packets when TCP sessions expire.
func (o LookupFirewallPolicy6ResultOutput) TimeoutSendRst() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TimeoutSendRst }).(pulumi.StringOutput)
}

// ToS (Type of Service) value used for comparison.
func (o LookupFirewallPolicy6ResultOutput) Tos() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Tos }).(pulumi.StringOutput)
}

// Non-zero bit positions are used for comparison while zero bit positions are ignored.
func (o LookupFirewallPolicy6ResultOutput) TosMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TosMask }).(pulumi.StringOutput)
}

// Enable negated TOS match.
func (o LookupFirewallPolicy6ResultOutput) TosNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TosNegate }).(pulumi.StringOutput)
}

// Reverse traffic shaper.
func (o LookupFirewallPolicy6ResultOutput) TrafficShaper() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TrafficShaper }).(pulumi.StringOutput)
}

// Reverse traffic shaper.
func (o LookupFirewallPolicy6ResultOutput) TrafficShaperReverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TrafficShaperReverse }).(pulumi.StringOutput)
}

// URL category ID list. The structure of `urlCategory` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) UrlCategories() GetFirewallPolicy6UrlCategoryArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6UrlCategory { return v.UrlCategories }).(GetFirewallPolicy6UrlCategoryArrayOutput)
}

// Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
func (o LookupFirewallPolicy6ResultOutput) Users() GetFirewallPolicy6UserArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6User { return v.Users }).(GetFirewallPolicy6UserArrayOutput)
}

// Enable AV/web/ips protection profile.
func (o LookupFirewallPolicy6ResultOutput) UtmStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.UtmStatus }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupFirewallPolicy6ResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest
func (o LookupFirewallPolicy6ResultOutput) VlanCosFwd() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.VlanCosFwd }).(pulumi.IntOutput)
}

// VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest
func (o LookupFirewallPolicy6ResultOutput) VlanCosRev() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.VlanCosRev }).(pulumi.IntOutput)
}

// Set VLAN filters.
func (o LookupFirewallPolicy6ResultOutput) VlanFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.VlanFilter }).(pulumi.StringOutput)
}

// Name of an existing VoIP profile.
func (o LookupFirewallPolicy6ResultOutput) VoipProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.VoipProfile }).(pulumi.StringOutput)
}

// Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
func (o LookupFirewallPolicy6ResultOutput) Vpntunnel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Vpntunnel }).(pulumi.StringOutput)
}

// Name of an existing Web application firewall profile.
func (o LookupFirewallPolicy6ResultOutput) WafProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WafProfile }).(pulumi.StringOutput)
}

// Enable/disable web cache.
func (o LookupFirewallPolicy6ResultOutput) Webcache() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Webcache }).(pulumi.StringOutput)
}

// Enable/disable web cache for HTTPS.
func (o LookupFirewallPolicy6ResultOutput) WebcacheHttps() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WebcacheHttps }).(pulumi.StringOutput)
}

// Name of an existing Web filter profile.
func (o LookupFirewallPolicy6ResultOutput) WebfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WebfilterProfile }).(pulumi.StringOutput)
}

// Web proxy forward server name.
func (o LookupFirewallPolicy6ResultOutput) WebproxyForwardServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WebproxyForwardServer }).(pulumi.StringOutput)
}

// Webproxy profile name.
func (o LookupFirewallPolicy6ResultOutput) WebproxyProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WebproxyProfile }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallPolicy6ResultOutput{})
}
