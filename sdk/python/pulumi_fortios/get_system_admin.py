# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetSystemAdminResult',
    'AwaitableGetSystemAdminResult',
    'get_system_admin',
    'get_system_admin_output',
]

@pulumi.output_type
class GetSystemAdminResult:
    """
    A collection of values returned by GetSystemAdmin.
    """
    def __init__(__self__, accprofile=None, accprofile_override=None, allow_remove_admin_session=None, comments=None, email_to=None, force_password_change=None, fortitoken=None, guest_auth=None, guest_lang=None, guest_usergroups=None, gui_dashboards=None, gui_global_menu_favorites=None, gui_new_feature_acknowledges=None, gui_vdom_menu_favorites=None, hidden=None, history0=None, history1=None, id=None, ip6_trusthost1=None, ip6_trusthost10=None, ip6_trusthost2=None, ip6_trusthost3=None, ip6_trusthost4=None, ip6_trusthost5=None, ip6_trusthost6=None, ip6_trusthost7=None, ip6_trusthost8=None, ip6_trusthost9=None, login_times=None, name=None, password=None, password_expire=None, peer_auth=None, peer_group=None, radius_vdom_override=None, remote_auth=None, remote_group=None, schedule=None, sms_custom_server=None, sms_phone=None, sms_server=None, ssh_certificate=None, ssh_public_key1=None, ssh_public_key2=None, ssh_public_key3=None, trusthost1=None, trusthost10=None, trusthost2=None, trusthost3=None, trusthost4=None, trusthost5=None, trusthost6=None, trusthost7=None, trusthost8=None, trusthost9=None, two_factor=None, two_factor_authentication=None, two_factor_notification=None, vdomparam=None, vdoms=None, wildcard=None):
        if accprofile and not isinstance(accprofile, str):
            raise TypeError("Expected argument 'accprofile' to be a str")
        pulumi.set(__self__, "accprofile", accprofile)
        if accprofile_override and not isinstance(accprofile_override, str):
            raise TypeError("Expected argument 'accprofile_override' to be a str")
        pulumi.set(__self__, "accprofile_override", accprofile_override)
        if allow_remove_admin_session and not isinstance(allow_remove_admin_session, str):
            raise TypeError("Expected argument 'allow_remove_admin_session' to be a str")
        pulumi.set(__self__, "allow_remove_admin_session", allow_remove_admin_session)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if email_to and not isinstance(email_to, str):
            raise TypeError("Expected argument 'email_to' to be a str")
        pulumi.set(__self__, "email_to", email_to)
        if force_password_change and not isinstance(force_password_change, str):
            raise TypeError("Expected argument 'force_password_change' to be a str")
        pulumi.set(__self__, "force_password_change", force_password_change)
        if fortitoken and not isinstance(fortitoken, str):
            raise TypeError("Expected argument 'fortitoken' to be a str")
        pulumi.set(__self__, "fortitoken", fortitoken)
        if guest_auth and not isinstance(guest_auth, str):
            raise TypeError("Expected argument 'guest_auth' to be a str")
        pulumi.set(__self__, "guest_auth", guest_auth)
        if guest_lang and not isinstance(guest_lang, str):
            raise TypeError("Expected argument 'guest_lang' to be a str")
        pulumi.set(__self__, "guest_lang", guest_lang)
        if guest_usergroups and not isinstance(guest_usergroups, list):
            raise TypeError("Expected argument 'guest_usergroups' to be a list")
        pulumi.set(__self__, "guest_usergroups", guest_usergroups)
        if gui_dashboards and not isinstance(gui_dashboards, list):
            raise TypeError("Expected argument 'gui_dashboards' to be a list")
        pulumi.set(__self__, "gui_dashboards", gui_dashboards)
        if gui_global_menu_favorites and not isinstance(gui_global_menu_favorites, list):
            raise TypeError("Expected argument 'gui_global_menu_favorites' to be a list")
        pulumi.set(__self__, "gui_global_menu_favorites", gui_global_menu_favorites)
        if gui_new_feature_acknowledges and not isinstance(gui_new_feature_acknowledges, list):
            raise TypeError("Expected argument 'gui_new_feature_acknowledges' to be a list")
        pulumi.set(__self__, "gui_new_feature_acknowledges", gui_new_feature_acknowledges)
        if gui_vdom_menu_favorites and not isinstance(gui_vdom_menu_favorites, list):
            raise TypeError("Expected argument 'gui_vdom_menu_favorites' to be a list")
        pulumi.set(__self__, "gui_vdom_menu_favorites", gui_vdom_menu_favorites)
        if hidden and not isinstance(hidden, int):
            raise TypeError("Expected argument 'hidden' to be a int")
        pulumi.set(__self__, "hidden", hidden)
        if history0 and not isinstance(history0, str):
            raise TypeError("Expected argument 'history0' to be a str")
        pulumi.set(__self__, "history0", history0)
        if history1 and not isinstance(history1, str):
            raise TypeError("Expected argument 'history1' to be a str")
        pulumi.set(__self__, "history1", history1)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip6_trusthost1 and not isinstance(ip6_trusthost1, str):
            raise TypeError("Expected argument 'ip6_trusthost1' to be a str")
        pulumi.set(__self__, "ip6_trusthost1", ip6_trusthost1)
        if ip6_trusthost10 and not isinstance(ip6_trusthost10, str):
            raise TypeError("Expected argument 'ip6_trusthost10' to be a str")
        pulumi.set(__self__, "ip6_trusthost10", ip6_trusthost10)
        if ip6_trusthost2 and not isinstance(ip6_trusthost2, str):
            raise TypeError("Expected argument 'ip6_trusthost2' to be a str")
        pulumi.set(__self__, "ip6_trusthost2", ip6_trusthost2)
        if ip6_trusthost3 and not isinstance(ip6_trusthost3, str):
            raise TypeError("Expected argument 'ip6_trusthost3' to be a str")
        pulumi.set(__self__, "ip6_trusthost3", ip6_trusthost3)
        if ip6_trusthost4 and not isinstance(ip6_trusthost4, str):
            raise TypeError("Expected argument 'ip6_trusthost4' to be a str")
        pulumi.set(__self__, "ip6_trusthost4", ip6_trusthost4)
        if ip6_trusthost5 and not isinstance(ip6_trusthost5, str):
            raise TypeError("Expected argument 'ip6_trusthost5' to be a str")
        pulumi.set(__self__, "ip6_trusthost5", ip6_trusthost5)
        if ip6_trusthost6 and not isinstance(ip6_trusthost6, str):
            raise TypeError("Expected argument 'ip6_trusthost6' to be a str")
        pulumi.set(__self__, "ip6_trusthost6", ip6_trusthost6)
        if ip6_trusthost7 and not isinstance(ip6_trusthost7, str):
            raise TypeError("Expected argument 'ip6_trusthost7' to be a str")
        pulumi.set(__self__, "ip6_trusthost7", ip6_trusthost7)
        if ip6_trusthost8 and not isinstance(ip6_trusthost8, str):
            raise TypeError("Expected argument 'ip6_trusthost8' to be a str")
        pulumi.set(__self__, "ip6_trusthost8", ip6_trusthost8)
        if ip6_trusthost9 and not isinstance(ip6_trusthost9, str):
            raise TypeError("Expected argument 'ip6_trusthost9' to be a str")
        pulumi.set(__self__, "ip6_trusthost9", ip6_trusthost9)
        if login_times and not isinstance(login_times, list):
            raise TypeError("Expected argument 'login_times' to be a list")
        pulumi.set(__self__, "login_times", login_times)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if password_expire and not isinstance(password_expire, str):
            raise TypeError("Expected argument 'password_expire' to be a str")
        pulumi.set(__self__, "password_expire", password_expire)
        if peer_auth and not isinstance(peer_auth, str):
            raise TypeError("Expected argument 'peer_auth' to be a str")
        pulumi.set(__self__, "peer_auth", peer_auth)
        if peer_group and not isinstance(peer_group, str):
            raise TypeError("Expected argument 'peer_group' to be a str")
        pulumi.set(__self__, "peer_group", peer_group)
        if radius_vdom_override and not isinstance(radius_vdom_override, str):
            raise TypeError("Expected argument 'radius_vdom_override' to be a str")
        pulumi.set(__self__, "radius_vdom_override", radius_vdom_override)
        if remote_auth and not isinstance(remote_auth, str):
            raise TypeError("Expected argument 'remote_auth' to be a str")
        pulumi.set(__self__, "remote_auth", remote_auth)
        if remote_group and not isinstance(remote_group, str):
            raise TypeError("Expected argument 'remote_group' to be a str")
        pulumi.set(__self__, "remote_group", remote_group)
        if schedule and not isinstance(schedule, str):
            raise TypeError("Expected argument 'schedule' to be a str")
        pulumi.set(__self__, "schedule", schedule)
        if sms_custom_server and not isinstance(sms_custom_server, str):
            raise TypeError("Expected argument 'sms_custom_server' to be a str")
        pulumi.set(__self__, "sms_custom_server", sms_custom_server)
        if sms_phone and not isinstance(sms_phone, str):
            raise TypeError("Expected argument 'sms_phone' to be a str")
        pulumi.set(__self__, "sms_phone", sms_phone)
        if sms_server and not isinstance(sms_server, str):
            raise TypeError("Expected argument 'sms_server' to be a str")
        pulumi.set(__self__, "sms_server", sms_server)
        if ssh_certificate and not isinstance(ssh_certificate, str):
            raise TypeError("Expected argument 'ssh_certificate' to be a str")
        pulumi.set(__self__, "ssh_certificate", ssh_certificate)
        if ssh_public_key1 and not isinstance(ssh_public_key1, str):
            raise TypeError("Expected argument 'ssh_public_key1' to be a str")
        pulumi.set(__self__, "ssh_public_key1", ssh_public_key1)
        if ssh_public_key2 and not isinstance(ssh_public_key2, str):
            raise TypeError("Expected argument 'ssh_public_key2' to be a str")
        pulumi.set(__self__, "ssh_public_key2", ssh_public_key2)
        if ssh_public_key3 and not isinstance(ssh_public_key3, str):
            raise TypeError("Expected argument 'ssh_public_key3' to be a str")
        pulumi.set(__self__, "ssh_public_key3", ssh_public_key3)
        if trusthost1 and not isinstance(trusthost1, str):
            raise TypeError("Expected argument 'trusthost1' to be a str")
        pulumi.set(__self__, "trusthost1", trusthost1)
        if trusthost10 and not isinstance(trusthost10, str):
            raise TypeError("Expected argument 'trusthost10' to be a str")
        pulumi.set(__self__, "trusthost10", trusthost10)
        if trusthost2 and not isinstance(trusthost2, str):
            raise TypeError("Expected argument 'trusthost2' to be a str")
        pulumi.set(__self__, "trusthost2", trusthost2)
        if trusthost3 and not isinstance(trusthost3, str):
            raise TypeError("Expected argument 'trusthost3' to be a str")
        pulumi.set(__self__, "trusthost3", trusthost3)
        if trusthost4 and not isinstance(trusthost4, str):
            raise TypeError("Expected argument 'trusthost4' to be a str")
        pulumi.set(__self__, "trusthost4", trusthost4)
        if trusthost5 and not isinstance(trusthost5, str):
            raise TypeError("Expected argument 'trusthost5' to be a str")
        pulumi.set(__self__, "trusthost5", trusthost5)
        if trusthost6 and not isinstance(trusthost6, str):
            raise TypeError("Expected argument 'trusthost6' to be a str")
        pulumi.set(__self__, "trusthost6", trusthost6)
        if trusthost7 and not isinstance(trusthost7, str):
            raise TypeError("Expected argument 'trusthost7' to be a str")
        pulumi.set(__self__, "trusthost7", trusthost7)
        if trusthost8 and not isinstance(trusthost8, str):
            raise TypeError("Expected argument 'trusthost8' to be a str")
        pulumi.set(__self__, "trusthost8", trusthost8)
        if trusthost9 and not isinstance(trusthost9, str):
            raise TypeError("Expected argument 'trusthost9' to be a str")
        pulumi.set(__self__, "trusthost9", trusthost9)
        if two_factor and not isinstance(two_factor, str):
            raise TypeError("Expected argument 'two_factor' to be a str")
        pulumi.set(__self__, "two_factor", two_factor)
        if two_factor_authentication and not isinstance(two_factor_authentication, str):
            raise TypeError("Expected argument 'two_factor_authentication' to be a str")
        pulumi.set(__self__, "two_factor_authentication", two_factor_authentication)
        if two_factor_notification and not isinstance(two_factor_notification, str):
            raise TypeError("Expected argument 'two_factor_notification' to be a str")
        pulumi.set(__self__, "two_factor_notification", two_factor_notification)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms and not isinstance(vdoms, list):
            raise TypeError("Expected argument 'vdoms' to be a list")
        pulumi.set(__self__, "vdoms", vdoms)
        if wildcard and not isinstance(wildcard, str):
            raise TypeError("Expected argument 'wildcard' to be a str")
        pulumi.set(__self__, "wildcard", wildcard)

    @property
    @pulumi.getter
    def accprofile(self) -> str:
        """
        Access profile for this administrator. Access profiles control administrator access to FortiGate features.
        """
        return pulumi.get(self, "accprofile")

    @property
    @pulumi.getter(name="accprofileOverride")
    def accprofile_override(self) -> str:
        """
        Enable to use the name of an access profile provided by the remote authentication server to control the FortiGate features that this administrator can access.
        """
        return pulumi.get(self, "accprofile_override")

    @property
    @pulumi.getter(name="allowRemoveAdminSession")
    def allow_remove_admin_session(self) -> str:
        """
        Enable/disable allow admin session to be removed by privileged admin users.
        """
        return pulumi.get(self, "allow_remove_admin_session")

    @property
    @pulumi.getter
    def comments(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="emailTo")
    def email_to(self) -> str:
        """
        This administrator's email address.
        """
        return pulumi.get(self, "email_to")

    @property
    @pulumi.getter(name="forcePasswordChange")
    def force_password_change(self) -> str:
        """
        Enable/disable force password change on next login.
        """
        return pulumi.get(self, "force_password_change")

    @property
    @pulumi.getter
    def fortitoken(self) -> str:
        """
        This administrator's FortiToken serial number.
        """
        return pulumi.get(self, "fortitoken")

    @property
    @pulumi.getter(name="guestAuth")
    def guest_auth(self) -> str:
        """
        Enable/disable guest authentication.
        """
        return pulumi.get(self, "guest_auth")

    @property
    @pulumi.getter(name="guestLang")
    def guest_lang(self) -> str:
        """
        Guest management portal language.
        """
        return pulumi.get(self, "guest_lang")

    @property
    @pulumi.getter(name="guestUsergroups")
    def guest_usergroups(self) -> Sequence['outputs.GetSystemAdminGuestUsergroupResult']:
        """
        Select guest user groups. The structure of `guest_usergroups` block is documented below.
        """
        return pulumi.get(self, "guest_usergroups")

    @property
    @pulumi.getter(name="guiDashboards")
    def gui_dashboards(self) -> Sequence['outputs.GetSystemAdminGuiDashboardResult']:
        """
        GUI dashboards. The structure of `gui_dashboard` block is documented below.
        """
        return pulumi.get(self, "gui_dashboards")

    @property
    @pulumi.getter(name="guiGlobalMenuFavorites")
    def gui_global_menu_favorites(self) -> Sequence['outputs.GetSystemAdminGuiGlobalMenuFavoriteResult']:
        """
        Favorite GUI menu IDs for the global VDOM. The structure of `gui_global_menu_favorites` block is documented below.
        """
        return pulumi.get(self, "gui_global_menu_favorites")

    @property
    @pulumi.getter(name="guiNewFeatureAcknowledges")
    def gui_new_feature_acknowledges(self) -> Sequence['outputs.GetSystemAdminGuiNewFeatureAcknowledgeResult']:
        """
        Acknowledgement of new features. The structure of `gui_new_feature_acknowledge` block is documented below.
        """
        return pulumi.get(self, "gui_new_feature_acknowledges")

    @property
    @pulumi.getter(name="guiVdomMenuFavorites")
    def gui_vdom_menu_favorites(self) -> Sequence['outputs.GetSystemAdminGuiVdomMenuFavoriteResult']:
        """
        Favorite GUI menu IDs for VDOMs. The structure of `gui_vdom_menu_favorites` block is documented below.
        """
        return pulumi.get(self, "gui_vdom_menu_favorites")

    @property
    @pulumi.getter
    def hidden(self) -> int:
        """
        Admin user hidden attribute.
        """
        return pulumi.get(self, "hidden")

    @property
    @pulumi.getter
    def history0(self) -> str:
        """
        history0
        """
        return pulumi.get(self, "history0")

    @property
    @pulumi.getter
    def history1(self) -> str:
        """
        history1
        """
        return pulumi.get(self, "history1")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ip6Trusthost1")
    def ip6_trusthost1(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost1")

    @property
    @pulumi.getter(name="ip6Trusthost10")
    def ip6_trusthost10(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost10")

    @property
    @pulumi.getter(name="ip6Trusthost2")
    def ip6_trusthost2(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost2")

    @property
    @pulumi.getter(name="ip6Trusthost3")
    def ip6_trusthost3(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost3")

    @property
    @pulumi.getter(name="ip6Trusthost4")
    def ip6_trusthost4(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost4")

    @property
    @pulumi.getter(name="ip6Trusthost5")
    def ip6_trusthost5(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost5")

    @property
    @pulumi.getter(name="ip6Trusthost6")
    def ip6_trusthost6(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost6")

    @property
    @pulumi.getter(name="ip6Trusthost7")
    def ip6_trusthost7(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost7")

    @property
    @pulumi.getter(name="ip6Trusthost8")
    def ip6_trusthost8(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost8")

    @property
    @pulumi.getter(name="ip6Trusthost9")
    def ip6_trusthost9(self) -> str:
        """
        Any IPv6 address from which the administrator can connect to the FortiGate unit. Default allows access from any IPv6 address.
        """
        return pulumi.get(self, "ip6_trusthost9")

    @property
    @pulumi.getter(name="loginTimes")
    def login_times(self) -> Sequence['outputs.GetSystemAdminLoginTimeResult']:
        """
        Record user login time. The structure of `login_time` block is documented below.
        """
        return pulumi.get(self, "login_times")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Select guest user groups.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        Admin user password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="passwordExpire")
    def password_expire(self) -> str:
        """
        Password expire time.
        """
        return pulumi.get(self, "password_expire")

    @property
    @pulumi.getter(name="peerAuth")
    def peer_auth(self) -> str:
        """
        Set to enable peer certificate authentication (for HTTPS admin access).
        """
        return pulumi.get(self, "peer_auth")

    @property
    @pulumi.getter(name="peerGroup")
    def peer_group(self) -> str:
        """
        Name of peer group defined under config user group which has PKI members. Used for peer certificate authentication (for HTTPS admin access).
        """
        return pulumi.get(self, "peer_group")

    @property
    @pulumi.getter(name="radiusVdomOverride")
    def radius_vdom_override(self) -> str:
        """
        Enable to use the names of VDOMs provided by the remote authentication server to control the VDOMs that this administrator can access.
        """
        return pulumi.get(self, "radius_vdom_override")

    @property
    @pulumi.getter(name="remoteAuth")
    def remote_auth(self) -> str:
        """
        Enable/disable authentication using a remote RADIUS, LDAP, or TACACS+ server.
        """
        return pulumi.get(self, "remote_auth")

    @property
    @pulumi.getter(name="remoteGroup")
    def remote_group(self) -> str:
        """
        User group name used for remote auth.
        """
        return pulumi.get(self, "remote_group")

    @property
    @pulumi.getter
    def schedule(self) -> str:
        """
        Firewall schedule used to restrict when the administrator can log in. No schedule means no restrictions.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="smsCustomServer")
    def sms_custom_server(self) -> str:
        """
        Custom SMS server to send SMS messages to.
        """
        return pulumi.get(self, "sms_custom_server")

    @property
    @pulumi.getter(name="smsPhone")
    def sms_phone(self) -> str:
        """
        Phone number on which the administrator receives SMS messages.
        """
        return pulumi.get(self, "sms_phone")

    @property
    @pulumi.getter(name="smsServer")
    def sms_server(self) -> str:
        """
        Send SMS messages using the FortiGuard SMS server or a custom server.
        """
        return pulumi.get(self, "sms_server")

    @property
    @pulumi.getter(name="sshCertificate")
    def ssh_certificate(self) -> str:
        """
        Select the certificate to be used by the FortiGate for authentication with an SSH client.
        """
        return pulumi.get(self, "ssh_certificate")

    @property
    @pulumi.getter(name="sshPublicKey1")
    def ssh_public_key1(self) -> str:
        """
        Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
        """
        return pulumi.get(self, "ssh_public_key1")

    @property
    @pulumi.getter(name="sshPublicKey2")
    def ssh_public_key2(self) -> str:
        """
        Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
        """
        return pulumi.get(self, "ssh_public_key2")

    @property
    @pulumi.getter(name="sshPublicKey3")
    def ssh_public_key3(self) -> str:
        """
        Public key of an SSH client. The client is authenticated without being asked for credentials. Create the public-private key pair in the SSH client application.
        """
        return pulumi.get(self, "ssh_public_key3")

    @property
    @pulumi.getter
    def trusthost1(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost1")

    @property
    @pulumi.getter
    def trusthost10(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost10")

    @property
    @pulumi.getter
    def trusthost2(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost2")

    @property
    @pulumi.getter
    def trusthost3(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost3")

    @property
    @pulumi.getter
    def trusthost4(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost4")

    @property
    @pulumi.getter
    def trusthost5(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost5")

    @property
    @pulumi.getter
    def trusthost6(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost6")

    @property
    @pulumi.getter
    def trusthost7(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost7")

    @property
    @pulumi.getter
    def trusthost8(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost8")

    @property
    @pulumi.getter
    def trusthost9(self) -> str:
        """
        Any IPv4 address or subnet address and netmask from which the administrator can connect to the FortiGate unit. Default allows access from any IPv4 address.
        """
        return pulumi.get(self, "trusthost9")

    @property
    @pulumi.getter(name="twoFactor")
    def two_factor(self) -> str:
        """
        Enable/disable two-factor authentication.
        """
        return pulumi.get(self, "two_factor")

    @property
    @pulumi.getter(name="twoFactorAuthentication")
    def two_factor_authentication(self) -> str:
        """
        Authentication method by FortiToken Cloud.
        """
        return pulumi.get(self, "two_factor_authentication")

    @property
    @pulumi.getter(name="twoFactorNotification")
    def two_factor_notification(self) -> str:
        """
        Notification method for user activation by FortiToken Cloud.
        """
        return pulumi.get(self, "two_factor_notification")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vdoms(self) -> Sequence['outputs.GetSystemAdminVdomResult']:
        """
        Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
        """
        return pulumi.get(self, "vdoms")

    @property
    @pulumi.getter
    def wildcard(self) -> str:
        """
        Enable/disable wildcard RADIUS authentication.
        """
        return pulumi.get(self, "wildcard")


class AwaitableGetSystemAdminResult(GetSystemAdminResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAdminResult(
            accprofile=self.accprofile,
            accprofile_override=self.accprofile_override,
            allow_remove_admin_session=self.allow_remove_admin_session,
            comments=self.comments,
            email_to=self.email_to,
            force_password_change=self.force_password_change,
            fortitoken=self.fortitoken,
            guest_auth=self.guest_auth,
            guest_lang=self.guest_lang,
            guest_usergroups=self.guest_usergroups,
            gui_dashboards=self.gui_dashboards,
            gui_global_menu_favorites=self.gui_global_menu_favorites,
            gui_new_feature_acknowledges=self.gui_new_feature_acknowledges,
            gui_vdom_menu_favorites=self.gui_vdom_menu_favorites,
            hidden=self.hidden,
            history0=self.history0,
            history1=self.history1,
            id=self.id,
            ip6_trusthost1=self.ip6_trusthost1,
            ip6_trusthost10=self.ip6_trusthost10,
            ip6_trusthost2=self.ip6_trusthost2,
            ip6_trusthost3=self.ip6_trusthost3,
            ip6_trusthost4=self.ip6_trusthost4,
            ip6_trusthost5=self.ip6_trusthost5,
            ip6_trusthost6=self.ip6_trusthost6,
            ip6_trusthost7=self.ip6_trusthost7,
            ip6_trusthost8=self.ip6_trusthost8,
            ip6_trusthost9=self.ip6_trusthost9,
            login_times=self.login_times,
            name=self.name,
            password=self.password,
            password_expire=self.password_expire,
            peer_auth=self.peer_auth,
            peer_group=self.peer_group,
            radius_vdom_override=self.radius_vdom_override,
            remote_auth=self.remote_auth,
            remote_group=self.remote_group,
            schedule=self.schedule,
            sms_custom_server=self.sms_custom_server,
            sms_phone=self.sms_phone,
            sms_server=self.sms_server,
            ssh_certificate=self.ssh_certificate,
            ssh_public_key1=self.ssh_public_key1,
            ssh_public_key2=self.ssh_public_key2,
            ssh_public_key3=self.ssh_public_key3,
            trusthost1=self.trusthost1,
            trusthost10=self.trusthost10,
            trusthost2=self.trusthost2,
            trusthost3=self.trusthost3,
            trusthost4=self.trusthost4,
            trusthost5=self.trusthost5,
            trusthost6=self.trusthost6,
            trusthost7=self.trusthost7,
            trusthost8=self.trusthost8,
            trusthost9=self.trusthost9,
            two_factor=self.two_factor,
            two_factor_authentication=self.two_factor_authentication,
            two_factor_notification=self.two_factor_notification,
            vdomparam=self.vdomparam,
            vdoms=self.vdoms,
            wildcard=self.wildcard)


def get_system_admin(name: Optional[str] = None,
                     vdomparam: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAdminResult:
    """
    Use this data source to get information on an fortios system admin


    :param str name: Specify the name of the desired system admin.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemAdmin:GetSystemAdmin', __args__, opts=opts, typ=GetSystemAdminResult).value

    return AwaitableGetSystemAdminResult(
        accprofile=__ret__.accprofile,
        accprofile_override=__ret__.accprofile_override,
        allow_remove_admin_session=__ret__.allow_remove_admin_session,
        comments=__ret__.comments,
        email_to=__ret__.email_to,
        force_password_change=__ret__.force_password_change,
        fortitoken=__ret__.fortitoken,
        guest_auth=__ret__.guest_auth,
        guest_lang=__ret__.guest_lang,
        guest_usergroups=__ret__.guest_usergroups,
        gui_dashboards=__ret__.gui_dashboards,
        gui_global_menu_favorites=__ret__.gui_global_menu_favorites,
        gui_new_feature_acknowledges=__ret__.gui_new_feature_acknowledges,
        gui_vdom_menu_favorites=__ret__.gui_vdom_menu_favorites,
        hidden=__ret__.hidden,
        history0=__ret__.history0,
        history1=__ret__.history1,
        id=__ret__.id,
        ip6_trusthost1=__ret__.ip6_trusthost1,
        ip6_trusthost10=__ret__.ip6_trusthost10,
        ip6_trusthost2=__ret__.ip6_trusthost2,
        ip6_trusthost3=__ret__.ip6_trusthost3,
        ip6_trusthost4=__ret__.ip6_trusthost4,
        ip6_trusthost5=__ret__.ip6_trusthost5,
        ip6_trusthost6=__ret__.ip6_trusthost6,
        ip6_trusthost7=__ret__.ip6_trusthost7,
        ip6_trusthost8=__ret__.ip6_trusthost8,
        ip6_trusthost9=__ret__.ip6_trusthost9,
        login_times=__ret__.login_times,
        name=__ret__.name,
        password=__ret__.password,
        password_expire=__ret__.password_expire,
        peer_auth=__ret__.peer_auth,
        peer_group=__ret__.peer_group,
        radius_vdom_override=__ret__.radius_vdom_override,
        remote_auth=__ret__.remote_auth,
        remote_group=__ret__.remote_group,
        schedule=__ret__.schedule,
        sms_custom_server=__ret__.sms_custom_server,
        sms_phone=__ret__.sms_phone,
        sms_server=__ret__.sms_server,
        ssh_certificate=__ret__.ssh_certificate,
        ssh_public_key1=__ret__.ssh_public_key1,
        ssh_public_key2=__ret__.ssh_public_key2,
        ssh_public_key3=__ret__.ssh_public_key3,
        trusthost1=__ret__.trusthost1,
        trusthost10=__ret__.trusthost10,
        trusthost2=__ret__.trusthost2,
        trusthost3=__ret__.trusthost3,
        trusthost4=__ret__.trusthost4,
        trusthost5=__ret__.trusthost5,
        trusthost6=__ret__.trusthost6,
        trusthost7=__ret__.trusthost7,
        trusthost8=__ret__.trusthost8,
        trusthost9=__ret__.trusthost9,
        two_factor=__ret__.two_factor,
        two_factor_authentication=__ret__.two_factor_authentication,
        two_factor_notification=__ret__.two_factor_notification,
        vdomparam=__ret__.vdomparam,
        vdoms=__ret__.vdoms,
        wildcard=__ret__.wildcard)


@_utilities.lift_output_func(get_system_admin)
def get_system_admin_output(name: Optional[pulumi.Input[str]] = None,
                            vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemAdminResult]:
    """
    Use this data source to get information on an fortios system admin


    :param str name: Specify the name of the desired system admin.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
