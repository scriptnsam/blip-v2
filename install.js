// install.js
const {platform, arch} = require("os");
const {copyFileSync, chmodSync} = require("fs");
const {join} = require("path");

let binName;
switch (platform()) {
  case "win32":
    binName = "blip-win.exe";
    break;
  case "darwin":
    binName = "blip-macos";
    break;
  case "linux":
    binName = "blip-linux";
    break;
  default:
    throw new Error(`Unsupported platform: ${platform()}`);
}

const src = join(__dirname, "bin", binName);
const dest = join(__dirname, "blip");

copyFileSync(src, dest);
chmodSync(dest, 0o755); // make it executable
console.log(`Installed blip binary for ${platform()}`);

