# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FortimanagerDVMDeviceArgs', 'FortimanagerDVMDevice']

@pulumi.input_type
class FortimanagerDVMDeviceArgs:
    def __init__(__self__, *,
                 device_name: pulumi.Input[str],
                 ipaddr: pulumi.Input[str],
                 userid: pulumi.Input[str],
                 adom: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FortimanagerDVMDevice resource.
        """
        pulumi.set(__self__, "device_name", device_name)
        pulumi.set(__self__, "ipaddr", ipaddr)
        pulumi.set(__self__, "userid", userid)
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter
    def ipaddr(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ipaddr")

    @ipaddr.setter
    def ipaddr(self, value: pulumi.Input[str]):
        pulumi.set(self, "ipaddr", value)

    @property
    @pulumi.getter
    def userid(self) -> pulumi.Input[str]:
        return pulumi.get(self, "userid")

    @userid.setter
    def userid(self, value: pulumi.Input[str]):
        pulumi.set(self, "userid", value)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


@pulumi.input_type
class _FortimanagerDVMDeviceState:
    def __init__(__self__, *,
                 adom: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 ipaddr: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 userid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FortimanagerDVMDevice resources.
        """
        if adom is not None:
            pulumi.set(__self__, "adom", adom)
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if ipaddr is not None:
            pulumi.set(__self__, "ipaddr", ipaddr)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if userid is not None:
            pulumi.set(__self__, "userid", userid)

    @property
    @pulumi.getter
    def adom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "adom")

    @adom.setter
    def adom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "adom", value)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter
    def ipaddr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipaddr")

    @ipaddr.setter
    def ipaddr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipaddr", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def userid(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "userid")

    @userid.setter
    def userid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "userid", value)


class FortimanagerDVMDevice(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 ipaddr: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 userid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FortimanagerDVMDevice resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FortimanagerDVMDeviceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FortimanagerDVMDevice resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FortimanagerDVMDeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortimanagerDVMDeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adom: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 ipaddr: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 userid: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortimanagerDVMDeviceArgs.__new__(FortimanagerDVMDeviceArgs)

            __props__.__dict__["adom"] = adom
            if device_name is None and not opts.urn:
                raise TypeError("Missing required property 'device_name'")
            __props__.__dict__["device_name"] = device_name
            if ipaddr is None and not opts.urn:
                raise TypeError("Missing required property 'ipaddr'")
            __props__.__dict__["ipaddr"] = ipaddr
            __props__.__dict__["password"] = password
            if userid is None and not opts.urn:
                raise TypeError("Missing required property 'userid'")
            __props__.__dict__["userid"] = userid
        super(FortimanagerDVMDevice, __self__).__init__(
            'fortios:index/fortimanagerDVMDevice:FortimanagerDVMDevice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            adom: Optional[pulumi.Input[str]] = None,
            device_name: Optional[pulumi.Input[str]] = None,
            ipaddr: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            userid: Optional[pulumi.Input[str]] = None) -> 'FortimanagerDVMDevice':
        """
        Get an existing FortimanagerDVMDevice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortimanagerDVMDeviceState.__new__(_FortimanagerDVMDeviceState)

        __props__.__dict__["adom"] = adom
        __props__.__dict__["device_name"] = device_name
        __props__.__dict__["ipaddr"] = ipaddr
        __props__.__dict__["password"] = password
        __props__.__dict__["userid"] = userid
        return FortimanagerDVMDevice(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def adom(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "adom")

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "device_name")

    @property
    @pulumi.getter
    def ipaddr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ipaddr")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def userid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "userid")

