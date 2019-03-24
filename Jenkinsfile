pipeline {
  agent { dockerfile true }
  stages {
    stage('Information') {
      steps {
        sh 'echo ${UID}'
        sh 'ls -la'
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
