{
  pkgs ? import <nixpkgs> {},
  raiserverBinary ? "",
  doCheck ? true
}:

let pkg = pkgs.buildGoPackage rec {
  name = "rai-server-go-client-sdk-${version}";
  version = "1.1.3";
  goPackagePath = "github.com/RelationalAI-oss/relationalai-sdk-go";
  src = ./.;
  goDeps = ./deps.nix;
};

in

pkg.overrideAttrs(oldAttrs: rec {
  nativeBuildInputs = oldAttrs.nativeBuildInputs ++ [ raiserverBinary ];
  checkPhase = ''
    mkdir home
    export HOME=$PWD/home
    rai-server server &
    PID=$!
    sleep 15s
    go test -v -count=1 -tags=integration github.com/RelationalAI-oss/relationalai-sdk-go/sdk || (kill -9 $PID && exit 1)
    echo "Shutting down rai-server server. Pid: $PID"
    kill -9 $PID
  '';
  inherit doCheck;
})