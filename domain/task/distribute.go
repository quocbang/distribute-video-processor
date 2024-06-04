package task

import (
	"github.com/quocbang/distribute-video-processor/proto/quality"
)

type ConvertHLSVideo struct {
	S3FilePath   string          `json:"s3_file_path"`
	TempDir      string          `json:"temp_dir"`
	VideoQuality quality.Quality `json:"video_quality"`
}

type resolution struct {
	Width   int
	Height  int
	Bitrate string
}

var qualityResolutions = map[quality.Quality]resolution{}

func (c ConvertHLSVideo) GetWidth() int {
	resolution := qualityResolutions[c.VideoQuality]
	return resolution.Width
}

func (c ConvertHLSVideo) GetHeight() int {
	resolution := qualityResolutions[c.VideoQuality]
	return resolution.Height
}

func (c ConvertHLSVideo) GetBitrate() string {
	resolution := qualityResolutions[c.VideoQuality]
	return resolution.Bitrate
}

func init() {
	qualityResolutions[quality.Quality_Q360P] = resolution{
		Width:   640,
		Height:  360,
		Bitrate: "800k",
	}
	qualityResolutions[quality.Quality_Q480P] = resolution{
		Width:   854,
		Height:  480,
		Bitrate: "1400k",
	}
	qualityResolutions[quality.Quality_Q720P] = resolution{
		Width:   1280,
		Height:  720,
		Bitrate: "2800k",
	}
	qualityResolutions[quality.Quality_Q1080P] = resolution{
		Width:   1920,
		Height:  1080,
		Bitrate: "5000k",
	}
}
