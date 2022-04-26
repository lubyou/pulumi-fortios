# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SwitchControllerRemoteLogArgs', 'SwitchControllerRemoteLog']

@pulumi.input_type
class SwitchControllerRemoteLogArgs:
    def __init__(__self__, *,
                 csv: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchControllerRemoteLog resource.
        :param pulumi.Input[str] csv: Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] facility: Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        :param pulumi.Input[str] name: Remote log name.
        :param pulumi.Input[int] port: Remote syslog server listening port.
        :param pulumi.Input[str] server: IPv4 address of the remote syslog server.
        :param pulumi.Input[str] severity: Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        :param pulumi.Input[str] status: Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if csv is not None:
            pulumi.set(__self__, "csv", csv)
        if facility is not None:
            pulumi.set(__self__, "facility", facility)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def csv(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "csv")

    @csv.setter
    def csv(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "csv", value)

    @property
    @pulumi.getter
    def facility(self) -> Optional[pulumi.Input[str]]:
        """
        Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        """
        return pulumi.get(self, "facility")

    @facility.setter
    def facility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "facility", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Remote log name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Remote syslog server listening port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of the remote syslog server.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
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


@pulumi.input_type
class _SwitchControllerRemoteLogState:
    def __init__(__self__, *,
                 csv: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchControllerRemoteLog resources.
        :param pulumi.Input[str] csv: Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] facility: Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        :param pulumi.Input[str] name: Remote log name.
        :param pulumi.Input[int] port: Remote syslog server listening port.
        :param pulumi.Input[str] server: IPv4 address of the remote syslog server.
        :param pulumi.Input[str] severity: Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        :param pulumi.Input[str] status: Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if csv is not None:
            pulumi.set(__self__, "csv", csv)
        if facility is not None:
            pulumi.set(__self__, "facility", facility)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if severity is not None:
            pulumi.set(__self__, "severity", severity)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def csv(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "csv")

    @csv.setter
    def csv(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "csv", value)

    @property
    @pulumi.getter
    def facility(self) -> Optional[pulumi.Input[str]]:
        """
        Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        """
        return pulumi.get(self, "facility")

    @facility.setter
    def facility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "facility", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Remote log name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Remote syslog server listening port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of the remote syslog server.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter
    def severity(self) -> Optional[pulumi.Input[str]]:
        """
        Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        """
        return pulumi.get(self, "severity")

    @severity.setter
    def severity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "severity", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
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


class SwitchControllerRemoteLog(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 csv: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure logging by FortiSwitch device to a remote syslog server. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController RemoteLog can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/switchControllerRemoteLog:SwitchControllerRemoteLog labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/switchControllerRemoteLog:SwitchControllerRemoteLog labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] csv: Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] facility: Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        :param pulumi.Input[str] name: Remote log name.
        :param pulumi.Input[int] port: Remote syslog server listening port.
        :param pulumi.Input[str] server: IPv4 address of the remote syslog server.
        :param pulumi.Input[str] severity: Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        :param pulumi.Input[str] status: Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchControllerRemoteLogArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure logging by FortiSwitch device to a remote syslog server. Applies to FortiOS Version `>= 6.2.4`.

        ## Import

        SwitchController RemoteLog can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/switchControllerRemoteLog:SwitchControllerRemoteLog labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/switchControllerRemoteLog:SwitchControllerRemoteLog labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchControllerRemoteLogArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchControllerRemoteLogArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 csv: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 severity: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
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
            __props__ = SwitchControllerRemoteLogArgs.__new__(SwitchControllerRemoteLogArgs)

            __props__.__dict__["csv"] = csv
            __props__.__dict__["facility"] = facility
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["server"] = server
            __props__.__dict__["severity"] = severity
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchControllerRemoteLog, __self__).__init__(
            'fortios:index/switchControllerRemoteLog:SwitchControllerRemoteLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            csv: Optional[pulumi.Input[str]] = None,
            facility: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            server: Optional[pulumi.Input[str]] = None,
            severity: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchControllerRemoteLog':
        """
        Get an existing SwitchControllerRemoteLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] csv: Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] facility: Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        :param pulumi.Input[str] name: Remote log name.
        :param pulumi.Input[int] port: Remote syslog server listening port.
        :param pulumi.Input[str] server: IPv4 address of the remote syslog server.
        :param pulumi.Input[str] severity: Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        :param pulumi.Input[str] status: Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchControllerRemoteLogState.__new__(_SwitchControllerRemoteLogState)

        __props__.__dict__["csv"] = csv
        __props__.__dict__["facility"] = facility
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["server"] = server
        __props__.__dict__["severity"] = severity
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchControllerRemoteLog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def csv(self) -> pulumi.Output[str]:
        """
        Enable/disable comma-separated value (CSV) strings. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "csv")

    @property
    @pulumi.getter
    def facility(self) -> pulumi.Output[str]:
        """
        Facility to log to remote syslog server. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
        """
        return pulumi.get(self, "facility")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Remote log name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Remote syslog server listening port.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        """
        IPv4 address of the remote syslog server.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter
    def severity(self) -> pulumi.Output[str]:
        """
        Severity of logs to be transferred to remote log server. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
        """
        return pulumi.get(self, "severity")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable logging by FortiSwitch device to a remote syslog server. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

