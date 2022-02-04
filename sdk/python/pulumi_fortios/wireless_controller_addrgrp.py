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

__all__ = ['WirelessControllerAddrgrpArgs', 'WirelessControllerAddrgrp']

@pulumi.input_type
class WirelessControllerAddrgrpArgs:
    def __init__(__self__, *,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAddrgrpAddressArgs']]]] = None,
                 default_policy: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerAddrgrp resource.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerAddrgrpAddressArgs']]] addresses: Manually selected group of addresses. The structure of `addresses` block is documented below.
        :param pulumi.Input[str] default_policy: Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fosid: ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if addresses is not None:
            pulumi.set(__self__, "addresses", addresses)
        if default_policy is not None:
            pulumi.set(__self__, "default_policy", default_policy)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAddrgrpAddressArgs']]]]:
        """
        Manually selected group of addresses. The structure of `addresses` block is documented below.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAddrgrpAddressArgs']]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="defaultPolicy")
    def default_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "default_policy")

    @default_policy.setter
    def default_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_policy", value)

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
    def fosid(self) -> Optional[pulumi.Input[str]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fosid", value)

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
class _WirelessControllerAddrgrpState:
    def __init__(__self__, *,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAddrgrpAddressArgs']]]] = None,
                 default_policy: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerAddrgrp resources.
        :param pulumi.Input[Sequence[pulumi.Input['WirelessControllerAddrgrpAddressArgs']]] addresses: Manually selected group of addresses. The structure of `addresses` block is documented below.
        :param pulumi.Input[str] default_policy: Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fosid: ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if addresses is not None:
            pulumi.set(__self__, "addresses", addresses)
        if default_policy is not None:
            pulumi.set(__self__, "default_policy", default_policy)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAddrgrpAddressArgs']]]]:
        """
        Manually selected group of addresses. The structure of `addresses` block is documented below.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerAddrgrpAddressArgs']]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="defaultPolicy")
    def default_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "default_policy")

    @default_policy.setter
    def default_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_policy", value)

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
    def fosid(self) -> Optional[pulumi.Input[str]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fosid", value)

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


class WirelessControllerAddrgrp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAddrgrpAddressArgs']]]]] = None,
                 default_policy: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure the MAC address group. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        WirelessController Addrgrp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAddrgrpAddressArgs']]]] addresses: Manually selected group of addresses. The structure of `addresses` block is documented below.
        :param pulumi.Input[str] default_policy: Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fosid: ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerAddrgrpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure the MAC address group. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        WirelessController Addrgrp can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelessControllerAddrgrpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerAddrgrpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAddrgrpAddressArgs']]]]] = None,
                 default_policy: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
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
            __props__ = WirelessControllerAddrgrpArgs.__new__(WirelessControllerAddrgrpArgs)

            __props__.__dict__["addresses"] = addresses
            __props__.__dict__["default_policy"] = default_policy
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerAddrgrp, __self__).__init__(
            'fortios:index/wirelessControllerAddrgrp:WirelessControllerAddrgrp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            addresses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAddrgrpAddressArgs']]]]] = None,
            default_policy: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerAddrgrp':
        """
        Get an existing WirelessControllerAddrgrp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerAddrgrpAddressArgs']]]] addresses: Manually selected group of addresses. The structure of `addresses` block is documented below.
        :param pulumi.Input[str] default_policy: Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[str] fosid: ID.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerAddrgrpState.__new__(_WirelessControllerAddrgrpState)

        __props__.__dict__["addresses"] = addresses
        __props__.__dict__["default_policy"] = default_policy
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerAddrgrp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerAddrgrpAddress']]]:
        """
        Manually selected group of addresses. The structure of `addresses` block is documented below.
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="defaultPolicy")
    def default_policy(self) -> pulumi.Output[str]:
        """
        Allow or block the clients with MAC addresses that are not in the group. Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "default_policy")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[str]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

