{
  description = "Chai Solutions backend";

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
          pname = "chai-solutions-backend";
          version = "0.0.1";
          src = pkgs.nix-gitignore.gitignoreSource [] ./.;

          vendorHash = "sha256-sIkfsjg/Sxo58gOFzpEQwv2kS9NFy094hOpP0xl9xYQ=";

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
            postgresql_16
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
              echo "GRANT ALL ON SCHEMA public TO chai;" | postgres --single -D $PGDATA postgres > /dev/null
              echo "CREATE DATABASE chai;" | postgres --single -D $PGDATA postgres > /dev/null
            fi
          '';
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
