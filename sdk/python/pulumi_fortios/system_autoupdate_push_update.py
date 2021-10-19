# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemAutoupdatePushUpdateArgs', 'SystemAutoupdatePushUpdate']

@pulumi.input_type
class SystemAutoupdatePushUpdateArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 override: pulumi.Input[str],
                 port: pulumi.Input[int],
                 status: pulumi.Input[str],
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemAutoupdatePushUpdate resource.
        :param pulumi.Input[str] address: Push update override server.
        :param pulumi.Input[str] override: Enable/disable push update override server. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] port: Push update override port. (Do not overlap with other service ports)
        :param pulumi.Input[str] status: Enable/disable push updates. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "override", override)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        Push update override server.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def override(self) -> pulumi.Input[str]:
        """
        Enable/disable push update override server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: pulumi.Input[str]):
        pulumi.set(self, "override", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        Push update override port. (Do not overlap with other service ports)
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        Enable/disable push updates. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

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
class _SystemAutoupdatePushUpdateState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 override: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemAutoupdatePushUpdate resources.
        :param pulumi.Input[str] address: Push update override server.
        :param pulumi.Input[str] override: Enable/disable push update override server. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] port: Push update override port. (Do not overlap with other service ports)
        :param pulumi.Input[str] status: Enable/disable push updates. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if override is not None:
            pulumi.set(__self__, "override", override)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Push update override server.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def override(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable push update override server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override")

    @override.setter
    def override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Push update override port. (Do not overlap with other service ports)
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable push updates. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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


class SystemAutoupdatePushUpdate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 override: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure push updates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemAutoupdatePushUpdate("trname",
            address="0.0.0.0",
            override="disable",
            port=9443,
            status="disable")
        ```

        ## Import

        SystemAutoupdate PushUpdate can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemAutoupdatePushUpdate:SystemAutoupdatePushUpdate labelname SystemAutoupdatePushUpdate
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Push update override server.
        :param pulumi.Input[str] override: Enable/disable push update override server. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] port: Push update override port. (Do not overlap with other service ports)
        :param pulumi.Input[str] status: Enable/disable push updates. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemAutoupdatePushUpdateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure push updates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemAutoupdatePushUpdate("trname",
            address="0.0.0.0",
            override="disable",
            port=9443,
            status="disable")
        ```

        ## Import

        SystemAutoupdate PushUpdate can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemAutoupdatePushUpdate:SystemAutoupdatePushUpdate labelname SystemAutoupdatePushUpdate
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemAutoupdatePushUpdateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemAutoupdatePushUpdateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 override: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemAutoupdatePushUpdateArgs.__new__(SystemAutoupdatePushUpdateArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            if override is None and not opts.urn:
                raise TypeError("Missing required property 'override'")
            __props__.__dict__["override"] = override
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemAutoupdatePushUpdate, __self__).__init__(
            'fortios:index/systemAutoupdatePushUpdate:SystemAutoupdatePushUpdate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            override: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemAutoupdatePushUpdate':
        """
        Get an existing SystemAutoupdatePushUpdate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Push update override server.
        :param pulumi.Input[str] override: Enable/disable push update override server. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] port: Push update override port. (Do not overlap with other service ports)
        :param pulumi.Input[str] status: Enable/disable push updates. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemAutoupdatePushUpdateState.__new__(_SystemAutoupdatePushUpdateState)

        __props__.__dict__["address"] = address
        __props__.__dict__["override"] = override
        __props__.__dict__["port"] = port
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemAutoupdatePushUpdate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Push update override server.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def override(self) -> pulumi.Output[str]:
        """
        Enable/disable push update override server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Push update override port. (Do not overlap with other service ports)
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable push updates. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

