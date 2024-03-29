# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemVdomSflowArgs', 'SystemVdomSflow']

@pulumi.input_type
class SystemVdomSflowArgs:
    def __init__(__self__, *,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 vdom_sflow: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemVdomSflow resource.
        """
        if collector_ip is not None:
            pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port is not None:
            pulumi.set(__self__, "collector_port", collector_port)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if vdom_sflow is not None:
            pulumi.set(__self__, "vdom_sflow", vdom_sflow)
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
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="vdomSflow")
    def vdom_sflow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdom_sflow")

    @vdom_sflow.setter
    def vdom_sflow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom_sflow", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemVdomSflowState:
    def __init__(__self__, *,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 vdom_sflow: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemVdomSflow resources.
        """
        if collector_ip is not None:
            pulumi.set(__self__, "collector_ip", collector_ip)
        if collector_port is not None:
            pulumi.set(__self__, "collector_port", collector_port)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if vdom_sflow is not None:
            pulumi.set(__self__, "vdom_sflow", vdom_sflow)
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
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="vdomSflow")
    def vdom_sflow(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdom_sflow")

    @vdom_sflow.setter
    def vdom_sflow(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom_sflow", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemVdomSflow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 vdom_sflow: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemVdomSflow resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemVdomSflowArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemVdomSflow resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemVdomSflowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemVdomSflowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collector_ip: Optional[pulumi.Input[str]] = None,
                 collector_port: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 vdom_sflow: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemVdomSflowArgs.__new__(SystemVdomSflowArgs)

            __props__.__dict__["collector_ip"] = collector_ip
            __props__.__dict__["collector_port"] = collector_port
            __props__.__dict__["interface"] = interface
            __props__.__dict__["interface_select_method"] = interface_select_method
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["vdom_sflow"] = vdom_sflow
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemVdomSflow, __self__).__init__(
            'fortios:index/systemVdomSflow:SystemVdomSflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            collector_ip: Optional[pulumi.Input[str]] = None,
            collector_port: Optional[pulumi.Input[int]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            interface_select_method: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            vdom_sflow: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemVdomSflow':
        """
        Get an existing SystemVdomSflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemVdomSflowState.__new__(_SystemVdomSflowState)

        __props__.__dict__["collector_ip"] = collector_ip
        __props__.__dict__["collector_port"] = collector_port
        __props__.__dict__["interface"] = interface
        __props__.__dict__["interface_select_method"] = interface_select_method
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["vdom_sflow"] = vdom_sflow
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemVdomSflow(resource_name, opts=opts, __props__=__props__)

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
    def interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface_select_method")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="vdomSflow")
    def vdom_sflow(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vdom_sflow")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

