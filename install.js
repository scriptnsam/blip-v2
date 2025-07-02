// install.js
const os = require("os");
const fs = require("fs");
const path = require("path");

const platform = os.platform();
let source = "";
let target = path.join(__dirname, "bin", "blip");

if (platform === "win32") {
  source = path.join(__dirname, "bin", "blip-win.exe");
  target += ".cmd"; // NPM expects a .cmd wrapper for Windows
} else if (platform === "darwin") {
  source = path.join(__dirname, "bin", "blip-macos");
} else if (platform === "linux") {
  source = path.join(__dirname, "bin", "blip-linux");
} else {
  console.error(`Unsupported platform: ${platform}`);
  process.exit(1);
}

fs.copyFileSync(source, target);
fs.chmodSync(target, 0o755);

