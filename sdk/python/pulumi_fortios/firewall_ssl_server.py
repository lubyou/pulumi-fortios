# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallSslServerArgs', 'FirewallSslServer']

@pulumi.input_type
class FirewallSslServerArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 port: pulumi.Input[int],
                 ssl_cert: pulumi.Input[str],
                 add_header_x_forwarded_proto: Optional[pulumi.Input[str]] = None,
                 mapped_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ssl_algorithm: Optional[pulumi.Input[str]] = None,
                 ssl_client_renegotiation: Optional[pulumi.Input[str]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 ssl_max_version: Optional[pulumi.Input[str]] = None,
                 ssl_min_version: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
                 url_rewrite: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallSslServer resource.
        :param pulumi.Input[str] ip: IPv4 address of the SSL server.
        :param pulumi.Input[int] port: Server service port (1 - 65535, default = 443).
        :param pulumi.Input[str] ssl_cert: Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        :param pulumi.Input[str] add_header_x_forwarded_proto: Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] mapped_port: Mapped server service port (1 - 65535, default = 80).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] ssl_algorithm: Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        :param pulumi.Input[str] ssl_client_renegotiation: Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] ssl_max_version: Highest SSL/TLS version to negotiate.
        :param pulumi.Input[str] ssl_min_version: Lowest SSL/TLS version to negotiate.
        :param pulumi.Input[str] ssl_mode: SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
        :param pulumi.Input[str] ssl_send_empty_frags: Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url_rewrite: Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "ssl_cert", ssl_cert)
        if add_header_x_forwarded_proto is not None:
            pulumi.set(__self__, "add_header_x_forwarded_proto", add_header_x_forwarded_proto)
        if mapped_port is not None:
            pulumi.set(__self__, "mapped_port", mapped_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ssl_algorithm is not None:
            pulumi.set(__self__, "ssl_algorithm", ssl_algorithm)
        if ssl_client_renegotiation is not None:
            pulumi.set(__self__, "ssl_client_renegotiation", ssl_client_renegotiation)
        if ssl_dh_bits is not None:
            pulumi.set(__self__, "ssl_dh_bits", ssl_dh_bits)
        if ssl_max_version is not None:
            pulumi.set(__self__, "ssl_max_version", ssl_max_version)
        if ssl_min_version is not None:
            pulumi.set(__self__, "ssl_min_version", ssl_min_version)
        if ssl_mode is not None:
            pulumi.set(__self__, "ssl_mode", ssl_mode)
        if ssl_send_empty_frags is not None:
            pulumi.set(__self__, "ssl_send_empty_frags", ssl_send_empty_frags)
        if url_rewrite is not None:
            pulumi.set(__self__, "url_rewrite", url_rewrite)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IPv4 address of the SSL server.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        Server service port (1 - 65535, default = 443).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="sslCert")
    def ssl_cert(self) -> pulumi.Input[str]:
        """
        Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        """
        return pulumi.get(self, "ssl_cert")

    @ssl_cert.setter
    def ssl_cert(self, value: pulumi.Input[str]):
        pulumi.set(self, "ssl_cert", value)

    @property
    @pulumi.getter(name="addHeaderXForwardedProto")
    def add_header_x_forwarded_proto(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "add_header_x_forwarded_proto")

    @add_header_x_forwarded_proto.setter
    def add_header_x_forwarded_proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "add_header_x_forwarded_proto", value)

    @property
    @pulumi.getter(name="mappedPort")
    def mapped_port(self) -> Optional[pulumi.Input[int]]:
        """
        Mapped server service port (1 - 65535, default = 80).
        """
        return pulumi.get(self, "mapped_port")

    @mapped_port.setter
    def mapped_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mapped_port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    @pulumi.getter(name="sslClientRenegotiation")
    def ssl_client_renegotiation(self) -> Optional[pulumi.Input[str]]:
        """
        Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
        """
        return pulumi.get(self, "ssl_client_renegotiation")

    @ssl_client_renegotiation.setter
    def ssl_client_renegotiation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_client_renegotiation", value)

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
    @pulumi.getter(name="sslMaxVersion")
    def ssl_max_version(self) -> Optional[pulumi.Input[str]]:
        """
        Highest SSL/TLS version to negotiate.
        """
        return pulumi.get(self, "ssl_max_version")

    @ssl_max_version.setter
    def ssl_max_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_max_version", value)

    @property
    @pulumi.getter(name="sslMinVersion")
    def ssl_min_version(self) -> Optional[pulumi.Input[str]]:
        """
        Lowest SSL/TLS version to negotiate.
        """
        return pulumi.get(self, "ssl_min_version")

    @ssl_min_version.setter
    def ssl_min_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_version", value)

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> Optional[pulumi.Input[str]]:
        """
        SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
        """
        return pulumi.get(self, "ssl_mode")

    @ssl_mode.setter
    def ssl_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_mode", value)

    @property
    @pulumi.getter(name="sslSendEmptyFrags")
    def ssl_send_empty_frags(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_send_empty_frags")

    @ssl_send_empty_frags.setter
    def ssl_send_empty_frags(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_send_empty_frags", value)

    @property
    @pulumi.getter(name="urlRewrite")
    def url_rewrite(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "url_rewrite")

    @url_rewrite.setter
    def url_rewrite(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_rewrite", value)

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
class _FirewallSslServerState:
    def __init__(__self__, *,
                 add_header_x_forwarded_proto: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 mapped_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 ssl_algorithm: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 ssl_client_renegotiation: Optional[pulumi.Input[str]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 ssl_max_version: Optional[pulumi.Input[str]] = None,
                 ssl_min_version: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
                 url_rewrite: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallSslServer resources.
        :param pulumi.Input[str] add_header_x_forwarded_proto: Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ip: IPv4 address of the SSL server.
        :param pulumi.Input[int] mapped_port: Mapped server service port (1 - 65535, default = 80).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[int] port: Server service port (1 - 65535, default = 443).
        :param pulumi.Input[str] ssl_algorithm: Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        :param pulumi.Input[str] ssl_cert: Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        :param pulumi.Input[str] ssl_client_renegotiation: Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] ssl_max_version: Highest SSL/TLS version to negotiate.
        :param pulumi.Input[str] ssl_min_version: Lowest SSL/TLS version to negotiate.
        :param pulumi.Input[str] ssl_mode: SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
        :param pulumi.Input[str] ssl_send_empty_frags: Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url_rewrite: Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if add_header_x_forwarded_proto is not None:
            pulumi.set(__self__, "add_header_x_forwarded_proto", add_header_x_forwarded_proto)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if mapped_port is not None:
            pulumi.set(__self__, "mapped_port", mapped_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if ssl_algorithm is not None:
            pulumi.set(__self__, "ssl_algorithm", ssl_algorithm)
        if ssl_cert is not None:
            pulumi.set(__self__, "ssl_cert", ssl_cert)
        if ssl_client_renegotiation is not None:
            pulumi.set(__self__, "ssl_client_renegotiation", ssl_client_renegotiation)
        if ssl_dh_bits is not None:
            pulumi.set(__self__, "ssl_dh_bits", ssl_dh_bits)
        if ssl_max_version is not None:
            pulumi.set(__self__, "ssl_max_version", ssl_max_version)
        if ssl_min_version is not None:
            pulumi.set(__self__, "ssl_min_version", ssl_min_version)
        if ssl_mode is not None:
            pulumi.set(__self__, "ssl_mode", ssl_mode)
        if ssl_send_empty_frags is not None:
            pulumi.set(__self__, "ssl_send_empty_frags", ssl_send_empty_frags)
        if url_rewrite is not None:
            pulumi.set(__self__, "url_rewrite", url_rewrite)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addHeaderXForwardedProto")
    def add_header_x_forwarded_proto(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "add_header_x_forwarded_proto")

    @add_header_x_forwarded_proto.setter
    def add_header_x_forwarded_proto(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "add_header_x_forwarded_proto", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of the SSL server.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="mappedPort")
    def mapped_port(self) -> Optional[pulumi.Input[int]]:
        """
        Mapped server service port (1 - 65535, default = 80).
        """
        return pulumi.get(self, "mapped_port")

    @mapped_port.setter
    def mapped_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mapped_port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Server service port (1 - 65535, default = 443).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

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
    @pulumi.getter(name="sslClientRenegotiation")
    def ssl_client_renegotiation(self) -> Optional[pulumi.Input[str]]:
        """
        Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
        """
        return pulumi.get(self, "ssl_client_renegotiation")

    @ssl_client_renegotiation.setter
    def ssl_client_renegotiation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_client_renegotiation", value)

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
    @pulumi.getter(name="sslMaxVersion")
    def ssl_max_version(self) -> Optional[pulumi.Input[str]]:
        """
        Highest SSL/TLS version to negotiate.
        """
        return pulumi.get(self, "ssl_max_version")

    @ssl_max_version.setter
    def ssl_max_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_max_version", value)

    @property
    @pulumi.getter(name="sslMinVersion")
    def ssl_min_version(self) -> Optional[pulumi.Input[str]]:
        """
        Lowest SSL/TLS version to negotiate.
        """
        return pulumi.get(self, "ssl_min_version")

    @ssl_min_version.setter
    def ssl_min_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_min_version", value)

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> Optional[pulumi.Input[str]]:
        """
        SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
        """
        return pulumi.get(self, "ssl_mode")

    @ssl_mode.setter
    def ssl_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_mode", value)

    @property
    @pulumi.getter(name="sslSendEmptyFrags")
    def ssl_send_empty_frags(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_send_empty_frags")

    @ssl_send_empty_frags.setter
    def ssl_send_empty_frags(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_send_empty_frags", value)

    @property
    @pulumi.getter(name="urlRewrite")
    def url_rewrite(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "url_rewrite")

    @url_rewrite.setter
    def url_rewrite(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_rewrite", value)

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


class FirewallSslServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_header_x_forwarded_proto: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 mapped_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 ssl_algorithm: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 ssl_client_renegotiation: Optional[pulumi.Input[str]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 ssl_max_version: Optional[pulumi.Input[str]] = None,
                 ssl_min_version: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
                 url_rewrite: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure SSL servers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallSslServer("trname",
            add_header_x_forwarded_proto="enable",
            ip="1.1.1.1",
            mapped_port=2234,
            port=32321,
            ssl_algorithm="high",
            ssl_cert="Fortinet_CA_SSL",
            ssl_client_renegotiation="allow",
            ssl_dh_bits="2048",
            ssl_max_version="tls-1.2",
            ssl_min_version="tls-1.1",
            ssl_mode="half",
            ssl_send_empty_frags="enable",
            url_rewrite="disable")
        ```

        ## Import

        Firewall SslServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallSslServer:FirewallSslServer labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] add_header_x_forwarded_proto: Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ip: IPv4 address of the SSL server.
        :param pulumi.Input[int] mapped_port: Mapped server service port (1 - 65535, default = 80).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[int] port: Server service port (1 - 65535, default = 443).
        :param pulumi.Input[str] ssl_algorithm: Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        :param pulumi.Input[str] ssl_cert: Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        :param pulumi.Input[str] ssl_client_renegotiation: Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] ssl_max_version: Highest SSL/TLS version to negotiate.
        :param pulumi.Input[str] ssl_min_version: Lowest SSL/TLS version to negotiate.
        :param pulumi.Input[str] ssl_mode: SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
        :param pulumi.Input[str] ssl_send_empty_frags: Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url_rewrite: Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallSslServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure SSL servers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallSslServer("trname",
            add_header_x_forwarded_proto="enable",
            ip="1.1.1.1",
            mapped_port=2234,
            port=32321,
            ssl_algorithm="high",
            ssl_cert="Fortinet_CA_SSL",
            ssl_client_renegotiation="allow",
            ssl_dh_bits="2048",
            ssl_max_version="tls-1.2",
            ssl_min_version="tls-1.1",
            ssl_mode="half",
            ssl_send_empty_frags="enable",
            url_rewrite="disable")
        ```

        ## Import

        Firewall SslServer can be imported using any of these accepted formats$ export "FORTIOS_IMPORT_TABLE"="true"

        ```sh
         $ pulumi import fortios:index/firewallSslServer:FirewallSslServer labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallSslServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallSslServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_header_x_forwarded_proto: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 mapped_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 ssl_algorithm: Optional[pulumi.Input[str]] = None,
                 ssl_cert: Optional[pulumi.Input[str]] = None,
                 ssl_client_renegotiation: Optional[pulumi.Input[str]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 ssl_max_version: Optional[pulumi.Input[str]] = None,
                 ssl_min_version: Optional[pulumi.Input[str]] = None,
                 ssl_mode: Optional[pulumi.Input[str]] = None,
                 ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
                 url_rewrite: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallSslServerArgs.__new__(FirewallSslServerArgs)

            __props__.__dict__["add_header_x_forwarded_proto"] = add_header_x_forwarded_proto
            if ip is None and not opts.urn:
                raise TypeError("Missing required property 'ip'")
            __props__.__dict__["ip"] = ip
            __props__.__dict__["mapped_port"] = mapped_port
            __props__.__dict__["name"] = name
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            __props__.__dict__["ssl_algorithm"] = ssl_algorithm
            if ssl_cert is None and not opts.urn:
                raise TypeError("Missing required property 'ssl_cert'")
            __props__.__dict__["ssl_cert"] = ssl_cert
            __props__.__dict__["ssl_client_renegotiation"] = ssl_client_renegotiation
            __props__.__dict__["ssl_dh_bits"] = ssl_dh_bits
            __props__.__dict__["ssl_max_version"] = ssl_max_version
            __props__.__dict__["ssl_min_version"] = ssl_min_version
            __props__.__dict__["ssl_mode"] = ssl_mode
            __props__.__dict__["ssl_send_empty_frags"] = ssl_send_empty_frags
            __props__.__dict__["url_rewrite"] = url_rewrite
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallSslServer, __self__).__init__(
            'fortios:index/firewallSslServer:FirewallSslServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_header_x_forwarded_proto: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            mapped_port: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            ssl_algorithm: Optional[pulumi.Input[str]] = None,
            ssl_cert: Optional[pulumi.Input[str]] = None,
            ssl_client_renegotiation: Optional[pulumi.Input[str]] = None,
            ssl_dh_bits: Optional[pulumi.Input[str]] = None,
            ssl_max_version: Optional[pulumi.Input[str]] = None,
            ssl_min_version: Optional[pulumi.Input[str]] = None,
            ssl_mode: Optional[pulumi.Input[str]] = None,
            ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
            url_rewrite: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallSslServer':
        """
        Get an existing FirewallSslServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] add_header_x_forwarded_proto: Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] ip: IPv4 address of the SSL server.
        :param pulumi.Input[int] mapped_port: Mapped server service port (1 - 65535, default = 80).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[int] port: Server service port (1 - 65535, default = 443).
        :param pulumi.Input[str] ssl_algorithm: Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        :param pulumi.Input[str] ssl_cert: Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        :param pulumi.Input[str] ssl_client_renegotiation: Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] ssl_max_version: Highest SSL/TLS version to negotiate.
        :param pulumi.Input[str] ssl_min_version: Lowest SSL/TLS version to negotiate.
        :param pulumi.Input[str] ssl_mode: SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
        :param pulumi.Input[str] ssl_send_empty_frags: Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] url_rewrite: Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallSslServerState.__new__(_FirewallSslServerState)

        __props__.__dict__["add_header_x_forwarded_proto"] = add_header_x_forwarded_proto
        __props__.__dict__["ip"] = ip
        __props__.__dict__["mapped_port"] = mapped_port
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["ssl_algorithm"] = ssl_algorithm
        __props__.__dict__["ssl_cert"] = ssl_cert
        __props__.__dict__["ssl_client_renegotiation"] = ssl_client_renegotiation
        __props__.__dict__["ssl_dh_bits"] = ssl_dh_bits
        __props__.__dict__["ssl_max_version"] = ssl_max_version
        __props__.__dict__["ssl_min_version"] = ssl_min_version
        __props__.__dict__["ssl_mode"] = ssl_mode
        __props__.__dict__["ssl_send_empty_frags"] = ssl_send_empty_frags
        __props__.__dict__["url_rewrite"] = url_rewrite
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallSslServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addHeaderXForwardedProto")
    def add_header_x_forwarded_proto(self) -> pulumi.Output[str]:
        """
        Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "add_header_x_forwarded_proto")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        IPv4 address of the SSL server.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="mappedPort")
    def mapped_port(self) -> pulumi.Output[int]:
        """
        Mapped server service port (1 - 65535, default = 80).
        """
        return pulumi.get(self, "mapped_port")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Server service port (1 - 65535, default = 443).
        """
        return pulumi.get(self, "port")

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
    @pulumi.getter(name="sslClientRenegotiation")
    def ssl_client_renegotiation(self) -> pulumi.Output[str]:
        """
        Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
        """
        return pulumi.get(self, "ssl_client_renegotiation")

    @property
    @pulumi.getter(name="sslDhBits")
    def ssl_dh_bits(self) -> pulumi.Output[str]:
        """
        Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        """
        return pulumi.get(self, "ssl_dh_bits")

    @property
    @pulumi.getter(name="sslMaxVersion")
    def ssl_max_version(self) -> pulumi.Output[str]:
        """
        Highest SSL/TLS version to negotiate.
        """
        return pulumi.get(self, "ssl_max_version")

    @property
    @pulumi.getter(name="sslMinVersion")
    def ssl_min_version(self) -> pulumi.Output[str]:
        """
        Lowest SSL/TLS version to negotiate.
        """
        return pulumi.get(self, "ssl_min_version")

    @property
    @pulumi.getter(name="sslMode")
    def ssl_mode(self) -> pulumi.Output[str]:
        """
        SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
        """
        return pulumi.get(self, "ssl_mode")

    @property
    @pulumi.getter(name="sslSendEmptyFrags")
    def ssl_send_empty_frags(self) -> pulumi.Output[str]:
        """
        Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_send_empty_frags")

    @property
    @pulumi.getter(name="urlRewrite")
    def url_rewrite(self) -> pulumi.Output[str]:
        """
        Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "url_rewrite")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

