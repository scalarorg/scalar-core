pipeline {
    agent any
    options {
        skipStagesAfterUnstable()
    }
    environment {
        BTC_PRIVATE_KEY = credentials('btc-private-key-zwg')
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
                sh 'task -t ~/tasks/e2e.yml bridge:pooling'
                sh 'task -t ~/tasks/e2e.yml bridge:upc'
            }
        }
        stage('Bridging verification') {
            steps {
                echo 'Bridging verification'
                sh 'task -t ~/tasks/e2e.yml bridge:verify'
            }
        }
        stage('Transfer') {
            steps {
                sh 'task -t ~/tasks/e2e.yml transfer'
            }
        }
        stage('Transfer verification') {
            steps {
                echo 'Transfer verification'
                sh 'task -t ~/tasks/e2e.yml transfer:verify'
            }
        }
        stage('Redeem') {
            steps {
                sh 'task -t ~/tasks/e2e.yml redeem:pooling'
                sh 'task -t ~/tasks/e2e.yml redeem:upc'
            }
        }
        stage('Redeem verification') {
            steps {
                echo 'Redeem verification'
                sh 'task -t ~/tasks/e2e.yml redeem:verify'
            }
        }
    }
}