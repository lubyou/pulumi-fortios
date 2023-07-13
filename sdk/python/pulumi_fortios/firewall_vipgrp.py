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

__all__ = ['FirewallVipgrpArgs', 'FirewallVipgrp']

@pulumi.input_type
class FirewallVipgrpArgs:
    def __init__(__self__, *,
                 interface: pulumi.Input[str],
                 members: pulumi.Input[Sequence[pulumi.Input['FirewallVipgrpMemberArgs']]],
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallVipgrp resource.
        """
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "members", members)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallVipgrpMemberArgs']]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallVipgrpMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallVipgrpState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVipgrpMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallVipgrp resources.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVipgrpMemberArgs']]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVipgrpMemberArgs']]]]):
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
    def uuid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallVipgrp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVipgrpMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallVipgrp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallVipgrpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallVipgrp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallVipgrpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallVipgrpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVipgrpMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallVipgrpArgs.__new__(FirewallVipgrpArgs)

            __props__.__dict__["color"] = color
            __props__.__dict__["comments"] = comments
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallVipgrp, __self__).__init__(
            'fortios:index/firewallVipgrp:FirewallVipgrp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[int]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVipgrpMemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallVipgrp':
        """
        Get an existing FirewallVipgrp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallVipgrpState.__new__(_FirewallVipgrpState)

        __props__.__dict__["color"] = color
        __props__.__dict__["comments"] = comments
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["interface"] = interface
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["uuid"] = uuid
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallVipgrp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[int]:
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.FirewallVipgrpMember']]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

