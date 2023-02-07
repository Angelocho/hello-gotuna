pipeline {
    agent any
    options { timestamps() 
	ansiColor('xterm')	
	}
    stages {
        
            stage('TestingDocker') {
		steps {
                        sh 'docker-compose config'
                }
            }
            stage('building') {
                steps {
                       sh 'docker-compose build'
                    }
                
            }
	        stage('starting') {
                steps {
                       sh '''docker-compose up -d
                       docker-compose logs -t --tail=10'''
                       
		       echo '\033[1;32m[Success] \033[0m'
                }
        }
        
    }
}
