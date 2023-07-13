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

__all__ = ['WirelessControllerSettingArgs', 'WirelessControllerSetting']

@pulumi.input_type
class WirelessControllerSettingArgs:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 darrp_optimize: Optional[pulumi.Input[int]] = None,
                 darrp_optimize_schedules: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]] = None,
                 device_holdoff: Optional[pulumi.Input[int]] = None,
                 device_idle: Optional[pulumi.Input[int]] = None,
                 device_weight: Optional[pulumi.Input[int]] = None,
                 duplicate_ssid: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fake_ssid_action: Optional[pulumi.Input[str]] = None,
                 fapc_compatibility: Optional[pulumi.Input[str]] = None,
                 firmware_provision_on_authorization: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 offending_ssids: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingOffendingSsidArgs']]]] = None,
                 phishing_ssid_detect: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wfa_compatibility: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelessControllerSetting resource.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if darrp_optimize is not None:
            pulumi.set(__self__, "darrp_optimize", darrp_optimize)
        if darrp_optimize_schedules is not None:
            pulumi.set(__self__, "darrp_optimize_schedules", darrp_optimize_schedules)
        if device_holdoff is not None:
            pulumi.set(__self__, "device_holdoff", device_holdoff)
        if device_idle is not None:
            pulumi.set(__self__, "device_idle", device_idle)
        if device_weight is not None:
            pulumi.set(__self__, "device_weight", device_weight)
        if duplicate_ssid is not None:
            pulumi.set(__self__, "duplicate_ssid", duplicate_ssid)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fake_ssid_action is not None:
            pulumi.set(__self__, "fake_ssid_action", fake_ssid_action)
        if fapc_compatibility is not None:
            pulumi.set(__self__, "fapc_compatibility", fapc_compatibility)
        if firmware_provision_on_authorization is not None:
            pulumi.set(__self__, "firmware_provision_on_authorization", firmware_provision_on_authorization)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if offending_ssids is not None:
            pulumi.set(__self__, "offending_ssids", offending_ssids)
        if phishing_ssid_detect is not None:
            pulumi.set(__self__, "phishing_ssid_detect", phishing_ssid_detect)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wfa_compatibility is not None:
            pulumi.set(__self__, "wfa_compatibility", wfa_compatibility)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter(name="darrpOptimize")
    def darrp_optimize(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "darrp_optimize")

    @darrp_optimize.setter
    def darrp_optimize(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "darrp_optimize", value)

    @property
    @pulumi.getter(name="darrpOptimizeSchedules")
    def darrp_optimize_schedules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]]:
        return pulumi.get(self, "darrp_optimize_schedules")

    @darrp_optimize_schedules.setter
    def darrp_optimize_schedules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]]):
        pulumi.set(self, "darrp_optimize_schedules", value)

    @property
    @pulumi.getter(name="deviceHoldoff")
    def device_holdoff(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "device_holdoff")

    @device_holdoff.setter
    def device_holdoff(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_holdoff", value)

    @property
    @pulumi.getter(name="deviceIdle")
    def device_idle(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "device_idle")

    @device_idle.setter
    def device_idle(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_idle", value)

    @property
    @pulumi.getter(name="deviceWeight")
    def device_weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "device_weight")

    @device_weight.setter
    def device_weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_weight", value)

    @property
    @pulumi.getter(name="duplicateSsid")
    def duplicate_ssid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "duplicate_ssid")

    @duplicate_ssid.setter
    def duplicate_ssid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "duplicate_ssid", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="fakeSsidAction")
    def fake_ssid_action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fake_ssid_action")

    @fake_ssid_action.setter
    def fake_ssid_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fake_ssid_action", value)

    @property
    @pulumi.getter(name="fapcCompatibility")
    def fapc_compatibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fapc_compatibility")

    @fapc_compatibility.setter
    def fapc_compatibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fapc_compatibility", value)

    @property
    @pulumi.getter(name="firmwareProvisionOnAuthorization")
    def firmware_provision_on_authorization(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "firmware_provision_on_authorization")

    @firmware_provision_on_authorization.setter
    def firmware_provision_on_authorization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firmware_provision_on_authorization", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="offendingSsids")
    def offending_ssids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingOffendingSsidArgs']]]]:
        return pulumi.get(self, "offending_ssids")

    @offending_ssids.setter
    def offending_ssids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingOffendingSsidArgs']]]]):
        pulumi.set(self, "offending_ssids", value)

    @property
    @pulumi.getter(name="phishingSsidDetect")
    def phishing_ssid_detect(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "phishing_ssid_detect")

    @phishing_ssid_detect.setter
    def phishing_ssid_detect(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phishing_ssid_detect", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="wfaCompatibility")
    def wfa_compatibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wfa_compatibility")

    @wfa_compatibility.setter
    def wfa_compatibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wfa_compatibility", value)


@pulumi.input_type
class _WirelessControllerSettingState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 darrp_optimize: Optional[pulumi.Input[int]] = None,
                 darrp_optimize_schedules: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]] = None,
                 device_holdoff: Optional[pulumi.Input[int]] = None,
                 device_idle: Optional[pulumi.Input[int]] = None,
                 device_weight: Optional[pulumi.Input[int]] = None,
                 duplicate_ssid: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fake_ssid_action: Optional[pulumi.Input[str]] = None,
                 fapc_compatibility: Optional[pulumi.Input[str]] = None,
                 firmware_provision_on_authorization: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 offending_ssids: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingOffendingSsidArgs']]]] = None,
                 phishing_ssid_detect: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wfa_compatibility: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelessControllerSetting resources.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if darrp_optimize is not None:
            pulumi.set(__self__, "darrp_optimize", darrp_optimize)
        if darrp_optimize_schedules is not None:
            pulumi.set(__self__, "darrp_optimize_schedules", darrp_optimize_schedules)
        if device_holdoff is not None:
            pulumi.set(__self__, "device_holdoff", device_holdoff)
        if device_idle is not None:
            pulumi.set(__self__, "device_idle", device_idle)
        if device_weight is not None:
            pulumi.set(__self__, "device_weight", device_weight)
        if duplicate_ssid is not None:
            pulumi.set(__self__, "duplicate_ssid", duplicate_ssid)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if fake_ssid_action is not None:
            pulumi.set(__self__, "fake_ssid_action", fake_ssid_action)
        if fapc_compatibility is not None:
            pulumi.set(__self__, "fapc_compatibility", fapc_compatibility)
        if firmware_provision_on_authorization is not None:
            pulumi.set(__self__, "firmware_provision_on_authorization", firmware_provision_on_authorization)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if offending_ssids is not None:
            pulumi.set(__self__, "offending_ssids", offending_ssids)
        if phishing_ssid_detect is not None:
            pulumi.set(__self__, "phishing_ssid_detect", phishing_ssid_detect)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if wfa_compatibility is not None:
            pulumi.set(__self__, "wfa_compatibility", wfa_compatibility)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter(name="darrpOptimize")
    def darrp_optimize(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "darrp_optimize")

    @darrp_optimize.setter
    def darrp_optimize(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "darrp_optimize", value)

    @property
    @pulumi.getter(name="darrpOptimizeSchedules")
    def darrp_optimize_schedules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]]:
        return pulumi.get(self, "darrp_optimize_schedules")

    @darrp_optimize_schedules.setter
    def darrp_optimize_schedules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]]):
        pulumi.set(self, "darrp_optimize_schedules", value)

    @property
    @pulumi.getter(name="deviceHoldoff")
    def device_holdoff(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "device_holdoff")

    @device_holdoff.setter
    def device_holdoff(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_holdoff", value)

    @property
    @pulumi.getter(name="deviceIdle")
    def device_idle(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "device_idle")

    @device_idle.setter
    def device_idle(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_idle", value)

    @property
    @pulumi.getter(name="deviceWeight")
    def device_weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "device_weight")

    @device_weight.setter
    def device_weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "device_weight", value)

    @property
    @pulumi.getter(name="duplicateSsid")
    def duplicate_ssid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "duplicate_ssid")

    @duplicate_ssid.setter
    def duplicate_ssid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "duplicate_ssid", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="fakeSsidAction")
    def fake_ssid_action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fake_ssid_action")

    @fake_ssid_action.setter
    def fake_ssid_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fake_ssid_action", value)

    @property
    @pulumi.getter(name="fapcCompatibility")
    def fapc_compatibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fapc_compatibility")

    @fapc_compatibility.setter
    def fapc_compatibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fapc_compatibility", value)

    @property
    @pulumi.getter(name="firmwareProvisionOnAuthorization")
    def firmware_provision_on_authorization(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "firmware_provision_on_authorization")

    @firmware_provision_on_authorization.setter
    def firmware_provision_on_authorization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firmware_provision_on_authorization", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter(name="offendingSsids")
    def offending_ssids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingOffendingSsidArgs']]]]:
        return pulumi.get(self, "offending_ssids")

    @offending_ssids.setter
    def offending_ssids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WirelessControllerSettingOffendingSsidArgs']]]]):
        pulumi.set(self, "offending_ssids", value)

    @property
    @pulumi.getter(name="phishingSsidDetect")
    def phishing_ssid_detect(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "phishing_ssid_detect")

    @phishing_ssid_detect.setter
    def phishing_ssid_detect(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phishing_ssid_detect", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="wfaCompatibility")
    def wfa_compatibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wfa_compatibility")

    @wfa_compatibility.setter
    def wfa_compatibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wfa_compatibility", value)


class WirelessControllerSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 darrp_optimize: Optional[pulumi.Input[int]] = None,
                 darrp_optimize_schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]]] = None,
                 device_holdoff: Optional[pulumi.Input[int]] = None,
                 device_idle: Optional[pulumi.Input[int]] = None,
                 device_weight: Optional[pulumi.Input[int]] = None,
                 duplicate_ssid: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fake_ssid_action: Optional[pulumi.Input[str]] = None,
                 fapc_compatibility: Optional[pulumi.Input[str]] = None,
                 firmware_provision_on_authorization: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 offending_ssids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerSettingOffendingSsidArgs']]]]] = None,
                 phishing_ssid_detect: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wfa_compatibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a WirelessControllerSetting resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelessControllerSettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a WirelessControllerSetting resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WirelessControllerSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelessControllerSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 darrp_optimize: Optional[pulumi.Input[int]] = None,
                 darrp_optimize_schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]]] = None,
                 device_holdoff: Optional[pulumi.Input[int]] = None,
                 device_idle: Optional[pulumi.Input[int]] = None,
                 device_weight: Optional[pulumi.Input[int]] = None,
                 duplicate_ssid: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 fake_ssid_action: Optional[pulumi.Input[str]] = None,
                 fapc_compatibility: Optional[pulumi.Input[str]] = None,
                 firmware_provision_on_authorization: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 offending_ssids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerSettingOffendingSsidArgs']]]]] = None,
                 phishing_ssid_detect: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 wfa_compatibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelessControllerSettingArgs.__new__(WirelessControllerSettingArgs)

            __props__.__dict__["account_id"] = account_id
            __props__.__dict__["country"] = country
            __props__.__dict__["darrp_optimize"] = darrp_optimize
            __props__.__dict__["darrp_optimize_schedules"] = darrp_optimize_schedules
            __props__.__dict__["device_holdoff"] = device_holdoff
            __props__.__dict__["device_idle"] = device_idle
            __props__.__dict__["device_weight"] = device_weight
            __props__.__dict__["duplicate_ssid"] = duplicate_ssid
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["fake_ssid_action"] = fake_ssid_action
            __props__.__dict__["fapc_compatibility"] = fapc_compatibility
            __props__.__dict__["firmware_provision_on_authorization"] = firmware_provision_on_authorization
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["offending_ssids"] = offending_ssids
            __props__.__dict__["phishing_ssid_detect"] = phishing_ssid_detect
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["wfa_compatibility"] = wfa_compatibility
        super(WirelessControllerSetting, __self__).__init__(
            'fortios:index/wirelessControllerSetting:WirelessControllerSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            country: Optional[pulumi.Input[str]] = None,
            darrp_optimize: Optional[pulumi.Input[int]] = None,
            darrp_optimize_schedules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerSettingDarrpOptimizeScheduleArgs']]]]] = None,
            device_holdoff: Optional[pulumi.Input[int]] = None,
            device_idle: Optional[pulumi.Input[int]] = None,
            device_weight: Optional[pulumi.Input[int]] = None,
            duplicate_ssid: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            fake_ssid_action: Optional[pulumi.Input[str]] = None,
            fapc_compatibility: Optional[pulumi.Input[str]] = None,
            firmware_provision_on_authorization: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            offending_ssids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WirelessControllerSettingOffendingSsidArgs']]]]] = None,
            phishing_ssid_detect: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            wfa_compatibility: Optional[pulumi.Input[str]] = None) -> 'WirelessControllerSetting':
        """
        Get an existing WirelessControllerSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelessControllerSettingState.__new__(_WirelessControllerSettingState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["country"] = country
        __props__.__dict__["darrp_optimize"] = darrp_optimize
        __props__.__dict__["darrp_optimize_schedules"] = darrp_optimize_schedules
        __props__.__dict__["device_holdoff"] = device_holdoff
        __props__.__dict__["device_idle"] = device_idle
        __props__.__dict__["device_weight"] = device_weight
        __props__.__dict__["duplicate_ssid"] = duplicate_ssid
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["fake_ssid_action"] = fake_ssid_action
        __props__.__dict__["fapc_compatibility"] = fapc_compatibility
        __props__.__dict__["firmware_provision_on_authorization"] = firmware_provision_on_authorization
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["offending_ssids"] = offending_ssids
        __props__.__dict__["phishing_ssid_detect"] = phishing_ssid_detect
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["wfa_compatibility"] = wfa_compatibility
        return WirelessControllerSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def country(self) -> pulumi.Output[str]:
        return pulumi.get(self, "country")

    @property
    @pulumi.getter(name="darrpOptimize")
    def darrp_optimize(self) -> pulumi.Output[int]:
        return pulumi.get(self, "darrp_optimize")

    @property
    @pulumi.getter(name="darrpOptimizeSchedules")
    def darrp_optimize_schedules(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerSettingDarrpOptimizeSchedule']]]:
        return pulumi.get(self, "darrp_optimize_schedules")

    @property
    @pulumi.getter(name="deviceHoldoff")
    def device_holdoff(self) -> pulumi.Output[int]:
        return pulumi.get(self, "device_holdoff")

    @property
    @pulumi.getter(name="deviceIdle")
    def device_idle(self) -> pulumi.Output[int]:
        return pulumi.get(self, "device_idle")

    @property
    @pulumi.getter(name="deviceWeight")
    def device_weight(self) -> pulumi.Output[int]:
        return pulumi.get(self, "device_weight")

    @property
    @pulumi.getter(name="duplicateSsid")
    def duplicate_ssid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "duplicate_ssid")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="fakeSsidAction")
    def fake_ssid_action(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fake_ssid_action")

    @property
    @pulumi.getter(name="fapcCompatibility")
    def fapc_compatibility(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fapc_compatibility")

    @property
    @pulumi.getter(name="firmwareProvisionOnAuthorization")
    def firmware_provision_on_authorization(self) -> pulumi.Output[str]:
        return pulumi.get(self, "firmware_provision_on_authorization")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter(name="offendingSsids")
    def offending_ssids(self) -> pulumi.Output[Optional[Sequence['outputs.WirelessControllerSettingOffendingSsid']]]:
        return pulumi.get(self, "offending_ssids")

    @property
    @pulumi.getter(name="phishingSsidDetect")
    def phishing_ssid_detect(self) -> pulumi.Output[str]:
        return pulumi.get(self, "phishing_ssid_detect")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="wfaCompatibility")
    def wfa_compatibility(self) -> pulumi.Output[str]:
        return pulumi.get(self, "wfa_compatibility")

