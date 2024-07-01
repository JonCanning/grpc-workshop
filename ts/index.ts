import { startServer } from "server";
import { startClient } from "client";
import * as readline from "readline";

readline.emitKeypressEvents(process.stdin);
if (process.stdin.isTTY) process.stdin.setRawMode(true);
console.log("Press 1 to start server, 2 to start client");
process.stdin.on("keypress", async (_, key) => {
  switch (key.name) {
    case "1":
      await startServer();
      break;
    case "2":
      await startClient(() => process.exit(0));
      break;
    default:
      console.log("Invalid option");
      process.exit(1);
  }
  process.stdin.setRawMode(false);
});
