// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios firewall proxypolicy
func LookupFirewallProxyPolicy(ctx *pulumi.Context, args *LookupFirewallProxyPolicyArgs, opts ...pulumi.InvokeOption) (*LookupFirewallProxyPolicyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFirewallProxyPolicyResult
	err := ctx.Invoke("fortios:index/getFirewallProxyPolicy:GetFirewallProxyPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallProxyPolicy.
type LookupFirewallProxyPolicyArgs struct {
	// Specify the policyid of the desired firewall proxypolicy.
	Policyid int `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallProxyPolicy.
type LookupFirewallProxyPolicyResult struct {
	// IPv4 access proxy. The structure of `accessProxy` block is documented below.
	AccessProxies []GetFirewallProxyPolicyAccessProxy `pulumi:"accessProxies"`
	// IPv6 access proxy. The structure of `accessProxy6` block is documented below.
	AccessProxy6s []GetFirewallProxyPolicyAccessProxy6 `pulumi:"accessProxy6s"`
	// Accept or deny traffic matching the policy parameters.
	Action string `pulumi:"action"`
	// Name of an existing Application list.
	ApplicationList string `pulumi:"applicationList"`
	// Name of an existing Antivirus profile.
	AvProfile string `pulumi:"avProfile"`
	// Enable/disable block notification.
	BlockNotification string `pulumi:"blockNotification"`
	// Name of an existing CIFS profile.
	CifsProfile string `pulumi:"cifsProfile"`
	// Optional comments.
	Comments string `pulumi:"comments"`
	// Decrypted traffic mirror.
	DecryptedTrafficMirror string `pulumi:"decryptedTrafficMirror"`
	// When enabled, the ownership enforcement will be done at policy level.
	DeviceOwnership string `pulumi:"deviceOwnership"`
	// Web proxy disclaimer setting: by domain, policy, or user.
	Disclaimer string `pulumi:"disclaimer"`
	// Name of an existing DLP sensor.
	DlpSensor string `pulumi:"dlpSensor"`
	// IPv6 destination address objects. The structure of `dstaddr6` block is documented below.
	Dstaddr6s []GetFirewallProxyPolicyDstaddr6 `pulumi:"dstaddr6s"`
	// When enabled, destination addresses match against any address EXCEPT the specified destination addresses.
	DstaddrNegate string `pulumi:"dstaddrNegate"`
	// Destination address objects. The structure of `dstaddr` block is documented below.
	Dstaddrs []GetFirewallProxyPolicyDstaddr `pulumi:"dstaddrs"`
	// Destination interface names. The structure of `dstintf` block is documented below.
	Dstintfs []GetFirewallProxyPolicyDstintf `pulumi:"dstintfs"`
	// Name of an existing email filter profile.
	EmailfilterProfile string `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile string `pulumi:"fileFilterProfile"`
	// Global web-based manager visible label.
	GlobalLabel string `pulumi:"globalLabel"`
	// Names of group objects. The structure of `groups` block is documented below.
	Groups []GetFirewallProxyPolicyGroup `pulumi:"groups"`
	// Enable/disable HTTP tunnel authentication.
	HttpTunnelAuth string `pulumi:"httpTunnelAuth"`
	// Name of an existing ICAP profile.
	IcapProfile string `pulumi:"icapProfile"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
	InternetService string `pulumi:"internetService"`
	// Custom Internet Service group name. The structure of `internetServiceCustomGroup` block is documented below.
	InternetServiceCustomGroups []GetFirewallProxyPolicyInternetServiceCustomGroup `pulumi:"internetServiceCustomGroups"`
	// Custom Internet Service name. The structure of `internetServiceCustom` block is documented below.
	InternetServiceCustoms []GetFirewallProxyPolicyInternetServiceCustom `pulumi:"internetServiceCustoms"`
	// Internet Service group name. The structure of `internetServiceGroup` block is documented below.
	InternetServiceGroups []GetFirewallProxyPolicyInternetServiceGroup `pulumi:"internetServiceGroups"`
	// Internet Service ID. The structure of `internetServiceId` block is documented below.
	InternetServiceIds []GetFirewallProxyPolicyInternetServiceId `pulumi:"internetServiceIds"`
	// Internet Service name. The structure of `internetServiceName` block is documented below.
	InternetServiceNames []GetFirewallProxyPolicyInternetServiceName `pulumi:"internetServiceNames"`
	// When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.
	InternetServiceNegate string `pulumi:"internetServiceNegate"`
	// Name of an existing IPS sensor.
	IpsSensor string `pulumi:"ipsSensor"`
	// VDOM-specific GUI visible label.
	Label string `pulumi:"label"`
	// Enable/disable logging traffic through the policy.
	Logtraffic string `pulumi:"logtraffic"`
	// Enable/disable policy log traffic start.
	LogtrafficStart string `pulumi:"logtrafficStart"`
	// Group name.
	Name string `pulumi:"name"`
	// Policy ID.
	Policyid int `pulumi:"policyid"`
	// Name of IP pool object. The structure of `poolname` block is documented below.
	Poolnames []GetFirewallProxyPolicyPoolname `pulumi:"poolnames"`
	// Name of profile group.
	ProfileGroup string `pulumi:"profileGroup"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions string `pulumi:"profileProtocolOptions"`
	// Determine whether the firewall policy allows security profile groups or single profiles only.
	ProfileType string `pulumi:"profileType"`
	// Type of explicit proxy.
	Proxy string `pulumi:"proxy"`
	// Redirect URL for further explicit web proxy processing.
	RedirectUrl string `pulumi:"redirectUrl"`
	// Authentication replacement message override group.
	ReplacemsgOverrideGroup string `pulumi:"replacemsgOverrideGroup"`
	// Enable/disable scanning of connections to Botnet servers.
	ScanBotnetConnections string `pulumi:"scanBotnetConnections"`
	// Name of schedule object.
	Schedule string `pulumi:"schedule"`
	// Name of an existing SCTP filter profile.
	SctpFilterProfile string `pulumi:"sctpFilterProfile"`
	// When enabled, services match against any service EXCEPT the specified destination services.
	ServiceNegate string `pulumi:"serviceNegate"`
	// Name of service objects. The structure of `service` block is documented below.
	Services []GetFirewallProxyPolicyService `pulumi:"services"`
	// TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
	SessionTtl int `pulumi:"sessionTtl"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile string `pulumi:"spamfilterProfile"`
	// IPv6 source address objects. The structure of `srcaddr6` block is documented below.
	Srcaddr6s []GetFirewallProxyPolicySrcaddr6 `pulumi:"srcaddr6s"`
	// When enabled, source addresses match against any address EXCEPT the specified source addresses.
	SrcaddrNegate string `pulumi:"srcaddrNegate"`
	// Source address objects. The structure of `srcaddr` block is documented below.
	Srcaddrs []GetFirewallProxyPolicySrcaddr `pulumi:"srcaddrs"`
	// Source interface names. The structure of `srcintf` block is documented below.
	Srcintfs []GetFirewallProxyPolicySrcintf `pulumi:"srcintfs"`
	// Name of an existing SSH filter profile.
	SshFilterProfile string `pulumi:"sshFilterProfile"`
	// Redirect SSH traffic to matching transparent proxy policy.
	SshPolicyRedirect string `pulumi:"sshPolicyRedirect"`
	// Name of an existing SSL SSH profile.
	SslSshProfile string `pulumi:"sslSshProfile"`
	// Enable/disable the active status of the policy.
	Status string `pulumi:"status"`
	// Enable to use the IP address of the client to connect to the server.
	Transparent string `pulumi:"transparent"`
	// Names of user objects. The structure of `users` block is documented below.
	Users []GetFirewallProxyPolicyUser `pulumi:"users"`
	// Enable the use of UTM profiles/sensors/lists.
	UtmStatus string `pulumi:"utmStatus"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid      string  `pulumi:"uuid"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Name of an existing VideoFilter profile.
	VideofilterProfile string `pulumi:"videofilterProfile"`
	// Name of an existing VoIP profile.
	VoipProfile string `pulumi:"voipProfile"`
	// Name of an existing Web application firewall profile.
	WafProfile string `pulumi:"wafProfile"`
	// Enable/disable web caching.
	Webcache string `pulumi:"webcache"`
	// Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).
	WebcacheHttps string `pulumi:"webcacheHttps"`
	// Name of an existing Web filter profile.
	WebfilterProfile string `pulumi:"webfilterProfile"`
	// Web proxy forward server name.
	WebproxyForwardServer string `pulumi:"webproxyForwardServer"`
	// Name of web proxy profile.
	WebproxyProfile string `pulumi:"webproxyProfile"`
	// ZTNA EMS Tag names. The structure of `ztnaEmsTag` block is documented below.
	ZtnaEmsTags []GetFirewallProxyPolicyZtnaEmsTag `pulumi:"ztnaEmsTags"`
	// ZTNA tag matching logic.
	ZtnaTagsMatchLogic string `pulumi:"ztnaTagsMatchLogic"`
}

func LookupFirewallProxyPolicyOutput(ctx *pulumi.Context, args LookupFirewallProxyPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallProxyPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallProxyPolicyResult, error) {
			args := v.(LookupFirewallProxyPolicyArgs)
			r, err := LookupFirewallProxyPolicy(ctx, &args, opts...)
			var s LookupFirewallProxyPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallProxyPolicyResultOutput)
}

// A collection of arguments for invoking GetFirewallProxyPolicy.
type LookupFirewallProxyPolicyOutputArgs struct {
	// Specify the policyid of the desired firewall proxypolicy.
	Policyid pulumi.IntInput `pulumi:"policyid"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupFirewallProxyPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallProxyPolicyArgs)(nil)).Elem()
}

