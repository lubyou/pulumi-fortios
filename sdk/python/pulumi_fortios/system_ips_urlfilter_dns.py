# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemIpsUrlfilterDnsArgs', 'SystemIpsUrlfilterDns']

@pulumi.input_type
class SystemIpsUrlfilterDnsArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 ipv6_capability: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemIpsUrlfilterDns resource.
        :param pulumi.Input[str] address: DNS server IP address.
        :param pulumi.Input[str] ipv6_capability: Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if ipv6_capability is not None:
            pulumi.set(__self__, "ipv6_capability", ipv6_capability)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        DNS server IP address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="ipv6Capability")
    def ipv6_capability(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ipv6_capability")

    @ipv6_capability.setter
    def ipv6_capability(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_capability", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemIpsUrlfilterDnsState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 ipv6_capability: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemIpsUrlfilterDns resources.
        :param pulumi.Input[str] address: DNS server IP address.
        :param pulumi.Input[str] ipv6_capability: Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if ipv6_capability is not None:
            pulumi.set(__self__, "ipv6_capability", ipv6_capability)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        DNS server IP address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="ipv6Capability")
    def ipv6_capability(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ipv6_capability")

    @ipv6_capability.setter
    def ipv6_capability(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_capability", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemIpsUrlfilterDns(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 ipv6_capability: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPS URL filter DNS servers.

        ## Import

        System IpsUrlfilterDns can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemIpsUrlfilterDns:SystemIpsUrlfilterDns labelname {{address}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: DNS server IP address.
        :param pulumi.Input[str] ipv6_capability: Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemIpsUrlfilterDnsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPS URL filter DNS servers.

        ## Import

        System IpsUrlfilterDns can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemIpsUrlfilterDns:SystemIpsUrlfilterDns labelname {{address}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemIpsUrlfilterDnsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemIpsUrlfilterDnsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 ipv6_capability: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemIpsUrlfilterDnsArgs.__new__(SystemIpsUrlfilterDnsArgs)

            __props__.__dict__["address"] = address
            __props__.__dict__["ipv6_capability"] = ipv6_capability
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemIpsUrlfilterDns, __self__).__init__(
            'fortios:index/systemIpsUrlfilterDns:SystemIpsUrlfilterDns',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            ipv6_capability: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemIpsUrlfilterDns':
        """
        Get an existing SystemIpsUrlfilterDns resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: DNS server IP address.
        :param pulumi.Input[str] ipv6_capability: Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] status: Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemIpsUrlfilterDnsState.__new__(_SystemIpsUrlfilterDnsState)

        __props__.__dict__["address"] = address
        __props__.__dict__["ipv6_capability"] = ipv6_capability
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemIpsUrlfilterDns(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        DNS server IP address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="ipv6Capability")
    def ipv6_capability(self) -> pulumi.Output[str]:
        """
        Enable/disable this server for IPv6 queries. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ipv6_capability")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable using this DNS server for IPS URL filter DNS queries. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

