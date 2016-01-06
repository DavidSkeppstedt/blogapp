package credentialsutil

import "errors"
import "io/ioutil"
import "log"
import "strings"

func readConfig(path string) (text string, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	text = string(data)
	return
}

func CreateConfig(path string) (dbConfig string, err error) {
	textfile, err := readConfig(path)
	if err != nil {
		log.Println("Missing config file at path", path)
		return "", errors.New("file missing")
	}
	config := strings.Split(textfile, "\n")
	config = config[:len(config)-1]
	if len(config) != 6 {
		log.Println("Missing arguments in file", path)
		return "", errors.New("Argument missmatch")
	}
	dbConfig = "host=" + config[0] +
		" port=" + config[1] +
		" user=" + config[2] +
		" password=" + config[3] +
		" dbname=" + config[4] +
		" sslmode=" + config[5]
	return dbConfig, err
}
