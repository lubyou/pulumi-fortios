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

__all__ = ['FirewallScheduleGroupArgs', 'FirewallScheduleGroup']

@pulumi.input_type
class FirewallScheduleGroupArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input['FirewallScheduleGroupMemberArgs']]],
                 color: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallScheduleGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallScheduleGroupMemberArgs']]] members: Schedules added to the schedule group. The structure of `member` block is documented below.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fabric_object: Security Fabric global object setting. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Schedule name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "members", members)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fabric_object is not None:
            pulumi.set(__self__, "fabric_object", fabric_object)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallScheduleGroupMemberArgs']]]:
        """
        Schedules added to the schedule group. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallScheduleGroupMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

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
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> Optional[pulumi.Input[str]]:
        """
        Security Fabric global object setting. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fabric_object")

    @fabric_object.setter
    def fabric_object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fabric_object", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _FirewallScheduleGroupState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallScheduleGroupMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallScheduleGroup resources.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fabric_object: Security Fabric global object setting. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallScheduleGroupMemberArgs']]] members: Schedules added to the schedule group. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Schedule name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fabric_object is not None:
            pulumi.set(__self__, "fabric_object", fabric_object)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

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
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> Optional[pulumi.Input[str]]:
        """
        Security Fabric global object setting. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fabric_object")

    @fabric_object.setter
    def fabric_object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fabric_object", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallScheduleGroupMemberArgs']]]]:
        """
        Schedules added to the schedule group. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallScheduleGroupMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class FirewallScheduleGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallScheduleGroupMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Schedule group configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.FirewallScheduleRecurring("trname1",
            color=0,
            day="sunday",
            end="00:00",
            start="00:00")
        trname = fortios.FirewallScheduleGroup("trname",
            color=0,
            members=[fortios.FirewallScheduleGroupMemberArgs(
                name=trname1.name,
            )])
        ```

        ## Import

        FirewallSchedule Group can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallScheduleGroup:FirewallScheduleGroup labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallScheduleGroup:FirewallScheduleGroup labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fabric_object: Security Fabric global object setting. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallScheduleGroupMemberArgs']]]] members: Schedules added to the schedule group. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Schedule name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallScheduleGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Schedule group configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname1 = fortios.FirewallScheduleRecurring("trname1",
            color=0,
            day="sunday",
            end="00:00",
            start="00:00")
        trname = fortios.FirewallScheduleGroup("trname",
            color=0,
            members=[fortios.FirewallScheduleGroupMemberArgs(
                name=trname1.name,
            )])
        ```

        ## Import

        FirewallSchedule Group can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallScheduleGroup:FirewallScheduleGroup labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallScheduleGroup:FirewallScheduleGroup labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallScheduleGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallScheduleGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallScheduleGroupMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallScheduleGroupArgs.__new__(FirewallScheduleGroupArgs)

            __props__.__dict__["color"] = color
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["fabric_object"] = fabric_object
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallScheduleGroup, __self__).__init__(
            'fortios:index/firewallScheduleGroup:FirewallScheduleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[int]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            fabric_object: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallScheduleGroupMemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallScheduleGroup':
        """
        Get an existing FirewallScheduleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fabric_object: Security Fabric global object setting. Valid values: `enable`, `disable`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallScheduleGroupMemberArgs']]]] members: Schedules added to the schedule group. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Schedule name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallScheduleGroupState.__new__(_FirewallScheduleGroupState)

        __props__.__dict__["color"] = color
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["fabric_object"] = fabric_object
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallScheduleGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[int]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> pulumi.Output[str]:
        """
        Security Fabric global object setting. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fabric_object")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.FirewallScheduleGroupMember']]:
        """
        Schedules added to the schedule group. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Schedule name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

