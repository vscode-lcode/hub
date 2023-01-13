const vscode = require("vscode");

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
  let u = `webdav://${uri.path.slice(1)}`;
  if (uri.query) u += "?" + uri.query;
  if (uri.fragment) u += "#" + uri.fragment;
  return u;
}

module.exports = {
  UriHandler,
  getWebdavUri,
};
