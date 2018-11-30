package collector

import (
	"io/ioutil"
	"os"
	"testing"
)

var mockData = `<domainCapabilities>
  <cpu>
    <mode name='host-passthrough' supported='yes'/>
    <mode name='host-model' supported='yes'>
      <model fallback='allow'>Skylake-Client-IBRS</model>
      <vendor>Intel</vendor>
      <feature policy='require' name='ds'/>
      <feature policy='require' name='acpi'/>
      <feature policy='require' name='ss'/>
    </mode>
    <mode name='custom' supported='yes'>
      <model usable='no'>EPYC-IBPB</model>
      <model>fake-model-without-usable</model>
      <model usable='yes'>Haswell</model>
    </mode>
  </cpu>
</domainCapabilities>`

func writeMockDataFile(path, data string) error {
	err := ioutil.WriteFile(path, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

func deleteMockFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// TestCollectData test CollectData function
func TestCollectData(t *testing.T) {
	filePath := "/tmp/virsh-domcapabilities.xml"
	err := writeMockDataFile(filePath, mockData)
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}

	result, err := CollectData(filePath)
	if err != nil {
		t.Error("CollectData should not throw error: " + err.Error())
	}

	if len(result) != 1 {
		t.Error("CollectData should return one cpu model")
	}

	if len(result) == 1 {
		if result[0] != "haswell" {
			t.Error("cpu model should equal to haswell")
		}

		if result[0] == "Haswell" {
			t.Error("cpu model should contain lower cased chars")
		}
	}

	err = deleteMockFile(filePath)
	if err != nil {
		t.Error("deleteMockFile should not throw error: " + err.Error())
	}

	result, err = CollectData("")
	if err == nil {
		t.Error("CollectData should throw error, because of empty path")
	}

	if result != nil {
		t.Error("CollectData should return nil")
	}

	err = writeMockDataFile(filePath, "pat a mat")
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}

	_, err = CollectData(filePath)
	if err == nil {
		t.Error("CollectData should throw error, because data in file are not in xml")
	}

	deleteMockFile(filePath)
}
