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

__all__ = ['WirelessControllerWtpGroupArgs', 'WirelessControllerWtpGroup']

@pulumi.input_type
class WirelessControllerWtpGroupArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wtps: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerWtpGroupWtpArgs']]]] = None):
        """
        The set of arguments for constructing a WirelessControllerWtpGroup resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_type is not None:
            pulumi.set(__self__, "platform_type", platform_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wtps is not None:
            pulumi.set(__self__, "wtps", wtps)

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
    @pulumi.getter(name="platformType")
    def platform_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "platform_type")

    @platform_type.setter
    def platform_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform_type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def wtps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerWtpGroupWtpArgs']]]]:
        return pulumi.get(self, "wtps")

    @wtps.setter
    def wtps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerWtpGroupWtpArgs']]]]):
        pulumi.set(self, "wtps", value)


@pulumi.input_type
class _WirelessControllerWtpGroupState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wtps: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerWtpGroupWtpArgs']]]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerWtpGroup resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform_type is not None:
            pulumi.set(__self__, "platform_type", platform_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wtps is not None:
            pulumi.set(__self__, "wtps", wtps)

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
    @pulumi.getter(name="platformType")
    def platform_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "platform_type")

    @platform_type.setter
    def platform_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "platform_type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter
    def wtps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerWtpGroupWtpArgs']]]]:
        return pulumi.get(self, "wtps")

    @wtps.setter
    def wtps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerWtpGroupWtpArgs']]]]):
        pulumi.set(self, "wtps", value)


class WirelessControllerWtpGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wtps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerWtpGroupWtpArgs']]]]] = None,
                 __props__=None):
        """
        Create a WirelessControllerWtpGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerWtpGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WirelessControllerWtpGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WirelessControllerWtpGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerWtpGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 platform_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wtps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerWtpGroupWtpArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelessControllerWtpGroupArgs.__new__(WirelessControllerWtpGroupArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["name"] = name
            __props__.__dict__["platform_type"] = platform_type
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["wtps"] = wtps
        super(WirelessControllerWtpGroup, __self__).__init__(
            'fortios:index/wirelessControllerWtpGroup:WirelessControllerWtpGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            platform_type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            wtps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerWtpGroupWtpArgs']]]]] = None) -> 'WirelessControllerWtpGroup':
        """
        Get an existing WirelessControllerWtpGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerWtpGroupState.__new__(_WirelessControllerWtpGroupState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["name"] = name
        __props__.__dict__["platform_type"] = platform_type
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["wtps"] = wtps
        return WirelessControllerWtpGroup(resource_name, opts=opts, __props__=__props__)

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
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformType")
    def platform_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "platform_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def wtps(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerWtpGroupWtp']]]:
        return pulumi.get(self, "wtps")

