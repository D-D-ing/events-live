pipeline {
  agent {
    docker {
      image 'golang:1.12.1-alpine3.9'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh 'cd /go/src && go build -o /go/bin/events-live'
      }
    }
  }
}