// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSystemFortiguard(ctx *pulumi.Context, args *LookupSystemFortiguardArgs, opts ...pulumi.InvokeOption) (*LookupSystemFortiguardResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemFortiguardResult
	err := ctx.Invoke("fortios:index/getSystemFortiguard:GetSystemFortiguard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemFortiguard.
type LookupSystemFortiguardArgs struct {
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemFortiguard.
type LookupSystemFortiguardResult struct {
	AntispamCache                string `pulumi:"antispamCache"`
	AntispamCacheMpercent        int    `pulumi:"antispamCacheMpercent"`
	AntispamCacheMpermille       int    `pulumi:"antispamCacheMpermille"`
	AntispamCacheTtl             int    `pulumi:"antispamCacheTtl"`
	AntispamExpiration           int    `pulumi:"antispamExpiration"`
	AntispamForceOff             string `pulumi:"antispamForceOff"`
	AntispamLicense              int    `pulumi:"antispamLicense"`
	AntispamTimeout              int    `pulumi:"antispamTimeout"`
	AnycastSdnsServerIp          string `pulumi:"anycastSdnsServerIp"`
	AnycastSdnsServerPort        int    `pulumi:"anycastSdnsServerPort"`
	AutoFirmwareUpgrade          string `pulumi:"autoFirmwareUpgrade"`
	AutoFirmwareUpgradeDay       string `pulumi:"autoFirmwareUpgradeDay"`
	AutoFirmwareUpgradeDelay     int    `pulumi:"autoFirmwareUpgradeDelay"`
	AutoFirmwareUpgradeEndHour   int    `pulumi:"autoFirmwareUpgradeEndHour"`
	AutoFirmwareUpgradeStartHour int    `pulumi:"autoFirmwareUpgradeStartHour"`
	AutoJoinForticloud           string `pulumi:"autoJoinForticloud"`
	DdnsServerIp                 string `pulumi:"ddnsServerIp"`
	DdnsServerIp6                string `pulumi:"ddnsServerIp6"`
	DdnsServerPort               int    `pulumi:"ddnsServerPort"`
	FdsLicenseExpiringDays       int    `pulumi:"fdsLicenseExpiringDays"`
	FortiguardAnycast            string `pulumi:"fortiguardAnycast"`
	FortiguardAnycastSource      string `pulumi:"fortiguardAnycastSource"`
	GuiPromptAutoUpgrade         string `pulumi:"guiPromptAutoUpgrade"`
	// The provider-assigned unique ID for this managed resource.
	Id                               string  `pulumi:"id"`
	Interface                        string  `pulumi:"interface"`
	InterfaceSelectMethod            string  `pulumi:"interfaceSelectMethod"`
	LoadBalanceServers               int     `pulumi:"loadBalanceServers"`
	OutbreakPreventionCache          string  `pulumi:"outbreakPreventionCache"`
	OutbreakPreventionCacheMpercent  int     `pulumi:"outbreakPreventionCacheMpercent"`
	OutbreakPreventionCacheMpermille int     `pulumi:"outbreakPreventionCacheMpermille"`
	OutbreakPreventionCacheTtl       int     `pulumi:"outbreakPreventionCacheTtl"`
	OutbreakPreventionExpiration     int     `pulumi:"outbreakPreventionExpiration"`
	OutbreakPreventionForceOff       string  `pulumi:"outbreakPreventionForceOff"`
	OutbreakPreventionLicense        int     `pulumi:"outbreakPreventionLicense"`
	OutbreakPreventionTimeout        int     `pulumi:"outbreakPreventionTimeout"`
	PersistentConnection             string  `pulumi:"persistentConnection"`
	Port                             string  `pulumi:"port"`
	Protocol                         string  `pulumi:"protocol"`
	ProxyPassword                    string  `pulumi:"proxyPassword"`
	ProxyServerIp                    string  `pulumi:"proxyServerIp"`
	ProxyServerPort                  int     `pulumi:"proxyServerPort"`
	ProxyUsername                    string  `pulumi:"proxyUsername"`
	SandboxInlineScan                string  `pulumi:"sandboxInlineScan"`
	SandboxRegion                    string  `pulumi:"sandboxRegion"`
	SdnsOptions                      string  `pulumi:"sdnsOptions"`
	SdnsServerIp                     string  `pulumi:"sdnsServerIp"`
	SdnsServerPort                   int     `pulumi:"sdnsServerPort"`
	ServiceAccountId                 string  `pulumi:"serviceAccountId"`
	SourceIp                         string  `pulumi:"sourceIp"`
	SourceIp6                        string  `pulumi:"sourceIp6"`
	UpdateBuildProxy                 string  `pulumi:"updateBuildProxy"`
	UpdateDldb                       string  `pulumi:"updateDldb"`
	UpdateExtdb                      string  `pulumi:"updateExtdb"`
	UpdateFfdb                       string  `pulumi:"updateFfdb"`
	UpdateServerLocation             string  `pulumi:"updateServerLocation"`
	UpdateUwdb                       string  `pulumi:"updateUwdb"`
	Vdom                             string  `pulumi:"vdom"`
	Vdomparam                        *string `pulumi:"vdomparam"`
	VideofilterExpiration            int     `pulumi:"videofilterExpiration"`
	VideofilterLicense               int     `pulumi:"videofilterLicense"`
	WebfilterCache                   string  `pulumi:"webfilterCache"`
	WebfilterCacheTtl                int     `pulumi:"webfilterCacheTtl"`
	WebfilterExpiration              int     `pulumi:"webfilterExpiration"`
	WebfilterForceOff                string  `pulumi:"webfilterForceOff"`
	WebfilterLicense                 int     `pulumi:"webfilterLicense"`
	WebfilterTimeout                 int     `pulumi:"webfilterTimeout"`
}

func LookupSystemFortiguardOutput(ctx *pulumi.Context, args LookupSystemFortiguardOutputArgs, opts ...pulumi.InvokeOption) LookupSystemFortiguardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemFortiguardResult, error) {
			args := v.(LookupSystemFortiguardArgs)
			r, err := LookupSystemFortiguard(ctx, &args, opts...)
			var s LookupSystemFortiguardResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemFortiguardResultOutput)
}

