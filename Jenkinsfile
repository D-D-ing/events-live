pipeline {
  agent {
    docker {
      image 'golang:1.12.1-stretch'
    }

  }
  stages {
    stage('Information') {
      steps {
        sh 'setenv XDG_CACHE_HOME /tmp/.cache'
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