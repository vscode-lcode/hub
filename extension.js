const vscode = require("vscode");
const { Hub } = require("./hub");
const { UriHandler } = require("./lcode-uri");
const { Opener } = require("./opener");

/**
 * @param {vscode.ExtensionContext} context
 */
async function activate(context) {
  const hub = new Hub();
  context.subscriptions.push(hub);

  const hubInit = hub.init();
  const lcodeUriHandler = new UriHandler(hubInit);
  context.subscriptions.push(vscode.window.registerUriHandler(lcodeUriHandler));

  const opener = new Opener(lcodeUriHandler);
  context.subscriptions.push(opener);
}

// this method is called when your extension is deactivated
function deactivate() {}

module.exports = {
  activate,
  deactivate,
};
