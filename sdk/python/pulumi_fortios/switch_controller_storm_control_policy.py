# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SwitchControllerStormControlPolicyArgs', 'SwitchControllerStormControlPolicy']

@pulumi.input_type
class SwitchControllerStormControlPolicyArgs:
    def __init__(__self__, *,
                 broadcast: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[int]] = None,
                 storm_control_mode: Optional[pulumi.Input[str]] = None,
                 unknown_multicast: Optional[pulumi.Input[str]] = None,
                 unknown_unicast: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerStormControlPolicy resource.
        """
        if broadcast is not None:
            pulumi.set(__self__, "broadcast", broadcast)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rate is not None:
            pulumi.set(__self__, "rate", rate)
        if storm_control_mode is not None:
            pulumi.set(__self__, "storm_control_mode", storm_control_mode)
        if unknown_multicast is not None:
            pulumi.set(__self__, "unknown_multicast", unknown_multicast)
        if unknown_unicast is not None:
            pulumi.set(__self__, "unknown_unicast", unknown_unicast)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def broadcast(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "broadcast")

    @broadcast.setter
    def broadcast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "broadcast", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rate(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "rate")

    @rate.setter
    def rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate", value)

    @property
    @pulumi.getter(name="stormControlMode")
    def storm_control_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "storm_control_mode")

    @storm_control_mode.setter
    def storm_control_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storm_control_mode", value)

    @property
    @pulumi.getter(name="unknownMulticast")
    def unknown_multicast(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "unknown_multicast")

    @unknown_multicast.setter
    def unknown_multicast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unknown_multicast", value)

    @property
    @pulumi.getter(name="unknownUnicast")
    def unknown_unicast(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "unknown_unicast")

    @unknown_unicast.setter
    def unknown_unicast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unknown_unicast", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _SwitchControllerStormControlPolicyState:
    def __init__(__self__, *,
                 broadcast: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[int]] = None,
                 storm_control_mode: Optional[pulumi.Input[str]] = None,
                 unknown_multicast: Optional[pulumi.Input[str]] = None,
                 unknown_unicast: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerStormControlPolicy resources.
        """
        if broadcast is not None:
            pulumi.set(__self__, "broadcast", broadcast)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rate is not None:
            pulumi.set(__self__, "rate", rate)
        if storm_control_mode is not None:
            pulumi.set(__self__, "storm_control_mode", storm_control_mode)
        if unknown_multicast is not None:
            pulumi.set(__self__, "unknown_multicast", unknown_multicast)
        if unknown_unicast is not None:
            pulumi.set(__self__, "unknown_unicast", unknown_unicast)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def broadcast(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "broadcast")

    @broadcast.setter
    def broadcast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "broadcast", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def rate(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "rate")

    @rate.setter
    def rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate", value)

    @property
    @pulumi.getter(name="stormControlMode")
    def storm_control_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "storm_control_mode")

    @storm_control_mode.setter
    def storm_control_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storm_control_mode", value)

    @property
    @pulumi.getter(name="unknownMulticast")
    def unknown_multicast(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "unknown_multicast")

    @unknown_multicast.setter
    def unknown_multicast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unknown_multicast", value)

    @property
    @pulumi.getter(name="unknownUnicast")
    def unknown_unicast(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "unknown_unicast")

    @unknown_unicast.setter
    def unknown_unicast(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unknown_unicast", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class SwitchControllerStormControlPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broadcast: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[int]] = None,
                 storm_control_mode: Optional[pulumi.Input[str]] = None,
                 unknown_multicast: Optional[pulumi.Input[str]] = None,
                 unknown_unicast: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SwitchControllerStormControlPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerStormControlPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SwitchControllerStormControlPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SwitchControllerStormControlPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerStormControlPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 broadcast: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[int]] = None,
                 storm_control_mode: Optional[pulumi.Input[str]] = None,
                 unknown_multicast: Optional[pulumi.Input[str]] = None,
                 unknown_unicast: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchControllerStormControlPolicyArgs.__new__(SwitchControllerStormControlPolicyArgs)

            __props__.__dict__["broadcast"] = broadcast
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["rate"] = rate
            __props__.__dict__["storm_control_mode"] = storm_control_mode
            __props__.__dict__["unknown_multicast"] = unknown_multicast
            __props__.__dict__["unknown_unicast"] = unknown_unicast
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerStormControlPolicy, __self__).__init__(
            'fortios:index/switchControllerStormControlPolicy:SwitchControllerStormControlPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            broadcast: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rate: Optional[pulumi.Input[int]] = None,
            storm_control_mode: Optional[pulumi.Input[str]] = None,
            unknown_multicast: Optional[pulumi.Input[str]] = None,
            unknown_unicast: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerStormControlPolicy':
        """
        Get an existing SwitchControllerStormControlPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerStormControlPolicyState.__new__(_SwitchControllerStormControlPolicyState)

        __props__.__dict__["broadcast"] = broadcast
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["rate"] = rate
        __props__.__dict__["storm_control_mode"] = storm_control_mode
        __props__.__dict__["unknown_multicast"] = unknown_multicast
        __props__.__dict__["unknown_unicast"] = unknown_unicast
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerStormControlPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def broadcast(self) -> pulumi.Output[str]:
        return pulumi.get(self, "broadcast")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rate(self) -> pulumi.Output[int]:
        return pulumi.get(self, "rate")

    @property
    @pulumi.getter(name="stormControlMode")
    def storm_control_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "storm_control_mode")

    @property
    @pulumi.getter(name="unknownMulticast")
    def unknown_multicast(self) -> pulumi.Output[str]:
        return pulumi.get(self, "unknown_multicast")

    @property
    @pulumi.getter(name="unknownUnicast")
    def unknown_unicast(self) -> pulumi.Output[str]:
        return pulumi.get(self, "unknown_unicast")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