// A collection of values returned by GetFirewallProxyPolicy.
type LookupFirewallProxyPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallProxyPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallProxyPolicyResult)(nil)).Elem()
}

func (o LookupFirewallProxyPolicyResultOutput) ToLookupFirewallProxyPolicyResultOutput() LookupFirewallProxyPolicyResultOutput {
	return o
}

func (o LookupFirewallProxyPolicyResultOutput) ToLookupFirewallProxyPolicyResultOutputWithContext(ctx context.Context) LookupFirewallProxyPolicyResultOutput {
	return o
}

// IPv4 access proxy. The structure of `accessProxy` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) AccessProxies() GetFirewallProxyPolicyAccessProxyArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyAccessProxy { return v.AccessProxies }).(GetFirewallProxyPolicyAccessProxyArrayOutput)
}

// IPv6 access proxy. The structure of `accessProxy6` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) AccessProxy6s() GetFirewallProxyPolicyAccessProxy6ArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyAccessProxy6 { return v.AccessProxy6s }).(GetFirewallProxyPolicyAccessProxy6ArrayOutput)
}

// Accept or deny traffic matching the policy parameters.
func (o LookupFirewallProxyPolicyResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Action }).(pulumi.StringOutput)
}

// Name of an existing Application list.
func (o LookupFirewallProxyPolicyResultOutput) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.ApplicationList }).(pulumi.StringOutput)
}

