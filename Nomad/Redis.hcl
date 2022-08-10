job "redis" {
  datacenters = ["dc1"]

  group "cache" {
    count = 1
    task "redis" {
      driver = "docker"

      config {
        image = "redis:3.2"

        port_map {
          db = 6379
        }
      }

      resources {
        cpu = 500
        memory = 256 
      }
    }
    network {
      port "db" {
        static = 6379
      } 
    }
  }
}

