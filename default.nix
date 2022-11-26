with import <nixpkgs> {};

stdenv.mkDerivation {

  name = "go-release";

  buildInputs = with pkgs; [
    go_1_18
    gnumake
  ];

  shellHook = ''
    export GOPATH=$HOME/go
    export PATH=$PATH:$HOME/go/bin
  '';

  hardeningDisable = [ "fortify" ];

}