// Name of an existing Antivirus profile.
func (o LookupFirewallProxyPolicyResultOutput) AvProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.AvProfile }).(pulumi.StringOutput)
}

// Enable/disable block notification.
func (o LookupFirewallProxyPolicyResultOutput) BlockNotification() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.BlockNotification }).(pulumi.StringOutput)
}

// Name of an existing CIFS profile.
func (o LookupFirewallProxyPolicyResultOutput) CifsProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.CifsProfile }).(pulumi.StringOutput)
}

// Optional comments.
func (o LookupFirewallProxyPolicyResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Comments }).(pulumi.StringOutput)
}

// Decrypted traffic mirror.
func (o LookupFirewallProxyPolicyResultOutput) DecryptedTrafficMirror() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.DecryptedTrafficMirror }).(pulumi.StringOutput)
}

// When enabled, the ownership enforcement will be done at policy level.
func (o LookupFirewallProxyPolicyResultOutput) DeviceOwnership() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.DeviceOwnership }).(pulumi.StringOutput)
}

// Web proxy disclaimer setting: by domain, policy, or user.
func (o LookupFirewallProxyPolicyResultOutput) Disclaimer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Disclaimer }).(pulumi.StringOutput)
}

// Name of an existing DLP sensor.
func (o LookupFirewallProxyPolicyResultOutput) DlpSensor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.DlpSensor }).(pulumi.StringOutput)
}

