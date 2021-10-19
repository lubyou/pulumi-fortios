# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetUserSamlResult',
    'AwaitableGetUserSamlResult',
    'get_user_saml',
]

@pulumi.output_type
class GetUserSamlResult:
    """
    A collection of values returned by GetUserSaml.
    """
    def __init__(__self__, cert=None, entity_id=None, group_name=None, id=None, idp_cert=None, idp_entity_id=None, idp_single_logout_url=None, idp_single_sign_on_url=None, name=None, single_logout_url=None, single_sign_on_url=None, user_name=None, vdomparam=None):
        if cert and not isinstance(cert, str):
            raise TypeError("Expected argument 'cert' to be a str")
        pulumi.set(__self__, "cert", cert)
        if entity_id and not isinstance(entity_id, str):
            raise TypeError("Expected argument 'entity_id' to be a str")
        pulumi.set(__self__, "entity_id", entity_id)
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idp_cert and not isinstance(idp_cert, str):
            raise TypeError("Expected argument 'idp_cert' to be a str")
        pulumi.set(__self__, "idp_cert", idp_cert)
        if idp_entity_id and not isinstance(idp_entity_id, str):
            raise TypeError("Expected argument 'idp_entity_id' to be a str")
        pulumi.set(__self__, "idp_entity_id", idp_entity_id)
        if idp_single_logout_url and not isinstance(idp_single_logout_url, str):
            raise TypeError("Expected argument 'idp_single_logout_url' to be a str")
        pulumi.set(__self__, "idp_single_logout_url", idp_single_logout_url)
        if idp_single_sign_on_url and not isinstance(idp_single_sign_on_url, str):
            raise TypeError("Expected argument 'idp_single_sign_on_url' to be a str")
        pulumi.set(__self__, "idp_single_sign_on_url", idp_single_sign_on_url)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if single_logout_url and not isinstance(single_logout_url, str):
            raise TypeError("Expected argument 'single_logout_url' to be a str")
        pulumi.set(__self__, "single_logout_url", single_logout_url)
        if single_sign_on_url and not isinstance(single_sign_on_url, str):
            raise TypeError("Expected argument 'single_sign_on_url' to be a str")
        pulumi.set(__self__, "single_sign_on_url", single_sign_on_url)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def cert(self) -> str:
        """
        Certificate to sign SAML messages.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> str:
        """
        SP entity ID.
        """
        return pulumi.get(self, "entity_id")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        """
        Group name in assertion statement.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idpCert")
    def idp_cert(self) -> str:
        """
        IDP Certificate name.
        """
        return pulumi.get(self, "idp_cert")

    @property
    @pulumi.getter(name="idpEntityId")
    def idp_entity_id(self) -> str:
        """
        IDP entity ID.
        """
        return pulumi.get(self, "idp_entity_id")

    @property
    @pulumi.getter(name="idpSingleLogoutUrl")
    def idp_single_logout_url(self) -> str:
        """
        IDP single logout url.
        """
        return pulumi.get(self, "idp_single_logout_url")

    @property
    @pulumi.getter(name="idpSingleSignOnUrl")
    def idp_single_sign_on_url(self) -> str:
        """
        IDP single sign-on URL.
        """
        return pulumi.get(self, "idp_single_sign_on_url")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        SAML server entry name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="singleLogoutUrl")
    def single_logout_url(self) -> str:
        """
        SP single logout URL.
        """
        return pulumi.get(self, "single_logout_url")

    @property
    @pulumi.getter(name="singleSignOnUrl")
    def single_sign_on_url(self) -> str:
        """
        SP single sign-on URL.
        """
        return pulumi.get(self, "single_sign_on_url")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        User name in assertion statement.
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetUserSamlResult(GetUserSamlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserSamlResult(
            cert=self.cert,
            entity_id=self.entity_id,
            group_name=self.group_name,
            id=self.id,
            idp_cert=self.idp_cert,
            idp_entity_id=self.idp_entity_id,
            idp_single_logout_url=self.idp_single_logout_url,
            idp_single_sign_on_url=self.idp_single_sign_on_url,
            name=self.name,
            single_logout_url=self.single_logout_url,
            single_sign_on_url=self.single_sign_on_url,
            user_name=self.user_name,
            vdomparam=self.vdomparam)


def get_user_saml(name: Optional[str] = None,
                  vdomparam: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserSamlResult:
    """
    Use this data source to get information on an fortios user saml


    :param str name: Specify the name of the desired user saml.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdomparam'] = vdomparam
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('fortios:index/getUserSaml:GetUserSaml', __args__, opts=opts, typ=GetUserSamlResult).value

    return AwaitableGetUserSamlResult(
        cert=__ret__.cert,
        entity_id=__ret__.entity_id,
        group_name=__ret__.group_name,
        id=__ret__.id,
        idp_cert=__ret__.idp_cert,
        idp_entity_id=__ret__.idp_entity_id,
        idp_single_logout_url=__ret__.idp_single_logout_url,
        idp_single_sign_on_url=__ret__.idp_single_sign_on_url,
        name=__ret__.name,
        single_logout_url=__ret__.single_logout_url,
        single_sign_on_url=__ret__.single_sign_on_url,
        user_name=__ret__.user_name,
        vdomparam=__ret__.vdomparam)
