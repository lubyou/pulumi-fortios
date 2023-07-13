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

__all__ = ['FirewallDosPolicyArgs', 'FirewallDosPolicy']

@pulumi.input_type
class FirewallDosPolicyArgs:
    def __init__(__self__, *,
                 dstaddrs: pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyDstaddrArgs']]],
                 interface: pulumi.Input[str],
                 srcaddrs: pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicySrcaddrArgs']]],
                 anomalies: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyAnomalyArgs']]]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policyid: Optional[pulumi.Input[int]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyServiceArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallDosPolicy resource.
        """
        pulumi.set(__self__, "dstaddrs", dstaddrs)
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "srcaddrs", srcaddrs)
        if anomalies is not None:
            pulumi.set(__self__, "anomalies", anomalies)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policyid is not None:
            pulumi.set(__self__, "policyid", policyid)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def dstaddrs(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyDstaddrArgs']]]:
        return pulumi.get(self, "dstaddrs")

    @dstaddrs.setter
    def dstaddrs(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyDstaddrArgs']]]):
        pulumi.set(self, "dstaddrs", value)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def srcaddrs(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicySrcaddrArgs']]]:
        return pulumi.get(self, "srcaddrs")

    @srcaddrs.setter
    def srcaddrs(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicySrcaddrArgs']]]):
        pulumi.set(self, "srcaddrs", value)

    @property
    @pulumi.getter
    def anomalies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyAnomalyArgs']]]]:
        return pulumi.get(self, "anomalies")

    @anomalies.setter
    def anomalies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyAnomalyArgs']]]]):
        pulumi.set(self, "anomalies", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

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
    @pulumi.getter
    def policyid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "policyid")

    @policyid.setter
    def policyid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policyid", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyServiceArgs']]]]:
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyServiceArgs']]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallDosPolicyState:
    def __init__(__self__, *,
                 anomalies: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyAnomalyArgs']]]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyDstaddrArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policyid: Optional[pulumi.Input[int]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyServiceArgs']]]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicySrcaddrArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallDosPolicy resources.
        """
        if anomalies is not None:
            pulumi.set(__self__, "anomalies", anomalies)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dstaddrs is not None:
            pulumi.set(__self__, "dstaddrs", dstaddrs)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policyid is not None:
            pulumi.set(__self__, "policyid", policyid)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if srcaddrs is not None:
            pulumi.set(__self__, "srcaddrs", srcaddrs)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def anomalies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyAnomalyArgs']]]]:
        return pulumi.get(self, "anomalies")

    @anomalies.setter
    def anomalies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyAnomalyArgs']]]]):
        pulumi.set(self, "anomalies", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def dstaddrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyDstaddrArgs']]]]:
        return pulumi.get(self, "dstaddrs")

    @dstaddrs.setter
    def dstaddrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyDstaddrArgs']]]]):
        pulumi.set(self, "dstaddrs", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policyid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "policyid")

    @policyid.setter
    def policyid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policyid", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyServiceArgs']]]]:
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicyServiceArgs']]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def srcaddrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicySrcaddrArgs']]]]:
        return pulumi.get(self, "srcaddrs")

    @srcaddrs.setter
    def srcaddrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDosPolicySrcaddrArgs']]]]):
        pulumi.set(self, "srcaddrs", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallDosPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anomalies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyAnomalyArgs']]]]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyDstaddrArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policyid: Optional[pulumi.Input[int]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyServiceArgs']]]]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicySrcaddrArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallDosPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallDosPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallDosPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallDosPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallDosPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 anomalies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyAnomalyArgs']]]]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyDstaddrArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policyid: Optional[pulumi.Input[int]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyServiceArgs']]]]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicySrcaddrArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallDosPolicyArgs.__new__(FirewallDosPolicyArgs)

            __props__.__dict__["anomalies"] = anomalies
            __props__.__dict__["comments"] = comments
            if dstaddrs is None and not opts.urn:
                raise TypeError("Missing required property 'dstaddrs'")
            __props__.__dict__["dstaddrs"] = dstaddrs
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            __props__.__dict__["name"] = name
            __props__.__dict__["policyid"] = policyid
            __props__.__dict__["services"] = services
            if srcaddrs is None and not opts.urn:
                raise TypeError("Missing required property 'srcaddrs'")
            __props__.__dict__["srcaddrs"] = srcaddrs
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallDosPolicy, __self__).__init__(
            'fortios:index/firewallDosPolicy:FirewallDosPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            anomalies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyAnomalyArgs']]]]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyDstaddrArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policyid: Optional[pulumi.Input[int]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicyServiceArgs']]]]] = None,
            srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDosPolicySrcaddrArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallDosPolicy':
        """
        Get an existing FirewallDosPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallDosPolicyState.__new__(_FirewallDosPolicyState)

        __props__.__dict__["anomalies"] = anomalies
        __props__.__dict__["comments"] = comments
        __props__.__dict__["dstaddrs"] = dstaddrs
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["interface"] = interface
        __props__.__dict__["name"] = name
        __props__.__dict__["policyid"] = policyid
        __props__.__dict__["services"] = services
        __props__.__dict__["srcaddrs"] = srcaddrs
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallDosPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def anomalies(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallDosPolicyAnomaly']]]:
        return pulumi.get(self, "anomalies")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def dstaddrs(self) -> pulumi.Output[Sequence['outputs.FirewallDosPolicyDstaddr']]:
        return pulumi.get(self, "dstaddrs")

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
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policyid(self) -> pulumi.Output[int]:
        return pulumi.get(self, "policyid")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallDosPolicyService']]]:
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def srcaddrs(self) -> pulumi.Output[Sequence['outputs.FirewallDosPolicySrcaddr']]:
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

