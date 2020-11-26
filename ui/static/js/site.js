function myAjaxUpload(files, action) {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/upload');

    // send data to server
    var formData = new FormData();
    files.forEach(f => {
        formData.append("filepond", f.file)
    });
    formData.append("action", action); // merge|img2pdf...

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
