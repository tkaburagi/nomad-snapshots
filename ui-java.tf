job "ui-java" {
  datacenters = ["dc1"]

  type = "service"

  group "ui-java" {
    count = 1
    task "ui-java" {
      driver = "java"
      artifact {
        source = "https://jar-tkaburagi.s3-ap-northeast-1.amazonaws.com/nomad-shapshots-0.0.1-SNAPSHOT.jar"
      }
      config {
        jar_path    = "local/nomad-shapshots-0.0.1-SNAPSHOT.jar"
        jvm_options = ["-Xmx2048m", "-Xms256m"]
      }
      resources {
        cpu    = 500
        memory = 512
        network {
          port "http" {}
        }
      }
    }
  }
}
