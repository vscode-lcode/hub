/**
 * @typedef {[string,string,string]} List
 */

/**@type {List[]} */
const lists = [
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

const vsce = require("vsce");
const glob = require("glob");

const preRelease = !!process.env["VSCE_PRE"];
/**
 * @param {boolean} publish
 */
async function main(publish) {
  for (let [OS, ARCH, target] of lists) {
    console.warn(`-----------------------------------`);
    console.warn(`start build target ${target}`);
    process.env["GOOS"] = OS;
    process.env["GOARCH"] = ARCH;
    await vsce.createVSIX({
      target: target,
      useYarn: true,
      preRelease: preRelease,
    });
    console.warn(`finish build target ${target}`);
  }
  if (publish) {
    const packages = glob.sync("*.vsix");
    await vsce.publishVSIX(packages, {
      useYarn: true,
      preRelease: preRelease,
    });
  }
}

main(process.argv[2] === "-p");
