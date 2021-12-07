

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

			subPermissions, err := permissions.SubPermissions(fip, operation, {{ .IsList }})
			if err != nil {
				return nil, errors.Wrap(err, "sub permissions")
			}

			return a.{{ .Name }}[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)
		{{- else if eq .BaseType "fixedchar" "varchar" }}
			a.{{ .Name }}[fip[1]] = string(data)
			return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- else if eq .BaseType "bin" }}
			if len(data) != {{ .BaseSize }} {
				return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), {{ .BaseSize }})
			}
			copy(a.{{ .Name }}[fip[1]], data)
			return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- else if eq .BaseType "varbin" }}
			a.{{ .Name }}[fip[1]] = data
			return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- else if eq .BaseType "uint" }}
			buf := bytes.NewBuffer(data)
			if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
			} else {
				a.{{ .Name }}[fip[1]] = {{ .GoSingularType }}(value)
			}
			return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- else if eq .BaseType "bool" }}
			if len(data) != 1 {
				return nil, fmt.Errorf("{{ .Name }} amendment value is wrong size : %d", len(data))
			}
			buf := bytes.NewBuffer(data)
			if err := binary.Read(buf, binary.LittleEndian, &a.{{ .Name }}[fip[1]]); err != nil {
				return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
			}
			return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- end }}

		case 1: // Add element

		{{- if .IsCompoundType }}
			if len(fip) > 2 { // includes list index
				// Add is for sub-object
				subPermissions, err := permissions.SubPermissions(fip, operation, {{ .IsList }})
				if err != nil {
					return nil, errors.Wrap(err, "sub permissions")
				}

				return a.{{ .Name }}[fip[1]].ApplyAmendment(fip[2:], operation, data, subPermissions)

			} else if len(fip) < 2 {
				return nil, fmt.Errorf("Amendment field index path wrong depth for add {{ .Name }} : %v",
					fip)
			}

			newValue := &{{ .GoSingularType }}{}
			if len(data) != 0 { // Leave default values if data is empty
				if err := proto.Unmarshal(data, newValue); err != nil {
					return nil, fmt.Errorf("Amendment addition to {{ .Name }} failed to deserialize : %s",
						err)
				}
			}
		{{- else }}
			if len(fip) != 2 { // includes list index
				return nil, fmt.Errorf("Amendment field index path wrong depth for add {{ .Name }} : %v",
					fip)
			}

			{{ if eq .BaseType "fixedchar" "varchar" }}
			newValue := string(data)
			{{- else if eq .BaseType "bin" }}
			if len(data) != {{ .BaseSize }} {
				return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), {{ .BaseSize }})
			}
			var newValue [{{ .BaseSize }}]byte
			copy(newValue, data)
			a.{{ .Name }} = append(a.{{ .Name }}, newValue)
			{{- else if eq .BaseType "varbin" }}
			newValue := data
			{{- else if eq .BaseType "uint" }}
			var newValue {{ .GoSingularType }}
			buf := bytes.NewBuffer(data)
			if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
				return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s",
					err)
			} else {
				newValue = {{ .GoSingularType }}(value)
			}
			{{- else if eq .BaseType "bool" }}
			if len(data) != 1 {
				return nil, fmt.Errorf("Amendment field index path too deep for {{ .Name }} : %x",
					data)
			}
			buf := bytes.NewBuffer(data)
			var newValue {{ .GoSingularType }}
			if err := binary.Read(buf, binary.LittleEndian, &newValue); err != nil {
				return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
			}
			{{- end }}
		{{- end }}

			if len(a.{{ .Name }}) <= int(fip[1]) {
				// Append item to the end
				a.{{ .Name }} = append(a.{{ .Name }}, newValue)
			} else {
				// Insert item at index specified by fip[1]
				before := a.{{ .Name }}[:fip[1]]
				after := make([]{{ .GoSingularTypeWithPointer }}, len(a.{{ .Name }})-int(fip[1]))
				copy(after, a.{{ .Name }}[fip[1]+1:]) // copy so slice reuse won't overwrite

				a.{{ .Name }} = append(before, newValue)
				a.{{ .Name }} = append(a.{{ .Name }}, after...)
			}
			return permissions.SubPermissions(fip, operation, {{ .IsList }})

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
			return permissions.SubPermissions(fip, operation, {{ .IsList }})
		}
	{{- else }}
		{{- if and (gt (len .BaseResource) 0) (ne .BaseResource "LegalSystems") (ne .BaseResource "Polities") }}
		if {{ .BaseResource }}Data(a.{{ .Name }}) == nil {
			return nil, fmt.Errorf("{{ .BaseResource }} resource value not defined : %v", a.{{ .Name }})
		}
		{{- end }}
		{{- if .IsCompoundType }}

		subPermissions, err := permissions.SubPermissions(fip, operation, {{ .IsList }})
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.{{ .Name }}.ApplyAmendment(fip[1:], operation, data, subPermissions)
		{{- else if eq .BaseType "fixedchar" "varchar" }}
		a.{{ .Name }} = string(data)
		return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- else if eq .BaseType "bin" }}
		if len(data) != {{ .BaseSize }} {
			return nil, fmt.Errorf("bin size wrong : got %d, want %d", len(data), {{ .BaseSize }})
		}
		copy(a.{{ .Name }}, data)
		return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- else if eq .BaseType "varbin" }}
		a.{{ .Name }} = data
		return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- else if eq .BaseType "uint" }}
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for {{ .Name }} : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("{{ .Name }} amendment value failed to deserialize : %s", err)
		} else {
			a.{{ .Name }} = {{ .GoSingularType }}(value)
		}
		return permissions.SubPermissions(fip, operation, {{ .IsList }})
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
		return permissions.SubPermissions(fip, operation, {{ .IsList }})
		{{- end }}
	{{- end }}
{{- end }}
