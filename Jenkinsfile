pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh 'cd /go/src && go build -o /go/bin/hello-app'
      }
    }
  }
}