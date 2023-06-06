git clone https://github.com/antlr/antlr4.git
pushd antlr4/docker
docker build -t antlr/antlr4 --platform linux/amd64 .
popd
rm -rf antlr4
docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/work antlr/antlr4 -Dlanguage=Go -visitor -o ./pkg/a2l/parser ./grammar/A2L.g4
