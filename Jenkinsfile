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
			echo "\033[1;32m[Success] \033[0m $1"
                    }
                }
        }
        
    }
}
