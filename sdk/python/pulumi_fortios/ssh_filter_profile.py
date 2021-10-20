# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SshFilterProfileArgs', 'SshFilterProfile']

@pulumi.input_type
class SshFilterProfileArgs:
    def __init__(__self__, *,
                 block: Optional[pulumi.Input[str]] = None,
                 default_command_log: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 file_filter: Optional[pulumi.Input['SshFilterProfileFileFilterArgs']] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 shell_commands: Optional[pulumi.Input[Sequence[pulumi.Input['SshFilterProfileShellCommandArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SshFilterProfile resource.
        :param pulumi.Input[str] block: SSH blocking options.
        :param pulumi.Input[str] default_command_log: Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input['SshFilterProfileFileFilterArgs'] file_filter: File filter. The structure of `file_filter` block is documented below.
        :param pulumi.Input[str] log: Enable/disable file filter logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: File type name.
        :param pulumi.Input[Sequence[pulumi.Input['SshFilterProfileShellCommandArgs']]] shell_commands: SSH command filter. The structure of `shell_commands` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if block is not None:
            pulumi.set(__self__, "block", block)
        if default_command_log is not None:
            pulumi.set(__self__, "default_command_log", default_command_log)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if file_filter is not None:
            pulumi.set(__self__, "file_filter", file_filter)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if shell_commands is not None:
            pulumi.set(__self__, "shell_commands", shell_commands)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def block(self) -> Optional[pulumi.Input[str]]:
        """
        SSH blocking options.
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "block", value)

    @property
    @pulumi.getter(name="defaultCommandLog")
    def default_command_log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "default_command_log")

    @default_command_log.setter
    def default_command_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_command_log", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="fileFilter")
    def file_filter(self) -> Optional[pulumi.Input['SshFilterProfileFileFilterArgs']]:
        """
        File filter. The structure of `file_filter` block is documented below.
        """
        return pulumi.get(self, "file_filter")

    @file_filter.setter
    def file_filter(self, value: Optional[pulumi.Input['SshFilterProfileFileFilterArgs']]):
        pulumi.set(self, "file_filter", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable file filter logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        File type name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="shellCommands")
    def shell_commands(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SshFilterProfileShellCommandArgs']]]]:
        """
        SSH command filter. The structure of `shell_commands` block is documented below.
        """
        return pulumi.get(self, "shell_commands")

    @shell_commands.setter
    def shell_commands(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SshFilterProfileShellCommandArgs']]]]):
        pulumi.set(self, "shell_commands", value)

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
class _SshFilterProfileState:
    def __init__(__self__, *,
                 block: Optional[pulumi.Input[str]] = None,
                 default_command_log: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 file_filter: Optional[pulumi.Input['SshFilterProfileFileFilterArgs']] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 shell_commands: Optional[pulumi.Input[Sequence[pulumi.Input['SshFilterProfileShellCommandArgs']]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SshFilterProfile resources.
        :param pulumi.Input[str] block: SSH blocking options.
        :param pulumi.Input[str] default_command_log: Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input['SshFilterProfileFileFilterArgs'] file_filter: File filter. The structure of `file_filter` block is documented below.
        :param pulumi.Input[str] log: Enable/disable file filter logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: File type name.
        :param pulumi.Input[Sequence[pulumi.Input['SshFilterProfileShellCommandArgs']]] shell_commands: SSH command filter. The structure of `shell_commands` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if block is not None:
            pulumi.set(__self__, "block", block)
        if default_command_log is not None:
            pulumi.set(__self__, "default_command_log", default_command_log)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if file_filter is not None:
            pulumi.set(__self__, "file_filter", file_filter)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if shell_commands is not None:
            pulumi.set(__self__, "shell_commands", shell_commands)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def block(self) -> Optional[pulumi.Input[str]]:
        """
        SSH blocking options.
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "block", value)

    @property
    @pulumi.getter(name="defaultCommandLog")
    def default_command_log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "default_command_log")

    @default_command_log.setter
    def default_command_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_command_log", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="fileFilter")
    def file_filter(self) -> Optional[pulumi.Input['SshFilterProfileFileFilterArgs']]:
        """
        File filter. The structure of `file_filter` block is documented below.
        """
        return pulumi.get(self, "file_filter")

    @file_filter.setter
    def file_filter(self, value: Optional[pulumi.Input['SshFilterProfileFileFilterArgs']]):
        pulumi.set(self, "file_filter", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable file filter logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        File type name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="shellCommands")
    def shell_commands(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SshFilterProfileShellCommandArgs']]]]:
        """
        SSH command filter. The structure of `shell_commands` block is documented below.
        """
        return pulumi.get(self, "shell_commands")

    @shell_commands.setter
    def shell_commands(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SshFilterProfileShellCommandArgs']]]]):
        pulumi.set(self, "shell_commands", value)

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


class SshFilterProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block: Optional[pulumi.Input[str]] = None,
                 default_command_log: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 file_filter: Optional[pulumi.Input[pulumi.InputType['SshFilterProfileFileFilterArgs']]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 shell_commands: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SshFilterProfileShellCommandArgs']]]]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        SSH filter profile.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SshFilterProfile("trname",
            block="x11",
            default_command_log="enable",
            log="x11")
        ```

        ## Import

        SshFilter Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/sshFilterProfile:SshFilterProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] block: SSH blocking options.
        :param pulumi.Input[str] default_command_log: Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[pulumi.InputType['SshFilterProfileFileFilterArgs']] file_filter: File filter. The structure of `file_filter` block is documented below.
        :param pulumi.Input[str] log: Enable/disable file filter logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: File type name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SshFilterProfileShellCommandArgs']]]] shell_commands: SSH command filter. The structure of `shell_commands` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SshFilterProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        SSH filter profile.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.SshFilterProfile("trname",
            block="x11",
            default_command_log="enable",
            log="x11")
        ```

        ## Import

        SshFilter Profile can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/sshFilterProfile:SshFilterProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SshFilterProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SshFilterProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block: Optional[pulumi.Input[str]] = None,
                 default_command_log: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 file_filter: Optional[pulumi.Input[pulumi.InputType['SshFilterProfileFileFilterArgs']]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 shell_commands: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SshFilterProfileShellCommandArgs']]]]] = None,
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
            __props__ = SshFilterProfileArgs.__new__(SshFilterProfileArgs)

            __props__.__dict__["block"] = block
            __props__.__dict__["default_command_log"] = default_command_log
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["file_filter"] = file_filter
            __props__.__dict__["log"] = log
            __props__.__dict__["name"] = name
            __props__.__dict__["shell_commands"] = shell_commands
            __props__.__dict__["vdomparam"] = vdomparam
        super(SshFilterProfile, __self__).__init__(
            'fortios:index/sshFilterProfile:SshFilterProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            block: Optional[pulumi.Input[str]] = None,
            default_command_log: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            file_filter: Optional[pulumi.Input[pulumi.InputType['SshFilterProfileFileFilterArgs']]] = None,
            log: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            shell_commands: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SshFilterProfileShellCommandArgs']]]]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SshFilterProfile':
        """
        Get an existing SshFilterProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] block: SSH blocking options.
        :param pulumi.Input[str] default_command_log: Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] dynamic_sort_subtable: true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        :param pulumi.Input[pulumi.InputType['SshFilterProfileFileFilterArgs']] file_filter: File filter. The structure of `file_filter` block is documented below.
        :param pulumi.Input[str] log: Enable/disable file filter logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: File type name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SshFilterProfileShellCommandArgs']]]] shell_commands: SSH command filter. The structure of `shell_commands` block is documented below.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SshFilterProfileState.__new__(_SshFilterProfileState)

        __props__.__dict__["block"] = block
        __props__.__dict__["default_command_log"] = default_command_log
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["file_filter"] = file_filter
        __props__.__dict__["log"] = log
        __props__.__dict__["name"] = name
        __props__.__dict__["shell_commands"] = shell_commands
        __props__.__dict__["vdomparam"] = vdomparam
        return SshFilterProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def block(self) -> pulumi.Output[str]:
        """
        SSH blocking options.
        """
        return pulumi.get(self, "block")

    @property
    @pulumi.getter(name="defaultCommandLog")
    def default_command_log(self) -> pulumi.Output[str]:
        """
        Enable/disable logging unmatched shell commands. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "default_command_log")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="fileFilter")
    def file_filter(self) -> pulumi.Output[Optional['outputs.SshFilterProfileFileFilter']]:
        """
        File filter. The structure of `file_filter` block is documented below.
        """
        return pulumi.get(self, "file_filter")

    @property
    @pulumi.getter
    def log(self) -> pulumi.Output[str]:
        """
        Enable/disable file filter logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        File type name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="shellCommands")
    def shell_commands(self) -> pulumi.Output[Optional[Sequence['outputs.SshFilterProfileShellCommand']]]:
        """
        SSH command filter. The structure of `shell_commands` block is documented below.
        """
        return pulumi.get(self, "shell_commands")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")
