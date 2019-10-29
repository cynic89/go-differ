package differ

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"testing"
)

func TestDiffer(t *testing.T)  {
	RegisterFailHandler(Fail)
	RunSpecs(t, "/")
}

var _  = Describe("Differ", func() {

	Context("happy path", func() {

		var(
			primaryYaml = `
a: A
b:
 c: C
d: This is extra
`
			toCompareJson = `{
	"a": "A",
	"b": { "c": "C"}
}
`

		)

		It("should return the difference in keys when two serialized objects are given", func() {

			expected := map[string]interface{}{"d": "This is extra"}
			diff, err := Diff(Yaml{[]byte(primaryYaml)}, Json{[]byte(toCompareJson)})

			Expect(err).To(BeNil())
			Expect(diff).To(BeEquivalentTo(expected))

		})
	})
})