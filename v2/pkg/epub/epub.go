package epub

import "archive/zip"

type Epub struct {
	zip *zip.ReadCloser

	Package  map[string]string
	Metadata []MetadataItem
	Manifest map[string]ManifestItem
	Spine    []SpineItem
	Guide    []GuideItem
}

func (e Epub) GetTitle() string {
	for _, item := range e.Metadata {
		if item.Name == "dc:title" {
			return item.Value
		}
	}
	return e.zip.File[0].Name
}

type MetadataItem struct {
	Name  string
	Value string
	Attrs map[string]string
}

type ManifestItem struct {
	ID        string
	Href      string
	MediaType string
}

type SpineItem struct {
	IDRef  string
	Linear bool
}

type GuideItem struct {
	Href  string
	Title string
	Type  string
}
