pipeline {
    agent any

    stages {
        
            stage('TestingDocker') {
                steps {
                        timestamps {
                        sh 'docker-compose config'
                        }
                }
            }
            stage('building') {
                steps {
                    timestamps {
                        sh 'docker-compose build'
                    }
                }
                
            }
	        stage('starting') {
                steps {
                    timestamps {
                        sh 'docker-compose up -d'
                    }
                }
        }
        
    }
}
