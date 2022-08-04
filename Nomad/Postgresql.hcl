job "database" {
  datacenters = ["dc1"]
  type = "service"

  group "db" {
    count = 1

    network {
      port "db" {
        static = 5432
      }
    }

    service {
      port = "postgres"
    }

    task "postgres" {
      driver = "docker"

      config {
        image = "postgres:14.4"
        network_mode = "host"
        port_map db {
          db = 5432
        } 
      }

      env {
        POSTGRES_DB = "CheckInBoard"
        POSTGRES_USER = "michzuerch"
        POSTGRES_PASSWORD = "lx0lc33a"
      }

      logs {
        max_files = 5
        max_file_size = 15
      }

      resources {
        cpu    = 1000 # MHz
        memory = 1024 # MB
      }

      service {
        name = "postgres"
        tags = ["postgres for CheckInBoard"]
        port = "db"
        /* check { */
        /*   name = "alive" */
        /*   type = "tcp" */
        /*   interval = "10s" */
        /*   timeout = "2s" */
        /* } */
      }
      restart {
        attempts = 10
        interval = "5m"
        delay = "25s"
        mode = "delay"
      }
    }
  }
  update {
    max_parallel = 1
    min_healthy_time = "5s"
    healthy_deadline = "3m"
    auto_revert = false
    canary = 0
  }
}
