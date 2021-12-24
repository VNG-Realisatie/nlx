#!/usr/bin/env node
/* eslint-disable @typescript-eslint/no-var-requires */
var os = require("os");
var child_process = require("child_process");
var config_file =
  "./conf/" + (process.env.E2E_CONFIG_FILE || "default") + ".conf.js";
var config = require(config_file).config;
var command = "/usr/bin/env";

process.argv[0] = "node";
process.argv[1] = "./node_modules/@cucumber/cucumber/bin/cucumber-js";
process.argv[2] = "features/**/*.feature";
process.argv[3] = "--parallel";
process.argv[4] = `${process.env.E2E_PARALLEL_COUNT || 8}`;
process.argv[5] = "--tags";
process.argv[6] = "not @ignore";

// Check if os is windows
if (os.platform() == "win32") {
  command = process.argv.shift();
}

for (var i in config.capabilities) {
  var env = Object.create(process.env);
  env.TASK_ID = i.toString();
  var p = child_process.spawn(command, process.argv, { env: env });
  p.stdout.pipe(process.stdout);
  p.stderr.pipe(process.stderr);
}
