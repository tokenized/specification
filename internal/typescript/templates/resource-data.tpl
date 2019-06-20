
export class Data {
{{range .}}
{{comment (print .Name " - " .Metadata.Description) "//"}}
	static yaml{{ .Name }} = `
{{ .Data }}
`;
{{ end }}
}