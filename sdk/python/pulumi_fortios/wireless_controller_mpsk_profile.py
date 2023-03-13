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

__all__ = ['WirelessControllerMpskProfileArgs', 'WirelessControllerMpskProfile']

@pulumi.input_type
class WirelessControllerMpskProfileArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
                 mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerMpskProfileMpskGroupArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerMpskProfile resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if mpsk_concurrent_clients is not None:
            pulumi.set(__self__, "mpsk_concurrent_clients", mpsk_concurrent_clients)
        if mpsk_groups is not None:
            pulumi.set(__self__, "mpsk_groups", mpsk_groups)
        if name is not None:
            pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="mpskConcurrentClients")
    def mpsk_concurrent_clients(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "mpsk_concurrent_clients")

    @mpsk_concurrent_clients.setter
    def mpsk_concurrent_clients(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mpsk_concurrent_clients", value)

    @property
    @pulumi.getter(name="mpskGroups")
    def mpsk_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerMpskProfileMpskGroupArgs']]]]:
        return pulumi.get(self, "mpsk_groups")

    @mpsk_groups.setter
    def mpsk_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerMpskProfileMpskGroupArgs']]]]):
        pulumi.set(self, "mpsk_groups", value)

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


@pulumi.input_type
class _WirelessControllerMpskProfileState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
                 mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerMpskProfileMpskGroupArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerMpskProfile resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if mpsk_concurrent_clients is not None:
            pulumi.set(__self__, "mpsk_concurrent_clients", mpsk_concurrent_clients)
        if mpsk_groups is not None:
            pulumi.set(__self__, "mpsk_groups", mpsk_groups)
        if name is not None:
            pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="mpskConcurrentClients")
    def mpsk_concurrent_clients(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "mpsk_concurrent_clients")

    @mpsk_concurrent_clients.setter
    def mpsk_concurrent_clients(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mpsk_concurrent_clients", value)

    @property
    @pulumi.getter(name="mpskGroups")
    def mpsk_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerMpskProfileMpskGroupArgs']]]]:
        return pulumi.get(self, "mpsk_groups")

    @mpsk_groups.setter
    def mpsk_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerMpskProfileMpskGroupArgs']]]]):
        pulumi.set(self, "mpsk_groups", value)

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


class WirelessControllerMpskProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
                 mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerMpskProfileMpskGroupArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a WirelessControllerMpskProfile resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerMpskProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WirelessControllerMpskProfile resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WirelessControllerMpskProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerMpskProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
                 mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerMpskProfileMpskGroupArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelessControllerMpskProfileArgs.__new__(WirelessControllerMpskProfileArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["mpsk_concurrent_clients"] = mpsk_concurrent_clients
            __props__.__dict__["mpsk_groups"] = mpsk_groups
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerMpskProfile, __self__).__init__(
            'fortios:index/wirelessControllerMpskProfile:WirelessControllerMpskProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            mpsk_concurrent_clients: Optional[pulumi.Input[int]] = None,
            mpsk_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerMpskProfileMpskGroupArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerMpskProfile':
        """
        Get an existing WirelessControllerMpskProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerMpskProfileState.__new__(_WirelessControllerMpskProfileState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["mpsk_concurrent_clients"] = mpsk_concurrent_clients
        __props__.__dict__["mpsk_groups"] = mpsk_groups
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerMpskProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="mpskConcurrentClients")
    def mpsk_concurrent_clients(self) -> pulumi.Output[int]:
        return pulumi.get(self, "mpsk_concurrent_clients")

    @property
    @pulumi.getter(name="mpskGroups")
    def mpsk_groups(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerMpskProfileMpskGroup']]]:
        return pulumi.get(self, "mpsk_groups")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

