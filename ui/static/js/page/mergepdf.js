const inputElement = document.querySelector('input[type="file"]');

const pond = FilePond.create(inputElement);

pond.setOptions({
    server: '/mergepdf',
    instantUpload: false, // 不自动上传；默认为自动。
    allowFileTypeValidation: false,
    acceptedFileTypes: "pdf",
    labelIdle: '拖放文件或点击',
    labelFileProcessing: '转换中',
    labelFileProcessingComplete: '完成',
    labelTapToCancel: '点击取消',
    labelTapToUndo: ''
});

// 上传成功或失败
pond.on('processfile', (error, file) => {
    if (error) {
        console.log('Oh no');
        return;
    }
    console.log('File processed', file); // success
    var fileName = file.filename.split(".")[0] + ".pdf";
    var a = document.createElement("a");
    var url = "http://localhost:8001/upload/";
    // var linkText = document.createTextNode("chaolianjie");
    // a.appendChild(linkText);
    // a.title = "this is a title of a"
    // a.classList.add("btn", "btn-primary")
    a.setAttribute('href', url + fileName);
    a.setAttribute('download', fileName);
    document.body.appendChild(a);
    a.click();
});