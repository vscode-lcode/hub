#!/usr/bin/env zx

/**@typedef {import("zx/globals")} */

const targets = [
  ["windows", "amd64", "win32-x64"],
  ["windows", "386", "win32-ia32"],
  ["windows", "arm64", "win32-arm64"],
  ["linux", "amd64", "linux-x64"],
  ["linux", "arm64", "linux-arm64"],
  // ["linux", "armhf", "linux-armhf"],
  ["linux", "amd64", "alpine-x64"],
  ["linux", "arm64", "alpine-arm64"],
  ["darwin", "amd64", "darwin-x64"],
  ["darwin", "arm64", "darwin-arm64"],
];

const baseUrl =
  // "https://github.com/vscode-lcode/lcode/releases/latest/download/";
  "https://github.com/vscode-lcode/lcode-hub/releases/download/v0.0.5/";

const downloadFailed = Symbol("download failed");

await $`rm -rf *.vsix`;

for (let index = 0; index < targets.length; index++) {
  let t = targets[index];
  console.log(`[${index + 1}/${targets.length}] progress`);
  let binLink = new URL(`lcode-hub-${t[0]}-${t[1]}`, baseUrl);
  let download = $`wget -qO bin/lcode-hub ${binLink}`;
  let result = await download.catch(() => downloadFailed);
  await $`chmod +x bin/lcode-hub`;
  if (result === downloadFailed) {
    console.error(`skip target ${t}`);
    continue;
  }
  await $`yarn vsce package --pre-release -t ${t[2]}`;
}
