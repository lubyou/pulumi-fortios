// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall policy
func LookupFirewallPolicy(ctx *pulumi.Context, args *LookupFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyResult, error) {
	var rv LookupFirewallPolicyResult
	err := ctx.Invoke("fortios:index/getFirewallPolicy:GetFirewallPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallPolicy.
type LookupFirewallPolicyArgs struct {
	// Specify the policyid of the desired firewall policy.
	Policyid int `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallPolicy.
type LookupFirewallPolicyResult struct {
	// Policy action (allow/deny/ipsec).
	Action string `pulumi:"action"`
	// Enable/disable anti-replay check.
	AntiReplay string `pulumi:"antiReplay"`
	// Application category ID list. The structure of `appCategory` block is documented below.
	AppCategories []GetFirewallPolicyAppCategory `pulumi:"appCategories"`
	// Application group names. The structure of `appGroup` block is documented below.
	AppGroups []GetFirewallPolicyAppGroup `pulumi:"appGroups"`
	// Name of an existing Application list.
	ApplicationList string `pulumi:"applicationList"`
	// Application ID list. The structure of `application` block is documented below.
	Applications []GetFirewallPolicyApplication `pulumi:"applications"`
	// HTTPS server certificate for policy authentication.
	AuthCert string `pulumi:"authCert"`
	// Enable/disable authentication-based routing.
	AuthPath string `pulumi:"authPath"`
	// HTTP-to-HTTPS redirect address for firewall authentication.
	AuthRedirectAddr string `pulumi:"authRedirectAddr"`
	// Enable/disable policy traffic ASIC offloading.
	AutoAsicOffload string `pulumi:"autoAsicOffload"`
	// Name of an existing Antivirus profile.
	AvProfile string `pulumi:"avProfile"`
	// Enable/disable block notification.
	BlockNotification string `pulumi:"blockNotification"`
	// Enable to exempt some users from the captive portal.
	CaptivePortalExempt string `pulumi:"captivePortalExempt"`
	// Enable/disable capture packets.
	CapturePacket string `pulumi:"capturePacket"`
	// Name of an existing CIFS profile.
	CifsProfile string `pulumi:"cifsProfile"`
	// Comment.
	Comments string `pulumi:"comments"`
	// Custom fields to append to log messages for this policy. The structure of `customLogFields` block is documented below.
	CustomLogFields []GetFirewallPolicyCustomLogField `pulumi:"customLogFields"`
	// Decrypted traffic mirror.
	DecryptedTrafficMirror string `pulumi:"decryptedTrafficMirror"`
	// Enable TCP NPU session delay to guarantee packet order of 3-way handshake.
	DelayTcpNpuSession string `pulumi:"delayTcpNpuSession"`
	// Names of devices or device groups that can be matched by the policy. The structure of `devices` block is documented below.
	Devices []GetFirewallPolicyDevice `pulumi:"devices"`
	// Enable to change packet's DiffServ values to the specified diffservcode-forward value.
	DiffservForward string `pulumi:"diffservForward"`
	// Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
	DiffservReverse string `pulumi:"diffservReverse"`
	// Change packet's DiffServ to this value.
	DiffservcodeForward string `pulumi:"diffservcodeForward"`
	// Change packet's reverse (reply) DiffServ to this value.
	DiffservcodeRev string `pulumi:"diffservcodeRev"`
	// Enable/disable user authentication disclaimer.
	Disclaimer string `pulumi:"disclaimer"`
	// Name of an existing DLP sensor.
	DlpSensor string `pulumi:"dlpSensor"`
	// Name of an existing DNS filter profile.
	DnsfilterProfile string `pulumi:"dnsfilterProfile"`
	// Enable DSRI to ignore HTTP server responses.
	Dsri string `pulumi:"dsri"`
	// Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.
	Dstaddr6s []GetFirewallPolicyDstaddr6 `pulumi:"dstaddr6s"`
	// When enabled dstaddr specifies what the destination address must NOT be.
	DstaddrNegate string `pulumi:"dstaddrNegate"`
	// Destination address and address group names. The structure of `dstaddr` block is documented below.
	Dstaddrs []GetFirewallPolicyDstaddr `pulumi:"dstaddrs"`
	// Outgoing (egress) interface. The structure of `dstintf` block is documented below.
	Dstintfs []GetFirewallPolicyDstintf `pulumi:"dstintfs"`
	// Enable/disable email collection.
	EmailCollect string `pulumi:"emailCollect"`
	// Name of an existing email filter profile.
	EmailfilterProfile string `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile string `pulumi:"fileFilterProfile"`
	// How to handle sessions if the configuration of this firewall policy changes.
	FirewallSessionDirty string `pulumi:"firewallSessionDirty"`
	// Enable to prevent source NAT from changing a session's source port.
	Fixedport string `pulumi:"fixedport"`
	// Enable/disable Fortinet Single Sign-On.
	Fsso string `pulumi:"fsso"`
	// FSSO agent to use for NTLM authentication.
	FssoAgentForNtlm string `pulumi:"fssoAgentForNtlm"`
	// Names of FSSO groups. The structure of `fssoGroups` block is documented below.
	FssoGroups []GetFirewallPolicyFssoGroup `pulumi:"fssoGroups"`
	// Enable/disable recognition of anycast IP addresses using the geography IP database.
	GeoipAnycast string `pulumi:"geoipAnycast"`
	// Match geography address based either on its physical location or registered location.
	GeoipMatch string `pulumi:"geoipMatch"`
	// Label for the policy that appears when the GUI is in Global View mode.
	GlobalLabel string `pulumi:"globalLabel"`
	// Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
	Groups []GetFirewallPolicyGroup `pulumi:"groups"`
	// Redirect HTTP(S) traffic to matching transparent web proxy policy.
	HttpPolicyRedirect string `pulumi:"httpPolicyRedirect"`
	// Name of an existing ICAP profile.
	IcapProfile string `pulumi:"icapProfile"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of identity-based routing rule.
	IdentityBasedRoute string `pulumi:"identityBasedRoute"`
	// Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
	Inbound string `pulumi:"inbound"`
	// Policy inspection mode (Flow/proxy). Default is Flow mode.
	InspectionMode string `pulumi:"inspectionMode"`
	// Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
	InternetService string `pulumi:"internetService"`
	// Custom Internet Service group name. The structure of `internetServiceCustomGroup` block is documented below.
	InternetServiceCustomGroups []GetFirewallPolicyInternetServiceCustomGroup `pulumi:"internetServiceCustomGroups"`
	// Custom Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms []GetFirewallPolicyInternetServiceCustom `pulumi:"internetServiceCustoms"`
	// Internet Service group name. The structure of `internetServiceGroup` block is documented below.
	InternetServiceGroups []GetFirewallPolicyInternetServiceGroup `pulumi:"internetServiceGroups"`
	// Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds []GetFirewallPolicyInternetServiceId `pulumi:"internetServiceIds"`
	// Internet Service name. The structure of `internetServiceName` block is documented below.
	InternetServiceNames []GetFirewallPolicyInternetServiceName `pulumi:"internetServiceNames"`
	// When enabled internet-service specifies what the service must NOT be.
	InternetServiceNegate string `pulumi:"internetServiceNegate"`
	// Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.
	InternetServiceSrc string `pulumi:"internetServiceSrc"`
	// Custom Internet Service source group name. The structure of `internetServiceSrcCustomGroup` block is documented below.
	InternetServiceSrcCustomGroups []GetFirewallPolicyInternetServiceSrcCustomGroup `pulumi:"internetServiceSrcCustomGroups"`
	// Custom Internet Service source name. The structure of `internetServiceSrcCustom` block is documented below.
	InternetServiceSrcCustoms []GetFirewallPolicyInternetServiceSrcCustom `pulumi:"internetServiceSrcCustoms"`
	// Internet Service source group name. The structure of `internetServiceSrcGroup` block is documented below.
	InternetServiceSrcGroups []GetFirewallPolicyInternetServiceSrcGroup `pulumi:"internetServiceSrcGroups"`
	// Internet Service source ID. The structure of `internetServiceSrcId` block is documented below.
	InternetServiceSrcIds []GetFirewallPolicyInternetServiceSrcId `pulumi:"internetServiceSrcIds"`
	// Internet Service source name. The structure of `internetServiceSrcName` block is documented below.
	InternetServiceSrcNames []GetFirewallPolicyInternetServiceSrcName `pulumi:"internetServiceSrcNames"`
	// When enabled internet-service-src specifies what the service must NOT be.
	InternetServiceSrcNegate string `pulumi:"internetServiceSrcNegate"`
	// Enable to use IP Pools for source NAT.
	Ippool string `pulumi:"ippool"`
	// Name of an existing IPS sensor.
	IpsSensor string `pulumi:"ipsSensor"`
	// Label for the policy that appears when the GUI is in Section View mode.
	Label string `pulumi:"label"`
	// Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated.
	LearningMode string `pulumi:"learningMode"`
	// Enable or disable logging. Log all sessions or security profile sessions.
	Logtraffic string `pulumi:"logtraffic"`
	// Record logs when a session starts.
	LogtrafficStart string `pulumi:"logtrafficStart"`
	// Enable to match packets that have had their destination addresses changed by a VIP.
	MatchVip string `pulumi:"matchVip"`
	// Enable/disable matching of only those packets that have had their destination addresses changed by a VIP.
	MatchVipOnly string `pulumi:"matchVipOnly"`
	// Mirror Interface name.
	Name string `pulumi:"name"`
	// Enable/disable source NAT.
	Nat string `pulumi:"nat"`
	// Policy-based IPsec VPN: apply destination NAT to inbound traffic.
	Natinbound string `pulumi:"natinbound"`
	// Policy-based IPsec VPN: source NAT IP address for outgoing traffic.
	Natip string `pulumi:"natip"`
	// Policy-based IPsec VPN: apply source NAT to outbound traffic.
	Natoutbound string `pulumi:"natoutbound"`
	// Enable/disable NTLM authentication.
	Ntlm string `pulumi:"ntlm"`
	// HTTP-User-Agent value of supported browsers. The structure of `ntlmEnabledBrowsers` block is documented below.
	NtlmEnabledBrowsers []GetFirewallPolicyNtlmEnabledBrowser `pulumi:"ntlmEnabledBrowsers"`
	// Enable/disable NTLM guest user access.
	NtlmGuest string `pulumi:"ntlmGuest"`
	// Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
	Outbound string `pulumi:"outbound"`
	// Per-IP traffic shaper.
	PerIpShaper string `pulumi:"perIpShaper"`
	// Accept UDP packets from any host.
	PermitAnyHost string `pulumi:"permitAnyHost"`
	// Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host.
	PermitStunHost string `pulumi:"permitStunHost"`
	// Policy ID.
	Policyid int `pulumi:"policyid"`
	// IPv6 pool names. The structure of `poolname6` block is documented below.
	Poolname6s []GetFirewallPolicyPoolname6 `pulumi:"poolname6s"`
	// IP Pool names. The structure of `poolname` block is documented below.
	Poolnames []GetFirewallPolicyPoolname `pulumi:"poolnames"`
	// Name of profile group.
	ProfileGroup string `pulumi:"profileGroup"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions string `pulumi:"profileProtocolOptions"`
	// Determine whether the firewall policy allows security profile groups or single profiles only.
	ProfileType string `pulumi:"profileType"`
	// Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server.
	RadiusMacAuthBypass string `pulumi:"radiusMacAuthBypass"`
	// URL users are directed to after seeing and accepting the disclaimer or authenticating.
	RedirectUrl string `pulumi:"redirectUrl"`
	// Override the default replacement message group for this policy.
	ReplacemsgOverrideGroup string `pulumi:"replacemsgOverrideGroup"`
	// Direction of the initial traffic for reputation to take effect.
	ReputationDirection string `pulumi:"reputationDirection"`
	// Minimum Reputation to take action.
	ReputationMinimum int `pulumi:"reputationMinimum"`
	// Enable/disable RADIUS single sign-on (RSSO).
	Rsso string `pulumi:"rsso"`
	// Address names if this is an RTP NAT policy. The structure of `rtpAddr` block is documented below.
	RtpAddrs []GetFirewallPolicyRtpAddr `pulumi:"rtpAddrs"`
	// Enable Real Time Protocol (RTP) NAT.
	RtpNat string `pulumi:"rtpNat"`
	// Block or monitor connections to Botnet servers or disable Botnet scanning.
	ScanBotnetConnections string `pulumi:"scanBotnetConnections"`
	// Schedule name.
	Schedule string `pulumi:"schedule"`
	// Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity.
	ScheduleTimeout string `pulumi:"scheduleTimeout"`
	// Enable to send a reply when a session is denied or blocked by a firewall policy.
	SendDenyPacket string `pulumi:"sendDenyPacket"`
	// When enabled service specifies what the service must NOT be.
	ServiceNegate string `pulumi:"serviceNegate"`
	// Service and service group names. The structure of `service` block is documented below.
	Services []GetFirewallPolicyService `pulumi:"services"`
	// TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
	SessionTtl int `pulumi:"sessionTtl"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile string `pulumi:"spamfilterProfile"`
	// Vendor MAC source ID. The structure of `srcVendorMac` block is documented below.
	SrcVendorMacs []GetFirewallPolicySrcVendorMac `pulumi:"srcVendorMacs"`
	// Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.
	Srcaddr6s []GetFirewallPolicySrcaddr6 `pulumi:"srcaddr6s"`
	// When enabled srcaddr specifies what the source address must NOT be.
	SrcaddrNegate string `pulumi:"srcaddrNegate"`
	// Source address and address group names. The structure of `srcaddr` block is documented below.
	Srcaddrs []GetFirewallPolicySrcaddr `pulumi:"srcaddrs"`
	// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
	Srcintfs []GetFirewallPolicySrcintf `pulumi:"srcintfs"`
	// Name of an existing SSH filter profile.
	SshFilterProfile string `pulumi:"sshFilterProfile"`
	// Redirect SSH traffic to matching transparent proxy policy.
	SshPolicyRedirect string `pulumi:"sshPolicyRedirect"`
	// Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring).
	SslMirror string `pulumi:"sslMirror"`
	// SSL mirror interface name. The structure of `sslMirrorIntf` block is documented below.
	SslMirrorIntfs []GetFirewallPolicySslMirrorIntf `pulumi:"sslMirrorIntfs"`
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
	// Traffic shaper.
	TrafficShaper string `pulumi:"trafficShaper"`
	// Reverse traffic shaper.
	TrafficShaperReverse string `pulumi:"trafficShaperReverse"`
	// URL category ID list. The structure of `urlCategory` block is documented below.
	UrlCategories []GetFirewallPolicyUrlCategory `pulumi:"urlCategories"`
	// Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
	Users []GetFirewallPolicyUser `pulumi:"users"`
	// Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
	UtmStatus string `pulumi:"utmStatus"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.
	VlanCosFwd int `pulumi:"vlanCosFwd"`
	// VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.
	VlanCosRev int `pulumi:"vlanCosRev"`
	// Set VLAN filters.
	VlanFilter string `pulumi:"vlanFilter"`
	// Name of an existing VoIP profile.
	VoipProfile string `pulumi:"voipProfile"`
	// Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
	Vpntunnel string `pulumi:"vpntunnel"`
	// Name of an existing Web application firewall profile.
	WafProfile string `pulumi:"wafProfile"`
	// Enable/disable WAN optimization.
	Wanopt string `pulumi:"wanopt"`
	// WAN optimization auto-detection mode.
	WanoptDetection string `pulumi:"wanoptDetection"`
	// WAN optimization passive mode options. This option decides what IP address will be used to connect server.
	WanoptPassiveOpt string `pulumi:"wanoptPassiveOpt"`
	// WAN optimization peer.
	WanoptPeer string `pulumi:"wanoptPeer"`
	// WAN optimization profile.
	WanoptProfile string `pulumi:"wanoptProfile"`
	// Enable/disable forwarding traffic matching this policy to a configured WCCP server.
	Wccp string `pulumi:"wccp"`
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
	// Enable/disable WiFi Single Sign On (WSSO).
	Wsso string `pulumi:"wsso"`
}