// IPv6 destination address objects. The structure of `dstaddr6` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Dstaddr6s() GetFirewallProxyPolicyDstaddr6ArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyDstaddr6 { return v.Dstaddr6s }).(GetFirewallProxyPolicyDstaddr6ArrayOutput)
}

// When enabled, destination addresses match against any address EXCEPT the specified destination addresses.
func (o LookupFirewallProxyPolicyResultOutput) DstaddrNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.DstaddrNegate }).(pulumi.StringOutput)
}

// Destination address objects. The structure of `dstaddr` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Dstaddrs() GetFirewallProxyPolicyDstaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyDstaddr { return v.Dstaddrs }).(GetFirewallProxyPolicyDstaddrArrayOutput)
}

// Destination interface names. The structure of `dstintf` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Dstintfs() GetFirewallProxyPolicyDstintfArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyDstintf { return v.Dstintfs }).(GetFirewallProxyPolicyDstintfArrayOutput)
}

// Name of an existing email filter profile.
func (o LookupFirewallProxyPolicyResultOutput) EmailfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.EmailfilterProfile }).(pulumi.StringOutput)
}

// Name of an existing file-filter profile.
func (o LookupFirewallProxyPolicyResultOutput) FileFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.FileFilterProfile }).(pulumi.StringOutput)
}

// Global web-based manager visible label.
func (o LookupFirewallProxyPolicyResultOutput) GlobalLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.GlobalLabel }).(pulumi.StringOutput)
}

// Names of group objects. The structure of `groups` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Groups() GetFirewallProxyPolicyGroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyGroup { return v.Groups }).(GetFirewallProxyPolicyGroupArrayOutput)
}

// Enable/disable HTTP tunnel authentication.
func (o LookupFirewallProxyPolicyResultOutput) HttpTunnelAuth() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.HttpTunnelAuth }).(pulumi.StringOutput)
}

// Name of an existing ICAP profile.
func (o LookupFirewallProxyPolicyResultOutput) IcapProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.IcapProfile }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallProxyPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
func (o LookupFirewallProxyPolicyResultOutput) InternetService() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.InternetService }).(pulumi.StringOutput)
}

// Custom Internet Service group name. The structure of `internetServiceCustomGroup` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) InternetServiceCustomGroups() GetFirewallProxyPolicyInternetServiceCustomGroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyInternetServiceCustomGroup {
		return v.InternetServiceCustomGroups
	}).(GetFirewallProxyPolicyInternetServiceCustomGroupArrayOutput)
}

// Custom Internet Service name. The structure of `internetServiceCustom` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) InternetServiceCustoms() GetFirewallProxyPolicyInternetServiceCustomArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyInternetServiceCustom {
		return v.InternetServiceCustoms
	}).(GetFirewallProxyPolicyInternetServiceCustomArrayOutput)
}

// Internet Service group name. The structure of `internetServiceGroup` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) InternetServiceGroups() GetFirewallProxyPolicyInternetServiceGroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyInternetServiceGroup {
		return v.InternetServiceGroups
	}).(GetFirewallProxyPolicyInternetServiceGroupArrayOutput)
}

// Internet Service ID. The structure of `internetServiceId` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) InternetServiceIds() GetFirewallProxyPolicyInternetServiceIdArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyInternetServiceId {
		return v.InternetServiceIds
	}).(GetFirewallProxyPolicyInternetServiceIdArrayOutput)
}

