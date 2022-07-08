const vscode = require("vscode");
const fetch = require("node-fetch-commonjs").default;

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
    uri = vscode.Uri.parse(u);
    await this.preTask;
    const opt = {};
    opt.forceNewWindow = true;

    const isDir = await fetch(u.replace("webdav", "http"))
      .then((r) => r.text())
      .then((xml) => /D:collection/.test(xml));
    if (!isDir) {
      await vscode.commands.executeCommand("vscode.open", uri, opt);
      return;
    }
    await vscode.commands.executeCommand("vscode.openFolder", uri, opt);
  }
}

/**
 * @param {vscode.Uri} uri
 * @returns {string}
 */
function getWebdavUri(uri) {
  let u = `webdav://127.0.0.1:4349/proxy${uri.path}`;
  if (uri.query) u = "?" + uri.query;
  if (uri.fragment) u = "#" + uri.fragment;
  return u;
}

module.exports = {
  UriHandler,
  getWebdavUri,
};
