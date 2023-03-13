# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallScheduleOnetimeArgs', 'FirewallScheduleOnetime']

@pulumi.input_type
class FirewallScheduleOnetimeArgs:
    def __init__(__self__, *,
                 end: pulumi.Input[str],
                 start: pulumi.Input[str],
                 color: Optional[pulumi.Input[int]] = None,
                 expiration_days: Optional[pulumi.Input[int]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallScheduleOnetime resource.
        """
        pulumi.set(__self__, "end", end)
        pulumi.set(__self__, "start", start)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if expiration_days is not None:
            pulumi.set(__self__, "expiration_days", expiration_days)
        if fabric_object is not None:
            pulumi.set(__self__, "fabric_object", fabric_object)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def end(self) -> pulumi.Input[str]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: pulumi.Input[str]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> pulumi.Input[str]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: pulumi.Input[str]):
        pulumi.set(self, "start", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter(name="expirationDays")
    def expiration_days(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "expiration_days")

    @expiration_days.setter
    def expiration_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiration_days", value)

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fabric_object")

    @fabric_object.setter
    def fabric_object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fabric_object", value)

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
class _FirewallScheduleOnetimeState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[int]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 expiration_days: Optional[pulumi.Input[int]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallScheduleOnetime resources.
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if end is not None:
            pulumi.set(__self__, "end", end)
        if expiration_days is not None:
            pulumi.set(__self__, "expiration_days", expiration_days)
        if fabric_object is not None:
            pulumi.set(__self__, "fabric_object", fabric_object)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if start is not None:
            pulumi.set(__self__, "start", start)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter(name="expirationDays")
    def expiration_days(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "expiration_days")

    @expiration_days.setter
    def expiration_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiration_days", value)

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fabric_object")

    @fabric_object.setter
    def fabric_object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fabric_object", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class FirewallScheduleOnetime(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 expiration_days: Optional[pulumi.Input[int]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FirewallScheduleOnetime resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallScheduleOnetimeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FirewallScheduleOnetime resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FirewallScheduleOnetimeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallScheduleOnetimeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 end: Optional[pulumi.Input[str]] = None,
                 expiration_days: Optional[pulumi.Input[int]] = None,
                 fabric_object: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallScheduleOnetimeArgs.__new__(FirewallScheduleOnetimeArgs)

            __props__.__dict__["color"] = color
            if end is None and not opts.urn:
                raise TypeError("Missing required property 'end'")
            __props__.__dict__["end"] = end
            __props__.__dict__["expiration_days"] = expiration_days
            __props__.__dict__["fabric_object"] = fabric_object
            __props__.__dict__["name"] = name
            if start is None and not opts.urn:
                raise TypeError("Missing required property 'start'")
            __props__.__dict__["start"] = start
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallScheduleOnetime, __self__).__init__(
            'fortios:index/firewallScheduleOnetime:FirewallScheduleOnetime',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[int]] = None,
            end: Optional[pulumi.Input[str]] = None,
            expiration_days: Optional[pulumi.Input[int]] = None,
            fabric_object: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            start: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallScheduleOnetime':
        """
        Get an existing FirewallScheduleOnetime resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallScheduleOnetimeState.__new__(_FirewallScheduleOnetimeState)

        __props__.__dict__["color"] = color
        __props__.__dict__["end"] = end
        __props__.__dict__["expiration_days"] = expiration_days
        __props__.__dict__["fabric_object"] = fabric_object
        __props__.__dict__["name"] = name
        __props__.__dict__["start"] = start
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallScheduleOnetime(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[int]:
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def end(self) -> pulumi.Output[str]:
        return pulumi.get(self, "end")

    @property
    @pulumi.getter(name="expirationDays")
    def expiration_days(self) -> pulumi.Output[int]:
        return pulumi.get(self, "expiration_days")

    @property
    @pulumi.getter(name="fabricObject")
    def fabric_object(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fabric_object")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def start(self) -> pulumi.Output[str]:
        return pulumi.get(self, "start")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

