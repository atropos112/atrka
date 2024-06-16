{
  pkgs,
  config,
  lib,
  ...
}: {
  pre-commit.hooks = {
    check-merge-conflicts.enable = true;
    check-added-large-files.enable = true;
    gotest.enable = true;
    govet.enable = true;
    gofmt.enable = true;
    golangci-lint.enable = true;
    revive.enable = true;
    staticcheck.enable = true;
  };

  languages.go = {
    enable = true;
    package = pkgs.go;
  };

  scripts = {
    run-producer = {
      exec = ''
        go run producer/main.go 2>&1 | jq
      '';
      description = "Run the producer";
    };
    run-consumer = {
      exec = ''
        go run consumer/main.go 2>&1 | jq
      '';
      description = "Run the consumer";
    };
  };

  enterShell = ''
    echo
    echo 🦾 Useful project scripts:
    echo 🦾
    ${pkgs.gnused}/bin/sed -e 's| |••|g' -e 's|=| |' <<EOF | ${pkgs.util-linuxMinimal}/bin/column -t | ${pkgs.gnused}/bin/sed -e 's|^|🦾 |' -e 's|••| |g'
    ${lib.generators.toKeyValue {} (lib.mapAttrs (_: value: value.description) config.scripts)}
    EOF
    echo
  '';
}
