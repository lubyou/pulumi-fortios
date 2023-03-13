# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallInternetServiceAppendArgs', 'FirewallInternetServiceAppend']

@pulumi.input_type
class FirewallInternetServiceAppendArgs:
    def __init__(__self__, *,
                 addr_mode: Optional[pulumi.Input[str]] = None,
                 append_port: Optional[pulumi.Input[int]] = None,
                 match_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallInternetServiceAppend resource.
        """
        if addr_mode is not None:
            pulumi.set(__self__, "addr_mode", addr_mode)
        if append_port is not None:
            pulumi.set(__self__, "append_port", append_port)
        if match_port is not None:
            pulumi.set(__self__, "match_port", match_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addrMode")
    def addr_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "addr_mode")

    @addr_mode.setter
    def addr_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "addr_mode", value)

    @property
    @pulumi.getter(name="appendPort")
    def append_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "append_port")

    @append_port.setter
    def append_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "append_port", value)

    @property
    @pulumi.getter(name="matchPort")
    def match_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "match_port")

    @match_port.setter
    def match_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "match_port", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallInternetServiceAppendState:
    def __init__(__self__, *,
                 addr_mode: Optional[pulumi.Input[str]] = None,
                 append_port: Optional[pulumi.Input[int]] = None,
                 match_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallInternetServiceAppend resources.
        """
        if addr_mode is not None:
            pulumi.set(__self__, "addr_mode", addr_mode)
        if append_port is not None:
            pulumi.set(__self__, "append_port", append_port)
        if match_port is not None:
            pulumi.set(__self__, "match_port", match_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addrMode")
    def addr_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "addr_mode")

    @addr_mode.setter
    def addr_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "addr_mode", value)

    @property
    @pulumi.getter(name="appendPort")
    def append_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "append_port")

    @append_port.setter
    def append_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "append_port", value)

    @property
    @pulumi.getter(name="matchPort")
    def match_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "match_port")

    @match_port.setter
    def match_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "match_port", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallInternetServiceAppend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addr_mode: Optional[pulumi.Input[str]] = None,
                 append_port: Optional[pulumi.Input[int]] = None,
                 match_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallInternetServiceAppend resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallInternetServiceAppendArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallInternetServiceAppend resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallInternetServiceAppendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallInternetServiceAppendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addr_mode: Optional[pulumi.Input[str]] = None,
                 append_port: Optional[pulumi.Input[int]] = None,
                 match_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallInternetServiceAppendArgs.__new__(FirewallInternetServiceAppendArgs)

            __props__.__dict__["addr_mode"] = addr_mode
            __props__.__dict__["append_port"] = append_port
            __props__.__dict__["match_port"] = match_port
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallInternetServiceAppend, __self__).__init__(
            'fortios:index/firewallInternetServiceAppend:FirewallInternetServiceAppend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            addr_mode: Optional[pulumi.Input[str]] = None,
            append_port: Optional[pulumi.Input[int]] = None,
            match_port: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallInternetServiceAppend':
        """
        Get an existing FirewallInternetServiceAppend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallInternetServiceAppendState.__new__(_FirewallInternetServiceAppendState)

        __props__.__dict__["addr_mode"] = addr_mode
        __props__.__dict__["append_port"] = append_port
        __props__.__dict__["match_port"] = match_port
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallInternetServiceAppend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addrMode")
    def addr_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "addr_mode")

    @property
    @pulumi.getter(name="appendPort")
    def append_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "append_port")

    @property
    @pulumi.getter(name="matchPort")
    def match_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "match_port")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

