export GOPATH=$PWD
echo $GOPATH

sh deploy/apt_update.sh
sh deploy/install_go.sh
sh deploy/install_deps.sh