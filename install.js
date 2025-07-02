const fs = require("fs");
const path = require("path");
const os = require("os");

const platform = os.platform();
let src;

if (platform === "win32") {
  src = path.join(__dirname, "bin", "blip-win.exe");
} else if (platform === "darwin") {
  src = path.join(__dirname, "bin", "blip-macos");
} else if (platform === "linux") {
  src = path.join(__dirname, "bin", "blip-linux");
} else {
  console.error("Unsupported platform:", platform);
  process.exit(1);
}

const dest = path.join(__dirname, "blip");

fs.copyFileSync(src, dest);
fs.chmodSync(dest, 0o755);

