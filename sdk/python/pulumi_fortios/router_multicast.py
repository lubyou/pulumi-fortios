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

__all__ = ['RouterMulticastArgs', 'RouterMulticast']

@pulumi.input_type
class RouterMulticastArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticastInterfaceArgs']]]] = None,
                 multicast_routing: Optional[pulumi.Input[str]] = None,
                 pim_sm_global: Optional[pulumi.Input['RouterMulticastPimSmGlobalArgs']] = None,
                 route_limit: Optional[pulumi.Input[int]] = None,
                 route_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouterMulticast resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if multicast_routing is not None:
            pulumi.set(__self__, "multicast_routing", multicast_routing)
        if pim_sm_global is not None:
            pulumi.set(__self__, "pim_sm_global", pim_sm_global)
        if route_limit is not None:
            pulumi.set(__self__, "route_limit", route_limit)
        if route_threshold is not None:
            pulumi.set(__self__, "route_threshold", route_threshold)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticastInterfaceArgs']]]]:
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticastInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "multicast_routing")

    @multicast_routing.setter
    def multicast_routing(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_routing", value)

    @property
    @pulumi.getter(name="pimSmGlobal")
    def pim_sm_global(self) -> Optional[pulumi.Input['RouterMulticastPimSmGlobalArgs']]:
        return pulumi.get(self, "pim_sm_global")

    @pim_sm_global.setter
    def pim_sm_global(self, value: Optional[pulumi.Input['RouterMulticastPimSmGlobalArgs']]):
        pulumi.set(self, "pim_sm_global", value)

    @property
    @pulumi.getter(name="routeLimit")
    def route_limit(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "route_limit")

    @route_limit.setter
    def route_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "route_limit", value)

    @property
    @pulumi.getter(name="routeThreshold")
    def route_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "route_threshold")

    @route_threshold.setter
    def route_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "route_threshold", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _RouterMulticastState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticastInterfaceArgs']]]] = None,
                 multicast_routing: Optional[pulumi.Input[str]] = None,
                 pim_sm_global: Optional[pulumi.Input['RouterMulticastPimSmGlobalArgs']] = None,
                 route_limit: Optional[pulumi.Input[int]] = None,
                 route_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouterMulticast resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if multicast_routing is not None:
            pulumi.set(__self__, "multicast_routing", multicast_routing)
        if pim_sm_global is not None:
            pulumi.set(__self__, "pim_sm_global", pim_sm_global)
        if route_limit is not None:
            pulumi.set(__self__, "route_limit", route_limit)
        if route_threshold is not None:
            pulumi.set(__self__, "route_threshold", route_threshold)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticastInterfaceArgs']]]]:
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticastInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "multicast_routing")

    @multicast_routing.setter
    def multicast_routing(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_routing", value)

    @property
    @pulumi.getter(name="pimSmGlobal")
    def pim_sm_global(self) -> Optional[pulumi.Input['RouterMulticastPimSmGlobalArgs']]:
        return pulumi.get(self, "pim_sm_global")

    @pim_sm_global.setter
    def pim_sm_global(self, value: Optional[pulumi.Input['RouterMulticastPimSmGlobalArgs']]):
        pulumi.set(self, "pim_sm_global", value)

    @property
    @pulumi.getter(name="routeLimit")
    def route_limit(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "route_limit")

    @route_limit.setter
    def route_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "route_limit", value)

    @property
    @pulumi.getter(name="routeThreshold")
    def route_threshold(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "route_threshold")

    @route_threshold.setter
    def route_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "route_threshold", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class RouterMulticast(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterMulticastInterfaceArgs']]]]] = None,
                 multicast_routing: Optional[pulumi.Input[str]] = None,
                 pim_sm_global: Optional[pulumi.Input[pulumi.InputType['RouterMulticastPimSmGlobalArgs']]] = None,
                 route_limit: Optional[pulumi.Input[int]] = None,
                 route_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a RouterMulticast resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RouterMulticastArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RouterMulticast resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RouterMulticastArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterMulticastArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterMulticastInterfaceArgs']]]]] = None,
                 multicast_routing: Optional[pulumi.Input[str]] = None,
                 pim_sm_global: Optional[pulumi.Input[pulumi.InputType['RouterMulticastPimSmGlobalArgs']]] = None,
                 route_limit: Optional[pulumi.Input[int]] = None,
                 route_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouterMulticastArgs.__new__(RouterMulticastArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["interfaces"] = interfaces
            __props__.__dict__["multicast_routing"] = multicast_routing
            __props__.__dict__["pim_sm_global"] = pim_sm_global
            __props__.__dict__["route_limit"] = route_limit
            __props__.__dict__["route_threshold"] = route_threshold
            __props__.__dict__["vdomparam"] = vdomparam
        super(RouterMulticast, __self__).__init__(
            'fortios:index/routerMulticast:RouterMulticast',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterMulticastInterfaceArgs']]]]] = None,
            multicast_routing: Optional[pulumi.Input[str]] = None,
            pim_sm_global: Optional[pulumi.Input[pulumi.InputType['RouterMulticastPimSmGlobalArgs']]] = None,
            route_limit: Optional[pulumi.Input[int]] = None,
            route_threshold: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'RouterMulticast':
        """
        Get an existing RouterMulticast resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouterMulticastState.__new__(_RouterMulticastState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["interfaces"] = interfaces
        __props__.__dict__["multicast_routing"] = multicast_routing
        __props__.__dict__["pim_sm_global"] = pim_sm_global
        __props__.__dict__["route_limit"] = route_limit
        __props__.__dict__["route_threshold"] = route_threshold
        __props__.__dict__["vdomparam"] = vdomparam
        return RouterMulticast(resource_name, opts=opts, __props__=__props__)

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
    def interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.RouterMulticastInterface']]]:
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> pulumi.Output[str]:
        return pulumi.get(self, "multicast_routing")

    @property
    @pulumi.getter(name="pimSmGlobal")
    def pim_sm_global(self) -> pulumi.Output['outputs.RouterMulticastPimSmGlobal']:
        return pulumi.get(self, "pim_sm_global")

    @property
    @pulumi.getter(name="routeLimit")
    def route_limit(self) -> pulumi.Output[int]:
        return pulumi.get(self, "route_limit")

    @property
    @pulumi.getter(name="routeThreshold")
    def route_threshold(self) -> pulumi.Output[int]:
        return pulumi.get(self, "route_threshold")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

