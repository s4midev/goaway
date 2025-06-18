# goaway
## A featherweight, Dockerised, headless application for removing stale or failed arr downloads
### And I mean *really* featherweight (31MB Average Idle)

## Example Compose
```yaml
services:
  goaway:
    container_name: goaway
    build: .
    restart: unless-stopped
    volumes:
      - ./:/goaway
    environment:
      - GOAWAY_RADARR_URL=http://192.168.0.85:7878
      - GOAWAY_RADARR_API_KEY=radarrKey
      - GOAWAY_SONARR_URL=http://192.168.0.85:8989
      - GOAWAY_SONARR_API_KEY=sonarrKey
      # ms
      - GOAWAY_CHECK_INTERVAL=1000
```