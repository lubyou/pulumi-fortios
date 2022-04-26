# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CertificateCaArgs', 'CertificateCa']

@pulumi.input_type
class CertificateCaArgs:
    def __init__(__self__, *,
                 ca: pulumi.Input[str],
                 auto_update_days: Optional[pulumi.Input[int]] = None,
                 auto_update_days_warning: Optional[pulumi.Input[int]] = None,
                 ca_identifier: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 scep_url: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_inspection_trusted: Optional[pulumi.Input[str]] = None,
                 trusted: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CertificateCa resource.
        :param pulumi.Input[str] ca: CA certificate as a PEM file.
        :param pulumi.Input[int] auto_update_days: Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
        :param pulumi.Input[int] auto_update_days_warning: Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
        :param pulumi.Input[str] ca_identifier: CA identifier of the SCEP server.
        :param pulumi.Input[int] last_updated: Time at which CA was last updated.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] scep_url: URL of the SCEP server.
        :param pulumi.Input[str] source: CA certificate source type.
        :param pulumi.Input[str] source_ip: Source IP address for communications to the SCEP server.
        :param pulumi.Input[str] ssl_inspection_trusted: Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] trusted: Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "ca", ca)
        if auto_update_days is not None:
            pulumi.set(__self__, "auto_update_days", auto_update_days)
        if auto_update_days_warning is not None:
            pulumi.set(__self__, "auto_update_days_warning", auto_update_days_warning)
        if ca_identifier is not None:
            pulumi.set(__self__, "ca_identifier", ca_identifier)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if range is not None:
            pulumi.set(__self__, "range", range)
        if scep_url is not None:
            pulumi.set(__self__, "scep_url", scep_url)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if ssl_inspection_trusted is not None:
            pulumi.set(__self__, "ssl_inspection_trusted", ssl_inspection_trusted)
        if trusted is not None:
            pulumi.set(__self__, "trusted", trusted)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ca(self) -> pulumi.Input[str]:
        """
        CA certificate as a PEM file.
        """
        return pulumi.get(self, "ca")

    @ca.setter
    def ca(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca", value)

    @property
    @pulumi.getter(name="autoUpdateDays")
    def auto_update_days(self) -> Optional[pulumi.Input[int]]:
        """
        Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
        """
        return pulumi.get(self, "auto_update_days")

    @auto_update_days.setter
    def auto_update_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_update_days", value)

    @property
    @pulumi.getter(name="autoUpdateDaysWarning")
    def auto_update_days_warning(self) -> Optional[pulumi.Input[int]]:
        """
        Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
        """
        return pulumi.get(self, "auto_update_days_warning")

    @auto_update_days_warning.setter
    def auto_update_days_warning(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_update_days_warning", value)

    @property
    @pulumi.getter(name="caIdentifier")
    def ca_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        CA identifier of the SCEP server.
        """
        return pulumi.get(self, "ca_identifier")

    @ca_identifier.setter
    def ca_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_identifier", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[int]]:
        """
        Time at which CA was last updated.
        """
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def range(self) -> Optional[pulumi.Input[str]]:
        """
        Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @range.setter
    def range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range", value)

    @property
    @pulumi.getter(name="scepUrl")
    def scep_url(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the SCEP server.
        """
        return pulumi.get(self, "scep_url")

    @scep_url.setter
    def scep_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scep_url", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate source type.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communications to the SCEP server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sslInspectionTrusted")
    def ssl_inspection_trusted(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_inspection_trusted")

    @ssl_inspection_trusted.setter
    def ssl_inspection_trusted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_inspection_trusted", value)

    @property
    @pulumi.getter
    def trusted(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "trusted")

    @trusted.setter
    def trusted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusted", value)

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
class _CertificateCaState:
    def __init__(__self__, *,
                 auto_update_days: Optional[pulumi.Input[int]] = None,
                 auto_update_days_warning: Optional[pulumi.Input[int]] = None,
                 ca: Optional[pulumi.Input[str]] = None,
                 ca_identifier: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 scep_url: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_inspection_trusted: Optional[pulumi.Input[str]] = None,
                 trusted: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CertificateCa resources.
        :param pulumi.Input[int] auto_update_days: Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
        :param pulumi.Input[int] auto_update_days_warning: Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
        :param pulumi.Input[str] ca: CA certificate as a PEM file.
        :param pulumi.Input[str] ca_identifier: CA identifier of the SCEP server.
        :param pulumi.Input[int] last_updated: Time at which CA was last updated.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] scep_url: URL of the SCEP server.
        :param pulumi.Input[str] source: CA certificate source type.
        :param pulumi.Input[str] source_ip: Source IP address for communications to the SCEP server.
        :param pulumi.Input[str] ssl_inspection_trusted: Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] trusted: Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if auto_update_days is not None:
            pulumi.set(__self__, "auto_update_days", auto_update_days)
        if auto_update_days_warning is not None:
            pulumi.set(__self__, "auto_update_days_warning", auto_update_days_warning)
        if ca is not None:
            pulumi.set(__self__, "ca", ca)
        if ca_identifier is not None:
            pulumi.set(__self__, "ca_identifier", ca_identifier)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if range is not None:
            pulumi.set(__self__, "range", range)
        if scep_url is not None:
            pulumi.set(__self__, "scep_url", scep_url)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)
        if ssl_inspection_trusted is not None:
            pulumi.set(__self__, "ssl_inspection_trusted", ssl_inspection_trusted)
        if trusted is not None:
            pulumi.set(__self__, "trusted", trusted)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="autoUpdateDays")
    def auto_update_days(self) -> Optional[pulumi.Input[int]]:
        """
        Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
        """
        return pulumi.get(self, "auto_update_days")

    @auto_update_days.setter
    def auto_update_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_update_days", value)

    @property
    @pulumi.getter(name="autoUpdateDaysWarning")
    def auto_update_days_warning(self) -> Optional[pulumi.Input[int]]:
        """
        Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
        """
        return pulumi.get(self, "auto_update_days_warning")

    @auto_update_days_warning.setter
    def auto_update_days_warning(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_update_days_warning", value)

    @property
    @pulumi.getter
    def ca(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate as a PEM file.
        """
        return pulumi.get(self, "ca")

    @ca.setter
    def ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca", value)

    @property
    @pulumi.getter(name="caIdentifier")
    def ca_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        CA identifier of the SCEP server.
        """
        return pulumi.get(self, "ca_identifier")

    @ca_identifier.setter
    def ca_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_identifier", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[int]]:
        """
        Time at which CA was last updated.
        """
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def range(self) -> Optional[pulumi.Input[str]]:
        """
        Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @range.setter
    def range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range", value)

    @property
    @pulumi.getter(name="scepUrl")
    def scep_url(self) -> Optional[pulumi.Input[str]]:
        """
        URL of the SCEP server.
        """
        return pulumi.get(self, "scep_url")

    @scep_url.setter
    def scep_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scep_url", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate source type.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Source IP address for communications to the SCEP server.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)

    @property
    @pulumi.getter(name="sslInspectionTrusted")
    def ssl_inspection_trusted(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_inspection_trusted")

    @ssl_inspection_trusted.setter
    def ssl_inspection_trusted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_inspection_trusted", value)

    @property
    @pulumi.getter
    def trusted(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "trusted")

    @trusted.setter
    def trusted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusted", value)

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


class CertificateCa(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_update_days: Optional[pulumi.Input[int]] = None,
                 auto_update_days_warning: Optional[pulumi.Input[int]] = None,
                 ca: Optional[pulumi.Input[str]] = None,
                 ca_identifier: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 scep_url: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_inspection_trusted: Optional[pulumi.Input[str]] = None,
                 trusted: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        CA certificate.

        ## Import

        Certificate Ca can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/certificateCa:CertificateCa labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/certificateCa:CertificateCa labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] auto_update_days: Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
        :param pulumi.Input[int] auto_update_days_warning: Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
        :param pulumi.Input[str] ca: CA certificate as a PEM file.
        :param pulumi.Input[str] ca_identifier: CA identifier of the SCEP server.
        :param pulumi.Input[int] last_updated: Time at which CA was last updated.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] scep_url: URL of the SCEP server.
        :param pulumi.Input[str] source: CA certificate source type.
        :param pulumi.Input[str] source_ip: Source IP address for communications to the SCEP server.
        :param pulumi.Input[str] ssl_inspection_trusted: Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] trusted: Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateCaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        CA certificate.

        ## Import

        Certificate Ca can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/certificateCa:CertificateCa labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/certificateCa:CertificateCa labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param CertificateCaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateCaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_update_days: Optional[pulumi.Input[int]] = None,
                 auto_update_days_warning: Optional[pulumi.Input[int]] = None,
                 ca: Optional[pulumi.Input[str]] = None,
                 ca_identifier: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 range: Optional[pulumi.Input[str]] = None,
                 scep_url: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 ssl_inspection_trusted: Optional[pulumi.Input[str]] = None,
                 trusted: Optional[pulumi.Input[str]] = None,
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
            __props__ = CertificateCaArgs.__new__(CertificateCaArgs)

            __props__.__dict__["auto_update_days"] = auto_update_days
            __props__.__dict__["auto_update_days_warning"] = auto_update_days_warning
            if ca is None and not opts.urn:
                raise TypeError("Missing required property 'ca'")
            __props__.__dict__["ca"] = ca
            __props__.__dict__["ca_identifier"] = ca_identifier
            __props__.__dict__["last_updated"] = last_updated
            __props__.__dict__["name"] = name
            __props__.__dict__["range"] = range
            __props__.__dict__["scep_url"] = scep_url
            __props__.__dict__["source"] = source
            __props__.__dict__["source_ip"] = source_ip
            __props__.__dict__["ssl_inspection_trusted"] = ssl_inspection_trusted
            __props__.__dict__["trusted"] = trusted
            __props__.__dict__["vdomparam"] = vdomparam
        super(CertificateCa, __self__).__init__(
            'fortios:index/certificateCa:CertificateCa',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_update_days: Optional[pulumi.Input[int]] = None,
            auto_update_days_warning: Optional[pulumi.Input[int]] = None,
            ca: Optional[pulumi.Input[str]] = None,
            ca_identifier: Optional[pulumi.Input[str]] = None,
            last_updated: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            range: Optional[pulumi.Input[str]] = None,
            scep_url: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            source_ip: Optional[pulumi.Input[str]] = None,
            ssl_inspection_trusted: Optional[pulumi.Input[str]] = None,
            trusted: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'CertificateCa':
        """
        Get an existing CertificateCa resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] auto_update_days: Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
        :param pulumi.Input[int] auto_update_days_warning: Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
        :param pulumi.Input[str] ca: CA certificate as a PEM file.
        :param pulumi.Input[str] ca_identifier: CA identifier of the SCEP server.
        :param pulumi.Input[int] last_updated: Time at which CA was last updated.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] range: Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
        :param pulumi.Input[str] scep_url: URL of the SCEP server.
        :param pulumi.Input[str] source: CA certificate source type.
        :param pulumi.Input[str] source_ip: Source IP address for communications to the SCEP server.
        :param pulumi.Input[str] ssl_inspection_trusted: Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] trusted: Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CertificateCaState.__new__(_CertificateCaState)

        __props__.__dict__["auto_update_days"] = auto_update_days
        __props__.__dict__["auto_update_days_warning"] = auto_update_days_warning
        __props__.__dict__["ca"] = ca
        __props__.__dict__["ca_identifier"] = ca_identifier
        __props__.__dict__["last_updated"] = last_updated
        __props__.__dict__["name"] = name
        __props__.__dict__["range"] = range
        __props__.__dict__["scep_url"] = scep_url
        __props__.__dict__["source"] = source
        __props__.__dict__["source_ip"] = source_ip
        __props__.__dict__["ssl_inspection_trusted"] = ssl_inspection_trusted
        __props__.__dict__["trusted"] = trusted
        __props__.__dict__["vdomparam"] = vdomparam
        return CertificateCa(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoUpdateDays")
    def auto_update_days(self) -> pulumi.Output[int]:
        """
        Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
        """
        return pulumi.get(self, "auto_update_days")

    @property
    @pulumi.getter(name="autoUpdateDaysWarning")
    def auto_update_days_warning(self) -> pulumi.Output[int]:
        """
        Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
        """
        return pulumi.get(self, "auto_update_days_warning")

    @property
    @pulumi.getter
    def ca(self) -> pulumi.Output[str]:
        """
        CA certificate as a PEM file.
        """
        return pulumi.get(self, "ca")

    @property
    @pulumi.getter(name="caIdentifier")
    def ca_identifier(self) -> pulumi.Output[str]:
        """
        CA identifier of the SCEP server.
        """
        return pulumi.get(self, "ca_identifier")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> pulumi.Output[int]:
        """
        Time at which CA was last updated.
        """
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def range(self) -> pulumi.Output[str]:
        """
        Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
        """
        return pulumi.get(self, "range")

    @property
    @pulumi.getter(name="scepUrl")
    def scep_url(self) -> pulumi.Output[str]:
        """
        URL of the SCEP server.
        """
        return pulumi.get(self, "scep_url")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        CA certificate source type.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[str]:
        """
        Source IP address for communications to the SCEP server.
        """
        return pulumi.get(self, "source_ip")

    @property
    @pulumi.getter(name="sslInspectionTrusted")
    def ssl_inspection_trusted(self) -> pulumi.Output[str]:
        """
        Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_inspection_trusted")

    @property
    @pulumi.getter
    def trusted(self) -> pulumi.Output[str]:
        """
        Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "trusted")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

