/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha5

import (
	"github.com/metal3-io/cluster-api-provider-metal3/api/v1beta1"
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// Constant variables
const (
	APIEndpointPort = "6443"
)

func (src *Metal3Cluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3Cluster)
	if err := Convert_v1alpha5_Metal3Cluster_To_v1beta1_Metal3Cluster(src, dst, nil); err != nil {
		return err
	}
	// Manually restore data.
	restored := &v1beta1.Metal3Cluster{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}
	dst.Status.Conditions = restored.Status.Conditions
	return nil
}

func (dst *Metal3Cluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3Cluster)
	if err := Convert_v1beta1_Metal3Cluster_To_v1alpha5_Metal3Cluster(src, dst, nil); err != nil {
		return err
	}
	// Preserve Hub data on down-conversion except for metadata
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}
	return nil
}

// Status.Conditions was introduced in v1beta1, thus requiring a custom conversion function; the values is going to be preserved in an annotation thus allowing roundtrip without losing information.
func Convert_v1beta1_Metal3ClusterStatus_To_v1alpha5_Metal3ClusterStatus(in *v1beta1.Metal3ClusterStatus, out *Metal3ClusterStatus, s apiconversion.Scope) error {
	return autoConvert_v1beta1_Metal3ClusterStatus_To_v1alpha5_Metal3ClusterStatus(in, out, s)
}

func (src *Metal3ClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3ClusterList)
	return Convert_v1alpha5_Metal3ClusterList_To_v1beta1_Metal3ClusterList(src, dst, nil)
}

func (dst *Metal3ClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3ClusterList)
	return Convert_v1beta1_Metal3ClusterList_To_v1alpha5_Metal3ClusterList(src, dst, nil)
}

func (src *Metal3Machine) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3Machine)
	if err := Convert_v1alpha5_Metal3Machine_To_v1beta1_Metal3Machine(src, dst, nil); err != nil {
		return err
	}
	// Manually restore data.
	restored := &v1beta1.Metal3Machine{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}
	dst.Status.Conditions = restored.Status.Conditions
	return nil
}

func (dst *Metal3Machine) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3Machine)
	if err := Convert_v1beta1_Metal3Machine_To_v1alpha5_Metal3Machine(src, dst, nil); err != nil {
		return err
	}
	// Preserve Hub data on down-conversion except for metadata
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}
	return nil
}

// Status.Conditions was introduced in v1beta1, thus requiring a custom conversion function; the values is going to be preserved in an annotation thus allowing roundtrip without losing information.
func Convert_v1beta1_Metal3MachineStatus_To_v1alpha5_Metal3MachineStatus(in *v1beta1.Metal3MachineStatus, out *Metal3MachineStatus, s apiconversion.Scope) error {
	return autoConvert_v1beta1_Metal3MachineStatus_To_v1alpha5_Metal3MachineStatus(in, out, s)
}

func (src *Metal3MachineList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3MachineList)
	return Convert_v1alpha5_Metal3MachineList_To_v1beta1_Metal3MachineList(src, dst, nil)
}

func (dst *Metal3MachineList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3MachineList)
	return Convert_v1beta1_Metal3MachineList_To_v1alpha5_Metal3MachineList(src, dst, nil)
}

func (src *Metal3MachineTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3MachineTemplate)
	if err := Convert_v1alpha5_Metal3MachineTemplate_To_v1beta1_Metal3MachineTemplate(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (dst *Metal3MachineTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3MachineTemplate)
	if err := Convert_v1beta1_Metal3MachineTemplate_To_v1alpha5_Metal3MachineTemplate(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *Metal3MachineTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3MachineTemplateList)
	return Convert_v1alpha5_Metal3MachineTemplateList_To_v1beta1_Metal3MachineTemplateList(src, dst, nil)
}

func (dst *Metal3MachineTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3MachineTemplateList)
	return Convert_v1beta1_Metal3MachineTemplateList_To_v1alpha5_Metal3MachineTemplateList(src, dst, nil)
}

func (src *Metal3Data) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3Data)
	if err := Convert_v1alpha5_Metal3Data_To_v1beta1_Metal3Data(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (dst *Metal3Data) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3Data)
	if err := Convert_v1beta1_Metal3Data_To_v1alpha5_Metal3Data(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *Metal3DataList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3DataList)
	return Convert_v1alpha5_Metal3DataList_To_v1beta1_Metal3DataList(src, dst, nil)
}

