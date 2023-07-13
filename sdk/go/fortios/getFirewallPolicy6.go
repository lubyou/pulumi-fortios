// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallPolicy6(ctx *pulumi.Context, args *LookupFirewallPolicy6Args, opts ...pulumi.InvokeOption) (*LookupFirewallPolicy6Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallPolicy6Result
	err := ctx.Invoke("fortios:index/getFirewallPolicy6:GetFirewallPolicy6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetFirewallPolicy6.
type LookupFirewallPolicy6Args struct {
	Policyid  int     `pulumi:"policyid"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetFirewallPolicy6.
type LookupFirewallPolicy6Result struct {
	Action                 string                             `pulumi:"action"`
	AntiReplay             string                             `pulumi:"antiReplay"`
	AppCategories          []GetFirewallPolicy6AppCategory    `pulumi:"appCategories"`
	AppGroups              []GetFirewallPolicy6AppGroup       `pulumi:"appGroups"`
	ApplicationList        string                             `pulumi:"applicationList"`
	Applications           []GetFirewallPolicy6Application    `pulumi:"applications"`
	AutoAsicOffload        string                             `pulumi:"autoAsicOffload"`
	AvProfile              string                             `pulumi:"avProfile"`
	CifsProfile            string                             `pulumi:"cifsProfile"`
	Comments               string                             `pulumi:"comments"`
	CustomLogFields        []GetFirewallPolicy6CustomLogField `pulumi:"customLogFields"`
	DecryptedTrafficMirror string                             `pulumi:"decryptedTrafficMirror"`
	Devices                []GetFirewallPolicy6Device         `pulumi:"devices"`
	DiffservForward        string                             `pulumi:"diffservForward"`
	DiffservReverse        string                             `pulumi:"diffservReverse"`
	DiffservcodeForward    string                             `pulumi:"diffservcodeForward"`
	DiffservcodeRev        string                             `pulumi:"diffservcodeRev"`
	DlpSensor              string                             `pulumi:"dlpSensor"`
	DnsfilterProfile       string                             `pulumi:"dnsfilterProfile"`
	Dsri                   string                             `pulumi:"dsri"`
	DstaddrNegate          string                             `pulumi:"dstaddrNegate"`
	Dstaddrs               []GetFirewallPolicy6Dstaddr        `pulumi:"dstaddrs"`
	Dstintfs               []GetFirewallPolicy6Dstintf        `pulumi:"dstintfs"`
	EmailfilterProfile     string                             `pulumi:"emailfilterProfile"`
	FirewallSessionDirty   string                             `pulumi:"firewallSessionDirty"`
	Fixedport              string                             `pulumi:"fixedport"`
	FssoGroups             []GetFirewallPolicy6FssoGroup      `pulumi:"fssoGroups"`
	GlobalLabel            string                             `pulumi:"globalLabel"`
	Groups                 []GetFirewallPolicy6Group          `pulumi:"groups"`
	HttpPolicyRedirect     string                             `pulumi:"httpPolicyRedirect"`
	IcapProfile            string                             `pulumi:"icapProfile"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string                            `pulumi:"id"`
	Inbound                 string                            `pulumi:"inbound"`
	InspectionMode          string                            `pulumi:"inspectionMode"`
	Ippool                  string                            `pulumi:"ippool"`
	IpsSensor               string                            `pulumi:"ipsSensor"`
	Label                   string                            `pulumi:"label"`
	Logtraffic              string                            `pulumi:"logtraffic"`
	LogtrafficStart         string                            `pulumi:"logtrafficStart"`
	Name                    string                            `pulumi:"name"`
	Nat                     string                            `pulumi:"nat"`
	Natinbound              string                            `pulumi:"natinbound"`
	Natoutbound             string                            `pulumi:"natoutbound"`
	Outbound                string                            `pulumi:"outbound"`
	PerIpShaper             string                            `pulumi:"perIpShaper"`
	Policyid                int                               `pulumi:"policyid"`
	Poolnames               []GetFirewallPolicy6Poolname      `pulumi:"poolnames"`
	ProfileGroup            string                            `pulumi:"profileGroup"`
	ProfileProtocolOptions  string                            `pulumi:"profileProtocolOptions"`
	ProfileType             string                            `pulumi:"profileType"`
	ReplacemsgOverrideGroup string                            `pulumi:"replacemsgOverrideGroup"`
	Rsso                    string                            `pulumi:"rsso"`
	Schedule                string                            `pulumi:"schedule"`
	SendDenyPacket          string                            `pulumi:"sendDenyPacket"`
	ServiceNegate           string                            `pulumi:"serviceNegate"`
	Services                []GetFirewallPolicy6Service       `pulumi:"services"`
	SessionTtl              int                               `pulumi:"sessionTtl"`
	SpamfilterProfile       string                            `pulumi:"spamfilterProfile"`
	SrcaddrNegate           string                            `pulumi:"srcaddrNegate"`
	Srcaddrs                []GetFirewallPolicy6Srcaddr       `pulumi:"srcaddrs"`
	Srcintfs                []GetFirewallPolicy6Srcintf       `pulumi:"srcintfs"`
	SshFilterProfile        string                            `pulumi:"sshFilterProfile"`
	SshPolicyRedirect       string                            `pulumi:"sshPolicyRedirect"`
	SslMirror               string                            `pulumi:"sslMirror"`
	SslMirrorIntfs          []GetFirewallPolicy6SslMirrorIntf `pulumi:"sslMirrorIntfs"`
	SslSshProfile           string                            `pulumi:"sslSshProfile"`
	Status                  string                            `pulumi:"status"`
	TcpMssReceiver          int                               `pulumi:"tcpMssReceiver"`
	TcpMssSender            int                               `pulumi:"tcpMssSender"`
	TcpSessionWithoutSyn    string                            `pulumi:"tcpSessionWithoutSyn"`
	TimeoutSendRst          string                            `pulumi:"timeoutSendRst"`
	Tos                     string                            `pulumi:"tos"`
	TosMask                 string                            `pulumi:"tosMask"`
	TosNegate               string                            `pulumi:"tosNegate"`
	TrafficShaper           string                            `pulumi:"trafficShaper"`
	TrafficShaperReverse    string                            `pulumi:"trafficShaperReverse"`
	UrlCategories           []GetFirewallPolicy6UrlCategory   `pulumi:"urlCategories"`
	Users                   []GetFirewallPolicy6User          `pulumi:"users"`
	UtmStatus               string                            `pulumi:"utmStatus"`
	Uuid                    string                            `pulumi:"uuid"`
	Vdomparam               *string                           `pulumi:"vdomparam"`
	VlanCosFwd              int                               `pulumi:"vlanCosFwd"`
	VlanCosRev              int                               `pulumi:"vlanCosRev"`
	VlanFilter              string                            `pulumi:"vlanFilter"`
	VoipProfile             string                            `pulumi:"voipProfile"`
	Vpntunnel               string                            `pulumi:"vpntunnel"`
	WafProfile              string                            `pulumi:"wafProfile"`
	Webcache                string                            `pulumi:"webcache"`
	WebcacheHttps           string                            `pulumi:"webcacheHttps"`
	WebfilterProfile        string                            `pulumi:"webfilterProfile"`
	WebproxyForwardServer   string                            `pulumi:"webproxyForwardServer"`
	WebproxyProfile         string                            `pulumi:"webproxyProfile"`
}

func LookupFirewallPolicy6Output(ctx *pulumi.Context, args LookupFirewallPolicy6OutputArgs, opts ...pulumi.InvokeOption) LookupFirewallPolicy6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallPolicy6Result, error) {
			args := v.(LookupFirewallPolicy6Args)
			r, err := LookupFirewallPolicy6(ctx, &args, opts...)
			var s LookupFirewallPolicy6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallPolicy6ResultOutput)
}

