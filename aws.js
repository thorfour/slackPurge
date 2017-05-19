var cp = require("child_process");
var DEBUG = false

exports.handler = function(event, context) {

    // Parse our the request from the body
    var queryStr = unescape(event.body)

    if (DEBUG) {
        console.log(queryStr)
    }

    // Spawn the go routine to lookup stock quote
    var proc = cp.spawnSync("./gostock", [queryStr], {stdio: 'pipe', encoding: "utf8"});
    var resp = proc.stdout;

    // Return json
    context.succeed(resp);
};
