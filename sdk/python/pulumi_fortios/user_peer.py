# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserPeerArgs', 'UserPeer']

@pulumi.input_type
class UserPeerArgs:
    def __init__(__self__, *,
                 ca: Optional[pulumi.Input[str]] = None,
                 cn: Optional[pulumi.Input[str]] = None,
                 cn_type: Optional[pulumi.Input[str]] = None,
                 ldap_mode: Optional[pulumi.Input[str]] = None,
                 ldap_password: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 ldap_username: Optional[pulumi.Input[str]] = None,
                 mandatory_ca_verify: Optional[pulumi.Input[str]] = None,
                 mfa_mode: Optional[pulumi.Input[str]] = None,
                 mfa_password: Optional[pulumi.Input[str]] = None,
                 mfa_server: Optional[pulumi.Input[str]] = None,
                 mfa_username: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ocsp_override_server: Optional[pulumi.Input[str]] = None,
                 passwd: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 two_factor: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserPeer resource.
        """
        if ca is not None:
            pulumi.set(__self__, "ca", ca)
        if cn is not None:
            pulumi.set(__self__, "cn", cn)
        if cn_type is not None:
            pulumi.set(__self__, "cn_type", cn_type)
        if ldap_mode is not None:
            pulumi.set(__self__, "ldap_mode", ldap_mode)
        if ldap_password is not None:
            pulumi.set(__self__, "ldap_password", ldap_password)
        if ldap_server is not None:
            pulumi.set(__self__, "ldap_server", ldap_server)
        if ldap_username is not None:
            pulumi.set(__self__, "ldap_username", ldap_username)
        if mandatory_ca_verify is not None:
            pulumi.set(__self__, "mandatory_ca_verify", mandatory_ca_verify)
        if mfa_mode is not None:
            pulumi.set(__self__, "mfa_mode", mfa_mode)
        if mfa_password is not None:
            pulumi.set(__self__, "mfa_password", mfa_password)
        if mfa_server is not None:
            pulumi.set(__self__, "mfa_server", mfa_server)
        if mfa_username is not None:
            pulumi.set(__self__, "mfa_username", mfa_username)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ocsp_override_server is not None:
            pulumi.set(__self__, "ocsp_override_server", ocsp_override_server)
        if passwd is not None:
            pulumi.set(__self__, "passwd", passwd)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if two_factor is not None:
            pulumi.set(__self__, "two_factor", two_factor)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ca(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ca")

    @ca.setter
    def ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca", value)

    @property
    @pulumi.getter
    def cn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cn")

    @cn.setter
    def cn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cn", value)

    @property
    @pulumi.getter(name="cnType")
    def cn_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cn_type")

    @cn_type.setter
    def cn_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cn_type", value)

    @property
    @pulumi.getter(name="ldapMode")
    def ldap_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldap_mode")

    @ldap_mode.setter
    def ldap_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_mode", value)

    @property
    @pulumi.getter(name="ldapPassword")
    def ldap_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldap_password")

    @ldap_password.setter
    def ldap_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_password", value)

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldap_server")

    @ldap_server.setter
    def ldap_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_server", value)

    @property
    @pulumi.getter(name="ldapUsername")
    def ldap_username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldap_username")

    @ldap_username.setter
    def ldap_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_username", value)

    @property
    @pulumi.getter(name="mandatoryCaVerify")
    def mandatory_ca_verify(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mandatory_ca_verify")

    @mandatory_ca_verify.setter
    def mandatory_ca_verify(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mandatory_ca_verify", value)

    @property
    @pulumi.getter(name="mfaMode")
    def mfa_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_mode")

    @mfa_mode.setter
    def mfa_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_mode", value)

    @property
    @pulumi.getter(name="mfaPassword")
    def mfa_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_password")

    @mfa_password.setter
    def mfa_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_password", value)

    @property
    @pulumi.getter(name="mfaServer")
    def mfa_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_server")

    @mfa_server.setter
    def mfa_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_server", value)

    @property
    @pulumi.getter(name="mfaUsername")
    def mfa_username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_username")

    @mfa_username.setter
    def mfa_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_username", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ocspOverrideServer")
    def ocsp_override_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ocsp_override_server")

    @ocsp_override_server.setter
    def ocsp_override_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ocsp_override_server", value)

    @property
    @pulumi.getter
    def passwd(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "passwd")

    @passwd.setter
    def passwd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passwd", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter(name="twoFactor")
    def two_factor(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "two_factor")

    @two_factor.setter
    def two_factor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "two_factor", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


@pulumi.input_type
class _UserPeerState:
    def __init__(__self__, *,
                 ca: Optional[pulumi.Input[str]] = None,
                 cn: Optional[pulumi.Input[str]] = None,
                 cn_type: Optional[pulumi.Input[str]] = None,
                 ldap_mode: Optional[pulumi.Input[str]] = None,
                 ldap_password: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 ldap_username: Optional[pulumi.Input[str]] = None,
                 mandatory_ca_verify: Optional[pulumi.Input[str]] = None,
                 mfa_mode: Optional[pulumi.Input[str]] = None,
                 mfa_password: Optional[pulumi.Input[str]] = None,
                 mfa_server: Optional[pulumi.Input[str]] = None,
                 mfa_username: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ocsp_override_server: Optional[pulumi.Input[str]] = None,
                 passwd: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 two_factor: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserPeer resources.
        """
        if ca is not None:
            pulumi.set(__self__, "ca", ca)
        if cn is not None:
            pulumi.set(__self__, "cn", cn)
        if cn_type is not None:
            pulumi.set(__self__, "cn_type", cn_type)
        if ldap_mode is not None:
            pulumi.set(__self__, "ldap_mode", ldap_mode)
        if ldap_password is not None:
            pulumi.set(__self__, "ldap_password", ldap_password)
        if ldap_server is not None:
            pulumi.set(__self__, "ldap_server", ldap_server)
        if ldap_username is not None:
            pulumi.set(__self__, "ldap_username", ldap_username)
        if mandatory_ca_verify is not None:
            pulumi.set(__self__, "mandatory_ca_verify", mandatory_ca_verify)
        if mfa_mode is not None:
            pulumi.set(__self__, "mfa_mode", mfa_mode)
        if mfa_password is not None:
            pulumi.set(__self__, "mfa_password", mfa_password)
        if mfa_server is not None:
            pulumi.set(__self__, "mfa_server", mfa_server)
        if mfa_username is not None:
            pulumi.set(__self__, "mfa_username", mfa_username)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ocsp_override_server is not None:
            pulumi.set(__self__, "ocsp_override_server", ocsp_override_server)
        if passwd is not None:
            pulumi.set(__self__, "passwd", passwd)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if two_factor is not None:
            pulumi.set(__self__, "two_factor", two_factor)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def ca(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ca")

    @ca.setter
    def ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca", value)

    @property
    @pulumi.getter
    def cn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cn")

    @cn.setter
    def cn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cn", value)

    @property
    @pulumi.getter(name="cnType")
    def cn_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cn_type")

    @cn_type.setter
    def cn_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cn_type", value)

    @property
    @pulumi.getter(name="ldapMode")
    def ldap_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldap_mode")

    @ldap_mode.setter
    def ldap_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_mode", value)

    @property
    @pulumi.getter(name="ldapPassword")
    def ldap_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldap_password")

    @ldap_password.setter
    def ldap_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_password", value)

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldap_server")

    @ldap_server.setter
    def ldap_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_server", value)

    @property
    @pulumi.getter(name="ldapUsername")
    def ldap_username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ldap_username")

    @ldap_username.setter
    def ldap_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_username", value)

    @property
    @pulumi.getter(name="mandatoryCaVerify")
    def mandatory_ca_verify(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mandatory_ca_verify")

    @mandatory_ca_verify.setter
    def mandatory_ca_verify(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mandatory_ca_verify", value)

    @property
    @pulumi.getter(name="mfaMode")
    def mfa_mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_mode")

    @mfa_mode.setter
    def mfa_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_mode", value)

    @property
    @pulumi.getter(name="mfaPassword")
    def mfa_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_password")

    @mfa_password.setter
    def mfa_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_password", value)

    @property
    @pulumi.getter(name="mfaServer")
    def mfa_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_server")

    @mfa_server.setter
    def mfa_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_server", value)

    @property
    @pulumi.getter(name="mfaUsername")
    def mfa_username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mfa_username")

    @mfa_username.setter
    def mfa_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mfa_username", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ocspOverrideServer")
    def ocsp_override_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ocsp_override_server")

    @ocsp_override_server.setter
    def ocsp_override_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ocsp_override_server", value)

    @property
    @pulumi.getter
    def passwd(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "passwd")

    @passwd.setter
    def passwd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passwd", value)

    @property
    @pulumi.getter
    def subject(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter(name="twoFactor")
    def two_factor(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "two_factor")

    @two_factor.setter
    def two_factor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "two_factor", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)


class UserPeer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca: Optional[pulumi.Input[str]] = None,
                 cn: Optional[pulumi.Input[str]] = None,
                 cn_type: Optional[pulumi.Input[str]] = None,
                 ldap_mode: Optional[pulumi.Input[str]] = None,
                 ldap_password: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 ldap_username: Optional[pulumi.Input[str]] = None,
                 mandatory_ca_verify: Optional[pulumi.Input[str]] = None,
                 mfa_mode: Optional[pulumi.Input[str]] = None,
                 mfa_password: Optional[pulumi.Input[str]] = None,
                 mfa_server: Optional[pulumi.Input[str]] = None,
                 mfa_username: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ocsp_override_server: Optional[pulumi.Input[str]] = None,
                 passwd: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 two_factor: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a UserPeer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserPeerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a UserPeer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param UserPeerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserPeerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca: Optional[pulumi.Input[str]] = None,
                 cn: Optional[pulumi.Input[str]] = None,
                 cn_type: Optional[pulumi.Input[str]] = None,
                 ldap_mode: Optional[pulumi.Input[str]] = None,
                 ldap_password: Optional[pulumi.Input[str]] = None,
                 ldap_server: Optional[pulumi.Input[str]] = None,
                 ldap_username: Optional[pulumi.Input[str]] = None,
                 mandatory_ca_verify: Optional[pulumi.Input[str]] = None,
                 mfa_mode: Optional[pulumi.Input[str]] = None,
                 mfa_password: Optional[pulumi.Input[str]] = None,
                 mfa_server: Optional[pulumi.Input[str]] = None,
                 mfa_username: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ocsp_override_server: Optional[pulumi.Input[str]] = None,
                 passwd: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 two_factor: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserPeerArgs.__new__(UserPeerArgs)

            __props__.__dict__["ca"] = ca
            __props__.__dict__["cn"] = cn
            __props__.__dict__["cn_type"] = cn_type
            __props__.__dict__["ldap_mode"] = ldap_mode
            __props__.__dict__["ldap_password"] = None if ldap_password is None else pulumi.Output.secret(ldap_password)
            __props__.__dict__["ldap_server"] = ldap_server
            __props__.__dict__["ldap_username"] = ldap_username
            __props__.__dict__["mandatory_ca_verify"] = mandatory_ca_verify
            __props__.__dict__["mfa_mode"] = mfa_mode
            __props__.__dict__["mfa_password"] = mfa_password
            __props__.__dict__["mfa_server"] = mfa_server
            __props__.__dict__["mfa_username"] = mfa_username
            __props__.__dict__["name"] = name
            __props__.__dict__["ocsp_override_server"] = ocsp_override_server
            __props__.__dict__["passwd"] = None if passwd is None else pulumi.Output.secret(passwd)
            __props__.__dict__["subject"] = subject
            __props__.__dict__["two_factor"] = two_factor
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["ldapPassword", "passwd"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(UserPeer, __self__).__init__(
            'fortios:index/userPeer:UserPeer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ca: Optional[pulumi.Input[str]] = None,
            cn: Optional[pulumi.Input[str]] = None,
            cn_type: Optional[pulumi.Input[str]] = None,
            ldap_mode: Optional[pulumi.Input[str]] = None,
            ldap_password: Optional[pulumi.Input[str]] = None,
            ldap_server: Optional[pulumi.Input[str]] = None,
            ldap_username: Optional[pulumi.Input[str]] = None,
            mandatory_ca_verify: Optional[pulumi.Input[str]] = None,
            mfa_mode: Optional[pulumi.Input[str]] = None,
            mfa_password: Optional[pulumi.Input[str]] = None,
            mfa_server: Optional[pulumi.Input[str]] = None,
            mfa_username: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ocsp_override_server: Optional[pulumi.Input[str]] = None,
            passwd: Optional[pulumi.Input[str]] = None,
            subject: Optional[pulumi.Input[str]] = None,
            two_factor: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'UserPeer':
        """
        Get an existing UserPeer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserPeerState.__new__(_UserPeerState)

        __props__.__dict__["ca"] = ca
        __props__.__dict__["cn"] = cn
        __props__.__dict__["cn_type"] = cn_type
        __props__.__dict__["ldap_mode"] = ldap_mode
        __props__.__dict__["ldap_password"] = ldap_password
        __props__.__dict__["ldap_server"] = ldap_server
        __props__.__dict__["ldap_username"] = ldap_username
        __props__.__dict__["mandatory_ca_verify"] = mandatory_ca_verify
        __props__.__dict__["mfa_mode"] = mfa_mode
        __props__.__dict__["mfa_password"] = mfa_password
        __props__.__dict__["mfa_server"] = mfa_server
        __props__.__dict__["mfa_username"] = mfa_username
        __props__.__dict__["name"] = name
        __props__.__dict__["ocsp_override_server"] = ocsp_override_server
        __props__.__dict__["passwd"] = passwd
        __props__.__dict__["subject"] = subject
        __props__.__dict__["two_factor"] = two_factor
        __props__.__dict__["vdomparam"] = vdomparam
        return UserPeer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ca(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ca")

    @property
    @pulumi.getter
    def cn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cn")

    @property
    @pulumi.getter(name="cnType")
    def cn_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cn_type")

    @property
    @pulumi.getter(name="ldapMode")
    def ldap_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ldap_mode")

    @property
    @pulumi.getter(name="ldapPassword")
    def ldap_password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ldap_password")

    @property
    @pulumi.getter(name="ldapServer")
    def ldap_server(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ldap_server")

    @property
    @pulumi.getter(name="ldapUsername")
    def ldap_username(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ldap_username")

    @property
    @pulumi.getter(name="mandatoryCaVerify")
    def mandatory_ca_verify(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mandatory_ca_verify")

    @property
    @pulumi.getter(name="mfaMode")
    def mfa_mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mfa_mode")

    @property
    @pulumi.getter(name="mfaPassword")
    def mfa_password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "mfa_password")

    @property
    @pulumi.getter(name="mfaServer")
    def mfa_server(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mfa_server")

    @property
    @pulumi.getter(name="mfaUsername")
    def mfa_username(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mfa_username")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ocspOverrideServer")
    def ocsp_override_server(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ocsp_override_server")

    @property
    @pulumi.getter
    def passwd(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "passwd")

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter(name="twoFactor")
    def two_factor(self) -> pulumi.Output[str]:
        return pulumi.get(self, "two_factor")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vdomparam")

