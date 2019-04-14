package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func NewProtocolResources(filenames []string) ProtocolResources {
	items := ProtocolResources{}

	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		m := ProtocolResource{}
		if err := yaml.Unmarshal(data, &m); err != nil {
			panic(fmt.Errorf("file %v : %s", filename, err))
		}

		// Strip header
		buf := bufio.NewReader(bytes.NewBuffer(data))
		var output bytes.Buffer
		started := false
		for {
			line, _, err := buf.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(fmt.Errorf("file %v : %s", filename, err))
			}
			if started {
				output.Write(line)
				output.WriteByte(byte('\n'))
			} else if strings.HasPrefix(string(line), "values:") {
				started = true
			}
		}

		m.Data = string(output.Bytes())
		items = append(items, m)
	}

	return items
}

type ProtocolResources []ProtocolResource

type ProtocolResource struct {
	Metadata Metadata
	Data     string
}

func (m ProtocolResource) Name() string {
	return strings.Replace(m.Metadata.Name, " ", "", -1)
}
