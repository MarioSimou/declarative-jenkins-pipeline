 def auditTools() {
     node {
        sh '''
            go version
            git --version
        '''
     }
 }