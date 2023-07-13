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

__all__ = ['IcapServerGroupArgs', 'IcapServerGroup']

@pulumi.input_type
class IcapServerGroupArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input['IcapServerGroupServerListArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IcapServerGroup resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ldb_method is not None:
            pulumi.set(__self__, "ldb_method", ldb_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_lists is not None:
            pulumi.set(__self__, "server_lists", server_lists)
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
    @pulumi.getter(name="ldbMethod")
    def ldb_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldb_method")

    @ldb_method.setter
    def ldb_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldb_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IcapServerGroupServerListArgs']]]]:
        return pulumi.get(self, "server_lists")

    @server_lists.setter
    def server_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IcapServerGroupServerListArgs']]]]):
        pulumi.set(self, "server_lists", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _IcapServerGroupState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input['IcapServerGroupServerListArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IcapServerGroup resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ldb_method is not None:
            pulumi.set(__self__, "ldb_method", ldb_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if server_lists is not None:
            pulumi.set(__self__, "server_lists", server_lists)
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
    @pulumi.getter(name="ldbMethod")
    def ldb_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldb_method")

    @ldb_method.setter
    def ldb_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldb_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IcapServerGroupServerListArgs']]]]:
        return pulumi.get(self, "server_lists")

    @server_lists.setter
    def server_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IcapServerGroupServerListArgs']]]]):
        pulumi.set(self, "server_lists", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class IcapServerGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IcapServerGroupServerListArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a IcapServerGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IcapServerGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IcapServerGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IcapServerGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IcapServerGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IcapServerGroupServerListArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IcapServerGroupArgs.__new__(IcapServerGroupArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["ldb_method"] = ldb_method
            __props__.__dict__["name"] = name
            __props__.__dict__["server_lists"] = server_lists
            __props__.__dict__["vdomparam"] = vdomparam
        super(IcapServerGroup, __self__).__init__(
            'fortios:index/icapServerGroup:IcapServerGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            ldb_method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IcapServerGroupServerListArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'IcapServerGroup':
        """
        Get an existing IcapServerGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IcapServerGroupState.__new__(_IcapServerGroupState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["ldb_method"] = ldb_method
        __props__.__dict__["name"] = name
        __props__.__dict__["server_lists"] = server_lists
        __props__.__dict__["vdomparam"] = vdomparam
        return IcapServerGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="ldbMethod")
    def ldb_method(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ldb_method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> pulumi.Output[Optional[Sequence['outputs.IcapServerGroupServerList']]]:
        return pulumi.get(self, "server_lists")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

