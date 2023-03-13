# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FortimanagerSystemNTPArgs', 'FortimanagerSystemNTP']

@pulumi.input_type
class FortimanagerSystemNTPArgs:
    def __init__(__self__, *,
                 server: pulumi.Input[str],
                 status: Optional[pulumi.Input[str]] = None,
                 sync_interval: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a FortimanagerSystemNTP resource.
        """
        pulumi.set(__self__, "server", server)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if sync_interval is not None:
            pulumi.set(__self__, "sync_interval", sync_interval)

    @property
    @pulumi.getter
    def server(self) -> pulumi.Input[str]:
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: pulumi.Input[str]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="syncInterval")
    def sync_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "sync_interval")

    @sync_interval.setter
    def sync_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sync_interval", value)


@pulumi.input_type
class _FortimanagerSystemNTPState:
    def __init__(__self__, *,
                 server: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 sync_interval: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering FortimanagerSystemNTP resources.
        """
        if server is not None:
            pulumi.set(__self__, "server", server)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if sync_interval is not None:
            pulumi.set(__self__, "sync_interval", sync_interval)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="syncInterval")
    def sync_interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "sync_interval")

    @sync_interval.setter
    def sync_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sync_interval", value)


class FortimanagerSystemNTP(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 sync_interval: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a FortimanagerSystemNTP resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FortimanagerSystemNTPArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FortimanagerSystemNTP resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FortimanagerSystemNTPArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortimanagerSystemNTPArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 sync_interval: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortimanagerSystemNTPArgs.__new__(FortimanagerSystemNTPArgs)

            if server is None and not opts.urn:
                raise TypeError("Missing required property 'server'")
            __props__.__dict__["server"] = server
            __props__.__dict__["status"] = status
            __props__.__dict__["sync_interval"] = sync_interval
        super(FortimanagerSystemNTP, __self__).__init__(
            'fortios:index/fortimanagerSystemNTP:FortimanagerSystemNTP',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            server: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            sync_interval: Optional[pulumi.Input[int]] = None) -> 'FortimanagerSystemNTP':
        """
        Get an existing FortimanagerSystemNTP resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortimanagerSystemNTPState.__new__(_FortimanagerSystemNTPState)

        __props__.__dict__["server"] = server
        __props__.__dict__["status"] = status
        __props__.__dict__["sync_interval"] = sync_interval
        return FortimanagerSystemNTP(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="syncInterval")
    def sync_interval(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "sync_interval")

