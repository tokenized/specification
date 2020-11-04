


{{- define "CreateAmendmentField" -}}
	{{- if .IsList }}
	{{ .Name }}Min := len(a.{{ .Name }})
	if {{ .Name }}Min > len(newValue.{{ .Name }}) {
		{{ .Name }}Min = len(newValue.{{ .Name }})
	}

	// Compare values
	for i:=0;i<{{ .Name }}Min;i++ {
		lfip := append(fip, uint32(i))
		{{- if .IsCompoundType }}
		{{ .Name }}Amendments, err := a.{{ .Name }}[i].CreateAmendments(lfip,
			newValue.{{ .Name }}[i])
		if err != nil {
			return nil, errors.Wrapf(err, "{{ .Name }}%d", i)
		}
		result = append(result, {{ .Name }}Amendments...)
		{{- else if eq .BaseType "fixedchar" "varchar" }}
		if a.{{ .Name }}[i] != newValue.{{ .Name }}[i] {
			result = append(result, &internal.Amendment{
				FIP: lfip,
				Operation: 0, // Modify element
				Data: []byte(newValue.{{ .Name }}[i]),
			})
		}
		{{- else if eq .BaseType "bin" }}
		if !bytes.Equal(a.{{ .Name }}[i][:], newValue.{{ .Name }}[i][:]) {
			result = append(result, &internal.Amendment{
				FIP: lfip,
				Operation: 0, // Modify element
				Data: newValue.{{ .Name }}[i][:],
			})
		}
		{{- else if eq .BaseType "varbin" }}
		if !bytes.Equal(a.{{ .Name }}[i], newValue.{{ .Name }}[i]) {
			result = append(result, &internal.Amendment{
				FIP: lfip,
				Operation: 0, // Modify element
				Data: newValue.{{ .Name }}[i],
			})
		}
		{{- else if eq .BaseType "uint" }}
		if a.{{ .Name }}[i] != newValue.{{ .Name }}[i] {
			var buf bytes.Buffer
			if err := WriteBase128VarInt(&buf, int(newValue.{{ .Name }}[i])); err != nil {
				return nil, errors.Wrapf(err, "{{ .Name }} %d", i)
			}

			result = append(result, &internal.Amendment{
				FIP: lfip,
				Operation: 0, // Modify element
				Data: buf.Bytes(),
			})
		}
		{{- else if eq .BaseType "bool" }}
		if a.{{ .Name }}[i] != newValue.{{ .Name }}[i] {
			var buf bytes.Buffer
			if err := binary.Write(&buf, binary.LittleEndian, newValue.{{ .Name }}[i]); err != nil {
				return nil, errors.Wrapf(err, "{{ .Name }} %d", i)
			}

			result = append(result, &internal.Amendment{
				FIP: lfip,
				Operation: 0, // Modify element
				Data: buf.Bytes(),
			})
		}
		{{- end }}
	}

	{{ .Name }}Max := len(a.{{ .Name }})
	if {{ .Name }}Max < len(newValue.{{ .Name }}) {
		{{ .Name }}Max = len(newValue.{{ .Name }})
	}

	// Add/Remove values
	for i:={{ .Name }}Min;i<{{ .Name }}Max;i++ {
		amendment := &internal.Amendment{
			FIP: append(fip, uint32(i)), // Add array index to path
		}

		if i < len(newValue.{{ .Name }}) {
			amendment.Operation = 1 // Add element
		{{- if .IsCompoundType }}
			b, err := proto.Marshal(newValue.{{ .Name }}[i])
			if err != nil {
				return nil, errors.Wrapf(err, "serialize {{ .Name }} %d", i)
			}
			amendment.Data = b
		{{- else if eq .BaseType "fixedchar" "varchar" }}
			amendment.Data = []byte(newValue.{{ .Name }}[i])
		{{- else if eq .BaseType "bin" }}
			amendment.Data = newValue.{{ .Name }}[i][:]
		{{- else if eq .BaseType "varbin" }}
			amendment.Data = newValue.{{ .Name }}[i]
		{{- else if eq .BaseType "uint" }}
			var buf bytes.Buffer
			if err := WriteBase128VarInt(&buf, int(newValue.{{ .Name }}[i])); err != nil {
				return nil, errors.Wrapf(err, "{{ .Name }} %d", i)
			}
			amendment.Data = buf.Bytes()
		{{- else if eq .BaseType "bool" }}
			var buf bytes.Buffer
			if err := binary.Write(&buf, binary.LittleEndian, newValue.{{ .Name }}[i]); err != nil {
				return nil, errors.Wrapf(err, "{{ .Name }} %d", i)
			}
			amendment.Data = buf.Bytes()
		{{- end }}
		} else {
			amendment.Operation = 2 // Remove element
		}

		result = append(result, amendment)
	}

	{{- else }}
		{{- if eq .Name "AssetType" }}
		// AssetType modifications not allowed
		{{- else if eq .Name "AssetPayload" }}
	if a.AssetType != newValue.AssetType {
		return nil, fmt.Errorf("Asset type modification not allowed : %s -> %s", a.AssetType,
			newValue.AssetType)
	}

	payloadAmendments, err := assets.CreatePayloadAmendments(fip, []byte(a.AssetType),
		a.AssetPayload, newValue.AssetPayload)
	if err != nil {
		return nil, errors.Wrap(err, "{{ .Name }}")
	}
	result = append(result, payloadAmendments...)
		{{- else if .IsCompoundType }}

	{{ .Name }}Amendments, err := a.{{ .Name }}.CreateAmendments(fip, newValue.{{ .Name }})
	if err != nil {
		return nil, errors.Wrap(err, "{{ .Name }}")
	}
	result = append(result, {{ .Name }}Amendments...)
		{{- else if eq .BaseType "fixedchar" "varchar" }}
	if a.{{ .Name }} != newValue.{{ .Name }} {
		result = append(result, &internal.Amendment{
			FIP: fip,
			Data: []byte(newValue.{{ .Name }}),
		})
	}
		{{- else if eq .BaseType "bin" }}
	if !bytes.Equal(a.{{ .Name }}[:], newValue.{{ .Name }}[:]) {
		result = append(result, &internal.Amendment{
			FIP: fip,
			Data: newValue.{{ .Name }}[:],
		})
	}
		{{- else if eq .BaseType "varbin" }}
	if !bytes.Equal(a.{{ .Name }}, newValue.{{ .Name }}) {
		result = append(result, &internal.Amendment{
			FIP: fip,
			Data: newValue.{{ .Name }},
		})
	}
		{{- else if eq .BaseType "uint" }}
	if a.{{ .Name }} != newValue.{{ .Name }} {
		var buf bytes.Buffer
		if err := WriteBase128VarInt(&buf, int(newValue.{{ .Name }})); err != nil {
			return nil, errors.Wrap(err, "{{ .Name }}")
		}

		result = append(result, &internal.Amendment{
			FIP: fip,
			Data: buf.Bytes(),
		})
	}
		{{- else if eq .BaseType "bool" }}
	if a.{{ .Name }} != newValue.{{ .Name }} {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.{{ .Name }}); err != nil {
			return nil, errors.Wrap(err, "{{ .Name }}")
		}

		result = append(result, &internal.Amendment{
			FIP: fip,
			Data: buf.Bytes(),
		})
	}
		{{- end }}
	{{- end }}
{{- end }}
