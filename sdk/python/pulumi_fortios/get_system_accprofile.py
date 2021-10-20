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
    'GetSystemAccprofileResult',
    'AwaitableGetSystemAccprofileResult',
    'get_system_accprofile',
]

@pulumi.output_type
class GetSystemAccprofileResult:
    """
    A collection of values returned by GetSystemAccprofile.
    """
    def __init__(__self__, admintimeout=None, admintimeout_override=None, authgrp=None, comments=None, ftviewgrp=None, fwgrp=None, fwgrp_permission=None, id=None, loggrp=None, loggrp_permission=None, name=None, netgrp=None, netgrp_permission=None, scope=None, secfabgrp=None, sysgrp=None, sysgrp_permission=None, system_diagnostics=None, utmgrp=None, utmgrp_permission=None, vdomparam=None, vpngrp=None, wanoptgrp=None, wifi=None):
        if admintimeout and not isinstance(admintimeout, int):
            raise TypeError("Expected argument 'admintimeout' to be a int")
        pulumi.set(__self__, "admintimeout", admintimeout)
        if admintimeout_override and not isinstance(admintimeout_override, str):
            raise TypeError("Expected argument 'admintimeout_override' to be a str")
        pulumi.set(__self__, "admintimeout_override", admintimeout_override)
        if authgrp and not isinstance(authgrp, str):
            raise TypeError("Expected argument 'authgrp' to be a str")
        pulumi.set(__self__, "authgrp", authgrp)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if ftviewgrp and not isinstance(ftviewgrp, str):
            raise TypeError("Expected argument 'ftviewgrp' to be a str")
        pulumi.set(__self__, "ftviewgrp", ftviewgrp)
        if fwgrp and not isinstance(fwgrp, str):
            raise TypeError("Expected argument 'fwgrp' to be a str")
        pulumi.set(__self__, "fwgrp", fwgrp)
        if fwgrp_permission and not isinstance(fwgrp_permission, dict):
            raise TypeError("Expected argument 'fwgrp_permission' to be a dict")
        pulumi.set(__self__, "fwgrp_permission", fwgrp_permission)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if loggrp and not isinstance(loggrp, str):
            raise TypeError("Expected argument 'loggrp' to be a str")
        pulumi.set(__self__, "loggrp", loggrp)
        if loggrp_permission and not isinstance(loggrp_permission, dict):
            raise TypeError("Expected argument 'loggrp_permission' to be a dict")
        pulumi.set(__self__, "loggrp_permission", loggrp_permission)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if netgrp and not isinstance(netgrp, str):
            raise TypeError("Expected argument 'netgrp' to be a str")
        pulumi.set(__self__, "netgrp", netgrp)
        if netgrp_permission and not isinstance(netgrp_permission, dict):
            raise TypeError("Expected argument 'netgrp_permission' to be a dict")
        pulumi.set(__self__, "netgrp_permission", netgrp_permission)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)
        if secfabgrp and not isinstance(secfabgrp, str):
            raise TypeError("Expected argument 'secfabgrp' to be a str")
        pulumi.set(__self__, "secfabgrp", secfabgrp)
        if sysgrp and not isinstance(sysgrp, str):
            raise TypeError("Expected argument 'sysgrp' to be a str")
        pulumi.set(__self__, "sysgrp", sysgrp)
        if sysgrp_permission and not isinstance(sysgrp_permission, dict):
            raise TypeError("Expected argument 'sysgrp_permission' to be a dict")
        pulumi.set(__self__, "sysgrp_permission", sysgrp_permission)
        if system_diagnostics and not isinstance(system_diagnostics, str):
            raise TypeError("Expected argument 'system_diagnostics' to be a str")
        pulumi.set(__self__, "system_diagnostics", system_diagnostics)
        if utmgrp and not isinstance(utmgrp, str):
            raise TypeError("Expected argument 'utmgrp' to be a str")
        pulumi.set(__self__, "utmgrp", utmgrp)
        if utmgrp_permission and not isinstance(utmgrp_permission, dict):
            raise TypeError("Expected argument 'utmgrp_permission' to be a dict")
        pulumi.set(__self__, "utmgrp_permission", utmgrp_permission)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if vpngrp and not isinstance(vpngrp, str):
            raise TypeError("Expected argument 'vpngrp' to be a str")
        pulumi.set(__self__, "vpngrp", vpngrp)
        if wanoptgrp and not isinstance(wanoptgrp, str):
            raise TypeError("Expected argument 'wanoptgrp' to be a str")
        pulumi.set(__self__, "wanoptgrp", wanoptgrp)
        if wifi and not isinstance(wifi, str):
            raise TypeError("Expected argument 'wifi' to be a str")
        pulumi.set(__self__, "wifi", wifi)

    @property
    @pulumi.getter
    def admintimeout(self) -> int:
        """
        Administrator timeout for this access profile (0 - 480 min, default = 10, 0 means never timeout).
        """
        return pulumi.get(self, "admintimeout")

    @property
    @pulumi.getter(name="admintimeoutOverride")
    def admintimeout_override(self) -> str:
        """
        Enable/disable overriding the global administrator idle timeout.
        """
        return pulumi.get(self, "admintimeout_override")

    @property
    @pulumi.getter
    def authgrp(self) -> str:
        """
        Administrator access to Users and Devices.
        """
        return pulumi.get(self, "authgrp")

    @property
    @pulumi.getter
    def comments(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def ftviewgrp(self) -> str:
        """
        FortiView.
        """
        return pulumi.get(self, "ftviewgrp")

    @property
    @pulumi.getter
    def fwgrp(self) -> str:
        """
        Administrator access to the Firewall configuration.
        """
        return pulumi.get(self, "fwgrp")

    @property
    @pulumi.getter(name="fwgrpPermission")
    def fwgrp_permission(self) -> 'outputs.GetSystemAccprofileFwgrpPermissionResult':
        """
        Custom firewall permission. The structure of `fwgrp_permission` block is documented below.
        """
        return pulumi.get(self, "fwgrp_permission")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def loggrp(self) -> str:
        """
        Administrator access to Logging and Reporting including viewing log messages.
        """
        return pulumi.get(self, "loggrp")

    @property
    @pulumi.getter(name="loggrpPermission")
    def loggrp_permission(self) -> 'outputs.GetSystemAccprofileLoggrpPermissionResult':
        """
        Custom Log & Report permission. The structure of `loggrp_permission` block is documented below.
        """
        return pulumi.get(self, "loggrp_permission")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netgrp(self) -> str:
        """
        Network Configuration.
        """
        return pulumi.get(self, "netgrp")

    @property
    @pulumi.getter(name="netgrpPermission")
    def netgrp_permission(self) -> 'outputs.GetSystemAccprofileNetgrpPermissionResult':
        """
        Custom network permission. The structure of `netgrp_permission` block is documented below.
        """
        return pulumi.get(self, "netgrp_permission")

    @property
    @pulumi.getter
    def scope(self) -> str:
        """
        Scope of admin access: global or specific VDOM(s).
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def secfabgrp(self) -> str:
        """
        Security Fabric.
        """
        return pulumi.get(self, "secfabgrp")

    @property
    @pulumi.getter
    def sysgrp(self) -> str:
        """
        System Configuration.
        """
        return pulumi.get(self, "sysgrp")

    @property
    @pulumi.getter(name="sysgrpPermission")
    def sysgrp_permission(self) -> 'outputs.GetSystemAccprofileSysgrpPermissionResult':
        """
        Custom system permission. The structure of `sysgrp_permission` block is documented below.
        """
        return pulumi.get(self, "sysgrp_permission")

    @property
    @pulumi.getter(name="systemDiagnostics")
    def system_diagnostics(self) -> str:
        """
        Enable/disable permission to run system diagnostic commands.
        """
        return pulumi.get(self, "system_diagnostics")

    @property
    @pulumi.getter
    def utmgrp(self) -> str:
        """
        Administrator access to Security Profiles.
        """
        return pulumi.get(self, "utmgrp")

    @property
    @pulumi.getter(name="utmgrpPermission")
    def utmgrp_permission(self) -> 'outputs.GetSystemAccprofileUtmgrpPermissionResult':
        """
        Custom Security Profile permissions. The structure of `utmgrp_permission` block is documented below.
        """
        return pulumi.get(self, "utmgrp_permission")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vpngrp(self) -> str:
        """
        Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        """
        return pulumi.get(self, "vpngrp")

    @property
    @pulumi.getter
    def wanoptgrp(self) -> str:
        """
        Administrator access to WAN Opt & Cache.
        """
        return pulumi.get(self, "wanoptgrp")

    @property
    @pulumi.getter
    def wifi(self) -> str:
        """
        Administrator access to the WiFi controller and Switch controller.
        """
        return pulumi.get(self, "wifi")


class AwaitableGetSystemAccprofileResult(GetSystemAccprofileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAccprofileResult(
            admintimeout=self.admintimeout,
            admintimeout_override=self.admintimeout_override,
            authgrp=self.authgrp,
            comments=self.comments,
            ftviewgrp=self.ftviewgrp,
            fwgrp=self.fwgrp,
            fwgrp_permission=self.fwgrp_permission,
            id=self.id,
            loggrp=self.loggrp,
            loggrp_permission=self.loggrp_permission,
            name=self.name,
            netgrp=self.netgrp,
            netgrp_permission=self.netgrp_permission,
            scope=self.scope,
            secfabgrp=self.secfabgrp,
            sysgrp=self.sysgrp,
            sysgrp_permission=self.sysgrp_permission,
            system_diagnostics=self.system_diagnostics,
            utmgrp=self.utmgrp,
            utmgrp_permission=self.utmgrp_permission,
            vdomparam=self.vdomparam,
            vpngrp=self.vpngrp,
            wanoptgrp=self.wanoptgrp,
            wifi=self.wifi)


def get_system_accprofile(name: Optional[str] = None,
                          vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAccprofileResult:
    """
    Use this data source to get information on an fortios system accprofile


    :param str name: Specify the name of the desired system accprofile.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemAccprofile:GetSystemAccprofile', __args__, opts=opts, typ=GetSystemAccprofileResult).value

    return AwaitableGetSystemAccprofileResult(
        admintimeout=__ret__.admintimeout,
        admintimeout_override=__ret__.admintimeout_override,
        authgrp=__ret__.authgrp,
        comments=__ret__.comments,
        ftviewgrp=__ret__.ftviewgrp,
        fwgrp=__ret__.fwgrp,
        fwgrp_permission=__ret__.fwgrp_permission,
        id=__ret__.id,
        loggrp=__ret__.loggrp,
        loggrp_permission=__ret__.loggrp_permission,
        name=__ret__.name,
        netgrp=__ret__.netgrp,
        netgrp_permission=__ret__.netgrp_permission,
        scope=__ret__.scope,
        secfabgrp=__ret__.secfabgrp,
        sysgrp=__ret__.sysgrp,
        sysgrp_permission=__ret__.sysgrp_permission,
        system_diagnostics=__ret__.system_diagnostics,
        utmgrp=__ret__.utmgrp,
        utmgrp_permission=__ret__.utmgrp_permission,
        vdomparam=__ret__.vdomparam,
        vpngrp=__ret__.vpngrp,
        wanoptgrp=__ret__.wanoptgrp,
        wifi=__ret__.wifi)