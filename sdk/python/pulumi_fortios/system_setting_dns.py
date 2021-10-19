# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemSettingDNSArgs', 'SystemSettingDNS']

@pulumi.input_type
class SystemSettingDNSArgs:
    def __init__(__self__, *,
                 dns_over_tls: Optional[pulumi.Input[str]] = None,
                 primary: Optional[pulumi.Input[str]] = None,
                 secondary: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemSettingDNS resource.
        :param pulumi.Input[str] dns_over_tls: Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        :param pulumi.Input[str] primary: Primary DNS server IP address.
        :param pulumi.Input[str] secondary: Secondary DNS server IP address.
        """
        if dns_over_tls is not None:
            pulumi.set(__self__, "dns_over_tls", dns_over_tls)
        if primary is not None:
            pulumi.set(__self__, "primary", primary)
        if secondary is not None:
            pulumi.set(__self__, "secondary", secondary)

    @property
    @pulumi.getter(name="dnsOverTls")
    def dns_over_tls(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        """
        return pulumi.get(self, "dns_over_tls")

    @dns_over_tls.setter
    def dns_over_tls(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_over_tls", value)

    @property
    @pulumi.getter
    def primary(self) -> Optional[pulumi.Input[str]]:
        """
        Primary DNS server IP address.
        """
        return pulumi.get(self, "primary")

    @primary.setter
    def primary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary", value)

    @property
    @pulumi.getter
    def secondary(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary DNS server IP address.
        """
        return pulumi.get(self, "secondary")

    @secondary.setter
    def secondary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary", value)


@pulumi.input_type
class _SystemSettingDNSState:
    def __init__(__self__, *,
                 dns_over_tls: Optional[pulumi.Input[str]] = None,
                 primary: Optional[pulumi.Input[str]] = None,
                 secondary: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemSettingDNS resources.
        :param pulumi.Input[str] dns_over_tls: Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        :param pulumi.Input[str] primary: Primary DNS server IP address.
        :param pulumi.Input[str] secondary: Secondary DNS server IP address.
        """
        if dns_over_tls is not None:
            pulumi.set(__self__, "dns_over_tls", dns_over_tls)
        if primary is not None:
            pulumi.set(__self__, "primary", primary)
        if secondary is not None:
            pulumi.set(__self__, "secondary", secondary)

    @property
    @pulumi.getter(name="dnsOverTls")
    def dns_over_tls(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        """
        return pulumi.get(self, "dns_over_tls")

    @dns_over_tls.setter
    def dns_over_tls(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_over_tls", value)

    @property
    @pulumi.getter
    def primary(self) -> Optional[pulumi.Input[str]]:
        """
        Primary DNS server IP address.
        """
        return pulumi.get(self, "primary")

    @primary.setter
    def primary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary", value)

    @property
    @pulumi.getter
    def secondary(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary DNS server IP address.
        """
        return pulumi.get(self, "secondary")

    @secondary.setter
    def secondary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary", value)


class SystemSettingDNS(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_over_tls: Optional[pulumi.Input[str]] = None,
                 primary: Optional[pulumi.Input[str]] = None,
                 secondary: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to configure DNS of FortiOS.

        !> **Warning:** The resource will be deprecated and replaced by new resource `SystemDns`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test1 = fortios.SystemSettingDNS("test1",
            dns_over_tls="disable",
            primary="208.91.112.53",
            secondary="208.91.112.22")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_over_tls: Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        :param pulumi.Input[str] primary: Primary DNS server IP address.
        :param pulumi.Input[str] secondary: Secondary DNS server IP address.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemSettingDNSArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to configure DNS of FortiOS.

        !> **Warning:** The resource will be deprecated and replaced by new resource `SystemDns`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test1 = fortios.SystemSettingDNS("test1",
            dns_over_tls="disable",
            primary="208.91.112.53",
            secondary="208.91.112.22")
        ```

        :param str resource_name: The name of the resource.
        :param SystemSettingDNSArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemSettingDNSArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_over_tls: Optional[pulumi.Input[str]] = None,
                 primary: Optional[pulumi.Input[str]] = None,
                 secondary: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemSettingDNSArgs.__new__(SystemSettingDNSArgs)

            __props__.__dict__["dns_over_tls"] = dns_over_tls
            __props__.__dict__["primary"] = primary
            __props__.__dict__["secondary"] = secondary
        super(SystemSettingDNS, __self__).__init__(
            'fortios:index/systemSettingDNS:SystemSettingDNS',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dns_over_tls: Optional[pulumi.Input[str]] = None,
            primary: Optional[pulumi.Input[str]] = None,
            secondary: Optional[pulumi.Input[str]] = None) -> 'SystemSettingDNS':
        """
        Get an existing SystemSettingDNS resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_over_tls: Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        :param pulumi.Input[str] primary: Primary DNS server IP address.
        :param pulumi.Input[str] secondary: Secondary DNS server IP address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemSettingDNSState.__new__(_SystemSettingDNSState)

        __props__.__dict__["dns_over_tls"] = dns_over_tls
        __props__.__dict__["primary"] = primary
        __props__.__dict__["secondary"] = secondary
        return SystemSettingDNS(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsOverTls")
    def dns_over_tls(self) -> pulumi.Output[str]:
        """
        Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]
        """
        return pulumi.get(self, "dns_over_tls")

    @property
    @pulumi.getter
    def primary(self) -> pulumi.Output[str]:
        """
        Primary DNS server IP address.
        """
        return pulumi.get(self, "primary")

    @property
    @pulumi.getter
    def secondary(self) -> pulumi.Output[str]:
        """
        Secondary DNS server IP address.
        """
        return pulumi.get(self, "secondary")

