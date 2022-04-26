# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallSshHostKeyArgs', 'FirewallSshHostKey']

@pulumi.input_type
class FirewallSshHostKeyArgs:
    def __init__(__self__, *,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nid: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallSshHostKey resource.
        :param pulumi.Input[str] hostname: Hostname of the SSH server.
        :param pulumi.Input[str] ip: IP address of the SSH server.
        :param pulumi.Input[str] name: SSH public key name.
        :param pulumi.Input[str] nid: Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        :param pulumi.Input[int] port: Port of the SSH server.
        :param pulumi.Input[str] public_key: SSH public key.
        :param pulumi.Input[str] status: Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        :param pulumi.Input[str] type: Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        :param pulumi.Input[str] usage: Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nid is not None:
            pulumi.set(__self__, "nid", nid)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if usage is not None:
            pulumi.set(__self__, "usage", usage)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the SSH server.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the SSH server.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        SSH public key name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nid(self) -> Optional[pulumi.Input[str]]:
        """
        Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        """
        return pulumi.get(self, "nid")

    @nid.setter
    def nid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nid", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port of the SSH server.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        SSH public key.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def usage(self) -> Optional[pulumi.Input[str]]:
        """
        Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        """
        return pulumi.get(self, "usage")

    @usage.setter
    def usage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage", value)

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
class _FirewallSshHostKeyState:
    def __init__(__self__, *,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nid: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallSshHostKey resources.
        :param pulumi.Input[str] hostname: Hostname of the SSH server.
        :param pulumi.Input[str] ip: IP address of the SSH server.
        :param pulumi.Input[str] name: SSH public key name.
        :param pulumi.Input[str] nid: Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        :param pulumi.Input[int] port: Port of the SSH server.
        :param pulumi.Input[str] public_key: SSH public key.
        :param pulumi.Input[str] status: Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        :param pulumi.Input[str] type: Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        :param pulumi.Input[str] usage: Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nid is not None:
            pulumi.set(__self__, "nid", nid)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if usage is not None:
            pulumi.set(__self__, "usage", usage)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname of the SSH server.
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the SSH server.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        SSH public key name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def nid(self) -> Optional[pulumi.Input[str]]:
        """
        Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        """
        return pulumi.get(self, "nid")

    @nid.setter
    def nid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nid", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port of the SSH server.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        SSH public key.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def usage(self) -> Optional[pulumi.Input[str]]:
        """
        Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        """
        return pulumi.get(self, "usage")

    @usage.setter
    def usage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage", value)

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


class FirewallSshHostKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nid: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        SSH proxy host public keys.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallSshHostKey("trname",
            hostname="testmachine",
            ip="1.1.1.1",
            port=22,
            status="trusted",
            type="RSA")
        ```

        ## Import

        FirewallSsh HostKey can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallSshHostKey:FirewallSshHostKey labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallSshHostKey:FirewallSshHostKey labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hostname: Hostname of the SSH server.
        :param pulumi.Input[str] ip: IP address of the SSH server.
        :param pulumi.Input[str] name: SSH public key name.
        :param pulumi.Input[str] nid: Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        :param pulumi.Input[int] port: Port of the SSH server.
        :param pulumi.Input[str] public_key: SSH public key.
        :param pulumi.Input[str] status: Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        :param pulumi.Input[str] type: Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        :param pulumi.Input[str] usage: Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallSshHostKeyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        SSH proxy host public keys.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_fortios as fortios

        trname = fortios.FirewallSshHostKey("trname",
            hostname="testmachine",
            ip="1.1.1.1",
            port=22,
            status="trusted",
            type="RSA")
        ```

        ## Import

        FirewallSsh HostKey can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallSshHostKey:FirewallSshHostKey labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallSshHostKey:FirewallSshHostKey labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallSshHostKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallSshHostKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nid: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[str]] = None,
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
            __props__ = FirewallSshHostKeyArgs.__new__(FirewallSshHostKeyArgs)

            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["ip"] = ip
            __props__.__dict__["name"] = name
            __props__.__dict__["nid"] = nid
            __props__.__dict__["port"] = port
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["status"] = status
            __props__.__dict__["type"] = type
            __props__.__dict__["usage"] = usage
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallSshHostKey, __self__).__init__(
            'fortios:index/firewallSshHostKey:FirewallSshHostKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nid: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            usage: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallSshHostKey':
        """
        Get an existing FirewallSshHostKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hostname: Hostname of the SSH server.
        :param pulumi.Input[str] ip: IP address of the SSH server.
        :param pulumi.Input[str] name: SSH public key name.
        :param pulumi.Input[str] nid: Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        :param pulumi.Input[int] port: Port of the SSH server.
        :param pulumi.Input[str] public_key: SSH public key.
        :param pulumi.Input[str] status: Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        :param pulumi.Input[str] type: Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        :param pulumi.Input[str] usage: Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallSshHostKeyState.__new__(_FirewallSshHostKeyState)

        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["ip"] = ip
        __props__.__dict__["name"] = name
        __props__.__dict__["nid"] = nid
        __props__.__dict__["port"] = port
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        __props__.__dict__["usage"] = usage
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallSshHostKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        Hostname of the SSH server.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        IP address of the SSH server.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        SSH public key name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nid(self) -> pulumi.Output[str]:
        """
        Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        """
        return pulumi.get(self, "nid")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Port of the SSH server.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[Optional[str]]:
        """
        SSH public key.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def usage(self) -> pulumi.Output[str]:
        """
        Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        """
        return pulumi.get(self, "usage")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

