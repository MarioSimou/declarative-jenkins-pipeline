 def auditTools() {
     sh '''
        go version
        git --version
     '''
 }

return this