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
    //res.status(200).json({response_type:'in_channel',text:'Would you like to delete these files?',attachments:[{fallback:'',color:'',pretext:'',author_name:'',author_link:'',author_icon:'',title:'',title_link:'',text:'James and the Giant Peach\t0 KiB\t1969-12-31 17:00:00 -0700 MST\nThe Hunt for Red October\t0 KiB\t1969-12-31 17:00:00 -0700 MST\nEnders Game\t0 KiB\t1969-12-31 17:00:00 -0700 MST\n',fields:'',image_url:'',thumb_url:'',footer:'',footer_icon:'',ts:'',actions:[{name:'yes,',text:'Yes',style:'',type:'button',value:'   ',confirm:{title:'',text:'Are you sure?',ok_text:'Deleted',dismiss_text:''}},{name:'no',text:'No',style:'',type:'button',value:'no',confirm:{title:'',text:'',ok_text:'',dismiss_text:''}}]}]})
};

