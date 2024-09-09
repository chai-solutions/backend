{
  description = "Description for the project";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    flake-parts.url = "github:hercules-ci/flake-parts";
  };

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      imports = [];

      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];

      perSystem = {pkgs, ...}: {
        packages.default = pkgs.buildGoModule {
          pname = "chai-solutions-backend";
          version = "0.0.1";
          src = pkgs.nix-gitignore.gitignoreSource [] ./.;

          vendorHash = "sha256-NA5fgx+uzQQKmAuJwGIBtTGU9QcrM54BiGQ+kz03pYk=";

          meta = with pkgs.lib; {
            description = "Chai Solutions backend";
            homepage = "https://chai-solutions.org";
            license = licenses.gpl3Only;
            maintainers = with maintainers; [water-sucks];
          };
        };

        devShells.default = pkgs.mkShell {
          name = "chai-solutions-backend";
          buildInputs = with pkgs; [
            go
            gotools
          ];
        };
      };
    };
}
