// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system admin
func LookupSystemAdmin(ctx *pulumi.Context, args *LookupSystemAdminArgs, opts ...pulumi.InvokeOption) (*LookupSystemAdminResult, error) {
	var rv LookupSystemAdminResult
	err := ctx.Invoke("fortios:index/getSystemAdmin:GetSystemAdmin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSystemAdmin.
type LookupSystemAdminArgs struct {
	// Specify the name of the desired system admin.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by GetSystemAdmin.
type LookupSystemAdminResult struct {
	// Access profile for this administrator. Access profiles control administrator access to FortiGate features.
	Accprofile string `pulumi:"accprofile"`
	// Enable to use the name of an access profile provided by the remote authentication server to control the FortiGate features that this administrator can access.
	AccprofileOverride string `pulumi:"accprofileOverride"`
	// Enable/disable allow admin session to be removed by privileged admin users.
	AllowRemoveAdminSession string `pulumi:"allowRemoveAdminSession"`
	// Comment.
	Comments string `pulumi:"comments"`
	// This administrator's email address.
	EmailTo string `pulumi:"emailTo"`
	// Enable/disable force password change on next login.
	ForcePasswordChange string `pulumi:"forcePasswordChange"`
	// This administrator's FortiToken serial number.
	Fortitoken string `pulumi:"fortitoken"`
	// Enable/disable guest authentication.
	GuestAuth string `pulumi:"guestAuth"`
	// Guest management portal language.
	GuestLang string `pulumi:"guestLang"`
	// Select guest user groups. The structure of `guestUsergroups` block is documented below.
	GuestUsergroups []GetSystemAdminGuestUsergroup `pulumi:"guestUsergroups"`
	// GUI dashboards. The structure of `guiDashboard` block is documented below.
	GuiDashboards []GetSystemAdminGuiDashboard `pulumi:"guiDashboards"`
	// Favorite GUI menu IDs for the global VDOM. The structure of `guiGlobalMenuFavorites` block is documented below.
	GuiGlobalMenuFavorites []GetSystemAdminGuiGlobalMenuFavorite `pulumi:"guiGlobalMenuFavorites"`
	// Acknowledgement of new features. The structure of `guiNewFeatureAcknowledge` block is documented below.
	GuiNewFeatureAcknowledges []GetSystemAdminGuiNewFeatureAcknowledge `pulumi:"guiNewFeatureAcknowledges"`
	// Favorite GUI menu IDs for VDOMs. The structure of `guiVdomMenuFavorites` block is documented below.
	GuiVdomMenuFavorites []GetSystemAdminGuiVdomMenuFavorite `pulumi:"guiVdomMenuFavorites"`
	// Admin user hidden attribute.
	Hidden int `pulumi:"hidden"`
	// history0
	History0 string `pulumi:"history0"`
	// history1
	History1 string `pulumi:"history1"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost1 string `pulumi:"ip6Trusthost1"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost10 string `pulumi:"ip6Trusthost10"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost2 string `pulumi:"ip6Trusthost2"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost3 string `pulumi:"ip6Trusthost3"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost4 string `pulumi:"ip6Trusthost4"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost5 string `pulumi:"ip6Trusthost5"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost6 string `pulumi:"ip6Trusthost6"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost7 string `pulumi:"ip6Trusthost7"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost8 string `pulumi:"ip6Trusthost8"`
	// Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
	Ip6Trusthost9 string `pulumi:"ip6Trusthost9"`
	// Record user login time. The structure of `loginTime` block is documented below.
	LoginTimes []GetSystemAdminLoginTime `pulumi:"loginTimes"`
	// Select guest user groups.
	Name string `pulumi:"name"`
	// Admin user password.
	Password string `pulumi:"password"`
	// Password expire time.
	PasswordExpire string `pulumi:"passwordExpire"`
	// Set to enable peer certificate authentication (for HTTPS admin access).
	PeerAuth string `pulumi:"peerAuth"`
	// Name of peer group defined under config user group which has PKI members. Used for peer certificate authentication (for HTTPS admin access).
	PeerGroup string `pulumi:"peerGroup"`
	// Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access.
	RadiusVdomOverride string `pulumi:"radiusVdomOverride"`
	// Enable/disable authentication using a remote RADIUS, LDAP, or TACACS+ server.
	RemoteAuth string `pulumi:"remoteAuth"`
	// User group name used for remote auth.
	RemoteGroup string `pulumi:"remoteGroup"`
	// Firewall schedule used to restrict when the administrator can log in. No schedule means no restrictions.
	Schedule string `pulumi:"schedule"`
	// Custom SMS server to send SMS messages to.
	SmsCustomServer string `pulumi:"smsCustomServer"`
	// Phone number on which the administrator receives SMS messages.
	SmsPhone string `pulumi:"smsPhone"`
	// Send SMS messages using the FortiGuard SMS server or a custom server.
	SmsServer string `pulumi:"smsServer"`
	// Select the certificate to be used by the FortiGate for authentication with an SSH client.
	SshCertificate string `pulumi:"sshCertificate"`
	// Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
	SshPublicKey1 string `pulumi:"sshPublicKey1"`
	// Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
	SshPublicKey2 string `pulumi:"sshPublicKey2"`
	// Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
	SshPublicKey3 string `pulumi:"sshPublicKey3"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost1 string `pulumi:"trusthost1"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost10 string `pulumi:"trusthost10"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost2 string `pulumi:"trusthost2"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost3 string `pulumi:"trusthost3"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost4 string `pulumi:"trusthost4"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost5 string `pulumi:"trusthost5"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost6 string `pulumi:"trusthost6"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost7 string `pulumi:"trusthost7"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost8 string `pulumi:"trusthost8"`
	// Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
	Trusthost9 string `pulumi:"trusthost9"`
	// Enable/disable two-factor authentication.
	TwoFactor string `pulumi:"twoFactor"`
	// Authentication method by FortiToken Cloud.
	TwoFactorAuthentication string `pulumi:"twoFactorAuthentication"`
	// Notification method for user activation by FortiToken Cloud.
	TwoFactorNotification string  `pulumi:"twoFactorNotification"`
	Vdomparam             *string `pulumi:"vdomparam"`
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms []GetSystemAdminVdom `pulumi:"vdoms"`
	// Enable/disable wildcard RADIUS authentication.
	Wildcard string `pulumi:"wildcard"`
}
