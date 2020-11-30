FilePond.registerPlugin(FilePondPluginImagePreview);
const inputElement = document.querySelector('input[type="file"]');

const pond = FilePond.create(inputElement);

pond.setOptions({
    // server: '/upload',
    // add additional data 
    server:{
        url: 'http://' + window.location.host, // http://ip:port
        process:{
            url: '/upload',
            method: 'POST',
            ondata:(formData)=>{
                formData.append('action','img2pdf');
                return formData;
            }
        }
    },
    instantUpload: false, // 不自动上传；默认为自动。自动上传one by one ，每次一个，循环。
    labelIdle: '拖放文件或点击',
    labelFileProcessing: '转换中',
    labelFileProcessingComplete: '完成',
    labelTapToCancel: '点击取消',
    labelTapToUndo: ''
});

// finished processing a file
// click button in imagePreview 
pond.on('processfile', (error, file) => {
    if (error) {
        console.log('Oh no');
        return;
    }
    console.log('File processed', file); // success

    var fileName = JSON.parse(file.serverId); // from server

    var a = document.createElement("a");
    var url = "http://" + window.location.host + "/download/" + fileName;
    //var linkText = document.createTextNode("single file");
    //a.appendChild(linkText);
    //a.title = "this is a title of a"
    //a.classList.add("btn", "btn-primary")
    a.setAttribute('href', url);
    a.setAttribute('download', fileName);
    document.body.appendChild(a);
    a.click();
});

// manual upload
const uploadBtn = document.getElementById('upload');
var spinner = document.getElementById('spinner');

uploadBtn.addEventListener("click", function() {

    var files = pond.getFiles();
    if (files.length == 0) {
        notify('请选择图片！');
        return;
    }
    spinner.style.display = "block";
    this.hidden = true;

    var obj ={};
    obj.action = "img2pdf";
    
    myAjaxUpload(files, obj);
});
