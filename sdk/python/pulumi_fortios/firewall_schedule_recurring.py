# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallScheduleRecurringArgs', 'FirewallScheduleRecurring']

@pulumi.input_type
class FirewallScheduleRecurringArgs:
    def __init__(__self__, *,
                 end: pulumi.Input[str],
                 start: pulumi.Input[str],
                 color: Optional[pulumi.Input[int]] = None,
                 day: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallScheduleRecurring resource.
        :param pulumi.Input[str] end: Time of day to end the schedule, format hh:mm.
        :param pulumi.Input[str] start: Time of day to start the schedule, format hh:mm.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] day: One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        :param pulumi.Input[str] name: Recurring schedule name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if day is not None:
            pulumi.set(__self__, "day", day)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def end(self) -> pulumi.Input[str]:
        """
        Time of day to end the schedule, format hh:mm.
        """
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: pulumi.Input[str]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> pulumi.Input[str]:
        """
        Time of day to start the schedule, format hh:mm.
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: pulumi.Input[str]):
        pulumi.set(self, "start", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def day(self) -> Optional[pulumi.Input[str]]:
        """
        One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        """
        return pulumi.get(self, "day")

    @day.setter
    def day(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Recurring schedule name.
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
class _FirewallScheduleRecurringState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[int]] = None,
                 day: Optional[pulumi.Input[str]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallScheduleRecurring resources.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] day: One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        :param pulumi.Input[str] end: Time of day to end the schedule, format hh:mm.
        :param pulumi.Input[str] name: Recurring schedule name.
        :param pulumi.Input[str] start: Time of day to start the schedule, format hh:mm.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if day is not None:
            pulumi.set(__self__, "day", day)
        if end is not None:
            pulumi.set(__self__, "end", end)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if start is not None:
            pulumi.set(__self__, "start", start)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def day(self) -> Optional[pulumi.Input[str]]:
        """
        One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        """
        return pulumi.get(self, "day")

    @day.setter
    def day(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day", value)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[str]]:
        """
        Time of day to end the schedule, format hh:mm.
        """
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Recurring schedule name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[str]]:
        """
        Time of day to start the schedule, format hh:mm.
        """
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start", value)

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


class FirewallScheduleRecurring(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 day: Optional[pulumi.Input[str]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Recurring schedule configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallScheduleRecurring("trname",
            color=0,
            day="sunday",
            end="00:00",
            start="00:00")
        ```

        ## Import

        FirewallSchedule Recurring can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallScheduleRecurring:FirewallScheduleRecurring labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] day: One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        :param pulumi.Input[str] end: Time of day to end the schedule, format hh:mm.
        :param pulumi.Input[str] name: Recurring schedule name.
        :param pulumi.Input[str] start: Time of day to start the schedule, format hh:mm.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallScheduleRecurringArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Recurring schedule configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallScheduleRecurring("trname",
            color=0,
            day="sunday",
            end="00:00",
            start="00:00")
        ```

        ## Import

        FirewallSchedule Recurring can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallScheduleRecurring:FirewallScheduleRecurring labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallScheduleRecurringArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallScheduleRecurringArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 day: Optional[pulumi.Input[str]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallScheduleRecurringArgs.__new__(FirewallScheduleRecurringArgs)

            __props__.__dict__["color"] = color
            __props__.__dict__["day"] = day
            if end is None and not opts.urn:
                raise TypeError("Missing required property 'end'")
            __props__.__dict__["end"] = end
            __props__.__dict__["name"] = name
            if start is None and not opts.urn:
                raise TypeError("Missing required property 'start'")
            __props__.__dict__["start"] = start
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallScheduleRecurring, __self__).__init__(
            'fortios:index/firewallScheduleRecurring:FirewallScheduleRecurring',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[int]] = None,
            day: Optional[pulumi.Input[str]] = None,
            end: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            start: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallScheduleRecurring':
        """
        Get an existing FirewallScheduleRecurring resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] day: One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        :param pulumi.Input[str] end: Time of day to end the schedule, format hh:mm.
        :param pulumi.Input[str] name: Recurring schedule name.
        :param pulumi.Input[str] start: Time of day to start the schedule, format hh:mm.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallScheduleRecurringState.__new__(_FirewallScheduleRecurringState)

        __props__.__dict__["color"] = color
        __props__.__dict__["day"] = day
        __props__.__dict__["end"] = end
        __props__.__dict__["name"] = name
        __props__.__dict__["start"] = start
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallScheduleRecurring(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[int]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def day(self) -> pulumi.Output[str]:
        """
        One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
        """
        return pulumi.get(self, "day")

    @property
    @pulumi.getter
    def end(self) -> pulumi.Output[str]:
        """
        Time of day to end the schedule, format hh:mm.
        """
        return pulumi.get(self, "end")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Recurring schedule name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def start(self) -> pulumi.Output[str]:
        """
        Time of day to start the schedule, format hh:mm.
        """
        return pulumi.get(self, "start")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

