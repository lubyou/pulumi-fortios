# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['LogFortiAnalyzerSettingLegacyArgs', 'LogFortiAnalyzerSettingLegacy']

@pulumi.input_type
class LogFortiAnalyzerSettingLegacyArgs:
    def __init__(__self__, *,
                 status: pulumi.Input[str],
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 hmac_algorithm: Optional[pulumi.Input[str]] = None,
                 reliable: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 upload_option: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogFortiAnalyzerSettingLegacy resource.
        """
        pulumi.set(__self__, "status", status)
        if enc_algorithm is not None:
            pulumi.set(__self__, "enc_algorithm", enc_algorithm)
        if hmac_algorithm is not None:
            pulumi.set(__self__, "hmac_algorithm", hmac_algorithm)
        if reliable is not None:
            pulumi.set(__self__, "reliable", reliable)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if upload_option is not None:
            pulumi.set(__self__, "upload_option", upload_option)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "enc_algorithm")

    @enc_algorithm.setter
    def enc_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enc_algorithm", value)

    @property
    @pulumi.getter(name="hmacAlgorithm")
    def hmac_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hmac_algorithm")

    @hmac_algorithm.setter
    def hmac_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hmac_algorithm", value)

    @property
    @pulumi.getter
    def reliable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "reliable")

    @reliable.setter
    def reliable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reliable", value)

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
    @pulumi.getter(name="uploadOption")
    def upload_option(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "upload_option")

    @upload_option.setter
    def upload_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_option", value)


@pulumi.input_type
class _LogFortiAnalyzerSettingLegacyState:
    def __init__(__self__, *,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 hmac_algorithm: Optional[pulumi.Input[str]] = None,
                 reliable: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 upload_option: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogFortiAnalyzerSettingLegacy resources.
        """
        if enc_algorithm is not None:
            pulumi.set(__self__, "enc_algorithm", enc_algorithm)
        if hmac_algorithm is not None:
            pulumi.set(__self__, "hmac_algorithm", hmac_algorithm)
        if reliable is not None:
            pulumi.set(__self__, "reliable", reliable)
        if server is not None:
            pulumi.set(__self__, "server", server)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if upload_option is not None:
            pulumi.set(__self__, "upload_option", upload_option)

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "enc_algorithm")

    @enc_algorithm.setter
    def enc_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enc_algorithm", value)

    @property
    @pulumi.getter(name="hmacAlgorithm")
    def hmac_algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hmac_algorithm")

    @hmac_algorithm.setter
    def hmac_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hmac_algorithm", value)

    @property
    @pulumi.getter
    def reliable(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "reliable")

    @reliable.setter
    def reliable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reliable", value)

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
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="uploadOption")
    def upload_option(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "upload_option")

    @upload_option.setter
    def upload_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upload_option", value)


class LogFortiAnalyzerSettingLegacy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 hmac_algorithm: Optional[pulumi.Input[str]] = None,
                 reliable: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 upload_option: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LogFortiAnalyzerSettingLegacy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogFortiAnalyzerSettingLegacyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LogFortiAnalyzerSettingLegacy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LogFortiAnalyzerSettingLegacyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogFortiAnalyzerSettingLegacyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enc_algorithm: Optional[pulumi.Input[str]] = None,
                 hmac_algorithm: Optional[pulumi.Input[str]] = None,
                 reliable: Optional[pulumi.Input[str]] = None,
                 server: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 upload_option: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogFortiAnalyzerSettingLegacyArgs.__new__(LogFortiAnalyzerSettingLegacyArgs)

            __props__.__dict__["enc_algorithm"] = enc_algorithm
            __props__.__dict__["hmac_algorithm"] = hmac_algorithm
            __props__.__dict__["reliable"] = reliable
            __props__.__dict__["server"] = server
            __props__.__dict__["source_ip"] = source_ip
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["upload_option"] = upload_option
        super(LogFortiAnalyzerSettingLegacy, __self__).__init__(
            'fortios:index/logFortiAnalyzerSettingLegacy:LogFortiAnalyzerSettingLegacy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enc_algorithm: Optional[pulumi.Input[str]] = None,
            hmac_algorithm: Optional[pulumi.Input[str]] = None,
            reliable: Optional[pulumi.Input[str]] = None,
            server: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            upload_option: Optional[pulumi.Input[str]] = None) -> 'LogFortiAnalyzerSettingLegacy':
        """
        Get an existing LogFortiAnalyzerSettingLegacy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogFortiAnalyzerSettingLegacyState.__new__(_LogFortiAnalyzerSettingLegacyState)

        __props__.__dict__["enc_algorithm"] = enc_algorithm
        __props__.__dict__["hmac_algorithm"] = hmac_algorithm
        __props__.__dict__["reliable"] = reliable
        __props__.__dict__["server"] = server
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["status"] = status
        __props__.__dict__["upload_option"] = upload_option
        return LogFortiAnalyzerSettingLegacy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="encAlgorithm")
    def enc_algorithm(self) -> pulumi.Output[str]:
        return pulumi.get(self, "enc_algorithm")

    @property
    @pulumi.getter(name="hmacAlgorithm")
    def hmac_algorithm(self) -> pulumi.Output[str]:
        return pulumi.get(self, "hmac_algorithm")

    @property
    @pulumi.getter
    def reliable(self) -> pulumi.Output[str]:
        return pulumi.get(self, "reliable")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[str]:
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="uploadOption")
    def upload_option(self) -> pulumi.Output[str]:
        return pulumi.get(self, "upload_option")

