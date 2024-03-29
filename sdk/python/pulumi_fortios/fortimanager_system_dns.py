# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FortimanagerSystemDNSArgs', 'FortimanagerSystemDNS']

@pulumi.input_type
class FortimanagerSystemDNSArgs:
    def __init__(__self__, *,
                 primary: Optional[pulumi.Input[str]] = None,
                 secondary: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FortimanagerSystemDNS resource.
        """
        if primary is not None:
            pulumi.set(__self__, "primary", primary)
        if secondary is not None:
            pulumi.set(__self__, "secondary", secondary)

    @property
    @pulumi.getter
    def primary(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "primary")

    @primary.setter
    def primary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary", value)

    @property
    @pulumi.getter
    def secondary(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secondary")

    @secondary.setter
    def secondary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary", value)


@pulumi.input_type
class _FortimanagerSystemDNSState:
    def __init__(__self__, *,
                 primary: Optional[pulumi.Input[str]] = None,
                 secondary: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FortimanagerSystemDNS resources.
        """
        if primary is not None:
            pulumi.set(__self__, "primary", primary)
        if secondary is not None:
            pulumi.set(__self__, "secondary", secondary)

    @property
    @pulumi.getter
    def primary(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "primary")

    @primary.setter
    def primary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary", value)

    @property
    @pulumi.getter
    def secondary(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secondary")

    @secondary.setter
    def secondary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary", value)


class FortimanagerSystemDNS(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 primary: Optional[pulumi.Input[str]] = None,
                 secondary: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FortimanagerSystemDNS resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortimanagerSystemDNSArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FortimanagerSystemDNS resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FortimanagerSystemDNSArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortimanagerSystemDNSArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 primary: Optional[pulumi.Input[str]] = None,
                 secondary: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortimanagerSystemDNSArgs.__new__(FortimanagerSystemDNSArgs)

            __props__.__dict__["primary"] = primary
            __props__.__dict__["secondary"] = secondary
        super(FortimanagerSystemDNS, __self__).__init__(
            'fortios:index/fortimanagerSystemDNS:FortimanagerSystemDNS',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            primary: Optional[pulumi.Input[str]] = None,
            secondary: Optional[pulumi.Input[str]] = None) -> 'FortimanagerSystemDNS':
        """
        Get an existing FortimanagerSystemDNS resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortimanagerSystemDNSState.__new__(_FortimanagerSystemDNSState)

        __props__.__dict__["primary"] = primary
        __props__.__dict__["secondary"] = secondary
        return FortimanagerSystemDNS(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def primary(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "primary")

    @property
    @pulumi.getter
    def secondary(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "secondary")

