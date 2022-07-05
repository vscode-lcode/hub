function getLcodeAddr() {
  return process.env["LCODE_ADDR"] || "http://127.0.0.1:4349";
}

module.exports = {
  getLcodeAddr,
};
