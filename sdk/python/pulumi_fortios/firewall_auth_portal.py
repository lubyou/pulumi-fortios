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

__all__ = ['FirewallAuthPortalArgs', 'FirewallAuthPortal']

@pulumi.input_type
class FirewallAuthPortalArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAuthPortalGroupArgs']]]] = None,
                 identity_based_route: Optional[pulumi.Input[str]] = None,
                 portal_addr: Optional[pulumi.Input[str]] = None,
                 portal_addr6: Optional[pulumi.Input[str]] = None,
                 proxy_auth: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallAuthPortal resource.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if identity_based_route is not None:
            pulumi.set(__self__, "identity_based_route", identity_based_route)
        if portal_addr is not None:
            pulumi.set(__self__, "portal_addr", portal_addr)
        if portal_addr6 is not None:
            pulumi.set(__self__, "portal_addr6", portal_addr6)
        if proxy_auth is not None:
            pulumi.set(__self__, "proxy_auth", proxy_auth)
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
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAuthPortalGroupArgs']]]]:
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAuthPortalGroupArgs']]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="identityBasedRoute")
    def identity_based_route(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "identity_based_route")

    @identity_based_route.setter
    def identity_based_route(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_based_route", value)

    @property
    @pulumi.getter(name="portalAddr")
    def portal_addr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "portal_addr")

    @portal_addr.setter
    def portal_addr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portal_addr", value)

    @property
    @pulumi.getter(name="portalAddr6")
    def portal_addr6(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "portal_addr6")

    @portal_addr6.setter
    def portal_addr6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portal_addr6", value)

    @property
    @pulumi.getter(name="proxyAuth")
    def proxy_auth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "proxy_auth")

    @proxy_auth.setter
    def proxy_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_auth", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallAuthPortalState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAuthPortalGroupArgs']]]] = None,
                 identity_based_route: Optional[pulumi.Input[str]] = None,
                 portal_addr: Optional[pulumi.Input[str]] = None,
                 portal_addr6: Optional[pulumi.Input[str]] = None,
                 proxy_auth: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallAuthPortal resources.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if identity_based_route is not None:
            pulumi.set(__self__, "identity_based_route", identity_based_route)
        if portal_addr is not None:
            pulumi.set(__self__, "portal_addr", portal_addr)
        if portal_addr6 is not None:
            pulumi.set(__self__, "portal_addr6", portal_addr6)
        if proxy_auth is not None:
            pulumi.set(__self__, "proxy_auth", proxy_auth)
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
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAuthPortalGroupArgs']]]]:
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallAuthPortalGroupArgs']]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="identityBasedRoute")
    def identity_based_route(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "identity_based_route")

    @identity_based_route.setter
    def identity_based_route(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_based_route", value)

    @property
    @pulumi.getter(name="portalAddr")
    def portal_addr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "portal_addr")

    @portal_addr.setter
    def portal_addr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portal_addr", value)

    @property
    @pulumi.getter(name="portalAddr6")
    def portal_addr6(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "portal_addr6")

    @portal_addr6.setter
    def portal_addr6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portal_addr6", value)

    @property
    @pulumi.getter(name="proxyAuth")
    def proxy_auth(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "proxy_auth")

    @proxy_auth.setter
    def proxy_auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_auth", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallAuthPortal(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAuthPortalGroupArgs']]]]] = None,
                 identity_based_route: Optional[pulumi.Input[str]] = None,
                 portal_addr: Optional[pulumi.Input[str]] = None,
                 portal_addr6: Optional[pulumi.Input[str]] = None,
                 proxy_auth: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallAuthPortal resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallAuthPortalArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallAuthPortal resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallAuthPortalArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallAuthPortalArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAuthPortalGroupArgs']]]]] = None,
                 identity_based_route: Optional[pulumi.Input[str]] = None,
                 portal_addr: Optional[pulumi.Input[str]] = None,
                 portal_addr6: Optional[pulumi.Input[str]] = None,
                 proxy_auth: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallAuthPortalArgs.__new__(FirewallAuthPortalArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["groups"] = groups
            __props__.__dict__["identity_based_route"] = identity_based_route
            __props__.__dict__["portal_addr"] = portal_addr
            __props__.__dict__["portal_addr6"] = portal_addr6
            __props__.__dict__["proxy_auth"] = proxy_auth
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallAuthPortal, __self__).__init__(
            'fortios:index/firewallAuthPortal:FirewallAuthPortal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallAuthPortalGroupArgs']]]]] = None,
            identity_based_route: Optional[pulumi.Input[str]] = None,
            portal_addr: Optional[pulumi.Input[str]] = None,
            portal_addr6: Optional[pulumi.Input[str]] = None,
            proxy_auth: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallAuthPortal':
        """
        Get an existing FirewallAuthPortal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallAuthPortalState.__new__(_FirewallAuthPortalState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["groups"] = groups
        __props__.__dict__["identity_based_route"] = identity_based_route
        __props__.__dict__["portal_addr"] = portal_addr
        __props__.__dict__["portal_addr6"] = portal_addr6
        __props__.__dict__["proxy_auth"] = proxy_auth
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallAuthPortal(resource_name, opts=opts, __props__=__props__)

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
    def groups(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallAuthPortalGroup']]]:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="identityBasedRoute")
    def identity_based_route(self) -> pulumi.Output[str]:
        return pulumi.get(self, "identity_based_route")

    @property
    @pulumi.getter(name="portalAddr")
    def portal_addr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "portal_addr")

    @property
    @pulumi.getter(name="portalAddr6")
    def portal_addr6(self) -> pulumi.Output[str]:
        return pulumi.get(self, "portal_addr6")

    @property
    @pulumi.getter(name="proxyAuth")
    def proxy_auth(self) -> pulumi.Output[str]:
        return pulumi.get(self, "proxy_auth")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

