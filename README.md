# go-video-streaming-practice

Practice project for video streaming in Go.

## Requirements

- Go

- Any modern browser that supports HTML5 video tag and has its own developer tools to inspect network traffic (e.g., Chrome)

## Setup

Add `big_video.mp4` to `media/` directory.

## Run

```bash
$ go run cmd/service/main.go
```

Open browser and open following files to see the result.

Optionally open Chrome DevTools and check the network tab to see how the video is streamed.

- public/progressive-download/index.html

## Future Plans

- [ ] Implement adaptive HTTP streaming
