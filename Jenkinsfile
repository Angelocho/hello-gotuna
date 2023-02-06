pipeline {
    agent any

    stages {
        
            stage('TestingDocker') {
                timestamps {
                    steps {
                        sh 'docker-compose config'
                    }
                }
            }
            stage('building') {
                timestamps {
                    steps {
                        sh 'docker-compose build'
                    }
                }
            }
	        stage('starting') {
	            timestamps {
                    steps {
                        sh 'docker-compose up -d'
                    }
	            }
        }
        
    }
}
