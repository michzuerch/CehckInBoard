job "hello-world" {
  datacenters = ["dc1"]
  group "hello-world" {
    network {
      port "http" {
        to = 5000
      }
    }
    task "hello-world" {
      driver = "docker"
      config {
        image = "${var.CI_GITHUB_IMAGE}"
        ports = [ "http" ]
      } 
    }
    service {
      name "hello-world"
      tags = [
        "urlprefix-${var.CI_PROJECT_PATH_SLUG}-${var.CI_COMMIT_REF_SLUG}.${var.BASE_DOMAIN}:443/"
      ]
      port = "http"
      check {
        type = "http"
        port = "http"
        path = "/"
        interval = "10s"
        timeout = "2s"
      }
    }
  }
}
