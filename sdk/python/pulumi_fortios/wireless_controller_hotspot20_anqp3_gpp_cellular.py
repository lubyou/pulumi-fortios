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

__all__ = ['WirelessControllerHotspot20Anqp3GppCellularArgs', 'WirelessControllerHotspot20Anqp3GppCellular']

@pulumi.input_type
class WirelessControllerHotspot20Anqp3GppCellularArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mcc_mnc_lists: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerHotspot20Anqp3GppCellular resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if mcc_mnc_lists is not None:
            pulumi.set(__self__, "mcc_mnc_lists", mcc_mnc_lists)
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
    @pulumi.getter(name="mccMncLists")
    def mcc_mnc_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]]:
        return pulumi.get(self, "mcc_mnc_lists")

    @mcc_mnc_lists.setter
    def mcc_mnc_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]]):
        pulumi.set(self, "mcc_mnc_lists", value)

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
class _WirelessControllerHotspot20Anqp3GppCellularState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mcc_mnc_lists: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerHotspot20Anqp3GppCellular resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if mcc_mnc_lists is not None:
            pulumi.set(__self__, "mcc_mnc_lists", mcc_mnc_lists)
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
    @pulumi.getter(name="mccMncLists")
    def mcc_mnc_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]]:
        return pulumi.get(self, "mcc_mnc_lists")

    @mcc_mnc_lists.setter
    def mcc_mnc_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]]):
        pulumi.set(self, "mcc_mnc_lists", value)

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


class WirelessControllerHotspot20Anqp3GppCellular(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mcc_mnc_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a WirelessControllerHotspot20Anqp3GppCellular resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerHotspot20Anqp3GppCellularArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WirelessControllerHotspot20Anqp3GppCellular resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WirelessControllerHotspot20Anqp3GppCellularArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerHotspot20Anqp3GppCellularArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 mcc_mnc_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelessControllerHotspot20Anqp3GppCellularArgs.__new__(WirelessControllerHotspot20Anqp3GppCellularArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["mcc_mnc_lists"] = mcc_mnc_lists
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerHotspot20Anqp3GppCellular, __self__).__init__(
            'fortios:index/wirelessControllerHotspot20Anqp3GppCellular:WirelessControllerHotspot20Anqp3GppCellular',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            mcc_mnc_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerHotspot20Anqp3GppCellularMccMncListArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerHotspot20Anqp3GppCellular':
        """
        Get an existing WirelessControllerHotspot20Anqp3GppCellular resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerHotspot20Anqp3GppCellularState.__new__(_WirelessControllerHotspot20Anqp3GppCellularState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["mcc_mnc_lists"] = mcc_mnc_lists
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerHotspot20Anqp3GppCellular(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="mccMncLists")
    def mcc_mnc_lists(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerHotspot20Anqp3GppCellularMccMncList']]]:
        return pulumi.get(self, "mcc_mnc_lists")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

