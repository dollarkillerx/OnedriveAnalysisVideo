const https = require('https');
let url = 'https://vipmail-my.sharepoint.cn/:v:/g/personal/dollarkiller_e1_tn/ETX644n2ap1GsZAovwE03pgBJay_mzDDf-75E9bYO3_LsQ?e=UWU8kX';

// url = 'https://vipmail-my.sharepoint.cn/personal/dollarkiller_e1_tn/_layouts/15/onedrive.aspx?id=%2Fpersonal%2Fdollarkiller%5Fe1%5Ftn%2FDocuments%2FRabbit%2FCrayonshinchan%2F%5BYGSUB%5D%5B1992%2E04%2E27%5D%5B003%5D%5BHDTV%5D%5BX264%2EMP4%5D%2Emp4&amp;parent=%2Fpersonal%2Fdollarkiller%5Fe1%5Ftn%2FDocuments%2FRabbit%2FCrayonshinchan&amp;originalPath=aHR0cHM6Ly92aXBtYWlsLW15LnNoYXJlcG9pbnQuY24vOnY6L2cvcGVyc29uYWwvZG9sbGFya2lsbGVyX2UxX3RuL0VUWDY0NG4yYXAxR3NaQW92d0UwM3BnQkpheV9tekREZi03NUU5YllPM19Mc1E_cnRpbWU9YmNPZkxFbFkxMGc'

https.get(url,function(res){

    res.on('cookie',function(cookie){
        console.log("====Start====")
        console.log(cookie)
        console.log("====End====")
    })

    res.on('data',function(chunk){
        console.log(chunk.toString());
    });

    res.on('end',function(){
        console.log('数据包传输完毕');
    })
})