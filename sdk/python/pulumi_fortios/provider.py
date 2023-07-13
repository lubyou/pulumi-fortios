# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 cabundlecontent: Optional[pulumi.Input[str]] = None,
                 cabundlefile: Optional[pulumi.Input[str]] = None,
                 cacert: Optional[pulumi.Input[str]] = None,
                 clientcert: Optional[pulumi.Input[str]] = None,
                 clientkey: Optional[pulumi.Input[str]] = None,
                 fmg_cabundlefile: Optional[pulumi.Input[str]] = None,
                 fmg_hostname: Optional[pulumi.Input[str]] = None,
                 fmg_insecure: Optional[pulumi.Input[bool]] = None,
                 fmg_passwd: Optional[pulumi.Input[str]] = None,
                 fmg_username: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 http_proxy: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 peerauth: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] cabundlecontent: CA Bundle file content
        :param pulumi.Input[str] cabundlefile: CA Bundle file
        :param pulumi.Input[str] cacert: CA certtificate(Optional)
        :param pulumi.Input[str] clientcert: User certificate
        :param pulumi.Input[str] clientkey: User private key
        :param pulumi.Input[str] fmg_cabundlefile: CA Bundle file
        :param pulumi.Input[str] fmg_hostname: Hostname/IP address of the FortiManager to connect to
        :param pulumi.Input[str] hostname: The hostname/IP address of the FortiOS to be connected
        :param pulumi.Input[str] http_proxy: HTTP proxy address
        :param pulumi.Input[str] password: The password of the user.
        :param pulumi.Input[str] peerauth: Enable/disable peer authentication, can be 'enable' or 'disable'
        :param pulumi.Input[str] username: The username of the user.
        """
        if cabundlecontent is not None:
            pulumi.set(__self__, "cabundlecontent", cabundlecontent)
        if cabundlefile is None:
            cabundlefile = _utilities.get_env('FORTIOS_CA_CABUNDLE')
        if cabundlefile is not None:
            pulumi.set(__self__, "cabundlefile", cabundlefile)
        if cacert is not None:
            pulumi.set(__self__, "cacert", cacert)
        if clientcert is not None:
            pulumi.set(__self__, "clientcert", clientcert)
        if clientkey is not None:
            pulumi.set(__self__, "clientkey", clientkey)
        if fmg_cabundlefile is None:
            fmg_cabundlefile = _utilities.get_env('FORTIOS_FMG_CABUNDLE')
        if fmg_cabundlefile is not None:
            pulumi.set(__self__, "fmg_cabundlefile", fmg_cabundlefile)
        if fmg_hostname is None:
            fmg_hostname = _utilities.get_env('FORTIOS_FMG_HOSTNAME')
        if fmg_hostname is not None:
            pulumi.set(__self__, "fmg_hostname", fmg_hostname)
        if fmg_insecure is None:
            fmg_insecure = _utilities.get_env_bool('FORTIOS_FMG_INSECURE')
        if fmg_insecure is not None:
            pulumi.set(__self__, "fmg_insecure", fmg_insecure)
        if fmg_passwd is None:
            fmg_passwd = _utilities.get_env('FORTIOS_FMG_PASSWORD')
        if fmg_passwd is not None:
            pulumi.set(__self__, "fmg_passwd", fmg_passwd)
        if fmg_username is None:
            fmg_username = _utilities.get_env('FORTIOS_FMG_USERNAME')
        if fmg_username is not None:
            pulumi.set(__self__, "fmg_username", fmg_username)
        if hostname is None:
            hostname = _utilities.get_env('FORTIOS_ACCESS_HOSTNAME')
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if http_proxy is not None:
            pulumi.set(__self__, "http_proxy", http_proxy)
        if insecure is None:
            insecure = _utilities.get_env_bool('FORTIOS_INSECURE')
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if peerauth is not None:
            pulumi.set(__self__, "peerauth", peerauth)
        if token is None:
            token = _utilities.get_env('FORTIOS_ACCESS_TOKEN')
        if token is not None:
            pulumi.set(__self__, "token", token)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)

    @property
    @pulumi.getter
    def cabundlecontent(self) -> Optional[pulumi.Input[str]]:
        """
        CA Bundle file content
        """
        return pulumi.get(self, "cabundlecontent")

    @cabundlecontent.setter
    def cabundlecontent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cabundlecontent", value)

    @property
    @pulumi.getter
    def cabundlefile(self) -> Optional[pulumi.Input[str]]:
        """
        CA Bundle file
        """
        return pulumi.get(self, "cabundlefile")

    @cabundlefile.setter
    def cabundlefile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cabundlefile", value)

    @property
    @pulumi.getter
    def cacert(self) -> Optional[pulumi.Input[str]]:
        """
        CA certtificate(Optional)
        """
        return pulumi.get(self, "cacert")

    @cacert.setter
    def cacert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cacert", value)

    @property
    @pulumi.getter
    def clientcert(self) -> Optional[pulumi.Input[str]]:
        """
        User certificate
        """
        return pulumi.get(self, "clientcert")

    @clientcert.setter
    def clientcert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clientcert", value)

    @property
    @pulumi.getter
    def clientkey(self) -> Optional[pulumi.Input[str]]:
        """
        User private key
        """
        return pulumi.get(self, "clientkey")

    @clientkey.setter
    def clientkey(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "clientkey", value)

    @property
    @pulumi.getter(name="fmgCabundlefile")
    def fmg_cabundlefile(self) -> Optional[pulumi.Input[str]]:
        """
        CA Bundle file
        """
        return pulumi.get(self, "fmg_cabundlefile")

    @fmg_cabundlefile.setter
    def fmg_cabundlefile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fmg_cabundlefile", value)

    @property
    @pulumi.getter(name="fmgHostname")
    def fmg_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Hostname/IP address of the FortiManager to connect to
        """
        return pulumi.get(self, "fmg_hostname")

    @fmg_hostname.setter
    def fmg_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fmg_hostname", value)

    @property
    @pulumi.getter(name="fmgInsecure")
    def fmg_insecure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "fmg_insecure")

    @fmg_insecure.setter
    def fmg_insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "fmg_insecure", value)

    @property
    @pulumi.getter(name="fmgPasswd")
    def fmg_passwd(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fmg_passwd")

    @fmg_passwd.setter
    def fmg_passwd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fmg_passwd", value)

    @property
    @pulumi.getter(name="fmgUsername")
    def fmg_username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fmg_username")

    @fmg_username.setter
    def fmg_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fmg_username", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        The hostname/IP address of the FortiOS to be connected
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="httpProxy")
    def http_proxy(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP proxy address
        """
        return pulumi.get(self, "http_proxy")

    @http_proxy.setter
    def http_proxy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_proxy", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def peerauth(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable peer authentication, can be 'enable' or 'disable'
        """
        return pulumi.get(self, "peerauth")

    @peerauth.setter
    def peerauth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peerauth", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username of the user.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cabundlecontent: Optional[pulumi.Input[str]] = None,
                 cabundlefile: Optional[pulumi.Input[str]] = None,
                 cacert: Optional[pulumi.Input[str]] = None,
                 clientcert: Optional[pulumi.Input[str]] = None,
                 clientkey: Optional[pulumi.Input[str]] = None,
                 fmg_cabundlefile: Optional[pulumi.Input[str]] = None,
                 fmg_hostname: Optional[pulumi.Input[str]] = None,
                 fmg_insecure: Optional[pulumi.Input[bool]] = None,
                 fmg_passwd: Optional[pulumi.Input[str]] = None,
                 fmg_username: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 http_proxy: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 peerauth: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the fortios package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cabundlecontent: CA Bundle file content
        :param pulumi.Input[str] cabundlefile: CA Bundle file
        :param pulumi.Input[str] cacert: CA certtificate(Optional)
        :param pulumi.Input[str] clientcert: User certificate
        :param pulumi.Input[str] clientkey: User private key
        :param pulumi.Input[str] fmg_cabundlefile: CA Bundle file
        :param pulumi.Input[str] fmg_hostname: Hostname/IP address of the FortiManager to connect to
        :param pulumi.Input[str] hostname: The hostname/IP address of the FortiOS to be connected
        :param pulumi.Input[str] http_proxy: HTTP proxy address
        :param pulumi.Input[str] password: The password of the user.
        :param pulumi.Input[str] peerauth: Enable/disable peer authentication, can be 'enable' or 'disable'
        :param pulumi.Input[str] username: The username of the user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the fortios package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cabundlecontent: Optional[pulumi.Input[str]] = None,
                 cabundlefile: Optional[pulumi.Input[str]] = None,
                 cacert: Optional[pulumi.Input[str]] = None,
                 clientcert: Optional[pulumi.Input[str]] = None,
                 clientkey: Optional[pulumi.Input[str]] = None,
                 fmg_cabundlefile: Optional[pulumi.Input[str]] = None,
                 fmg_hostname: Optional[pulumi.Input[str]] = None,
                 fmg_insecure: Optional[pulumi.Input[bool]] = None,
                 fmg_passwd: Optional[pulumi.Input[str]] = None,
                 fmg_username: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 http_proxy: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 peerauth: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["cabundlecontent"] = cabundlecontent
            if cabundlefile is None:
                cabundlefile = _utilities.get_env('FORTIOS_CA_CABUNDLE')
            __props__.__dict__["cabundlefile"] = cabundlefile
            __props__.__dict__["cacert"] = cacert
            __props__.__dict__["clientcert"] = clientcert
            __props__.__dict__["clientkey"] = clientkey
            if fmg_cabundlefile is None:
                fmg_cabundlefile = _utilities.get_env('FORTIOS_FMG_CABUNDLE')
            __props__.__dict__["fmg_cabundlefile"] = fmg_cabundlefile
            if fmg_hostname is None:
                fmg_hostname = _utilities.get_env('FORTIOS_FMG_HOSTNAME')
            __props__.__dict__["fmg_hostname"] = fmg_hostname
            if fmg_insecure is None:
                fmg_insecure = _utilities.get_env_bool('FORTIOS_FMG_INSECURE')
            __props__.__dict__["fmg_insecure"] = pulumi.Output.from_input(fmg_insecure).apply(pulumi.runtime.to_json) if fmg_insecure is not None else None
            if fmg_passwd is None:
                fmg_passwd = _utilities.get_env('FORTIOS_FMG_PASSWORD')
            __props__.__dict__["fmg_passwd"] = fmg_passwd
            if fmg_username is None:
                fmg_username = _utilities.get_env('FORTIOS_FMG_USERNAME')
            __props__.__dict__["fmg_username"] = fmg_username
            if hostname is None:
                hostname = _utilities.get_env('FORTIOS_ACCESS_HOSTNAME')
            __props__.__dict__["hostname"] = hostname
            __props__.__dict__["http_proxy"] = http_proxy
            if insecure is None:
                insecure = _utilities.get_env_bool('FORTIOS_INSECURE')
            __props__.__dict__["insecure"] = pulumi.Output.from_input(insecure).apply(pulumi.runtime.to_json) if insecure is not None else None
            __props__.__dict__["password"] = password
            __props__.__dict__["peerauth"] = peerauth
            if token is None:
                token = _utilities.get_env('FORTIOS_ACCESS_TOKEN')
            __props__.__dict__["token"] = token
            __props__.__dict__["username"] = username
            __props__.__dict__["vdom"] = vdom
        super(Provider, __self__).__init__(
            'fortios',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter
    def cabundlecontent(self) -> pulumi.Output[Optional[str]]:
        """
        CA Bundle file content
        """
        return pulumi.get(self, "cabundlecontent")

    @property
    @pulumi.getter
    def cabundlefile(self) -> pulumi.Output[Optional[str]]:
        """
        CA Bundle file
        """
        return pulumi.get(self, "cabundlefile")

    @property
    @pulumi.getter
    def cacert(self) -> pulumi.Output[Optional[str]]:
        """
        CA certtificate(Optional)
        """
        return pulumi.get(self, "cacert")

    @property
    @pulumi.getter
    def clientcert(self) -> pulumi.Output[Optional[str]]:
        """
        User certificate
        """
        return pulumi.get(self, "clientcert")

    @property
    @pulumi.getter
    def clientkey(self) -> pulumi.Output[Optional[str]]:
        """
        User private key
        """
        return pulumi.get(self, "clientkey")

    @property
    @pulumi.getter(name="fmgCabundlefile")
    def fmg_cabundlefile(self) -> pulumi.Output[Optional[str]]:
        """
        CA Bundle file
        """
        return pulumi.get(self, "fmg_cabundlefile")

    @property
    @pulumi.getter(name="fmgHostname")
    def fmg_hostname(self) -> pulumi.Output[Optional[str]]:
        """
        Hostname/IP address of the FortiManager to connect to
        """
        return pulumi.get(self, "fmg_hostname")

    @property
    @pulumi.getter(name="fmgPasswd")
    def fmg_passwd(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "fmg_passwd")

    @property
    @pulumi.getter(name="fmgUsername")
    def fmg_username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "fmg_username")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[Optional[str]]:
        """
        The hostname/IP address of the FortiOS to be connected
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="httpProxy")
    def http_proxy(self) -> pulumi.Output[Optional[str]]:
        """
        HTTP proxy address
        """
        return pulumi.get(self, "http_proxy")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password of the user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def peerauth(self) -> pulumi.Output[Optional[str]]:
        """
        Enable/disable peer authentication, can be 'enable' or 'disable'
        """
        return pulumi.get(self, "peerauth")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        The username of the user.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vdom(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdom")

