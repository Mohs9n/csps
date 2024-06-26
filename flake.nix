# This flake was initially generated by fh, the CLI for FlakeHub (version 0.1.8)
{

  # Flake inputs
  inputs = {
    flake-compat.url = "https://flakehub.com/f/edolstra/flake-compat/*.tar.gz";

    flake-schemas.url = "https://flakehub.com/f/DeterminateSystems/flake-schemas/*.tar.gz";

    nixpkgs.url = "https://flakehub.com/f/NixOS/nixpkgs/*.tar.gz";
  };

  # Flake outputs that other flakes can use
  outputs = { self, flake-compat, flake-schemas, nixpkgs }:
    let
      # Helpers for producing system-specific outputs
      supportedSystems = [ "x86_64-linux" ];
      forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
        pkgs = import nixpkgs { inherit system; };
      });
    in {
      # Schemas tell Nix about the structure of your flake's outputs
      schemas = flake-schemas.schemas;

      # Development environments
      devShells = forEachSupportedSystem ({ pkgs }: {
        default = pkgs.mkShell {
          # Pinned packages available in the environment
          packages = with pkgs; [
            go_1_21
            nixpkgs-fmt
          ];

          # Environment variables
          env = {
            PORT = "8080";
          };

          # A hook run every time you enter the environment
          shellHook = ''
            echo "HELLO"
          '';
        };
        

      });

      packages = forEachSupportedSystem ({ pkgs }: {
        default = pkgs.buildGoModule {
          name = "csps";
          src = ./.;
          vendorHash = "sha256-CWiegMXsvWBCUtOCkuK3/VBJeEpthVcr1J1D067np+c=";
        };
      });
    };
}
