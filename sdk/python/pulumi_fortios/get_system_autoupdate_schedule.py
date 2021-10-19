# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSystemAutoupdateScheduleResult',
    'AwaitableGetSystemAutoupdateScheduleResult',
    'get_system_autoupdate_schedule',
]

@pulumi.output_type
class GetSystemAutoupdateScheduleResult:
    """
    A collection of values returned by GetSystemAutoupdateSchedule.
    """
    def __init__(__self__, day=None, frequency=None, id=None, status=None, time=None, vdomparam=None):
        if day and not isinstance(day, str):
            raise TypeError("Expected argument 'day' to be a str")
        pulumi.set(__self__, "day", day)
        if frequency and not isinstance(frequency, str):
            raise TypeError("Expected argument 'frequency' to be a str")
        pulumi.set(__self__, "frequency", frequency)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if time and not isinstance(time, str):
            raise TypeError("Expected argument 'time' to be a str")
        pulumi.set(__self__, "time", time)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def day(self) -> str:
        """
        Update day.
        """
        return pulumi.get(self, "day")

    @property
    @pulumi.getter
    def frequency(self) -> str:
        """
        Update frequency.
        """
        return pulumi.get(self, "frequency")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable scheduled updates.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def time(self) -> str:
        """
        Update time.
        """
        return pulumi.get(self, "time")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetSystemAutoupdateScheduleResult(GetSystemAutoupdateScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSystemAutoupdateScheduleResult(
            day=self.day,
            frequency=self.frequency,
            id=self.id,
            status=self.status,
            time=self.time,
            vdomparam=self.vdomparam)


def get_system_autoupdate_schedule(vdomparam: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSystemAutoupdateScheduleResult:
    """
    Use this data source to get information on fortios systemautoupdate schedule


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getSystemAutoupdateSchedule:GetSystemAutoupdateSchedule', __args__, opts=opts, typ=GetSystemAutoupdateScheduleResult).value

    return AwaitableGetSystemAutoupdateScheduleResult(
        day=__ret__.day,
        frequency=__ret__.frequency,
        id=__ret__.id,
        status=__ret__.status,
        time=__ret__.time,
        vdomparam=__ret__.vdomparam)