// A collection of arguments for invoking GetFirewallPolicy6.
type LookupFirewallPolicy6OutputArgs struct {
	Policyid  pulumi.IntInput       `pulumi:"policyid"`
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

func (o LookupFirewallPolicy6ResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Action }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) AntiReplay() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.AntiReplay }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) AppCategories() GetFirewallPolicy6AppCategoryArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6AppCategory { return v.AppCategories }).(GetFirewallPolicy6AppCategoryArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) AppGroups() GetFirewallPolicy6AppGroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6AppGroup { return v.AppGroups }).(GetFirewallPolicy6AppGroupArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ApplicationList }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Applications() GetFirewallPolicy6ApplicationArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Application { return v.Applications }).(GetFirewallPolicy6ApplicationArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) AutoAsicOffload() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.AutoAsicOffload }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) AvProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.AvProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) CifsProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.CifsProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Comments }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) CustomLogFields() GetFirewallPolicy6CustomLogFieldArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6CustomLogField { return v.CustomLogFields }).(GetFirewallPolicy6CustomLogFieldArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) DecryptedTrafficMirror() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DecryptedTrafficMirror }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Devices() GetFirewallPolicy6DeviceArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Device { return v.Devices }).(GetFirewallPolicy6DeviceArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) DiffservForward() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DiffservForward }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) DiffservReverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DiffservReverse }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) DiffservcodeForward() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DiffservcodeForward }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) DiffservcodeRev() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DiffservcodeRev }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) DlpSensor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DlpSensor }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) DnsfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DnsfilterProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Dsri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Dsri }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) DstaddrNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.DstaddrNegate }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Dstaddrs() GetFirewallPolicy6DstaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Dstaddr { return v.Dstaddrs }).(GetFirewallPolicy6DstaddrArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Dstintfs() GetFirewallPolicy6DstintfArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Dstintf { return v.Dstintfs }).(GetFirewallPolicy6DstintfArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) EmailfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.EmailfilterProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) FirewallSessionDirty() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.FirewallSessionDirty }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Fixedport() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Fixedport }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) FssoGroups() GetFirewallPolicy6FssoGroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6FssoGroup { return v.FssoGroups }).(GetFirewallPolicy6FssoGroupArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) GlobalLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.GlobalLabel }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Groups() GetFirewallPolicy6GroupArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Group { return v.Groups }).(GetFirewallPolicy6GroupArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) HttpPolicyRedirect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.HttpPolicyRedirect }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) IcapProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.IcapProfile }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallPolicy6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Inbound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Inbound }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) InspectionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.InspectionMode }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Ippool() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Ippool }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.IpsSensor }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Label }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Logtraffic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Logtraffic }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) LogtrafficStart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.LogtrafficStart }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Nat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Nat }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Natinbound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Natinbound }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Natoutbound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Natoutbound }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Outbound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Outbound }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) PerIpShaper() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.PerIpShaper }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Policyid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.Policyid }).(pulumi.IntOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Poolnames() GetFirewallPolicy6PoolnameArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Poolname { return v.Poolnames }).(GetFirewallPolicy6PoolnameArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) ProfileGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ProfileGroup }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) ProfileProtocolOptions() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ProfileProtocolOptions }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) ProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ProfileType }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) ReplacemsgOverrideGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ReplacemsgOverrideGroup }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Rsso() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Rsso }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Schedule }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SendDenyPacket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SendDenyPacket }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) ServiceNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.ServiceNegate }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Services() GetFirewallPolicy6ServiceArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Service { return v.Services }).(GetFirewallPolicy6ServiceArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SessionTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.SessionTtl }).(pulumi.IntOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SpamfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SpamfilterProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SrcaddrNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SrcaddrNegate }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Srcaddrs() GetFirewallPolicy6SrcaddrArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Srcaddr { return v.Srcaddrs }).(GetFirewallPolicy6SrcaddrArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Srcintfs() GetFirewallPolicy6SrcintfArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6Srcintf { return v.Srcintfs }).(GetFirewallPolicy6SrcintfArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SshFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SshFilterProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SshPolicyRedirect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SshPolicyRedirect }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SslMirror() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SslMirror }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SslMirrorIntfs() GetFirewallPolicy6SslMirrorIntfArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6SslMirrorIntf { return v.SslMirrorIntfs }).(GetFirewallPolicy6SslMirrorIntfArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) SslSshProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.SslSshProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) TcpMssReceiver() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.TcpMssReceiver }).(pulumi.IntOutput)
}

