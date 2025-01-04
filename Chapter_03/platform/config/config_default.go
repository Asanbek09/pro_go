package config

import "strings"

type DefaultConfig struct {
	configData map[string]interface{}
}

func (c *DefaultConfig) get(name string) (result interface{}, found bool) {
	data := c.configData
	for _, key := range strings.Split(name, ":") {
		var ok bool
		result, ok = data[key]
		if !ok { // Если ключ не найден
			found = false
			return
		}
		// Если это карта, спускаемся на уровень ниже
		if newSection, isMap := result.(map[string]interface{}); isMap {
			data = newSection
		} else if len(strings.Split(name, ":")) > 1 { // Если ожидаем вложенность, но её нет
			found = false
			return
		}
	}
	found = true
	return
}

func (c *DefaultConfig) GetSection(name string) (section Configuration, found bool) {
	value, found := c.get(name)
	if found {
		if sectionData, ok := value.(map[string]interface{}); ok {
			section = &DefaultConfig{configData: sectionData}
		} 
	}
	return
}

func (c *DefaultConfig) GetString(name string) (string, bool) {
	value, found := c.get(name)
	if !found {
		return "", false
	}
	// Приводим значение к строке
	strValue, ok := value.(string)
	if !ok {
		return "", false
	}
	return strValue, true
}

func (c *DefaultConfig) GetInt(name string) (result int, found bool) {
	value, found := c.get(name)
	if(found){ result = int(value.(float64)) }
	return
}

func (c *DefaultConfig) GetBool(name string) (result bool, found bool) {
	value, found := c.get(name)
	if (found) { result = value.(bool)}
	return
}

func (c *DefaultConfig) GetFloat(name string) (result float64, found bool) {
	value, found := c.get(name)
	if (found) { result = value.(float64)}
	return
}