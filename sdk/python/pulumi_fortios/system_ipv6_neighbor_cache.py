# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemIpv6NeighborCacheArgs', 'SystemIpv6NeighborCache']

@pulumi.input_type
class SystemIpv6NeighborCacheArgs:
    def __init__(__self__, *,
                 fosid: pulumi.Input[int],
                 interface: pulumi.Input[str],
                 ipv6: pulumi.Input[str],
                 mac: pulumi.Input[str],
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemIpv6NeighborCache resource.
        :param pulumi.Input[int] fosid: Unique integer ID of the entry.
        :param pulumi.Input[str] interface: Select the associated interface name from available options.
        :param pulumi.Input[str] ipv6: IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        :param pulumi.Input[str] mac: MAC address (format: xx:xx:xx:xx:xx:xx).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "fosid", fosid)
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "ipv6", ipv6)
        pulumi.set(__self__, "mac", mac)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Input[int]:
        """
        Unique integer ID of the entry.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: pulumi.Input[int]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        """
        Select the associated interface name from available options.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def ipv6(self) -> pulumi.Input[str]:
        """
        IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: pulumi.Input[str]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Input[str]:
        """
        MAC address (format: xx:xx:xx:xx:xx:xx).
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: pulumi.Input[str]):
        pulumi.set(self, "mac", value)

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
class _SystemIpv6NeighborCacheState:
    def __init__(__self__, *,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemIpv6NeighborCache resources.
        :param pulumi.Input[int] fosid: Unique integer ID of the entry.
        :param pulumi.Input[str] interface: Select the associated interface name from available options.
        :param pulumi.Input[str] ipv6: IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        :param pulumi.Input[str] mac: MAC address (format: xx:xx:xx:xx:xx:xx).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Unique integer ID of the entry.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Select the associated interface name from available options.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        """
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address (format: xx:xx:xx:xx:xx:xx).
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

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


class SystemIpv6NeighborCache(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure IPv6 neighbor cache table.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemIpv6NeighborCache("trname",
            fosid=1,
            interface="port2",
            ipv6="fe80::b11a:5ae3:198:ba1c",
            mac="00:00:00:00:00:00")
        ```

        ## Import

        System Ipv6NeighborCache can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemIpv6NeighborCache:SystemIpv6NeighborCache labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] fosid: Unique integer ID of the entry.
        :param pulumi.Input[str] interface: Select the associated interface name from available options.
        :param pulumi.Input[str] ipv6: IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        :param pulumi.Input[str] mac: MAC address (format: xx:xx:xx:xx:xx:xx).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemIpv6NeighborCacheArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure IPv6 neighbor cache table.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemIpv6NeighborCache("trname",
            fosid=1,
            interface="port2",
            ipv6="fe80::b11a:5ae3:198:ba1c",
            mac="00:00:00:00:00:00")
        ```

        ## Import

        System Ipv6NeighborCache can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemIpv6NeighborCache:SystemIpv6NeighborCache labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemIpv6NeighborCacheArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemIpv6NeighborCacheArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemIpv6NeighborCacheArgs.__new__(SystemIpv6NeighborCacheArgs)

            if fosid is None and not opts.urn:
                raise TypeError("Missing required property 'fosid'")
            __props__.__dict__["fosid"] = fosid
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            if ipv6 is None and not opts.urn:
                raise TypeError("Missing required property 'ipv6'")
            __props__.__dict__["ipv6"] = ipv6
            if mac is None and not opts.urn:
                raise TypeError("Missing required property 'mac'")
            __props__.__dict__["mac"] = mac
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemIpv6NeighborCache, __self__).__init__(
            'fortios:index/systemIpv6NeighborCache:SystemIpv6NeighborCache',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            ipv6: Optional[pulumi.Input[str]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemIpv6NeighborCache':
        """
        Get an existing SystemIpv6NeighborCache resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] fosid: Unique integer ID of the entry.
        :param pulumi.Input[str] interface: Select the associated interface name from available options.
        :param pulumi.Input[str] ipv6: IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        :param pulumi.Input[str] mac: MAC address (format: xx:xx:xx:xx:xx:xx).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemIpv6NeighborCacheState.__new__(_SystemIpv6NeighborCacheState)

        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["interface"] = interface
        __props__.__dict__["ipv6"] = ipv6
        __props__.__dict__["mac"] = mac
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemIpv6NeighborCache(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Unique integer ID of the entry.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Select the associated interface name from available options.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def ipv6(self) -> pulumi.Output[str]:
        """
        IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        MAC address (format: xx:xx:xx:xx:xx:xx).
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
