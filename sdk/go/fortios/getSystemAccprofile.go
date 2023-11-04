// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/lubyou/pulumi-fortios/sdk/go/fortios/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupSystemAccprofile(ctx *pulumi.Context, args *LookupSystemAccprofileArgs, opts ...pulumi.InvokeOption) (*LookupSystemAccprofileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemAccprofileResult
	err := ctx.Invoke("fortios:index/getSystemAccprofile:GetSystemAccprofile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAccprofile.
type LookupSystemAccprofileArgs struct {
	Name      string  `pulumi:"name"`
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAccprofile.
type LookupSystemAccprofileResult struct {
	Admintimeout         int                                  `pulumi:"admintimeout"`
	AdmintimeoutOverride string                               `pulumi:"admintimeoutOverride"`
	Authgrp              string                               `pulumi:"authgrp"`
	Comments             string                               `pulumi:"comments"`
	Ftviewgrp            string                               `pulumi:"ftviewgrp"`
	Fwgrp                string                               `pulumi:"fwgrp"`
	FwgrpPermissions     []GetSystemAccprofileFwgrpPermission `pulumi:"fwgrpPermissions"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                                `pulumi:"id"`
	Loggrp              string                                `pulumi:"loggrp"`
	LoggrpPermissions   []GetSystemAccprofileLoggrpPermission `pulumi:"loggrpPermissions"`
	Name                string                                `pulumi:"name"`
	Netgrp              string                                `pulumi:"netgrp"`
	NetgrpPermissions   []GetSystemAccprofileNetgrpPermission `pulumi:"netgrpPermissions"`
	Scope               string                                `pulumi:"scope"`
	Secfabgrp           string                                `pulumi:"secfabgrp"`
	Sysgrp              string                                `pulumi:"sysgrp"`
	SysgrpPermissions   []GetSystemAccprofileSysgrpPermission `pulumi:"sysgrpPermissions"`
	SystemDiagnostics   string                                `pulumi:"systemDiagnostics"`
	SystemExecuteSsh    string                                `pulumi:"systemExecuteSsh"`
	SystemExecuteTelnet string                                `pulumi:"systemExecuteTelnet"`
	Utmgrp              string                                `pulumi:"utmgrp"`
	UtmgrpPermissions   []GetSystemAccprofileUtmgrpPermission `pulumi:"utmgrpPermissions"`
	Vdomparam           *string                               `pulumi:"vdomparam"`
	Vpngrp              string                                `pulumi:"vpngrp"`
	Wanoptgrp           string                                `pulumi:"wanoptgrp"`
	Wifi                string                                `pulumi:"wifi"`
}

func LookupSystemAccprofileOutput(ctx *pulumi.Context, args LookupSystemAccprofileOutputArgs, opts ...pulumi.InvokeOption) LookupSystemAccprofileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemAccprofileResult, error) {
			args := v.(LookupSystemAccprofileArgs)
			r, err := LookupSystemAccprofile(ctx, &args, opts...)
			var s LookupSystemAccprofileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemAccprofileResultOutput)
}

// A collection of arguments for invoking GetSystemAccprofile.
type LookupSystemAccprofileOutputArgs struct {
	Name      pulumi.StringInput    `pulumi:"name"`
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemAccprofileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAccprofileArgs)(nil)).Elem()
}

// A collection of values returned by GetSystemAccprofile.
type LookupSystemAccprofileResultOutput struct{ *pulumi.OutputState }

func (LookupSystemAccprofileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemAccprofileResult)(nil)).Elem()
}

func (o LookupSystemAccprofileResultOutput) ToLookupSystemAccprofileResultOutput() LookupSystemAccprofileResultOutput {
	return o
}

func (o LookupSystemAccprofileResultOutput) ToLookupSystemAccprofileResultOutputWithContext(ctx context.Context) LookupSystemAccprofileResultOutput {
	return o
}

func (o LookupSystemAccprofileResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSystemAccprofileResult] {
	return pulumix.Output[LookupSystemAccprofileResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupSystemAccprofileResultOutput) Admintimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) int { return v.Admintimeout }).(pulumi.IntOutput)
}

func (o LookupSystemAccprofileResultOutput) AdmintimeoutOverride() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.AdmintimeoutOverride }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Authgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Authgrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Comments }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Ftviewgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Ftviewgrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Fwgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Fwgrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) FwgrpPermissions() GetSystemAccprofileFwgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileFwgrpPermission { return v.FwgrpPermissions }).(GetSystemAccprofileFwgrpPermissionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemAccprofileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Loggrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Loggrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) LoggrpPermissions() GetSystemAccprofileLoggrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileLoggrpPermission { return v.LoggrpPermissions }).(GetSystemAccprofileLoggrpPermissionArrayOutput)
}

func (o LookupSystemAccprofileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Netgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Netgrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) NetgrpPermissions() GetSystemAccprofileNetgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileNetgrpPermission { return v.NetgrpPermissions }).(GetSystemAccprofileNetgrpPermissionArrayOutput)
}

func (o LookupSystemAccprofileResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Secfabgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Secfabgrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Sysgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Sysgrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) SysgrpPermissions() GetSystemAccprofileSysgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileSysgrpPermission { return v.SysgrpPermissions }).(GetSystemAccprofileSysgrpPermissionArrayOutput)
}

func (o LookupSystemAccprofileResultOutput) SystemDiagnostics() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.SystemDiagnostics }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) SystemExecuteSsh() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.SystemExecuteSsh }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) SystemExecuteTelnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.SystemExecuteTelnet }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Utmgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Utmgrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) UtmgrpPermissions() GetSystemAccprofileUtmgrpPermissionArrayOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) []GetSystemAccprofileUtmgrpPermission { return v.UtmgrpPermissions }).(GetSystemAccprofileUtmgrpPermissionArrayOutput)
}

func (o LookupSystemAccprofileResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func (o LookupSystemAccprofileResultOutput) Vpngrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Vpngrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Wanoptgrp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Wanoptgrp }).(pulumi.StringOutput)
}

func (o LookupSystemAccprofileResultOutput) Wifi() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemAccprofileResult) string { return v.Wifi }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemAccprofileResultOutput{})
}
