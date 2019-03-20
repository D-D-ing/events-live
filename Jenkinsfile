pipeline {
  agent {
    docker {
      image 'golang:1.12.1-alpine3.9'
    }

  }
  stages {
    stage('Test') {
      steps {
        sh 'go test ./cmd/web'
      }
    }
    stage('Build') {
      steps {
        sh 'go build -o ./cmd/web'
      }
    }
  }
}