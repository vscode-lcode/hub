const vscode = require("vscode");
const config = require("./config");

/**
 * @implements {vscode.UriHandler}
 */
class UriHandler {
  /**
   * @param {Promise<any>} preTask
   */
  constructor(preTask) {
    /**
     * @type {Promise<any>}
     */
    this.preTask = preTask;
  }
  /**
   * @param {vscode.Uri} uri
   */
  async handleUri(uri) {
    let u = getWebdavUri(uri);
    if (uri.fragment === "file") {
      require("child_process").execSync(`code --file-uri ${u}`);
      return;
    }

    uri = vscode.Uri.parse(u);
    await this.preTask;
    const opt = {};
    opt.forceNewWindow = true;
    opt.noRecentEntry = true;

    await vscode.commands.executeCommand("vscode.openFolder", uri, opt);
  }
}

/**
 * @param {vscode.Uri} uri
 * @returns {string}
 */
function getWebdavUri(uri) {
  const addr = config.getLcodeHubAddr().slice("http://".length);
  let u = `webdav://${addr}/proxy${uri.path}`;
  if (uri.query) u += "?" + uri.query;
  if (uri.fragment) u += "#" + uri.fragment;
  return u;
}

module.exports = {
  UriHandler,
  getWebdavUri,
};
