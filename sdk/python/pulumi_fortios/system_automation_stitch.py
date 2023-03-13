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

__all__ = ['SystemAutomationStitchArgs', 'SystemAutomationStitch']

@pulumi.input_type
class SystemAutomationStitchArgs:
    def __init__(__self__, *,
                 status: pulumi.Input[str],
                 trigger: pulumi.Input[str],
                 action: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destinations: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchDestinationArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemAutomationStitch resource.
        """
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "trigger", trigger)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destinations is not None:
            pulumi.set(__self__, "destinations", destinations)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def trigger(self) -> pulumi.Input[str]:
        return pulumi.get(self, "trigger")

    @trigger.setter
    def trigger(self, value: pulumi.Input[str]):
        pulumi.set(self, "trigger", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def destinations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchDestinationArgs']]]]:
        return pulumi.get(self, "destinations")

    @destinations.setter
    def destinations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchDestinationArgs']]]]):
        pulumi.set(self, "destinations", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SystemAutomationStitchState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destinations: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchDestinationArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 trigger: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemAutomationStitch resources.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destinations is not None:
            pulumi.set(__self__, "destinations", destinations)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if trigger is not None:
            pulumi.set(__self__, "trigger", trigger)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchActionArgs']]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def destinations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchDestinationArgs']]]]:
        return pulumi.get(self, "destinations")

    @destinations.setter
    def destinations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemAutomationStitchDestinationArgs']]]]):
        pulumi.set(self, "destinations", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def trigger(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trigger")

    @trigger.setter
    def trigger(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SystemAutomationStitch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchActionArgs']]]]] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchActionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destinations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchDestinationArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 trigger: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SystemAutomationStitch resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemAutomationStitchArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SystemAutomationStitch resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SystemAutomationStitchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemAutomationStitchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchActionArgs']]]]] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchActionArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destinations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchDestinationArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 trigger: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemAutomationStitchArgs.__new__(SystemAutomationStitchArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["actions"] = actions
            __props__.__dict__["description"] = description
            __props__.__dict__["destinations"] = destinations
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            if trigger is None and not opts.urn:
                raise TypeError("Missing required property 'trigger'")
            __props__.__dict__["trigger"] = trigger
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemAutomationStitch, __self__).__init__(
            'fortios:index/systemAutomationStitch:SystemAutomationStitch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchActionArgs']]]]] = None,
            actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchActionArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            destinations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemAutomationStitchDestinationArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            trigger: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemAutomationStitch':
        """
        Get an existing SystemAutomationStitch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemAutomationStitchState.__new__(_SystemAutomationStitchState)

        __props__.__dict__["action"] = action
        __props__.__dict__["actions"] = actions
        __props__.__dict__["description"] = description
        __props__.__dict__["destinations"] = destinations
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["status"] = status
        __props__.__dict__["trigger"] = trigger
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemAutomationStitch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[Sequence['outputs.SystemAutomationStitchAction']]]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Optional[Sequence['outputs.SystemAutomationStitchAction']]]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destinations(self) -> pulumi.Output[Optional[Sequence['outputs.SystemAutomationStitchDestination']]]:
        return pulumi.get(self, "destinations")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def trigger(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trigger")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

