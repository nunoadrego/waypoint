project = "waypoint-on-minikube"

labels = { "tech" = "golang" }

app "hello-waypoint" {
  build {
    use "pack" {}
    registry {
      use "docker" {
        image = "hello-waypoint"
        tag   = "0.0.1"
        local = true
      }
    }
  }

  deploy {
    use "kubernetes" {
      probe_path   = "/"
      replicas     = 1
      service_port = 8080
      probe {
        initial_delay = 4
      }
      labels = {
        env = "local"
      }
      annotations = {
        demo = "yes"
      }
    }
  }

  release {
    use "kubernetes" {
      port          = 8080
      load_balancer = true
    }
  }
}
