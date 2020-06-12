job "api-go" {
  datacenters = ["dc1"]

  type = "service"

  group "api-go" {
    count = 3
    task "api-go" {
      driver = "raw_exec"
      config {
        command = "/usr/local/bin/go"
        args = [
          "run",
        "/Users/kabu/hashicorp/nomad/snapshots-demo/api.go"]
      }
    }
  }
}
