with import <nixpkgs> {};

stdenv.mkDerivation {

  name = "go-release";

  buildInputs = with pkgs; [
    jetbrains.goland
  ];

  shellHook = ''
    goland .
    exit
  '';
}