func (dst *Metal3DataList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3DataList)
	return Convert_v1beta1_Metal3DataList_To_v1alpha5_Metal3DataList(src, dst, nil)
}

func (src *Metal3DataTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3DataTemplate)
	if err := Convert_v1alpha5_Metal3DataTemplate_To_v1beta1_Metal3DataTemplate(src, dst, nil); err != nil {
		return err
	}

	restored := &v1beta1.Metal3DataTemplate{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	if dst.Spec.MetaData != nil && restored.Spec.MetaData != nil {
		for k := range dst.Spec.MetaData.IPAddressesFromPool {
			dst.Spec.MetaData.IPAddressesFromPool[k].APIGroup = restored.Spec.MetaData.IPAddressesFromPool[k].APIGroup
			dst.Spec.MetaData.IPAddressesFromPool[k].Kind = restored.Spec.MetaData.IPAddressesFromPool[k].Kind
		}
		for k := range dst.Spec.MetaData.GatewaysFromPool {
			dst.Spec.MetaData.GatewaysFromPool[k].APIGroup = restored.Spec.MetaData.GatewaysFromPool[k].APIGroup
			dst.Spec.MetaData.GatewaysFromPool[k].Kind = restored.Spec.MetaData.GatewaysFromPool[k].Kind
		}
		for k := range dst.Spec.MetaData.PrefixesFromPool {
			dst.Spec.MetaData.PrefixesFromPool[k].APIGroup = restored.Spec.MetaData.PrefixesFromPool[k].APIGroup
			dst.Spec.MetaData.PrefixesFromPool[k].Kind = restored.Spec.MetaData.PrefixesFromPool[k].Kind
		}
		for k := range dst.Spec.MetaData.DNSServersFromPool {
			dst.Spec.MetaData.DNSServersFromPool[k].APIGroup = restored.Spec.MetaData.DNSServersFromPool[k].APIGroup
			dst.Spec.MetaData.DNSServersFromPool[k].Kind = restored.Spec.MetaData.DNSServersFromPool[k].Kind
		}
	}
	if dst.Spec.NetworkData != nil && restored.Spec.NetworkData != nil {
		for k := range dst.Spec.NetworkData.Links.Ethernets {
			dst.Spec.NetworkData.Links.Ethernets[k].MACAddress = restored.Spec.NetworkData.Links.Ethernets[k].MACAddress
		}
		for k := range dst.Spec.NetworkData.Links.Vlans {
			dst.Spec.NetworkData.Links.Vlans[k].MACAddress = restored.Spec.NetworkData.Links.Vlans[k].MACAddress
		}
		for k := range dst.Spec.NetworkData.Links.Bonds {
			dst.Spec.NetworkData.Links.Bonds[k].MACAddress = restored.Spec.NetworkData.Links.Bonds[k].MACAddress
		}
		for k := range dst.Spec.NetworkData.Networks.IPv4 {
			dst.Spec.NetworkData.Networks.IPv4[k].FromPoolRef = restored.Spec.NetworkData.Networks.IPv4[k].FromPoolRef
		}
		for k := range dst.Spec.NetworkData.Networks.IPv6 {
			dst.Spec.NetworkData.Networks.IPv6[k].FromPoolRef = restored.Spec.NetworkData.Networks.IPv6[k].FromPoolRef
		}
		for k := range dst.Spec.NetworkData.Links.Bonds {
			dst.Spec.NetworkData.Links.Bonds[k].BondXmitHashPolicy = restored.Spec.NetworkData.Links.Bonds[k].BondXmitHashPolicy
		}
	}

	return nil
}

func (dst *Metal3DataTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3DataTemplate)
	if err := Convert_v1beta1_Metal3DataTemplate_To_v1alpha5_Metal3DataTemplate(src, dst, nil); err != nil {
		return err
	}

	return utilconversion.MarshalData(src, dst)
}

func Convert_v1beta1_NetworkLinkEthernetMac_To_v1alpha5_NetworkLinkEthernetMac(in *v1beta1.NetworkLinkEthernetMac, out *NetworkLinkEthernetMac, s apiconversion.Scope) error {
	// fromAnnotation was added with v1betaX.
	return autoConvert_v1beta1_NetworkLinkEthernetMac_To_v1alpha5_NetworkLinkEthernetMac(in, out, s)
}

