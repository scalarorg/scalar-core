pipeline {
    agent any
    options {
        skipStagesAfterUnstable()
    }
    stages {
        stage('Build') {
            steps {
                sh 'make USER_ID=$(id -u jenkins) GROUP_ID=$(id -g jenkins) docker-image-test'
            }
        }
        stage('Start scalar network'){
            steps {
                sh 'export IMAGE_TAG_SCALAR_CORE=$(git log -1 --format="%H") && task -t ~/tasks/e2e.yml scalar:up'
                sh 'task -t ~/tasks/e2e.yml scalar:multisig'
                sh 'task -t ~/tasks/e2e.yml scalar:token-deploy'
            }
        }
         stage('Start relayer'){
            steps {
                sh 'task -t ~/tasks/e2e.yml relayer:up'
            }
        }
        stage('Bridging') {
            steps {
                echo 'Bridging tasks'
                // sh 'task -t ~/tasks/e2e.yml bridge:pooling'
                // sh 'task -t ~/tasks/e2e.yml bridge:upc'
            }
        }
        stage('Bridging verification') {
            steps {
                echo 'Bridging verification'
            }
        }
    }
}