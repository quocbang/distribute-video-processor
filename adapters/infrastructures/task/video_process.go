package task

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hibiken/asynq"
	"github.com/quocbang/distribute-video-processor/domain/task"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

func (t *Task) ConvertHLSVideoProcess(ctx context.Context, aTask *asynq.Task) error {
	// get payload
	var payload task.ConvertHLSVideo
	if err := json.Unmarshal(aTask.Payload(), &payload); err != nil {
		return err
	}

	zap.L().Info("starting to open file", zap.String("quality", payload.VideoQuality.String()))
	f, err := os.Open(payload.TempDir)
	if err != nil {
		return err
	}
	zap.L().Info("file is opened", zap.String("quality", payload.VideoQuality.String()))

	// Create a temporary directory for the HLS output
	tempDir, err := os.MkdirTemp("", payload.VideoQuality.String())
	if err != nil {
		return fmt.Errorf("error creating temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up the temp directory after use

	// Processing video with specific resolution
	segmentPattern := filepath.Join(tempDir, payload.VideoQuality.String()+"_segment_%03d.ts")
	playlistPath := filepath.Join(tempDir, payload.VideoQuality.String()+".m3u8")

	// input := bytes.NewReader(f.)
	err = ffmpeg.Input("pipe:").WithInput(f).
		Output(playlistPath, ffmpeg.KwArgs{
			"c:a":                  "aac",
			"b:a":                  "128k",
			"c:v":                  "libx264",
			"b:v":                  payload.GetBitrate(),
			"s":                    fmt.Sprintf("%dx%d", payload.GetWidth(), payload.GetHeight()),
			"hls_list_size":        "0",
			"hls_time":             "10",
			"hls_segment_filename": segmentPattern,
		}).
		OverWriteOutput().Run()
	if err != nil {
		return fmt.Errorf("error processing video (%s): %v", payload.VideoQuality.String(), err)
	}

	// load all file in temp dir and upload to s3
	file, err := os.Open(tempDir)
	if err != nil {
		return fmt.Errorf("error opening file %s: %w", tempDir, err)
	}
	defer file.Close()

	dirs, err := os.ReadDir(tempDir)
	if err != nil {
		log.Printf("failed to read dir, error: %v", err)
		return err
	}
	for _, dir := range dirs {
		f, err := os.Open(tempDir + "/" + dir.Name())
		if err != nil {
			log.Printf("failed to open file in temp dir, error: %v \n", err)
			return err
		}

		buff := &bytes.Buffer{}
		_, err = buff.ReadFrom(f)
		if err != nil {
			log.Printf("failed to read from file, error: %v", err)
			return err
		}
		_, err = t.S3.UploadLargeObject(fmt.Sprintf("video/hls/%s/%s", payload.VideoQuality.String(), dir.Name()), buff.Bytes())
		if err != nil {
			log.Printf("failed to upload video, error: %v", err)
			return err
		}
	}

	log.Println("Successfully uploaded all playlists and segments to S3")

	return nil
}
