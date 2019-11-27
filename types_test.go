package metalcloud

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCheckID(t *testing.T) {
	RegisterTestingT(t)

	//check int ok
	Expect(checkID(100)).To(BeNil())

	//check wrong id.
	Expect(checkID(-100).Error()).To(ContainSubstring("less than 0"))

	//check good string
	Expect(checkID("test-as1")).To(BeNil())

	//check wrong string
	Expect(checkID("_asdad_").Error()).To(ContainSubstring("must be a label format"))

	//check partially correct string
	Expect(checkID("test_id").Error()).To(ContainSubstring("must be a label format"))

	//check wrong type
	Expect(checkID(100.0).Error()).To(ContainSubstring("must be an int"))

}
