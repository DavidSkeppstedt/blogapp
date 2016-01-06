package credentialsutil

import "testing"

func TestOpenNonExisitingFile(t *testing.T) {
	_, err := readConfig("main.env")
	if err == nil {
		t.Errorf("Not handling none exisiting file")
	}
}
func TetsOpenExisitingFile(t *testing.T) {
	_, err := readConfig("../config/db.conf")
	if err != nil {
		t.Errorf("Not handling exisiting file correctly")
	}
}
func TestCreateConfigFromExisitingFile(t *testing.T) {
	_, err := CreateConfig("../config/db.conf")
	if err != nil {
		if err.Error() == "file" {
			t.Errorf("Not hanlding exisiting file correctly")
		}
	}
}
func TestCreateConfigToFewArguments(t *testing.T) {
	_, err := CreateConfig("../config/db.conf")
	if err != nil {
		if err.Error() != "Argument missmatch" {
			t.Errorf("Casting wrong error on exisiting file")
		}
	}
}

func TestCreateConfigFromNonExisitingFile(t *testing.T) {
	_, err := CreateConfig("db.conf")
	if err == nil || err.Error() != "file missing" {
		t.Errorf("Not reporting file missing error ...")
	}
}
