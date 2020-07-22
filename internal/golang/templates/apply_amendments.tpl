

{{- define "ApplyAmendmentField" -}}
	{{- if .IsList }}
		switch operation {
		case 0: // Modify
			{{- if .IsCompoundType }}
			if len(fip) < 3 { // includes list index and subfield index
			{{- else }}
			if len(fip) != 2 { // includes list index
			{{- end }}
				return nil, fmt.Errorf("Amendment field index path incorrect depth for modify {{ .Name }} : %v",
					fip)
			}
			if int(fip[1]) >= len(a.{{ .Name }}) {
				return nil, fmt.Errorf("Amendment element index out of range for modify {{ .Name }} : %d", fip[1])
			}

			{{- if .IsCompoundType }}
			result, err := a.{{ .Name }}[fip[1]].ApplyAmendment(fip[2:], operation, data)
			return append(fip[:1], result...), err
			{{- else if eq .BaseType "fixedchar" "varchar" }}
			a.{{ .Name }}[fip[1]] = string(data)
			return append(fip[:1], fip[2:]...), nil
			{{- else if eq .BaseType "bin" }}
			if len(data) != {{ .Size }} {
				return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), {{ .Size }})
			}
			copy(a.{{ .Name }}[fip[1]], data)
			return append(fip[:1], fip[2:]...), nil
			{{- else if eq .BaseType "varbin" }}
			a.{{ .Name }}[fip[1]] = data
			return append(fip[:1], fip[2:]...), nil
			{{- else if eq .BaseType "uint" }}
			buf := bytes.NewBuffer(data)
			if value, err := ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
			} else {
				a.{{ .Name }}[fip[1]] = {{ .GoSingularType }}(value)
			}
			return append(fip[:1], fip[2:]...), nil
			{{- else if eq .BaseType "bool" }}
			if len(data) != 1 {
				return nil, fmt.Errorf("{{ .Name }} amendment value is wrong size : %d", len(data))
			}
			buf := bytes.NewBuffer(data)
			if err := binary.Read(buf, binary.LittleEndian, &a.{{ .Name }}[fip[1]]); err != nil {
				return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
			}
			return append(fip[:1], fip[2:]...), nil
			{{- end }}

		case 1: // Add element
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for add {{ .Name }} : %v",
					fip)
			}

			{{- if .IsCompoundType }}
			newValue := &{{ .GoSingularType }}{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to {{ .Name }} failed to deserialize : %s",
						err)
				}
			}
			{{- else if eq .BaseType "fixedchar" "varchar" }}
			newValue := string(data)
			{{- else if eq .BaseType "bin" }}
			if len(data) != {{ .Size }} {
				return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), {{ .Size }})
			}
			var newValue [{{ .BaseSize }}]byte
			copy(newValue, data)
			a.{{ .Name }} = append(a.{{ .Name }}, newValue)
			{{- else if eq .BaseType "varbin" }}
			newValue := data
			{{- else if eq .BaseType "uint" }}
			var newValue {{ .GoSingularType }}
			buf := bytes.NewBuffer(data)
			if value, err := ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
			} else {
				newValue = {{ .GoSingularType }}(value)
			}
			{{- else if eq .BaseType "bool" }}
			if len(fip) > 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for {{ .Name }} : %v", fip)
			}
			if len(data) != {{ .BaseSize }} {
				return nil, fmt.Errorf("{{ .Name }} amendment value is wrong size : %d", len(data))
			}
			buf := bytes.NewBuffer(data)
			var newValue {{ .GoSingularType }}
			if err := binary.Read(buf, binary.LittleEndian, &newValue); err != nil {
				return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
			}
			{{- end }}
			a.{{ .Name }} = append(a.{{ .Name }}, newValue)
			return fip[:], nil

		case 2: // Delete element
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path incorrect depth for delete {{ .Name }} : %v",
					fip)
			}
			if int(fip[1]) >= len(a.{{ .Name }}) {
				return nil, fmt.Errorf("Amendment element index out of range for delete {{ .Name }} : %d", fip[1])
			}

			// Remove item from list
			a.{{ .Name }} = append(a.{{ .Name }}[:fip[1]], a.{{ .Name }}[fip[1]+1:]...)
			return append(fip[:1], fip[2:]...), nil
		}
	{{- else }}
		{{- if gt (len .BaseResource) 0 }}
		if {{ .BaseResource }}Data(a.{{ .Name }}) == nil {
			return nil, fmt.Errorf("{{ .BaseResource }} resource value not defined : %v", a.{{ .Name }})
		}
		{{- end }}
		{{- if .IsCompoundType }}
		return a.{{ .Name }}.ApplyAmendment(fip[1:], operation, data)
		{{- else if eq .BaseType "fixedchar" "varchar" }}
		a.{{ .Name }} = string(data)
		return fip[:], nil
		{{- else if eq .BaseType "bin" }}
		if len(data) != {{ .Size }} {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), {{ .Size }})
		}
		copy(a.{{ .Name }}, data)
		return fip[:], nil
		{{- else if eq .BaseType "varbin" }}
		a.{{ .Name }} = data
		return fip[:], nil
		{{- else if eq .BaseType "uint" }}
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for {{ .Name }} : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
		} else {
			a.{{ .Name }} = {{ .GoSingularType }}(value)
		}
		return fip[:], nil
		{{- else if eq .BaseType "bool" }}
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for {{ .Name }} : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("{{ .Name }} amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.{{ .Name }}); err != nil {
			return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
		}
		return fip[:], nil
		{{- end }}
	{{- end }}
{{- end }}
