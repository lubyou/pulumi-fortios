# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WirelessControllerBleProfileArgs', 'WirelessControllerBleProfile']

@pulumi.input_type
class WirelessControllerBleProfileArgs:
    def __init__(__self__, *,
                 advertising: Optional[pulumi.Input[str]] = None,
                 beacon_interval: Optional[pulumi.Input[int]] = None,
                 ble_scanning: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 eddystone_instance: Optional[pulumi.Input[str]] = None,
                 eddystone_namespace: Optional[pulumi.Input[str]] = None,
                 eddystone_url: Optional[pulumi.Input[str]] = None,
                 eddystone_url_encode_hex: Optional[pulumi.Input[str]] = None,
                 ibeacon_uuid: Optional[pulumi.Input[str]] = None,
                 major_id: Optional[pulumi.Input[int]] = None,
                 minor_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 txpower: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerBleProfile resource.
        :param pulumi.Input[str] advertising: Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
        :param pulumi.Input[int] beacon_interval: Beacon interval (default = 100 msec).
        :param pulumi.Input[str] ble_scanning: Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] eddystone_instance: Eddystone instance ID.
        :param pulumi.Input[str] eddystone_namespace: Eddystone namespace ID.
        :param pulumi.Input[str] eddystone_url: Eddystone URL.
        :param pulumi.Input[str] eddystone_url_encode_hex: Eddystone encoded URL hexadecimal string
        :param pulumi.Input[str] ibeacon_uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[int] major_id: Major ID.
        :param pulumi.Input[int] minor_id: Minor ID.
        :param pulumi.Input[str] name: Bluetooth Low Energy profile name.
        :param pulumi.Input[str] txpower: Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if advertising is not None:
            pulumi.set(__self__, "advertising", advertising)
        if beacon_interval is not None:
            pulumi.set(__self__, "beacon_interval", beacon_interval)
        if ble_scanning is not None:
            pulumi.set(__self__, "ble_scanning", ble_scanning)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if eddystone_instance is not None:
            pulumi.set(__self__, "eddystone_instance", eddystone_instance)
        if eddystone_namespace is not None:
            pulumi.set(__self__, "eddystone_namespace", eddystone_namespace)
        if eddystone_url is not None:
            pulumi.set(__self__, "eddystone_url", eddystone_url)
        if eddystone_url_encode_hex is not None:
            pulumi.set(__self__, "eddystone_url_encode_hex", eddystone_url_encode_hex)
        if ibeacon_uuid is not None:
            pulumi.set(__self__, "ibeacon_uuid", ibeacon_uuid)
        if major_id is not None:
            pulumi.set(__self__, "major_id", major_id)
        if minor_id is not None:
            pulumi.set(__self__, "minor_id", minor_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if txpower is not None:
            pulumi.set(__self__, "txpower", txpower)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def advertising(self) -> Optional[pulumi.Input[str]]:
        """
        Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
        """
        return pulumi.get(self, "advertising")

    @advertising.setter
    def advertising(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "advertising", value)

    @property
    @pulumi.getter(name="beaconInterval")
    def beacon_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Beacon interval (default = 100 msec).
        """
        return pulumi.get(self, "beacon_interval")

    @beacon_interval.setter
    def beacon_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "beacon_interval", value)

    @property
    @pulumi.getter(name="bleScanning")
    def ble_scanning(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ble_scanning")

    @ble_scanning.setter
    def ble_scanning(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ble_scanning", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="eddystoneInstance")
    def eddystone_instance(self) -> Optional[pulumi.Input[str]]:
        """
        Eddystone instance ID.
        """
        return pulumi.get(self, "eddystone_instance")

    @eddystone_instance.setter
    def eddystone_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eddystone_instance", value)

    @property
    @pulumi.getter(name="eddystoneNamespace")
    def eddystone_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Eddystone namespace ID.
        """
        return pulumi.get(self, "eddystone_namespace")

    @eddystone_namespace.setter
    def eddystone_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eddystone_namespace", value)

    @property
    @pulumi.getter(name="eddystoneUrl")
    def eddystone_url(self) -> Optional[pulumi.Input[str]]:
        """
        Eddystone URL.
        """
        return pulumi.get(self, "eddystone_url")

    @eddystone_url.setter
    def eddystone_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eddystone_url", value)

    @property
    @pulumi.getter(name="eddystoneUrlEncodeHex")
    def eddystone_url_encode_hex(self) -> Optional[pulumi.Input[str]]:
        """
        Eddystone encoded URL hexadecimal string
        """
        return pulumi.get(self, "eddystone_url_encode_hex")

    @eddystone_url_encode_hex.setter
    def eddystone_url_encode_hex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eddystone_url_encode_hex", value)

    @property
    @pulumi.getter(name="ibeaconUuid")
    def ibeacon_uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "ibeacon_uuid")

    @ibeacon_uuid.setter
    def ibeacon_uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ibeacon_uuid", value)

    @property
    @pulumi.getter(name="majorId")
    def major_id(self) -> Optional[pulumi.Input[int]]:
        """
        Major ID.
        """
        return pulumi.get(self, "major_id")

    @major_id.setter
    def major_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "major_id", value)

    @property
    @pulumi.getter(name="minorId")
    def minor_id(self) -> Optional[pulumi.Input[int]]:
        """
        Minor ID.
        """
        return pulumi.get(self, "minor_id")

    @minor_id.setter
    def minor_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minor_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Bluetooth Low Energy profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def txpower(self) -> Optional[pulumi.Input[str]]:
        """
        Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
        """
        return pulumi.get(self, "txpower")

    @txpower.setter
    def txpower(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "txpower", value)

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
class _WirelessControllerBleProfileState:
    def __init__(__self__, *,
                 advertising: Optional[pulumi.Input[str]] = None,
                 beacon_interval: Optional[pulumi.Input[int]] = None,
                 ble_scanning: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 eddystone_instance: Optional[pulumi.Input[str]] = None,
                 eddystone_namespace: Optional[pulumi.Input[str]] = None,
                 eddystone_url: Optional[pulumi.Input[str]] = None,
                 eddystone_url_encode_hex: Optional[pulumi.Input[str]] = None,
                 ibeacon_uuid: Optional[pulumi.Input[str]] = None,
                 major_id: Optional[pulumi.Input[int]] = None,
                 minor_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 txpower: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerBleProfile resources.
        :param pulumi.Input[str] advertising: Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
        :param pulumi.Input[int] beacon_interval: Beacon interval (default = 100 msec).
        :param pulumi.Input[str] ble_scanning: Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] eddystone_instance: Eddystone instance ID.
        :param pulumi.Input[str] eddystone_namespace: Eddystone namespace ID.
        :param pulumi.Input[str] eddystone_url: Eddystone URL.
        :param pulumi.Input[str] eddystone_url_encode_hex: Eddystone encoded URL hexadecimal string
        :param pulumi.Input[str] ibeacon_uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[int] major_id: Major ID.
        :param pulumi.Input[int] minor_id: Minor ID.
        :param pulumi.Input[str] name: Bluetooth Low Energy profile name.
        :param pulumi.Input[str] txpower: Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if advertising is not None:
            pulumi.set(__self__, "advertising", advertising)
        if beacon_interval is not None:
            pulumi.set(__self__, "beacon_interval", beacon_interval)
        if ble_scanning is not None:
            pulumi.set(__self__, "ble_scanning", ble_scanning)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if eddystone_instance is not None:
            pulumi.set(__self__, "eddystone_instance", eddystone_instance)
        if eddystone_namespace is not None:
            pulumi.set(__self__, "eddystone_namespace", eddystone_namespace)
        if eddystone_url is not None:
            pulumi.set(__self__, "eddystone_url", eddystone_url)
        if eddystone_url_encode_hex is not None:
            pulumi.set(__self__, "eddystone_url_encode_hex", eddystone_url_encode_hex)
        if ibeacon_uuid is not None:
            pulumi.set(__self__, "ibeacon_uuid", ibeacon_uuid)
        if major_id is not None:
            pulumi.set(__self__, "major_id", major_id)
        if minor_id is not None:
            pulumi.set(__self__, "minor_id", minor_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if txpower is not None:
            pulumi.set(__self__, "txpower", txpower)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def advertising(self) -> Optional[pulumi.Input[str]]:
        """
        Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
        """
        return pulumi.get(self, "advertising")

    @advertising.setter
    def advertising(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "advertising", value)

    @property
    @pulumi.getter(name="beaconInterval")
    def beacon_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Beacon interval (default = 100 msec).
        """
        return pulumi.get(self, "beacon_interval")

    @beacon_interval.setter
    def beacon_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "beacon_interval", value)

    @property
    @pulumi.getter(name="bleScanning")
    def ble_scanning(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ble_scanning")

    @ble_scanning.setter
    def ble_scanning(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ble_scanning", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="eddystoneInstance")
    def eddystone_instance(self) -> Optional[pulumi.Input[str]]:
        """
        Eddystone instance ID.
        """
        return pulumi.get(self, "eddystone_instance")

    @eddystone_instance.setter
    def eddystone_instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eddystone_instance", value)

    @property
    @pulumi.getter(name="eddystoneNamespace")
    def eddystone_namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Eddystone namespace ID.
        """
        return pulumi.get(self, "eddystone_namespace")

    @eddystone_namespace.setter
    def eddystone_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eddystone_namespace", value)

    @property
    @pulumi.getter(name="eddystoneUrl")
    def eddystone_url(self) -> Optional[pulumi.Input[str]]:
        """
        Eddystone URL.
        """
        return pulumi.get(self, "eddystone_url")

    @eddystone_url.setter
    def eddystone_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eddystone_url", value)

    @property
    @pulumi.getter(name="eddystoneUrlEncodeHex")
    def eddystone_url_encode_hex(self) -> Optional[pulumi.Input[str]]:
        """
        Eddystone encoded URL hexadecimal string
        """
        return pulumi.get(self, "eddystone_url_encode_hex")

    @eddystone_url_encode_hex.setter
    def eddystone_url_encode_hex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eddystone_url_encode_hex", value)

    @property
    @pulumi.getter(name="ibeaconUuid")
    def ibeacon_uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "ibeacon_uuid")

    @ibeacon_uuid.setter
    def ibeacon_uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ibeacon_uuid", value)

    @property
    @pulumi.getter(name="majorId")
    def major_id(self) -> Optional[pulumi.Input[int]]:
        """
        Major ID.
        """
        return pulumi.get(self, "major_id")

    @major_id.setter
    def major_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "major_id", value)

    @property
    @pulumi.getter(name="minorId")
    def minor_id(self) -> Optional[pulumi.Input[int]]:
        """
        Minor ID.
        """
        return pulumi.get(self, "minor_id")

    @minor_id.setter
    def minor_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minor_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Bluetooth Low Energy profile name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def txpower(self) -> Optional[pulumi.Input[str]]:
        """
        Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
        """
        return pulumi.get(self, "txpower")

    @txpower.setter
    def txpower(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "txpower", value)

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


class WirelessControllerBleProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advertising: Optional[pulumi.Input[str]] = None,
                 beacon_interval: Optional[pulumi.Input[int]] = None,
                 ble_scanning: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 eddystone_instance: Optional[pulumi.Input[str]] = None,
                 eddystone_namespace: Optional[pulumi.Input[str]] = None,
                 eddystone_url: Optional[pulumi.Input[str]] = None,
                 eddystone_url_encode_hex: Optional[pulumi.Input[str]] = None,
                 ibeacon_uuid: Optional[pulumi.Input[str]] = None,
                 major_id: Optional[pulumi.Input[int]] = None,
                 minor_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 txpower: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure Bluetooth Low Energy profile.

        ## Import

        WirelessController BleProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerBleProfile:WirelessControllerBleProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] advertising: Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
        :param pulumi.Input[int] beacon_interval: Beacon interval (default = 100 msec).
        :param pulumi.Input[str] ble_scanning: Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] eddystone_instance: Eddystone instance ID.
        :param pulumi.Input[str] eddystone_namespace: Eddystone namespace ID.
        :param pulumi.Input[str] eddystone_url: Eddystone URL.
        :param pulumi.Input[str] eddystone_url_encode_hex: Eddystone encoded URL hexadecimal string
        :param pulumi.Input[str] ibeacon_uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[int] major_id: Major ID.
        :param pulumi.Input[int] minor_id: Minor ID.
        :param pulumi.Input[str] name: Bluetooth Low Energy profile name.
        :param pulumi.Input[str] txpower: Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerBleProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure Bluetooth Low Energy profile.

        ## Import

        WirelessController BleProfile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/wirelessControllerBleProfile:WirelessControllerBleProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelessControllerBleProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerBleProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 advertising: Optional[pulumi.Input[str]] = None,
                 beacon_interval: Optional[pulumi.Input[int]] = None,
                 ble_scanning: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 eddystone_instance: Optional[pulumi.Input[str]] = None,
                 eddystone_namespace: Optional[pulumi.Input[str]] = None,
                 eddystone_url: Optional[pulumi.Input[str]] = None,
                 eddystone_url_encode_hex: Optional[pulumi.Input[str]] = None,
                 ibeacon_uuid: Optional[pulumi.Input[str]] = None,
                 major_id: Optional[pulumi.Input[int]] = None,
                 minor_id: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 txpower: Optional[pulumi.Input[str]] = None,
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
            __props__ = WirelessControllerBleProfileArgs.__new__(WirelessControllerBleProfileArgs)

            __props__.__dict__["advertising"] = advertising
            __props__.__dict__["beacon_interval"] = beacon_interval
            __props__.__dict__["ble_scanning"] = ble_scanning
            __props__.__dict__["comment"] = comment
            __props__.__dict__["eddystone_instance"] = eddystone_instance
            __props__.__dict__["eddystone_namespace"] = eddystone_namespace
            __props__.__dict__["eddystone_url"] = eddystone_url
            __props__.__dict__["eddystone_url_encode_hex"] = eddystone_url_encode_hex
            __props__.__dict__["ibeacon_uuid"] = ibeacon_uuid
            __props__.__dict__["major_id"] = major_id
            __props__.__dict__["minor_id"] = minor_id
            __props__.__dict__["name"] = name
            __props__.__dict__["txpower"] = txpower
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelessControllerBleProfile, __self__).__init__(
            'fortios:index/wirelessControllerBleProfile:WirelessControllerBleProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            advertising: Optional[pulumi.Input[str]] = None,
            beacon_interval: Optional[pulumi.Input[int]] = None,
            ble_scanning: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            eddystone_instance: Optional[pulumi.Input[str]] = None,
            eddystone_namespace: Optional[pulumi.Input[str]] = None,
            eddystone_url: Optional[pulumi.Input[str]] = None,
            eddystone_url_encode_hex: Optional[pulumi.Input[str]] = None,
            ibeacon_uuid: Optional[pulumi.Input[str]] = None,
            major_id: Optional[pulumi.Input[int]] = None,
            minor_id: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            txpower: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerBleProfile':
        """
        Get an existing WirelessControllerBleProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] advertising: Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
        :param pulumi.Input[int] beacon_interval: Beacon interval (default = 100 msec).
        :param pulumi.Input[str] ble_scanning: Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] eddystone_instance: Eddystone instance ID.
        :param pulumi.Input[str] eddystone_namespace: Eddystone namespace ID.
        :param pulumi.Input[str] eddystone_url: Eddystone URL.
        :param pulumi.Input[str] eddystone_url_encode_hex: Eddystone encoded URL hexadecimal string
        :param pulumi.Input[str] ibeacon_uuid: Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        :param pulumi.Input[int] major_id: Major ID.
        :param pulumi.Input[int] minor_id: Minor ID.
        :param pulumi.Input[str] name: Bluetooth Low Energy profile name.
        :param pulumi.Input[str] txpower: Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerBleProfileState.__new__(_WirelessControllerBleProfileState)

        __props__.__dict__["advertising"] = advertising
        __props__.__dict__["beacon_interval"] = beacon_interval
        __props__.__dict__["ble_scanning"] = ble_scanning
        __props__.__dict__["comment"] = comment
        __props__.__dict__["eddystone_instance"] = eddystone_instance
        __props__.__dict__["eddystone_namespace"] = eddystone_namespace
        __props__.__dict__["eddystone_url"] = eddystone_url
        __props__.__dict__["eddystone_url_encode_hex"] = eddystone_url_encode_hex
        __props__.__dict__["ibeacon_uuid"] = ibeacon_uuid
        __props__.__dict__["major_id"] = major_id
        __props__.__dict__["minor_id"] = minor_id
        __props__.__dict__["name"] = name
        __props__.__dict__["txpower"] = txpower
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelessControllerBleProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def advertising(self) -> pulumi.Output[str]:
        """
        Advertising type. Valid values: `ibeacon`, `eddystone-uid`, `eddystone-url`.
        """
        return pulumi.get(self, "advertising")

    @property
    @pulumi.getter(name="beaconInterval")
    def beacon_interval(self) -> pulumi.Output[int]:
        """
        Beacon interval (default = 100 msec).
        """
        return pulumi.get(self, "beacon_interval")

    @property
    @pulumi.getter(name="bleScanning")
    def ble_scanning(self) -> pulumi.Output[str]:
        """
        Enable/disable Bluetooth Low Energy (BLE) scanning. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ble_scanning")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[str]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="eddystoneInstance")
    def eddystone_instance(self) -> pulumi.Output[str]:
        """
        Eddystone instance ID.
        """
        return pulumi.get(self, "eddystone_instance")

    @property
    @pulumi.getter(name="eddystoneNamespace")
    def eddystone_namespace(self) -> pulumi.Output[str]:
        """
        Eddystone namespace ID.
        """
        return pulumi.get(self, "eddystone_namespace")

    @property
    @pulumi.getter(name="eddystoneUrl")
    def eddystone_url(self) -> pulumi.Output[str]:
        """
        Eddystone URL.
        """
        return pulumi.get(self, "eddystone_url")

    @property
    @pulumi.getter(name="eddystoneUrlEncodeHex")
    def eddystone_url_encode_hex(self) -> pulumi.Output[str]:
        """
        Eddystone encoded URL hexadecimal string
        """
        return pulumi.get(self, "eddystone_url_encode_hex")

    @property
    @pulumi.getter(name="ibeaconUuid")
    def ibeacon_uuid(self) -> pulumi.Output[str]:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "ibeacon_uuid")

    @property
    @pulumi.getter(name="majorId")
    def major_id(self) -> pulumi.Output[int]:
        """
        Major ID.
        """
        return pulumi.get(self, "major_id")

    @property
    @pulumi.getter(name="minorId")
    def minor_id(self) -> pulumi.Output[int]:
        """
        Minor ID.
        """
        return pulumi.get(self, "minor_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Bluetooth Low Energy profile name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def txpower(self) -> pulumi.Output[str]:
        """
        Transmit power level (default = 0). Valid values: `0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`.
        """
        return pulumi.get(self, "txpower")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

