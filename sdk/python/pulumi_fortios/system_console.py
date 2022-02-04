# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemConsoleArgs', 'SystemConsole']

@pulumi.input_type
class SystemConsoleArgs:
    def __init__(__self__, *,
                 baudrate: Optional[pulumi.Input[str]] = None,
                 fortiexplorer: Optional[pulumi.Input[str]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 output: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemConsole resource.
        :param pulumi.Input[str] baudrate: Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
        :param pulumi.Input[str] fortiexplorer: Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] login: Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] mode: Console mode. Valid values: `batch`, `line`.
        :param pulumi.Input[str] output: Console output mode. Valid values: `standard`, `more`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if baudrate is not None:
            pulumi.set(__self__, "baudrate", baudrate)
        if fortiexplorer is not None:
            pulumi.set(__self__, "fortiexplorer", fortiexplorer)
        if login is not None:
            pulumi.set(__self__, "login", login)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if output is not None:
            pulumi.set(__self__, "output", output)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def baudrate(self) -> Optional[pulumi.Input[str]]:
        """
        Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
        """
        return pulumi.get(self, "baudrate")

    @baudrate.setter
    def baudrate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "baudrate", value)

    @property
    @pulumi.getter
    def fortiexplorer(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiexplorer")

    @fortiexplorer.setter
    def fortiexplorer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiexplorer", value)

    @property
    @pulumi.getter
    def login(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Console mode. Valid values: `batch`, `line`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def output(self) -> Optional[pulumi.Input[str]]:
        """
        Console output mode. Valid values: `standard`, `more`.
        """
        return pulumi.get(self, "output")

    @output.setter
    def output(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output", value)

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
class _SystemConsoleState:
    def __init__(__self__, *,
                 baudrate: Optional[pulumi.Input[str]] = None,
                 fortiexplorer: Optional[pulumi.Input[str]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 output: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemConsole resources.
        :param pulumi.Input[str] baudrate: Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
        :param pulumi.Input[str] fortiexplorer: Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] login: Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] mode: Console mode. Valid values: `batch`, `line`.
        :param pulumi.Input[str] output: Console output mode. Valid values: `standard`, `more`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if baudrate is not None:
            pulumi.set(__self__, "baudrate", baudrate)
        if fortiexplorer is not None:
            pulumi.set(__self__, "fortiexplorer", fortiexplorer)
        if login is not None:
            pulumi.set(__self__, "login", login)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if output is not None:
            pulumi.set(__self__, "output", output)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def baudrate(self) -> Optional[pulumi.Input[str]]:
        """
        Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
        """
        return pulumi.get(self, "baudrate")

    @baudrate.setter
    def baudrate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "baudrate", value)

    @property
    @pulumi.getter
    def fortiexplorer(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiexplorer")

    @fortiexplorer.setter
    def fortiexplorer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fortiexplorer", value)

    @property
    @pulumi.getter
    def login(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Console mode. Valid values: `batch`, `line`.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def output(self) -> Optional[pulumi.Input[str]]:
        """
        Console output mode. Valid values: `standard`, `more`.
        """
        return pulumi.get(self, "output")

    @output.setter
    def output(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output", value)

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


class SystemConsole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 baudrate: Optional[pulumi.Input[str]] = None,
                 fortiexplorer: Optional[pulumi.Input[str]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 output: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure console.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemConsole("trname",
            baudrate="9600",
            login="enable",
            mode="line",
            output="more")
        ```

        ## Import

        System Console can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemConsole:SystemConsole labelname SystemConsole
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] baudrate: Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
        :param pulumi.Input[str] fortiexplorer: Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] login: Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] mode: Console mode. Valid values: `batch`, `line`.
        :param pulumi.Input[str] output: Console output mode. Valid values: `standard`, `more`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemConsoleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure console.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SystemConsole("trname",
            baudrate="9600",
            login="enable",
            mode="line",
            output="more")
        ```

        ## Import

        System Console can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/systemConsole:SystemConsole labelname SystemConsole
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemConsoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemConsoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 baudrate: Optional[pulumi.Input[str]] = None,
                 fortiexplorer: Optional[pulumi.Input[str]] = None,
                 login: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 output: Optional[pulumi.Input[str]] = None,
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
            __props__ = SystemConsoleArgs.__new__(SystemConsoleArgs)

            __props__.__dict__["baudrate"] = baudrate
            __props__.__dict__["fortiexplorer"] = fortiexplorer
            __props__.__dict__["login"] = login
            __props__.__dict__["mode"] = mode
            __props__.__dict__["output"] = output
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemConsole, __self__).__init__(
            'fortios:index/systemConsole:SystemConsole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            baudrate: Optional[pulumi.Input[str]] = None,
            fortiexplorer: Optional[pulumi.Input[str]] = None,
            login: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            output: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemConsole':
        """
        Get an existing SystemConsole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] baudrate: Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
        :param pulumi.Input[str] fortiexplorer: Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] login: Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] mode: Console mode. Valid values: `batch`, `line`.
        :param pulumi.Input[str] output: Console output mode. Valid values: `standard`, `more`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemConsoleState.__new__(_SystemConsoleState)

        __props__.__dict__["baudrate"] = baudrate
        __props__.__dict__["fortiexplorer"] = fortiexplorer
        __props__.__dict__["login"] = login
        __props__.__dict__["mode"] = mode
        __props__.__dict__["output"] = output
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemConsole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def baudrate(self) -> pulumi.Output[str]:
        """
        Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
        """
        return pulumi.get(self, "baudrate")

    @property
    @pulumi.getter
    def fortiexplorer(self) -> pulumi.Output[str]:
        """
        Enable/disable access for FortiExplorer. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "fortiexplorer")

    @property
    @pulumi.getter
    def login(self) -> pulumi.Output[str]:
        """
        Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "login")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        Console mode. Valid values: `batch`, `line`.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def output(self) -> pulumi.Output[str]:
        """
        Console output mode. Valid values: `standard`, `more`.
        """
        return pulumi.get(self, "output")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

