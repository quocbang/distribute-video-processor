package api

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/quocbang/distribute-video-processor/adapters/interfaces/api/model"
	"github.com/quocbang/distribute-video-processor/domain/task"
	"github.com/quocbang/distribute-video-processor/proto/quality"
)

func (a API) NewVideoHandlerRoutes(g *echo.Group) {
	g.POST("/upload", a.UploadVideo)
}

func (a API) UploadVideo(c echo.Context) error {
	var (
		req model.UploadVideoRequest
		// ctx = context.Background()
	)

	fh, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"error_code": "8129372",
				"info":       "failed to open request file",
			},
		})
	}

	if err := req.ParseMultipleQualitiesQuery(c.QueryParams()); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"error_code": "8129376",
				"info":       "failed to parse qualities",
				"details":    err,
			},
		})
	}

	f, err := fh.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"error_code": "1000023",
				"info":       "failed to open request file",
			},
		})
	}
	defer f.Close()

	// upload video to s3
	input, err := io.ReadAll(f)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"error_code": "100000",
				"info":       "failed to init task",
			},
		})
	}
	objectKey, err := a.Cloud.S3.UploadLargeObject("video/hls/mp4/test.mp4", input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"error_code": "100000",
				"info":       "failed to upload video",
			},
		})
	}

	// save to disk
	dir := "./tmp/test.mp4"
	file, err := os.Create(dir)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"error_code": "100000",
				"info":       "failed to create file",
				"detail":     err,
			},
		})
	}
	defer file.Close()
	// defer os.Remove(dir)

	_, err = file.Write(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"error_code": "100000",
				"info":       "failed to write file",
				"detail":     err,
			},
		})
	}

	log.Println(dir)

	for _, q := range req.Qualities {
		task, err := a.Task.Distributor.ConvertHLSVideoTask(context.Background(), task.ConvertHLSVideo{
			S3FilePath:   objectKey, // maybe change to s3 url
			TempDir:      dir,
			VideoQuality: quality.Quality(q),
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"data": map[string]interface{}{
					"error_code": "100000",
					"info":       "failed to init task",
				},
			})
		}

		info, err := a.Task.Client.Enqueue(task, asynq.MaxRetry(4))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"data": map[string]interface{}{
					"error_code": "100000",
					"info":       "failed to init task",
				},
			})
		}

		log.Println("queue info", info)
	}

	// Create and upload the master playlist
	// masterPlaylist := `
	// 		#EXTM3U
	// 		#EXT-X-STREAM-INF:BANDWIDTH=800000,RESOLUTION=640x360
	// 		360p.m3u8
	// 		#EXT-X-STREAM-INF:BANDWIDTH=1400000,RESOLUTION=854x480
	// 		480p.m3u8
	// 		#EXT-X-STREAM-INF:BANDWIDTH=2800000,RESOLUTION=1280x720
	// 		720p.m3u8
	// 		#EXT-X-STREAM-INF:BANDWIDTH=5000000,RESOLUTION=1920x1080
	// 		1080p.m3u8
	// `
	// masterPlaylistPath := filepath.Join(tempDir, "master.m3u8")
	// err = os.WriteFile(masterPlaylistPath, []byte(masterPlaylist), 0644)
	// if err != nil {
	// 	fmt.Printf("Error writing master playlist: %v\n", err)
	// 	return err
	// }

	// err = uploadFileToS3(masterPlaylistPath, bucket, "video/hls/master.m3u8")
	// if err != nil {
	// 	fmt.Printf("Error uploading master playlist to S3: %v\n", err)
	// 	return err
	// }

	return c.JSON(http.StatusAccepted, "OK")
}
