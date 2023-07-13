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

__all__ = ['SystemSessionTtlArgs', 'SystemSessionTtl']

@pulumi.input_type
class SystemSessionTtlArgs:
    def __init__(__self__, *,
                 default: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input['SystemSessionTtlPortArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemSessionTtl resource.
        """
        if default is not None:
            pulumi.set(__self__, "default", default)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default", value)

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
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemSessionTtlPortArgs']]]]:
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemSessionTtlPortArgs']]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemSessionTtlState:
    def __init__(__self__, *,
                 default: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input['SystemSessionTtlPortArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemSessionTtl resources.
        """
        if default is not None:
            pulumi.set(__self__, "default", default)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default", value)

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
    def ports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemSessionTtlPortArgs']]]]:
        return pulumi.get(self, "ports")

    @ports.setter
    def ports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemSessionTtlPortArgs']]]]):
        pulumi.set(self, "ports", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemSessionTtl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemSessionTtlPortArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemSessionTtl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemSessionTtlArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemSessionTtl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemSessionTtlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemSessionTtlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemSessionTtlPortArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemSessionTtlArgs.__new__(SystemSessionTtlArgs)

            __props__.__dict__["default"] = default
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["ports"] = ports
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemSessionTtl, __self__).__init__(
            'fortios:index/systemSessionTtl:SystemSessionTtl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            ports: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemSessionTtlPortArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemSessionTtl':
        """
        Get an existing SystemSessionTtl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemSessionTtlState.__new__(_SystemSessionTtlState)

        __props__.__dict__["default"] = default
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["ports"] = ports
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemSessionTtl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def default(self) -> pulumi.Output[str]:
        return pulumi.get(self, "default")

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
    def ports(self) -> pulumi.Output[Optional[Sequence['outputs.SystemSessionTtlPort']]]:
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

