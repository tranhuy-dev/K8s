function build() {
    docker login -u "tqhuy1996developer" -p "34f93368-d658-44d7-89eb-cd0519f90b03"
    docker build . -t "tqhuy1996developer/istock"
    docker push "tqhuy1996developer/istock"
}

build