func (o LookupFirewallPolicy6ResultOutput) TcpMssSender() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.TcpMssSender }).(pulumi.IntOutput)
}

func (o LookupFirewallPolicy6ResultOutput) TcpSessionWithoutSyn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TcpSessionWithoutSyn }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) TimeoutSendRst() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TimeoutSendRst }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Tos() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Tos }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) TosMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TosMask }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) TosNegate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TosNegate }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) TrafficShaper() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TrafficShaper }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) TrafficShaperReverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.TrafficShaperReverse }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) UrlCategories() GetFirewallPolicy6UrlCategoryArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6UrlCategory { return v.UrlCategories }).(GetFirewallPolicy6UrlCategoryArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Users() GetFirewallPolicy6UserArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) []GetFirewallPolicy6User { return v.Users }).(GetFirewallPolicy6UserArrayOutput)
}

func (o LookupFirewallPolicy6ResultOutput) UtmStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.UtmStatus }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicy6ResultOutput) VlanCosFwd() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.VlanCosFwd }).(pulumi.IntOutput)
}

func (o LookupFirewallPolicy6ResultOutput) VlanCosRev() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) int { return v.VlanCosRev }).(pulumi.IntOutput)
}

func (o LookupFirewallPolicy6ResultOutput) VlanFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.VlanFilter }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) VoipProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.VoipProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Vpntunnel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Vpntunnel }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) WafProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WafProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) Webcache() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.Webcache }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) WebcacheHttps() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WebcacheHttps }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) WebfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WebfilterProfile }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) WebproxyForwardServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WebproxyForwardServer }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicy6ResultOutput) WebproxyProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicy6Result) string { return v.WebproxyProfile }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallPolicy6ResultOutput{})
}
