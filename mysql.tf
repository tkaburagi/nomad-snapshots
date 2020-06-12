job "mysql-5.7.22" {
  datacenters = ["dc1"]

  type = "service"

  group "mysql-group" {
    count = 1

    volume "mysql-vol" {
      type = "host"
      read_only = false
      source = "mysql-vol"
    }

    task "mysql-task" {
      driver = "docker"
      volume_mount {
        volume = "mysql-vol"
        destination = "/var/lib/mysql"
        read_only = false
      }

      config {
        image = "mysql:5.7.28"
        port_map {
          db = 3306
        }
      }
      env {
        "MYSQL_ROOT_PASSWORD" = "rooooot"
      }
      resources {
        cpu = 500
        memory = 512

        network {
          mbits = 10
          port "db" {
            static = 3306
          }
        }
      }
    }
  }
}