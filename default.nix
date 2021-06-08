{
  pkgs ? import <nixpkgs> {},
  delveBinary ? "",
  doCheck ? true
}:

let pkg = pkgs.buildGoPackage rec {
  name = "delve-go-client-sdk-${version}";
  version = "1.1.3";
  goPackagePath = "github.com/RelationalAI/go-client-sdk";
  src = ./.;
  goDeps = ./deps.nix;
};

in

pkg.overrideAttrs(oldAttrs: rec {
  nativeBuildInputs = oldAttrs.nativeBuildInputs ++ [ delveBinary ];
  checkPhase = ''
    mkdir home
    export HOME=$PWD/home
    delve server &
    PID=$!
    sleep 15s
    go test -v -count=1 -tags=integration github.com/RelationalAI/go-client-sdk/sdk || (kill -9 $PID && exit 1)
    echo "Shutting down delve server. Pid: $PID"
    kill -9 $PID
  '';
  inherit doCheck;
})