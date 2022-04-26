# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FirewallTtlPolicyArgs', 'FirewallTtlPolicy']

@pulumi.input_type
class FirewallTtlPolicyArgs:
    def __init__(__self__, *,
                 fosid: pulumi.Input[int],
                 schedule: pulumi.Input[str],
                 services: pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicyServiceArgs']]],
                 srcaddrs: pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicySrcaddrArgs']]],
                 srcintf: pulumi.Input[str],
                 ttl: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallTtlPolicy resource.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] schedule: Schedule object from available options.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicyServiceArgs']]] services: Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicySrcaddrArgs']]] srcaddrs: Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        :param pulumi.Input[str] srcintf: Source interface name from available interfaces.
        :param pulumi.Input[str] ttl: Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        :param pulumi.Input[str] action: Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] status: Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "fosid", fosid)
        pulumi.set(__self__, "schedule", schedule)
        pulumi.set(__self__, "services", services)
        pulumi.set(__self__, "srcaddrs", srcaddrs)
        pulumi.set(__self__, "srcintf", srcintf)
        pulumi.set(__self__, "ttl", ttl)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Input[int]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: pulumi.Input[int]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[str]:
        """
        Schedule object from available options.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def services(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicyServiceArgs']]]:
        """
        Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicyServiceArgs']]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def srcaddrs(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicySrcaddrArgs']]]:
        """
        Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        """
        return pulumi.get(self, "srcaddrs")

    @srcaddrs.setter
    def srcaddrs(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicySrcaddrArgs']]]):
        pulumi.set(self, "srcaddrs", value)

    @property
    @pulumi.getter
    def srcintf(self) -> pulumi.Input[str]:
        """
        Source interface name from available interfaces.
        """
        return pulumi.get(self, "srcintf")

    @srcintf.setter
    def srcintf(self, value: pulumi.Input[str]):
        pulumi.set(self, "srcintf", value)

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Input[str]:
        """
        Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: pulumi.Input[str]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _FirewallTtlPolicyState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicyServiceArgs']]]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicySrcaddrArgs']]]] = None,
                 srcintf: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallTtlPolicy resources.
        :param pulumi.Input[str] action: Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] schedule: Schedule object from available options.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicyServiceArgs']]] services: Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicySrcaddrArgs']]] srcaddrs: Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        :param pulumi.Input[str] srcintf: Source interface name from available interfaces.
        :param pulumi.Input[str] status: Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ttl: Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if srcaddrs is not None:
            pulumi.set(__self__, "srcaddrs", srcaddrs)
        if srcintf is not None:
            pulumi.set(__self__, "srcintf", srcintf)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        Schedule object from available options.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicyServiceArgs']]]]:
        """
        Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicyServiceArgs']]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def srcaddrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicySrcaddrArgs']]]]:
        """
        Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        """
        return pulumi.get(self, "srcaddrs")

    @srcaddrs.setter
    def srcaddrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTtlPolicySrcaddrArgs']]]]):
        pulumi.set(self, "srcaddrs", value)

    @property
    @pulumi.getter
    def srcintf(self) -> Optional[pulumi.Input[str]]:
        """
        Source interface name from available interfaces.
        """
        return pulumi.get(self, "srcintf")

    @srcintf.setter
    def srcintf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "srcintf", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallTtlPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicyServiceArgs']]]]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicySrcaddrArgs']]]]] = None,
                 srcintf: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure TTL policies.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallTtlPolicy("trname",
            action="accept",
            fosid=1,
            schedule="always",
            services=[fortios.FirewallTtlPolicyServiceArgs(
                name="ALL",
            )],
            srcaddrs=[fortios.FirewallTtlPolicySrcaddrArgs(
                name="all",
            )],
            srcintf="port3",
            status="enable",
            ttl="23")
        ```

        ## Import

        Firewall TtlPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallTtlPolicy:FirewallTtlPolicy labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallTtlPolicy:FirewallTtlPolicy labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] schedule: Schedule object from available options.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicyServiceArgs']]]] services: Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicySrcaddrArgs']]]] srcaddrs: Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        :param pulumi.Input[str] srcintf: Source interface name from available interfaces.
        :param pulumi.Input[str] status: Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ttl: Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallTtlPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure TTL policies.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallTtlPolicy("trname",
            action="accept",
            fosid=1,
            schedule="always",
            services=[fortios.FirewallTtlPolicyServiceArgs(
                name="ALL",
            )],
            srcaddrs=[fortios.FirewallTtlPolicySrcaddrArgs(
                name="all",
            )],
            srcintf="port3",
            status="enable",
            ttl="23")
        ```

        ## Import

        Firewall TtlPolicy can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallTtlPolicy:FirewallTtlPolicy labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallTtlPolicy:FirewallTtlPolicy labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallTtlPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallTtlPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicyServiceArgs']]]]] = None,
                 srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicySrcaddrArgs']]]]] = None,
                 srcintf: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallTtlPolicyArgs.__new__(FirewallTtlPolicyArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if fosid is None and not opts.urn:
                raise TypeError("Missing required property 'fosid'")
            __props__.__dict__["fosid"] = fosid
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            if services is None and not opts.urn:
                raise TypeError("Missing required property 'services'")
            __props__.__dict__["services"] = services
            if srcaddrs is None and not opts.urn:
                raise TypeError("Missing required property 'srcaddrs'")
            __props__.__dict__["srcaddrs"] = srcaddrs
            if srcintf is None and not opts.urn:
                raise TypeError("Missing required property 'srcintf'")
            __props__.__dict__["srcintf"] = srcintf
            __props__.__dict__["status"] = status
            if ttl is None and not opts.urn:
                raise TypeError("Missing required property 'ttl'")
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallTtlPolicy, __self__).__init__(
            'fortios:index/firewallTtlPolicy:FirewallTtlPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicyServiceArgs']]]]] = None,
            srcaddrs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicySrcaddrArgs']]]]] = None,
            srcintf: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallTtlPolicy':
        """
        Get an existing FirewallTtlPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] schedule: Schedule object from available options.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicyServiceArgs']]]] services: Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTtlPolicySrcaddrArgs']]]] srcaddrs: Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        :param pulumi.Input[str] srcintf: Source interface name from available interfaces.
        :param pulumi.Input[str] status: Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ttl: Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallTtlPolicyState.__new__(_FirewallTtlPolicyState)

        __props__.__dict__["action"] = action
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["services"] = services
        __props__.__dict__["srcaddrs"] = srcaddrs
        __props__.__dict__["srcintf"] = srcintf
        __props__.__dict__["status"] = status
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallTtlPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Action to be performed on traffic matching this policy (default = deny). Valid values: `accept`, `deny`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        Schedule object from available options.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Sequence['outputs.FirewallTtlPolicyService']]:
        """
        Service object(s) from available options. Separate multiple names with a space. The structure of `service` block is documented below.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def srcaddrs(self) -> pulumi.Output[Sequence['outputs.FirewallTtlPolicySrcaddr']]:
        """
        Source address object(s) from available options. Separate multiple names with a space. The structure of `srcaddr` block is documented below.
        """
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter
    def srcintf(self) -> pulumi.Output[str]:
        """
        Source interface name from available interfaces.
        """
        return pulumi.get(self, "srcintf")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable this TTL policy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[str]:
        """
        Value/range to match against the packet's Time to Live value (format: ttl[ - ttl_high], 1 - 255).
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

