function myAjaxUpload(files, obj) {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/upload');

    // send data to server
    var formData = new FormData();
    files.forEach(f => {
        formData.append("filepond", f.file)
    });
    formData.append("action", obj.action); // merge|img2pdf...

    // additional data
    formData.append("cw",obj.cw?obj.cw:""); // rotate pdf
    formData.append("ccw",obj.ccw?obj.ccw:""); // rotate
    formData.append("degree",obj.degree?obj.degree:""); // rotate
    formData.append("format",obj.format?obj.format:""); // pdf2img
    formData.append("combine",obj.combine); // [img2pdf] combine multiple images to single pdf
    formData.append("pdf2oneimg",obj.pdf2oneimg); // [pdf2img]
    
    xhr.send(formData);
    xhr.onload = function() {
        if ((xhr.status >= 200 && xhr.status < 300) || xhr.status == 304) {
            console.log('upload success');
            // switch button status
            document.getElementById('spinner').style.display = 'none';
            document.getElementById('upload').hidden = false;

            var result = JSON.parse(xhr.responseText);
            var href = "http://" + window.location.host;
            window.location.href = href + "/download/" + result;
        }
    };
    xhr.onerror = function(e) {
        console.log('error', e)
    };
}

function notify(message){
    $.notify({
        icon: 'fa fa-info-circle',
        message: message,
    },{
        type: "info",
        allow_dismiss: true,
        // delay: 50000,
        placement:{
            from:"top",
            align: "center"
        },
        animate: {
			enter: "animate__animated animate__fadeInDown",
			exit: "animate__animated animate__fadeOutUp"
		}
    });
}
