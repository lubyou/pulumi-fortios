# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LogSyslogdSettingArgs', 'LogSyslogdSetting']

@pulumi.input_type
class LogSyslogdSettingArgs:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input[str]] = None,
                 custom_field_names: Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdSettingCustomFieldNameArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 max_log_rate: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 syslog_type: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogSyslogdSetting resource.
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if custom_field_names is not None:
            pulumi.set(__self__, "custom_field_names", custom_field_names)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if enc_algorithm is not None:
            pulumi.set(__self__, "enc_algorithm", enc_algorithm)
        if facility is not None:
            pulumi.set(__self__, "facility", facility)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if max_log_rate is not None:
            pulumi.set(__self__, "max_log_rate", max_log_rate)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if syslog_type is not None:
            pulumi.set(__self__, "syslog_type", syslog_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="customFieldNames")
    def custom_field_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdSettingCustomFieldNameArgs']]]]:
        return pulumi.get(self, "custom_field_names")

    @custom_field_names.setter
    def custom_field_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdSettingCustomFieldNameArgs']]]]):
        pulumi.set(self, "custom_field_names", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "enc_algorithm")

    @enc_algorithm.setter
    def enc_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enc_algorithm", value)

    @property
    @pulumi.getter
    def facility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "facility")

    @facility.setter
    def facility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "facility", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter(name="maxLogRate")
    def max_log_rate(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_log_rate")

    @max_log_rate.setter
    def max_log_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_log_rate", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="syslogType")
    def syslog_type(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "syslog_type")

    @syslog_type.setter
    def syslog_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "syslog_type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _LogSyslogdSettingState:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input[str]] = None,
                 custom_field_names: Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdSettingCustomFieldNameArgs']]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 max_log_rate: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 syslog_type: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogSyslogdSetting resources.
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if custom_field_names is not None:
            pulumi.set(__self__, "custom_field_names", custom_field_names)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if enc_algorithm is not None:
            pulumi.set(__self__, "enc_algorithm", enc_algorithm)
        if facility is not None:
            pulumi.set(__self__, "facility", facility)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if get_all_tables is not None:
            pulumi.set(__self__, "get_all_tables", get_all_tables)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if interface_select_method is not None:
            pulumi.set(__self__, "interface_select_method", interface_select_method)
        if max_log_rate is not None:
            pulumi.set(__self__, "max_log_rate", max_log_rate)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if ssl_min_proto_version is not None:
            pulumi.set(__self__, "ssl_min_proto_version", ssl_min_proto_version)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if syslog_type is not None:
            pulumi.set(__self__, "syslog_type", syslog_type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="customFieldNames")
    def custom_field_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdSettingCustomFieldNameArgs']]]]:
        return pulumi.get(self, "custom_field_names")

    @custom_field_names.setter
    def custom_field_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LogSyslogdSettingCustomFieldNameArgs']]]]):
        pulumi.set(self, "custom_field_names", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "enc_algorithm")

    @enc_algorithm.setter
    def enc_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enc_algorithm", value)

    @property
    @pulumi.getter
    def facility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "facility")

    @facility.setter
    def facility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "facility", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "get_all_tables")

    @get_all_tables.setter
    def get_all_tables(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "get_all_tables", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interface_select_method")

    @interface_select_method.setter
    def interface_select_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_select_method", value)

    @property
    @pulumi.getter(name="maxLogRate")
    def max_log_rate(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_log_rate")

    @max_log_rate.setter
    def max_log_rate(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_log_rate", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ssl_min_proto_version")

    @ssl_min_proto_version.setter
    def ssl_min_proto_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_proto_version", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="syslogType")
    def syslog_type(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "syslog_type")

    @syslog_type.setter
    def syslog_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "syslog_type", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class LogSyslogdSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 custom_field_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogSyslogdSettingCustomFieldNameArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 max_log_rate: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 syslog_type: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LogSyslogdSetting resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[LogSyslogdSettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LogSyslogdSetting resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LogSyslogdSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogSyslogdSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 custom_field_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogSyslogdSettingCustomFieldNameArgs']]]]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 facility: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 get_all_tables: Optional[pulumi.Input[str]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 interface_select_method: Optional[pulumi.Input[str]] = None,
                 max_log_rate: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 syslog_type: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogSyslogdSettingArgs.__new__(LogSyslogdSettingArgs)

            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["custom_field_names"] = custom_field_names
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["enc_algorithm"] = enc_algorithm
            __props__.__dict__["facility"] = facility
            __props__.__dict__["format"] = format
            __props__.__dict__["get_all_tables"] = get_all_tables
            __props__.__dict__["interface"] = interface
            __props__.__dict__["interface_select_method"] = interface_select_method
            __props__.__dict__["max_log_rate"] = max_log_rate
            __props__.__dict__["mode"] = mode
            __props__.__dict__["port"] = port
            __props__.__dict__["priority"] = priority
            __props__.__dict__["server"] = server
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
            __props__.__dict__["status"] = status
            __props__.__dict__["syslog_type"] = syslog_type
            __props__.__dict__["vdomparam"] = vdomparam
        super(LogSyslogdSetting, __self__).__init__(
            'fortios:index/logSyslogdSetting:LogSyslogdSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            custom_field_names: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LogSyslogdSettingCustomFieldNameArgs']]]]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            enc_algorithm: Optional[pulumi.Input[str]] = None,
            facility: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            get_all_tables: Optional[pulumi.Input[str]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            interface_select_method: Optional[pulumi.Input[str]] = None,
            max_log_rate: Optional[pulumi.Input[int]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            priority: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            ssl_min_proto_version: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            syslog_type: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'LogSyslogdSetting':
        """
        Get an existing LogSyslogdSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogSyslogdSettingState.__new__(_LogSyslogdSettingState)

        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["custom_field_names"] = custom_field_names
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["enc_algorithm"] = enc_algorithm
        __props__.__dict__["facility"] = facility
        __props__.__dict__["format"] = format
        __props__.__dict__["get_all_tables"] = get_all_tables
        __props__.__dict__["interface"] = interface
        __props__.__dict__["interface_select_method"] = interface_select_method
        __props__.__dict__["max_log_rate"] = max_log_rate
        __props__.__dict__["mode"] = mode
        __props__.__dict__["port"] = port
        __props__.__dict__["priority"] = priority
        __props__.__dict__["server"] = server
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["ssl_min_proto_version"] = ssl_min_proto_version
        __props__.__dict__["status"] = status
        __props__.__dict__["syslog_type"] = syslog_type
        __props__.__dict__["vdomparam"] = vdomparam
        return LogSyslogdSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="customFieldNames")
    def custom_field_names(self) -> pulumi.Output[Optional[Sequence['outputs.LogSyslogdSettingCustomFieldName']]]:
        return pulumi.get(self, "custom_field_names")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> pulumi.Output[str]:
        return pulumi.get(self, "enc_algorithm")

    @property
    @pulumi.getter
    def facility(self) -> pulumi.Output[str]:
        return pulumi.get(self, "facility")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[str]:
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="getAllTables")
    def get_all_tables(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "get_all_tables")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="interfaceSelectMethod")
    def interface_select_method(self) -> pulumi.Output[str]:
        return pulumi.get(self, "interface_select_method")

    @property
    @pulumi.getter(name="maxLogRate")
    def max_log_rate(self) -> pulumi.Output[int]:
        return pulumi.get(self, "max_log_rate")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[str]:
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sslMinProtoVersion")
    def ssl_min_proto_version(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ssl_min_proto_version")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="syslogType")
    def syslog_type(self) -> pulumi.Output[int]:
        return pulumi.get(self, "syslog_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

