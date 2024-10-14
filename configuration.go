package configuration

import (
	"strconv"
)

/*XX_tip.kv can not set 'NOW' text, like JUSTNOW, RIGHTNOW ... etc. */

const (
	KILOBYTE = 1024
)

var mapKeyID map[string]int
var mapIDValue map[string]string // [configurationID-languageID]value	/	[configurationID]value
var encryptionID []int
var max_buf_size int
var mem_usage int

func init() {
	mapKeyID = make(map[string]int)
	mapIDValue = make(map[string]string)
	max_buf_size = 100 * KILOBYTE
	mem_usage = 0
}

func SetMaxBufSize(kb int) {
	old_size := max_buf_size
	max_buf_size = kb * KILOBYTE
	if max_buf_size < old_size {
		//reduceBuffer
	}
}

func SetKey(key string, configuration_id int, encryption bool) {
	mapKeyID[key] = configuration_id
	if encryption {
		encryptionID = append(encryptionID, configuration_id)
	}
}

func NeedEncryption(configuration_id int) (flag bool) {
	flag = false
	for _, v := range encryptionID {
		if configuration_id == v {
			flag = true
			break
		}
	}
	return
}

func KeyID(key string) (configuration_id int) {
	configuration_id = 0
	if v, ok := mapKeyID[key]; ok {
		configuration_id = v
	}
	return
}

func Set(configuration_id, language_id, cl_id, priority int, value string) {
	/*sz := len(value)
	if mem_usage+sz <= max_buf_size {

	}*/
	v := value
	if priority < 7 { //only save the pointer of database record id.
		v = "@" + strconv.Itoa(configuration_id) + "-" + strconv.Itoa(cl_id) + "@"
	}
	mapIDValue[strconv.Itoa(configuration_id)+"-"+strconv.Itoa(language_id)] = v
}

func SetSimple(configuration_id, priority int, value string) {
	v := value
	if priority < 7 { //only save the pointer of database record id.
		v = "@" + strconv.Itoa(configuration_id) + "@"
	}
	mapIDValue[strconv.Itoa(configuration_id)] = v
}

func GetValue(configuration_id int, language_id string) (value string) {
	value = ""
	if v, ok := mapIDValue[strconv.Itoa(configuration_id)+"-"+language_id]; ok {
		value = v
	}
	return
}

func GetSimpleValue(configuration_id int) (value string) {
	value = ""
	if v, ok := mapIDValue[strconv.Itoa(configuration_id)]; ok {
		value = v
	}
	return
}

func Get(key, language_id string) (value string) {
	configuration_id := KeyID(key)
	if configuration_id > 0 {
		value = GetValue(configuration_id, language_id)
	}
	return
}

func GetSimple(key string) (value string) {
	configuration_id := KeyID(key)
	if configuration_id > 0 {
		value = GetSimpleValue(configuration_id)
	}
	return
}
