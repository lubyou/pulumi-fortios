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

__all__ = [
    'GetFirewallDosPolicyResult',
    'AwaitableGetFirewallDosPolicyResult',
    'get_firewall_dos_policy',
    'get_firewall_dos_policy_output',
]

@pulumi.output_type
class GetFirewallDosPolicyResult:
    """
    A collection of values returned by GetFirewallDosPolicy.
    """
    def __init__(__self__, anomalies=None, comments=None, dstaddrs=None, id=None, interface=None, name=None, policyid=None, services=None, srcaddrs=None, status=None, vdomparam=None):
        if anomalies and not isinstance(anomalies, list):
            raise TypeError("Expected argument 'anomalies' to be a list")
        pulumi.set(__self__, "anomalies", anomalies)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if dstaddrs and not isinstance(dstaddrs, list):
            raise TypeError("Expected argument 'dstaddrs' to be a list")
        pulumi.set(__self__, "dstaddrs", dstaddrs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policyid and not isinstance(policyid, int):
            raise TypeError("Expected argument 'policyid' to be a int")
        pulumi.set(__self__, "policyid", policyid)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
        if srcaddrs and not isinstance(srcaddrs, list):
            raise TypeError("Expected argument 'srcaddrs' to be a list")
        pulumi.set(__self__, "srcaddrs", srcaddrs)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def anomalies(self) -> Sequence['outputs.GetFirewallDosPolicyAnomalyResult']:
        return pulumi.get(self, "anomalies")

    @property
    @pulumi.getter
    def comments(self) -> str:
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def dstaddrs(self) -> Sequence['outputs.GetFirewallDosPolicyDstaddrResult']:
        return pulumi.get(self, "dstaddrs")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policyid(self) -> int:
        return pulumi.get(self, "policyid")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetFirewallDosPolicyServiceResult']:
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def srcaddrs(self) -> Sequence['outputs.GetFirewallDosPolicySrcaddrResult']:
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallDosPolicyResult(GetFirewallDosPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallDosPolicyResult(
            anomalies=self.anomalies,
            comments=self.comments,
            dstaddrs=self.dstaddrs,
            id=self.id,
            interface=self.interface,
            name=self.name,
            policyid=self.policyid,
            services=self.services,
            srcaddrs=self.srcaddrs,
            status=self.status,
            vdomparam=self.vdomparam)


def get_firewall_dos_policy(policyid: Optional[int] = None,
                            vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallDosPolicyResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['policyid'] = policyid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallDosPolicy:GetFirewallDosPolicy', __args__, opts=opts, typ=GetFirewallDosPolicyResult).value

    return AwaitableGetFirewallDosPolicyResult(
        anomalies=pulumi.get(__ret__, 'anomalies'),
        comments=pulumi.get(__ret__, 'comments'),
        dstaddrs=pulumi.get(__ret__, 'dstaddrs'),
        id=pulumi.get(__ret__, 'id'),
        interface=pulumi.get(__ret__, 'interface'),
        name=pulumi.get(__ret__, 'name'),
        policyid=pulumi.get(__ret__, 'policyid'),
        services=pulumi.get(__ret__, 'services'),
        srcaddrs=pulumi.get(__ret__, 'srcaddrs'),
        status=pulumi.get(__ret__, 'status'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_firewall_dos_policy)
def get_firewall_dos_policy_output(policyid: Optional[pulumi.Input[int]] = None,
                                   vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallDosPolicyResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
