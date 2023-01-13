const vscode = require("vscode");
const { UriHandler } = require("./lcode-uri");

/**
 * @param {vscode.ExtensionContext} context
 */
async function activate(context) {
  // const hub = new Hub();
  // context.subscriptions.push(hub);

  // const hubInit = hub.init();
  const hubInit = Promise.resolve(0);
  const lcodeUriHandler = new UriHandler(hubInit);
  context.subscriptions.push(vscode.window.registerUriHandler(lcodeUriHandler));

  // lcodeUriHandler.preTask.then(() => {
  //   const opener = new Opener(lcodeUriHandler);
  //   context.subscriptions.push(opener);
  // });
}

// this method is called when your extension is deactivated
function deactivate() {}

module.exports = {
  activate,
  deactivate,
};
