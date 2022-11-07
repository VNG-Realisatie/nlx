#!/usr/bin/env node
/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

/* eslint-disable @typescript-eslint/no-var-requires */
var os = require("os");
var child_process = require("child_process");
var config_file =
  "./conf/" + (process.env.E2E_CONFIG_FILE || "default") + ".conf.js";
var config = require(config_file).config;
var command = "/usr/bin/env";

process.argv = [
  "node",
  "./node_modules/@cucumber/cucumber/bin/cucumber",
  "features/**/*.feature",
  "--parallel",
  `${process.env.E2E_PARALLEL_COUNT || 8}`,
  "--tags",
  `${process.env.E2E_TESTS_TAGS || ""}`,
];

const retries = process.env.E2E_FLAKY_RETRIES || 0;
if (retries > 0) {
  process.argv.concat([
    "--retry",
    `${retries}`,
    "--retry-tag-filter",
    "@flaky",
  ]);
}

// Check if os is windows
if (os.platform() === "win32") {
  command = process.argv.shift();
}

const spawnSubProcess = (options) =>
  new Promise((resolve, reject) => {
    const p = child_process.spawn(command, process.argv, options);
    p.stdout.pipe(process.stdout);
    p.stderr.pipe(process.stderr);

    p.on("close", (code) => {
      if (code > 0) {
        reject(code);
      } else {
        resolve(code);
      }
    });
  });

const subProcesses = config.capabilities.map((capability, i) => {
  const env = Object.create(process.env);
  env.TASK_ID = i.toString();
  return spawnSubProcess({ env });
});

Promise.all(subProcesses)
  .then(() => {
    process.exit();
  })
  .catch(() => {
    process.exit(1);
  });
