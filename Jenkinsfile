pipeline {
    agent any

    stages {
        timestamps {
            stage('TestingDocker') {
                steps {
                   sh 'docker-compose config'
                }
            }
        }
        timestamps {
            stage('building') {
                steps {
                    sh 'docker-compose build'
                }
            }
        }
        timestamps {
	        stage('starting') {
                steps {
                    sh 'docker-compose up -d'
                }
            }
        }
        
    }
}
