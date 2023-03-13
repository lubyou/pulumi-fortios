# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FirewallMulticastAddressArgs', 'FirewallMulticastAddress']

@pulumi.input_type
class FirewallMulticastAddressArgs:
    def __init__(__self__, *,
                 associated_interface: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastAddressTaggingArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallMulticastAddress resource.
        """
        if associated_interface is not None:
            pulumi.set(__self__, "associated_interface", associated_interface)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if taggings is not None:
            pulumi.set(__self__, "taggings", taggings)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter(name="associatedInterface")
    def associated_interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "associated_interface")

    @associated_interface.setter
    def associated_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associated_interface", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def taggings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastAddressTaggingArgs']]]]:
        return pulumi.get(self, "taggings")

    @taggings.setter
    def taggings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastAddressTaggingArgs']]]]):
        pulumi.set(self, "taggings", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


@pulumi.input_type
class _FirewallMulticastAddressState:
    def __init__(__self__, *,
                 associated_interface: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastAddressTaggingArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallMulticastAddress resources.
        """
        if associated_interface is not None:
            pulumi.set(__self__, "associated_interface", associated_interface)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if end_ip is not None:
            pulumi.set(__self__, "end_ip", end_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if start_ip is not None:
            pulumi.set(__self__, "start_ip", start_ip)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if taggings is not None:
            pulumi.set(__self__, "taggings", taggings)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter(name="associatedInterface")
    def associated_interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "associated_interface")

    @associated_interface.setter
    def associated_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "associated_interface", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end_ip")

    @end_ip.setter
    def end_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "start_ip")

    @start_ip.setter
    def start_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_ip", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def taggings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastAddressTaggingArgs']]]]:
        return pulumi.get(self, "taggings")

    @taggings.setter
    def taggings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallMulticastAddressTaggingArgs']]]]):
        pulumi.set(self, "taggings", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


class FirewallMulticastAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associated_interface: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastAddressTaggingArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallMulticastAddress resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallMulticastAddressArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallMulticastAddress resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallMulticastAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallMulticastAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associated_interface: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 end_ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start_ip: Optional[pulumi.Input[str]] = None,
                 subnet: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastAddressTaggingArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallMulticastAddressArgs.__new__(FirewallMulticastAddressArgs)

            __props__.__dict__["associated_interface"] = associated_interface
            __props__.__dict__["color"] = color
            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["end_ip"] = end_ip
            __props__.__dict__["name"] = name
            __props__.__dict__["start_ip"] = start_ip
            __props__.__dict__["subnet"] = subnet
            __props__.__dict__["taggings"] = taggings
            __props__.__dict__["type"] = type
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["visibility"] = visibility
        super(FirewallMulticastAddress, __self__).__init__(
            'fortios:index/firewallMulticastAddress:FirewallMulticastAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            associated_interface: Optional[pulumi.Input[str]] = None,
            color: Optional[pulumi.Input[int]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            end_ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            start_ip: Optional[pulumi.Input[str]] = None,
            subnet: Optional[pulumi.Input[str]] = None,
            taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallMulticastAddressTaggingArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            visibility: Optional[pulumi.Input[str]] = None) -> 'FirewallMulticastAddress':
        """
        Get an existing FirewallMulticastAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallMulticastAddressState.__new__(_FirewallMulticastAddressState)

        __props__.__dict__["associated_interface"] = associated_interface
        __props__.__dict__["color"] = color
        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["end_ip"] = end_ip
        __props__.__dict__["name"] = name
        __props__.__dict__["start_ip"] = start_ip
        __props__.__dict__["subnet"] = subnet
        __props__.__dict__["taggings"] = taggings
        __props__.__dict__["type"] = type
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["visibility"] = visibility
        return FirewallMulticastAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="associatedInterface")
    def associated_interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "associated_interface")

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[int]:
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="endIp")
    def end_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "end_ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startIp")
    def start_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "start_ip")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def taggings(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallMulticastAddressTagging']]]:
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[str]:
        return pulumi.get(self, "visibility")

