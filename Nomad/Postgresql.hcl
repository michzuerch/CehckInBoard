job "database" {
  datacenters = ["dc1"]
  type = "service"

  group "db" {
    count = 1

    task "postgres" {
      driver = "docker"

      config {
        image = "postgres:14.4"
        network_mode = "host"
        port_map {
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
    }
  }
}
