pipeline {
    agent any

    stages {
        
            stage('TestingDocker') {
                options {
		    timestamps()
		}
		steps {
                        sh 'docker-compose config'
                }
            }
            stage('building') {
                options {
		    timestamps()
		}
                steps {
                    timestamps {
                        sh 'docker-compose build'
                    }
                }
                
            }
	        stage('starting') {
                options {
		    timestamps()
		}
                steps {
                    timestamps {
                        sh 'docker-compose up -d'
                    }
                }
        }
        
    }
}
