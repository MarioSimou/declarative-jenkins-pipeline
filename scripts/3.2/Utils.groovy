 def auditTools() {
     sh '''
        go version
        git --version
     '''
 }