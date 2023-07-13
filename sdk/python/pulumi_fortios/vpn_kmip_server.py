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

__all__ = ['VpnKmipServerArgs', 'VpnKmipServer']

@pulumi.input_type
class VpnKmipServerArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_identity_check: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input['VpnKmipServerServerListArgs']]]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpnKmipServer resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if server_identity_check is not None:
            pulumi.set(__self__, "server_identity_check", server_identity_check)
        if server_lists is not None:
            pulumi.set(__self__, "server_lists", server_lists)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if username is not None:
            pulumi.set(__self__, "username", username)
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
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="serverIdentityCheck")
    def server_identity_check(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_identity_check")

    @server_identity_check.setter
    def server_identity_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_identity_check", value)

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnKmipServerServerListArgs']]]]:
        return pulumi.get(self, "server_lists")

    @server_lists.setter
    def server_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnKmipServerServerListArgs']]]]):
        pulumi.set(self, "server_lists", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _VpnKmipServerState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_identity_check: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input['VpnKmipServerServerListArgs']]]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpnKmipServer resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if server_identity_check is not None:
            pulumi.set(__self__, "server_identity_check", server_identity_check)
        if server_lists is not None:
            pulumi.set(__self__, "server_lists", server_lists)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if username is not None:
            pulumi.set(__self__, "username", username)
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
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="serverIdentityCheck")
    def server_identity_check(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server_identity_check")

    @server_identity_check.setter
    def server_identity_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_identity_check", value)

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpnKmipServerServerListArgs']]]]:
        return pulumi.get(self, "server_lists")

    @server_lists.setter
    def server_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpnKmipServerServerListArgs']]]]):
        pulumi.set(self, "server_lists", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class VpnKmipServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_identity_check: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnKmipServerServerListArgs']]]]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VpnKmipServer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VpnKmipServerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VpnKmipServer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VpnKmipServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpnKmipServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 server_identity_check: Optional[pulumi.Input[str]] = None,
                 server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnKmipServerServerListArgs']]]]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpnKmipServerArgs.__new__(VpnKmipServerArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["interface"] = interface
            __props__.__dict__["interface_select_method"] = interface_select_method
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = password
            __props__.__dict__["server_identity_check"] = server_identity_check
            __props__.__dict__["server_lists"] = server_lists
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
            __props__.__dict__["username"] = username
            __props__.__dict__["vdomparam"] = vdomparam
        super(VpnKmipServer, __self__).__init__(
            'fortios:index/vpnKmipServer:VpnKmipServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            interface_select_method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            server_identity_check: Optional[pulumi.Input[str]] = None,
            server_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpnKmipServerServerListArgs']]]]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'VpnKmipServer':
        """
        Get an existing VpnKmipServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpnKmipServerState.__new__(_VpnKmipServerState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["interface"] = interface
        __props__.__dict__["interface_select_method"] = interface_select_method
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["server_identity_check"] = server_identity_check
        __props__.__dict__["server_lists"] = server_lists
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
        __props__.__dict__["username"] = username
        __props__.__dict__["vdomparam"] = vdomparam
        return VpnKmipServer(resource_name, opts=opts, __props__=__props__)

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
    def interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface_select_method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="serverIdentityCheck")
    def server_identity_check(self) -> pulumi.Output[str]:
        return pulumi.get(self, "server_identity_check")

    @property
    @pulumi.getter(name="serverLists")
    def server_lists(self) -> pulumi.Output[Optional[Sequence['outputs.VpnKmipServerServerList']]]:
        return pulumi.get(self, "server_lists")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ssl_min_proto_version")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

