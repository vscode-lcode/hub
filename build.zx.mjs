#!/usr/bin/env zx

/**@typedef {import("zx/globals")} */

let targets = [
  ["windows", "amd64", "win32-x64"],
  // ["windows", "386", "win32-ia32"],
  ["windows", "arm64", "win32-arm64"],
  ["linux", "amd64", "linux-x64"],
  ["linux", "arm64", "linux-arm64"],
  // ["linux", "armhf", "linux-armhf"],
  ["linux", "amd64", "alpine-x64"],
  ["linux", "arm64", "alpine-arm64"],
  ["darwin", "amd64", "darwin-x64"],
  ["darwin", "arm64", "darwin-arm64"],
];

await $`rm -rf *.vsix`;

for (let index = 0; index < targets.length; index++) {
  let t = targets[index];
  console.log(`[${index + 1}/${targets.length}] progress`);
  await $`GOOS=${t[0]} GOARCH=${t[1]} go build -ldflags="-X 'main.Version=$(git describe --tags --always --dirty)' -s -w" -o bin/lcode-hub .`;
  await $`yarn vsce package -t ${t[2]}`;
}

// 如何发布:
// yarn vsce publish -i *.vsix
