pipeline {
    agent any

    environment {
        DEMO="1.3"
    }

    stages {
        stage('First stage') {
            steps {
                echo "This is $BUILD_NUMBER of demo $DEMO"
                sh 'echo "hello world" > test.sh && chmod +x ./test.sh && ./test.sh'
            }
        }
    }
}