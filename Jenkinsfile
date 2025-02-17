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
                sh '~/tvl_maker --db-path ~/tvl_maker.db tx --test-env ~/envs/e2e/tools.env  --service-tag pools bridge custodian-only \
--amount 2000 \
--wallet-address tb1q2rwweg2c48y8966qt4fzj0f4zyg9wty7tykzwg \
--private-key ${BTC_PRIVATE_KEY} \
--destination-chain 0100000000AA36A7 \
--destination-token-address 156E2E35CCA2e106003a53c7bF96584A83A8BaBd \
--destination-recipient-address 982321eb5693cdbAadFfe97056BEce07D09Ba49f'
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