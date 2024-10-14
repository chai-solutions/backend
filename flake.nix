{
  description = "Chai Solutions backend daemon";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    flake-parts.url = "github:hercules-ci/flake-parts";

    process-compose-flake.url = "github:Platonic-Systems/process-compose-flake";
  };

  outputs = {flake-parts, ...} @ inputs:
    flake-parts.lib.mkFlake {inherit inputs;} {
      imports = [inputs.process-compose-flake.flakeModule];

      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];

      perSystem = {pkgs, ...}: {
        packages.default = pkgs.buildGoModule {
          pname = "chaid";
          version = "0.0.1";
          src = pkgs.nix-gitignore.gitignoreSource [] ./.;

          vendorHash = "sha256-IWoL6Xt1Z+zYgK7Hq90VG26/J5oXQsxdNclvxUnq2yI=";

          postInstall = ''
            mv $out/bin/cmd $out/bin/chaid
          '';

          meta = with pkgs.lib; {
            description = "Chai Solutions backend daemon";
            homepage = "https://chai-solutions.org";
            license = licenses.gpl3Only;
            maintainers = with maintainers; [water-sucks];
          };
        };

        devShells = {
          default = pkgs.mkShell {
            name = "chaid-shell";
            buildInputs = with pkgs; [
              go
              gotools
              golangci-lint
              postgresql_16
              dbmate
              sqlc
            ];

            shellHook = ''
              root_dir="$(git rev-parse --show-toplevel)"

              export PGHOST="$root_dir/.postgres"
              export PGDATA="$PGHOST/data"
              export PGDATABASE="chai"
              export PGLOG="$PGHOST/postgres.log"

              if [ ! -d $PGDATA ]; then
                mkdir -p $PGDATA
                initdb -U postgres $PGDATA --auth=trust --no-locale --encoding=UTF8 > /dev/null

                echo "CREATE USER chai;" | postgres --single -D $PGDATA postgres > /dev/null
                echo "CREATE DATABASE chai OWNER chai;" | postgres --single -D $PGDATA postgres > /dev/null
              fi
            '';
          };
          ci = pkgs.mkShell {
            name = "chaid-ci-shell";
            buildInputs = with pkgs; [
              go
              golangci-lint
            ];
          };
        };

        process-compose = {
          localpg = {
            settings.processes = {
              postgres-server = {
                command = ''pg_ctl start -l '$PGLOG' -o "--unix_socket_directories='$PGHOST'"'';
                is_daemon = true;
                shutdown = {command = "pg_ctl stop";};
              };
            };
          };
        };
      };
    };
}
