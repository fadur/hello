{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go_1_22
    pkgs.git
  ];
  shellHook = ''
    export GO111MODULE=on
    export GOPATH=$(mktemp -d)
    export PATH=$PATH:$GOPATH/bin
  '';

}
