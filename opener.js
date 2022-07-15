const EventSource = require("eventsource");
const vscode = require("vscode");
const { getLcodeHubAddr } = require("./config");

class Opener {
  /**
   * @param {vscode.UriHandler} handler
   */
  constructor(handler) {
    const api = new URL("/open-handler", getLcodeHubAddr());
    const sse = new EventSource(api.toString());
    sse.addEventListener("open-webdav", (e) => {
      // ignore open webdav event
      if (true) {
        return;
      }
      let uri = vscode.Uri.parse(e.data);
      handler.handleUri(uri);
    });
    /**@type {EventSource} */
    this.sse = sse;
  }
  dispose() {
    this.sse.close();
  }
}

module.exports = {
  Opener,
};
