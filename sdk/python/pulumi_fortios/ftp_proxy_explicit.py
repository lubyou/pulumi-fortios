# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FtpProxyExplicitArgs', 'FtpProxyExplicit']

@pulumi.input_type
class FtpProxyExplicitArgs:
    def __init__(__self__, *,
                 incoming_ip: Optional[pulumi.Input[str]] = None,
                 incoming_port: Optional[pulumi.Input[str]] = None,
                 outgoing_ip: Optional[pulumi.Input[str]] = None,
                 sec_default_action: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[str]] = None,
                 ssl_algorithm: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FtpProxyExplicit resource.
        :param pulumi.Input[str] incoming_ip: Accept incoming FTP requests from this IP address. An interface must have this IP address.
        :param pulumi.Input[str] incoming_port: Accept incoming FTP requests on one or more ports.
        :param pulumi.Input[str] outgoing_ip: Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        :param pulumi.Input[str] sec_default_action: Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        :param pulumi.Input[str] ssl: Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssl_algorithm: Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        :param pulumi.Input[str] ssl_cert: Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] status: Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if incoming_ip is not None:
            pulumi.set(__self__, "incoming_ip", incoming_ip)
        if incoming_port is not None:
            pulumi.set(__self__, "incoming_port", incoming_port)
        if outgoing_ip is not None:
            pulumi.set(__self__, "outgoing_ip", outgoing_ip)
        if sec_default_action is not None:
            pulumi.set(__self__, "sec_default_action", sec_default_action)
        if ssl is not None:
            pulumi.set(__self__, "ssl", ssl)
        if ssl_algorithm is not None:
            pulumi.set(__self__, "ssl_algorithm", ssl_algorithm)
        if ssl_cert is not None:
            pulumi.set(__self__, "ssl_cert", ssl_cert)
        if ssl_dh_bits is not None:
            pulumi.set(__self__, "ssl_dh_bits", ssl_dh_bits)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="incomingIp")
    def incoming_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Accept incoming FTP requests from this IP address. An interface must have this IP address.
        """
        return pulumi.get(self, "incoming_ip")

    @incoming_ip.setter
    def incoming_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incoming_ip", value)

    @property
    @pulumi.getter(name="incomingPort")
    def incoming_port(self) -> Optional[pulumi.Input[str]]:
        """
        Accept incoming FTP requests on one or more ports.
        """
        return pulumi.get(self, "incoming_port")

    @incoming_port.setter
    def incoming_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incoming_port", value)

    @property
    @pulumi.getter(name="outgoingIp")
    def outgoing_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        """
        return pulumi.get(self, "outgoing_ip")

    @outgoing_ip.setter
    def outgoing_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outgoing_ip", value)

    @property
    @pulumi.getter(name="secDefaultAction")
    def sec_default_action(self) -> Optional[pulumi.Input[str]]:
        """
        Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        """
        return pulumi.get(self, "sec_default_action")

    @sec_default_action.setter
    def sec_default_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sec_default_action", value)

    @property
    @pulumi.getter
    def ssl(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl")

    @ssl.setter
    def ssl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl", value)

    @property
    @pulumi.getter(name="sslAlgorithm")
    def ssl_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        """
        return pulumi.get(self, "ssl_algorithm")

    @ssl_algorithm.setter
    def ssl_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_algorithm", value)

    @property
    @pulumi.getter(name="sslCert")
    def ssl_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        """
        return pulumi.get(self, "ssl_cert")

    @ssl_cert.setter
    def ssl_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_cert", value)

    @property
    @pulumi.getter(name="sslDhBits")
    def ssl_dh_bits(self) -> Optional[pulumi.Input[str]]:
        """
        Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        """
        return pulumi.get(self, "ssl_dh_bits")

    @ssl_dh_bits.setter
    def ssl_dh_bits(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_dh_bits", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
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
class _FtpProxyExplicitState:
    def __init__(__self__, *,
                 incoming_ip: Optional[pulumi.Input[str]] = None,
                 incoming_port: Optional[pulumi.Input[str]] = None,
                 outgoing_ip: Optional[pulumi.Input[str]] = None,
                 sec_default_action: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[str]] = None,
                 ssl_algorithm: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FtpProxyExplicit resources.
        :param pulumi.Input[str] incoming_ip: Accept incoming FTP requests from this IP address. An interface must have this IP address.
        :param pulumi.Input[str] incoming_port: Accept incoming FTP requests on one or more ports.
        :param pulumi.Input[str] outgoing_ip: Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        :param pulumi.Input[str] sec_default_action: Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        :param pulumi.Input[str] ssl: Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssl_algorithm: Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        :param pulumi.Input[str] ssl_cert: Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] status: Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if incoming_ip is not None:
            pulumi.set(__self__, "incoming_ip", incoming_ip)
        if incoming_port is not None:
            pulumi.set(__self__, "incoming_port", incoming_port)
        if outgoing_ip is not None:
            pulumi.set(__self__, "outgoing_ip", outgoing_ip)
        if sec_default_action is not None:
            pulumi.set(__self__, "sec_default_action", sec_default_action)
        if ssl is not None:
            pulumi.set(__self__, "ssl", ssl)
        if ssl_algorithm is not None:
            pulumi.set(__self__, "ssl_algorithm", ssl_algorithm)
        if ssl_cert is not None:
            pulumi.set(__self__, "ssl_cert", ssl_cert)
        if ssl_dh_bits is not None:
            pulumi.set(__self__, "ssl_dh_bits", ssl_dh_bits)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="incomingIp")
    def incoming_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Accept incoming FTP requests from this IP address. An interface must have this IP address.
        """
        return pulumi.get(self, "incoming_ip")

    @incoming_ip.setter
    def incoming_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incoming_ip", value)

    @property
    @pulumi.getter(name="incomingPort")
    def incoming_port(self) -> Optional[pulumi.Input[str]]:
        """
        Accept incoming FTP requests on one or more ports.
        """
        return pulumi.get(self, "incoming_port")

    @incoming_port.setter
    def incoming_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "incoming_port", value)

    @property
    @pulumi.getter(name="outgoingIp")
    def outgoing_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        """
        return pulumi.get(self, "outgoing_ip")

    @outgoing_ip.setter
    def outgoing_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "outgoing_ip", value)

    @property
    @pulumi.getter(name="secDefaultAction")
    def sec_default_action(self) -> Optional[pulumi.Input[str]]:
        """
        Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        """
        return pulumi.get(self, "sec_default_action")

    @sec_default_action.setter
    def sec_default_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sec_default_action", value)

    @property
    @pulumi.getter
    def ssl(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl")

    @ssl.setter
    def ssl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl", value)

    @property
    @pulumi.getter(name="sslAlgorithm")
    def ssl_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        """
        return pulumi.get(self, "ssl_algorithm")

    @ssl_algorithm.setter
    def ssl_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_algorithm", value)

    @property
    @pulumi.getter(name="sslCert")
    def ssl_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        """
        return pulumi.get(self, "ssl_cert")

    @ssl_cert.setter
    def ssl_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_cert", value)

    @property
    @pulumi.getter(name="sslDhBits")
    def ssl_dh_bits(self) -> Optional[pulumi.Input[str]]:
        """
        Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        """
        return pulumi.get(self, "ssl_dh_bits")

    @ssl_dh_bits.setter
    def ssl_dh_bits(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_dh_bits", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
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


class FtpProxyExplicit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 incoming_ip: Optional[pulumi.Input[str]] = None,
                 incoming_port: Optional[pulumi.Input[str]] = None,
                 outgoing_ip: Optional[pulumi.Input[str]] = None,
                 sec_default_action: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[str]] = None,
                 ssl_algorithm: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure explicit FTP proxy settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FtpProxyExplicit("trname",
            incoming_ip="0.0.0.0",
            sec_default_action="deny",
            status="disable")
        ```

        ## Import

        FtpProxy Explicit can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/ftpProxyExplicit:FtpProxyExplicit labelname FtpProxyExplicit
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] incoming_ip: Accept incoming FTP requests from this IP address. An interface must have this IP address.
        :param pulumi.Input[str] incoming_port: Accept incoming FTP requests on one or more ports.
        :param pulumi.Input[str] outgoing_ip: Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        :param pulumi.Input[str] sec_default_action: Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        :param pulumi.Input[str] ssl: Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssl_algorithm: Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        :param pulumi.Input[str] ssl_cert: Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] status: Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FtpProxyExplicitArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure explicit FTP proxy settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FtpProxyExplicit("trname",
            incoming_ip="0.0.0.0",
            sec_default_action="deny",
            status="disable")
        ```

        ## Import

        FtpProxy Explicit can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/ftpProxyExplicit:FtpProxyExplicit labelname FtpProxyExplicit
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FtpProxyExplicitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FtpProxyExplicitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 incoming_ip: Optional[pulumi.Input[str]] = None,
                 incoming_port: Optional[pulumi.Input[str]] = None,
                 outgoing_ip: Optional[pulumi.Input[str]] = None,
                 sec_default_action: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[str]] = None,
                 ssl_algorithm: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
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
            __props__ = FtpProxyExplicitArgs.__new__(FtpProxyExplicitArgs)

            __props__.__dict__["incoming_ip"] = incoming_ip
            __props__.__dict__["incoming_port"] = incoming_port
            __props__.__dict__["outgoing_ip"] = outgoing_ip
            __props__.__dict__["sec_default_action"] = sec_default_action
            __props__.__dict__["ssl"] = ssl
            __props__.__dict__["ssl_algorithm"] = ssl_algorithm
            __props__.__dict__["ssl_cert"] = ssl_cert
            __props__.__dict__["ssl_dh_bits"] = ssl_dh_bits
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(FtpProxyExplicit, __self__).__init__(
            'fortios:index/ftpProxyExplicit:FtpProxyExplicit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            incoming_ip: Optional[pulumi.Input[str]] = None,
            incoming_port: Optional[pulumi.Input[str]] = None,
            outgoing_ip: Optional[pulumi.Input[str]] = None,
            sec_default_action: Optional[pulumi.Input[str]] = None,
            ssl: Optional[pulumi.Input[str]] = None,
            ssl_algorithm: Optional[pulumi.Input[str]] = None,
            ssl_cert: Optional[pulumi.Input[str]] = None,
            ssl_dh_bits: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FtpProxyExplicit':
        """
        Get an existing FtpProxyExplicit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] incoming_ip: Accept incoming FTP requests from this IP address. An interface must have this IP address.
        :param pulumi.Input[str] incoming_port: Accept incoming FTP requests on one or more ports.
        :param pulumi.Input[str] outgoing_ip: Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        :param pulumi.Input[str] sec_default_action: Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        :param pulumi.Input[str] ssl: Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ssl_algorithm: Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        :param pulumi.Input[str] ssl_cert: Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] status: Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FtpProxyExplicitState.__new__(_FtpProxyExplicitState)

        __props__.__dict__["incoming_ip"] = incoming_ip
        __props__.__dict__["incoming_port"] = incoming_port
        __props__.__dict__["outgoing_ip"] = outgoing_ip
        __props__.__dict__["sec_default_action"] = sec_default_action
        __props__.__dict__["ssl"] = ssl
        __props__.__dict__["ssl_algorithm"] = ssl_algorithm
        __props__.__dict__["ssl_cert"] = ssl_cert
        __props__.__dict__["ssl_dh_bits"] = ssl_dh_bits
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return FtpProxyExplicit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="incomingIp")
    def incoming_ip(self) -> pulumi.Output[str]:
        """
        Accept incoming FTP requests from this IP address. An interface must have this IP address.
        """
        return pulumi.get(self, "incoming_ip")

    @property
    @pulumi.getter(name="incomingPort")
    def incoming_port(self) -> pulumi.Output[str]:
        """
        Accept incoming FTP requests on one or more ports.
        """
        return pulumi.get(self, "incoming_port")

    @property
    @pulumi.getter(name="outgoingIp")
    def outgoing_ip(self) -> pulumi.Output[str]:
        """
        Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        """
        return pulumi.get(self, "outgoing_ip")

    @property
    @pulumi.getter(name="secDefaultAction")
    def sec_default_action(self) -> pulumi.Output[str]:
        """
        Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        """
        return pulumi.get(self, "sec_default_action")

    @property
    @pulumi.getter
    def ssl(self) -> pulumi.Output[str]:
        """
        Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl")

    @property
    @pulumi.getter(name="sslAlgorithm")
    def ssl_algorithm(self) -> pulumi.Output[str]:
        """
        Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        """
        return pulumi.get(self, "ssl_algorithm")

    @property
    @pulumi.getter(name="sslCert")
    def ssl_cert(self) -> pulumi.Output[str]:
        """
        Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        """
        return pulumi.get(self, "ssl_cert")

    @property
    @pulumi.getter(name="sslDhBits")
    def ssl_dh_bits(self) -> pulumi.Output[str]:
        """
        Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        """
        return pulumi.get(self, "ssl_dh_bits")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

