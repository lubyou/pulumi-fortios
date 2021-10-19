// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system accprofile
func LookupSystemAccprofile(ctx *pulumi.Context, args *LookupSystemAccprofileArgs, opts ...pulumi.InvokeOption) (*LookupSystemAccprofileResult, error) {
	var rv LookupSystemAccprofileResult
	err := ctx.Invoke("fortios:index/getSystemAccprofile:GetSystemAccprofile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAccprofile.
type LookupSystemAccprofileArgs struct {
	// Specify the name of the desired system accprofile.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAccprofile.
type LookupSystemAccprofileResult struct {
	// Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
	Admintimeout int `pulumi:"admintimeout"`
	// Enable/disable overriding the global administrator idle timeout.
	AdmintimeoutOverride string `pulumi:"admintimeoutOverride"`
	// Administrator access to Users and Devices.
	Authgrp string `pulumi:"authgrp"`
	// Comment.
	Comments string `pulumi:"comments"`
	// FortiView.
	Ftviewgrp string `pulumi:"ftviewgrp"`
	// Administrator access to the Firewall configuration.
	Fwgrp string `pulumi:"fwgrp"`
	// Custom firewall permission. The structure of `fwgrpPermission` block is documented below.
	FwgrpPermission GetSystemAccprofileFwgrpPermission `pulumi:"fwgrpPermission"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Administrator access to Logging and Reporting including viewing log messages.
	Loggrp string `pulumi:"loggrp"`
	// Custom Log & Report permission. The structure of `loggrpPermission` block is documented below.
	LoggrpPermission GetSystemAccprofileLoggrpPermission `pulumi:"loggrpPermission"`
	// Profile name.
	Name string `pulumi:"name"`
	// Network Configuration.
	Netgrp string `pulumi:"netgrp"`
	// Custom network permission. The structure of `netgrpPermission` block is documented below.
	NetgrpPermission GetSystemAccprofileNetgrpPermission `pulumi:"netgrpPermission"`
	// Scope of admin access: global or specific VDOM(s).
	Scope string `pulumi:"scope"`
	// Security Fabric.
	Secfabgrp string `pulumi:"secfabgrp"`
	// System Configuration.
	Sysgrp string `pulumi:"sysgrp"`
	// Custom system permission. The structure of `sysgrpPermission` block is documented below.
	SysgrpPermission GetSystemAccprofileSysgrpPermission `pulumi:"sysgrpPermission"`
	// Enable/disable permission to run system diagnostic commands.
	SystemDiagnostics string `pulumi:"systemDiagnostics"`
	// Administrator access to Security Profiles.
	Utmgrp string `pulumi:"utmgrp"`
	// Custom Security Profile permissions. The structure of `utmgrpPermission` block is documented below.
	UtmgrpPermission GetSystemAccprofileUtmgrpPermission `pulumi:"utmgrpPermission"`
	Vdomparam        *string                             `pulumi:"vdomparam"`
	// Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
	Vpngrp string `pulumi:"vpngrp"`
	// Administrator access to WAN Opt & Cache.
	Wanoptgrp string `pulumi:"wanoptgrp"`
	// Administrator access to the WiFi controller and Switch controller.
	Wifi string `pulumi:"wifi"`
}
