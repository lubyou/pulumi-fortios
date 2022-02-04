# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemAdminProfilesArgs', 'SystemAdminProfiles']

@pulumi.input_type
class SystemAdminProfilesArgs:
    def __init__(__self__, *,
                 admintimeout_override: Optional[pulumi.Input[str]] = None,
                 authgrp: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftviewgrp: Optional[pulumi.Input[str]] = None,
                 fwgrp: Optional[pulumi.Input[str]] = None,
                 loggrp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netgrp: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 secfabgrp: Optional[pulumi.Input[str]] = None,
                 sysgrp: Optional[pulumi.Input[str]] = None,
                 utmgrp: Optional[pulumi.Input[str]] = None,
                 vpngrp: Optional[pulumi.Input[str]] = None,
                 wanoptgrp: Optional[pulumi.Input[str]] = None,
                 wifi: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemAdminProfiles resource.
        :param pulumi.Input[str] admintimeout_override: Enable/disable overriding the global administrator idle timeout.
        :param pulumi.Input[str] authgrp: Administrator access to Users and Devices.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] ftviewgrp: FortiView.
        :param pulumi.Input[str] fwgrp: Administrator access to the Firewall configuration.
        :param pulumi.Input[str] loggrp: Administrator access to Logging and Reporting including viewing log messages.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[str] netgrp: Network Configuration.
        :param pulumi.Input[str] scope: Scope of admin access.
        :param pulumi.Input[str] secfabgrp: Security Fabric.
        :param pulumi.Input[str] sysgrp: System Configuration.
        :param pulumi.Input[str] utmgrp: Administrator access to Security Profiles.
        :param pulumi.Input[str] vpngrp: Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        :param pulumi.Input[str] wanoptgrp: Administrator access to WAN Opt & Cache.
        :param pulumi.Input[str] wifi: Administrator access to the WiFi controller and Switch controller.
        """
        if admintimeout_override is not None:
            pulumi.set(__self__, "admintimeout_override", admintimeout_override)
        if authgrp is not None:
            pulumi.set(__self__, "authgrp", authgrp)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if ftviewgrp is not None:
            pulumi.set(__self__, "ftviewgrp", ftviewgrp)
        if fwgrp is not None:
            pulumi.set(__self__, "fwgrp", fwgrp)
        if loggrp is not None:
            pulumi.set(__self__, "loggrp", loggrp)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if netgrp is not None:
            pulumi.set(__self__, "netgrp", netgrp)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if secfabgrp is not None:
            pulumi.set(__self__, "secfabgrp", secfabgrp)
        if sysgrp is not None:
            pulumi.set(__self__, "sysgrp", sysgrp)
        if utmgrp is not None:
            pulumi.set(__self__, "utmgrp", utmgrp)
        if vpngrp is not None:
            pulumi.set(__self__, "vpngrp", vpngrp)
        if wanoptgrp is not None:
            pulumi.set(__self__, "wanoptgrp", wanoptgrp)
        if wifi is not None:
            pulumi.set(__self__, "wifi", wifi)

    @property
    @pulumi.getter(name="admintimeoutOverride")
    def admintimeout_override(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable overriding the global administrator idle timeout.
        """
        return pulumi.get(self, "admintimeout_override")

    @admintimeout_override.setter
    def admintimeout_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admintimeout_override", value)

    @property
    @pulumi.getter
    def authgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to Users and Devices.
        """
        return pulumi.get(self, "authgrp")

    @authgrp.setter
    def authgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authgrp", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def ftviewgrp(self) -> Optional[pulumi.Input[str]]:
        """
        FortiView.
        """
        return pulumi.get(self, "ftviewgrp")

    @ftviewgrp.setter
    def ftviewgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ftviewgrp", value)

    @property
    @pulumi.getter
    def fwgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to the Firewall configuration.
        """
        return pulumi.get(self, "fwgrp")

    @fwgrp.setter
    def fwgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fwgrp", value)

    @property
    @pulumi.getter
    def loggrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to Logging and Reporting including viewing log messages.
        """
        return pulumi.get(self, "loggrp")

    @loggrp.setter
    def loggrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "loggrp", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def netgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Network Configuration.
        """
        return pulumi.get(self, "netgrp")

    @netgrp.setter
    def netgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netgrp", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Scope of admin access.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def secfabgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Security Fabric.
        """
        return pulumi.get(self, "secfabgrp")

    @secfabgrp.setter
    def secfabgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secfabgrp", value)

    @property
    @pulumi.getter
    def sysgrp(self) -> Optional[pulumi.Input[str]]:
        """
        System Configuration.
        """
        return pulumi.get(self, "sysgrp")

    @sysgrp.setter
    def sysgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sysgrp", value)

    @property
    @pulumi.getter
    def utmgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to Security Profiles.
        """
        return pulumi.get(self, "utmgrp")

    @utmgrp.setter
    def utmgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "utmgrp", value)

    @property
    @pulumi.getter
    def vpngrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        """
        return pulumi.get(self, "vpngrp")

    @vpngrp.setter
    def vpngrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpngrp", value)

    @property
    @pulumi.getter
    def wanoptgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to WAN Opt & Cache.
        """
        return pulumi.get(self, "wanoptgrp")

    @wanoptgrp.setter
    def wanoptgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wanoptgrp", value)

    @property
    @pulumi.getter
    def wifi(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to the WiFi controller and Switch controller.
        """
        return pulumi.get(self, "wifi")

    @wifi.setter
    def wifi(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wifi", value)


@pulumi.input_type
class _SystemAdminProfilesState:
    def __init__(__self__, *,
                 admintimeout_override: Optional[pulumi.Input[str]] = None,
                 authgrp: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftviewgrp: Optional[pulumi.Input[str]] = None,
                 fwgrp: Optional[pulumi.Input[str]] = None,
                 loggrp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netgrp: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 secfabgrp: Optional[pulumi.Input[str]] = None,
                 sysgrp: Optional[pulumi.Input[str]] = None,
                 utmgrp: Optional[pulumi.Input[str]] = None,
                 vpngrp: Optional[pulumi.Input[str]] = None,
                 wanoptgrp: Optional[pulumi.Input[str]] = None,
                 wifi: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemAdminProfiles resources.
        :param pulumi.Input[str] admintimeout_override: Enable/disable overriding the global administrator idle timeout.
        :param pulumi.Input[str] authgrp: Administrator access to Users and Devices.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] ftviewgrp: FortiView.
        :param pulumi.Input[str] fwgrp: Administrator access to the Firewall configuration.
        :param pulumi.Input[str] loggrp: Administrator access to Logging and Reporting including viewing log messages.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[str] netgrp: Network Configuration.
        :param pulumi.Input[str] scope: Scope of admin access.
        :param pulumi.Input[str] secfabgrp: Security Fabric.
        :param pulumi.Input[str] sysgrp: System Configuration.
        :param pulumi.Input[str] utmgrp: Administrator access to Security Profiles.
        :param pulumi.Input[str] vpngrp: Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        :param pulumi.Input[str] wanoptgrp: Administrator access to WAN Opt & Cache.
        :param pulumi.Input[str] wifi: Administrator access to the WiFi controller and Switch controller.
        """
        if admintimeout_override is not None:
            pulumi.set(__self__, "admintimeout_override", admintimeout_override)
        if authgrp is not None:
            pulumi.set(__self__, "authgrp", authgrp)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if ftviewgrp is not None:
            pulumi.set(__self__, "ftviewgrp", ftviewgrp)
        if fwgrp is not None:
            pulumi.set(__self__, "fwgrp", fwgrp)
        if loggrp is not None:
            pulumi.set(__self__, "loggrp", loggrp)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if netgrp is not None:
            pulumi.set(__self__, "netgrp", netgrp)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if secfabgrp is not None:
            pulumi.set(__self__, "secfabgrp", secfabgrp)
        if sysgrp is not None:
            pulumi.set(__self__, "sysgrp", sysgrp)
        if utmgrp is not None:
            pulumi.set(__self__, "utmgrp", utmgrp)
        if vpngrp is not None:
            pulumi.set(__self__, "vpngrp", vpngrp)
        if wanoptgrp is not None:
            pulumi.set(__self__, "wanoptgrp", wanoptgrp)
        if wifi is not None:
            pulumi.set(__self__, "wifi", wifi)

    @property
    @pulumi.getter(name="admintimeoutOverride")
    def admintimeout_override(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable overriding the global administrator idle timeout.
        """
        return pulumi.get(self, "admintimeout_override")

    @admintimeout_override.setter
    def admintimeout_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admintimeout_override", value)

    @property
    @pulumi.getter
    def authgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to Users and Devices.
        """
        return pulumi.get(self, "authgrp")

    @authgrp.setter
    def authgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authgrp", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def ftviewgrp(self) -> Optional[pulumi.Input[str]]:
        """
        FortiView.
        """
        return pulumi.get(self, "ftviewgrp")

    @ftviewgrp.setter
    def ftviewgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ftviewgrp", value)

    @property
    @pulumi.getter
    def fwgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to the Firewall configuration.
        """
        return pulumi.get(self, "fwgrp")

    @fwgrp.setter
    def fwgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fwgrp", value)

    @property
    @pulumi.getter
    def loggrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to Logging and Reporting including viewing log messages.
        """
        return pulumi.get(self, "loggrp")

    @loggrp.setter
    def loggrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "loggrp", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def netgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Network Configuration.
        """
        return pulumi.get(self, "netgrp")

    @netgrp.setter
    def netgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netgrp", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        Scope of admin access.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def secfabgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Security Fabric.
        """
        return pulumi.get(self, "secfabgrp")

    @secfabgrp.setter
    def secfabgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secfabgrp", value)

    @property
    @pulumi.getter
    def sysgrp(self) -> Optional[pulumi.Input[str]]:
        """
        System Configuration.
        """
        return pulumi.get(self, "sysgrp")

    @sysgrp.setter
    def sysgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sysgrp", value)

    @property
    @pulumi.getter
    def utmgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to Security Profiles.
        """
        return pulumi.get(self, "utmgrp")

    @utmgrp.setter
    def utmgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "utmgrp", value)

    @property
    @pulumi.getter
    def vpngrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        """
        return pulumi.get(self, "vpngrp")

    @vpngrp.setter
    def vpngrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpngrp", value)

    @property
    @pulumi.getter
    def wanoptgrp(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to WAN Opt & Cache.
        """
        return pulumi.get(self, "wanoptgrp")

    @wanoptgrp.setter
    def wanoptgrp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wanoptgrp", value)

    @property
    @pulumi.getter
    def wifi(self) -> Optional[pulumi.Input[str]]:
        """
        Administrator access to the WiFi controller and Switch controller.
        """
        return pulumi.get(self, "wifi")

    @wifi.setter
    def wifi(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wifi", value)


class SystemAdminProfiles(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admintimeout_override: Optional[pulumi.Input[str]] = None,
                 authgrp: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftviewgrp: Optional[pulumi.Input[str]] = None,
                 fwgrp: Optional[pulumi.Input[str]] = None,
                 loggrp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netgrp: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 secfabgrp: Optional[pulumi.Input[str]] = None,
                 sysgrp: Optional[pulumi.Input[str]] = None,
                 utmgrp: Optional[pulumi.Input[str]] = None,
                 vpngrp: Optional[pulumi.Input[str]] = None,
                 wanoptgrp: Optional[pulumi.Input[str]] = None,
                 wifi: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to configure access profiles of FortiOS.

        !> **Warning:** The resource will be deprecated and replaced by new resource `SystemAccprofile`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test1 = fortios.SystemAdminProfiles("test1",
            admintimeout_override="disable",
            authgrp="none",
            comments="test",
            ftviewgrp="read",
            fwgrp="none",
            loggrp="none",
            netgrp="none",
            scope="vdom",
            secfabgrp="read-write",
            sysgrp="read",
            utmgrp="none",
            vpngrp="none",
            wanoptgrp="none",
            wifi="none")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admintimeout_override: Enable/disable overriding the global administrator idle timeout.
        :param pulumi.Input[str] authgrp: Administrator access to Users and Devices.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] ftviewgrp: FortiView.
        :param pulumi.Input[str] fwgrp: Administrator access to the Firewall configuration.
        :param pulumi.Input[str] loggrp: Administrator access to Logging and Reporting including viewing log messages.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[str] netgrp: Network Configuration.
        :param pulumi.Input[str] scope: Scope of admin access.
        :param pulumi.Input[str] secfabgrp: Security Fabric.
        :param pulumi.Input[str] sysgrp: System Configuration.
        :param pulumi.Input[str] utmgrp: Administrator access to Security Profiles.
        :param pulumi.Input[str] vpngrp: Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        :param pulumi.Input[str] wanoptgrp: Administrator access to WAN Opt & Cache.
        :param pulumi.Input[str] wifi: Administrator access to the WiFi controller and Switch controller.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemAdminProfilesArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to configure access profiles of FortiOS.

        !> **Warning:** The resource will be deprecated and replaced by new resource `SystemAccprofile`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test1 = fortios.SystemAdminProfiles("test1",
            admintimeout_override="disable",
            authgrp="none",
            comments="test",
            ftviewgrp="read",
            fwgrp="none",
            loggrp="none",
            netgrp="none",
            scope="vdom",
            secfabgrp="read-write",
            sysgrp="read",
            utmgrp="none",
            vpngrp="none",
            wanoptgrp="none",
            wifi="none")
        ```

        :param str resource_name: The name of the resource.
        :param SystemAdminProfilesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemAdminProfilesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admintimeout_override: Optional[pulumi.Input[str]] = None,
                 authgrp: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 ftviewgrp: Optional[pulumi.Input[str]] = None,
                 fwgrp: Optional[pulumi.Input[str]] = None,
                 loggrp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netgrp: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 secfabgrp: Optional[pulumi.Input[str]] = None,
                 sysgrp: Optional[pulumi.Input[str]] = None,
                 utmgrp: Optional[pulumi.Input[str]] = None,
                 vpngrp: Optional[pulumi.Input[str]] = None,
                 wanoptgrp: Optional[pulumi.Input[str]] = None,
                 wifi: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemAdminProfilesArgs.__new__(SystemAdminProfilesArgs)

            __props__.__dict__["admintimeout_override"] = admintimeout_override
            __props__.__dict__["authgrp"] = authgrp
            __props__.__dict__["comments"] = comments
            __props__.__dict__["ftviewgrp"] = ftviewgrp
            __props__.__dict__["fwgrp"] = fwgrp
            __props__.__dict__["loggrp"] = loggrp
            __props__.__dict__["name"] = name
            __props__.__dict__["netgrp"] = netgrp
            __props__.__dict__["scope"] = scope
            __props__.__dict__["secfabgrp"] = secfabgrp
            __props__.__dict__["sysgrp"] = sysgrp
            __props__.__dict__["utmgrp"] = utmgrp
            __props__.__dict__["vpngrp"] = vpngrp
            __props__.__dict__["wanoptgrp"] = wanoptgrp
            __props__.__dict__["wifi"] = wifi
        super(SystemAdminProfiles, __self__).__init__(
            'fortios:index/systemAdminProfiles:SystemAdminProfiles',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admintimeout_override: Optional[pulumi.Input[str]] = None,
            authgrp: Optional[pulumi.Input[str]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            ftviewgrp: Optional[pulumi.Input[str]] = None,
            fwgrp: Optional[pulumi.Input[str]] = None,
            loggrp: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            netgrp: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            secfabgrp: Optional[pulumi.Input[str]] = None,
            sysgrp: Optional[pulumi.Input[str]] = None,
            utmgrp: Optional[pulumi.Input[str]] = None,
            vpngrp: Optional[pulumi.Input[str]] = None,
            wanoptgrp: Optional[pulumi.Input[str]] = None,
            wifi: Optional[pulumi.Input[str]] = None) -> 'SystemAdminProfiles':
        """
        Get an existing SystemAdminProfiles resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admintimeout_override: Enable/disable overriding the global administrator idle timeout.
        :param pulumi.Input[str] authgrp: Administrator access to Users and Devices.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] ftviewgrp: FortiView.
        :param pulumi.Input[str] fwgrp: Administrator access to the Firewall configuration.
        :param pulumi.Input[str] loggrp: Administrator access to Logging and Reporting including viewing log messages.
        :param pulumi.Input[str] name: Profile name.
        :param pulumi.Input[str] netgrp: Network Configuration.
        :param pulumi.Input[str] scope: Scope of admin access.
        :param pulumi.Input[str] secfabgrp: Security Fabric.
        :param pulumi.Input[str] sysgrp: System Configuration.
        :param pulumi.Input[str] utmgrp: Administrator access to Security Profiles.
        :param pulumi.Input[str] vpngrp: Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        :param pulumi.Input[str] wanoptgrp: Administrator access to WAN Opt & Cache.
        :param pulumi.Input[str] wifi: Administrator access to the WiFi controller and Switch controller.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemAdminProfilesState.__new__(_SystemAdminProfilesState)

        __props__.__dict__["admintimeout_override"] = admintimeout_override
        __props__.__dict__["authgrp"] = authgrp
        __props__.__dict__["comments"] = comments
        __props__.__dict__["ftviewgrp"] = ftviewgrp
        __props__.__dict__["fwgrp"] = fwgrp
        __props__.__dict__["loggrp"] = loggrp
        __props__.__dict__["name"] = name
        __props__.__dict__["netgrp"] = netgrp
        __props__.__dict__["scope"] = scope
        __props__.__dict__["secfabgrp"] = secfabgrp
        __props__.__dict__["sysgrp"] = sysgrp
        __props__.__dict__["utmgrp"] = utmgrp
        __props__.__dict__["vpngrp"] = vpngrp
        __props__.__dict__["wanoptgrp"] = wanoptgrp
        __props__.__dict__["wifi"] = wifi
        return SystemAdminProfiles(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="admintimeoutOverride")
    def admintimeout_override(self) -> pulumi.Output[str]:
        """
        Enable/disable overriding the global administrator idle timeout.
        """
        return pulumi.get(self, "admintimeout_override")

    @property
    @pulumi.getter
    def authgrp(self) -> pulumi.Output[str]:
        """
        Administrator access to Users and Devices.
        """
        return pulumi.get(self, "authgrp")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def ftviewgrp(self) -> pulumi.Output[str]:
        """
        FortiView.
        """
        return pulumi.get(self, "ftviewgrp")

    @property
    @pulumi.getter
    def fwgrp(self) -> pulumi.Output[str]:
        """
        Administrator access to the Firewall configuration.
        """
        return pulumi.get(self, "fwgrp")

    @property
    @pulumi.getter
    def loggrp(self) -> pulumi.Output[str]:
        """
        Administrator access to Logging and Reporting including viewing log messages.
        """
        return pulumi.get(self, "loggrp")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netgrp(self) -> pulumi.Output[str]:
        """
        Network Configuration.
        """
        return pulumi.get(self, "netgrp")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Scope of admin access.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def secfabgrp(self) -> pulumi.Output[str]:
        """
        Security Fabric.
        """
        return pulumi.get(self, "secfabgrp")

    @property
    @pulumi.getter
    def sysgrp(self) -> pulumi.Output[str]:
        """
        System Configuration.
        """
        return pulumi.get(self, "sysgrp")

    @property
    @pulumi.getter
    def utmgrp(self) -> pulumi.Output[str]:
        """
        Administrator access to Security Profiles.
        """
        return pulumi.get(self, "utmgrp")

    @property
    @pulumi.getter
    def vpngrp(self) -> pulumi.Output[str]:
        """
        Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
        """
        return pulumi.get(self, "vpngrp")

    @property
    @pulumi.getter
    def wanoptgrp(self) -> pulumi.Output[str]:
        """
        Administrator access to WAN Opt & Cache.
        """
        return pulumi.get(self, "wanoptgrp")

    @property
    @pulumi.getter
    def wifi(self) -> pulumi.Output[str]:
        """
        Administrator access to the WiFi controller and Switch controller.
        """
        return pulumi.get(self, "wifi")

