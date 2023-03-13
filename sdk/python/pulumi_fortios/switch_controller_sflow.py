# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SwitchControllerSflowArgs', 'SwitchControllerSflow']

@pulumi.input_type
class SwitchControllerSflowArgs:
    def __init__(__self__, *,
                 collector_ip: pulumi.Input[str],
                 collector_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerSflow resource.
        """
        pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port is not None:
            pulumi.set(__self__, "collector_port", collector_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "collector_ip")

    @collector_ip.setter
    def collector_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "collector_ip", value)

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "collector_port")

    @collector_port.setter
    def collector_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "collector_port", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SwitchControllerSflowState:
    def __init__(__self__, *,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerSflow resources.
        """
        if collector_ip is not None:
            pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port is not None:
            pulumi.set(__self__, "collector_port", collector_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "collector_ip")

    @collector_ip.setter
    def collector_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collector_ip", value)

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "collector_port")

    @collector_port.setter
    def collector_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "collector_port", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SwitchControllerSflow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SwitchControllerSflow resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SwitchControllerSflowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SwitchControllerSflow resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SwitchControllerSflowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerSflowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchControllerSflowArgs.__new__(SwitchControllerSflowArgs)

            if collector_ip is None and not opts.urn:
                raise TypeError("Missing required property 'collector_ip'")
            __props__.__dict__["collector_ip"] = collector_ip
            __props__.__dict__["collector_port"] = collector_port
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerSflow, __self__).__init__(
            'fortios:index/switchControllerSflow:SwitchControllerSflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            collector_ip: Optional[pulumi.Input[str]] = None,
            collector_port: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerSflow':
        """
        Get an existing SwitchControllerSflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerSflowState.__new__(_SwitchControllerSflowState)

        __props__.__dict__["collector_ip"] = collector_ip
        __props__.__dict__["collector_port"] = collector_port
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerSflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="collectorIp")
    def collector_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "collector_ip")

    @property
    @pulumi.getter(name="collectorPort")
    def collector_port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "collector_port")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

