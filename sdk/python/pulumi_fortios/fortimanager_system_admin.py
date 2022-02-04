# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FortimanagerSystemAdminArgs', 'FortimanagerSystemAdmin']

@pulumi.input_type
class FortimanagerSystemAdminArgs:
    def __init__(__self__, *,
                 http_port: Optional[pulumi.Input[int]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a FortimanagerSystemAdmin resource.
        :param pulumi.Input[int] http_port: Http port.
        :param pulumi.Input[int] https_port: Https port.
        :param pulumi.Input[int] idle_timeout: Idle Timeout(1-480 minute).
        """
        if http_port is not None:
            pulumi.set(__self__, "http_port", http_port)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)
        if idle_timeout is not None:
            pulumi.set(__self__, "idle_timeout", idle_timeout)

    @property
    @pulumi.getter(name="httpPort")
    def http_port(self) -> Optional[pulumi.Input[int]]:
        """
        Http port.
        """
        return pulumi.get(self, "http_port")

    @http_port.setter
    def http_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_port", value)

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> Optional[pulumi.Input[int]]:
        """
        Https port.
        """
        return pulumi.get(self, "https_port")

    @https_port.setter
    def https_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "https_port", value)

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Idle Timeout(1-480 minute).
        """
        return pulumi.get(self, "idle_timeout")

    @idle_timeout.setter
    def idle_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "idle_timeout", value)


@pulumi.input_type
class _FortimanagerSystemAdminState:
    def __init__(__self__, *,
                 http_port: Optional[pulumi.Input[int]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering FortimanagerSystemAdmin resources.
        :param pulumi.Input[int] http_port: Http port.
        :param pulumi.Input[int] https_port: Https port.
        :param pulumi.Input[int] idle_timeout: Idle Timeout(1-480 minute).
        """
        if http_port is not None:
            pulumi.set(__self__, "http_port", http_port)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)
        if idle_timeout is not None:
            pulumi.set(__self__, "idle_timeout", idle_timeout)

    @property
    @pulumi.getter(name="httpPort")
    def http_port(self) -> Optional[pulumi.Input[int]]:
        """
        Http port.
        """
        return pulumi.get(self, "http_port")

    @http_port.setter
    def http_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_port", value)

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> Optional[pulumi.Input[int]]:
        """
        Https port.
        """
        return pulumi.get(self, "https_port")

    @https_port.setter
    def https_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "https_port", value)

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Idle Timeout(1-480 minute).
        """
        return pulumi.get(self, "idle_timeout")

    @idle_timeout.setter
    def idle_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "idle_timeout", value)


class FortimanagerSystemAdmin(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_port: Optional[pulumi.Input[int]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        This resource supports modifying system admin setting for FortiManager.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test1 = fortios.FortimanagerSystemAdmin("test1",
            http_port=80,
            https_port=443,
            idle_timeout=20)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] http_port: Http port.
        :param pulumi.Input[int] https_port: Https port.
        :param pulumi.Input[int] idle_timeout: Idle Timeout(1-480 minute).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FortimanagerSystemAdminArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports modifying system admin setting for FortiManager.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        test1 = fortios.FortimanagerSystemAdmin("test1",
            http_port=80,
            https_port=443,
            idle_timeout=20)
        ```

        :param str resource_name: The name of the resource.
        :param FortimanagerSystemAdminArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FortimanagerSystemAdminArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_port: Optional[pulumi.Input[int]] = None,
                 https_port: Optional[pulumi.Input[int]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FortimanagerSystemAdminArgs.__new__(FortimanagerSystemAdminArgs)

            __props__.__dict__["http_port"] = http_port
            __props__.__dict__["https_port"] = https_port
            __props__.__dict__["idle_timeout"] = idle_timeout
        super(FortimanagerSystemAdmin, __self__).__init__(
            'fortios:index/fortimanagerSystemAdmin:FortimanagerSystemAdmin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            http_port: Optional[pulumi.Input[int]] = None,
            https_port: Optional[pulumi.Input[int]] = None,
            idle_timeout: Optional[pulumi.Input[int]] = None) -> 'FortimanagerSystemAdmin':
        """
        Get an existing FortimanagerSystemAdmin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] http_port: Http port.
        :param pulumi.Input[int] https_port: Https port.
        :param pulumi.Input[int] idle_timeout: Idle Timeout(1-480 minute).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FortimanagerSystemAdminState.__new__(_FortimanagerSystemAdminState)

        __props__.__dict__["http_port"] = http_port
        __props__.__dict__["https_port"] = https_port
        __props__.__dict__["idle_timeout"] = idle_timeout
        return FortimanagerSystemAdmin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="httpPort")
    def http_port(self) -> pulumi.Output[Optional[int]]:
        """
        Http port.
        """
        return pulumi.get(self, "http_port")

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> pulumi.Output[Optional[int]]:
        """
        Https port.
        """
        return pulumi.get(self, "https_port")

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Idle Timeout(1-480 minute).
        """
        return pulumi.get(self, "idle_timeout")