// Internet Service name. The structure of `internetServiceName` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) InternetServiceNames() GetFirewallProxyPolicyInternetServiceNameArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyInternetServiceName {
		return v.InternetServiceNames
	}).(GetFirewallProxyPolicyInternetServiceNameArrayOutput)
}

// When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.
func (o LookupFirewallProxyPolicyResultOutput) InternetServiceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.InternetServiceNegate }).(pulumi.StringOutput)
}

// Name of an existing IPS sensor.
func (o LookupFirewallProxyPolicyResultOutput) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.IpsSensor }).(pulumi.StringOutput)
}

// VDOM-specific GUI visible label.
func (o LookupFirewallProxyPolicyResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Label }).(pulumi.StringOutput)
}

// Enable/disable logging traffic through the policy.
func (o LookupFirewallProxyPolicyResultOutput) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Logtraffic }).(pulumi.StringOutput)
}

// Enable/disable policy log traffic start.
func (o LookupFirewallProxyPolicyResultOutput) LogtrafficStart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.LogtrafficStart }).(pulumi.StringOutput)
}

// Group name.
func (o LookupFirewallProxyPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy ID.
func (o LookupFirewallProxyPolicyResultOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) int { return v.Policyid }).(pulumi.IntOutput)
}

// Name of IP pool object. The structure of `poolname` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Poolnames() GetFirewallProxyPolicyPoolnameArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyPoolname { return v.Poolnames }).(GetFirewallProxyPolicyPoolnameArrayOutput)
}

// Name of profile group.
func (o LookupFirewallProxyPolicyResultOutput) ProfileGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.ProfileGroup }).(pulumi.StringOutput)
}

// Name of an existing Protocol options profile.
func (o LookupFirewallProxyPolicyResultOutput) ProfileProtocolOptions() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.ProfileProtocolOptions }).(pulumi.StringOutput)
}

// Determine whether the firewall policy allows security profile groups or single profiles only.
func (o LookupFirewallProxyPolicyResultOutput) ProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.ProfileType }).(pulumi.StringOutput)
}

// Type of explicit proxy.
func (o LookupFirewallProxyPolicyResultOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Proxy }).(pulumi.StringOutput)
}

// Redirect URL for further explicit web proxy processing.
func (o LookupFirewallProxyPolicyResultOutput) RedirectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.RedirectUrl }).(pulumi.StringOutput)
}

// Authentication replacement message override group.
func (o LookupFirewallProxyPolicyResultOutput) ReplacemsgOverrideGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.ReplacemsgOverrideGroup }).(pulumi.StringOutput)
}

// Enable/disable scanning of connections to Botnet servers.
func (o LookupFirewallProxyPolicyResultOutput) ScanBotnetConnections() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.ScanBotnetConnections }).(pulumi.StringOutput)
}

// Name of schedule object.
func (o LookupFirewallProxyPolicyResultOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Schedule }).(pulumi.StringOutput)
}

// Name of an existing SCTP filter profile.
func (o LookupFirewallProxyPolicyResultOutput) SctpFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.SctpFilterProfile }).(pulumi.StringOutput)
}

// When enabled, services match against any service EXCEPT the specified destination services.
func (o LookupFirewallProxyPolicyResultOutput) ServiceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.ServiceNegate }).(pulumi.StringOutput)
}

// Name of service objects. The structure of `service` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Services() GetFirewallProxyPolicyServiceArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyService { return v.Services }).(GetFirewallProxyPolicyServiceArrayOutput)
}

// TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
func (o LookupFirewallProxyPolicyResultOutput) SessionTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) int { return v.SessionTtl }).(pulumi.IntOutput)
}

// Name of an existing Spam filter profile.
func (o LookupFirewallProxyPolicyResultOutput) SpamfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.SpamfilterProfile }).(pulumi.StringOutput)
}

// IPv6 source address objects. The structure of `srcaddr6` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Srcaddr6s() GetFirewallProxyPolicySrcaddr6ArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicySrcaddr6 { return v.Srcaddr6s }).(GetFirewallProxyPolicySrcaddr6ArrayOutput)
}

