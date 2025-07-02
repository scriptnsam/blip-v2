#!/usr/bin/env node

const {execFileSync} = require('child_process');
const path = require('path');
const os = require('os');

const getBinaryPath = () => {
  const platform = os.platform();
  const arch = os.arch();

  let binaryName = 'blip';
  let binaryPath = '';

  switch (`${platform}-${arch}`) {
    case 'linux-x64':
      binaryPath = path.join(__dirname, 'bin', `${binaryName}-linux`);
      break;
    case 'win32-x64':
      binaryPath = path.join(__dirname, 'bin', `${binaryName}-windows.exe`);
      break;
    case 'darwin-x64': // Intel macOS
      binaryPath = path.join(__dirname, 'bin', `${binaryName}-macos`);
      break;
    default:
      console.error(`Unsupported platform/architecture: ${platform}-${arch}`);
      process.exit(1);
  }
  return binaryPath;
};

try {
  const binary = getBinaryPath();
  execFileSync(binary, process.argv.slice(2), {stdio: 'inherit'});
} catch (error) {
  console.error(`Error executing CLI: ${error.message}`);
  process.exit(1);
}
