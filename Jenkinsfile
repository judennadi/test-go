pipeline {
  agent {
    label {
      label 'docker-go-agent'
      retries 2
    }
  }
  
  stages {
    stage('Build') {
      steps {
          echo 'Building...'
          sh '''
          go build -o server .
          '''
      }
    }

    stage('Test') {
      steps {
          echo 'Testing...'
          sh '''
          go test -v
          '''
      }
    }

    stage('Deploy') {
      steps {
          echo 'Deploying...'
          
      }
    }
  }
}
