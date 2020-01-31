package metalcloud

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestStageDefinitionUnmarshalTest(t *testing.T) {
	RegisterTestingT(t)

	var obj StageDefinition
	err := json.Unmarshal([]byte(_stageDefinitionsFixture), &obj)
	Expect(err).To(BeNil())
	Expect(obj).NotTo(BeNil())
	Expect(obj.StageDefinition).ToNot((BeNil()))
	Expect(obj.StageDefinitionID).To(Equal(402))
	Expect(obj.UserIDOwner).To(Equal(8030))

	req := obj.StageDefinition.(HTTPRequest)
	Expect(req.URL).To(Equal("https://raw.githubusercontent.com/bigstepinc/terraform-provider-metalcloud/master/README.md"))
}

func TestStageDefinitionMarshalTest(t *testing.T) {
	RegisterTestingT(t)

	//get an object

	ab := AnsibleBundle{
		AnsibleBundleArchiveFilename: "test",
	}

	obj := StageDefinition{
		StageDefinitionID:    100,
		StageDefinitionLabel: "test",
		StageDefinition:      ab,
		StageDefinitionType:  "AnsibleBundle",
	}

	//we get bytes
	b, err := json.Marshal(obj)
	Expect(err).To(BeNil())
	Expect(b).NotTo(BeNil())

	//unmarshal
	var obj2 StageDefinition
	err = json.Unmarshal(b, &obj2)
	Expect(err).To(BeNil())

	//check that we have ansible bundle
	ab2 := obj2.StageDefinition.(AnsibleBundle)
	Expect(ab2.AnsibleBundleArchiveFilename).To(Equal(ab.AnsibleBundleArchiveFilename))

}

const _stageDefinitionsFixture = "{\"stage_definition_id\":402,\"user_id_owner\":8030,\"user_id_authenticated\":8030,\"stage_definition_is_deprecated\":false,\"stage_definition_label\":\"my-test-stage-3\",\"stage_definition_label_unique\":\"my-test-stage-3@8030\",\"stage_definition_title\":\"my title\",\"stage_definition_description\":\"\",\"stage_definition_type\":\"HTTPRequest\",\"icon_asset_data_uri\":null,\"stage_definition_created_timestamp\":\"2020-01-28T17:20:03Z\",\"stage_definition_updated_timestamp\":\"2020-01-28T17:20:03Z\",\"stage_definition_variable_names_required\":[],\"stage_definition\":{\"url\":\"https://raw.githubusercontent.com/bigstepinc/terraform-provider-metalcloud/master/README.md\",\"type\":\"HTTPRequest\",\"options\":{\"body\":null,\"size\":67108864,\"follow\":40,\"method\":\"GET\",\"headers\":{\"Pragma\":\"no-cache\",\"Connection\":\"close\",\"User-Agent\":\"MetalCloud/1.0 (WebFetchAPI)\",\"Cache-Control\":\"no-cache\"},\"timeout\":240000,\"compress\":true,\"redirect\":\"follow\",\"bodyBufferBase64\":null}}}"
