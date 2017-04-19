# SOON\_ FM 2.0 Stats Event Relay

Captures events from the SOON\_ FM Pub/Sub service, gathers data and places
a job onto an appropriate queue for stats / metric tracking.

* `player:play`
  When this event is received the system will place the Playlist ID on
  an appropriate queue for Spotify, Google or Sound Cloud to call the
  appropriate API for information about the track

* `player:stop`
  When this event is received the system will place the Playlist ID on
  the queue to remove the metric

## Development

1. Create your `go` workspace: `mkdir -p ~/go/src/eventstatsrelay && cd ~/go/src/eventstatsrelay`
2. Clone the repo into the workspace: `git clone git@github.com:thisissoon-fm/statseventrelay.git .`
3. Edit Code
4. Run tests: `make test`
5. Build: `go build` or `make build`
