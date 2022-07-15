const vscode = require("vscode");

/**
 *
 * @returns {string}
 */
function getLcodeHubAddr() {
  const hubAddr = vscode.workspace.getConfiguration("lcode.hub").get("addr");
  return hubAddr;
}

module.exports = {
  getLcodeHubAddr,
};
