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

__all__ = ['WirelessControllerHotspot20AnqpRoamingConsortiumArgs', 'WirelessControllerHotspot20AnqpRoamingConsortium']

@pulumi.input_type
class WirelessControllerHotspot20AnqpRoamingConsortiumArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oi_lists: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerHotspot20AnqpRoamingConsortium resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Roaming consortium name.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]] oi_lists: Organization identifier list. The structure of `oi_list` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if oi_lists is not None:
            pulumi.set(__self__, "oi_lists", oi_lists)
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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Roaming consortium name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="oiLists")
    def oi_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]]:
        """
        Organization identifier list. The structure of `oi_list` block is documented below.
        """
        return pulumi.get(self, "oi_lists")

    @oi_lists.setter
    def oi_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]]):
        pulumi.set(self, "oi_lists", value)

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
class _WirelessControllerHotspot20AnqpRoamingConsortiumState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oi_lists: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerHotspot20AnqpRoamingConsortium resources.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Roaming consortium name.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]] oi_lists: Organization identifier list. The structure of `oi_list` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if oi_lists is not None:
            pulumi.set(__self__, "oi_lists", oi_lists)
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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Roaming consortium name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="oiLists")
    def oi_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]]:
        """
        Organization identifier list. The structure of `oi_list` block is documented below.
        """
        return pulumi.get(self, "oi_lists")

    @oi_lists.setter
    def oi_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]]):
        pulumi.set(self, "oi_lists", value)

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


class WirelessControllerHotspot20AnqpRoamingConsortium(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oi_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure roaming consortium.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WirelessControllerHotspot20AnqpRoamingConsortium("trname")
        ```

        ## Import

        WirelessControllerHotspot20 AnqpRoamingConsortium can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Roaming consortium name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]] oi_lists: Organization identifier list. The structure of `oi_list` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerHotspot20AnqpRoamingConsortiumArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure roaming consortium.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.WirelessControllerHotspot20AnqpRoamingConsortium("trname")
        ```

        ## Import

        WirelessControllerHotspot20 AnqpRoamingConsortium can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelessControllerHotspot20AnqpRoamingConsortiumArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerHotspot20AnqpRoamingConsortiumArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oi_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]]] = None,
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
            __props__ = WirelessControllerHotspot20AnqpRoamingConsortiumArgs.__new__(WirelessControllerHotspot20AnqpRoamingConsortiumArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["oi_lists"] = oi_lists
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerHotspot20AnqpRoamingConsortium, __self__).__init__(
            'fortios:index/wirelessControllerHotspot20AnqpRoamingConsortium:WirelessControllerHotspot20AnqpRoamingConsortium',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            oi_lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerHotspot20AnqpRoamingConsortium':
        """
        Get an existing WirelessControllerHotspot20AnqpRoamingConsortium resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] name: Roaming consortium name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerHotspot20AnqpRoamingConsortiumOiListArgs']]]] oi_lists: Organization identifier list. The structure of `oi_list` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerHotspot20AnqpRoamingConsortiumState.__new__(_WirelessControllerHotspot20AnqpRoamingConsortiumState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["oi_lists"] = oi_lists
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerHotspot20AnqpRoamingConsortium(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Roaming consortium name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="oiLists")
    def oi_lists(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerHotspot20AnqpRoamingConsortiumOiList']]]:
        """
        Organization identifier list. The structure of `oi_list` block is documented below.
        """
        return pulumi.get(self, "oi_lists")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