func Convert_v1beta1_NetworkDataIPv6_To_v1alpha5_NetworkDataIPv6(in *v1beta1.NetworkDataIPv6, out *NetworkDataIPv6, s apiconversion.Scope) error {
	// fromPoolRef was added with v1beta1.
	return autoConvert_v1beta1_NetworkDataIPv6_To_v1alpha5_NetworkDataIPv6(in, out, s)
}

func Convert_v1beta1_NetworkDataIPv4_To_v1alpha5_NetworkDataIPv4(in *v1beta1.NetworkDataIPv4, out *NetworkDataIPv4, s apiconversion.Scope) error {
	// fromPoolRef was added with v1beta1.
	return autoConvert_v1beta1_NetworkDataIPv4_To_v1alpha5_NetworkDataIPv4(in, out, s)
}

func Convert_v1beta1_FromPool_To_v1alpha5_FromPool(in *v1beta1.FromPool, out *FromPool, s apiconversion.Scope) error {
	// apiGroup and kind was added with v1beta1.
	return autoConvert_v1beta1_FromPool_To_v1alpha5_FromPool(in, out, s)
}

func Convert_v1beta1_NetworkDataLinkBond_To_v1alpha5_NetworkDataLinkBond(in *v1beta1.NetworkDataLinkBond, out *NetworkDataLinkBond, s apiconversion.Scope) error {
	return autoConvert_v1beta1_NetworkDataLinkBond_To_v1alpha5_NetworkDataLinkBond(in, out, s)
}

func (src *Metal3DataTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3DataTemplateList)
	return Convert_v1alpha5_Metal3DataTemplateList_To_v1beta1_Metal3DataTemplateList(src, dst, nil)
}

func (dst *Metal3DataTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3DataTemplateList)
	return Convert_v1beta1_Metal3DataTemplateList_To_v1alpha5_Metal3DataTemplateList(src, dst, nil)
}

func (src *Metal3DataClaim) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3DataClaim)
	if err := Convert_v1alpha5_Metal3DataClaim_To_v1beta1_Metal3DataClaim(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (dst *Metal3DataClaim) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3DataClaim)
	if err := Convert_v1beta1_Metal3DataClaim_To_v1alpha5_Metal3DataClaim(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *Metal3DataClaimList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3DataClaimList)
	return Convert_v1alpha5_Metal3DataClaimList_To_v1beta1_Metal3DataClaimList(src, dst, nil)
}

func (dst *Metal3DataClaimList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3DataClaimList)
	return Convert_v1beta1_Metal3DataClaimList_To_v1alpha5_Metal3DataClaimList(src, dst, nil)
}

func (src *Metal3Remediation) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3Remediation)
	return Convert_v1alpha5_Metal3Remediation_To_v1beta1_Metal3Remediation(src, dst, nil)
}

func (dst *Metal3Remediation) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3Remediation)
	return Convert_v1beta1_Metal3Remediation_To_v1alpha5_Metal3Remediation(src, dst, nil)
}

func (src *Metal3RemediationList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3RemediationList)
	return Convert_v1alpha5_Metal3RemediationList_To_v1beta1_Metal3RemediationList(src, dst, nil)
}

func (dst *Metal3RemediationList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3RemediationList)
	return Convert_v1beta1_Metal3RemediationList_To_v1alpha5_Metal3RemediationList(src, dst, nil)
}

func (src *Metal3RemediationTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3RemediationTemplate)
	return Convert_v1alpha5_Metal3RemediationTemplate_To_v1beta1_Metal3RemediationTemplate(src, dst, nil)
}

func (dst *Metal3RemediationTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3RemediationTemplate)
	return Convert_v1beta1_Metal3RemediationTemplate_To_v1alpha5_Metal3RemediationTemplate(src, dst, nil)
}

func (src *Metal3RemediationTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Metal3RemediationTemplateList)
	return Convert_v1alpha5_Metal3RemediationTemplateList_To_v1beta1_Metal3RemediationTemplateList(src, dst, nil)
}

func (dst *Metal3RemediationTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Metal3RemediationTemplateList)
	return Convert_v1beta1_Metal3RemediationTemplateList_To_v1alpha5_Metal3RemediationTemplateList(src, dst, nil)
}
