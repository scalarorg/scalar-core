pipeline {
    agent any
    options {
        skipStagesAfterUnstable()
    }
    stages {
        stage('Preparation') {
            steps {
                echo "env:  ${env.getEnvironment()}"
            }
        }
        stage('Build') {
            steps {
                sh 'make docker-image'
            }
        }
        stage('Start'){
            steps {
                sh 'task -t ~/tasks/e2e.yml scalar:up'
            }
        }
        stage('Bridging') {
            steps {
                sh 'task -t ~/tasks/e2e.yml bridge:pooling'
                sh 'task -t ~/tasks/e2e.yml bridge:upc'
            }
        }
        stage('Bridging verification') {
            steps {
                echo 'Bridging verification'
            }
        }
    }
}