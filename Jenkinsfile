pipeline {
    agent any
    options { timestamps() }
    stages {
        
            stage('TestingDocker') {
		steps {
                        sh 'docker-compose config'
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
