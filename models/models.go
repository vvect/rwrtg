package models

// TODO(cduplooy): I had to absolutely butcher this.
// TODO(cduploy): Can do with a rewrite at some point.

type Profile struct {
	Metadata struct {
		Runtime string `json:"runtime"`
	} `json:"meta"`
}

type AndroidProfile struct {
	Classes map[string]AndroidClass `json:"classes"`
}

type TempAndroidProfile struct {
	Entries []AndroidEntry `json:"data"`
}

type AndroidEntry struct {
	Loader  *string `json:"loader"`
	Classes []struct {
		Name      string   `json:"name"`
		Methods   []string `json:"methods"`
		Overloads map[string]struct {
			ArgumentTypes [][]AndroidParameter `json:"argTypes"`
			ReturnTypes   *[]interface{}       `json:"returnType"`
		} `json:"overloads"`
	} `json:"classes"`
}

func (tap TempAndroidProfile) GetAndroidProfile() AndroidProfile {
	p := AndroidProfile{
		Classes: make(map[string]AndroidClass),
	}
	for _, entry := range tap.Entries {
		for _, class := range entry.Classes {
			var loader string
			if entry.Loader == nil {
				loader = "default"
			} else {
				loader = *entry.Loader
			}
			p.Classes[class.Name] = AndroidClass{
				Name:    class.Name,
				Loader:  loader,
				Methods: make(map[string]AndroidMethod),
			}
			for _, method := range class.Methods {
				returnTypes := make([]string, 0)
				if class.Overloads[method].ReturnTypes == nil {
					returnTypes = append(returnTypes, "void")
				} else {
					for _, r := range *class.Overloads[method].ReturnTypes {
						mmap, _ := r.(map[string]interface{})
						if mmap["className"] != nil {
							returnTypes = append(returnTypes, mmap["className"].(string))
						} else {
							returnTypes = append(returnTypes, class.Name)
						}
					}
				}
				p.Classes[class.Name].Methods[method] = AndroidMethod{
					Name:          method,
					ArgumentTypes: class.Overloads[method].ArgumentTypes,
					ReturnTypes:   &returnTypes,
				}
			}
		}
	}
	return p
}

type AndroidClass struct {
	Name    string                   `json:"name"`
	Methods map[string]AndroidMethod `json:"methods"`
	Loader  string                   `json:"loader"`
}

type AndroidMethod struct {
	Name          string               `json:"name"`
	ArgumentTypes [][]AndroidParameter `json:"argTypes"`
	ReturnTypes   *[]string            `json:"returnType"`
}

type AndroidParameter struct {
	ClassName string `json:"className"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
}

type TempIOSEntry struct {
	Name    string `json:"name"`
	Address uint64 `json:"address"`
}

type TempIOSProfile struct {
	Classes []TempIOSEntry `json:"classes"`
}

type IOSProfile struct {
	Classes map[string][]string `json:"classes"`
}
