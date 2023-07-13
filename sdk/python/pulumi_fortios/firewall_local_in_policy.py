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

__all__ = ['FirewallLocalInPolicyArgs', 'FirewallLocalInPolicy']

@pulumi.input_type
class FirewallLocalInPolicyArgs:
    def __init__(__self__, *,
                 dstaddrs: pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyDstaddrArgs']]],
                 schedule: pulumi.Input[str],
                 srcaddrs: pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicySrcaddrArgs']]],
                 action: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dstaddr_negate: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ha_mgmt_intf_only: Optional[pulumi.Input[str]] = None,
                 intf: Optional[pulumi.Input[str]] = None,
                 policyid: Optional[pulumi.Input[int]] = None,
                 service_negate: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyServiceArgs']]]] = None,
                 srcaddr_negate: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 virtual_patch: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallLocalInPolicy resource.
        """
        pulumi.set(__self__, "dstaddrs", dstaddrs)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "srcaddrs", srcaddrs)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dstaddr_negate is not None:
            pulumi.set(__self__, "dstaddr_negate", dstaddr_negate)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ha_mgmt_intf_only is not None:
            pulumi.set(__self__, "ha_mgmt_intf_only", ha_mgmt_intf_only)
        if intf is not None:
            pulumi.set(__self__, "intf", intf)
        if policyid is not None:
            pulumi.set(__self__, "policyid", policyid)
        if service_negate is not None:
            pulumi.set(__self__, "service_negate", service_negate)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if srcaddr_negate is not None:
            pulumi.set(__self__, "srcaddr_negate", srcaddr_negate)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if virtual_patch is not None:
            pulumi.set(__self__, "virtual_patch", virtual_patch)

    @property
    @pulumi.getter
    def dstaddrs(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyDstaddrArgs']]]:
        return pulumi.get(self, "dstaddrs")

    @dstaddrs.setter
    def dstaddrs(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyDstaddrArgs']]]):
        pulumi.set(self, "dstaddrs", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[str]:
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def srcaddrs(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicySrcaddrArgs']]]:
        return pulumi.get(self, "srcaddrs")

    @srcaddrs.setter
    def srcaddrs(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicySrcaddrArgs']]]):
        pulumi.set(self, "srcaddrs", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="dstaddrNegate")
    def dstaddr_negate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dstaddr_negate")

    @dstaddr_negate.setter
    def dstaddr_negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dstaddr_negate", value)

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
    @pulumi.getter(name="haMgmtIntfOnly")
    def ha_mgmt_intf_only(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ha_mgmt_intf_only")

    @ha_mgmt_intf_only.setter
    def ha_mgmt_intf_only(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ha_mgmt_intf_only", value)

    @property
    @pulumi.getter
    def intf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "intf")

    @intf.setter
    def intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "intf", value)

    @property
    @pulumi.getter
    def policyid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "policyid")

    @policyid.setter
    def policyid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policyid", value)

    @property
    @pulumi.getter(name="serviceNegate")
    def service_negate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_negate")

    @service_negate.setter
    def service_negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_negate", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyServiceArgs']]]]:
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyServiceArgs']]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter(name="srcaddrNegate")
    def srcaddr_negate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "srcaddr_negate")

    @srcaddr_negate.setter
    def srcaddr_negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "srcaddr_negate", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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

    @property
    @pulumi.getter(name="virtualPatch")
    def virtual_patch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "virtual_patch")

    @virtual_patch.setter
    def virtual_patch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_patch", value)


@pulumi.input_type
class _FirewallLocalInPolicyState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dstaddr_negate: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyDstaddrArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ha_mgmt_intf_only: Optional[pulumi.Input[str]] = None,
                 intf: Optional[pulumi.Input[str]] = None,
                 policyid: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 service_negate: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyServiceArgs']]]] = None,
                 srcaddr_negate: Optional[pulumi.Input[str]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicySrcaddrArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 virtual_patch: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallLocalInPolicy resources.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dstaddr_negate is not None:
            pulumi.set(__self__, "dstaddr_negate", dstaddr_negate)
        if dstaddrs is not None:
            pulumi.set(__self__, "dstaddrs", dstaddrs)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if ha_mgmt_intf_only is not None:
            pulumi.set(__self__, "ha_mgmt_intf_only", ha_mgmt_intf_only)
        if intf is not None:
            pulumi.set(__self__, "intf", intf)
        if policyid is not None:
            pulumi.set(__self__, "policyid", policyid)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if service_negate is not None:
            pulumi.set(__self__, "service_negate", service_negate)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if srcaddr_negate is not None:
            pulumi.set(__self__, "srcaddr_negate", srcaddr_negate)
        if srcaddrs is not None:
            pulumi.set(__self__, "srcaddrs", srcaddrs)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if virtual_patch is not None:
            pulumi.set(__self__, "virtual_patch", virtual_patch)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter(name="dstaddrNegate")
    def dstaddr_negate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dstaddr_negate")

    @dstaddr_negate.setter
    def dstaddr_negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dstaddr_negate", value)

    @property
    @pulumi.getter
    def dstaddrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyDstaddrArgs']]]]:
        return pulumi.get(self, "dstaddrs")

    @dstaddrs.setter
    def dstaddrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyDstaddrArgs']]]]):
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
    @pulumi.getter(name="haMgmtIntfOnly")
    def ha_mgmt_intf_only(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ha_mgmt_intf_only")

    @ha_mgmt_intf_only.setter
    def ha_mgmt_intf_only(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ha_mgmt_intf_only", value)

    @property
    @pulumi.getter
    def intf(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "intf")

    @intf.setter
    def intf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "intf", value)

    @property
    @pulumi.getter
    def policyid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "policyid")

    @policyid.setter
    def policyid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "policyid", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="serviceNegate")
    def service_negate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_negate")

    @service_negate.setter
    def service_negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_negate", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyServiceArgs']]]]:
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicyServiceArgs']]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter(name="srcaddrNegate")
    def srcaddr_negate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "srcaddr_negate")

    @srcaddr_negate.setter
    def srcaddr_negate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "srcaddr_negate", value)

    @property
    @pulumi.getter
    def srcaddrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicySrcaddrArgs']]]]:
        return pulumi.get(self, "srcaddrs")

    @srcaddrs.setter
    def srcaddrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallLocalInPolicySrcaddrArgs']]]]):
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

    @property
    @pulumi.getter(name="virtualPatch")
    def virtual_patch(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "virtual_patch")

    @virtual_patch.setter
    def virtual_patch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_patch", value)


class FirewallLocalInPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dstaddr_negate: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicyDstaddrArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ha_mgmt_intf_only: Optional[pulumi.Input[str]] = None,
                 intf: Optional[pulumi.Input[str]] = None,
                 policyid: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 service_negate: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicyServiceArgs']]]]] = None,
                 srcaddr_negate: Optional[pulumi.Input[str]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicySrcaddrArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 virtual_patch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallLocalInPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallLocalInPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallLocalInPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallLocalInPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallLocalInPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dstaddr_negate: Optional[pulumi.Input[str]] = None,
                 dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicyDstaddrArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 ha_mgmt_intf_only: Optional[pulumi.Input[str]] = None,
                 intf: Optional[pulumi.Input[str]] = None,
                 policyid: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 service_negate: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicyServiceArgs']]]]] = None,
                 srcaddr_negate: Optional[pulumi.Input[str]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicySrcaddrArgs']]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 virtual_patch: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallLocalInPolicyArgs.__new__(FirewallLocalInPolicyArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["comments"] = comments
            __props__.__dict__["dstaddr_negate"] = dstaddr_negate
            if dstaddrs is None and not opts.urn:
                raise TypeError("Missing required property 'dstaddrs'")
            __props__.__dict__["dstaddrs"] = dstaddrs
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["ha_mgmt_intf_only"] = ha_mgmt_intf_only
            __props__.__dict__["intf"] = intf
            __props__.__dict__["policyid"] = policyid
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["service_negate"] = service_negate
            __props__.__dict__["services"] = services
            __props__.__dict__["srcaddr_negate"] = srcaddr_negate
            if srcaddrs is None and not opts.urn:
                raise TypeError("Missing required property 'srcaddrs'")
            __props__.__dict__["srcaddrs"] = srcaddrs
            __props__.__dict__["status"] = status
            __props__.__dict__["uuid"] = uuid
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["virtual_patch"] = virtual_patch
        super(FirewallLocalInPolicy, __self__).__init__(
            'fortios:index/firewallLocalInPolicy:FirewallLocalInPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            dstaddr_negate: Optional[pulumi.Input[str]] = None,
            dstaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicyDstaddrArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            ha_mgmt_intf_only: Optional[pulumi.Input[str]] = None,
            intf: Optional[pulumi.Input[str]] = None,
            policyid: Optional[pulumi.Input[int]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            service_negate: Optional[pulumi.Input[str]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicyServiceArgs']]]]] = None,
            srcaddr_negate: Optional[pulumi.Input[str]] = None,
            srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallLocalInPolicySrcaddrArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            virtual_patch: Optional[pulumi.Input[str]] = None) -> 'FirewallLocalInPolicy':
        """
        Get an existing FirewallLocalInPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallLocalInPolicyState.__new__(_FirewallLocalInPolicyState)

        __props__.__dict__["action"] = action
        __props__.__dict__["comments"] = comments
        __props__.__dict__["dstaddr_negate"] = dstaddr_negate
        __props__.__dict__["dstaddrs"] = dstaddrs
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["ha_mgmt_intf_only"] = ha_mgmt_intf_only
        __props__.__dict__["intf"] = intf
        __props__.__dict__["policyid"] = policyid
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["service_negate"] = service_negate
        __props__.__dict__["services"] = services
        __props__.__dict__["srcaddr_negate"] = srcaddr_negate
        __props__.__dict__["srcaddrs"] = srcaddrs
        __props__.__dict__["status"] = status
        __props__.__dict__["uuid"] = uuid
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["virtual_patch"] = virtual_patch
        return FirewallLocalInPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="dstaddrNegate")
    def dstaddr_negate(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dstaddr_negate")

    @property
    @pulumi.getter
    def dstaddrs(self) -> pulumi.Output[Sequence['outputs.FirewallLocalInPolicyDstaddr']]:
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
    @pulumi.getter(name="haMgmtIntfOnly")
    def ha_mgmt_intf_only(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ha_mgmt_intf_only")

    @property
    @pulumi.getter
    def intf(self) -> pulumi.Output[str]:
        return pulumi.get(self, "intf")

    @property
    @pulumi.getter
    def policyid(self) -> pulumi.Output[int]:
        return pulumi.get(self, "policyid")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="serviceNegate")
    def service_negate(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_negate")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallLocalInPolicyService']]]:
        return pulumi.get(self, "services")

    @property
    @pulumi.getter(name="srcaddrNegate")
    def srcaddr_negate(self) -> pulumi.Output[str]:
        return pulumi.get(self, "srcaddr_negate")

    @property
    @pulumi.getter
    def srcaddrs(self) -> pulumi.Output[Sequence['outputs.FirewallLocalInPolicySrcaddr']]:
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="virtualPatch")
    def virtual_patch(self) -> pulumi.Output[str]:
        return pulumi.get(self, "virtual_patch")

