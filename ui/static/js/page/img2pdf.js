FilePond.registerPlugin(FilePondPluginImagePreview);
const inputElement = document.querySelector('input[type="file"]');

const pond = FilePond.create(inputElement);

pond.setOptions({
    server: '/upload',
    instantUpload: false, // 不自动上传；默认为自动。自动上传one by one ，每次一个，循环。
    labelIdle: '拖放文件或点击',
    labelFileProcessing: '转换中',
    labelFileProcessingComplete: '完成',
    labelTapToCancel: '点击取消',
    labelTapToUndo: ''
});

// manual upload
const uploadBtn = document.getElementById('upload');
uploadBtn.addEventListener("click", function () {
    // 依次上传，不是文件列表。
    // pond.processFiles().then(files => {
    //     files.forEach(file => {
    //         console.log(file.filename);
    //     });
    // });

    var files = pond.getFiles();
    var formData = new FormData();
    files.forEach(f => {
        formData.append("filepond", f.file)
    });

    // https://segmentfault.com/a/1190000004322487
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/upload");
    xhr.send(formData);
    xhr.onload = function () {
        if ((xhr.status >= 200 && xhr.status < 300) || xhr.status == 304) {
            console.log('upload success');
            // go to download page
           // window.location.href="./download/"
         

        }
    };
    xhr.onerror = function (e) {
        console.log('error', e)
    };
})

// 上传成功或失败 (trigger when click single file upload)
pond.on('processfile', (error, file) => {
    if (error) {
        console.log('Oh no');
        return;
    }
    console.log('File processed', file); // success
    var fileName = file.filename.split(".")[0] + ".pdf";
    var a = document.createElement("a");
    var url = "http://localhost:1234/upload/";
    var linkText = document.createTextNode("single file");
    a.appendChild(linkText);
    a.title = "this is a title of a"
    a.classList.add("btn", "btn-primary")
    a.setAttribute('href', url + fileName);
    a.setAttribute('download', fileName);
    document.body.appendChild(a);
    // a.click();
});
