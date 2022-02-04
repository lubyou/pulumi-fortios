# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LogGuiDisplayArgs', 'LogGuiDisplay']

@pulumi.input_type
class LogGuiDisplayArgs:
    def __init__(__self__, *,
                 fortiview_unscanned_apps: Optional[pulumi.Input[str]] = None,
                 resolve_apps: Optional[pulumi.Input[str]] = None,
                 resolve_hosts: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogGuiDisplay resource.
        :param pulumi.Input[str] fortiview_unscanned_apps: Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] resolve_apps: Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] resolve_hosts: Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fortiview_unscanned_apps is not None:
            pulumi.set(__self__, "fortiview_unscanned_apps", fortiview_unscanned_apps)
        if resolve_apps is not None:
            pulumi.set(__self__, "resolve_apps", resolve_apps)
        if resolve_hosts is not None:
            pulumi.set(__self__, "resolve_hosts", resolve_hosts)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="fortiviewUnscannedApps")
    def fortiview_unscanned_apps(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiview_unscanned_apps")

    @fortiview_unscanned_apps.setter
    def fortiview_unscanned_apps(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiview_unscanned_apps", value)

    @property
    @pulumi.getter(name="resolveApps")
    def resolve_apps(self) -> Optional[pulumi.Input[str]]:
        """
        Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "resolve_apps")

    @resolve_apps.setter
    def resolve_apps(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resolve_apps", value)

    @property
    @pulumi.getter(name="resolveHosts")
    def resolve_hosts(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "resolve_hosts")

    @resolve_hosts.setter
    def resolve_hosts(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resolve_hosts", value)

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
class _LogGuiDisplayState:
    def __init__(__self__, *,
                 fortiview_unscanned_apps: Optional[pulumi.Input[str]] = None,
                 resolve_apps: Optional[pulumi.Input[str]] = None,
                 resolve_hosts: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogGuiDisplay resources.
        :param pulumi.Input[str] fortiview_unscanned_apps: Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] resolve_apps: Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] resolve_hosts: Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fortiview_unscanned_apps is not None:
            pulumi.set(__self__, "fortiview_unscanned_apps", fortiview_unscanned_apps)
        if resolve_apps is not None:
            pulumi.set(__self__, "resolve_apps", resolve_apps)
        if resolve_hosts is not None:
            pulumi.set(__self__, "resolve_hosts", resolve_hosts)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="fortiviewUnscannedApps")
    def fortiview_unscanned_apps(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiview_unscanned_apps")

    @fortiview_unscanned_apps.setter
    def fortiview_unscanned_apps(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiview_unscanned_apps", value)

    @property
    @pulumi.getter(name="resolveApps")
    def resolve_apps(self) -> Optional[pulumi.Input[str]]:
        """
        Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "resolve_apps")

    @resolve_apps.setter
    def resolve_apps(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resolve_apps", value)

    @property
    @pulumi.getter(name="resolveHosts")
    def resolve_hosts(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "resolve_hosts")

    @resolve_hosts.setter
    def resolve_hosts(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resolve_hosts", value)

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


class LogGuiDisplay(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fortiview_unscanned_apps: Optional[pulumi.Input[str]] = None,
                 resolve_apps: Optional[pulumi.Input[str]] = None,
                 resolve_hosts: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure how log messages are displayed on the GUI.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.LogGuiDisplay("trname",
            fortiview_unscanned_apps="disable",
            resolve_apps="enable",
            resolve_hosts="enable")
        ```

        ## Import

        Log GuiDisplay can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/logGuiDisplay:LogGuiDisplay labelname LogGuiDisplay
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fortiview_unscanned_apps: Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] resolve_apps: Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] resolve_hosts: Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LogGuiDisplayArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure how log messages are displayed on the GUI.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.LogGuiDisplay("trname",
            fortiview_unscanned_apps="disable",
            resolve_apps="enable",
            resolve_hosts="enable")
        ```

        ## Import

        Log GuiDisplay can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/logGuiDisplay:LogGuiDisplay labelname LogGuiDisplay
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param LogGuiDisplayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogGuiDisplayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fortiview_unscanned_apps: Optional[pulumi.Input[str]] = None,
                 resolve_apps: Optional[pulumi.Input[str]] = None,
                 resolve_hosts: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
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
            __props__ = LogGuiDisplayArgs.__new__(LogGuiDisplayArgs)

            __props__.__dict__["fortiview_unscanned_apps"] = fortiview_unscanned_apps
            __props__.__dict__["resolve_apps"] = resolve_apps
            __props__.__dict__["resolve_hosts"] = resolve_hosts
            __props__.__dict__["vdomparam"] = vdomparam
        super(LogGuiDisplay, __self__).__init__(
            'fortios:index/logGuiDisplay:LogGuiDisplay',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fortiview_unscanned_apps: Optional[pulumi.Input[str]] = None,
            resolve_apps: Optional[pulumi.Input[str]] = None,
            resolve_hosts: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'LogGuiDisplay':
        """
        Get an existing LogGuiDisplay resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fortiview_unscanned_apps: Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] resolve_apps: Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] resolve_hosts: Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogGuiDisplayState.__new__(_LogGuiDisplayState)

        __props__.__dict__["fortiview_unscanned_apps"] = fortiview_unscanned_apps
        __props__.__dict__["resolve_apps"] = resolve_apps
        __props__.__dict__["resolve_hosts"] = resolve_hosts
        __props__.__dict__["vdomparam"] = vdomparam
        return LogGuiDisplay(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fortiviewUnscannedApps")
    def fortiview_unscanned_apps(self) -> pulumi.Output[str]:
        """
        Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiview_unscanned_apps")

    @property
    @pulumi.getter(name="resolveApps")
    def resolve_apps(self) -> pulumi.Output[str]:
        """
        Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "resolve_apps")

    @property
    @pulumi.getter(name="resolveHosts")
    def resolve_hosts(self) -> pulumi.Output[str]:
        """
        Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "resolve_hosts")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

