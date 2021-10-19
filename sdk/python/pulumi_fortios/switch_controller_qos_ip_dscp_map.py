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

__all__ = ['SwitchControllerQosIpDscpMapArgs', 'SwitchControllerQosIpDscpMap']

@pulumi.input_type
class SwitchControllerQosIpDscpMapArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 maps: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosIpDscpMapMapArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerQosIpDscpMap resource.
        :param pulumi.Input[str] description: Description of the ip-dscp map name.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosIpDscpMapMapArgs']]] maps: Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        :param pulumi.Input[str] name: Dscp mapping entry name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if maps is not None:
            pulumi.set(__self__, "maps", maps)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the ip-dscp map name.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    def maps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosIpDscpMapMapArgs']]]]:
        """
        Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        """
        return pulumi.get(self, "maps")

    @maps.setter
    def maps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosIpDscpMapMapArgs']]]]):
        pulumi.set(self, "maps", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Dscp mapping entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _SwitchControllerQosIpDscpMapState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 maps: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosIpDscpMapMapArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerQosIpDscpMap resources.
        :param pulumi.Input[str] description: Description of the ip-dscp map name.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosIpDscpMapMapArgs']]] maps: Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        :param pulumi.Input[str] name: Dscp mapping entry name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if maps is not None:
            pulumi.set(__self__, "maps", maps)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the ip-dscp map name.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    def maps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosIpDscpMapMapArgs']]]]:
        """
        Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        """
        return pulumi.get(self, "maps")

    @maps.setter
    def maps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchControllerQosIpDscpMapMapArgs']]]]):
        pulumi.set(self, "maps", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Dscp mapping entry name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class SwitchControllerQosIpDscpMap(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 maps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosIpDscpMapMapArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FortiSwitch QoS IP precedence/DSCP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SwitchControllerQosIpDscpMap("trname",
            description="DEIW",
            maps=[fortios.SwitchControllerQosIpDscpMapMapArgs(
                cos_queue=3,
                diffserv="CS0 CS1 AF11",
                name="1",
            )])
        ```

        ## Import

        SwitchControllerQos IpDscpMap can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/switchControllerQosIpDscpMap:SwitchControllerQosIpDscpMap labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the ip-dscp map name.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosIpDscpMapMapArgs']]]] maps: Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        :param pulumi.Input[str] name: Dscp mapping entry name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerQosIpDscpMapArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FortiSwitch QoS IP precedence/DSCP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SwitchControllerQosIpDscpMap("trname",
            description="DEIW",
            maps=[fortios.SwitchControllerQosIpDscpMapMapArgs(
                cos_queue=3,
                diffserv="CS0 CS1 AF11",
                name="1",
            )])
        ```

        ## Import

        SwitchControllerQos IpDscpMap can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/switchControllerQosIpDscpMap:SwitchControllerQosIpDscpMap labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchControllerQosIpDscpMapArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerQosIpDscpMapArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 maps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosIpDscpMapMapArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = SwitchControllerQosIpDscpMapArgs.__new__(SwitchControllerQosIpDscpMapArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["maps"] = maps
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerQosIpDscpMap, __self__).__init__(
            'fortios:index/switchControllerQosIpDscpMap:SwitchControllerQosIpDscpMap',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            maps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosIpDscpMapMapArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerQosIpDscpMap':
        """
        Get an existing SwitchControllerQosIpDscpMap resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the ip-dscp map name.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchControllerQosIpDscpMapMapArgs']]]] maps: Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        :param pulumi.Input[str] name: Dscp mapping entry name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerQosIpDscpMapState.__new__(_SwitchControllerQosIpDscpMapState)

        __props__.__dict__["description"] = description
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["maps"] = maps
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerQosIpDscpMap(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the ip-dscp map name.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def maps(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchControllerQosIpDscpMapMap']]]:
        """
        Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
        """
        return pulumi.get(self, "maps")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Dscp mapping entry name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

