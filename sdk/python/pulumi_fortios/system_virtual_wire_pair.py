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

__all__ = ['SystemVirtualWirePairArgs', 'SystemVirtualWirePair']

@pulumi.input_type
class SystemVirtualWirePairArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input['SystemVirtualWirePairMemberArgs']]],
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan_filter: Optional[pulumi.Input[str]] = None,
                 wildcard_vlan: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemVirtualWirePair resource.
        """
        pulumi.set(__self__, "members", members)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan_filter is not None:
            pulumi.set(__self__, "vlan_filter", vlan_filter)
        if wildcard_vlan is not None:
            pulumi.set(__self__, "wildcard_vlan", wildcard_vlan)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['SystemVirtualWirePairMemberArgs']]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['SystemVirtualWirePairMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="vlanFilter")
    def vlan_filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vlan_filter")

    @vlan_filter.setter
    def vlan_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan_filter", value)

    @property
    @pulumi.getter(name="wildcardVlan")
    def wildcard_vlan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wildcard_vlan")

    @wildcard_vlan.setter
    def wildcard_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wildcard_vlan", value)


@pulumi.input_type
class _SystemVirtualWirePairState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['SystemVirtualWirePairMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan_filter: Optional[pulumi.Input[str]] = None,
                 wildcard_vlan: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemVirtualWirePair resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan_filter is not None:
            pulumi.set(__self__, "vlan_filter", vlan_filter)
        if wildcard_vlan is not None:
            pulumi.set(__self__, "wildcard_vlan", wildcard_vlan)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemVirtualWirePairMemberArgs']]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemVirtualWirePairMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="vlanFilter")
    def vlan_filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vlan_filter")

    @vlan_filter.setter
    def vlan_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan_filter", value)

    @property
    @pulumi.getter(name="wildcardVlan")
    def wildcard_vlan(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wildcard_vlan")

    @wildcard_vlan.setter
    def wildcard_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wildcard_vlan", value)


class SystemVirtualWirePair(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemVirtualWirePairMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan_filter: Optional[pulumi.Input[str]] = None,
                 wildcard_vlan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemVirtualWirePair resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemVirtualWirePairArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemVirtualWirePair resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemVirtualWirePairArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemVirtualWirePairArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemVirtualWirePairMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan_filter: Optional[pulumi.Input[str]] = None,
                 wildcard_vlan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemVirtualWirePairArgs.__new__(SystemVirtualWirePairArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vlan_filter"] = vlan_filter
            __props__.__dict__["wildcard_vlan"] = wildcard_vlan
        super(SystemVirtualWirePair, __self__).__init__(
            'fortios:index/systemVirtualWirePair:SystemVirtualWirePair',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemVirtualWirePairMemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vlan_filter: Optional[pulumi.Input[str]] = None,
            wildcard_vlan: Optional[pulumi.Input[str]] = None) -> 'SystemVirtualWirePair':
        """
        Get an existing SystemVirtualWirePair resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemVirtualWirePairState.__new__(_SystemVirtualWirePairState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vlan_filter"] = vlan_filter
        __props__.__dict__["wildcard_vlan"] = wildcard_vlan
        return SystemVirtualWirePair(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.SystemVirtualWirePairMember']]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="vlanFilter")
    def vlan_filter(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vlan_filter")

    @property
    @pulumi.getter(name="wildcardVlan")
    def wildcard_vlan(self) -> pulumi.Output[str]:
        return pulumi.get(self, "wildcard_vlan")

