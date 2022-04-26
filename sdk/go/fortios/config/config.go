// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// CA Bundle file
func GetCabundlefile(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "fortios:cabundlefile")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "FORTIOS_CA_CABUNDLE").(string)
}

// CA certtificate(Optional)
func GetCacert(ctx *pulumi.Context) string {
	return config.Get(ctx, "fortios:cacert")
}

// User certificate
func GetClientcert(ctx *pulumi.Context) string {
	return config.Get(ctx, "fortios:clientcert")
}

// User private key
func GetClientkey(ctx *pulumi.Context) string {
	return config.Get(ctx, "fortios:clientkey")
}

// CA Bundle file
func GetFmgCabundlefile(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "fortios:fmgCabundlefile")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "FORTIOS_FMG_CABUNDLE").(string)
}

// Hostname/IP address of the FortiManager to connect to
func GetFmgHostname(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "fortios:fmgHostname")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "FORTIOS_FMG_HOSTNAME").(string)
}
func GetFmgInsecure(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "fortios:fmgInsecure")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "FORTIOS_FMG_INSECURE").(bool)
}
func GetFmgPasswd(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "fortios:fmgPasswd")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "FORTIOS_FMG_PASSWORD").(string)
}
func GetFmgUsername(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "fortios:fmgUsername")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "FORTIOS_FMG_USERNAME").(string)
}

// The hostname/IP address of the FortiOS to be connected
func GetHostname(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "fortios:hostname")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "FORTIOS_ACCESS_HOSTNAME").(string)
}
func GetInsecure(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "fortios:insecure")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "FORTIOS_INSECURE").(bool)
}

// Enable/disable peer authentication, can be 'enable' or 'disable'
func GetPeerauth(ctx *pulumi.Context) string {
	return config.Get(ctx, "fortios:peerauth")
}
func GetToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "fortios:token")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "FORTIOS_ACCESS_TOKEN").(string)
}
func GetVdom(ctx *pulumi.Context) string {
	return config.Get(ctx, "fortios:vdom")
}
