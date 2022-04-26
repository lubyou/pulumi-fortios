# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FirewallWildcardFqdnGroupArgs', 'FirewallWildcardFqdnGroup']

@pulumi.input_type
class FirewallWildcardFqdnGroupArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input['FirewallWildcardFqdnGroupMemberArgs']]],
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallWildcardFqdnGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallWildcardFqdnGroupMemberArgs']]] members: Address group members. The structure of `member` block is documented below.
        :param pulumi.Input[int] color: GUI icon color.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Address name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] visibility: Enable/disable address visibility. Valid values: `enable`, `disable`.
        """
        pulumi.set(__self__, "members", members)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallWildcardFqdnGroupMemberArgs']]]:
        """
        Address group members. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallWildcardFqdnGroupMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        GUI icon color.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Address name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

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

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable address visibility. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


@pulumi.input_type
class _FirewallWildcardFqdnGroupState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallWildcardFqdnGroupMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallWildcardFqdnGroup resources.
        :param pulumi.Input[int] color: GUI icon color.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallWildcardFqdnGroupMemberArgs']]] members: Address group members. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Address name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] visibility: Enable/disable address visibility. Valid values: `enable`, `disable`.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        GUI icon color.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallWildcardFqdnGroupMemberArgs']]]]:
        """
        Address group members. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallWildcardFqdnGroupMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Address name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)

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

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable address visibility. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


class FirewallWildcardFqdnGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallWildcardFqdnGroupMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Config global Wildcard FQDN address groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.FirewallWildcardFqdnCustom("trname1",
            color=0,
            visibility="enable",
            wildcard_fqdn="*.ms.com")
        trname = fortios.FirewallWildcardFqdnGroup("trname",
            color=0,
            visibility="enable",
            members=[fortios.FirewallWildcardFqdnGroupMemberArgs(
                name=trname1.name,
            )])
        ```

        ## Import

        FirewallWildcardFqdn Group can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] color: GUI icon color.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallWildcardFqdnGroupMemberArgs']]]] members: Address group members. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Address name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] visibility: Enable/disable address visibility. Valid values: `enable`, `disable`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallWildcardFqdnGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Config global Wildcard FQDN address groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.FirewallWildcardFqdnCustom("trname1",
            color=0,
            visibility="enable",
            wildcard_fqdn="*.ms.com")
        trname = fortios.FirewallWildcardFqdnGroup("trname",
            color=0,
            visibility="enable",
            members=[fortios.FirewallWildcardFqdnGroupMemberArgs(
                name=trname1.name,
            )])
        ```

        ## Import

        FirewallWildcardFqdn Group can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallWildcardFqdnGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallWildcardFqdnGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallWildcardFqdnGroupMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallWildcardFqdnGroupArgs.__new__(FirewallWildcardFqdnGroupArgs)

            __props__.__dict__["color"] = color
            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["visibility"] = visibility
        super(FirewallWildcardFqdnGroup, __self__).__init__(
            'fortios:index/firewallWildcardFqdnGroup:FirewallWildcardFqdnGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[int]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallWildcardFqdnGroupMemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            visibility: Optional[pulumi.Input[str]] = None) -> 'FirewallWildcardFqdnGroup':
        """
        Get an existing FirewallWildcardFqdnGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] color: GUI icon color.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallWildcardFqdnGroupMemberArgs']]]] members: Address group members. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Address name.
        :param pulumi.Input[str] uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] visibility: Enable/disable address visibility. Valid values: `enable`, `disable`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallWildcardFqdnGroupState.__new__(_FirewallWildcardFqdnGroupState)

        __props__.__dict__["color"] = color
        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["uuid"] = uuid
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["visibility"] = visibility
        return FirewallWildcardFqdnGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[int]:
        """
        GUI icon color.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.FirewallWildcardFqdnGroupMember']]:
        """
        Address group members. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Address name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[str]:
        """
        Enable/disable address visibility. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "visibility")

