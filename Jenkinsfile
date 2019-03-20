pipeline {
  agent {
    docker {
      image 'golang:1.12.1-alpine3.9'
    }

  }
  stages {
    stage('Information') {
      steps {
        sh 'ls -a'
        sh 'go version'
        sh 'go env'
      }
    }
    stage('Test') {
      steps {
        sh 'go test ./cmd/web'
      }
    }
    stage('Build') {
      steps {
        sh 'go test ./cmd/web'
      }
    }
  }
}