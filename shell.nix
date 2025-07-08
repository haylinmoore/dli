{ pkgs ? import <nixpkgs> { } }:
let
  treefmt-nix = import (builtins.fetchTarball "https://github.com/numtide/treefmt-nix/archive/main.tar.gz");
  treefmt = treefmt-nix.mkWrapper pkgs {
    projectRootFile = ".git/config";
    programs = {
      gofmt.enable = true;
      prettier.enable = true;
      nixpkgs-fmt.enable = true;
    };
  };
in
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    gopls
    gotools
    go-tools
    delve
    treefmt
  ];

  shellHook = ''
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$PATH
  '';
}
