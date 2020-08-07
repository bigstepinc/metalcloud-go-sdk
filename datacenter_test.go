package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestDatacenterConfiglUnmarshalTest(t *testing.T) {

	RegisterTestingT(t)

	var dc DatacenterConfig

	err := json.Unmarshal([]byte(_DCConfigFixture), &dc)

	Expect(err).To(BeNil())
	Expect(dc).NotTo(BeNil())
	/*
		Expect(dc.SANRoutedSubnet).To(Equal("100.64.0.0/21"))
		Expect(dc.SwitchProvisioner.Type).To(Equal("VPLSProvisioner"))
		Expect(dc.SwitchProvisioner.Provisioner.(VPLSProvisioner).ACLSAN).To(Equal(3999))
	*/

	Expect(dc.SANRoutedSubnet).To(Equal("100.64.0.0/21"))
	Expect(dc.SwitchProvisioner["type"]).To(Equal("VPLSProvisioner"))
	Expect(dc.SwitchProvisioner["ACLSAN"]).To(Equal(3999.0))

}

const _DCConfigFixture = "{\"SANRoutedSubnet\":\"100.64.0.0/21\",\"BSIVRRPListenIPv4\":\"172.16.10.6\",\"BSIMachineListenIPv4List\":[\"172.16.10.6\"],\"BSIMachinesSubnetIPv4CIDR\":\"10.255.226.0/24\",\"BSIExternallyVisibleIPv4\":\"89.36.24.2\",\"repoURLRoot\":\"https://repointegrationpublic.bigstepcloud.com\",\"repoURLRootQuarantineNetwork\":\"https://repointegrationpublic.bigstepcloud.com\",\"DNSServers\":[\"84.40.63.27\"],\"NTPServers\":[\"84.40.58.44\",\"84.40.58.45\"],\"KMS\":\"\",\"TFTPServerWANVRRPListenIPv4\":\"172.16.10.6\",\"dataLakeEnabled\":false,\"monitoringGraphitePlainTextSocketHost\":\"\",\"monitoringGraphiteRenderURLHost\":\"\",\"latitude\":0,\"longitude\":0,\"address\":\"\",\"switchProvisioner\":{\"type\":\"VPLSProvisioner\",\"ACLSAN\":3999,\"ACLWAN\":3399,\"SANACLRange\":\"3700-3998\",\"ToRLANVLANRange\":\"400-699\",\"ToRSANVLANRange\":\"700-999\",\"ToRWANVLANRange\":\"100-300\",\"quarantineVLANID\":5,\"NorthWANVLANRange\":\"1001-2000\"},\"childDatacentersConfigDefault\":[]}"

func TestDatacenterConfigMarshalTest(t *testing.T) {

	RegisterTestingT(t)

	var dc DatacenterConfig

	err := json.Unmarshal([]byte(_DCConfigFixture), &dc)

	Expect(err).To(BeNil())
	Expect(dc).NotTo(BeNil())

	b, err := json.Marshal(dc)
	Expect(err).To(BeNil())
	Expect(b).NotTo(BeNil())

	var dc2 DatacenterConfig
	err = json.Unmarshal(b, &dc2)
	Expect(err).To(BeNil())

	Expect(dc2.SANRoutedSubnet).To(Equal("100.64.0.0/21"))
	//Expect(dc2.SwitchProvisioner.Type).To(Equal("VPLSProvisioner"))
	Expect(dc2.SwitchProvisioner["type"]).To(Equal("VPLSProvisioner"))
	Expect(dc2.SwitchProvisioner["ACLSAN"]).To(Equal(3999.0))
	//Expect(dc2.SwitchProvisioner.Provisioner.(VPLSProvisioner).ACLSAN).To(Equal(3999))

}
