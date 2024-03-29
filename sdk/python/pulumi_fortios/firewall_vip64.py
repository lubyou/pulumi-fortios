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

__all__ = ['FirewallVip64Args', 'FirewallVip64']

@pulumi.input_type
class FirewallVip64Args:
    def __init__(__self__, *,
                 extip: pulumi.Input[str],
                 mappedip: pulumi.Input[str],
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extport: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 mappedport: Optional[pulumi.Input[str]] = None,
                 monitors: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64MonitorArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portforward: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 realservers: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64RealserverArgs']]]] = None,
                 server_type: Optional[pulumi.Input[str]] = None,
                 src_filters: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64SrcFilterArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallVip64 resource.
        """
        pulumi.set(__self__, "extip", extip)
        pulumi.set(__self__, "mappedip", mappedip)
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if extport is not None:
            pulumi.set(__self__, "extport", extport)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ldb_method is not None:
            pulumi.set(__self__, "ldb_method", ldb_method)
        if mappedport is not None:
            pulumi.set(__self__, "mappedport", mappedport)
        if monitors is not None:
            pulumi.set(__self__, "monitors", monitors)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if portforward is not None:
            pulumi.set(__self__, "portforward", portforward)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if realservers is not None:
            pulumi.set(__self__, "realservers", realservers)
        if server_type is not None:
            pulumi.set(__self__, "server_type", server_type)
        if src_filters is not None:
            pulumi.set(__self__, "src_filters", src_filters)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def extip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "extip")

    @extip.setter
    def extip(self, value: pulumi.Input[str]):
        pulumi.set(self, "extip", value)

    @property
    @pulumi.getter
    def mappedip(self) -> pulumi.Input[str]:
        return pulumi.get(self, "mappedip")

    @mappedip.setter
    def mappedip(self, value: pulumi.Input[str]):
        pulumi.set(self, "mappedip", value)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

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
    def extport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "extport")

    @extport.setter
    def extport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extport", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

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
    def mappedport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mappedport")

    @mappedport.setter
    def mappedport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mappedport", value)

    @property
    @pulumi.getter
    def monitors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64MonitorArgs']]]]:
        return pulumi.get(self, "monitors")

    @monitors.setter
    def monitors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64MonitorArgs']]]]):
        pulumi.set(self, "monitors", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def portforward(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "portforward")

    @portforward.setter
    def portforward(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portforward", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def realservers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64RealserverArgs']]]]:
        return pulumi.get(self, "realservers")

    @realservers.setter
    def realservers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64RealserverArgs']]]]):
        pulumi.set(self, "realservers", value)

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_type")

    @server_type.setter
    def server_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_type", value)

    @property
    @pulumi.getter(name="srcFilters")
    def src_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64SrcFilterArgs']]]]:
        return pulumi.get(self, "src_filters")

    @src_filters.setter
    def src_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64SrcFilterArgs']]]]):
        pulumi.set(self, "src_filters", value)

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


@pulumi.input_type
class _FirewallVip64State:
    def __init__(__self__, *,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extip: Optional[pulumi.Input[str]] = None,
                 extport: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 mappedip: Optional[pulumi.Input[str]] = None,
                 mappedport: Optional[pulumi.Input[str]] = None,
                 monitors: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64MonitorArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portforward: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 realservers: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64RealserverArgs']]]] = None,
                 server_type: Optional[pulumi.Input[str]] = None,
                 src_filters: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64SrcFilterArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallVip64 resources.
        """
        if arp_reply is not None:
            pulumi.set(__self__, "arp_reply", arp_reply)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if extip is not None:
            pulumi.set(__self__, "extip", extip)
        if extport is not None:
            pulumi.set(__self__, "extport", extport)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ldb_method is not None:
            pulumi.set(__self__, "ldb_method", ldb_method)
        if mappedip is not None:
            pulumi.set(__self__, "mappedip", mappedip)
        if mappedport is not None:
            pulumi.set(__self__, "mappedport", mappedport)
        if monitors is not None:
            pulumi.set(__self__, "monitors", monitors)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if portforward is not None:
            pulumi.set(__self__, "portforward", portforward)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if realservers is not None:
            pulumi.set(__self__, "realservers", realservers)
        if server_type is not None:
            pulumi.set(__self__, "server_type", server_type)
        if src_filters is not None:
            pulumi.set(__self__, "src_filters", src_filters)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "arp_reply")

    @arp_reply.setter
    def arp_reply(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arp_reply", value)

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
    def extip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "extip")

    @extip.setter
    def extip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extip", value)

    @property
    @pulumi.getter
    def extport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "extport")

    @extport.setter
    def extport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "extport", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

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
    def mappedip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mappedip")

    @mappedip.setter
    def mappedip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mappedip", value)

    @property
    @pulumi.getter
    def mappedport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mappedport")

    @mappedport.setter
    def mappedport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mappedport", value)

    @property
    @pulumi.getter
    def monitors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64MonitorArgs']]]]:
        return pulumi.get(self, "monitors")

    @monitors.setter
    def monitors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64MonitorArgs']]]]):
        pulumi.set(self, "monitors", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def portforward(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "portforward")

    @portforward.setter
    def portforward(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portforward", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def realservers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64RealserverArgs']]]]:
        return pulumi.get(self, "realservers")

    @realservers.setter
    def realservers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64RealserverArgs']]]]):
        pulumi.set(self, "realservers", value)

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_type")

    @server_type.setter
    def server_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_type", value)

    @property
    @pulumi.getter(name="srcFilters")
    def src_filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64SrcFilterArgs']]]]:
        return pulumi.get(self, "src_filters")

    @src_filters.setter
    def src_filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallVip64SrcFilterArgs']]]]):
        pulumi.set(self, "src_filters", value)

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


class FirewallVip64(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extip: Optional[pulumi.Input[str]] = None,
                 extport: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 mappedip: Optional[pulumi.Input[str]] = None,
                 mappedport: Optional[pulumi.Input[str]] = None,
                 monitors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64MonitorArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portforward: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 realservers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64RealserverArgs']]]]] = None,
                 server_type: Optional[pulumi.Input[str]] = None,
                 src_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64SrcFilterArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallVip64 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallVip64Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallVip64 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallVip64Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallVip64Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arp_reply: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 extip: Optional[pulumi.Input[str]] = None,
                 extport: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ldb_method: Optional[pulumi.Input[str]] = None,
                 mappedip: Optional[pulumi.Input[str]] = None,
                 mappedport: Optional[pulumi.Input[str]] = None,
                 monitors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64MonitorArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portforward: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 realservers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64RealserverArgs']]]]] = None,
                 server_type: Optional[pulumi.Input[str]] = None,
                 src_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64SrcFilterArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallVip64Args.__new__(FirewallVip64Args)

            __props__.__dict__["arp_reply"] = arp_reply
            __props__.__dict__["color"] = color
            __props__.__dict__["comment"] = comment
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if extip is None and not opts.urn:
                raise TypeError("Missing required property 'extip'")
            __props__.__dict__["extip"] = extip
            __props__.__dict__["extport"] = extport
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["ldb_method"] = ldb_method
            if mappedip is None and not opts.urn:
                raise TypeError("Missing required property 'mappedip'")
            __props__.__dict__["mappedip"] = mappedip
            __props__.__dict__["mappedport"] = mappedport
            __props__.__dict__["monitors"] = monitors
            __props__.__dict__["name"] = name
            __props__.__dict__["portforward"] = portforward
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["realservers"] = realservers
            __props__.__dict__["server_type"] = server_type
            __props__.__dict__["src_filters"] = src_filters
            __props__.__dict__["type"] = type
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallVip64, __self__).__init__(
            'fortios:index/firewallVip64:FirewallVip64',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arp_reply: Optional[pulumi.Input[str]] = None,
            color: Optional[pulumi.Input[int]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            extip: Optional[pulumi.Input[str]] = None,
            extport: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            ldb_method: Optional[pulumi.Input[str]] = None,
            mappedip: Optional[pulumi.Input[str]] = None,
            mappedport: Optional[pulumi.Input[str]] = None,
            monitors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64MonitorArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            portforward: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            realservers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64RealserverArgs']]]]] = None,
            server_type: Optional[pulumi.Input[str]] = None,
            src_filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallVip64SrcFilterArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallVip64':
        """
        Get an existing FirewallVip64 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallVip64State.__new__(_FirewallVip64State)

        __props__.__dict__["arp_reply"] = arp_reply
        __props__.__dict__["color"] = color
        __props__.__dict__["comment"] = comment
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["extip"] = extip
        __props__.__dict__["extport"] = extport
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["ldb_method"] = ldb_method
        __props__.__dict__["mappedip"] = mappedip
        __props__.__dict__["mappedport"] = mappedport
        __props__.__dict__["monitors"] = monitors
        __props__.__dict__["name"] = name
        __props__.__dict__["portforward"] = portforward
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["realservers"] = realservers
        __props__.__dict__["server_type"] = server_type
        __props__.__dict__["src_filters"] = src_filters
        __props__.__dict__["type"] = type
        __props__.__dict__["uuid"] = uuid
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallVip64(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="arpReply")
    def arp_reply(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arp_reply")

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
    def extip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "extip")

    @property
    @pulumi.getter
    def extport(self) -> pulumi.Output[str]:
        return pulumi.get(self, "extport")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        return pulumi.get(self, "fosid")

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
    def mappedip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mappedip")

    @property
    @pulumi.getter
    def mappedport(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mappedport")

    @property
    @pulumi.getter
    def monitors(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallVip64Monitor']]]:
        return pulumi.get(self, "monitors")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def portforward(self) -> pulumi.Output[str]:
        return pulumi.get(self, "portforward")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def realservers(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallVip64Realserver']]]:
        return pulumi.get(self, "realservers")

    @property
    @pulumi.getter(name="serverType")
    def server_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "server_type")

    @property
    @pulumi.getter(name="srcFilters")
    def src_filters(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallVip64SrcFilter']]]:
        return pulumi.get(self, "src_filters")

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

