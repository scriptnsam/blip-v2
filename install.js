const fs = require("fs");
const os = require("os");
const path = require("path");

const platform = os.platform();

let binName;

if (platform === "win32") {
  binName = "blip-win.exe";
} else if (platform === "darwin") {
  binName = "blip-macos";
} else {
  binName = "blip-linux";
}

const src = path.join(__dirname, "bin", binName);
const dest = path.join(__dirname, "blip");

try {
  fs.copyFileSync(src, dest);
  fs.chmodSync(dest, 0o755);
  console.log(`Installed blip binary for ${platform}`);
} catch (err) {
  console.error("Error installing binary:", err.message);
  process.exit(1);
}

