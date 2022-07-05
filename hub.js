const { spawn } = require("child_process");
const fetch = require("node-fetch-commonjs").default;
const { getLcodeAddr } = require("./config");

class Hub {
  constructor() {
    this.timeout = 10 * 1e3;
    /**
     * @private
     */
    this.initPromise = null;
  }
  init() {
    if (this.initPromise) return this.initPromise;
    this.initPromise = Promise.resolve().then(() => this._init());
    return this.initPromise;
  }
  /**
   * @private
   */
  async _init() {
    const healthApi = new URL("/health", getLcodeAddr()).toString();
    let resp = await fetch(healthApi).catch(() => null);
    if (resp !== null && resp.status === 200) {
      return;
    }
    const binPath = require.resolve("./bin/lcode-hub");
    const proc = spawn(binPath, { detached: true });
    let start = Date.now();
    while (true) {
      let now = Date.now();
      if (now - start > this.timeout) {
        throw new Error("start lcode-hub timeout");
      }
      resp = await fetch(healthApi).catch(() => null);
      if (resp !== null && resp.status === 200) {
        return;
      }
      if (proc.exitCode !== null) {
        throw new Error(`start lcode-hub failed. exit code: ${proc.exitCode}`);
      }
    }
  }
  async dispose() {
    const exitApi = new URL("/exit", getLcodeAddr()).toString();
    await fetch(exitApi);
  }
}
module.exports = {
  Hub,
};
