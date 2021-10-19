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

__all__ = ['RouterMulticast6Args', 'RouterMulticast6']

@pulumi.input_type
class RouterMulticast6Args:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticast6InterfaceArgs']]]] = None,
                 multicast_pmtu: Optional[pulumi.Input[str]] = None,
                 multicast_routing: Optional[pulumi.Input[str]] = None,
                 pim_sm_global: Optional[pulumi.Input['RouterMulticast6PimSmGlobalArgs']] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouterMulticast6 resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['RouterMulticast6InterfaceArgs']]] interfaces: Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        :param pulumi.Input[str] multicast_pmtu: Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] multicast_routing: Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        :param pulumi.Input['RouterMulticast6PimSmGlobalArgs'] pim_sm_global: PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if multicast_pmtu is not None:
            pulumi.set(__self__, "multicast_pmtu", multicast_pmtu)
        if multicast_routing is not None:
            pulumi.set(__self__, "multicast_routing", multicast_routing)
        if pim_sm_global is not None:
            pulumi.set(__self__, "pim_sm_global", pim_sm_global)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticast6InterfaceArgs']]]]:
        """
        Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticast6InterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="multicastPmtu")
    def multicast_pmtu(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_pmtu")

    @multicast_pmtu.setter
    def multicast_pmtu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_pmtu", value)

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_routing")

    @multicast_routing.setter
    def multicast_routing(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_routing", value)

    @property
    @pulumi.getter(name="pimSmGlobal")
    def pim_sm_global(self) -> Optional[pulumi.Input['RouterMulticast6PimSmGlobalArgs']]:
        """
        PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        """
        return pulumi.get(self, "pim_sm_global")

    @pim_sm_global.setter
    def pim_sm_global(self, value: Optional[pulumi.Input['RouterMulticast6PimSmGlobalArgs']]):
        pulumi.set(self, "pim_sm_global", value)

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
class _RouterMulticast6State:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticast6InterfaceArgs']]]] = None,
                 multicast_pmtu: Optional[pulumi.Input[str]] = None,
                 multicast_routing: Optional[pulumi.Input[str]] = None,
                 pim_sm_global: Optional[pulumi.Input['RouterMulticast6PimSmGlobalArgs']] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouterMulticast6 resources.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['RouterMulticast6InterfaceArgs']]] interfaces: Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        :param pulumi.Input[str] multicast_pmtu: Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] multicast_routing: Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        :param pulumi.Input['RouterMulticast6PimSmGlobalArgs'] pim_sm_global: PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if multicast_pmtu is not None:
            pulumi.set(__self__, "multicast_pmtu", multicast_pmtu)
        if multicast_routing is not None:
            pulumi.set(__self__, "multicast_routing", multicast_routing)
        if pim_sm_global is not None:
            pulumi.set(__self__, "pim_sm_global", pim_sm_global)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticast6InterfaceArgs']]]]:
        """
        Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterMulticast6InterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="multicastPmtu")
    def multicast_pmtu(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_pmtu")

    @multicast_pmtu.setter
    def multicast_pmtu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_pmtu", value)

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_routing")

    @multicast_routing.setter
    def multicast_routing(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "multicast_routing", value)

    @property
    @pulumi.getter(name="pimSmGlobal")
    def pim_sm_global(self) -> Optional[pulumi.Input['RouterMulticast6PimSmGlobalArgs']]:
        """
        PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        """
        return pulumi.get(self, "pim_sm_global")

    @pim_sm_global.setter
    def pim_sm_global(self, value: Optional[pulumi.Input['RouterMulticast6PimSmGlobalArgs']]):
        pulumi.set(self, "pim_sm_global", value)

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


class RouterMulticast6(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterMulticast6InterfaceArgs']]]]] = None,
                 multicast_pmtu: Optional[pulumi.Input[str]] = None,
                 multicast_routing: Optional[pulumi.Input[str]] = None,
                 pim_sm_global: Optional[pulumi.Input[pulumi.InputType['RouterMulticast6PimSmGlobalArgs']]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPv6 multicast.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.RouterMulticast6("trname",
            multicast_pmtu="disable",
            multicast_routing="disable",
            pim_sm_global=fortios.RouterMulticast6PimSmGlobalArgs(
                register_rate_limit=0,
            ))
        ```

        ## Import

        Router Multicast6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/routerMulticast6:RouterMulticast6 labelname RouterMulticast6
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterMulticast6InterfaceArgs']]]] interfaces: Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        :param pulumi.Input[str] multicast_pmtu: Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] multicast_routing: Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        :param pulumi.Input[pulumi.InputType['RouterMulticast6PimSmGlobalArgs']] pim_sm_global: PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RouterMulticast6Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPv6 multicast.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.RouterMulticast6("trname",
            multicast_pmtu="disable",
            multicast_routing="disable",
            pim_sm_global=fortios.RouterMulticast6PimSmGlobalArgs(
                register_rate_limit=0,
            ))
        ```

        ## Import

        Router Multicast6 can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/routerMulticast6:RouterMulticast6 labelname RouterMulticast6
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param RouterMulticast6Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterMulticast6Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterMulticast6InterfaceArgs']]]]] = None,
                 multicast_pmtu: Optional[pulumi.Input[str]] = None,
                 multicast_routing: Optional[pulumi.Input[str]] = None,
                 pim_sm_global: Optional[pulumi.Input[pulumi.InputType['RouterMulticast6PimSmGlobalArgs']]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouterMulticast6Args.__new__(RouterMulticast6Args)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["interfaces"] = interfaces
            __props__.__dict__["multicast_pmtu"] = multicast_pmtu
            __props__.__dict__["multicast_routing"] = multicast_routing
            __props__.__dict__["pim_sm_global"] = pim_sm_global
            __props__.__dict__["vdomparam"] = vdomparam
        super(RouterMulticast6, __self__).__init__(
            'fortios:index/routerMulticast6:RouterMulticast6',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterMulticast6InterfaceArgs']]]]] = None,
            multicast_pmtu: Optional[pulumi.Input[str]] = None,
            multicast_routing: Optional[pulumi.Input[str]] = None,
            pim_sm_global: Optional[pulumi.Input[pulumi.InputType['RouterMulticast6PimSmGlobalArgs']]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'RouterMulticast6':
        """
        Get an existing RouterMulticast6 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterMulticast6InterfaceArgs']]]] interfaces: Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        :param pulumi.Input[str] multicast_pmtu: Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] multicast_routing: Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        :param pulumi.Input[pulumi.InputType['RouterMulticast6PimSmGlobalArgs']] pim_sm_global: PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouterMulticast6State.__new__(_RouterMulticast6State)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["interfaces"] = interfaces
        __props__.__dict__["multicast_pmtu"] = multicast_pmtu
        __props__.__dict__["multicast_routing"] = multicast_routing
        __props__.__dict__["pim_sm_global"] = pim_sm_global
        __props__.__dict__["vdomparam"] = vdomparam
        return RouterMulticast6(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.RouterMulticast6Interface']]]:
        """
        Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="multicastPmtu")
    def multicast_pmtu(self) -> pulumi.Output[str]:
        """
        Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_pmtu")

    @property
    @pulumi.getter(name="multicastRouting")
    def multicast_routing(self) -> pulumi.Output[str]:
        """
        Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "multicast_routing")

    @property
    @pulumi.getter(name="pimSmGlobal")
    def pim_sm_global(self) -> pulumi.Output[Optional['outputs.RouterMulticast6PimSmGlobal']]:
        """
        PIM sparse-mode global settings. The structure of `pim_sm_global` block is documented below.
        """
        return pulumi.get(self, "pim_sm_global")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

