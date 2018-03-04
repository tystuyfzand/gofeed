package ext

import (
	"strconv"
	"strings"
)

// MediaExtension represents a feed extension
// for the Media specification.
type MediaExtension struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
	Categories []*MediaCategory `json:"categories,omitempty"`
	Thumbnails []*MediaThumbnail `json:"thumbnails,omitempty"`
	Hashes []*MediaHash `json:"hashes,omitempty"`
}

type MediaContent struct {
	URL string `json:"url,omitempty"`
	FileSize int64 `json:"fileSize,omitempty"`
	Type string `json:"type,omitempty"`
	Medium string `json:"medium,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	Expression string `json:"expression,omitempty"`
	Bitrate int `json:"bitrate,omitempty"`
	Framerate int `json:"framerate,omitempty"`
	SamplingRate float32 `json:"samplingrate,omitempty"`
	Channels int `json:"channels,omitempty"`
	Duration int `json:"duration,omitempty"`
	Width int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
	Language string `json:"lang,omitempty"`
}

type MediaCategory struct {
	Scheme string `json:"scheme,omitempty"`
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type MediaHash struct {
	Algorithm string `json:"algo,omitempty"`
	Hash string `json:"hash,omitempty"`
}

type MediaCredit struct {
	Role string `json:"role,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	Credit string `json:"credit,omitempty"`
}

type MediaThumbnail struct {
	Width int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
	URL string `json:"url,omitempty"`
}

// NewMediaExtension creates a new MediaExtension
// given the generic extension map for the "media" prefix.
func NewMediaExtension(extensions map[string][]Extension) *MediaExtension {
	media := &MediaExtension{}
	media.Title = parseTextExtension("title", extensions)
	media.Description = parseTextExtension("description", extensions)
	media.Keywords = parseDelimitedTextExtension("keywords", ",", extensions)
	media.Thumbnails = parseMediaThumbnails(extensions)
	media.Categories = parseMediaCategories(extensions)
	media.Hashes = parseMediaHashes(extensions)
	return media
}

func parseMediaThumbnails(extensions map[string][]Extension) (thumbnails []*MediaThumbnail) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["thumbnail"]
	if !ok || len(matches) == 0 {
		return
	}

	thumbnails = []*MediaThumbnail{}
	for _, thumb := range matches {
		m := &MediaThumbnail{}

		if url, ok := thumb.Attrs["url"]; ok {
			m.URL = url
		}

		if width, ok := thumb.Attrs["width"]; ok {
			m.Width, _ = strconv.Atoi(width)
		}

		if height, ok := thumb.Attrs["width"]; ok {
			m.Height, _ = strconv.Atoi(height)
		}

		thumbnails = append(thumbnails, m)
	}
	return
}

func parseMediaCategories(extensions map[string][]Extension) (categories []*MediaCategory) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["category"]
	if !ok || len(matches) == 0 {
		return
	}

	categories = []*MediaCategory{}
	for _, cat := range matches {
		c := &MediaCategory{}

		if scheme, ok := cat.Attrs["scheme"]; ok {
			c.Scheme = scheme
		}

		if label, ok := cat.Attrs["label"]; ok {
			c.Label = label
		}

		c.Value = cat.Value

		categories = append(categories, c)
	}
	return
}

func parseMediaHashes(extensions map[string][]Extension) (hashes []*MediaHash) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["hash"]
	if !ok || len(matches) == 0 {
		return
	}

	hashes = []*MediaHash{}
	for _, hash := range matches {
		h := &MediaHash{}

		if algo, ok := hash.Attrs["algo"]; ok {
			h.Algorithm = algo
		}

		h.Hash = hash.Value

		hashes = append(hashes, h)
	}
	return
}

func parseDelimitedTextExtension(name, delimiter string, extensions map[string][]Extension) []string {
	val := parseTextExtension(name, extensions)

	vals := strings.Split(val, delimiter)

	ret := make([]string, 0)

	for _, val := range vals {
		ret = append(ret, strings.TrimSpace(val))
	}

	return ret
}