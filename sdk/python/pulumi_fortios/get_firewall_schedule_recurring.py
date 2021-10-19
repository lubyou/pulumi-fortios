# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFirewallScheduleRecurringResult',
    'AwaitableGetFirewallScheduleRecurringResult',
    'get_firewall_schedule_recurring',
]

@pulumi.output_type
class GetFirewallScheduleRecurringResult:
    """
    A collection of values returned by GetFirewallScheduleRecurring.
    """
    def __init__(__self__, color=None, day=None, end=None, id=None, name=None, start=None, vdomparam=None):
        if color and not isinstance(color, int):
            raise TypeError("Expected argument 'color' to be a int")
        pulumi.set(__self__, "color", color)
        if day and not isinstance(day, str):
            raise TypeError("Expected argument 'day' to be a str")
        pulumi.set(__self__, "day", day)
        if end and not isinstance(end, str):
            raise TypeError("Expected argument 'end' to be a str")
        pulumi.set(__self__, "end", end)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start and not isinstance(start, str):
            raise TypeError("Expected argument 'start' to be a str")
        pulumi.set(__self__, "start", start)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def color(self) -> int:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def day(self) -> str:
        """
        One or more days of the week on which the schedule is valid. Separate the names of the days with a space.
        """
        return pulumi.get(self, "day")

    @property
    @pulumi.getter
    def end(self) -> str:
        """
        Time of day to end the schedule, format hh:mm.
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Recurring schedule name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def start(self) -> str:
        """
        Time of day to start the schedule, format hh:mm.
        """
        return pulumi.get(self, "start")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallScheduleRecurringResult(GetFirewallScheduleRecurringResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallScheduleRecurringResult(
            color=self.color,
            day=self.day,
            end=self.end,
            id=self.id,
            name=self.name,
            start=self.start,
            vdomparam=self.vdomparam)


def get_firewall_schedule_recurring(name: Optional[str] = None,
                                    vdomparam: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallScheduleRecurringResult:
    """
    Use this data source to get information on an fortios firewallschedule recurring


    :param str name: Specify the name of the desired firewallschedule recurring.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallScheduleRecurring:GetFirewallScheduleRecurring', __args__, opts=opts, typ=GetFirewallScheduleRecurringResult).value

    return AwaitableGetFirewallScheduleRecurringResult(
        color=__ret__.color,
        day=__ret__.day,
        end=__ret__.end,
        id=__ret__.id,
        name=__ret__.name,
        start=__ret__.start,
        vdomparam=__ret__.vdomparam)
