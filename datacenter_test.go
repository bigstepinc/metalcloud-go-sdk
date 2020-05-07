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

	Expect(dc.SANRoutedSubnet).To(Equal("100.96.0.0/24"))
	Expect(dc.VLANProvisioner.LANVLANRange).To(Equal("200-299"))

}

const _DCConfigFixture = "{\"SANRoutedSubnet\":\"100.96.0.0/24\",\"BSIVRRPListenIPv4\":\"172.31.240.126\",\"BSIMachineListenIPv4List\":[\"172.31.240.124\",\"172.31.240.125\"],\"BSIMachinesSubnetIPv4CIDR\":\"172.31.240.96/27\",\"BSIExternallyVisibleIPv4\":\"10.255.231.54\",\"repoURLRoot\":\"https://repointegration.bigstepcloud.com\",\"repoURLRootQuarantineNetwork\":\"http://10.255.239.35\",\"DNSServers\":[\"10.255.231.44\",\"10.255.231.45\"],\"NTPServers\":[\"10.255.231.28\",\"10.255.231.29\"],\"KMS\":\"10.255.235.41:1688\",\"TFTPServerWANVRRPListenIPv4\":\"172.31.240.126\",\"dataLakeEnabled\":true,\"monitoringGraphitePlainTextSocketHost\":\"172.31.240.148:2003\",\"monitoringGraphiteRenderURLHost\":\"172.31.240.157:80\",\"latitude\":0,\"longitude\":0,\"address\":\"\",\"VLANProvisioner\":{\"LANVLANRange\":\"200-299\",\"WANVLANRange\":\"100-199\",\"quarantineVLANID\":5},\"VPLSProvisioner\":{\"ACLSAN\":3999,\"ACLWAN\":3399,\"SANACLRange\":\"3700-3998\",\"ToRLANVLANRange\":\"400-699\",\"ToRSANVLANRange\":\"700-999\",\"ToRWANVLANRange\":\"100-399\",\"quarantineVLANID\":5,\"NorthWANVLANRange\":\"1001-2000\"},\"childDatacentersConfigDefault\":[]}"
