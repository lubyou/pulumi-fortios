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

__all__ = [
    'GetSystemAutomationTriggerResult',
    'AwaitableGetSystemAutomationTriggerResult',
    'get_system_automation_trigger',
    'get_system_automation_trigger_output',
]

@pulumi.output_type
class GetSystemAutomationTriggerResult:
    """
    A collection of values returned by GetSystemAutomationTrigger.
    """
    def __init__(__self__, description=None, event_type=None, fabric_event_name=None, fabric_event_severity=None, faz_event_name=None, faz_event_severity=None, faz_event_tags=None, fields=None, id=None, ioc_level=None, license_type=None, logid=None, logid_blocks=None, name=None, report_type=None, serial=None, trigger_datetime=None, trigger_day=None, trigger_frequency=None, trigger_hour=None, trigger_minute=None, trigger_type=None, trigger_weekday=None, vdomparam=None, vdoms=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if event_type and not isinstance(event_type, str):
            raise TypeError("Expected argument 'event_type' to be a str")
        pulumi.set(__self__, "event_type", event_type)
        if fabric_event_name and not isinstance(fabric_event_name, str):
            raise TypeError("Expected argument 'fabric_event_name' to be a str")
        pulumi.set(__self__, "fabric_event_name", fabric_event_name)
        if fabric_event_severity and not isinstance(fabric_event_severity, str):
            raise TypeError("Expected argument 'fabric_event_severity' to be a str")
        pulumi.set(__self__, "fabric_event_severity", fabric_event_severity)
        if faz_event_name and not isinstance(faz_event_name, str):
            raise TypeError("Expected argument 'faz_event_name' to be a str")
        pulumi.set(__self__, "faz_event_name", faz_event_name)
        if faz_event_severity and not isinstance(faz_event_severity, str):
            raise TypeError("Expected argument 'faz_event_severity' to be a str")
        pulumi.set(__self__, "faz_event_severity", faz_event_severity)
        if faz_event_tags and not isinstance(faz_event_tags, str):
            raise TypeError("Expected argument 'faz_event_tags' to be a str")
        pulumi.set(__self__, "faz_event_tags", faz_event_tags)
        if fields and not isinstance(fields, list):
            raise TypeError("Expected argument 'fields' to be a list")
        pulumi.set(__self__, "fields", fields)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ioc_level and not isinstance(ioc_level, str):
            raise TypeError("Expected argument 'ioc_level' to be a str")
        pulumi.set(__self__, "ioc_level", ioc_level)
        if license_type and not isinstance(license_type, str):
            raise TypeError("Expected argument 'license_type' to be a str")
        pulumi.set(__self__, "license_type", license_type)
        if logid and not isinstance(logid, int):
            raise TypeError("Expected argument 'logid' to be a int")
        pulumi.set(__self__, "logid", logid)
        if logid_blocks and not isinstance(logid_blocks, list):
            raise TypeError("Expected argument 'logid_blocks' to be a list")
        pulumi.set(__self__, "logid_blocks", logid_blocks)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if report_type and not isinstance(report_type, str):
            raise TypeError("Expected argument 'report_type' to be a str")
        pulumi.set(__self__, "report_type", report_type)
        if serial and not isinstance(serial, str):
            raise TypeError("Expected argument 'serial' to be a str")
        pulumi.set(__self__, "serial", serial)
        if trigger_datetime and not isinstance(trigger_datetime, str):
            raise TypeError("Expected argument 'trigger_datetime' to be a str")
        pulumi.set(__self__, "trigger_datetime", trigger_datetime)
        if trigger_day and not isinstance(trigger_day, int):
            raise TypeError("Expected argument 'trigger_day' to be a int")
        pulumi.set(__self__, "trigger_day", trigger_day)
        if trigger_frequency and not isinstance(trigger_frequency, str):
            raise TypeError("Expected argument 'trigger_frequency' to be a str")
        pulumi.set(__self__, "trigger_frequency", trigger_frequency)
        if trigger_hour and not isinstance(trigger_hour, int):
            raise TypeError("Expected argument 'trigger_hour' to be a int")
        pulumi.set(__self__, "trigger_hour", trigger_hour)
        if trigger_minute and not isinstance(trigger_minute, int):
            raise TypeError("Expected argument 'trigger_minute' to be a int")
        pulumi.set(__self__, "trigger_minute", trigger_minute)
        if trigger_type and not isinstance(trigger_type, str):
            raise TypeError("Expected argument 'trigger_type' to be a str")
        pulumi.set(__self__, "trigger_type", trigger_type)
        if trigger_weekday and not isinstance(trigger_weekday, str):
            raise TypeError("Expected argument 'trigger_weekday' to be a str")
        pulumi.set(__self__, "trigger_weekday", trigger_weekday)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)
        if vdoms and not isinstance(vdoms, list):
            raise TypeError("Expected argument 'vdoms' to be a list")
        pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eventType")
    def event_type(self) -> str:
        return pulumi.get(self, "event_type")

    @property
    @pulumi.getter(name="fabricEventName")
    def fabric_event_name(self) -> str:
        return pulumi.get(self, "fabric_event_name")

    @property
    @pulumi.getter(name="fabricEventSeverity")
    def fabric_event_severity(self) -> str:
        return pulumi.get(self, "fabric_event_severity")

    @property
    @pulumi.getter(name="fazEventName")
    def faz_event_name(self) -> str:
        return pulumi.get(self, "faz_event_name")

    @property
    @pulumi.getter(name="fazEventSeverity")
    def faz_event_severity(self) -> str:
        return pulumi.get(self, "faz_event_severity")

    @property
    @pulumi.getter(name="fazEventTags")
    def faz_event_tags(self) -> str:
        return pulumi.get(self, "faz_event_tags")

    @property
    @pulumi.getter
    def fields(self) -> Sequence['outputs.GetSystemAutomationTriggerFieldResult']:
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="iocLevel")
    def ioc_level(self) -> str:
        return pulumi.get(self, "ioc_level")

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> str:
        return pulumi.get(self, "license_type")

    @property
    @pulumi.getter
    def logid(self) -> int:
        return pulumi.get(self, "logid")

    @property
    @pulumi.getter(name="logidBlocks")
    def logid_blocks(self) -> Sequence['outputs.GetSystemAutomationTriggerLogidBlockResult']:
        return pulumi.get(self, "logid_blocks")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="reportType")
    def report_type(self) -> str:
        return pulumi.get(self, "report_type")

    @property
    @pulumi.getter
    def serial(self) -> str:
        return pulumi.get(self, "serial")

    @property
    @pulumi.getter(name="triggerDatetime")
    def trigger_datetime(self) -> str:
        return pulumi.get(self, "trigger_datetime")

    @property
    @pulumi.getter(name="triggerDay")
    def trigger_day(self) -> int:
        return pulumi.get(self, "trigger_day")

    @property
    @pulumi.getter(name="triggerFrequency")
    def trigger_frequency(self) -> str:
        return pulumi.get(self, "trigger_frequency")

    @property
    @pulumi.getter(name="triggerHour")
    def trigger_hour(self) -> int:
        return pulumi.get(self, "trigger_hour")

    @property
    @pulumi.getter(name="triggerMinute")
    def trigger_minute(self) -> int:
        return pulumi.get(self, "trigger_minute")

    @property
    @pulumi.getter(name="triggerType")
    def trigger_type(self) -> str:
        return pulumi.get(self, "trigger_type")

    @property
    @pulumi.getter(name="triggerWeekday")
    def trigger_weekday(self) -> str:
        return pulumi.get(self, "trigger_weekday")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vdoms(self) -> Sequence['outputs.GetSystemAutomationTriggerVdomResult']:
        return pulumi.get(self, "vdoms")