// When enabled, source addresses match against any address EXCEPT the specified source addresses.
func (o LookupFirewallProxyPolicyResultOutput) SrcaddrNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.SrcaddrNegate }).(pulumi.StringOutput)
}

// Source address objects. The structure of `srcaddr` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Srcaddrs() GetFirewallProxyPolicySrcaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicySrcaddr { return v.Srcaddrs }).(GetFirewallProxyPolicySrcaddrArrayOutput)
}

// Source interface names. The structure of `srcintf` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Srcintfs() GetFirewallProxyPolicySrcintfArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicySrcintf { return v.Srcintfs }).(GetFirewallProxyPolicySrcintfArrayOutput)
}

// Name of an existing SSH filter profile.
func (o LookupFirewallProxyPolicyResultOutput) SshFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.SshFilterProfile }).(pulumi.StringOutput)
}

// Redirect SSH traffic to matching transparent proxy policy.
func (o LookupFirewallProxyPolicyResultOutput) SshPolicyRedirect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.SshPolicyRedirect }).(pulumi.StringOutput)
}

// Name of an existing SSL SSH profile.
func (o LookupFirewallProxyPolicyResultOutput) SslSshProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.SslSshProfile }).(pulumi.StringOutput)
}

// Enable/disable the active status of the policy.
func (o LookupFirewallProxyPolicyResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Status }).(pulumi.StringOutput)
}

// Enable to use the IP address of the client to connect to the server.
func (o LookupFirewallProxyPolicyResultOutput) Transparent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Transparent }).(pulumi.StringOutput)
}

// Names of user objects. The structure of `users` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) Users() GetFirewallProxyPolicyUserArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyUser { return v.Users }).(GetFirewallProxyPolicyUserArrayOutput)
}

// Enable the use of UTM profiles/sensors/lists.
func (o LookupFirewallProxyPolicyResultOutput) UtmStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.UtmStatus }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o LookupFirewallProxyPolicyResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupFirewallProxyPolicyResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Name of an existing VideoFilter profile.
func (o LookupFirewallProxyPolicyResultOutput) VideofilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.VideofilterProfile }).(pulumi.StringOutput)
}

// Name of an existing VoIP profile.
func (o LookupFirewallProxyPolicyResultOutput) VoipProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.VoipProfile }).(pulumi.StringOutput)
}

// Name of an existing Web application firewall profile.
func (o LookupFirewallProxyPolicyResultOutput) WafProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.WafProfile }).(pulumi.StringOutput)
}

// Enable/disable web caching.
func (o LookupFirewallProxyPolicyResultOutput) Webcache() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.Webcache }).(pulumi.StringOutput)
}

// Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).
func (o LookupFirewallProxyPolicyResultOutput) WebcacheHttps() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.WebcacheHttps }).(pulumi.StringOutput)
}

// Name of an existing Web filter profile.
func (o LookupFirewallProxyPolicyResultOutput) WebfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.WebfilterProfile }).(pulumi.StringOutput)
}

// Web proxy forward server name.
func (o LookupFirewallProxyPolicyResultOutput) WebproxyForwardServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.WebproxyForwardServer }).(pulumi.StringOutput)
}

// Name of web proxy profile.
func (o LookupFirewallProxyPolicyResultOutput) WebproxyProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.WebproxyProfile }).(pulumi.StringOutput)
}

// ZTNA EMS Tag names. The structure of `ztnaEmsTag` block is documented below.
func (o LookupFirewallProxyPolicyResultOutput) ZtnaEmsTags() GetFirewallProxyPolicyZtnaEmsTagArrayOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) []GetFirewallProxyPolicyZtnaEmsTag { return v.ZtnaEmsTags }).(GetFirewallProxyPolicyZtnaEmsTagArrayOutput)
}

// ZTNA tag matching logic.
func (o LookupFirewallProxyPolicyResultOutput) ZtnaTagsMatchLogic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallProxyPolicyResult) string { return v.ZtnaTagsMatchLogic }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallProxyPolicyResultOutput{})
}