// A collection of arguments for invoking GetSystemFortiguard.
type LookupSystemFortiguardOutputArgs struct {
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemFortiguardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFortiguardArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemFortiguard.
type LookupSystemFortiguardResultOutput struct{ *pulumi.OutputState }

func (LookupSystemFortiguardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemFortiguardResult)(nil)).Elem()
}

func (o LookupSystemFortiguardResultOutput) ToLookupSystemFortiguardResultOutput() LookupSystemFortiguardResultOutput {
	return o
}

func (o LookupSystemFortiguardResultOutput) ToLookupSystemFortiguardResultOutputWithContext(ctx context.Context) LookupSystemFortiguardResultOutput {
	return o
}

func (o LookupSystemFortiguardResultOutput) AntispamCache() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.AntispamCache }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) AntispamCacheMpercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AntispamCacheMpercent }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AntispamCacheMpermille() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AntispamCacheMpermille }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AntispamCacheTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AntispamCacheTtl }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AntispamExpiration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AntispamExpiration }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AntispamForceOff() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.AntispamForceOff }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) AntispamLicense() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AntispamLicense }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AntispamTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AntispamTimeout }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AnycastSdnsServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.AnycastSdnsServerIp }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) AnycastSdnsServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AnycastSdnsServerPort }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AutoFirmwareUpgrade() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.AutoFirmwareUpgrade }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) AutoFirmwareUpgradeDay() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.AutoFirmwareUpgradeDay }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) AutoFirmwareUpgradeDelay() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AutoFirmwareUpgradeDelay }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AutoFirmwareUpgradeEndHour() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AutoFirmwareUpgradeEndHour }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AutoFirmwareUpgradeStartHour() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.AutoFirmwareUpgradeStartHour }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) AutoJoinForticloud() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.AutoJoinForticloud }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) DdnsServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.DdnsServerIp }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) DdnsServerIp6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.DdnsServerIp6 }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) DdnsServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.DdnsServerPort }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) FdsLicenseExpiringDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.FdsLicenseExpiringDays }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) FortiguardAnycast() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.FortiguardAnycast }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) FortiguardAnycastSource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.FortiguardAnycastSource }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) GuiPromptAutoUpgrade() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.GuiPromptAutoUpgrade }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemFortiguardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.Interface }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) LoadBalanceServers() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.LoadBalanceServers }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) OutbreakPreventionCache() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.OutbreakPreventionCache }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) OutbreakPreventionCacheMpercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.OutbreakPreventionCacheMpercent }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) OutbreakPreventionCacheMpermille() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.OutbreakPreventionCacheMpermille }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) OutbreakPreventionCacheTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.OutbreakPreventionCacheTtl }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) OutbreakPreventionExpiration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.OutbreakPreventionExpiration }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) OutbreakPreventionForceOff() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.OutbreakPreventionForceOff }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) OutbreakPreventionLicense() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.OutbreakPreventionLicense }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) OutbreakPreventionTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.OutbreakPreventionTimeout }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) PersistentConnection() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.PersistentConnection }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.Port }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) ProxyPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.ProxyPassword }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) ProxyServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.ProxyServerIp }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) ProxyServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.ProxyServerPort }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) ProxyUsername() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.ProxyUsername }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) SandboxInlineScan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.SandboxInlineScan }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) SandboxRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.SandboxRegion }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) SdnsOptions() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.SdnsOptions }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) SdnsServerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.SdnsServerIp }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) SdnsServerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.SdnsServerPort }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.ServiceAccountId }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) SourceIp6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.SourceIp6 }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) UpdateBuildProxy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.UpdateBuildProxy }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) UpdateDldb() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.UpdateDldb }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) UpdateExtdb() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.UpdateExtdb }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) UpdateFfdb() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.UpdateFfdb }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) UpdateServerLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.UpdateServerLocation }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) UpdateUwdb() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.UpdateUwdb }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) Vdom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.Vdom }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupSystemFortiguardResultOutput) VideofilterExpiration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.VideofilterExpiration }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) VideofilterLicense() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.VideofilterLicense }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) WebfilterCache() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.WebfilterCache }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) WebfilterCacheTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.WebfilterCacheTtl }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) WebfilterExpiration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.WebfilterExpiration }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) WebfilterForceOff() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) string { return v.WebfilterForceOff }).(pulumi.StringOutput)
}

func (o LookupSystemFortiguardResultOutput) WebfilterLicense() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.WebfilterLicense }).(pulumi.IntOutput)
}

func (o LookupSystemFortiguardResultOutput) WebfilterTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemFortiguardResult) int { return v.WebfilterTimeout }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemFortiguardResultOutput{})
}
