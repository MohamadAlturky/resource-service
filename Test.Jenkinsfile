pipeline {
    agent any

    environment {
        REPO_URL = 'https://github.com/MohamadAlturky/resource-service.git'
        BRANCH = 'test'
        CREDENTIALS_ID = 'github'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: "${BRANCH}", url: "${REPO_URL}", credentialsId: "${CREDENTIALS_ID}"
            }
        }

        stage('Run Docker Compose') {
            steps {
                script {
                    sh 'k6 --version'
                    sh "docker-compose -f docker-compose.test.yaml up --build -d"
                    sh "k6 tests/k6.js"
                    sh "docker-compose -f docker-compose.test.yaml down"
                }
            }
        }
    }
}