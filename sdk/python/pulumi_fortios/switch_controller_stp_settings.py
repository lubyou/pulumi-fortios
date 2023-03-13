# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SwitchControllerStpSettingsArgs', 'SwitchControllerStpSettings']

@pulumi.input_type
class SwitchControllerStpSettingsArgs:
    def __init__(__self__, *,
                 forward_time: Optional[pulumi.Input[int]] = None,
                 hello_time: Optional[pulumi.Input[int]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pending_timer: Optional[pulumi.Input[int]] = None,
                 revision: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerStpSettings resource.
        """
        if forward_time is not None:
            pulumi.set(__self__, "forward_time", forward_time)
        if hello_time is not None:
            pulumi.set(__self__, "hello_time", hello_time)
        if max_age is not None:
            pulumi.set(__self__, "max_age", max_age)
        if max_hops is not None:
            pulumi.set(__self__, "max_hops", max_hops)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pending_timer is not None:
            pulumi.set(__self__, "pending_timer", pending_timer)
        if revision is not None:
            pulumi.set(__self__, "revision", revision)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="forwardTime")
    def forward_time(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forward_time")

    @forward_time.setter
    def forward_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forward_time", value)

    @property
    @pulumi.getter(name="helloTime")
    def hello_time(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "hello_time")

    @hello_time.setter
    def hello_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hello_time", value)

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_age")

    @max_age.setter
    def max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age", value)

    @property
    @pulumi.getter(name="maxHops")
    def max_hops(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_hops")

    @max_hops.setter
    def max_hops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_hops", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pendingTimer")
    def pending_timer(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "pending_timer")

    @pending_timer.setter
    def pending_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pending_timer", value)

    @property
    @pulumi.getter
    def revision(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "revision", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SwitchControllerStpSettingsState:
    def __init__(__self__, *,
                 forward_time: Optional[pulumi.Input[int]] = None,
                 hello_time: Optional[pulumi.Input[int]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pending_timer: Optional[pulumi.Input[int]] = None,
                 revision: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerStpSettings resources.
        """
        if forward_time is not None:
            pulumi.set(__self__, "forward_time", forward_time)
        if hello_time is not None:
            pulumi.set(__self__, "hello_time", hello_time)
        if max_age is not None:
            pulumi.set(__self__, "max_age", max_age)
        if max_hops is not None:
            pulumi.set(__self__, "max_hops", max_hops)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pending_timer is not None:
            pulumi.set(__self__, "pending_timer", pending_timer)
        if revision is not None:
            pulumi.set(__self__, "revision", revision)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="forwardTime")
    def forward_time(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "forward_time")

    @forward_time.setter
    def forward_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "forward_time", value)

    @property
    @pulumi.getter(name="helloTime")
    def hello_time(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "hello_time")

    @hello_time.setter
    def hello_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hello_time", value)

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_age")

    @max_age.setter
    def max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age", value)

    @property
    @pulumi.getter(name="maxHops")
    def max_hops(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_hops")

    @max_hops.setter
    def max_hops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_hops", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pendingTimer")
    def pending_timer(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "pending_timer")

    @pending_timer.setter
    def pending_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pending_timer", value)

    @property
    @pulumi.getter
    def revision(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "revision", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SwitchControllerStpSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 forward_time: Optional[pulumi.Input[int]] = None,
                 hello_time: Optional[pulumi.Input[int]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pending_timer: Optional[pulumi.Input[int]] = None,
                 revision: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SwitchControllerStpSettings resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerStpSettingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SwitchControllerStpSettings resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SwitchControllerStpSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerStpSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 forward_time: Optional[pulumi.Input[int]] = None,
                 hello_time: Optional[pulumi.Input[int]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_hops: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pending_timer: Optional[pulumi.Input[int]] = None,
                 revision: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchControllerStpSettingsArgs.__new__(SwitchControllerStpSettingsArgs)

            __props__.__dict__["forward_time"] = forward_time
            __props__.__dict__["hello_time"] = hello_time
            __props__.__dict__["max_age"] = max_age
            __props__.__dict__["max_hops"] = max_hops
            __props__.__dict__["name"] = name
            __props__.__dict__["pending_timer"] = pending_timer
            __props__.__dict__["revision"] = revision
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerStpSettings, __self__).__init__(
            'fortios:index/switchControllerStpSettings:SwitchControllerStpSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            forward_time: Optional[pulumi.Input[int]] = None,
            hello_time: Optional[pulumi.Input[int]] = None,
            max_age: Optional[pulumi.Input[int]] = None,
            max_hops: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pending_timer: Optional[pulumi.Input[int]] = None,
            revision: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerStpSettings':
        """
        Get an existing SwitchControllerStpSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerStpSettingsState.__new__(_SwitchControllerStpSettingsState)

        __props__.__dict__["forward_time"] = forward_time
        __props__.__dict__["hello_time"] = hello_time
        __props__.__dict__["max_age"] = max_age
        __props__.__dict__["max_hops"] = max_hops
        __props__.__dict__["name"] = name
        __props__.__dict__["pending_timer"] = pending_timer
        __props__.__dict__["revision"] = revision
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerStpSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="forwardTime")
    def forward_time(self) -> pulumi.Output[int]:
        return pulumi.get(self, "forward_time")

    @property
    @pulumi.getter(name="helloTime")
    def hello_time(self) -> pulumi.Output[int]:
        return pulumi.get(self, "hello_time")

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> pulumi.Output[int]:
        return pulumi.get(self, "max_age")

    @property
    @pulumi.getter(name="maxHops")
    def max_hops(self) -> pulumi.Output[int]:
        return pulumi.get(self, "max_hops")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pendingTimer")
    def pending_timer(self) -> pulumi.Output[int]:
        return pulumi.get(self, "pending_timer")

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Output[int]:
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

