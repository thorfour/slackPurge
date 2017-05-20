/**
 * Responds to any HTTP request that can provide a "message" field in the body.
 *
 * @param {!Object} req Cloud Function request context.
 * @param {!Object} res Cloud Function response context.
 */
var cp = require("child_process");

exports.handler = function(req, res) {
    var queryStr = unescape(req.body);
    var proc = cp.spawnSync("./slackPurge", [queryStr], {stdio: 'pipe', encoding: "utf8"});
    var resp = proc.stdout;
    console.log(req.body.message);
    res.status(200).json(JSON.parse(resp));
};

