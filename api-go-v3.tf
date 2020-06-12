job "api-go" {
  datacenters = ["dc1"]
  update {
    max_parallel      = 2
    health_check      = "task_states"
    min_healthy_time  = "5s"
    healthy_deadline  = "10m"
    progress_deadline = "20m"
    canary            = 1
  }
  type = "service"

  group "api-go" {
    count = 5
    task "api-go" {
      driver = "raw_exec"
      config {
        command = "/usr/local/bin/go"
        args = [
          "run",
        "/Users/kabu/hashicorp/nomad/snapshots-demo/api-go-v3.go"]
      }
    }
  }
}
