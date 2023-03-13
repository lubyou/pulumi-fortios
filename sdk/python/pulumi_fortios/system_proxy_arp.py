# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemProxyArpArgs', 'SystemProxyArp']

@pulumi.input_type
class SystemProxyArpArgs:
    def __init__(__self__, *,
                 fosid: pulumi.Input[int],
                 interface: pulumi.Input[str],
                 ip: pulumi.Input[str],
                 end_ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemProxyArp resource.
        """
        pulumi.set(__self__, "fosid", fosid)
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "ip", ip)
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Input[int]:
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: pulumi.Input[int]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemProxyArpState:
    def __init__(__self__, *,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemProxyArp resources.
        """
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemProxyArp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemProxyArp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemProxyArpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemProxyArp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemProxyArpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemProxyArpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemProxyArpArgs.__new__(SystemProxyArpArgs)

            __props__.__dict__["end_ip"] = end_ip
            if fosid is None and not opts.urn:
                raise TypeError("Missing required property 'fosid'")
            __props__.__dict__["fosid"] = fosid
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            if ip is None and not opts.urn:
                raise TypeError("Missing required property 'ip'")
            __props__.__dict__["ip"] = ip
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemProxyArp, __self__).__init__(
            'fortios:index/systemProxyArp:SystemProxyArp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            end_ip: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemProxyArp':
        """
        Get an existing SystemProxyArp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemProxyArpState.__new__(_SystemProxyArpState)

        __props__.__dict__["end_ip"] = end_ip
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["interface"] = interface
        __props__.__dict__["ip"] = ip
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemProxyArp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