class AwaitableGetSystemAutomationTriggerResult(GetSystemAutomationTriggerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAutomationTriggerResult(
            description=self.description,
            event_type=self.event_type,
            fabric_event_name=self.fabric_event_name,
            fabric_event_severity=self.fabric_event_severity,
            faz_event_name=self.faz_event_name,
            faz_event_severity=self.faz_event_severity,
            faz_event_tags=self.faz_event_tags,
            fields=self.fields,
            id=self.id,
            ioc_level=self.ioc_level,
            license_type=self.license_type,
            logid=self.logid,
            logid_blocks=self.logid_blocks,
            name=self.name,
            report_type=self.report_type,
            serial=self.serial,
            trigger_datetime=self.trigger_datetime,
            trigger_day=self.trigger_day,
            trigger_frequency=self.trigger_frequency,
            trigger_hour=self.trigger_hour,
            trigger_minute=self.trigger_minute,
            trigger_type=self.trigger_type,
            trigger_weekday=self.trigger_weekday,
            vdomparam=self.vdomparam,
            vdoms=self.vdoms)


def get_system_automation_trigger(name: Optional[str] = None,
                                  vdomparam: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAutomationTriggerResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemAutomationTrigger:GetSystemAutomationTrigger', __args__, opts=opts, typ=GetSystemAutomationTriggerResult).value

    return AwaitableGetSystemAutomationTriggerResult(
        description=__ret__.description,
        event_type=__ret__.event_type,
        fabric_event_name=__ret__.fabric_event_name,
        fabric_event_severity=__ret__.fabric_event_severity,
        faz_event_name=__ret__.faz_event_name,
        faz_event_severity=__ret__.faz_event_severity,
        faz_event_tags=__ret__.faz_event_tags,
        fields=__ret__.fields,
        id=__ret__.id,
        ioc_level=__ret__.ioc_level,
        license_type=__ret__.license_type,
        logid=__ret__.logid,
        logid_blocks=__ret__.logid_blocks,
        name=__ret__.name,
        report_type=__ret__.report_type,
        serial=__ret__.serial,
        trigger_datetime=__ret__.trigger_datetime,
        trigger_day=__ret__.trigger_day,
        trigger_frequency=__ret__.trigger_frequency,
        trigger_hour=__ret__.trigger_hour,
        trigger_minute=__ret__.trigger_minute,
        trigger_type=__ret__.trigger_type,
        trigger_weekday=__ret__.trigger_weekday,
        vdomparam=__ret__.vdomparam,
        vdoms=__ret__.vdoms)


@_utilities.lift_output_func(get_system_automation_trigger)
def get_system_automation_trigger_output(name: Optional[pulumi.Input[str]] = None,
                                         vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSystemAutomationTriggerResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
