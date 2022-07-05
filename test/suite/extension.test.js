const assert = require("assert");

// You can import and use all API from the 'vscode' module
// as well as import your extension to test it
const vscode = require("vscode");
const { getWebdavUri } = require("../../lcode-uri");

suite("Extension Test Suite", () => {
  vscode.window.showInformationMessage("Start all tests.");

  test("Sample test", () => {
    const uri = vscode.Uri.parse(
      "vscode://lcode.hub/localhost_0e_0e_0e_0e_0e/home/username/editdir"
    );
    const wlink = getWebdavUri(uri);
    assert.equal(
      wlink,
      "webdav://127.0.0.1:4349/proxy/localhost_0e_0e_0e_0e_0e/home/username/editdir"
    );
  });
});
