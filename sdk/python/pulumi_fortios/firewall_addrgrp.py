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

__all__ = ['FirewallAddrgrpArgs', 'FirewallAddrgrp']

@pulumi.input_type
class FirewallAddrgrpArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpMemberArgs']]],
                 allow_routing: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 exclude: Optional[pulumi.Input[str]] = None,
                 exclude_members: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpExcludeMemberArgs']]]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpTaggingArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallAddrgrp resource.
        """
        pulumi.set(__self__, "members", members)
        if allow_routing is not None:
            pulumi.set(__self__, "allow_routing", allow_routing)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if exclude is not None:
            pulumi.set(__self__, "exclude", exclude)
        if exclude_members is not None:
            pulumi.set(__self__, "exclude_members", exclude_members)
        if fabric_object is not None:
            pulumi.set(__self__, "fabric_object", fabric_object)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if taggings is not None:
            pulumi.set(__self__, "taggings", taggings)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpMemberArgs']]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="allowRouting")
    def allow_routing(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "allow_routing")

    @allow_routing.setter
    def allow_routing(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_routing", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

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
    @pulumi.getter
    def exclude(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exclude")

    @exclude.setter
    def exclude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclude", value)

    @property
    @pulumi.getter(name="excludeMembers")
    def exclude_members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpExcludeMemberArgs']]]]:
        return pulumi.get(self, "exclude_members")

    @exclude_members.setter
    def exclude_members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpExcludeMemberArgs']]]]):
        pulumi.set(self, "exclude_members", value)

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fabric_object")

    @fabric_object.setter
    def fabric_object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fabric_object", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def taggings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpTaggingArgs']]]]:
        return pulumi.get(self, "taggings")

    @taggings.setter
    def taggings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpTaggingArgs']]]]):
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

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


@pulumi.input_type
class _FirewallAddrgrpState:
    def __init__(__self__, *,
                 allow_routing: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 exclude: Optional[pulumi.Input[str]] = None,
                 exclude_members: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpExcludeMemberArgs']]]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpTaggingArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallAddrgrp resources.
        """
        if allow_routing is not None:
            pulumi.set(__self__, "allow_routing", allow_routing)
        if category is not None:
            pulumi.set(__self__, "category", category)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if exclude is not None:
            pulumi.set(__self__, "exclude", exclude)
        if exclude_members is not None:
            pulumi.set(__self__, "exclude_members", exclude_members)
        if fabric_object is not None:
            pulumi.set(__self__, "fabric_object", fabric_object)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if taggings is not None:
            pulumi.set(__self__, "taggings", taggings)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter(name="allowRouting")
    def allow_routing(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "allow_routing")

    @allow_routing.setter
    def allow_routing(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_routing", value)

    @property
    @pulumi.getter
    def category(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "category")

    @category.setter
    def category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category", value)

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
    @pulumi.getter
    def exclude(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exclude")

    @exclude.setter
    def exclude(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclude", value)

    @property
    @pulumi.getter(name="excludeMembers")
    def exclude_members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpExcludeMemberArgs']]]]:
        return pulumi.get(self, "exclude_members")

    @exclude_members.setter
    def exclude_members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpExcludeMemberArgs']]]]):
        pulumi.set(self, "exclude_members", value)

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fabric_object")

    @fabric_object.setter
    def fabric_object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fabric_object", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpMemberArgs']]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpMemberArgs']]]]):
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
    def taggings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpTaggingArgs']]]]:
        return pulumi.get(self, "taggings")

    @taggings.setter
    def taggings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAddrgrpTaggingArgs']]]]):
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

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


class FirewallAddrgrp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_routing: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 exclude: Optional[pulumi.Input[str]] = None,
                 exclude_members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpExcludeMemberArgs']]]]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpTaggingArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallAddrgrp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallAddrgrpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallAddrgrp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallAddrgrpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallAddrgrpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_routing: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 exclude: Optional[pulumi.Input[str]] = None,
                 exclude_members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpExcludeMemberArgs']]]]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpTaggingArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallAddrgrpArgs.__new__(FirewallAddrgrpArgs)

            __props__.__dict__["allow_routing"] = allow_routing
            __props__.__dict__["category"] = category
            __props__.__dict__["color"] = color
            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["exclude"] = exclude
            __props__.__dict__["exclude_members"] = exclude_members
            __props__.__dict__["fabric_object"] = fabric_object
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["taggings"] = taggings
            __props__.__dict__["type"] = type
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["visibility"] = visibility
        super(FirewallAddrgrp, __self__).__init__(
            'fortios:index/firewallAddrgrp:FirewallAddrgrp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_routing: Optional[pulumi.Input[str]] = None,
            category: Optional[pulumi.Input[str]] = None,
            color: Optional[pulumi.Input[int]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            exclude: Optional[pulumi.Input[str]] = None,
            exclude_members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpExcludeMemberArgs']]]]] = None,
            fabric_object: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpMemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            taggings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAddrgrpTaggingArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            visibility: Optional[pulumi.Input[str]] = None) -> 'FirewallAddrgrp':
        """
        Get an existing FirewallAddrgrp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallAddrgrpState.__new__(_FirewallAddrgrpState)

        __props__.__dict__["allow_routing"] = allow_routing
        __props__.__dict__["category"] = category
        __props__.__dict__["color"] = color
        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["exclude"] = exclude
        __props__.__dict__["exclude_members"] = exclude_members
        __props__.__dict__["fabric_object"] = fabric_object
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["taggings"] = taggings
        __props__.__dict__["type"] = type
        __props__.__dict__["uuid"] = uuid
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["visibility"] = visibility
        return FirewallAddrgrp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowRouting")
    def allow_routing(self) -> pulumi.Output[str]:
        return pulumi.get(self, "allow_routing")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[str]:
        return pulumi.get(self, "category")

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
    @pulumi.getter
    def exclude(self) -> pulumi.Output[str]:
        return pulumi.get(self, "exclude")

    @property
    @pulumi.getter(name="excludeMembers")
    def exclude_members(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallAddrgrpExcludeMember']]]:
        return pulumi.get(self, "exclude_members")

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fabric_object")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.FirewallAddrgrpMember']]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def taggings(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallAddrgrpTagging']]]:
        return pulumi.get(self, "taggings")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[str]:
        return pulumi.get(self, "visibility")

