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

__all__ = ['FirewallDecryptedTrafficMirrorArgs', 'FirewallDecryptedTrafficMirror']

@pulumi.input_type
class FirewallDecryptedTrafficMirrorArgs:
    def __init__(__self__, *,
                 dstmac: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDecryptedTrafficMirrorInterfaceArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 traffic_source: Optional[pulumi.Input[str]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallDecryptedTrafficMirror resource.
        :param pulumi.Input[str] dstmac: Set destination MAC address for mirrored traffic.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallDecryptedTrafficMirrorInterfaceArgs']]] interfaces: Decrypted traffic mirror interface The structure of `interface` block is documented below.
        :param pulumi.Input[str] name: Decrypted traffic mirror interface.
        :param pulumi.Input[str] traffic_source: Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
        :param pulumi.Input[str] traffic_type: Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dstmac is not None:
            pulumi.set(__self__, "dstmac", dstmac)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if traffic_source is not None:
            pulumi.set(__self__, "traffic_source", traffic_source)
        if traffic_type is not None:
            pulumi.set(__self__, "traffic_type", traffic_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def dstmac(self) -> Optional[pulumi.Input[str]]:
        """
        Set destination MAC address for mirrored traffic.
        """
        return pulumi.get(self, "dstmac")

    @dstmac.setter
    def dstmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dstmac", value)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDecryptedTrafficMirrorInterfaceArgs']]]]:
        """
        Decrypted traffic mirror interface The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDecryptedTrafficMirrorInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Decrypted traffic mirror interface.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="trafficSource")
    def traffic_source(self) -> Optional[pulumi.Input[str]]:
        """
        Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
        """
        return pulumi.get(self, "traffic_source")

    @traffic_source.setter
    def traffic_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_source", value)

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> Optional[pulumi.Input[str]]:
        """
        Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
        """
        return pulumi.get(self, "traffic_type")

    @traffic_type.setter
    def traffic_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_type", value)

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
class _FirewallDecryptedTrafficMirrorState:
    def __init__(__self__, *,
                 dstmac: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDecryptedTrafficMirrorInterfaceArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 traffic_source: Optional[pulumi.Input[str]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallDecryptedTrafficMirror resources.
        :param pulumi.Input[str] dstmac: Set destination MAC address for mirrored traffic.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input['FirewallDecryptedTrafficMirrorInterfaceArgs']]] interfaces: Decrypted traffic mirror interface The structure of `interface` block is documented below.
        :param pulumi.Input[str] name: Decrypted traffic mirror interface.
        :param pulumi.Input[str] traffic_source: Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
        :param pulumi.Input[str] traffic_type: Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dstmac is not None:
            pulumi.set(__self__, "dstmac", dstmac)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if interfaces is not None:
            pulumi.set(__self__, "interfaces", interfaces)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if traffic_source is not None:
            pulumi.set(__self__, "traffic_source", traffic_source)
        if traffic_type is not None:
            pulumi.set(__self__, "traffic_type", traffic_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def dstmac(self) -> Optional[pulumi.Input[str]]:
        """
        Set destination MAC address for mirrored traffic.
        """
        return pulumi.get(self, "dstmac")

    @dstmac.setter
    def dstmac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dstmac", value)

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
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDecryptedTrafficMirrorInterfaceArgs']]]]:
        """
        Decrypted traffic mirror interface The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallDecryptedTrafficMirrorInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Decrypted traffic mirror interface.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="trafficSource")
    def traffic_source(self) -> Optional[pulumi.Input[str]]:
        """
        Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
        """
        return pulumi.get(self, "traffic_source")

    @traffic_source.setter
    def traffic_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_source", value)

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> Optional[pulumi.Input[str]]:
        """
        Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
        """
        return pulumi.get(self, "traffic_type")

    @traffic_type.setter
    def traffic_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_type", value)

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


class FirewallDecryptedTrafficMirror(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dstmac: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDecryptedTrafficMirrorInterfaceArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 traffic_source: Optional[pulumi.Input[str]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure decrypted traffic mirror. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        Firewall DecryptedTrafficMirror can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallDecryptedTrafficMirror:FirewallDecryptedTrafficMirror labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallDecryptedTrafficMirror:FirewallDecryptedTrafficMirror labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dstmac: Set destination MAC address for mirrored traffic.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDecryptedTrafficMirrorInterfaceArgs']]]] interfaces: Decrypted traffic mirror interface The structure of `interface` block is documented below.
        :param pulumi.Input[str] name: Decrypted traffic mirror interface.
        :param pulumi.Input[str] traffic_source: Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
        :param pulumi.Input[str] traffic_type: Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallDecryptedTrafficMirrorArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure decrypted traffic mirror. Applies to FortiOS Version `>= 6.4.0`.

        ## Import

        Firewall DecryptedTrafficMirror can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallDecryptedTrafficMirror:FirewallDecryptedTrafficMirror labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallDecryptedTrafficMirror:FirewallDecryptedTrafficMirror labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallDecryptedTrafficMirrorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallDecryptedTrafficMirrorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dstmac: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDecryptedTrafficMirrorInterfaceArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 traffic_source: Optional[pulumi.Input[str]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallDecryptedTrafficMirrorArgs.__new__(FirewallDecryptedTrafficMirrorArgs)

            __props__.__dict__["dstmac"] = dstmac
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["interfaces"] = interfaces
            __props__.__dict__["name"] = name
            __props__.__dict__["traffic_source"] = traffic_source
            __props__.__dict__["traffic_type"] = traffic_type
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallDecryptedTrafficMirror, __self__).__init__(
            'fortios:index/firewallDecryptedTrafficMirror:FirewallDecryptedTrafficMirror',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dstmac: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDecryptedTrafficMirrorInterfaceArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            traffic_source: Optional[pulumi.Input[str]] = None,
            traffic_type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallDecryptedTrafficMirror':
        """
        Get an existing FirewallDecryptedTrafficMirror resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dstmac: Set destination MAC address for mirrored traffic.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallDecryptedTrafficMirrorInterfaceArgs']]]] interfaces: Decrypted traffic mirror interface The structure of `interface` block is documented below.
        :param pulumi.Input[str] name: Decrypted traffic mirror interface.
        :param pulumi.Input[str] traffic_source: Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
        :param pulumi.Input[str] traffic_type: Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallDecryptedTrafficMirrorState.__new__(_FirewallDecryptedTrafficMirrorState)

        __props__.__dict__["dstmac"] = dstmac
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["interfaces"] = interfaces
        __props__.__dict__["name"] = name
        __props__.__dict__["traffic_source"] = traffic_source
        __props__.__dict__["traffic_type"] = traffic_type
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallDecryptedTrafficMirror(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dstmac(self) -> pulumi.Output[str]:
        """
        Set destination MAC address for mirrored traffic.
        """
        return pulumi.get(self, "dstmac")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallDecryptedTrafficMirrorInterface']]]:
        """
        Decrypted traffic mirror interface The structure of `interface` block is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Decrypted traffic mirror interface.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="trafficSource")
    def traffic_source(self) -> pulumi.Output[str]:
        """
        Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
        """
        return pulumi.get(self, "traffic_source")

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> pulumi.Output[str]:
        """
        Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
        """
        return pulumi.get(self, "traffic_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

